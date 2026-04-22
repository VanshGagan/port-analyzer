package utils

import (
	"fmt"
	"time"
)

var KeepLoading = true

func Loading() {
	dots := []string{".  ", ".. ", "...", ".. "}
	for i := 0; ; i++ {
		if !KeepLoading {
			fmt.Print("\r")
			return
		}
		fmt.Printf("\r                     \r")
		fmt.Printf("\rScanning%s", dots[i%len(dots)])

		time.Sleep(200 * time.Millisecond)
	}
}
