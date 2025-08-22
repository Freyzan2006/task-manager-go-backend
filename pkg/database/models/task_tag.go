package models

type TaskTag struct {
    TaskID uint `gorm:"primaryKey"`
    TagID  uint `gorm:"primaryKey"`
}