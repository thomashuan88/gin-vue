package dto

type ShutdownHostDTO struct {
	HostIP string `json:"host_ip" binding:"required" message:"host_ip is required"`
}
