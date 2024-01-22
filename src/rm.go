package linuxutils

import (
	"fmt"
	"os"
)

func Rm(file_name string){
	err := os.Remove(file_name)
	if err != nil{
		fmt.Println("could not remove the file named ", file_name)
		os.Exit(1)
	}
}
