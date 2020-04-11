package test

import (
	"fmt"
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
		fmt.Printf("%+v", user)
	}
}

func clearAdminUsers() error {
	return testSrv.DAOBO.DeleteByQuery(&boModels.User{}, map[string]interface{}{
		"user_name like ?": data.TestPrefix + "%",
	})
}
