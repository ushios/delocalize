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
	delete    bool
)

func main() {
	flag.StringVar(&directory, "target", "./", "target directory")
	flag.StringVar(&directory, "t", "./", "target directory")
	flag.BoolVar(&delete, "delete", false, "delete .localized filed")
	flag.BoolVar(&delete, "d", false, "delete .localized filed")
	flag.Parse()

	d, err := filepath.Abs(directory)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("target:", d)

	dd := delocalize.NewDirectoryDispatcher(
		directoryDispatcherMode(delete),
		DirectoryWorkerNum*5,
		DirectoryWorkerNum,
	)

	dd.Add(d)

	dd.Start()
	dd.Wait()
}

func directoryDispatcherMode(delete bool) delocalize.ExecuteMode {
	if delete {
		return delocalize.ExecuteModeDelete
	}

	return delocalize.ExecuteModeDebugPrint
}
