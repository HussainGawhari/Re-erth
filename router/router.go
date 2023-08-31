package router

import (
	"client-admin/controller"
	"client-admin/middlewares"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Update with your frontend URL
	config.AllowMethods = []string{"*", "GET", "POST", "OPTIONS", "DELETE", "PUT"}
	config.AllowHeaders = []string{"Content-Type"}
	r.Use(cors.New(config))
	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/login", controller.LoginUser)
	r.POST("/singup", controller.CreateUser)
	r.GET("/users", controller.GetUsers)

	v1 := r.Group("/v1").Use(middlewares.Auth())
	{
		v1.GET("/health", controller.GetToken)
		v1.POST("/addclient", controller.Addclient)
		v1.DELETE("/client/:id", controller.DeleteClient)
		v1.GET("/clients", controller.GetAllclients)
		v1.GET("/getclient", controller.Getclient)
		v1.PUT("/client/:id", controller.EditClient)
	}

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound, "message": "Route Not Found12",
		})
	})

}
