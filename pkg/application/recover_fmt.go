package application

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"time"
)

func RecoverFmt(str string) {
	if r := recover(); r != nil {
		stack := make([]byte, 8096)
		stack = stack[:runtime.Stack(stack, false)]
		fmt.Printf("RecoverLog: %v %v %v", str, time.Now(), r)
	}
}

func RecoverFmtWithStack(str string) {
	if r := recover(); r != nil {
		stack := make([]byte, 8096)
		stack = stack[:runtime.Stack(stack, false)]
		fmt.Printf("RecoverLogWithStack: %v %v %v\n", str, time.Now(), r)
		fmt.Printf("%s\n", string(stack))
	}
}

func RecoverFmtFunc(str string) {
	// applicationInstance.Logger.Logger.Debug().Msg("RecoverFunc")
	if r := recover(); r != nil {
		err := fmt.Errorf("%v %v %v", str, time.Now(), r)
		_ = ioutil.WriteFile("error.txt", []byte(err.Error()), 0644)
		fmt.Printf("RecoverFunc:%s\n", err)
	}
}

func RecoverFmtExit(str string) {
	if r := recover(); r != nil {
		err := fmt.Errorf("%v %v %v", str, time.Now(), r)
		_ = ioutil.WriteFile("error.txt", []byte(err.Error()), 0644)
		fmt.Printf("RecoverExit:%s\n", err)
		os.Exit(1)
	}
}
