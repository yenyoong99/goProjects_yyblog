package exception

import (
	"fmt"
	"github.com/yenyoong99/mcube/tools/pretty"
)

func NewAPIException(code int, reason string) *APIException {
	return &APIException{
		Code:   code,
		Reason: reason,
	}
}

// APIException Custom API Error Exception
type APIException struct {
	HttpCode int    `json:"-"`
	Code     int    `json:"code"`
	Reason   string `json:"reason"`
	Message  string `json:"message"`
}

func (e *APIException) Error() string {
	return fmt.Sprintf("%s, %s", e.Reason, e.Message)
}

func (e *APIException) String() string {
	return pretty.ToJSON(e)
}

func (e *APIException) WithMessage(msg string) *APIException {
	e.Message = msg
	return e
}

func (e *APIException) WithHttpCode(code int) *APIException {
	e.HttpCode = code
	return e
}

func (e *APIException) WithMessagef(format string, a ...any) *APIException {
	e.Message = fmt.Sprintf(format, a...)
	return e
}

// IsException judge
func IsException(err error, e *APIException) bool {
	if target, ok := err.(*APIException); ok {
		return target.Code == e.Code
	}

	return false
}
