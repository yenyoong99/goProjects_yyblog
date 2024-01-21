package user

type Role int

// ROLE_MEMBER/ROLE_ADMIN
const (
	RoleMember Role = iota
	RoleAdmin
)
