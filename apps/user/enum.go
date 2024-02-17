package user

type Role int

// ROLE_MEMBER/ROLE_ADMIN
const (
	RoleVisitor Role = iota
	RoleAdmin
)
