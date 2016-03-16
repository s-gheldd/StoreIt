package main

import (
	"os"
	"fmt"
	"github.com/s-gheldd/StoreIt/DataStructure"
)

func main() {

	dir, _ := os.Getwd()
	file, _ := os.Open(fmt.Sprintf("%s/%s", dir, "StoreIt.go"))
	savedFile, err := DataStructure.NewStoredFile(file)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(0)
	}


	fmt.Println(savedFile.VerboseString())

}
