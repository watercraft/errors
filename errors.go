// Copyright 2017 F. Alan Jones.  All rights reserved.

package errors

type Error struct {
	msg  string
	code int
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Is(code int) bool {
	return e.code == code
}

func New(code int, msg string) *Error {
	return &Error{
		msg:  msg,
		code: code,
	}
}
