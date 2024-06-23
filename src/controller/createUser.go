package controller

import (
	"fmt"

	"github.com/MateusGuedess/go-crud/src/configuration/validation"
	"github.com/MateusGuedess/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest :=  validation.ValidateUserError(err)
	
	
		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)

}

