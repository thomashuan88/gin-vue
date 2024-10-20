package api

import (
	"errors"
	"fmt"
	"gin-vue/service/dto"
	"gin-vue/utils"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserApi struct {
	BaseApi
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
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
	errs := c.ShouldBind(&iUserLoginDTO)
	fmt.Printf("iUserLoginDTO: %+v\n", iUserLoginDTO)
	if errs != nil {
		Fail(c, ResponseJson{
			Msg: parseValidateErrors(errs.(validator.ValidationErrors), &iUserLoginDTO).Error(),
		})
		return
	}

	if err := u.BuildRequest({Ctx: c, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	OK(c, ResponseJson{
		Data: iUserLoginDTO,
	})

	// OK(c, ResponseJson{
	// 	Msg: "Login Success",
	// })

	// Fail(c, ResponseJson{
	// 	Code: 9001,
	// 	Msg:  "Login Fail",
	// })

}

func parseValidateErrors(errs validator.ValidationErrors, target any) error {

	var errResult error

	// from reflection gain the pointer type of target
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errs {

		field, _ := fields.FieldByName(fieldErr.Field())
		// return fieldErr.Field() & field
		// return errors.New(fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag()))

		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())

		errMessage := field.Tag.Get(errMessageTag)
		// return errors.New(errMessage + " hh")
		if errMessage != "" {
			errMessage = field.Tag.Get("message")
		}
		// return errors.New(errMessage + " hhh")
		if errMessage == "" {
			errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}

		// if utils.AppendError(errResult, errors.New(errMessage)) no return value to errResult,
		// like this will cause runtime error: invalid memory address
		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}

	return errResult
}
