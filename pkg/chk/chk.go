package chk

import "fmt"

// SE SystemError
func SE(err error, msg string) {
	if err != nil {
		panic(fmt.Errorf("err:%s\nmsg:%s", err.Error(), msg))
	}
}
