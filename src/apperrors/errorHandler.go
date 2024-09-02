package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError

	// errを第二引数のstructに変換可能であれば実行
	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: UnKnownError,
			Message: "internal process failed",
			Err:     err,
		}
	}

	var statusCode int

	switch appErr.ErrCode {
	case NotAvailableData, BadParam:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
