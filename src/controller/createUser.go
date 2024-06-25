package controller

import (
	"net/http"

	"github.com/MateusGuedess/go-crud/src/configuration/logger"
	"github.com/MateusGuedess/go-crud/src/configuration/validation"
	"github.com/MateusGuedess/go-crud/src/controller/model/request"
	"github.com/MateusGuedess/go-crud/src/controller/model/response" // Import the missing package
	"github.com/MateusGuedess/go-crud/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context){
	var userRequest request.UserRequest
	logger.Info("Init CreateUser controller...", zap.String("journey", "createUser"))
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		errRest :=  validation.ValidateUserError(err)
	
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	
	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created succesfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)


}

