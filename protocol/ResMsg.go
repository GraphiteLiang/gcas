package protocol

type LoginResMessage struct {
	UserName string
	Token    string
}

type CreditPayableMessage struct {
	UserName           string
	UserNickName       string
	CreditPayableItems []CreditPayableItem
}
type StatementMessage struct {
	StatementList []Statement
}
type UserQueryResMessage struct {
	UserInfos []UserInfo
}

type Statement struct {
	UserName      string
	UserNickName  string
	CashOut       string
	AAPeopleCount string
	Type          string
	Remark        string
}

type CreditPayableItem struct {
	DbQid        string // 对应数据库的QID
	ItemType     string // 类型，1-应付款 2-应收款
	UserName     string // 款项应收/付人
	UserNickName string // 款项应收/付人中文名
	Cash         string // 应付款金额
	Remark       string // 备注
}

type UserInfo struct {
	UserName string
	UserNickName string
}
