package models

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpMock struct {
	Port     int        `json:"port"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	Get      []Response `json:"get"`
	Post     []Response `json:"post"`
}

type Response struct {
	Url  string `json:"url"`
	Resp string `json:"response"`
}

func (m *HttpMock) RegisterEndpoints(engine *gin.Engine) {
	router := engine.Group("/")
	if len(m.Username) > 0 && len(m.Password) > 0 {
		router = engine.Group("/", gin.BasicAuth(gin.Accounts{m.Username: m.Password}))
	}

	for _, g := range m.Get {
		router.GET(g.Url, func(c *gin.Context) {
			c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(g.Resp))
		})
	}

	for _, p := range m.Post {
		router.POST(p.Url, func(c *gin.Context) {
			c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(p.Resp))
		})
	}

	engine.Run(fmt.Sprintf(":%d", m.Port))
}
