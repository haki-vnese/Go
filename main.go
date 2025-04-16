package main

import (
	"go-rest-api/routes"
)

func main() {
	r := routes.SetupRoutes()
	r.Run(":8080")
}
