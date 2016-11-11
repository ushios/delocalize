package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	wd, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(wd)
}
