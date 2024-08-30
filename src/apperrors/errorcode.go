package apperrors

type ErrCode string

const (
	UnKnownError     ErrCode = "U000"
	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NotAvailableData ErrCode = "S003"
	NoTargetData     ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"
)
