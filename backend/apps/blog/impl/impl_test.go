package impl_test

import (
	"context"

	"github.com/yenyoong99/goProjects_yyblog/apps/blog"
	"github.com/yenyoong99/goProjects_yyblog/ioc"

	_ "github.com/yenyoong99/goProjects_yyblog/apps"
)

var (
	impl blog.Service
	ctx  = context.Background()
)

func init() {
	// 2. ioc get object
	impl = ioc.Controller().Get(blog.AppName).(blog.Service)

	// ioc need init to fill the db attribute
	if err := ioc.Init(); err != nil {
		panic(err)
	}
}
