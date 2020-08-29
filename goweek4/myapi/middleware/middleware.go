package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	log.Println("start middleware")
	authKey := c.GetHeader("Authorization")
	if authKey != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		c.Abort()
		return
	}
	c.Next()
	log.Println("end middleware")
}
