package linuxutils

import (
	"fmt"
	"io/fs"
	"os"
)

func Touch(s string){
	if len(s) < 1 {
		fmt.Println("no file name given, could not create anything")
		os.Exit(1)
	}
	os.WriteFile(s, nil, fs.ModeDevice)
	os.Exit(0)
}

