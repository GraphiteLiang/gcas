package dao

import (
	"CashAAService/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func GetAllUser() ([]model.GcasUser, error) {
	var records []model.GcasUser
	err := db.Select(&records, "select * from graphite.gcas_user")
	if err != nil {
		fmt.Println("exec failed, ")
		return nil, err
	}
	fmt.Println("select succ:", records)
	return records, nil
}

func InsertUser(user model.GcasUser) error {
	stmt, err := db.Prepare("insert into graphite.gcas_user" +
		" (QID,USER_NAME,USER_NICK_NAME,PASSWORD,UPDATE_TIME,USER_RECEIPT_CODE)" +
		" values (?,?,?,?,?,?)")
	if err != nil {
		log.Println("open mysql failed,")
		return err
	}
	user.Qid, _ = GetSequenceNextVal()
	res, err := stmt.Exec(user.Qid, user.UserName, user.UserNickname, user.Password, time.Now().Format("2006-01-02 15:04:05"), user.UserReceiptCode)
	id, err := res.LastInsertId()
	if err != nil {
		log.Println("open mysql failed,")
		return err
	}
	fmt.Println(id)
	return nil
}

func ValidateUser(userName string, password string) bool {
	var records []model.GcasUser
	err := db.Select(&records, "select * from graphite.gcas_user where USER_NAME=? AND PASSWORD=?",
		userName, password)

	if err != nil {
		fmt.Println("exec failed, ", err)
		return false
	}
	fmt.Println("select succ:", records)
	if len(records) >= 1 {
		return true
	}
	return false
}
