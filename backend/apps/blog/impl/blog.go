package impl

import (
	"context"

	"dario.cat/mergo"
	"github.com/yenyoong99/goProjects_yyblog/common"

	"github.com/yenyoong99/goProjects_yyblog/apps/blog"
	"github.com/yenyoong99/goProjects_yyblog/exception"
)

// CreateBlog Create blog
func (i *blogServiceImpl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.ErrBadRequest.WithMessagef("Creation Failed, %s", err)
	}

	ins := blog.NewBlog(req)

	err := i.db.WithContext(ctx).Create(ins).Error
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// QueryBlog Query Blog
func (i *blogServiceImpl) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	set := blog.NewBlogSet()

	query := i.db.WithContext(ctx).Model(blog.Blog{})

	if in.CreateBy != "" {
		query = query.Where("create_by = ?", in.CreateBy)
	}

	if in.Keywords != "" {
		query = query.Where("title like ?", "%"+in.Keywords+"%")
	}

	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.
		Limit(in.Limit()).
		Offset(in.Offset()).
		Find(&set.Items).
		Error

	if err != nil {
		return nil, err
	}

	return set, nil
}

// DescribeBlog Describe Blog Details
func (i *blogServiceImpl) DescribeBlog(ctx context.Context, in *blog.DescribeBlogRequest) (*blog.Blog, error) {
	query := i.db.WithContext(ctx).Model(&blog.Blog{}).Where("id = ?", in.Id)

	ins := blog.NewBlog(blog.NewCreateBlogRequest())
	if err := query.First(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *blogServiceImpl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) (*blog.Blog, error) {
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.Id))
	if err != nil {
		return nil, err
	}

	err = i.db.
		WithContext(ctx).
		Model(&blog.Blog{}).
		Where("id = ?", req.Id).
		Delete(ins).Error

	if err != nil {
		return nil, err
	}

	return ins, nil
}

// UpdateBlog update blog
func (i *blogServiceImpl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.Id))
	if err != nil {
		return nil, err
	}

	// object update
	switch req.UpdateMode {
	case common.UPDATE_MODE_PATCH:
		// if req.Author != "" {
		// 	ins.Author = req.Author
		// }
		// if req.Title != "" {
		// 	ins.Title = req.Title
		// }
		// https://github.com/darccio/mergo
		// // WithOverride will make merge override non-empty dst attributes with non-empty src attributes values.
		//
		if err := mergo.MapWithOverwrite(ins.CreateBlogRequest, req.CreateBlogRequest); err != nil {
			return nil, err
		}
	default:
		ins.CreateBlogRequest = req.CreateBlogRequest
	}

	if err := ins.Validate(); err != nil {
		return nil, exception.ErrBadRequest.WithMessagef("Verification update request failed: %s", err)
	}

	// update databases
	stmt := i.db.WithContext(ctx).Model(&blog.Blog{}).Where("id = ?", ins.Id)
	if req.CreateBy != "" {
		stmt = stmt.Where("create_by = ?", ins.CreateBy)
	}
	err = stmt.Updates(ins).Error

	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ChangedBlogStatus Article status amend, Ex: Published
func (i *blogServiceImpl) ChangedBlogStatus(ctx context.Context, req *blog.ChangedBlogStatusRequest) (*blog.Blog, error) {
	return nil, nil
}

// AuditBlog review
func (i *blogServiceImpl) AuditBlog(ctx context.Context, req *blog.AuditInfo) (*blog.Blog, error) {
	return nil, nil
}
