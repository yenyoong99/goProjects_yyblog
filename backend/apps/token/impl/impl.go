package impl

import (
	"github.com/yenyoong99/goProjects_yyblog/apps/token"
	"github.com/yenyoong99/goProjects_yyblog/apps/user"
	"github.com/yenyoong99/goProjects_yyblog/conf"
	"github.com/yenyoong99/goProjects_yyblog/ioc"
	"gorm.io/gorm"
)

var (
	_ token.Service = (*TokenServiceImpl)(nil)
)

func init() {
	// new(TokenServiceImpl)
	ioc.Controller().Registry(token.AppName, &TokenServiceImpl{})
}

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

func (i *TokenServiceImpl) Init() error {
	i.db = conf.C().DB()
	i.user = ioc.Controller().Get(user.AppName).(user.Service)
	return nil
}

func (i *TokenServiceImpl) Destroy() error {
	return nil
}
