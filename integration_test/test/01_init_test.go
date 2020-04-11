package test01async

import (
	"testing"

	s "github.com/bluesky2106/eWallet-backend/integration_test/servers"
	"github.com/stretchr/testify/assert"
)

func TestInitConfigTest(t *testing.T) {
	assert := assert.New(t)

	testSrv := s.NewTestServer()
	testData := s.GetTestingData()

	assert.NotNil(testSrv)
	assert.NotNil(testData)
}
