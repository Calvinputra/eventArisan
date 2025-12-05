package helper

import (
	"fmt"
	"log"
	"runtime"
)

func ThrowWithoutMessage(err error) {
	if err != nil {
		fmt.Println("Err:", err)
		panic(err)
	}
}

func ThrowWithMessage(err error, msg string) {
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		log.Printf("[error] %s on %s:%d %v", msg, filename, line, err)
		panic(err)
	}
}
