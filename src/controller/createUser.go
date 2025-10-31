package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gustavoeguedes/api-hun-coding/src/configuration/rest_err"
	"github.com/gustavoeguedes/api-hun-coding/src/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s", err.Error())
		errRest := rest_err.NewBadRequestError("Some fields are incorrect")
		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}