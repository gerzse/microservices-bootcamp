package components

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var userStorage = make(map[string]User)

func Create(c *gin.Context){
	uuid := c.Params.ByName("uuid")
	if _, ok := userStorage[uuid]; ok{
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("With Id: %s already exists", uuid),
		})
	}
	var user User;
	err := c.BindJSON(&user)
	if err != nil{
		c.JSON(500, gin.H{
			"message": "Mapping error",
		})
	}else{
		userStorage[uuid] = user;
	}


}

func Update(c *gin.Context){

}

func Read(c *gin.Context){
	uuid := c.Params.ByName("uuid")
	if _, ok := userStorage[uuid]; ok{
		c.JSON(200, userStorage[uuid])
	}else{
		c.String(500, "User with id: %s dose not exist", uuid)
	}
}

func Detele(c *gin.Context){

	uuid := c.Params.ByName("uuid")
	if _, ok := userStorage[uuid]; ok{
		delete(userStorage, uuid)
		c.String(200, "User with id: %s was removed from storage", uuid)
	}else{
		c.String(500, "User with id: %s dose not exist", uuid)
	}
}

func GetAllUsers(c *gin.Context){

	users := make([]User, len(userStorage))

	for _, u := range userStorage{
		users = append(users, u)
	}
	c.JSON(200,  users)
}