package global

import (
	"gin-vue/conf"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger      *zap.SugaredLogger
	DB          *gorm.DB
	RedisClient *conf.RedisClient
)
