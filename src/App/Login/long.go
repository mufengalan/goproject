package Login

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//异步
func Async(c *gin.Context) {
	copyContext := c.Copy()
	//异步处理
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("异步执行：" + copyContext.Request.URL.Path)
	}()
}

//同步
func Sync(c *gin.Context) {
	time.Sleep(3 * time.Second)
	log.Println("同步执行：" + c.Request.URL.Path)
}
