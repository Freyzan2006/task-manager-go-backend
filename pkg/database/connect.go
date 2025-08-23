package database

import (
    "log"
)

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

import (
    "task-manager/pkg/database/models"
)

type DB struct {
    database *gorm.DB
    name     string
}

func NewDB(name string) *DB {
    return &DB{
        database: nil,
        name:     name,
    }
}

func (db *DB) ConnectSQLite() error {
    gdb, err := gorm.Open(sqlite.Open(db.name), &gorm.Config{})
    if err != nil {
        return err
    }
    db.database = gdb
    return nil
}

func (db *DB) GetDB() *gorm.DB {
    return db.database
}

func (db *DB) AutoMigrate() error {
    err := db.database.AutoMigrate(
        &models.User{},
        &models.Task{},
        &models.Tag{},
        &models.TaskTag{},
    )
    if err != nil {
        return err
    }

    
    return db.Seed()
}

func (db *DB) Seed() error {
    // Проверяем пользователей
    var count int64
    db.database.Model(&models.User{}).Count(&count)
    if count == 0 {
        users := []models.User{
            {Username: "Admin", Email: "admin@example.com", PasswordHash: "admin"},
            {Username: "Demo User", Email: "demo@example.com", PasswordHash: "demo"},
        }
        if err := db.database.Create(&users).Error; err != nil {
            return err
        }
    }

    // Проверяем теги
    db.database.Model(&models.Tag{}).Count(&count)
    if count == 0 {
        tags := []models.Tag{
            {Name: "Important"},
            {Name: "Urgent"},
            {Name: "Optional"},
        }
        if err := db.database.Create(&tags).Error; err != nil {
            return err
        }
    }

    // Проверяем задачи 
    db.database.Model(&models.Task{}).Count(&count)
    if count == 0 {
        var tags []models.Tag
        db.database.Find(&tags)

        tasks := []models.Task{
            {
                UserID: 1,
                Title: "Task 1",
                Description: "Description 1",
                Status: "todo",
                Priority: 0,
                Tags: []models.Tag{tags[0]}, // Important
            },
            {
                UserID: 1,
                Title: "Task 2",
                Description: "Description 2",
                Status: "in_progress",
                Priority: 1,
                Tags: []models.Tag{tags[1]}, // Urgent
            },
            {
                UserID: 1,
                Title: "Task 3",
                Description: "Description 3",
                Status: "done",
                Priority: 2,
                Tags: []models.Tag{tags[2]}, // Optional
            },
        }

        for i := range tasks {
            if err := db.database.Create(&tasks[i]).Error; err != nil {
                return err
            }
            for _, tag := range tasks[i].Tags {
                if err := db.database.Model(&tasks[i]).Association("Tags").Append(&tag); err != nil {
                    return err
                }
            }
        }

        // 🔑 теперь тянем с Preload, иначе в структуре будет пусто
        var checkTasks []models.Task
        if err := db.database.Preload("Tags").Find(&checkTasks).Error; err != nil {
            return err
        }
        log.Printf("✅ Seeded tasks with tags: %+v\n", checkTasks)
    }

    return nil
}


// -- users
// CREATE TABLE users (
//     id INTEGER PRIMARY KEY AUTOINCREMENT,
//     username TEXT NOT NULL UNIQUE,
//     password_hash TEXT NOT NULL
// );

// -- tasks
// CREATE TABLE tasks (
//     id INTEGER PRIMARY KEY AUTOINCREMENT,
//     user_id INTEGER NOT NULL,
//     title TEXT NOT NULL,
//     description TEXT,
//     status TEXT CHECK(status IN ('todo','in_progress','done')) DEFAULT 'todo',
//     priority INTEGER DEFAULT 0,
//     created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
//     FOREIGN KEY(user_id) REFERENCES users(id)
// );

// -- tags
// CREATE TABLE tags (
//     id INTEGER PRIMARY KEY AUTOINCREMENT,
//     name TEXT NOT NULL UNIQUE
// );

// -- tasks_tags (many-to-many)
// CREATE TABLE task_tags (
//     task_id INTEGER,
//     tag_id INTEGER,
//     PRIMARY KEY(task_id, tag_id),
//     FOREIGN KEY(task_id) REFERENCES tasks(id),
//     FOREIGN KEY(tag_id) REFERENCES tags(id)
// );
