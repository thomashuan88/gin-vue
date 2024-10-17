package router

import "github.com/gin-gonic/gin"

func InitUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		rgPublic.POST("/login", func(c *gin.Context) {
			c.AbortWithStatusJSON(200, gin.H{
				"message": "login success",
			})
		})

		rgAuthUser := rgAuth.Group("/user")
		rgAuthUser.GET("", func(c *gin.Context) {
			c.AbortWithStatusJSON(200, gin.H{
				"data": []map[string]any{
					{"id": 1, "name": "admin"},
					{"id": 2, "name": "user"},
				},
			})
		})

		rgAuthUser.GET("/:id", func(c *gin.Context) {
			c.AbortWithStatusJSON(200, gin.H{
				"data": map[string]any{"id": 1, "name": "admin"},
			})
		})
	})

}
