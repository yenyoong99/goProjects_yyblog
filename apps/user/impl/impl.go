package impl

import (
	"goProjects/yyblog/conf"
	"gorm.io/gorm"
)

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
