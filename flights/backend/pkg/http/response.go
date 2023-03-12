package http

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	Context *gin.Context
}

type Response struct {
	Code          int         `json:"code"`
	Message       string      `json:"msg"`
	TimestampUnix int64       `json:"timestamp"` // Unix timestamp, to convert it to Date in js you can do this: new Date(timestamp * 1000)
	Data          interface{} `json:"data"`
}

func sendResponse(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(code, Response{
		Code:          code,
		Message:       msg,
		TimestampUnix: time.Now().Unix(),
		Data:          data,
	})
}

func (gin *Gin) GenericResponse(code int, message string, data interface{}) {
	gin.Context.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func (gin *Gin) OK(data interface{}) {
	sendResponse(gin.Context, 200, "OK", data)
}

func (gin *Gin) Created(data interface{}) {
	sendResponse(gin.Context, 201, "Created", data)
}

func (gin *Gin) Accepted(data interface{}) {
	sendResponse(gin.Context, 202, "Accepted", data)
}

func (gin *Gin) NoContent(data interface{}) {
	sendResponse(gin.Context, 204, "No Content", data)
}

func (gin *Gin) BadRequest(data interface{}) {
	sendResponse(gin.Context, 400, "Bad Request", data)
}

func (gin *Gin) Unauthorized(data interface{}) {
	sendResponse(gin.Context, 401, "Unauthorized", data)
}

func (gin *Gin) NotFound(data interface{}) {
	sendResponse(gin.Context, 404, "Not Found", data)
}
