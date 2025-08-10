package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	for i, rang := range os.Args {
		fmt.Println(i, rang)
	}
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Invalid command: %v\n", os.Args[1])
	}
}

func hash(password string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error hashing: %v", password)
		return
	}
	fmt.Println(string(hashedBytes))
}

func compare(password, hash string) {
	fmt.Printf("Todo: compare the password %q with the hash %q\n", password, hash)
}