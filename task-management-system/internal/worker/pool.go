package worker

import (
    "context"
    "log"
    "sync"
)

type Task func(ctx context.Context) error

type Worker struct {
    id       int
    taskChan chan Task
    quit     chan bool
    wg       *sync.WaitGroup
}

func NewWorker(id int, taskChan chan Task, wg *sync.WaitGroup) *Worker {
    return &Worker{
        id:       id,
        taskChan: taskChan,
        quit:     make(chan bool),
        wg:       wg,
    }
}

func (w *Worker) Start(ctx context.Context) {
    w.wg.Add(1)
    go func() {
        defer w.wg.Done()
        for {
            select {
            case task := <-w.taskChan:
                if err := task(ctx); err != nil {
                    log.Printf("Worker %d: task failed: %v", w.id, err)
                }
            case <-w.quit:
                return
            case <-ctx.Done():
                return
            }
        }
    }()
}

func (w *Worker) Stop() {
    close(w.quit)
}

type Pool struct {
    workers  []*Worker
    taskChan chan Task
    ctx      context.Context
    cancel   context.CancelFunc
    wg       sync.WaitGroup
}

func NewPool(workerCount int) *Pool {
    ctx, cancel := context.WithCancel(context.Background())
    taskChan := make(chan Task, 100)
    
    pool := &Pool{
        workers:  make([]*Worker, workerCount),
        taskChan: taskChan,
        ctx:      ctx,
        cancel:   cancel,
    }
    
    for i := 0; i < workerCount; i++ {
        pool.workers[i] = NewWorker(i, taskChan, &pool.wg)
        pool.workers[i].Start(ctx)
    }
    
    return pool
}

func (p *Pool) Submit(task Task) {
    p.taskChan <- task
}

func (p *Pool) Shutdown() {
    p.cancel()
    for _, worker := range p.workers {
        worker.Stop()
    }
    p.wg.Wait()
    close(p.taskChan)
}