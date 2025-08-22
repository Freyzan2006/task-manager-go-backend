package models

type UserModel struct {
    ID           uint        `gorm:"primaryKey"`
    Username     string      `gorm:"unique;not null"`
    PasswordHash string      `gorm:"not null"`
    Tasks        []TaskModel `gorm:"foreignKey:UserID"`
}