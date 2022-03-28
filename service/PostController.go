package service

import (
	"CashAAService/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

// GetAllCreditOut
//  @Description: 根据用户，获取所有该用户的应付款项
//  @param w
//  @param r
func GetAllCreditOut(c *gin.Context) {
	var msg = model.CreditPayableMessage{}
	var req = model.CreditPayableReqMsg{}
	reqByte, err := c.GetRawData()
	err = json.Unmarshal(reqByte, &req)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(req)
	msg.UserName = "graphite"
	msg.CreditPayableItems = make([]model.CreditPayableItem, 0, 1000)
	msg.CreditPayableItems = append(msg.CreditPayableItems, model.CreditPayableItem{"zcs", "60"})
	msg.CreditPayableItems = append(msg.CreditPayableItems, model.CreditPayableItem{"pxl", "60"})

	msgByte, err := json.Marshal(msg)
	msgStr := string(msgByte)
	log.Println(msgStr)
	if err != nil {
		log.Println("消息头获取失败！")
		return
	}
	c.JSON(200, msg)
}
