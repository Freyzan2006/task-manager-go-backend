package server

import (
	"github.com/gin-gonic/gin"
)


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