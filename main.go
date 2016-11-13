package main

import (
	"log"
	"path/filepath"

	"github.com/ushios/delocalize/lib/delocalize"
)

var (
	// DirectoryWorkerNum directory worker num
	DirectoryWorkerNum = 4
)

func main() {
	wd, err := filepath.Abs("/Users/shugo")
	if err != nil {
		log.Fatal(err)
	}

	dd := delocalize.NewDirectoryDispatcher(
		DirectoryWorkerNum*5,
		DirectoryWorkerNum,
	)

	dd.Add(wd)

	dd.Start()
	dd.Wait()
}
