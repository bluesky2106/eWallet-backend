package servers

import "github.com/bluesky2106/eWallet-backend/libs/mysql"

// User : struct
type User struct {
	adminDAO *mysql.DAO
}

// NewUser : new user server
func NewUser(adminDAO *mysql.DAO) *User {
	return &User{
		adminDAO: adminDAO,
	}
}
