package cache

var tokenMap map[string]string

func TokenInit() {
	tokenMap = make(map[string]string)
}

func PutToken(userName string, token string) {
	tokenMap[userName] = token
}

func CheckToken(userName string, token string) bool {
	if tk, ok := tokenMap[userName]; ok && token == tk {
		return true
	}
	return false
}