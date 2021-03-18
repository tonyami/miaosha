package code

import (
	"strconv"
)

var errcode = map[int]string{}

func code(code int, msg string) Code {
	if _, ok := errcode[code]; ok {
		panic("Code already exist: " + strconv.Itoa(code))
	}
	errcode[code] = msg
	return Code(code)
}

type CodeMsg interface {
	Code() int
	Error() string
}

type Code int

func (e Code) Error() string {
	return errcode[e.Code()]
}

func (e Code) Code() int {
	return int(e)
}
