package server

import (
	"log"
)

func (s *server) connectInfrastructure() {
	if err := s.cfg.Database.ConnectSQLite(); err != nil {
        log.Fatal("DB connect error:", err)
    }

    if err := s.cfg.Database.AutoMigrate(); err != nil {
        log.Fatal("Migration error:", err)
    }
}