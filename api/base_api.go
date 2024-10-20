package api

import (
	"errors"
	"fmt"
	"gin-vue/global"
	"gin-vue/utils"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() *BaseApi {
	return &BaseApi{
		Logger: global.Logger,
	}
}

type BuildRequestOption struct {
	Ctx               *gin.Context
	DTO               any
	BindParamsFromUri bool
}

func (b *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {
	var errResult error

	// bind context
	b.Ctx = option.Ctx

	// bind params
	if option.DTO != nil {
		if option.BindParamsFromUri {
			errResult = b.Ctx.ShouldBindUri(option.DTO)
		} else {
			errResult = b.Ctx.ShouldBind(option.DTO)
		}

		if errResult != nil {
			errResult = b.ParseValidateErrors(errResult.(validator.ValidationErrors), option.DTO)
			b.AddError(errResult)
			Fail(b.Ctx, ResponseJson{
				Msg: b.GetError().Error(),
			})
		}
	}

	return b
}

func (b *BaseApi) AddError(errNew error) {
	b.Errors = utils.AppendError(b.Errors, errNew)
}

func (b *BaseApi) GetError() error {
	return b.Errors
}

func (b *BaseApi) ParseValidateErrors(errs validator.ValidationErrors, target any) error {

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
