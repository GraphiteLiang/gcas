// GCAS = Graphite's Cash AA Service
// TODO 分页查询
package main

import (
	"CashAAService/cache"
	"CashAAService/controller"
	"CashAAService/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func Init() {
	cache.TokenInit()
	dao.Init()
	cache.UserCacheInit()
}
func main() {
	Init()
	r := gin.Default()
	r.Use(Cors())
	// 获取应收、付款项
	r.GET("/api/gcas/test", controller.Test)
	// 获取应收、付款项
	r.GET("/api/gcas/allcreditrelated", controller.GetCreditStatementRelated)
	// 获取流水记录（包括出账记录和aa转账记录）
	r.GET("/api/gcas/statementquery", controller.GetFullStatement)
	// 获取收到的信息 TODO
	r.GET("/api/gcas/mailquery", controller.GetCreditStatementRelated)
	// 获取成员信息
	r.GET("/api/gcas/userquery", controller.GetAllUser)
	// 添加出账记录
	r.POST("/api/gcas/recordinput", controller.InputCashOutRecord)
	// 添加成员 TODO
	r.POST("/api/gcas/userinput", controller.GetCreditStatementRelated)
	// 添加付款码 TODO
	r.POST("/api/gcas/receiptcodeinput", controller.GetCreditStatementRelated)
	// 清空款项记录 TODO
	r.POST("/api/gcas/recordtruncate", controller.GetCreditStatementRelated)
	// 登入
	r.POST("/api/gcas/login", controller.Login)
	// 登出 TODO
	r.POST("/api/gcas/logout", controller.GetCreditStatementRelated)
	// 进行付款
	r.POST("/api/gcas/payinput", controller.InputCashPay)
	// 提醒付款 TODO
	r.POST("/api/gcas/hintpay", controller.GetCreditStatementRelated)

	r.Run(":8080")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method      //请求方法
		origin := c.Request.Header.Get("Origin")        //请求头部
		var headerKeys []string                             // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")       // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()        //  处理请求
	}
}