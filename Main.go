// GCAS = Graphite's Cash AA Service
package main

import (
	"CashAAService/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 获取应付款项
	r.POST("/api/gcas/allcreditpayable", service.GetAllCreditOut)
	// 获取应收款项
	r.POST("/api/gcas/allcreditreceivable", service.GetAllCreditOut)
	//获取流水记录（包括出账记录和aa转账记录）
	r.POST("/api/gcas/statementquery", service.GetAllCreditOut)
	// 获取收到的信息
	r.POST("/api/gcas/mailquery", service.GetAllCreditOut)
	// 获取成员信息
	r.POST("/api/gcas/userquery", service.GetAllCreditOut)
	// 添加出账记录
	r.POST("/api/gcas/recordinput", service.GetAllCreditOut)
	// 添加成员
	r.POST("/api/gcas/userinput", service.GetAllCreditOut)
	// 添加付款码
	r.POST("/api/gcas/receiptcodeinput", service.GetAllCreditOut)
	// 情况款项记录
	r.POST("/api/gcas/recordtruncate", service.GetAllCreditOut)
	// 登入
	r.POST("/api/gcas/login", service.GetAllCreditOut)
	// 登出
	r.POST("/api/gcas/logout", service.GetAllCreditOut)

	r.Run(":8080")
}