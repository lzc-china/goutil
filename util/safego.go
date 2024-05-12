package util

import (
	"fmt"
	"runtime"
)

func Go(f func()) {

	go func() {

		defer func() {
			if err := recover(); err != nil {
				// 获取panic发生时的上下文信息
				stackTrace := make([]byte, 4096)
				length := runtime.Stack(stackTrace, false)

				fmt.Printf("runtime panic: %+v ; stackTrace: %+v", err, string(stackTrace[:length]))
			}
		}()

		f()
	}()

}
