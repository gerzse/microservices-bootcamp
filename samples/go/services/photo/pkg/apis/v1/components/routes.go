package components

import "github.com/gin-gonic/gin"

func BuildRoutes(r *gin.RouterGroup){

	photo := r.Group("/photo")
	{
		photo.GET("/:uuid", GetPhoto)
		photo.POST("/:uuid", PostPhoto)
	}

}
