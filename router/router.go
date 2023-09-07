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
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	r.Use(cors.New(config))
	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// Authentication for admin
	v1 := r.Group("/v1").Use(middlewares.Auth())
	{
		v1.POST("/addclient", controller.Addclient)
		v1.DELETE("/client/:id", controller.DeleteClient)   //delete based on ID
		v1.PUT("/client/:id", controller.EditClient)        //update based on ID
		v1.PUT("/toggle/:id", controller.ActivateDeactivat) // This route can change the client to active and deactive
		v1.GET("/details/:id", controller.GetClientBasedID) //Get based on ID of client and display the details

	}
	// Authentication for user's
	v2 := r.Group("/v2").Use(middlewares.AuthForUsers())
	{
		v2.GET("/health", controller.GetToken)
		v2.GET("/clients", controller.GetAllclients)       //returns all the client which are actives
		v2.GET("/getclient/", controller.Getclient)        // this rout can search for the client based on last name, country,postal code,city and so on..
		v2.GET("/count", controller.CountClients)          // return the number of client in DB
		v2.GET("/clienthistory", controller.ClientHistory) // return history of the client

	}
	// User's signup and login routes
	r.POST("/login", controller.LoginUser)
	r.POST("/singup", controller.CreateUser)
	r.GET("/users", controller.GetUsers)

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound, "message": "Route Not Found12",
		})
	})

}
