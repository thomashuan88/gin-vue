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
			rgAuthUser.POST("", userApi.AddUser)
			rgAuthUser.GET("/:id", userApi.GetUserById)
		}

	})

}
