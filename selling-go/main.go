package main

import (
	"selling-go/databases"
	"selling-go/routers"
)

func main() {
	db := databases.StartDB()

	r := routers.StartApp(db)

	r.Run(":8090")
}
