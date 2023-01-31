package midware

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/httpd"
)

func OutputHandle() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		// 输出错误信息

		if err, exists := c.Get("Error"); exists && err != nil {
			c.AbortWithStatusJSON(errorCode(c), httpd.NewMessage(err))
			return
		}

		// 输出请求结果

		if res, exists := c.Get("Payload"); exists && res != nil {
			c.AbortWithStatusJSON(200, httpd.NewPayload(res))
			return
		}

		// 输出HTML内容

		if res, exists := c.Get("HTML"); exists && res != nil {
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, res.(string))
			c.Abort()
			return
		}

		// 捕获异常返回

		c.AbortWithStatusJSON(500, httpd.NewMessage("内部错误"))

	}

}

func errorCode(c *gin.Context) int {

	if code := c.GetInt("ErrorCode"); code > 400 {
		return code
	}

	return 400

}
