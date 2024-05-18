package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yenyoong99/goProjects_yyblog/apps/blog"
	"github.com/yenyoong99/goProjects_yyblog/apps/user"
	"github.com/yenyoong99/goProjects_yyblog/ioc"
	"github.com/yenyoong99/goProjects_yyblog/middleware"
)

func init() {
	ioc.Api().Registry(blog.AppName, &blogApiHandler{})
}

type blogApiHandler struct {
	svc blog.Service
}

func (i *blogApiHandler) Init() error {
	i.svc = ioc.Controller().Get(blog.AppName).(blog.Service)
	return nil
}

func (i *blogApiHandler) Destroy() error {
	return nil
}

func (i *blogApiHandler) Registry(rr gin.IRouter) {
	r := rr.Group(blog.AppName)

	// guest api, no need permission
	r.GET("/", i.QueryBlog)
	r.GET("/:id", i.DescribeBlog)

	// add authenticate
	r.Use(middleware.Auth)

	// back-end api, need permission
	r.POST("/", i.CreateBlog)
	r.PATCH("/:id", i.PatchBlog)
	r.PUT("/:id", i.UpdateBlog)
	r.POST("/upload", i.UploadBlogImg) // Add the upload route

	// only admin role allow to delete
	r.DELETE("/:id", middleware.Required(user.RoleAdmin), i.DeleteBlog)
}
