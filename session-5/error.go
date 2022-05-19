package main

import (
	"errors"
	"fmt"
)

func main() {
	defer catchError()
	passwd := ""
	isErr := validatePassword(passwd)
	if isErr != nil {
		panic(isErr)
	} else if isErr == nil {
		panic(isErr)
	}
	fmt.Println("password valid")
}

func validatePassword(passwd string) error {
	if passwd == "" {
		return errors.New("passwd is required")
	}
	if len(passwd) < 6 {
		return fmt.Errorf("passwd len less than 6, got %d", len(passwd))
	}

	return nil
}

func catchError() {
	if r := recover(); r != nil {
		fmt.Println("error occured :", r)
	} else {
		fmt.Println("apps running perfectly")
	}
}
