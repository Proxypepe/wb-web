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

func (server *Server) Run(addr ...string) {
	err := server.route.Run(addr...)
	if err != nil {
		log.Fatal(err)
		return
	}
}
