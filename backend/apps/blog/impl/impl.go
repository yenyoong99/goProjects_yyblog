package impl

import (
	"github.com/yenyoong99/goProjects_yyblog/apps/blog"
	"github.com/yenyoong99/goProjects_yyblog/conf"
	"github.com/yenyoong99/goProjects_yyblog/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.Controller().Registry(blog.AppName, &blogServiceImpl{})
}

// blog.Service interface, direct register to Ioc, no need export
type blogServiceImpl struct {
	// A connection pool object that relies on a database operation
	db *gorm.DB
}

func (i *blogServiceImpl) Init() error {
	i.db = conf.C().DB()
	return nil
}

func (i *blogServiceImpl) Destroy() error {
	return nil
}
