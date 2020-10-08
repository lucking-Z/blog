package uuid

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GenerateUUID 生成UUID
func GenerateUUID() (val string) {
	u := uuid.NewV4()
	val = fmt.Sprintf("%s", u)
	return
}

// GetXRequestID 获取请求ID
func GetXRequestID(c *gin.Context) string {
	xRequestID, exists := c.Get("XRequestID")
	if !exists {
		return GenerateUUID()
	}
	return xRequestID.(string)
}