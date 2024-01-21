package impl

import (
	"goProjects/yyblog/apps/user"

	"context"
)

// user.Service
// &UserServiceImpl{} memory will be allocated. How can we not allocate memory?
// nil *UserServiceImpl
// (*UserServiceImpl)(nil) ---> int8 1  int32(1)  (int32)(1)
// nil is *UserServiceImpl null pointer
var _ user.Service = (*UserServiceImpl)(nil)

func (i *UserServiceImpl) CreateUser(
	ctx context.Context,
	in *user.CreateUserRequest) (
	*user.User, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	// create user instance object
	u := user.NewUser(in)
	// Persist the object (put it in the database)
	// object --> row
	// INSERT INTO `users` (`created_at`,`updated_at`,`username`,`password`,`role`,`label`) VALUES (1705130308,1705130308,'admin','123456','1','{}') RETURNING `id`
	// For example, in create user, if the request has not returned after 4 seconds, the user cancels the request. Will the backend end because the request exits?
	// The program does not have the ability to interrupt database operations, and ctx is carried through WithContext.
	if err := i.db.WithContext(ctx).Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil

}

// User service implement

// query the user list
// check multiple data in databases

func (i *UserServiceImpl) QueryUser(
	ctx context.Context,
	in *user.QueryUserRequest) (
	*user.UserSet, error) {
	// Construct a mysql conditional query statement  select * from users where ....
	query := i.db.WithContext(ctx).Model(&user.User{})

	// Construct conditional where username = ""
	if in.Username != "" {
		// query Generates a new statement without amend the query object itself
		query = query.Where("username = ?", in.Username)
	}

	set := user.NewUserSet()

	// Count
	// select COUNT(*) from users where ....
	err := query.
		Count(&set.Total).
		Error
	if err != nil {
		return nil, err
	}

	// page query: sql  LIMIT offset,limit
	// limit 20,20 query the second page
	// use Find to query multiple rows of data, use []User to return
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

// DescribeUser query user detail, use user id check
func (i *UserServiceImpl) DescribeUser(
	ctx context.Context,
	in *user.DescribeUserRequest) (
	*user.User, error) {
	query := i.db.WithContext(ctx).Model(&user.User{}).Where("id = ?", in.UserId)

	u := user.NewUser(user.NewCreateUserRequest())
	if err := query.First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (i *UserServiceImpl) AlterUser(
	ctx context.Context,
	in *user.AlterUserRequest) (
	*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (i *UserServiceImpl) DeleteUser(
	ctx context.Context,
	in *user.DeleteUserRequest) error {
	//TODO implement me
	panic("implement me")
}
