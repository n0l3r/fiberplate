package responder

import (
	"github.com/gofiber/fiber/v2"
)

// ResponseOutput holds success response info
type ResponseOutput struct {
	c            *fiber.Ctx
	Code         int    `json:"code" example:"200"`
	Data         any    `json:"data"`
	Message      string `json:"message" example:"OK"`
	Meta         any    `json:"meta,omitempty"`
	Status       string `json:"status,omitempty"`
	customHeader map[string]string
}

// NewHTTPSuccess creates a new ResponseOutput
func NewHTTPSuccess(c *fiber.Ctx, code int, data any, message string) *ResponseOutput {
	if message == "" {
		message = "OK"
	}
	if code == 0 {
		code = fiber.StatusOK
	}
	return &ResponseOutput{
		c:            c,
		Code:         code,
		Data:         data,
		Message:      message,
		customHeader: make(map[string]string),
	}
}

// WithMeta attaches meta info
func (ro *ResponseOutput) WithMeta(meta any) *ResponseOutput {
	ro.Meta = meta
	return ro
}

// WithStatus attaches status string
func (ro *ResponseOutput) WithStatus(status string) *ResponseOutput {
	ro.Status = status
	return ro
}

// WithHeader attaches custom header
func (ro *ResponseOutput) WithHeader(key, val string) *ResponseOutput {
	ro.customHeader[key] = val
	return ro
}

// JSON formats and sends JSON response
func (ro *ResponseOutput) JSON() error {
	for k, v := range ro.customHeader {
		ro.c.Set(k, v)
	}

	return ro.c.Status(ro.Code).JSON(ro)
}
