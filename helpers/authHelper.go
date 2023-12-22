package helpers

import (
	"error"
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func CheckUserType(c *gin.Context, role string) (err error){
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context,userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uuid")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	err = CheckUserType(c,userType)
	return err
}