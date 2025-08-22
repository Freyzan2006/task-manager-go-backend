package models

import (
	"time"
)

type TaskModel struct {
    ID          uint      `gorm:"primaryKey"`
    UserID      uint      `gorm:"not null;index"`
    Title       string    `gorm:"not null"`
    Description string
    Status      string    `gorm:"type:text;default:'todo'"`
    Priority    int
    CreatedAt   time.Time `gorm:"autoCreateTime"`
    Tags        []TagModel `gorm:"many2many:task_tags;"`
}
