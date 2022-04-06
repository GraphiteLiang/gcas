package cache

import (
	"CashAAService/dao"
	"CashAAService/model"
)

// TODO:保存用户中文名

var userMap map[string]model.GcasUser

func UserCacheInit() {
	userMap = make(map[string]model.GcasUser)
	users, _ := dao.GetAllUser()
	for _, user := range users {
		userMap[user.UserName] = user
	}
}

func GetUserNickName(userName string) string {
	if user, ok := userMap[userName]; ok {
		return user.UserNickname.String
	} else {
		return userName
	}
}
