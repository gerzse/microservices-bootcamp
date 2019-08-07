package components

import "github.com/gin-gonic/gin"

func BuildRoutes(r *gin.RouterGroup){

	user := r.Group("/users")
	{
		user.GET("/user/:uuid", Read)
		user.POST("/user/:uuid", Create)
		user.DELETE("/user/:uuid", Detele)
		user.PUT("/user/:uuid", Update)

		user.GET("/all/", GetAllUsers)
	}
}
