package dao

import (
	"CashAAService/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetFullStatementRecord() (error, []model.FullStatementRecord) {
	var records []model.FullStatementRecord
	err := db.Select(&records, "select * from graphite.full_statement_record")
	if err != nil {
		fmt.Println("exec failed, ")
		return err, nil
	}
	fmt.Println("select succ:", records)
	return nil, records
}

func InsertFullStatementRecord(record model.FullStatementRecord) error {
	stmt, err := db.Prepare("insert into graphite.full_statement_record" +
		" (QID,USER_NAME,CASH_OUT,AA_PEOPLE_COUNT,TYPE,REMARK)" +
		" values (?,?,?,?,?,?)")
	if err != nil {
		log.Println("open mysql failed,")
		return err
	}
	record.Qid, _ = GetSequenceNextVal()
	res, err := stmt.Exec(record.Qid, record.UserName, record.CashOut, record.AAPeopleCount, record.Type, record.Remark)
	id, err := res.LastInsertId()
	if err != nil {
		log.Println("open mysql failed,")
		return err
	}
	fmt.Println(id)
	return nil
}
