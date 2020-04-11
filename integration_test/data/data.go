package data

import (
	"fmt"

	boModels "github.com/bluesky2106/eWallet-backend/bo_entry_store/models"
)

// TestData : test data
type TestData struct {
	AdminUsers boModels.Users
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
	td.AdminUsers = make(boModels.Users, NumberOfAdminUsers)
	for i := 0; i < NumberOfAdminUsers; i++ {
		adminName := fmt.Sprintf("%sadmin_%d", TestPrefix, i+1)
		adminEmail := adminName + EmailSuffix
		td.AdminUsers[i] = &boModels.User{
			UserName: adminName,
			FullName: adminName,
			Email:    adminEmail,
		}
	}
}
