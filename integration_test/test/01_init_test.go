package test

import (
	"testing"

	"github.com/bluesky2106/eWallet-backend/integration_test/data"
	"github.com/bluesky2106/eWallet-backend/integration_test/servers"
	s "github.com/bluesky2106/eWallet-backend/integration_test/servers"
	"github.com/stretchr/testify/assert"
)

var (
	testSrv  *servers.TestServer
	testData *data.TestData
)

func TestInitTestServer(t *testing.T) {
	assert := assert.New(t)

	testSrv = s.NewTestServer()
	assert.NotNil(testSrv)
	assert.NotNil(testSrv.Conf)
	assert.NotNil(testSrv.DAO)
	assert.NotNil(testSrv.DAOBO)
}

func TestInitTestData(t *testing.T) {
	assert := assert.New(t)

	testData = data.GetTestData()
	assert.NotNil(testData)
	assert.NotNil(testData.AdminUsers)
	assert.Equal(len(testData.AdminUsers), data.NumberOfAdminUsers)
}
