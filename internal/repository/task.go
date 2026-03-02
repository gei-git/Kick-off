package repository

import (
	"github.com/gei-git/Kick-off/internal/model"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) AutoMigrate() error {
	return r.db.AutoMigrate(&model.Task{})
}

func (r *TaskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Preload("User").Find(&tasks).Error // 后续会加 User 关联
	return tasks, err
}

func (r *TaskRepository) FindAllWithFilter(page, limit int, priority string) ([]model.Task, int64, error) {
	var tasks []model.Task
	var total int64

	query := r.db.Model(&model.Task{})

	if priority != "" {
		query = query.Where("priority = ?", priority)
	}

	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Find(&tasks).Error

	return tasks, total, err
}
