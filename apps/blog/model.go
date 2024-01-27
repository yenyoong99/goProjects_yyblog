package blog

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/tools/pretty"
)

var (
	v = validator.New()
)

func NewBlog(req *CreateBlogRequest) *Blog {
	return &Blog{
		CreateBlogRequest:        req,
		ChangedBlogStatusRequest: NewChangedBlogStatusRequest(),
		AuditInfo:                NewAuditInfo(),
	}
}

type Blog struct {
	// user id
	Id        int   `json:"id" gorm:"column:id"`
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	*CreateBlogRequest
	*ChangedBlogStatusRequest
	*AuditInfo
}

func (u *Blog) String() string {
	return pretty.ToJSON(u)
}

func NewChangedBlogStatusRequest() *ChangedBlogStatusRequest {
	return &ChangedBlogStatusRequest{
		Status: STATUS_DRAFT,
	}
}

type ChangedBlogStatusRequest struct {
	PublishedAt int64 `json:"published_at" gorm:"column:published_at"`
	// article status: Draft/Published
	Status Status `json:"status" gorm:"column:status"`
}

func NewAuditInfo() *AuditInfo {
	return &AuditInfo{}
}

type AuditInfo struct {
	AuditAt     int64 `json:"audit_at" gorm:"column:audit_at"`
	IsAuditPass bool  `json:"is_audit_pass" gorm:"column:is_audit_pass"`
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags: map[string]string{},
	}
}

type CreateBlogRequest struct {
	Title    string            `json:"title" gorm:"column:title" validate:"required"`
	Author   string            `json:"author" gorm:"column:author" validate:"required"`
	Content  string            `json:"content" gorm:"column:content" validate:"required"`
	Summary  string            `json:"summary" gorm:"column:summary"`
	CreateBy string            `json:"create_by" gorm:"column:create_by"`
	Tags     map[string]string `json:"tags" gorm:"column:tags;serializer:json"`
}

func (req *CreateBlogRequest) Validate() error {
	return v.Struct(req)
}

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

type BlogSet struct {
	Total int64   `json:"total"`
	Items []*Blog `json:"items"`
}

func (u *BlogSet) String() string {
	return pretty.ToJSON(u)
}
