package zerrors

import "fmt"

type BaseError struct {
	Code    int32
	Message string
}

func (b *BaseError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", b.Code, b.Message)
}

func NewError(code int32, message string) *BaseError {
	return &BaseError{Code: code, Message: message}
}

var ErrGeneralAccessDenied = NewError(10000, "no permission to this operation")
var ErrGeneralInvalidPath = NewError(10001, "invalid uri path")
var ErrGeneralInvalidType = NewError(10002, "invalid type")
var ErrHelpContentEmpty = NewError(10003, "Content is empty")
