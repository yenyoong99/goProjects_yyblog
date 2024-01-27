package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yenyoong99/goProjects_yyblog/apps/blog"
	"github.com/yenyoong99/goProjects_yyblog/ioc"
	"github.com/yenyoong99/goProjects_yyblog/middleware"
)

func init() {
	ioc.Api().Registry(blog.AppName, &blogApiHandler{})
}

type blogApiHandler struct {
	svc blog.Service
}

func (h *blogApiHandler) Init() error {
	h.svc = ioc.Controller().Get(blog.AppName).(blog.Service)
	return nil
}

func (h *blogApiHandler) Destroy() error {
	return nil
}

func (h *blogApiHandler) Registry(rr gin.IRouter) {
	r := rr.Group(blog.AppName)

	// guest api, no need permission
	r.GET("/", h.QueryBlog)
	r.GET("/:id", h.DescribeBlog)

	// add authenticate
	r.Use(middleware.Auth)

	// back-end api, need permission
	r.POST("/", h.CreateBlog)
	r.PATCH("/:id", h.PatchBlog)
	r.PUT("/:id", h.UpdateBlog)
	r.DELETE("/:id", h.DeleteBlog)
}
