package controller

import (
	"CashAAService/cache"
	"CashAAService/dao"
	"CashAAService/protocol"
	"CashAAService/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

// GetCreditStatementRelated
//  @Description: 根据用户，获取所有该用户的应付款项
func GetCreditStatementRelated(c *gin.Context) {
	var err = util.Validate(c)
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	var resMsg = protocol.CreditPayableMessage{}
	var userName = c.Query("userName")
	log.Printf("收到对应用户%s的查询", userName)
	err, records := dao.GetAAStatementRecordByUserName(userName)
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	resMsg.UserName = userName
	resMsg.UserNickName = cache.GetUserNickName(userName)
	resMsg.CreditPayableItems = make([]protocol.CreditPayableItem, len(records))
	for i, record := range records {
		if record.PayerUserName == userName {
			resMsg.CreditPayableItems[i] = protocol.CreditPayableItem{
				DbQid: record.Qid,
				ItemType:     "1",
				UserName : record.ReceiverUserName,
				UserNickName: cache.GetUserNickName(record.ReceiverUserName),
				Cash:         record.CashToPay,
			Remark: record.Remark.String}
		} else {
			resMsg.CreditPayableItems[i] = protocol.CreditPayableItem{
				DbQid: record.Qid,
				ItemType:     "2",
				UserName : record.PayerUserName,
				UserNickName: cache.GetUserNickName(record.PayerUserName),
				Cash:         record.CashToPay,
				Remark: record.Remark.String}
		}
	}
	msgByte, _ := json.Marshal(resMsg)
	msgStr := string(msgByte)
	log.Println(msgStr)
	c.JSON(200, resMsg)
}
//
// GetFullStatement
//  @Description: 获取所有流水记录
//  @param c
//
func GetFullStatement(c *gin.Context) {
	var err = util.Validate(c)
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	var resMsg = protocol.StatementMessage{}
	err, records := dao.GetFullStatementRecord()
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	resMsg.StatementList = make([]protocol.Statement, len(records))
	for i, record := range records {
		resMsg.StatementList[i] = protocol.Statement{
			UserName:      record.UserName,
			UserNickName:  "从缓存中取或者sql语句写一个关联查找或者使用视图",
			CashOut:       record.CashOut,
			AAPeopleCount: record.AAPeopleCount,
			Type:          record.Type,
			Remark:        record.Remark.String}
	}
	msgByte, _ := json.Marshal(resMsg)
	msgStr := string(msgByte)
	log.Println(msgStr)
	c.JSON(200, resMsg)
}

func GetAllUser(c *gin.Context) {
	var err = util.Validate(c)
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	var resMsg = protocol.UserQueryResMessage{}
	records, err := dao.GetAllUser()
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	resMsg.UserInfos = make([]protocol.UserInfo, len(records))
	for i, record := range records {
		resMsg.UserInfos[i] = protocol.UserInfo{
			UserName:     record.UserName,
			UserNickName: record.UserNickname.String,
		}
	}
	msgByte, _ := json.Marshal(resMsg)
	msgStr := string(msgByte)
	log.Println(msgStr)
	c.JSON(200, resMsg)
}
func Test(c *gin.Context) {
	c.JSON(200, protocol.CreditPayableItem{
		ItemType:     "2",
		UserNickName: "测试数据",
		Cash:         "测试数据"})
}

