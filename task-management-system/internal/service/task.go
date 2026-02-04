package service

import (
    "errors"
    "task-management-system/internal/model"
    "task-management-system/internal/repository"
    "task-management-system/pkg/cache"
    
    "context"
    "time"
)

type TaskService struct {
    taskRepo *repository.TaskRepository
    cache    *cache.RedisClient
    pool     *worker.Pool
}

func NewTaskService(taskRepo *repository.TaskRepository, cache *cache.RedisClient, pool *worker.Pool) *TaskService {
    return &TaskService{
        taskRepo: taskRepo,
        cache:    cache,
        pool:     pool,
    }
}

func (s *TaskService) CreateTask(task *model.Task) error {
    if err := s.taskRepo.Create(task); err != nil {
        return err
    }
    
    s.pool.Submit(func(ctx context.Context) error {
        return s.cache.Set(ctx, taskCacheKey(task.ID), task, 5*time.Minute)
    })
    
    return nil
}

func (s *TaskService) GetTasks(userID uint, status, priority string) ([]model.Task, error) {
    return s.taskRepo.GetByUserIDWithFilters(userID, status, priority)
}

func (s *TaskService) GetTaskByID(id uint) (*model.Task, error) {
    var task model.Task
    err := s.cache.Get(context.Background(), taskCacheKey(id), &task)
    if err == nil {
        return &task, nil
    }
    
    task, err = s.taskRepo.GetByID(id)
    if err != nil {
        return nil, err
    }
    
    s.pool.Submit(func(ctx context.Context) error {
        return s.cache.Set(ctx, taskCacheKey(task.ID), task, 5*time.Minute)
    })
    
    return &task, nil
}

func (s *TaskService) UpdateTask(id uint, req UpdateTaskRequest) (*model.Task, error) {
    task, err := s.taskRepo.GetByID(id)
    if err != nil {
        return nil, errors.New("task not found")
    }
    
    if req.Title != "" {
        task.Title = req.Title
    }
    if req.Description != "" {
        task.Description = req.Description
    }
    if req.Status != "" {
        task.Status = req.Status
    }
    if req.Priority != "" {
        task.Priority = req.Priority
    }
    
    if err := s.taskRepo.Update(task); err != nil {
        return nil, err
    }
    
    s.pool.Submit(func(ctx context.Context) error {
        return s.cache.Delete(ctx, taskCacheKey(id))
    })
    
    return task, nil
}

func (s *TaskService) DeleteTask(id uint) error {
    if err := s.taskRepo.Delete(id); err != nil {
        return err
    }
    
    s.pool.Submit(func(ctx context.Context) error {
        return s.cache.Delete(ctx, taskCacheKey(id))
    })
    
    return nil
}

func taskCacheKey(id uint) string {
    return fmt.Sprintf("task:%d", id)
}