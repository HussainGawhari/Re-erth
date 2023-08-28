package api

import (
	"client-admin/middlewares"

	"github.com/gin-gonic/gin"
)

func Routings(route *gin.Engine) {

	v1 := route.Group("/v1")

	//client
	//
	client := v1.Group("/client").Use(middlewares.Auth())

	{
		client.POST("/client", Addclient)
		client.GET("/client", Getclients)
		client.GET("/client/:name", GetClient)
		client.PUT("/client/:id", EditClient)
		client.PUT("/")
	}

	//User
	// user := v1.Group("/user")
	// user.POST("/", CreateUser)
	// user.POST("/login", LoginUser)

}
