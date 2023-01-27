package main

import (
	"fmt"
	"testing/internal/guiapp"
	"time"
)

func main() {
	// Run
	for {
		exit := guiapp.Run()
		fmt.Printf("main() for Run exit = %v\n", exit)
		time.Sleep(1 * time.Second)
		if exit == 0 {
			break
		}
	}
	fmt.Printf("END OF GAME!\n")
}
