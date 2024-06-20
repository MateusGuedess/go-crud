package controller

import (
	"fmt"

	"github.com/MateusGuedess/go-crud/src/configuration/rest_err"
	"github.com/MateusGuedess/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error()), nil)
	
	
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)

}

