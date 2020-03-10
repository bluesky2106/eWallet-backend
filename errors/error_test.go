package errors

import (
	"fmt"
	"testing"
)

func connectToMongo(successful bool) error {
	if !successful {
		return New(ECMongoConnection)
	}
	return nil
}

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

func TestWithMessage(t *testing.T) {
	fmt.Println("Test Function WithMessage")
	SetService("backend")
	err := login("admin", "admin")
	if err != nil {
		t.Error(err)
	}

	err = login("admin", "test")
	if err == nil {
		t.Error(err)
	} else {
		// fmt.Println(err)
		e := FromError(err)
		fmt.Println(e)
	}
}
