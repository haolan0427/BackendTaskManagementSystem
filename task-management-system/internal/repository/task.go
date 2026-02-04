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
