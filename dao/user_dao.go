package dao

import "gin-vue/model"

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
