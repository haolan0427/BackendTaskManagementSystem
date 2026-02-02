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