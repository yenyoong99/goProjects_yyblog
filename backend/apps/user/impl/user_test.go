package impl_test

import (
	"context"
	"testing"

	"github.com/yenyoong99/goProjects_yyblog/apps/user"
	"github.com/yenyoong99/goProjects_yyblog/apps/user/impl"
)

var (
	i   user.Service
	ctx = context.Background()
)

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "guest"
	req.Password = "123456"
	req.Role = user.RoleAdmin

	u, err := i.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	if u == nil {
		t.Fatal("user not created")
	}

	t.Log(u)
}

func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	ul, err := i.QueryUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ul)
}

func TestDescribeUser(t *testing.T) {
	req := user.NewDescribeUserRequest("8")
	ul, err := i.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ul)

	t.Log(ul.CheckPassword("123456"))
}

func init() {
	i = impl.NewUserServiceImpl()
}
