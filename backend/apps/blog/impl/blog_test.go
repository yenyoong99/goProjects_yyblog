package impl_test

import (
	"testing"

	"github.com/yenyoong99/goProjects_yyblog/apps/blog"
	"github.com/yenyoong99/goProjects_yyblog/common"
)

func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.Title = "first yyblog"
	req.Author = "yy"
	req.Content = "yyblog xxxx"
	req.Summary = "xxxx"
	req.Tags["Menu"] = "Default"
	ins, err := impl.CreateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryBlog(t *testing.T) {
	req := blog.NewQueryBlogRequest()
	set, err := impl.QueryBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDescribeBlog(t *testing.T) {
	req := blog.NewDescribeBlogRequest("48")
	set, err := impl.DescribeBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

// req.Title = "first yyblog"
// req.Author = "yy"
// req.Content = "yyblog xxxx"
// req.Summary = "xxxx"
// req.Tags["Menu"] = "Default"
func TestUpdateBlogPatchMod(t *testing.T) {
	req := blog.NewUpdateBlogRequest("48")
	req.UpdateMode = common.UPDATE_MODE_PATCH
	req.Title = "first yyblog v2"
	set, err := impl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestUpdateBlogPutMod(t *testing.T) {
	req := blog.NewUpdateBlogRequest("48")
	req.Title = "first yyblog v3"
	req.Content = "v3"
	req.Author = "v3"
	set, err := impl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

//func TestDeleteBlog(t *testing.T) {
//	req := blog.NewDeleteBlogRequest("47")
//	set, err := impl.DeleteBlog(ctx, req)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log(set)
//}
