package model

import (
	"fmt"

	"github.com/MateusGuedess/go-crud/src/configuration/rest_err"
)

func (userDomain *userDomain) UpdateUser(string) *rest_err.RestErr {
	// Update user
	fmt.Println("User updated: ", userDomain)
	return nil
}