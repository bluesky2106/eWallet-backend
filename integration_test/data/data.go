package data

import (
	"fmt"

	"github.com/bluesky2106/eWallet-backend/bo_entry_store/models"
)

// TestData : test data
type TestData struct {
	AdminUsers models.Users
}

// NewTestData : new test data
func NewTestData() *TestData {
	td := new(TestData)
	td.init()
	return td
}

func (td *TestData) init() {
	td.initTestAdminUsers()
}

func (td *TestData) initTestAdminUsers() {
	td.AdminUsers = make(models.Users, NumberOfAdminUsers)
	for i := 0; i < NumberOfAdminUsers; i++ {
		adminName := fmt.Sprintf("%sadmin_%d", TestPrefix, i)
		adminEmail := adminName + EmailSuffix
		td.AdminUsers[i] = &models.User{
			UserName: adminName,
			Email:    adminEmail,
		}
	}
}
