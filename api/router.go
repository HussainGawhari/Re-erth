package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {

	// Gin Instance
	router := gin.New()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Add All route
	Routings(router)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound, "message": "Route Not Found",
		})
	})
	return router
}
