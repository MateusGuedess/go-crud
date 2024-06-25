package service

import (
	"github.com/MateusGuedess/go-crud/src/configuration/rest_err"
	"github.com/MateusGuedess/go-crud/src/model"
)


func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {

}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}