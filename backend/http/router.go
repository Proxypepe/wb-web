package http

import (
	"context"
	"github.com/Proxypepe/wb-web/backend/cache"
	"github.com/Proxypepe/wb-web/backend/db"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (server *Server) initRouter() {
	server.router = gin.Default()
	server.router.Use(CORSMiddleware())
	server.router.GET("/order", server.getOrder)
}

func (server *Server) getOrder(c *gin.Context) {
	uid := c.Query("order_uid")
	order, err := cache.GetOrder(uid)

	if err == redis.Nil {
		ctx := context.Background()
		orderDb, errorDb := db.GetOrderByUID(ctx, uid)
		if errorDb != nil {
			log.Print(errorDb.Error())
			c.String(http.StatusNotFound, "Unknown uid")
			return
		}
		err := cache.SaveOrder(orderDb)
		if err != nil {
			log.Print(errorDb.Error())
			c.String(http.StatusNotFound, "Unknown uid")
			return
		}
		c.JSON(http.StatusOK, orderDb)
		return
	}

	if err != nil {
		log.Print(err)
		c.String(http.StatusNotFound, "Unknown uid")
		return
	}
	if order == nil {
		log.Print(err)
		c.String(http.StatusNotFound, "Unknown uid")
		return
	}
	c.JSON(http.StatusOK, order)
}
