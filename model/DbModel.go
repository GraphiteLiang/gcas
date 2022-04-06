package model

import "database/sql"

type FullStatementRecord struct {
	Qid           string         `db:"QID"`
	UserName      string         `db:"USER_NAME"`
	CashOut       string         `db:"CASH_OUT"`
	AAPeopleCount string         `db:"AA_PEOPLE_COUNT"`
	Type          string         `db:"TYPE"`
	Remark        sql.NullString `db:"REMARK"`
}

type AAStatementRecord struct {
	Qid              string         `db:"QID"`
	PayerUserName    string         `db:"PAYER_USER_NAME"`
	ReceiverUserName string         `db:"RECEIVER_USER_NAME"`
	CashToPay        string         `db:"CASH_TO_PAY"`
	Remark           sql.NullString `db:"REMARK"`
	Completed        string         `db:"COMPLETED"`
}

type GcasUser struct {
	Qid             string         `db:"QID"`
	UserName        string         `db:"USER_NAME"`
	UserNickname    sql.NullString `db:"USER_NICK_NAME"`
	Password        string         `db:"PASSWORD"`
	UpdateTime      string         `db:"UPDATE_TIME"`
	UserReceiptCode sql.NullString `db:"USER_RECEIPT_CODE"`
}

type GcasUserMail struct {
	Qid        string `db:"QID”`
	UserId     string `db:"USER_ID”`
	MailInfo   string `db:"MAIL_INFO”`
	Read       string `db:"READ”`
	UpdateTime string `db:"UPDATE_TIME”`
	RecvTime   string `db:"RECV_TIME”`
}
