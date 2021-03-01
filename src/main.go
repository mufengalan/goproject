package main

import (
	"github.com/gin-gonic/gin"
	"project/App/Login"
)

func main() {
	r := gin.Default()

	r.POST("/loginForm", Login.AppLogin) //表单参数解析
	r.GET("/redirect", Login.Redirect)   //重定向
	r.GET("/long_async", Login.Async)
	r.GET("/long_sync", Login.Sync)

	r.Run(":8000")
}
