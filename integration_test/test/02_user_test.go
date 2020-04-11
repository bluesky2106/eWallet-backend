package test

import (
	"testing"

	"github.com/bluesky2106/eWallet-backend/bo_controller/serializers"
	boModels "github.com/bluesky2106/eWallet-backend/bo_entry_store/models"
	"github.com/bluesky2106/eWallet-backend/integration_test/data"
	"github.com/stretchr/testify/assert"
)

func TestAdminUserRegister(t *testing.T) {
	assert := assert.New(t)

	for i := 0; i < data.NumberOfAdminUsers; i++ {
		req := &serializers.UserRegisterReq{
			Email:           testData.AdminUsers[i].Email,
			FullName:        testData.AdminUsers[i].FullName,
			Password:        data.RegisterPwd,
			ConfirmPassword: data.RegisterPwd,
		}
		user, err := testSrv.UserSrv.AdminUserRegister(req)
		assert.Nil(err)
		assert.NotNil(user)
		// fmt.Printf("%+v\n\n", user)
		testData.AdminUsers[i].ID = user.ID
	}
}

func TestAdminUserLogin(t *testing.T) {
	assert := assert.New(t)

	for i := 0; i < data.NumberOfAdminUsers; i++ {
		req := &serializers.UserLoginReq{
			Email:    testData.AdminUsers[i].Email,
			Password: data.RegisterPwd,
		}
		resp, err := testSrv.UserSrv.AdminUserLogin(req)
		assert.Nil(err)
		assert.NotNil(resp)
		// fmt.Printf("%+v\n\n", resp)
		testData.AdminUsers[i].Token = resp.Token
	}
}

func TestAdminUserProfile(t *testing.T) {
	assert := assert.New(t)

	for i := 0; i < data.NumberOfAdminUsers; i++ {
		user, err := testSrv.UserSrv.AdminUserProfile(testData.AdminUsers[i].Token)
		assert.Nil(err)
		assert.NotNil(user)
		assert.Equal(testData.AdminUsers[i].ID, user.ID, "ID mismatched !")
		assert.Equal(testData.AdminUsers[i].Email, user.Email, "Email mismatched !")
		assert.Equal(testData.AdminUsers[i].FullName, user.FullName, "Fullname mismatched !")
		assert.Equal(testData.AdminUsers[i].UserName, user.UserName, "Username mismatched !")
		// fmt.Printf("%+v\n\n", user)
	}
}

func clearAdminUsers() error {
	return testSrv.DAOBO.DeleteByQuery(&boModels.User{}, map[string]interface{}{
		"user_name like ?": data.TestPrefix + "%",
	})
}
