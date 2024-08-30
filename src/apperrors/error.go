package apperrors

type MyAppError struct {
	// フィールド名を省略すると型名がそのままフィールド名になる
	ErrCode
	Message string
	Err     error
}

// Errorメソッドを実装することでerrorインターフェースを実装する
func (e *MyAppError) Error() string {
	return e.Err.Error()
}

func (e *MyAppError) Unwrap() error {
	return e.Err
}

func (code ErrCode) Wrap(err error, message string) *MyAppError {
	return &MyAppError{ErrCode: code, Err: err}
}
