package models

type User struct {
    ID           uint        `gorm:"primaryKey"`
    Email        string      `gorm:"unique;not null"`
    Username     string      `gorm:"unique;not null"`
    PasswordHash string      `gorm:"not null"`
    Tasks        []Task      `gorm:"foreignKey:UserID"`
}