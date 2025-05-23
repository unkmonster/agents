// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnknownEventType(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UNKNOWN_EVENT_TYPE.String() && e.Code == 400
}

func ErrorUnknownEventType(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_UNKNOWN_EVENT_TYPE.String(), fmt.Sprintf(format, args...))
}
