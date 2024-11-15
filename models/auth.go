package models

const (
	RoleViewer  = 1 << 0 // 1
	RoleEditor  = 1 << 1 // 2
	RoleManager = 1 << 2 // 4
	RoleAdmin   = 1 << 3 // 8
)

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}