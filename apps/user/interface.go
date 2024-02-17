package user

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/tools/pretty"
	"golang.org/x/crypto/bcrypt"
)

const (
	// AppName module name
	AppName = "users"
)

var (
	v = validator.New()
)

// Service object-oriented
// user.Service
// define interface, Considering compatibility, the parameters of the interface cannot be changed.
type Service interface {
	// CreateUser
	// (username, password, role string, lable map[string]string)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// QueryUser Query user list, object list [{}]
	QueryUser(context.Context, *QueryUserRequest) (*UserSet, error)
	// DescribeUser Query user details, use id check
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
	AlterUser(context.Context, *AlterUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) error
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Role:  RoleVisitor,
		Label: map[string]string{},
	}
}

// CreateUserRequest create user param
type CreateUserRequest struct {
	Username string `json:"username" validate:"required" gorm:"column:username"`
	Password string `json:"password" validate:"required" gorm:"column:password"`
	Role     Role   `json:"role" gorm:"column:role"`
	// https://gorm.io/docs/serializer.html
	Label map[string]string `json:"label" gorm:"column:label;serializer:json"`
}

func (req *CreateUserRequest) hashPassword() {
	hp, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	req.Password = string(hp)

}

func (req *CreateUserRequest) CheckPassword(pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(pass))
}

func (req *CreateUserRequest) Validate() error {
	//if req.Username = "" {
	//	return fmt.Errorf("username is required")
	//}
	//
	//if req.Password = "" {
	//	return fmt.Errorf("password is required")
	//}

	// validator library validator.New()
	return v.Struct(req)
}

func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		PageSize:   20,
		PageNumber: 1,
	}

}

// QueryUserRequest Query user list
type QueryUserRequest struct {
	// Page size
	PageSize int
	// Current page
	PageNumber int
	// Find user by changing user name
	Username string
}

func (req *QueryUserRequest) Limit() int {
	return req.PageSize
}

func (req *QueryUserRequest) Offset() int {
	return req.PageSize * (req.PageNumber - 1)

}

func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

type UserSet struct {
	// How many there in total
	Total int64 `json:"total"`
	// List of data currently queried
	Items []*User `json:"items"`
}

func (u *UserSet) String() string {
	return pretty.ToJSON(u)
}

func NewDescribeUserRequest(uid string) *DescribeUserRequest {
	return &DescribeUserRequest{
		UserId: uid,
	}

}

type DescribeUserRequest struct {
	UserId string
}

// alter user func

// AlterUserRequest modify the user data
type AlterUserRequest struct {
	*DescribeUserRequest
	Password string
	Role     string
	Label    map[string]string
	UpdateAt string
}

// delete user func

// DeleteUserRequest delete user by id
type DeleteUserRequest struct {
	UserId int
}
