package util

import (
	"CashAAService/constants"
	"CashAAService/dao"
	"CashAAService/model"
	"CashAAService/protocol"
	"database/sql"
	"fmt"
	"strconv"
)

func GetFullStatementRecordFromOutMsg(msg protocol.CashOutInputMessage) model.FullStatementRecord {
	var qid, _ = dao.GetSequenceNextVal()
	return model.FullStatementRecord{Qid: qid,
		UserName:      msg.UserName,
		CashOut:       msg.CashOut,
		AAPeopleCount: msg.AAPeopleCount,
		Type:          strconv.Itoa(constants.TypeCashOut),
		Remark:        sql.NullString{String: msg.Remark}}
}

func GetAAStatementRecordsFromOutMsg(msg protocol.CashOutInputMessage) []model.AAStatementRecord {
	aaCount, _ := strconv.Atoi(msg.AAPeopleCount)
	cashAA := Calculate(msg.CashOut, msg.AAPeopleCount, constants.Division)
	var res = make([]model.AAStatementRecord, 0, aaCount)
	for i, record := range msg.AAPeopleUserName {
		if record != msg.UserName {
			res = append(res, model.AAStatementRecord{
				PayerUserName:    msg.AAPeopleUserName[i],
				ReceiverUserName: msg.UserName,
				CashToPay:        cashAA,
				Remark:           sql.NullString{String: msg.Remark},
				Completed:        "N",
			})
		}
	}
	return res
}

func GetFullStatementRecordFromPayMsg(msg protocol.CashPayInputMessage) model.FullStatementRecord {
	var qid, _ = dao.GetSequenceNextVal()
	return model.FullStatementRecord{Qid: qid,
		UserName:      msg.UserName,
		CashOut:       msg.CashPay,
		AAPeopleCount: "",
		Type:          strconv.Itoa(constants.TypeAAPay),
		Remark:        sql.NullString{String:fmt.Sprintf("解决statementQid=%s的AA应付款记录", msg.StatementQid)}}
}