package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	server.initRouter()
	return server
}

func (server *Server) Run(addr ...string) {
	err := server.router.Run(addr...)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}
