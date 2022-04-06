package controller

import (
	"CashAAService/cache"
	"CashAAService/dao"
	"CashAAService/protocol"
	"CashAAService/util"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"strconv"
	"time"
)

func Login(c *gin.Context) {
	reqBytes, err := c.GetRawData()
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	var reqMsg protocol.LoginMessage
	var resMsg protocol.LoginResMessage
	err = json.Unmarshal(reqBytes, &reqMsg)
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	if dao.ValidateUser(reqMsg.UserName, reqMsg.PassWord) {
		resMsg.Token = GetToken(reqMsg.UserName, reqMsg.PassWord)
		resMsg.UserName = reqMsg.UserName
		cache.PutToken(reqMsg.UserName, resMsg.Token)
		c.JSON(200, resMsg)
	} else {
		log.Printf("登陆失败！%s-%s", reqMsg.UserName, reqMsg.PassWord)
		c.JSON(400, util.GetErrorResponse(errors.New("登录失败！")))
	}
}

func GetToken(userName string, password string) string {
	curTime := time.Now().Unix()
	h := md5.New()
	tokenStr := userName + password + strconv.FormatInt(curTime, 10)
	io.WriteString(h, tokenStr)
	token := fmt.Sprintf("%x", h.Sum(nil))

	return token
}
