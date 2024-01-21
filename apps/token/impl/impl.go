package impl

import (
	"goProjects/yyblog/apps/token"
	"goProjects/yyblog/apps/user"
	"goProjects/yyblog/conf"
	"gorm.io/gorm"
)

var (
	_ token.Service = (*TokenServiceImpl)(nil)
)

func NewTokenServiceImpl(userServiceImpl user.Service) *TokenServiceImpl {
	return &TokenServiceImpl{
		db:   conf.C().DB(),
		user: userServiceImpl,
	}
}

type TokenServiceImpl struct {
	db *gorm.DB

	user user.Service
}
