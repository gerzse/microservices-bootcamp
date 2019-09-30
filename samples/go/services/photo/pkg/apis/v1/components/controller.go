package components

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

var photoStorage = make(map[string]string)

func validToken(uuid string)(bool){

	userServiceAddress := os.Getenv("USER_SERVICE_BASE_ADDRESS")
	resp, err := http.Get(userServiceAddress + "/user/" + uuid)
	if err != nil{
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200{
		return true
	}
	return false
}

func GetPhoto(c *gin.Context){
	uuid := c.Params.ByName("uuid")
	if _, ok := photoStorage[uuid]; ok{
		photoFilename := photoStorage[uuid]
		c.File(photoFilename)
	}
}

func PostPhoto(c *gin.Context){

	imageStorage := os.Getenv("MEDIA_STORAGE_LOCATION_PATH")
	uuid := c.Params.ByName("uuid")
	if _, ok := photoStorage[uuid]; ok{
		c.String(200, "Photo was already " +
									"stored in the media storage")
	}
	if validToken(uuid){
		file, handler, err := c.Request.FormFile("photo")
		filename := handler.Filename
		filePath := path.Join(imageStorage,  "/" + path.Base(filename))
		out, err := os.Create(filePath)
		if err != nil{
			log.Fatalln(err)
		}
		defer out.Close()
		photoStorage[uuid] = filePath
		_, err = io.Copy(out, file)
		if err != nil{
			log.Fatalln(err)
		}
	}
}