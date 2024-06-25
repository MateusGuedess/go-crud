package model

import (
	"fmt"

	"github.com/MateusGuedess/go-crud/src/configuration/logger"
	"github.com/MateusGuedess/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (userDomain *userDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println("User created: ", userDomain)

	return nil
}