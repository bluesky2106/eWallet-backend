package servers

import (
	"fmt"

	"github.com/bluesky2106/eWallet-backend/bo_entry_store/models"
)

var (
	testData *TestData
)

// TestData : test data
type TestData struct {
	AdminUsers models.Users
}

// GetTestingData : get testing data
func GetTestingData() *TestData {
	return testData
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

func initTestData() {
	testData = new(TestData)
	testData.initTestAdminUsers()

	fmt.Printf("Testing data: %+v", testData)
}
