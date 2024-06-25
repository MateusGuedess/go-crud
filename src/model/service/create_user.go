package service

import (
	"fmt"

	"github.com/MateusGuedess/go-crud/src/configuration/logger"
	"github.com/MateusGuedess/go-crud/src/configuration/rest_err"
	"github.com/MateusGuedess/go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println("User created: ", userDomain.GetPassword())

	return nil
}