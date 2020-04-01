package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// ReadPassword : read string from stdio terminal
func ReadPassword() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	text, _ := reader.ReadString('\n')
	return text[:len(text)-1]
}

func main() {
	for {
		pwd := ReadPassword()
		// fmt.Println(pwd)

		hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
		if err != nil {
			fmt.Errorf("Cannot hash password: %s", err.Error())
		}
		fmt.Println(string(hashed))

		// err = bcrypt.CompareHashAndPassword(hashed, []byte("1234"))
		// if err != nil {
		// 	fmt.Println("Password is not 1234")
		// }
	}
}
