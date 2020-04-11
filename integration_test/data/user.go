package data

// User : user testing data
type User struct {
	ID       uint64
	FullName string
	UserName string
	Email    string
	Token    string
}

// Users : slice of User
type Users []*User
