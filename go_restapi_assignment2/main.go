package main

import (
	"go_restapi_assignment2/config"
	"go_restapi_assignment2/routers"
	// "go_restapi_assignment2/routers"
)

func main() {
	const PORT = ":8090"
	config.DBInit()
	routers.StartServer().Run(PORT)
}
