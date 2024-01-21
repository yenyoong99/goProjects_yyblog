package user

// Store the data structure (PO) that requires entry

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
