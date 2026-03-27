package main

import (
	"github.com/artur-sres/golang-event-booking-api/db"
	"github.com/artur-sres/golang-event-booking-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost: 8080
}
