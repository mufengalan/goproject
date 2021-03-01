package Login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
}
