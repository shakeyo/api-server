package models

// Predefined model error codes.
const (
	ErrDatabase = -1
	ErrSystem   = -2
	ErrDupRows  = -3
	ErrNotFound = -4
)

// CodeInfo definiton.
type ResponseInfo struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Data string `json:"data"`
}

// NewErrorInfo return a CodeInfo represents error.
func NewErrorInfo(info string) *ResponseInfo {
	return &ResponseInfo{-1, info, ""}
}

// NewNormalInfo return a CodeInfo represents OK.
func NewNormalInfo(info string, data string) *ResponseInfo {
	return &ResponseInfo{0, info, data}
}
