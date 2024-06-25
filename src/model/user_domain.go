package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/MateusGuedess/go-crud/src/configuration/rest_err"
)

func NewUserDomain(email, password, name string, age int8) *userDomain {
 return &userDomain{
	email, password, name, age,
 }
}

type userDomain struct {
	Email string 
	Password string 
	Name string 
	Age int8 
}

func (userDomain *userDomain) EncryptPassword() {
	// Encrypt password
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(userDomain.Password))
	userDomain.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*userDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}