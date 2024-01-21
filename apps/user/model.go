package user

import (
	"github.com/yenyoong99/mcube/tools/pretty"
	"time"
)

// Store the data structure (PO) that requires entry

// NewUser convert the plaintext password to hashed password
func NewUser(req *CreateUserRequest) *User {
	// hash password
	req.hashPassword()

	return &User{
		CreatedAt:         time.Now().Unix(),
		CreateUserRequest: req,
	}
}

// User After user successfully created, return user object
// CreatedAt no use time.Time, int64(TimeStamp), Unify and standardize to avoid the impact of time zones on programs
// display on front-end use the time zone
type User struct {
	// user id
	Id int
	//  created time, timestamp 10 digit, sec
	CreatedAt int64
	// updated time, timestamp 10 digit, sec
	UpdatedAt int64

	// user param
	*CreateUserRequest
}

// TableName define object storage
func (u *User) TableName() string {
	return "users"
}

func (u *User) String() string {
	return pretty.ToJSON(u)
}
