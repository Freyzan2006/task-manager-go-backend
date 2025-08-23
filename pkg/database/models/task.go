package models

import (
	"time"
)

type Task struct {
    ID          uint      `gorm:"primaryKey"`
    UserID      uint      `gorm:"not null;index"`
    Title       string    `gorm:"not null"`
    Description string
    Status      string    `gorm:"type:text;default:'todo'"`
    Priority    int
    CreatedAt   time.Time `gorm:"autoCreateTime"`
    Tags        []Tag     `gorm:"many2many:task_tags;joinForeignKey:TaskID;joinReferences:TagID"`
}
