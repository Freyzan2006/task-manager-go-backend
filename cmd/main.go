package main

import (
    "log"
    "fmt"
)

import (
	"task-manager/pkg/database"
    "task-manager/internal/server"
    "task-manager/pkg/console"
)

func main() {
    flags := console.NewFlags()
    log.Println("✅ Flags created successfully")

    db := database.NewDB(flags.Database)
    log.Println("✅ Database connected and migrated successfully")

    addr := fmt.Sprintf("%s:%s", flags.Host, flags.Port)
    cfg := server.NewConfigServer(flags.Mode, db, addr, flags.Host, flags.Port, flags.Protocol)
    log.Println("✅ Config created successfully")

    server := server.NewServer(cfg)
    server.Run()
    log.Println("✅ Server started successfully")

    
}
