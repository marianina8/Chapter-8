package main

import (
	"fmt"
	"os"
)

func IsAtty() {
	fileInfo, _ := os.Stdout.Stat()
	if (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("Is a TTY")
	} else {
		fmt.Println("Is not a TTY")
	}
}
