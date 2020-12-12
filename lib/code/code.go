package code

import (
	"strconv"
)

var codes = map[int]string{}

func code(code int, msg string) Code {
	if _, ok := codes[code]; ok {
		panic("Code already exist: " + strconv.Itoa(code))
	}
	codes[code] = msg
	return Code(code)
}

type Codes interface {
	Code() int
	Error() string
	Message() string
}

type Code int

func (e Code) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

func (e Code) Code() int {
	return int(e)
}

func (e Code) Message() string {
	return codes[e.Code()]
}

func String(e string) Code {
	i, err := strconv.Atoi(e)
	if err != nil {
		return SystemErr
	}
	return Code(i)
}
