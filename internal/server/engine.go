package server


import (
	"task-manager/internal/modules/task"
	"task-manager/internal/core"
	"task-manager/internal/core/logger"
)

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server interface {
    Run() error
}

type server struct {
	cfg *core.Config
	log logger.Logger
} 

func New(cfg *core.Config, log logger.Logger) Server {
    return &server{
		cfg: cfg,
		log: log,
    }
}

func (s *server) Run() error {
	s.modeServer();
	s.connectInfrastructure();

	router := gin.New()
	router.Use(gin.Logger()) 
	router.Use(gin.Recovery())
	router.SetTrustedProxies([]string{"10.0.0.0/8"})

	taskHandler := task.InitModule(s.cfg.Database.GetDB())
	taskHandler.RegisterRoutes(router)
	
	if s.cfg.Mode != "production" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	
	s.log.Infof("✅ Server running at %s://%s", s.cfg.Protocol, s.cfg.Addr)
	s.log.Infof("✅ Swagger UI: %s://%s/swagger/index.html", s.cfg.Protocol, s.cfg.Addr)

	router.Run(s.cfg.Addr) 

	return nil
}





