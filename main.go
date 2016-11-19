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

	deld := delocalize.NewDeleteDispatcher(
		deleteMode(delete), 10, 5,
	)

	dird := delocalize.NewDirectoryDispatcher(
		DirectoryWorkerNum*5,
		DirectoryWorkerNum,
		deld,
	)

	dird.Add(d)

	dird.Start()
	deld.Start()

	dird.Wait()
	deld.Wait()
}

func deleteMode(delete bool) delocalize.DeleteMode {
	if delete {
		return delocalize.DeleteModeDelete
	}

	return delocalize.DeleteModeDebugPrint
}
