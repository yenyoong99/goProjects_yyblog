package user_test

import (
	"testing"

	"github.com/yenyoong99/goProjects_yyblog/apps/user"
)

// $2a$10$1MvkjvWOS0/Rf.cEKKxeie/Y7ADz9XZTq09Wd/bKwX/vUv0kdYJ4.
// $2a$10$IyB.w1NVOrBmZ9WOsT6gEuruaynjse2CNmce9399yUErnufV10DX2
// https://gitee.com/infraboard/go-course/blob/master/day09/go-hash.md#bcrypt
func TestHashPassword(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Password = "123456"
	u := user.NewUser(req)
	t.Log(u.Password)

	t.Log(u.CheckPassword("123456"))
}
