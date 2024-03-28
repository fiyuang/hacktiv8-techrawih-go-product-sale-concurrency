package response

import "github.com/gin-gonic/gin"

func Error(c *gin.Context, httpCode int, err error) {
	c.JSON(httpCode, gin.H{
		"success":    false,
		"statusCode": httpCode,
		"message":    err.Error(),
	})
}

func Success(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"success":    true,
		"statusCode": httpCode,
		"message":    "success",
		"data":       data,
	})
}

func SuccessCustomMessage(c *gin.Context, httpCode int, message string) {
	c.JSON(httpCode, gin.H{
		"success":    true,
		"statusCode": httpCode,
		"message":    message,
	})
}
