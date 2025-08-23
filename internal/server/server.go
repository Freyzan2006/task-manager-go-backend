package server

import (
	"log"
)

import (
	"task-manager/internal/modules/task"
)

import "github.com/gin-gonic/gin"

type Server interface {
    Run()
}

type server struct {
	cfg *ConfigServer
} 

func NewServer(cfg *ConfigServer) Server {
    return &server{
		cfg: cfg,
    }
}

func (s *server) Run() {
	s.modeServer();
	s.connectInfrastructure();


	router := gin.Default()


	taskHandler := task.InitModule(s.cfg.Database.GetDB())
	taskHandler.RegisterRoutes(router)


	log.Printf("✅ Server running at %s://%s", s.cfg.Protocol, s.cfg.Addr)
	router.Run(s.cfg.Addr) 
}

func (s *server) connectInfrastructure() {
	if err := s.cfg.Database.ConnectSQLite(); err != nil {
        log.Fatal("DB connect error:", err)
    }

    if err := s.cfg.Database.AutoMigrate(); err != nil {
        log.Fatal("Migration error:", err)
    }
}

func (s *server) modeServer() {
	switch s.cfg.Mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

