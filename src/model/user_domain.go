package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	EncryptPassword()
}

func NewUserDomain(email, password, name string, age int8) *userDomain {
 return &userDomain{
	email, password, name, age,
 }
}

type userDomain struct {
	email string 
	password string 
	name string 
	age int8 
}

func (userDomain *userDomain) GetEmail() string {
	return userDomain.email
}

func (userDomain *userDomain) GetPassword() string {
	return userDomain.password
}

func (userDomain *userDomain) GetName() string {
	return userDomain.name
}

func (userDomain *userDomain) GetAge() int8 {
	return userDomain.age
}

func (userDomain *userDomain) EncryptPassword() {
	// Encrypt password
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(userDomain.password))
	userDomain.password = hex.EncodeToString(hash.Sum(nil))
}

