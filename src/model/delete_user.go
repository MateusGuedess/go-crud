package model

import (
	"fmt"

	"github.com/MateusGuedess/go-crud/src/configuration/rest_err"
)

func (userDomain *userDomain) DeleteUser(string) *rest_err.RestErr {
	//Delete user
	fmt.Println("User deleted: ", userDomain)
	return nil
}