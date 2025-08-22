package models 

type TagModel struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"unique;not null"`
    Tasks []TaskModel `gorm:"many2many:task_tags;"`
}