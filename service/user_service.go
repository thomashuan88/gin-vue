package service

import (
	"errors"
	"gin-vue/dao"
	"gin-vue/model"
	"gin-vue/service/dto"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

func (u *UserService) Login(iUserDTO dto.UserLoginDTO) (model.User, error) {
	var errResult error
	iUser := u.Dao.GetUserByNameAndPassword(iUserDTO.Name, iUserDTO.Password)
	if iUser.ID == 0 {
		errResult = errors.New("invalid username or password")
	}
	return iUser, errResult
}

func (u *UserService) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	if u.Dao.CheckUserExist(iUserAddDTO.Name) {
		return errors.New("user already exists")
	}
	return u.Dao.AddUser(iUserAddDTO)
}

func (u *UserService) GetUserById(iCommonIDDTO *dto.CommonIDDTO) (model.User, error) {
	return u.Dao.GetUserById(iCommonIDDTO.ID)
}
