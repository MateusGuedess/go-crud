package service

import (
	"github.com/MateusGuedess/go-crud/src/configuration/rest_err"
	"github.com/MateusGuedess/go-crud/src/model"
)

func (*userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	// Update user
	// fmt.Println("User updated: ", userDomain)
	return nil
}