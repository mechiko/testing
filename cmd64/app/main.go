package main

import (
	"fmt"
	"testing/internal/cmdapp"
	"time"
)

func main() {
	// Run
	for {
		exit := cmdapp.Run()
		fmt.Printf("main() for Run exit = %v\n", exit)
		time.Sleep(3 * time.Second)
		if exit == 0 {
			break
		}
	}
	fmt.Printf("END OF GAME!\n")
}
