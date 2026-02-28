package service

import (
	"fmt"

	"github.com/gei-git/Kick-off/internal/model"
	"github.com/gei-git/Kick-off/internal/repository"
	"gorm.io/gorm"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{repo: repository.NewTaskRepository(db)}
}

func (s *TaskService) CreateTask(task *model.Task) (*model.Task, error) {
	if task.Title == "" {
		return nil, fmt.Errorf("任务标题不能为空")
	}
	if err := s.repo.Create(task); err != nil {
		return nil, fmt.Errorf("创建任务失败: %w", err)
	}

	return task, nil
}

func (s *TaskService) ListTask() ([]model.Task, error) {
	return s.repo.FindAll()
}
