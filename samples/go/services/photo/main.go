package main

import (
	"fmt"
	"github.com/bigp/microservices-bootcamp/services/photo/pkg/apis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main(){

	err := godotenv.Load()
	if err != nil{
		log.Fatalln(err)
	}
	port := os.Getenv("PORT")

	router := gin.Default()
	apis.BuildRoutes(router)

	err = router.Run(fmt.Sprintf(":%s", port))
	if err != nil{
		log.Fatalln("Could not run server on port %s", port)
	}
}