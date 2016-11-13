package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/ushios/delocalize/lib/delocalize"
)

func main() {
	wd, err := filepath.Abs("/Users/shugo")
	if err != nil {
		log.Fatal(err)
	}

	dd := delocalize.NewDirectoryDispatcher(10, 2)

	dd.Add(wd)

	dd.Start()

	fmt.Println(wd)
}
