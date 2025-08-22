package server 

import (
	"task-manager/pkg/database"
)

type ConfigServer struct {
	Mode 		string
	Database 	*database.DB
	Addr   		string
	Host 		string
	Port 		string
	Protocol    string // "http", "wss", "hybrid"
}

func NewConfigServer(mode string, db *database.DB, addr string, host string, port string, protocol string) *ConfigServer {
	return &ConfigServer{
		Mode: mode,
		Database: db,
		Addr: addr,
		Host: host,
		Port: port,
		Protocol: protocol,
	}
}