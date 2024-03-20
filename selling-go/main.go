package main

import (
	"selling-go/databases"
	"selling-go/routers"
)

func main() {
	databases.StartDB()

	r := routers.StartApp()

	r.Run(":8090")
}
