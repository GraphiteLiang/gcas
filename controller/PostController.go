package controller

import (
	"CashAAService/dao"
	"CashAAService/model"
	"CashAAService/protocol"
	"CashAAService/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)
//
// InputCashOutRecord
//  @Description: 添加出账记录
//  @param c
//
func InputCashOutRecord(c *gin.Context) {
	if err := util.Validate(c); err != nil {
		util.ProcessError(err, c)
		return
	}
	reqBytes, err := c.GetRawData()
	log.Println("参数" + string(reqBytes))
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	var reqMsg protocol.CashOutInputMessage
	err = json.Unmarshal(reqBytes, &reqMsg)
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	// 根据录入的消息，获取出账流水记录
	var fullStatementRecord = util.GetFullStatementRecordFromOutMsg(reqMsg)
	// 根据录入的消息，获取应付款记录
	var aaStatementRecords = util.GetAAStatementRecordsFromOutMsg(reqMsg)
	// 查询收款人为出账记录人的AA应付款记录
	err, aaStatementRecordsOrg := dao.GetAAStatementRecordByUserName(reqMsg.UserName)
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	// 合并应付款记录到一个AA应付款记录上
	aaStatementRecords = mergeAAStatementRecord(aaStatementRecords, aaStatementRecordsOrg)
	// 出账记录插入表
	if err = dao.InsertFullStatementRecord(fullStatementRecord); err != nil {
		util.ProcessError(err, c)
		return
	}
	// 更新AA记录表
	if err = dao.UpdateBatchAAStatementRecord(aaStatementRecords); err != nil {
		util.ProcessError(err, c)
		return
	}
	c.JSON(200, util.GetSuccessResponse())
}

//
// InputCashPay
//  @Description: 付款，把对应qid的statement的completed置为"Y"
//  @param c
//
func InputCashPay(c *gin.Context) {
	if err := util.Validate(c); err != nil {
		util.ProcessError(err, c)
		return
	}
	reqBytes, err := c.GetRawData()
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	var reqMsg protocol.CashPayInputMessage
	err = json.Unmarshal(reqBytes, &reqMsg)
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	// 查询收款人/付款人为出账记录人的AA应付款记录
	err, statementRecord := dao.GetAAStatementRecordByQid(reqMsg.StatementQid)
	if err != nil {
		util.ProcessError(err, c)
		return
	}
	var fullStatementRecord = util.GetFullStatementRecordFromPayMsg(reqMsg)
	if err = dao.InsertFullStatementRecord(fullStatementRecord); err != nil {
		util.ProcessError(err, c)
		return
	}
	statementRecord.Completed = "Y"
	// 更新AA记录表
	if err = dao.UpdateStatementOnCompleted(statementRecord); err != nil {
		util.ProcessError(err, c)
		return
	}
	c.JSON(200, util.GetSuccessResponse())
}

//
// mergeAAStatementRecord
//  @Description: 将record中的AA记录合并到recordsNew，他们的收款人应该是相同的
//  @param record
//  @param recordOrg
//
func mergeAAStatementRecord(recordsNew []model.AAStatementRecord, recordsOrg []model.AAStatementRecord) []model.AAStatementRecord {
	var tmpMap = make(map[string]map[string]int, len(recordsOrg))
	for i, record := range recordsOrg {
		if tmpMap[record.PayerUserName] == nil {
			tmpMap[record.PayerUserName] = make(map[string]int)
		}
		tmpMap[record.PayerUserName][record.ReceiverUserName] = i
	}
	for _, record := range recordsNew {
		var cashNew, _ = strconv.ParseFloat(record.CashToPay, 64)
		if orgIndex, ok := tmpMap[record.PayerUserName][record.ReceiverUserName];ok {
			var cashOrg, _ = strconv.ParseFloat(recordsOrg[orgIndex].CashToPay, 64)
			var cashCur = cashOrg + cashNew
			recordsOrg[orgIndex].CashToPay = fmt.Sprintf("%2f", cashCur)
		} else if orgIndex, ok = tmpMap[record.ReceiverUserName][record.PayerUserName];ok {
			var cashOrg, _ = strconv.ParseFloat(recordsOrg[orgIndex].CashToPay, 64)
			var cashCur = cashOrg - cashNew
			if cashCur > 0.01 {
				recordsOrg[orgIndex].CashToPay = fmt.Sprintf("%2f", cashCur)
			} else if cashCur < -0.01 {
				// 将原来的标记为已付款，新增一条待付款
				recordsOrg[orgIndex].Completed = "Y"
				record.CashToPay = fmt.Sprintf("%2f", -cashCur)
				recordsOrg = append(recordsOrg, record)
			} else {
				recordsOrg[orgIndex].Completed = "Y"
			}
		} else {
			recordsOrg = append(recordsOrg, record)
		}
	}
	return recordsOrg
}
