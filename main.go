package main

import (
	"m-server-api/bootstrap"
	"m-server-api/initializers"
	"m-server-api/pkg/shutdown"

	_ "m-server-api/modules/admin/routes"
)

func main() {
	// Initialize Db
	initializers.InitDatabase()
	// Initialize Gin engine
	s := bootstrap.NewServer()
	// Start the server
	go s.Run()
	// Graceful shutdown
	shutdown.Close(func() {
		if err := initializers.CloseMysql(); err != nil {
			panic(err)
		}
		if err := initializers.CloseRedis(); err != nil {
			panic(err)
		}
	})
}
