package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HttpMock struct {
	Port     int        `json:"port"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	Get      []Response `json:"get"`
	Post     []Response `json:"post"`
	Put      []Response `json:"put"`
	Patch    []Response `json:"patch"`
	Delete   []Response `json:"delete"`
}

type Response struct {
	Url      string `json:"url"`
	Resp     string `json:"response"`
	RespCode int    `json:"response_code"`
}

func (r *Response) HandleResponse(c *gin.Context) {
	if r.RespCode < 100 || r.RespCode > 527 {
		r.RespCode = 200
	}
	c.Data(r.RespCode, "application/json; charset=utf-8", []byte(r.Resp))
}

func (m *HttpMock) RegisterEndpoints(engine *gin.Engine) {
	router := engine.Group("/")
	if len(m.Username) > 0 && len(m.Password) > 0 {
		router = engine.Group("/", gin.BasicAuth(gin.Accounts{m.Username: m.Password}))
	}

	for _, r := range m.Get {
		// https://stackoverflow.com/questions/48826460/using-pointers-in-a-for-loop-golang
		resp := r
		router.GET(resp.Url, resp.HandleResponse)
	}

	for _, r := range m.Post {
		resp := r
		router.POST(resp.Url, resp.HandleResponse)
	}

	for _, r := range m.Put {
		resp := r
		router.PUT(resp.Url, resp.HandleResponse)
	}

	for _, r := range m.Patch {
		resp := r
		router.PATCH(resp.Url, resp.HandleResponse)
	}

	for _, r := range m.Delete {
		resp := r
		router.DELETE(resp.Url, resp.HandleResponse)
	}

	engine.Run(fmt.Sprintf(":%d", m.Port))
}
