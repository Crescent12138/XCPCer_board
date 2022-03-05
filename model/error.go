package model

import (
	"github.com/FengZhg/errs"
)

const (
	errorHTTPResponseStatus = 259001
)

var (
	HTTPResponseStatusError = errs.NewError(errorHTTPResponseStatus, "HTTP请求相应状态错误")
)
