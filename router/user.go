package router

import (
	"gin-vue/api"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()

		rgPublicUser := rgPublic.Group("/user")
		{
			rgPublicUser.POST("/login", userApi.Login)
		}

		rgAuthUser := rgAuth.Group("/user")
		{
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
		}

	})

}
