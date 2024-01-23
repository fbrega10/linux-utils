package linuxutils

import (
	"fmt"
	"os"
	"io"
)

func Touch(s string){
	if len(s) < 1 {
		fmt.Println("no file name given, could not create anything")
		os.Exit(1)
	}
	os.Create(s)
	os.Exit(0)
}

func Rm(file_name string){
	err := os.Remove(file_name)
	if err != nil{
		fmt.Println("could not remove the file named ", file_name)
		os.Exit(1)
	}
}

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
