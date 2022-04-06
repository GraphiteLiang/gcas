package util

import (
	"CashAAService/cache"
	"CashAAService/constants"
	"CashAAService/model"
	"CashAAService/protocol"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func ProcessError(err error, c *gin.Context) {
	if err != nil {
		log.Println(err)
		c.JSON(400, GetErrorResponse(err))
	}
}

func GetGcasError(err error, ErrorCode string, ErrorMsg string) model.GcasError {
	return model.GcasError{Error: err, ErrorCode: ErrorCode, ErrorMsg: ErrorMsg}
}

func GetErrorResponse(err error) protocol.ErrorMessage{
	return protocol.ErrorMessage{ErrorCode: "-1000",ErrorMessage: err.Error()}
}

func GetSuccessResponse() protocol.ErrorMessage{
	return protocol.ErrorMessage{ErrorCode: "0",ErrorMessage: ""}
}

func Validate(c *gin.Context) error {
	var userName, token = c.Query("userName"), c.Query("token")
	if cache.CheckToken(userName, token) {
		return nil
	} else {
		err := errors.New(fmt.Sprintf("登录信息超时，请重新登陆,userName=%s", userName))
		c.JSON(400, GetErrorResponse(err))
		return err
	}
}

func Calculate(a string, b string, calType int) string {
	af, _ := strconv.ParseFloat(a, 64)
	bf, _ := strconv.ParseFloat(b, 64)
	var resf float64
	switch calType {
	case constants.Plus:
		resf = af + bf
	case constants.Minus:
		resf = af - bf
	case constants.Multiply:
		resf = af * bf
	case constants.Division:
		resf = af / bf
	}
	return fmt.Sprintf("%.2f", resf)
}


func GetStringVal(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}