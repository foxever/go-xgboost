package core

/*
#cgo pkg-config: xgboost

#include <xgboost/c_api.h>
*/
import "C"
import (
	"errors"
)

func checkError(res C.int) error {
	if int(res) != 0 {
		errStr := C.XGBGetLastError()
		return errors.New(C.GoString(errStr))
	}

	return nil
}