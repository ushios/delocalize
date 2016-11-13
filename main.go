package main

import (
	"log"
	"path/filepath"

	"github.com/ushios/delocalize/lib/delocalize"
)

func main() {
	wd, err := filepath.Abs("/Users/shugo")
	if err != nil {
		log.Fatal(err)
	}

	dd := delocalize.NewDirectoryDispatcher(1000, 200)

	dd.Add(wd)

	dd.Start()
	dd.Wait()
}
