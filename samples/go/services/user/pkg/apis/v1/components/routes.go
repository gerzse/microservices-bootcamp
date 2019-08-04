package components

import "github.com/gin-gonic/gin"

func BuildRoutes(r *gin.RouterGroup){

	user := r.Group("/user")
	{
		user.GET("/:uuid", Read)
		user.POST("/:uuid", Create)
		user.DELETE("/:uuid", Detele)
		user.PUT("/:uuid", Update)
	}
}