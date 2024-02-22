package main

import (
	"api/routes"

	"github.com/gin-gonic/gin"
)

var Server = gin.Default()

func main() {
	Server.SetTrustedProxies(nil)
	routes.AddRoutes(Server)
	Server.Run(":8080")
}
