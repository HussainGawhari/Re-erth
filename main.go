package main

import (
	"client-admin/pkg/helperdb"
	"client-admin/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := helperdb.ConnectToPostgres()
	if err != nil {
		// TODO Error log need to done
		log.Fatal(err)
		return
	}
	defer db.Close()

	r := gin.New()
	router.RegisterRoutes(r)

	r.Run(":8000")
}
