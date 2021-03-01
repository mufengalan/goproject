package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
    r := gin.Default()
    r.GET("/get", func(context *gin.Context) {
		filename := context.Query("filename")
		context.String(http.StatusOK,"%s ",filename)
	})
    r.Run()
}
