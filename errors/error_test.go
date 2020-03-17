package errors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func connectToMongo(successful bool) error {
	if !successful {
		return New(ECMongoConnection)
	}
	return nil
}

// fake login : returns true if and only if username and password are admin
func login(username, password string) error {
	successful := false
	if username == "admin" && password == "admin" {
		successful = true
	}
	err := connectToMongo(successful)
	if err != nil {
		return WithMessage(err, "login")
	}
	return nil
}

func TestNew(t *testing.T) {
	assert := assert.New(t)

	// rpc error: code = Code(900) desc = mongodb connection error
	err := New(ECMongoConnection)
	errStr := err.Error()

	assert.True(strings.Contains(errStr, EMMongoConnection))
}

func TestWithMessage(t *testing.T) {
	assert := assert.New(t)

	err := login("admin", "admin")
	assert.Nil(err)

	err = login("admin", "test")
	assert.NotNil(err)
	assert.Equal(err, WithMessage(New(ECMongoConnection), "login"))
}

func TestFromError(t *testing.T) {
	assert := assert.New(t)

	// rpc error: code = Code(900) desc = mongodb connection error
	err := New(ECMongoConnection)
	actualErr := FromError(err)
	expectedErr := &Error{
		Code:    ECMongoConnection,
		Message: EMMongoConnection,
		Service: SrvUnknown,
	}
	assert.Equal(expectedErr, actualErr)
}

func TestGrpcEncode(t *testing.T) {
	assert := assert.New(t)

	err := &Error{
		Code:    ECMongoConnection,
		Message: EMMongoConnection,
		Service: "entry store service",
	}
	err.grpcEncode()

	assert.Equal(int32(ECMongoConnection), err.s.GetCode(), "Wrong error code")
	assert.Equal(EMMongoConnection, err.s.GetMessage(), "Wrong error message")
	assert.Equal("entry store service", err.s.GetDetails()[0].GetTypeUrl(), "Wrong error message")
}
