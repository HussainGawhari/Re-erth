package router

import (
	"client-admin/controller"
	"client-admin/middlewares"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	fmt.Print("Routing \n")
	// Middleware
	// router.Use(gin.Logger())
	// router.Use(gin.Recovery())
	//Add All route
	// Routings(router)

	v1 := r.Group("/v1").Use(middlewares.Auth())
	v1.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"status": http.StatusAccepted, "message": "Route Not Found10",
		})
	})

	v1.GET("/health", controller.GetToken)
	v1.POST("/singup", controller.CreateUser)
	v1.POST("/login", controller.LoginUser)

	r.POST("/addclient", controller.Addclient)
	r.GET("/getclient/:data", controller.Getclients)
	// v1.GET("/client/:name", controller.GetClient)
	// v1.PUT("/client/:id", controller.EditClient)
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound, "message": "Route Not Found12",
		})
	})

}
