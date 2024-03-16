package blog

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yenyoong99/goProjects_yyblog/common"
	"strconv"
)

const (
	AppName = "blogs"
)

// Service Define blog Service interface
type Service interface {
	// CreateBlog Created blog
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)
	ChangedBlogStatus(context.Context, *ChangedBlogStatusRequest) (*Blog, error)
	AuditBlog(context.Context, *AuditInfo) (*Blog, error)
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

func NewQueryBlogRequestFromGin(c *gin.Context) *QueryBlogRequest {
	req := NewQueryBlogRequest()
	req.CreateBy = c.Query("create_by")
	req.Keywords = c.Query("keywords")
	ps := c.Query("page_size")
	if ps != "" {
		req.PageSize, _ = strconv.Atoi(ps)
	}
	pn := c.Query("page_number")
	if pn != "" {
		req.PageNumber, _ = strconv.Atoi(pn)
	}
	return req
}

type QueryBlogRequest struct {
	PageSize   int
	PageNumber int
	Keywords   string
	CreateBy   string
}

func (req *QueryBlogRequest) Limit() int {
	return req.PageSize
}

// Offset
// 1,  0
// 2, 20,
// 3, 20 * 2
// 4, 20 * 3
func (req *QueryBlogRequest) Offset() int {
	return req.PageSize * (req.PageNumber - 1)
}

func NewDescribeBlogRequest(id string) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		Id: id,
	}
}

type DescribeBlogRequest struct {
	Id string
}

func NewUpdateBlogRequest(id string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		Id:                id,
		UpdateMode:        common.UPDATE_MODE_PUT,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

type UpdateBlogRequest struct {
	Id         string            `json:"id"`
	UpdateMode common.UpdateMode `json:"update_mode"`
	*CreateBlogRequest
}

func NewDeleteBlogRequest(id string) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		Id: id,
	}
}

type DeleteBlogRequest struct {
	Id string
}
