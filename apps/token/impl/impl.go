package impl

import (
	"github.com/yenyoong99/goProjects_yyblog/apps/token"
	"github.com/yenyoong99/goProjects_yyblog/apps/user"
	"github.com/yenyoong99/goProjects_yyblog/conf"
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
