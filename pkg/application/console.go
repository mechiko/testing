package application

import (
	"debug/pe"
	"fmt"
	"os"
)

// these constants are copied from https://github.com/golang/go/blob/6219b48e11f36329de801f62f18448bb4b1cd1a5/src/cmd/link/internal/ld/pe.go#L92-L93
// https://stackoverflow.com/questions/58813512/is-it-possible-to-detect-if-go-binary-was-compiled-with-h-windowsgui-at-runtime

const (
	IMAGE_SUBSYSTEM_WINDOWS_GUI = 2
	IMAGE_SUBSYSTEM_WINDOWS_CUI = 3
)

func isConsole() bool {
	defer RecoverFmtFunc("pkg:application.isConsole()")

	fileName, err := os.Executable()
	if err != nil {
		os.Exit(1)
	}
	fl, err := pe.Open(fileName)
	if err != nil {
		fmt.Printf("Ошибка isConsole() pe.Open(fileName) fileName:%s err:%s\n", fileName, err.Error())
		os.Exit(1)
	}
	defer fl.Close()

	var subsystem uint16
	if header, ok := fl.OptionalHeader.(*pe.OptionalHeader64); ok {
		subsystem = header.Subsystem
	} else if header, ok := fl.OptionalHeader.(*pe.OptionalHeader32); ok {
		subsystem = header.Subsystem
	}

	if subsystem == IMAGE_SUBSYSTEM_WINDOWS_GUI {
		// fmt.Println("it is windows GUI")
		return false
	} else if subsystem == IMAGE_SUBSYSTEM_WINDOWS_CUI {
		// fmt.Println("it is windows CUI")
		return true
	}
	// else {
	//   fmt.Println("binary type unknown")
	// }
	return false
}
