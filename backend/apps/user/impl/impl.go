package impl

import (
	"github.com/yenyoong99/goProjects_yyblog/apps/user"
	"github.com/yenyoong99/goProjects_yyblog/conf"
	"github.com/yenyoong99/goProjects_yyblog/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.Controller().Registry(user.AppName, &UserServiceImpl{})
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		//get global configurations db object
		db: conf.C().DB(),
	}

}

// UserServiceImpl
// define UserServiceImpl to implement the interface
type UserServiceImpl struct {
	db *gorm.DB
}

func (i *UserServiceImpl) Init() error {
	i.db = conf.C().DB()
	return nil
}

func (i *UserServiceImpl) Destroy() error {
	// delete map
	return nil
}
