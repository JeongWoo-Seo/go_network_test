package errors

import (
	"fmt"
)

const (
	NotFoundUser int64 = iota
	DBError
)

var errMassage = map[int64]string{
	NotFoundUser: "not found user",
	DBError:      "db error : ",
}

func Errorf(code int64, args ...interface{}) error {
	if message, ok := errMassage[code]; ok {
		return fmt.Errorf("%s : %s", message, args)
	} else {
		return fmt.Errorf("not found err code")
	}
}
