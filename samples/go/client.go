package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
)

type Name struct{
	First string `json:"first"`
	Last  string `json:"last"`
}

type User struct{
	Name Name `json:"name"`
	Gender string `json:"gender"`
	BirthYear uint16 `json:"born"`
}

func generateUUID(size uint) string{

	// Generate UUID(size uint) -> will generate random uuid
	b := make([]byte, size)
	_ , err := rand.Read(b)
	if err != nil{
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid

}

func readJsonFile(filename string) []byte{
	jsonFile, err := os.Open(filename)
	if err != nil{
		log.Fatal(err)
	}
	defer jsonFile.Close()
	content, _ := ioutil.ReadAll(jsonFile)
	return content

}
func main(){

	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error while loading .env file")
	}
	userServiceAddress := os.Getenv("USER_SERVICE_BASE_ADDRESS")
	dir, err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}
	baseDir := path.Dir(dir)
	fmt.Println(baseDir)

	mockDataRoot := path.Join(baseDir, "/mock-data")
	fmt.Println(mockDataRoot)

	files, err := ioutil.ReadDir(mockDataRoot)
	if err != nil{
		log.Fatal(err)
	}
	for _, file := range files{
		if strings.HasSuffix(file.Name(), ".json"){
			filePath := path.Join(mockDataRoot, file.Name())
			content := readJsonFile(filePath)
			uuid := generateUUID(16)
			resp, err := http.Post(userServiceAddress + "/user/" + uuid, "application/json", bytes.NewBuffer(content))
			if err != nil{
				log.Fatal(err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil{
				log.Fatal()
			}
			log.Println(string(body))
		}
	}
}