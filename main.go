package main

import (
	"com.sal/main/db"
	"com.sal/main/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost
}
