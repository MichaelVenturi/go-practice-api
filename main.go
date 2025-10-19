package main

import (
	"github.com/MichaelVenturi/go-practice-api/db"
	"github.com/MichaelVenturi/go-practice-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost 8080
}
