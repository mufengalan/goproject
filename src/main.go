package main

import (
	"github.com/gin-gonic/gin"
	"project/App/Login"
	"project/MiddleWare"
)

func main() {
	r := gin.Default()

	r.POST("/loginForm", Login.AppLogin)        //表单参数解析
	r.GET("/redirect", Login.Redirect)          //重定向
	r.GET("/long_async", Login.Async)           //异步处理
	r.GET("/long_sync", Login.Sync)             //同步处理
	r.GET("/home", MiddleWare.AuthMiddleWare()) //中间件验证cookie
	r.GET("/login", Login.ALogin)               //设置cookie
	r.Run(":8000")
}
