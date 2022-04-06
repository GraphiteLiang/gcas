package dao

import (
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

// Init
//  @Description: 初始化数据库连接
func Init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:11452)/graphite")
	if err != nil {
		log.Println("open mysql failed,", err)
		return
	}
	db = database
}
