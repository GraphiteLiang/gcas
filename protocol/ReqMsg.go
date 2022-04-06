package protocol

type ErrorMessage struct {
	ErrorCode    string
	ErrorMessage string
}

type LoginMessage struct {
	UserName string
	PassWord string
}

type CashOutInputMessage struct {
	UserName         string
	CashOut          string
	AAPeopleCount    string
	Remark           string
	AAPeopleUserName []string
}

type CashPayInputMessage struct {
	UserName     string
	CashPay      string
	StatementQid string
}
