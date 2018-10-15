package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/difoul/rest-api-mock/helpers"
	"github.com/difoul/rest-api-mock/models"
	"github.com/gin-gonic/gin"
)

var resp_c string

func main() {
	fileFlag := flag.String("resp-file", "", "The path to the configuration file that contains the response")
	flag.Parse()
	fmt.Println(*fileFlag)
	if *fileFlag == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	resp_file := *fileFlag
	if !helpers.IsFileOrDirExists(resp_file) {
		os.Stderr.WriteString(fmt.Sprintf("The file [%s] does not exist!\n", resp_file))
		os.Exit(1)
	}
	content := helpers.ReadFile("conf/response.json")
	//fmt.Println(content)
	var resp models.HttpGetMock
	err := json.Unmarshal(content, &resp)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	resp_c = resp.Response
	fmt.Println(resp_c)

	router := gin.Default()
	router.GET(resp.Url, GetResp)
	router.Run(fmt.Sprintf(":%d", resp.Port)) // listen and serve on 0.0.0.0:8080
}

func GetResp(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{"message": resp_c})
	//c.JSON(http.StatusOK, resp_c)
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(resp_c))
}
