package routes

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(server *gin.Engine) {
	v1 := server.Group("/v1")

	addHelloRoutes(v1)
}
