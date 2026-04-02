package functest

import (
	"errors"
)

func getDeferfunc() func(err *error) {
	clear := func(err *error) {
		if err == nil {
			println("err is nil")
		} else {
			println((*err).Error())
		}
	}
	return clear
}

func GetDeferTest() (err error) {

	deferfunc := getDeferfunc()
	defer deferfunc(&err)

	return errors.New("test error")
}
