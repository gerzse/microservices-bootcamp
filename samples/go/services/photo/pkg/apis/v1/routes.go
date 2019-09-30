package v1

import (
	"github.com/bigp/microservices-bootcamp/services/photo/pkg/apis/v1/components"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func BuildRoutes(r *gin.RouterGroup){

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", ping)
		components.BuildRoutes(v1)
	}
}