package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/difoul/rest-api-mock/helpers"
	"github.com/difoul/rest-api-mock/models"
	"github.com/gin-gonic/gin"
)

var resp_c string

func main() {
	fileFlag := flag.String("resp-file", "", "The path to the configuration file that contains the response")
	flag.Parse()
	if *fileFlag == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	resp_file := *fileFlag
	if !helpers.IsFileOrDirExists(resp_file) {
		os.Stderr.WriteString(fmt.Sprintf("The file [%s] does not exist!\n", resp_file))
		os.Exit(1)
	}
	content := helpers.ReadFile(resp_file)
	//fmt.Println(content)
	var mock models.HttpMock
	err := json.Unmarshal(content, &mock)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	router := gin.Default()
	mock.RegisterEndpoints(router)
}
