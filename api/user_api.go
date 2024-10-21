package api

import (
	"gin-vue/service"
	"gin-vue/service/dto"
	"gin-vue/utils"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
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
func (u UserApi) Login(c *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO

	if err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	iUser, err := u.Service.Login(iUserLoginDTO)
	if err != nil {
		u.Fail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}

	token, _ := utils.GenerateToken(iUser.ID, iUser.Name)

	u.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
	})

	// OK(c, ResponseJson{
	// 	Msg: "Login Success",
	// })

	// Fail(c, ResponseJson{
	// 	Code: 9001,
	// 	Msg:  "Login Fail",
	// })

}
