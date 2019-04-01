package error

import (
	"bytes"
	"fmt"
)

// === ERRORS
func GeneralError(err error) ChatAPIError {
	return &apiError{
		errorCode:    0,
		err:          err,
		errorMessage: "error",
	}
}

func FailToCreateChatRooom(description string) ChatAPIError {
	return &apiError{
		errorCode:    0,
		err:          nil,
		errorMessage: "faild to create" + description,
	}
}

func NestedError(err []ChatAPIError) ChatAPIError {
	buf := bytes.NewBufferString("")
	for _, e := range err {
		fmt.Fprintf(buf, "%s \n", e.Error())
	}

	return &apiError{
		errorCode:    0,
		err:          nil,
		errorMessage: buf.String(),
	}
}

func ErroRDBConnection(err error) ChatAPIError {
	return &apiError{
		errorCode:    1,
		err:          err,
		errorMessage: "rdb connection failed",
	}
}

func ErroRDBMigration(err error) ChatAPIError {
	return &apiError{
		errorCode:    2,
		err:          err,
		errorMessage: "rdb migration failed",
	}
}
func ErrorNoPermission() ChatAPIError {
	return &apiError{
		errorCode:    3,
		err:          nil,
		errorMessage: "invalid login request",
	}
}
func ErrorLoginAuthFail(err error) ChatAPIError {
	return &apiError{
		errorCode:    3,
		err:          err,
		errorMessage: "invalid login request",
	}
}

// === ControllFlow
type ChatAPIError interface {
	Error() string
	Localize() string
}

type apiError struct {
	errorCode             int
	err                   error
	errorMessage          string
	localizedErrorMessage string
}

func (e *apiError) Error() string {
	sysMsg := ""
	if e.err != nil {
		sysMsg = e.err.Error()
	}

	return fmt.Sprintf("ERROR %d: %s, (%s)", e.errorCode, e.errorMessage, sysMsg)
}

func (e *apiError) Localize() string {
	if e.localizedErrorMessage != "" {
		return e.localizedErrorMessage
	}

	return fmt.Sprintf("サーバーエラーが起きています: %d", e.errorCode)
}
