package dto

import "gin-vue/model"

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required" message:"name wrong format" required_err:"name is required"`
	Password string `json:"password" binding:"required" message:"password is required"`
}

// Add user DTO
type UserAddDTO struct {
	ID       uint
	Name     string `json:"name" form:"name" binding:"required" message:"name is required"`
	RealName string `json:"real_name" form:"real_name"`
	Avatar   string
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"password is required"`
}

func (u *UserAddDTO) ConvertToModel(iUser *model.User) {
	iUser.Name = u.Name
	iUser.RealName = u.RealName
	iUser.Mobile = u.Mobile
	iUser.Email = u.Email
	iUser.Password = u.Password
}
