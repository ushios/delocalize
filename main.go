package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/ushios/delocalize/lib/delocalize"
)

var (
	// DirectoryWorkerNum directory worker num
	DirectoryWorkerNum = 4

	directory string
	debug     bool
)

func main() {
	flag.StringVar(&directory, "target", "./", "target directory")
	flag.StringVar(&directory, "t", "./", "target directory")
	flag.BoolVar(&debug, "debug", false, "not delete only print .localized files")
	flag.BoolVar(&debug, "d", false, "not delete only print .localized files")
	flag.Parse()

	d, err := filepath.Abs(directory)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("target:", d)

	dd := delocalize.NewDirectoryDispatcher(
		DirectoryWorkerNum*5,
		DirectoryWorkerNum,
	)

	if debug {
		dd.ExecuteMode = delocalize.ExecuteModeDebugPrint
	} else {
		dd.ExecuteMode = delocalize.ExecuteModeDelete
	}

	dd.Add(d)

	dd.Start()
	dd.Wait()
}
