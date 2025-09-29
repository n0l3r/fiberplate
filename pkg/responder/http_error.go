package responder

import (
	"github.com/gofiber/fiber/v3"
)

// ErrorOutput holds error response info
type ErrorOutput struct {
	c                fiber.Ctx
	Code             int                `json:"code"`
	Data             []*ErrorOutputData `json:"data,omitempty"`
	Message          string             `json:"message"`
	Status           string             `json:"status,omitempty"`
	ErrorTracingCode string             `json:"error_code,omitempty"`
	AppName          string             `json:"app_name,omitempty"`
	DataInterface    any                `json:"-"`
}

type ErrorOutputData struct {
	Quota       string `json:"quota,omitempty"`
	Field       string `json:"field,omitempty"`
	Description string `json:"description,omitempty"`
	Data        any    `json:"data,omitempty"`
}

// NewHTTPError creates and initializes ErrorOutput
func NewHTTPError(c fiber.Ctx, code int, message string) *ErrorOutput {
	if code == 0 {
		code = fiber.StatusInternalServerError
	}
	if message == "" {
		message = "Internal Server Error"
	}
	return &ErrorOutput{
		c:       c,
		Code:    code,
		Message: message,
	}
}

// WithStatus attaches status string
func (eo *ErrorOutput) WithStatus(status string) *ErrorOutput {
	eo.Status = status
	return eo
}

// WithTracingCode attaches error tracing code
func (eo *ErrorOutput) WithTracingCode(code string) *ErrorOutput {
	eo.ErrorTracingCode = code
	return eo
}

// WithAppName attaches app name
func (eo *ErrorOutput) WithAppName(name string) *ErrorOutput {
	eo.AppName = name
	return eo
}

// WithData attaches error detail data
func (eo *ErrorOutput) WithData(data []*ErrorOutputData) *ErrorOutput {
	eo.Data = data
	return eo
}

// WithDataInterface attaches custom error data (for legacy/compat)
func (eo *ErrorOutput) WithDataInterface(data any) *ErrorOutput {
	eo.DataInterface = data
	return eo
}

// JSON formats and sends error as JSON
func (eo *ErrorOutput) JSON() error {
	if eo.DataInterface != nil {
		return eo.c.Status(eo.Code).JSON(fiber.Map{
			"code":       eo.Code,
			"data":       eo.DataInterface,
			"message":    eo.Message,
			"error_code": eo.ErrorTracingCode,
			"app_name":   eo.AppName,
		})
	}
	return eo.c.Status(eo.Code).JSON(eo)
}

// XML formats and sends error as XML
func (eo *ErrorOutput) XML() error {
	return eo.c.Status(eo.Code).XML(eo)
}
