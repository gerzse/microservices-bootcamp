package main

import (
	"bytes"
	_ "crypto/rand"
	_ "encoding/hex"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/satori/go.uuid"
	_ "image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)
func postFile(baseName string, storagePath string, uuid string, photoAddr string) error{

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	filePath := path.Join(storagePath, baseName + ".jpg")
	fileWriter, err := bodyWriter.CreateFormFile("photo", filePath)
	if err != nil{
		return err
	}
	fileHandler, err := os.Open(filePath)
	if err != nil{
		return err
	}
	defer fileHandler.Close()
	_, err = io.Copy(fileWriter, fileHandler)
	if err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	_ = bodyWriter.Close()

	resp, err := http.Post(photoAddr + "/" + uuid, contentType, bodyBuf)
	if err != nil{
		return err
	}
	defer resp.Body.Close()
	fmt.Println(fmt.Sprintf("Set image for user: %s", baseName))
	return nil
}

func addUser(baseName string, mockStoragePath string, photoAddr string, userAddr string){

	uuid := uuid.NewV4().String()
	userData, _ := ioutil.ReadFile(path.Join(mockStoragePath, baseName+".json"))

	resp, err := http.Post(userAddr + "/user/" + uuid, "application/json", bytes.NewBuffer(userData))
	if err != nil{
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Println(fmt.Sprintf("Created user: %s with UUID %s ", baseName, uuid))

	err = postFile(baseName, mockStoragePath, uuid, photoAddr)
	if err != nil{
		log.Fatalln(err)
	}
}
func main(){

	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error while loading .env file")
	}
	userServiceAddress := os.Getenv("USER_SERVICE_BASE_ADDRESS")
	photoServiceAdress := os.Getenv("PHOTO_SERVICE_BASE_ADDRESS")

	dir, err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}
	baseDir := path.Dir(dir)
	mockDataRoot := path.Join(baseDir, "/mock-data")

	files, err := ioutil.ReadDir(mockDataRoot)
	if err != nil{
		log.Fatal(err)
	}
	for _, file := range files{
		if strings.HasSuffix(file.Name(), ".json"){
			filePath := path.Join(mockDataRoot, file.Name())
			fileName := path.Base(filePath)
			baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
			addUser(baseName, mockDataRoot, photoServiceAdress,userServiceAddress)
		}
	}
}