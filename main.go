package main

import (
	"client-admin/api"
	db "client-admin/pkg/database"
	"fmt"
	"log"
)

func main() {

	_, err := db.ConnectToPostgres()
	if err != nil {
		// TODO Error log need to done
		log.Fatal(err)
		return
	}

	api := api.RegisterRoutes()
	fmt.Println("hello jawan")
	//start server
	api.Run(":8000")

}
