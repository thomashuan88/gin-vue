package dao

import (
	"gin-vue/model"
	"gin-vue/service/dto"
)

var userDao *UserDao

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{BaseDao: NewBaseDao()}
	}
	return userDao
}

type UserDao struct {
	BaseDao
}

func (u *UserDao) GetUserByNameAndPassword(stUserName, stPassword string) model.User {
	var iUser model.User
	u.Orm.Model(&iUser).Where("name = ? AND password = ?", stUserName, stPassword).Find(&iUser)
	return iUser
}

func (u *UserDao) CheckUserExist(stUserName string) bool {
	var iUser model.User
	u.Orm.Model(&iUser).Where("name = ?", stUserName).Find(&iUser)
	return iUser.ID > 0
}

func (u *UserDao) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	var iUser model.User
	iUserAddDTO.ConvertToModel(&iUser)

	err := u.Orm.Save(&iUser).Error
	if err == nil {
		iUserAddDTO.ID = iUser.ID
		iUserAddDTO.Password = ""
	}
	return err
}

func (u *UserDao) GetUserById(id uint) (model.User, error) {
	var iUser model.User
	err := u.Orm.First(&iUser, id).Error
	return iUser, err
}
