package apis

import (
	v1 "github.com/bigp/microservices-bootcamp/services/user/pkg/apis/v1"
	"github.com/gin-gonic/gin"
)

func BuildRoutes(r *gin.Engine){

	api := r.Group("/api")
	{
		v1.BuildRoutes(api)
	}
}