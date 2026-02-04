# Go后端项目实战：任务管理系统

## 项目概述

这是一个难度中上的Go后端项目，旨在帮助你巩固已学知识并学习进阶概念。

### 项目特点

- **70%已学知识**：基础语法、结构体、接口、HTTP服务器、数据库操作、错误处理
- **30%新知识**：JWT认证、Redis缓存、Worker Pool、中间件、配置管理、单元测试、Docker

---

## 项目架构

```
task-management-system/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handler/
│   │   ├── auth.go
│   │   ├── task.go
│   │   └── user.go
│   ├── middleware/
│   │   ├── auth.go
│   │   ├── cors.go
│   │   ├── logger.go
│   │   └── ratelimit.go
│   ├── model/
│   │   ├── task.go
│   │   └── user.go
│   ├── repository/
│   │   ├── task.go
│   │   └── user.go
│   ├── service/
│   │   ├── auth.go
│   │   ├── task.go
│   │   └── user.go
│   └── worker/
│       └── pool.go
├── pkg/
│   ├── cache/
│   │   └── redis.go
│   ├── database/
│   │   └── mysql.go
│   └── jwt/
│       └── jwt.go
├── configs/
│   └── config.yaml
├── tests/
│   ├── handler_test.go
│   └── service_test.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum
```

---

## 核心功能

### 1. 用户认证系统

- 用户注册和登录
- JWT Token生成和验证
- 密码加密存储
- Token刷新机制

### 2. 任务管理

- 任务的CRUD操作
- 任务状态管理（待办、进行中、已完成）
- 任务优先级设置
- 任务分配给用户
- 任务搜索和过滤

### 3. 异步任务处理

- Worker Pool模式处理后台任务
- 任务队列管理
- 任务执行状态跟踪
- 失败重试机制

### 4. 缓存系统

- Redis缓存热点数据
- 缓存失效策略
- 缓存预热机制

### 5. API限流

- 基于IP的限流
- 基于用户的限流
- 限流算法实现

---

## 技术栈

### 核心技术

- **语言**: Go 1.21+
- **Web框架**: Gin
- **数据库**: MySQL
- **ORM**: GORM
- **缓存**: Redis
- **认证**: JWT

### 工具库

- **配置管理**: Viper
- **日志**: Zap
- **验证**: Go-playground/validator
- **测试**: Testify
- **Docker**: Docker & Docker Compose

---

## 实现步骤

### 阶段一：项目初始化（基础）

1. 创建项目结构
2. 初始化Go模块
3. 配置Docker环境
4. 数据库连接和模型定义

### 阶段二：用户认证（新知识30%）

1. 用户注册和登录API
2. JWT Token生成和验证
3. 中间件实现（认证、CORS、日志）
4. 密码加密和验证

### 阶段三：任务管理（基础70%）

1. 任务CRUD API
2. 任务状态和优先级管理
3. 任务搜索和过滤
4. 单元测试编写

### 阶段四：高级功能（新知识30%）

1. Worker Pool实现
2. Redis缓存集成
3. API限流中间件
4. 异步任务处理

### 阶段五：优化和部署（新知识30%）

1. 性能优化
2. 错误处理完善
3. Docker容器化
4. 文档编写

---

## 详细实现指南

### 1. 项目初始化

#### 创建项目结构 ✅

```bash
mkdir task-management-system
cd task-management-system
go mod init task-management-system

# 创建目录结构
mkdir -p cmd/server internal/{config,handler,middleware,model,repository,service,worker} pkg/{cache,database,jwt} configs tests
```

#### 配置文件 ✅

```yaml
# configs/config.yaml
server:
  port: 8080
  mode: debug

database:
  host: localhost
  port: 3306
  user: root
  password: root
  dbname: taskdb
  charset: utf8mb4

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0

jwt:
  secret: your-secret-key
  expire: 24h

rate_limit:
  requests_per_minute: 60
```

### 2. 数据库模型

#### 用户模型 ✅

```go
package model

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        uint           `gorm:"primarykey" json:"id"`
    Username  string         `gorm:"uniqueIndex;not null" json:"username"`
    Email     string         `gorm:"uniqueIndex;not null" json:"email"`
    Password  string         `gorm:"not null" json:"-"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
    Tasks     []Task         `gorm:"foreignKey:UserID" json:"tasks,omitempty"`
}
```

#### 任务模型 ✅

```go
package model

import (
    "time"
    "gorm.io/gorm"
)

type Task struct {
    ID          uint           `gorm:"primarykey" json:"id"`
    Title       string         `gorm:"not null" json:"title"`
    Description string         `json:"description"`
    Status      string         `gorm:"default:'pending'" json:"status"`
    Priority    string         `gorm:"default:'medium'" json:"priority"`
    UserID      uint           `gorm:"not null" json:"user_id"`
    User        User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
    DueDate     *time.Time     `json:"due_date"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type TaskStatus string

const (
    StatusPending    TaskStatus = "pending"
    StatusInProgress TaskStatus = "in_progress"
    StatusCompleted  TaskStatus = "completed"
)

type TaskPriority string

const (
    PriorityLow    TaskPriority = "low"
    PriorityMedium TaskPriority = "medium"
    PriorityHigh   TaskPriority = "high"
)
```

### 3. 配置管理 ✅

```go
package config

import (
    "fmt"
    "github.com/spf13/viper"
)

type Config struct {
    Server    ServerConfig    `mapstructure:"server"`
    Database  DatabaseConfig  `mapstructure:"database"`
    Redis     RedisConfig     `mapstructure:"redis"`
    JWT       JWTConfig       `mapstructure:"jwt"`
    RateLimit RateLimitConfig `mapstructure:"rate_limit"`
}

type ServerConfig struct {
    Port string `mapstructure:"port"`
    Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     string `mapstructure:"port"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    DBName   string `mapstructure:"dbname"`
    Charset  string `mapstructure:"charset"`
}

type RedisConfig struct {
    Host     string `mapstructure:"host"`
    Port     string `mapstructure:"port"`
    Password string `mapstructure:"password"`
    DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
    Secret string        `mapstructure:"secret"`
    Expire time.Duration `mapstructure:"expire"`
}

type RateLimitConfig struct {
    RequestsPerMinute int `mapstructure:"requests_per_minute"`
}

func LoadConfig(path string) (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(path)
    
    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }
    
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, fmt.Errorf("failed to unmarshal config: %w", err)
    }
    
    return &config, nil
}
```

### 4. 数据库连接 ✅

```go
package database

import (
    "fmt"
    "task-management-system/internal/config"
    "task-management-system/internal/model"
    
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func NewMySQLDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
        cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Charset,
    )
    
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }
    
    if err := db.AutoMigrate(&model.User{}, &model.Task{}); err != nil {
        return nil, fmt.Errorf("failed to migrate database: %w", err)
    }
    
    return db, nil
}
```

### 5. Redis缓存 ✅

```go
package cache

import (
    "context"
    "encoding/json"
    "fmt"
    "task-management-system/internal/config"
    "time"
    
    "github.com/redis/go-redis/v9"
)

type RedisClient struct {
    client *redis.Client
}

func NewRedisClient(cfg *config.RedisConfig) *RedisClient {
    return &RedisClient{
        client: redis.NewClient(&redis.Options{
            Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
            Password: cfg.Password,
            DB:       cfg.DB,
        }),
    }
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
    jsonValue, err := json.Marshal(value)
    if err != nil {
        return err
    }
    return r.client.Set(ctx, key, jsonValue, expiration).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string, dest interface{}) error {
    val, err := r.client.Get(ctx, key).Result()
    if err != nil {
        return err
    }
    return json.Unmarshal([]byte(val), dest)
}

func (r *RedisClient) Delete(ctx context.Context, key string) error {
    return r.client.Del(ctx, key).Err()
}

func (r *RedisClient) Close() error {
    return r.client.Close()
}
```

### 6. JWT工具 ✅ 

```go
package jwt

import (
    "errors"
    "task-management-system/internal/config"
    "time"
    
    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID uint   `json:"user_id"`
    Email  string `json:"email"`
    jwt.RegisteredClaims
}

type JWTManager struct {
    secret string
    expire time.Duration
}

func NewJWTManager(cfg *config.JWTConfig) *JWTManager {
    return &JWTManager{
        secret: cfg.Secret,
        expire: cfg.Expire,
    }
}

func (j *JWTManager) GenerateToken(userID uint, email string) (string, error) {
    claims := Claims{
        UserID: userID,
        Email:  email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expire)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(j.secret))
}

func (j *JWTManager) ValidateToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(j.secret), nil
    })
    
    if err != nil {
        return nil, err
    }
    
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, errors.New("invalid token")
}
```

### 7. 中间件

#### 认证中间件 ✅ 

```go
package middleware

import (
    "net/http"
    "strings"
    "task-management-system/pkg/jwt"
    
    "github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtManager *jwt.JWTManager) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }
        
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := jwtManager.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        
        c.Set("user_id", claims.UserID)
        c.Set("email", claims.Email)
        c.Next()
    }
}
```

#### 限流中间件 ✅

```go
package middleware

import (
    "net/http"
    "sync"
    "time"
    
    "github.com/gin-gonic/gin"
)

type RateLimiter struct {
    requests map[string][]time.Time
    mu       sync.Mutex
    limit    int
    window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
    return &RateLimiter{
        requests: make(map[string][]time.Time),
        limit:    limit,
        window:   window,
    }
}

func (r *RateLimiter) Allow(ip string) bool {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    now := time.Now()
    requests := r.requests[ip]
    
    var validRequests []time.Time
    for _, req := range requests {
        if now.Sub(req) < r.window {
            validRequests = append(validRequests, req)
        }
    }
    
    r.requests[ip] = validRequests
    
    if len(validRequests) >= r.limit {
        return false
    }
    
    r.requests[ip] = append(validRequests, now)
    return true
}

func RateLimitMiddleware(limiter *RateLimiter) gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP()
        if !limiter.Allow(ip) {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
            c.Abort()
            return
        }
        c.Next()
    }
}
```

### 8. Worker Pool ✅

```go
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
```

### 9. Repository实现

#### 用户Repository ✅

```go
package repository

import (
	"task-management-system/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

```

#### 任务Repository ✅

```go
package repository

import (
	"task-management-system/internal/model"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) FindByID(id uint) (*model.Task, error) {
	var task model.Task
	err := r.db.Preload("User").First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) FindByUserID(userID uint) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) FindByUserIDAndStatus(userID uint, status string) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Where("user_id = ? AND status = ?", userID, status).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) FindByUserIDAndPriority(userID uint, priority string) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Where("user_id = ? AND priority = ?", userID, priority).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) FindByUserIDWithFilters(userID uint, status, priority string) ([]model.Task, error) {
	var tasks []model.Task
	query := r.db.Where("user_id = ?", userID)

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if priority != "" {
		query = query.Where("priority = ?", priority)
	}

	err := query.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepository) Delete(id uint) error {
	return r.db.Delete(&model.Task{}, id).Error
}

func (r *TaskRepository) CountByUserID(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&model.Task{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

```

### 10. Service层

#### 用户Service ✅

```go
package service

import (
    "errors"
    "task-management-system/internal/model"
    "task-management-system/internal/repository"
    "task-management-system/pkg/jwt"
    
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    userRepo *repository.UserRepository
    jwtMgr   *jwt.JWTManager
}

func NewUserService(userRepo *repository.UserRepository, jwtMgr *jwt.JWTManager) *UserService {
    return &UserService{
        userRepo: userRepo,
        jwtMgr:   jwtMgr,
    }
}

func (s *UserService) Register(user *model.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    
    return s.userRepo.Create(user)
}

func (s *UserService) Login(email, password string) (string, *model.User, error) {
    user, err := s.userRepo.GetByEmail(email)
    if err != nil {
        return "", nil, errors.New("invalid credentials")
    }
    
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", nil, errors.New("invalid credentials")
    }
    
    token, err := s.jwtMgr.GenerateToken(user.ID, user.Email)
    if err != nil {
        return "", nil, err
    }
    
    return token, user, nil
}
```

#### 任务Service ✅

```go
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
```

### 11. Handler实现

#### 用户Handler

```go
package handler

import (
    "net/http"
    "task-management-system/internal/model"
    "task-management-system/internal/service"
    
    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

type RegisterRequest struct {
    Username string `json:"username" binding:"required,min=3,max=20"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func (h *UserHandler) Register(c *gin.Context) {
    var req RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    user := &model.User{
        Username: req.Username,
        Email:    req.Email,
        Password: req.Password,
    }
    
    if err := h.userService.Register(user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{
        "message": "User registered successfully",
        "user": gin.H{
            "id":       user.ID,
            "username": user.Username,
            "email":    user.Email,
        },
    })
}

func (h *UserHandler) Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    token, user, err := h.userService.Login(req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "token": token,
        "user": gin.H{
            "id":       user.ID,
            "username": user.Username,
            "email":    user.Email,
        },
    })
}
```

#### 任务Handler

```go
package handler

import (
    "net/http"
    "strconv"
    "task-management-system/internal/model"
    "task-management-system/internal/service"
    
    "github.com/gin-gonic/gin"
)

type TaskHandler struct {
    taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
    return &TaskHandler{taskService: taskService}
}

type CreateTaskRequest struct {
    Title       string     `json:"title" binding:"required"`
    Description string     `json:"description"`
    Priority    string     `json:"priority" binding:"required,oneof=low medium high"`
    DueDate     *string    `json:"due_date"`
}

type UpdateTaskRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status" binding:"omitempty,oneof=pending in_progress completed"`
    Priority    string `json:"priority" binding:"omitempty,oneof=low medium high"`
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
    userID := c.GetUint("user_id")
    
    var req CreateTaskRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    task := &model.Task{
        Title:       req.Title,
        Description: req.Description,
        Priority:    req.Priority,
        Status:      string(model.StatusPending),
        UserID:      userID,
    }
    
    if err := h.taskService.CreateTask(task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
    userID := c.GetUint("user_id")
    status := c.Query("status")
    priority := c.Query("priority")
    
    tasks, err := h.taskService.GetTasks(userID, status, priority)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTask(c *gin.Context) {
    taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    
    task, err := h.taskService.GetTaskByID(uint(taskID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
    taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    
    var req UpdateTaskRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    task, err := h.taskService.UpdateTask(uint(taskID), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
    taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    
    if err := h.taskService.DeleteTask(uint(taskID)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
```

### 12. 主程序

```go
package main

import (
    "log"
    "task-management-system/internal/config"
    "task-management-system/internal/handler"
    "task-management-system/internal/middleware"
    "task-management-system/internal/worker"
    "task-management-system/pkg/cache"
    "task-management-system/pkg/database"
    "task-management-system/pkg/jwt"
    
    "github.com/gin-gonic/gin"
)

func main() {
    cfg, err := config.LoadConfig("./configs")
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }
    
    db, err := database.NewMySQLDB(&cfg.Database)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    
    redisClient := cache.NewRedisClient(&cfg.Redis)
    defer redisClient.Close()
    
    jwtManager := jwt.NewJWTManager(&cfg.JWT)
    
    workerPool := worker.NewPool(5)
    defer workerPool.Shutdown()
    
    rateLimiter := middleware.NewRateLimiter(cfg.RateLimit.RequestsPerMinute, time.Minute)
    
    router := gin.Default()
    
    router.Use(middleware.LoggerMiddleware())
    router.Use(middleware.CORSMiddleware())
    router.Use(middleware.RateLimitMiddleware(rateLimiter))
    
    userRepo := repository.NewUserRepository(db)
    taskRepo := repository.NewTaskRepository(db)
    
    userService := service.NewUserService(userRepo, jwtManager)
    taskService := service.NewTaskService(taskRepo, redisClient, workerPool)
    
    userHandler := handler.NewUserHandler(userService)
    taskHandler := handler.NewTaskHandler(taskService)
    
    api := router.Group("/api/v1")
    {
        auth := api.Group("/auth")
        {
            auth.POST("/register", userHandler.Register)
            auth.POST("/login", userHandler.Login)
        }
        
        tasks := api.Group("/tasks")
        tasks.Use(middleware.AuthMiddleware(jwtManager))
        {
            tasks.POST("", taskHandler.CreateTask)
            tasks.GET("", taskHandler.GetTasks)
            tasks.GET("/:id", taskHandler.GetTask)
            tasks.PUT("/:id", taskHandler.UpdateTask)
            tasks.DELETE("/:id", taskHandler.DeleteTask)
        }
    }
    
    log.Printf("Server starting on port %s", cfg.Server.Port)
    if err := router.Run(":" + cfg.Server.Port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
```

### 13. Docker配置

#### Dockerfile

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
```

#### docker-compose.yml

```yaml
version: '3.8'

services:
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: taskdb
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

  app:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - mysql
      - redis
    environment:
      - DATABASE_HOST=mysql
      - REDIS_HOST=redis
    volumes:
      - ./configs:/app/configs

volumes:
  mysql_data:
  redis_data:
```

---

## 学习目标

### 已学知识巩固（70%）

1. Go基础语法和数据结构
2. 结构体和方法
3. 接口和多态
4. 错误处理
5. HTTP服务器开发
6. 数据库操作
7. 并发编程基础

### 新知识学习（30%）

1. JWT认证和授权
2. Redis缓存集成
3. Worker Pool模式
4. 中间件设计
5. 配置管理
6. 单元测试和集成测试
7. Docker容器化

---

## API文档

### 认证接口

#### 注册用户

```
POST /api/v1/auth/register
Content-Type: application/json

{
  "username": "alice",
  "email": "alice@example.com",
  "password": "password123"
}
```

#### 用户登录

```
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "alice@example.com",
  "password": "password123"
}
```

### 任务接口

#### 创建任务

```
POST /api/v1/tasks
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "完成项目文档",
  "description": "编写项目的技术文档",
  "priority": "high",
  "due_date": "2026-02-01T00:00:00Z"
}
```

#### 获取任务列表

```
GET /api/v1/tasks?status=pending&priority=high
Authorization: Bearer <token>
```

#### 获取单个任务

```
GET /api/v1/tasks/1
Authorization: Bearer <token>
```

#### 更新任务

```
PUT /api/v1/tasks/1
Authorization: Bearer <token>
Content-Type: application/json

{
  "status": "in_progress"
}
```

#### 删除任务

```
DELETE /api/v1/tasks/1
Authorization: Bearer <token>
```

---

## 测试指南

### 单元测试示例

```go
package service

import (
    "testing"
    "task-management-system/internal/model"
    
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) Create(user *model.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *MockUserRepository) GetByEmail(email string) (*model.User, error) {
    args := m.Called(email)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*model.User), args.Error(1)
}

func TestUserService_Register(t *testing.T) {
    mockRepo := new(MockUserRepository)
    jwtMgr := jwt.NewJWTManager(&config.JWTConfig{Secret: "test", Expire: time.Hour})
    service := NewUserService(mockRepo, jwtMgr)
    
    user := &model.User{
        Username: "testuser",
        Email:    "test@example.com",
        Password: "password123",
    }
    
    mockRepo.On("Create", mock.AnythingOfType("*model.User")).Return(nil)
    
    err := service.Register(user)
    
    assert.NoError(t, err)
    assert.NotEqual(t, "password123", user.Password)
    mockRepo.AssertExpectations(t)
}
```

---

## 项目扩展建议

完成基础功能后，可以考虑以下扩展：

1. **WebSocket支持**：实时任务更新通知
2. **文件上传**：任务附件功能
3. **标签系统**：任务分类和标签
4. **评论系统**：任务讨论功能
5. **统计报表**：任务完成率统计
6. **邮件通知**：任务提醒功能
7. **OAuth集成**：第三方登录
8. **GraphQL API**：替代REST API

---

## 总结

这个项目涵盖了Go后端开发的核心概念，难度适中，既能巩固已学知识，又能学习新的技术栈。通过完成这个项目，你将掌握：

1. 完整的Go后端项目架构
2. RESTful API设计和实现
3. 数据库设计和ORM使用
4. 认证和授权机制
5. 缓存和性能优化
6. 并发编程实践
7. 测试和部署

祝你学习愉快！