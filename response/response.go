package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"body,omitempty"`
}

func respond(ctx *gin.Context, code int, response Response) {
	ctx.JSONP(code, response)
}

func FormatResponse(ctx *gin.Context, code int, message string, body interface{}) {
	response := Response{}
	if body != nil {
		response.Body = body
	}
	response.Code = code
	response.Message = message

	respond(ctx, code, response)
}
