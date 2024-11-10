package service

import (
	"context"
	"fmt"
	"gin-vue/service/dto"

	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/spf13/viper"
)

var hostService *HostService

type HostService struct {
	BaseService
}

func NewHostService() *HostService {
	if hostService == nil {
		hostService = &HostService{}
	}
	return hostService
}

func (h *HostService) Shutdown(iShutdownHostDTO dto.ShutdownHostDTO) error {
	var errResult error

	stHostIP := iShutdownHostDTO.HostIP
	fmt.Println(stHostIP)

	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "ssh",
		User:       viper.GetString("ansible.user"),
	}

	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  fmt.Sprintf("%s,", stHostIP),
		ModuleName: "command",
		Args:       viper.GetString("ansible.ShutdownHost.Args"),
		ExtraVars: map[string]any{
			"ansible_password": viper.GetString("ansible.password"),
		},
	}

	adhoc := &adhoc.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
		StdoutCallback:    "oneline",
	}

	errResult = adhoc.Run(context.TODO())
	if errResult != nil {
		return errResult
	}

	return errResult
}
