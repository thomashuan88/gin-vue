package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	cfg := cors.Config{
		AllowOriginFunc: func(origin string) bool { return true },
		// AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept", "token"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}

	return cors.New(cfg)
}
