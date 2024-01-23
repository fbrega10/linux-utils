package main

import (
	"github.com/fbrega10/linux-utils/src"
	"os"
	"fmt"
)


func main(){
	const file_path = "person.json"
	const tmpfl = "newfile.json"
	linuxutils.Cat(file_path)
	linuxutils.Touch(tmpfl)
	linuxutils.Rm(tmpfl)
	err := os.Remove(tmpfl)
	if err != nil{
		fmt.Println("error occurred deleting file", err)
	}

}
