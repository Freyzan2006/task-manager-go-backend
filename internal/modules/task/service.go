package task

import "task-manager/pkg/database/models"


type Service interface {
	GetTasks() ([]models.Task, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetTasks() ([]models.Task, error) {
	return s.repo.FindAll()
}
