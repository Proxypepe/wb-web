package http

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	route *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	server.initRouter()
	return server
}

func (server *Server) Run() {
	err := server.route.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
