package api

import (
	"gin-vue/service"
	"gin-vue/service/dto"
	"gin-vue/utils"

	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_ADD_USER       = 10011
	ERR_CODE_GET_USER_BY_ID = 10012
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

func (u UserApi) AddUser(c *gin.Context) {
	var iUserAddDTO dto.UserAddDTO

	if err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserAddDTO}).GetError(); err != nil {
		return
	}

	if err := u.Service.AddUser(&iUserAddDTO); err != nil {
		u.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})
	}

	u.OK(ResponseJson{
		Data: iUserAddDTO,
	})
}

func (u UserApi) GetUserById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO

	if err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindParamsFromUri: true}).GetError(); err != nil {
		return
	}

	iUser, err := u.Service.GetUserById(&iCommonIDDTO)
	if err != nil {
		u.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})
		return
	}

	u.OK(ResponseJson{
		Data: iUser,
	})
}
