package common

import "github.com/pkg/errors"

type ServiceResult interface {
	SetCode(code int)
	GetCode() int
	SetMessage(msg string)
	GetMessage() string
	GetData() any
}

type BaseServiceResult struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}

func (r *BaseServiceResult) GetData() any {
	return r.Data
}

func (r *BaseServiceResult) SetCode(code int) {
	r.Code = code
}

func (r *BaseServiceResult) GetCode() int {
	return r.Code
}

func (r *BaseServiceResult) SetMessage(msg string) {
	r.Message = msg
}

func (r *BaseServiceResult) GetMessage() string {
	return r.Message
}
func (r *BaseServiceResult) SetError(err *ServiceError, internalErr ...error) {
	r.Code = err.Code

	if len(internalErr) == 0 {
		r.Message = err.Error()
	} else {
		r.Message = errors.Wrapf(err, "reson: %s", internalErr).Error()
	}
}
