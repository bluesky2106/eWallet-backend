package data

import (
	"fmt"
)

var td *TestData

// TestData : test data
type TestData struct {
	AdminUsers Users
}

// GetTestData : create new test data if not exist
func GetTestData() *TestData {
	if td == nil {
		td = new(TestData)
		td.init()
	}

	return td
}

func (td *TestData) init() {
	td.initTestAdminUsers()
}

func (td *TestData) initTestAdminUsers() {
	td.AdminUsers = make(Users, NumberOfAdminUsers)
	for i := 0; i < NumberOfAdminUsers; i++ {
		adminName := fmt.Sprintf("%sadmin_%d", TestPrefix, i+1)
		adminEmail := adminName + EmailSuffix
		td.AdminUsers[i] = &User{
			UserName: adminName,
			FullName: adminName,
			Email:    adminEmail,
		}
	}
}
