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
	// r.db.AutoMigrate(...) 是 GORM 提供的超级强大的功能：
	// 它会根据 model.Task 结构体自动创建表（如果表不存在）。
	// 如果你修改了 model.Task 的字段（加字段、改类型），再次运行时它会自动更新表结构（加列、改列等）。
	return r.db.AutoMigrate(&model.Task{})
}

func (r *TaskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}
