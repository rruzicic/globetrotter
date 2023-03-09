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

func (gin *Gin) OK() {
	sendResponse(gin.Context, 200, "OK", nil)
}

func (gin *Gin) OKObject(data interface{}) {
	sendResponse(gin.Context, 200, "OK", data)
}

func (gin *Gin) Created() {
	sendResponse(gin.Context, 201, "Created", nil)
}

func (gin *Gin) CreatedObject(data interface{}) {
	sendResponse(gin.Context, 201, "Created", data)
}

func (gin *Gin) Accepted() {
	sendResponse(gin.Context, 202, "Accepted", nil)
}

func (gin *Gin) AcceptedObject(data interface{}) {
	sendResponse(gin.Context, 202, "Accepted", data)
}

func (gin *Gin) NoContent() {
	sendResponse(gin.Context, 204, "No Content", nil)
}

func (gin *Gin) BadRequest() {
	sendResponse(gin.Context, 400, "Bad Request", nil)
}

func (gin *Gin) Unauthorized() {
	sendResponse(gin.Context, 401, "Unauthorized", nil)
}

func (gin *Gin) NotFound() {
	sendResponse(gin.Context, 404, "Not Found", nil)
}
