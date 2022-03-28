package model

type Message interface {
	GetHeader() Header
}
type Header struct {

}

type LoginMessage struct {
	UserName string
	PassWord string
}

type CreditPayableReqMsg struct {
	UserName string
}

type CreditPayableMessage struct {
	Header
	UserName string
	UserNickName string
	CreditPayableItems []CreditPayableItem
}

type CreditPayableItem struct {
	UserNickName string // 款项应付人
	CashToPay string // 应付款金额
}

func (msg *CreditPayableMessage) GetHeader() Header {
	return msg.Header
}

func (msg *LoginMessage) GetHeader() Header {
	return Header{}
}