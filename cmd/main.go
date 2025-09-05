package main

import (
	_ "task-manager/docs"

	"task-manager/internal/core"
	"task-manager/internal/core/logger"
	"task-manager/internal/server"
)

// @title           Task manager
// @version         1.0
// @description     Микросервис для управления задачами
// @host            localhost:${port}
// @BasePath        /api/v1
func main() {
	// Загружаем конфиг и сразу создаём объект базы внутри
	cfg := core.NewConfig()
	logger := logger.New("app.log", "debug", "task-manager-go-backend", cfg.Mode)
	logger.Info("✅ Config loaded successfully")

	// Создаём и запускаем сервер
	srv := server.New(cfg, logger)
	if err := srv.Run(); err != nil {
		logger.Fatalf("❌ failed to run server: %v", err)
	}
}

// task-manager-go-backend/internal/server/engine.go:49