package dao

import (
	"fmt"
	"strconv"
)

func GetSequenceNextVal() (string, error) {
	var currVal []string
	sql := "select curr_val from graphite.sequence where sequence_name='GCAS_SEQUENCE'"
	err := db.Select(&currVal, sql)
	if err != nil {
		fmt.Println("exec failed, ", sql)
		return "", err
	}
	currValInt, _ := strconv.Atoi(currVal[0])
	currValInt = currValInt + 1
	stmt, err := db.Prepare("update graphite.sequence set curr_val=?")
	if err != nil {
		fmt.Println("exec failed", stmt)
		return "", err
	}
	_, err = stmt.Exec(currValInt)
	if err != nil {
		return "", err
	}
	fmt.Println("select succ:", currVal)
	return currVal[0], nil
}