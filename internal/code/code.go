package code

import (
	"strconv"
)

var (
	_codes = map[int]string{}
)

func New(code int, msg string) Code {
	if _, ok := _codes[code]; ok {
		panic("Code already exist: " + strconv.Itoa(code))
	}
	_codes[code] = msg
	return Code(code)
}

type Wrapper interface {
	Code() int
	Message() string
	Error() string
}

type Code int

func (code Code) Code() int {
	return int(code)
}

func (code Code) Message() string {
	return _codes[code.Code()]
}

func (code Code) Error() string {
	return strconv.FormatInt(int64(code), 10)
}

func String(c string) Code {
	code, err := strconv.Atoi(c)
	if err != nil {
		code = SystemErr.Code()
	}
	return Code(code)
}
