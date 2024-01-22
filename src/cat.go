package cat

import (
	"fmt"
	"io"
	"os"
)

func Cat(s string) {
	file, err := os.Open(s)
	defer file.Close()
	if err != nil {
		fmt.Println("file doesn't exist", err)
		return
	}
	for {
		data := make([]byte, 100)
		_, err = file.Read(data)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("error occurred reading file", err)
			os.Exit(1)
		}
		os.Stdout.Write(data)
	}
}
