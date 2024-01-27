package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yenyoong99/goProjects_yyblog/apps/blog"
	"github.com/yenyoong99/goProjects_yyblog/apps/token"
	"github.com/yenyoong99/goProjects_yyblog/common"
	"github.com/yenyoong99/goProjects_yyblog/response"
)

// CreateBlog POST /yyblog/api/v1/blogs
func (i *blogApiHandler) CreateBlog(c *gin.Context) {
	// h.tk.Validate()
	req := blog.NewCreateBlogRequest()

	// get middleware info
	if v, ok := c.Get(token.TOKEN_MIDDLEWARE_KEY); ok {
		req.CreateBy = v.(*token.Token).UserName
	}

	if err := c.BindJSON(req); err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := i.svc.CreateBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// PatchBlog PATCH /yyblog/api/v1/blogs/:id
// /yyblog/api/v1/blogs/10 --> id = 10
// /yyblog/api/v1/blogs/20 --> id = 20
// c.Param("id") get request param value
func (i *blogApiHandler) PatchBlog(c *gin.Context) {
	req := blog.NewUpdateBlogRequest(c.Param("id"))
	req.UpdateMode = common.UPDATE_MODE_PATCH

	if err := c.BindJSON(req.CreateBlogRequest); err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := i.svc.UpdateBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, ins)

}

// UpdateBlog PUT /yyblog/api/v1/blogs/:id
func (i *blogApiHandler) UpdateBlog(c *gin.Context) {
	req := blog.NewUpdateBlogRequest(c.Param("id"))
	req.UpdateMode = common.UPDATE_MODE_PUT
	if err := c.BindJSON(req.CreateBlogRequest); err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := i.svc.UpdateBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// DeleteBlog DELETE /yyblog/api/v1/blogs/:id
func (i *blogApiHandler) DeleteBlog(c *gin.Context) {
	req := blog.NewDeleteBlogRequest(c.Param("id"))
	ins, err := i.svc.DeleteBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// QueryBlog GET /yyblog/api/v1/blogs?page_size=10&page_number=2
func (i *blogApiHandler) QueryBlog(c *gin.Context) {
	req := blog.NewQueryBlogRequestFromGin(c)
	set, err := i.svc.QueryBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, set)
}

// DescribeBlog GET /yyblog/api/v1/blogs/:id
func (i *blogApiHandler) DescribeBlog(c *gin.Context) {
	req := blog.NewDescribeBlogRequest(c.Param("id"))
	ins, err := i.svc.DescribeBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}
