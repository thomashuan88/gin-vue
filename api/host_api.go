package api

import (
	"gin-vue/service"
	"gin-vue/service/dto"

	"github.com/gin-gonic/gin"
)

type HostApi struct {
	BaseApi
	Service *service.HostService
}

func NewHostApi() HostApi {
	return HostApi{
		BaseApi: NewBaseApi(),
		Service: service.NewHostService(),
	}
}

func (h HostApi) Shutdown(c *gin.Context) {
	var iShutdownHostDTO dto.ShutdownHostDTO

	if err := h.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iShutdownHostDTO}).GetError(); err != nil {
		return
	}

	if err := h.Service.Shutdown(iShutdownHostDTO); err != nil {
		h.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}

	h.OK(ResponseJson{
		Msg: "shutdown success",
	})
}
