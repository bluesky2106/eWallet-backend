package main

import (
	"fmt"

	"github.com/pkg/errors"

	_ "github.com/bluesky2106/eWallet-backend/config"
)

func main() {
	// conf := config.ParseConfig("config.json", "config")
	// conf.Print()

	err1 := errors.New("Error reason 1")
	err1 = errors.WithMessage(err1, "Error 1")
	fmt.Println("err1 :", err1)

	err2 := errors.New("Error reason 2")
	err2 = errors.WithMessage(err2, "Error 2")
	fmt.Println("err2 :", err2)

	err3 := errors.New("Error reason 3")
	err3 = errors.Wrap(err3, "Error 3")
	fmt.Println("err3 :", err3)

	err3 = errors.Wrap(err3, "Error 4")
	fmt.Println("err3 :", err3)

	err3 = errors.WithMessage(err3, "Error 5")
	fmt.Println("err3 :", err3)
}
