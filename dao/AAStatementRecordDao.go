package dao

import (
	"CashAAService/model"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/**

INSERT INTO t
VALUES
  (1, 20, 'a'),
  (2, 26, 'b');
*/

func InsertAAStatementRecord(record model.AAStatementRecord) {
	stmt, err := db.Prepare("insert into aa_statement_record" +
		" (PAYER_USER_NAME,RECEIVER_USER_NAME,CASH_TO_PAY,REMARK,COMPLETED)" +
		" values (?,?,?,?,?)")
	if err != nil {
		log.Println("open mysql failed,", err)
		return
	}
	res, err := stmt.Exec(record)
	id, err := res.LastInsertId()
	if err != nil {
		log.Println("open mysql failed,", err)
		return
	}
	fmt.Println(id)
}

func UpdateStatementOnCompleted(record model.AAStatementRecord) error {
	stmt, err := db.Prepare("update aa_statement_record" +
		" set COMPLETED = ?" +
		" where QID = ?")
	if err != nil {
		log.Println("open mysql failed,", err)
		return err
	}
	res, err := stmt.Exec(record.Completed, record.Qid)
	id, err := res.LastInsertId()
	if err != nil {
		log.Println("open mysql failed,", err)
		return err
	}
	fmt.Println(id)
	return nil
}

//
// UpdateBatchAAStatementRecord
//  @Description: dao层不应该关注业务逻辑， 这个方法应该只实现：更新，如果没有，则插入
//  TODO 要形成事务，出错应不提交
//  @param records
//  @return error
//
func UpdateBatchAAStatementRecord(records []model.AAStatementRecord) error {
	var err error
	var count int
	updateSql := "update graphite.aa_statement_record set " +
		" CASH_TO_PAY = ?, " +
		" REMARK = ?, " +
		" COMPLETED = ?" +
		" where PAYER_USER_NAME = ? AND" +
		" RECEIVER_USER_NAME = ?"
	insertSql := "insert into graphite.aa_statement_record" +
		" (QID,PAYER_USER_NAME,RECEIVER_USER_NAME,CASH_TO_PAY,REMARK,COMPLETED) VALUES" +
		" (?,?,?,?,?,?)"
	var updateStmt, insertStmt *sql.Stmt
	if updateStmt, err = db.Prepare(updateSql); err != nil {
		return err
	}
	if insertStmt, err = db.Prepare(insertSql); err != nil {
		return err
	}
	for _, record := range records {
		res, err := updateStmt.Exec(record.CashToPay, record.Remark, record.Completed, record.PayerUserName, record.ReceiverUserName)
		if err != nil {
			log.Println("更新数据库失败")
			return err
		}
		rowCount, err := res.RowsAffected()
		count = count + int(rowCount)
		if rowCount <= 0 {
			Qid, err := GetSequenceNextVal()
			res, err = insertStmt.Exec(Qid, record.PayerUserName, record.ReceiverUserName, record.CashToPay, record.Remark, record.Completed)
			if err != nil {
				log.Println("插入数据库失败")
				return err
			}
			rowCount, err = res.RowsAffected()
			count = count + int(rowCount)
		}
	}
	return nil
}

func GetAAStatementRecordByUserName(userName string) (error, []model.AAStatementRecord) {
	var records []model.AAStatementRecord
	err := db.Select(&records, "select * from graphite.aa_statement_record "+
		"where (PAYER_USER_NAME=? OR RECEIVER_USER_NAME=?) "+
		"AND COMPLETED='N'", userName, userName)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return err, nil
	}
	fmt.Println("select succ:", records)
	return nil, records
}

func GetAAStatementRecordByQid(qid string) (error, model.AAStatementRecord){
	var records []model.AAStatementRecord
	err := db.Select(&records, "select * from graphite.aa_statement_record "+
		"where QID=? AND COMPLETED = 'N'", qid)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return err, model.AAStatementRecord{}
	}
	fmt.Println("select succ:", records)
	if len(records) > 1 || len(records) < 1 {
		return errors.New("数据库中记录条数错误"), model.AAStatementRecord{}
	}
	return nil, records[0]
}

func GetAAStatementRecordAll(userName string) (error, []model.AAStatementRecord){
	var records []model.AAStatementRecord
	err := db.Select(&records, "select * from graphite.aa_statement_record "+
		"where PAYER_USER_NAME=? OR RECEIVER_USER_NAME=? ", userName, userName)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return err, nil
	}
	fmt.Println("select succ:", records)
	return nil, records
}

func GetAAStatementRecordByPayerReceiver(payerUserName string, receiverUserName string) (error, *model.AAStatementRecord) {
	var records []model.AAStatementRecord
	err := db.Select(&records, "select * from graphite.aa_statement_record "+
		"where (PAYER_USER_NAME=? AND RECEIVER_USER_NAME=?) ",
		payerUserName, receiverUserName)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return err, nil
	}
	fmt.Println("select succ:", records)
	if len(records) > 1 {
		return errors.New("获取应AA记录失败，记录条数错误"), nil
	}
	if len(records) < 1 {
		return nil, nil
	}
	return nil, &records[0]
}

func GetAAStatementRecordByReceiver(receiverName string) (error, []model.AAStatementRecord) {
	var records []model.AAStatementRecord
	err := db.Select(&records, "select * from graphite.aa_statement_record "+
		"where RECEIVER_USER_NAME=?", receiverName)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return err, nil
	}
	fmt.Println("select succ:", records)
	return nil, records
}
