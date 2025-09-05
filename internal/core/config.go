package core

import (
	"fmt"
	"task-manager/pkg/console"
	"task-manager/pkg/database"
)

type Config struct {
	Mode        string
	Host        string
	Port        string
	Protocol    string
	Addr        string
	DatabaseDSN string
	Database    *database.DB 
}

func NewConfig() *Config {
	flags := console.NewFlags()

	// создаём объект базы заранее
	db := database.NewDB(flags.Database)

	return &Config{
		Mode:        flags.Mode,
		Host:        flags.Host,
		Port:        flags.Port,
		Protocol:    flags.Protocol,
		Addr:        fmt.Sprintf("%s:%s", flags.Host, flags.Port),
		DatabaseDSN: flags.Database,
		Database:    db, 
	}
}
