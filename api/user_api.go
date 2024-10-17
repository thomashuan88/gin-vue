package api

import "github.com/gin-gonic/gin"

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

// @Summary login
// @Description login
// @Param name formData string true "name"
// @Param password formData string true "password"
// @Tags User Api
// @Accept json
// @Produce json
// @Success 200 {string} string "login success"
// @Failure 401 {string} string "login fail"
// @Router /api/v1/public/user/login [post]
func (u UserApi) Login(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(200, gin.H{
		"message": "login success",
	})

}
