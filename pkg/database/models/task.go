package models

import "time"

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey" example:"1"`
	UserID      uint      `json:"user_id" gorm:"not null;index" example:"42"`
	Title       string    `json:"title" gorm:"not null" example:"Купить хлеб"`
	Description string    `json:"description,omitempty" example:"Сходить в магазин за хлебом"`
	Status      string    `json:"status" gorm:"type:text;default:'todo'" example:"todo"`
	Priority    int       `json:"priority" example:"1"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" example:"2025-09-05T14:25:00Z"`
	Tags        []Tag     `json:"tags" gorm:"many2many:task_tags;joinForeignKey:TaskID;joinReferences:TagID"`
}

