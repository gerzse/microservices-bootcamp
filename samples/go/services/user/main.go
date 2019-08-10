package main

import (
	"github.com/bigp/microservices-bootcamp/services/user/pkg/apis"
	"github.com/gin-gonic/gin"
	"log"
)

func main(){

	router := gin.Default()

	apis.BuildRoutes(router)

	err := router.Run(":8080")
	if err != nil{
		log.Fatal("Could not run server on port: 8080")
	}
}