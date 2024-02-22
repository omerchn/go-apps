package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addHelloRoutes(rg *gin.RouterGroup) {
	hello := rg.Group("/hello")

	hello.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello! add /{name} to get a personalized hello!")
	})

	hello.GET("/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.JSON(http.StatusOK, fmt.Sprintf("Hello %v!", name))
	})
}
