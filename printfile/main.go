package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for idx, input := range os.Args {
		if idx > 0 {
			f, err := os.Open(input)

			if err != nil {
				fmt.Println("Error:", err)
			}

			io.Copy(os.Stdout, f)
		}
	}
}
