package dto

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required,first_is_a" message:"name wrong format" required_err:"name is required"`
	Password string `json:"password" binding:"required" message:"password is required"`
}
