package constants

// 出金类型
const (
	TypeCashOut = iota // 普通出金
	TypeAAPay          // AA出金
)

// 计算符号类型
const (
	Plus = iota
	Minus
	Multiply
	Division
)
