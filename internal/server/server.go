package server

import (
	"log"
	"net/http"
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
	switch s.cfg.Mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

    if err := s.cfg.Database.ConnectSQLite(); err != nil {
        log.Fatal("DB connect error:", err)
    }

    if err := s.cfg.Database.AutoMigrate(); err != nil {
        log.Fatal("Migration error:", err)
    }

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })


	log.Printf("✅ Server running at %s://%s", s.cfg.Protocol, s.cfg.Addr)
	r.Run(s.cfg.Addr) 
}
