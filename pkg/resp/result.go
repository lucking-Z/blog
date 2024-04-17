package resp

import (
	"blogs/utils/uuid"
	"github.com/gin-gonic/gin"
)

// Result 通用返回结构
type Result struct {
	ReqID   string        `json:"reqid"`
	Code    int           `json:"code"`
	Message string        `json:"msg"`
	Data    interface{}   `json:"data,omitempty"`
	Details []interface{} `json:"-"`
}

//// AbortWithStatus 中断中间件调用链返回
//func AbortWithStatus(c *gin.Context, httpCode int, err error, obj interface{}) {
//	status := ecode.Cause(err)
//	message := status.MMessage
//	originMsg, err := strconv.Atoi(message)
//	if err == nil && originMsg == status.MCode {
//		message = ""
//	}
//	if len(message) == 0 {
//		message = code.ErrText(status.MCode)
//	}
//	result := &Result{
//		ReqID:   code.GetXRequestID(c),
//		Code:    status.MCode,
//		Message: message,
//		Data:    obj,
//	}
//	c.Set("response", result)
//	c.AbortWithStatusJSON(httpCode, result)
//}

// Render 返回结果
func Render(c *gin.Context, err *Error, obj interface{}) {
	if err == nil {
		err = ErrSuccess
	}
	result := &Result{
		ReqID:   uuid.GetXRequestID(c),
		Code:    err.ErrNo,
		Message: err.ErrMsg,
		Data:    obj,
	}
	c.Header("RequestID")
	c.Set("response", result)
	c.JSON(200, result)
}
