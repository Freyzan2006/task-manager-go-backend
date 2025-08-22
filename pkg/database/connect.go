package database

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

func (db *DB) AutoMigrate() error {
    return db.database.AutoMigrate(
        &models.UserModel{},
        &models.TaskModel{},
        &models.TagModel{},
        &models.TaskTag{},
    )
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
