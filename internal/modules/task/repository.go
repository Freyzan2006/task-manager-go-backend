package task

import "gorm.io/gorm"

import (
	"task-manager/pkg/database/models"
)

type Repository interface {
	FindAll() ([]models.Task, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Preload("Tags").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
