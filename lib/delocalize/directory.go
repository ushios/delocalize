package delocalize

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type (
	// ExecuteMode is DirectoryDispatcher mode
	ExecuteMode uint8

	// DirectoryDispatcher management workers
	DirectoryDispatcher struct {
		ExecuteMode

		pool    chan *directoryWorker
		queue   chan string
		workers []*directoryWorker
		wg      sync.WaitGroup
		quit    chan struct{}
	}

	directoryWorker struct {
		dispather *DirectoryDispatcher
		data      chan string
		quit      chan struct{}
	}
)

const (
	ExecuteModeDebugPrint = iota
	ExecuteModeDelete
)

// NewDirectoryDispatcher .
func NewDirectoryDispatcher(maxQueues, maxWorkers int) *DirectoryDispatcher {
	d := &DirectoryDispatcher{
		pool:  make(chan *directoryWorker, maxWorkers),
		queue: make(chan string, maxQueues),
		quit:  make(chan struct{}),
	}

	d.workers = make([]*directoryWorker, cap(d.pool))
	for i := 0; i < cap(d.pool); i++ {
		w := directoryWorker{
			dispather: d,
			data:      make(chan string),
			quit:      make(chan struct{}),
		}
		d.workers[i] = &w
	}

	return d
}

// Add value to queue
func (d *DirectoryDispatcher) Add(path string) {
	d.wg.Add(1)
	go d.queueing(path)
}

func (d *DirectoryDispatcher) queueing(path string) {
	d.queue <- path
}

// Wait for worker
func (d *DirectoryDispatcher) Wait() {
	d.wg.Wait()
}

// Start dispacher
func (d *DirectoryDispatcher) Start() {
	for _, w := range d.workers {
		w.start()
	}

	go func() {
		for {
			select {
			case v := <-d.queue:
				worker := <-d.pool
				worker.data <- v
			case <-d.quit:
				return
			}
		}
	}()
}

func (w *directoryWorker) start() {
	go func() {
		for {
			w.dispather.pool <- w

			select {
			case path := <-w.data:
				func() {
					defer w.dispather.wg.Done()

					dl, err := directories(path)
					if err != nil {
						panic(err)
					}

					for _, d := range dl {
						fullpath := filepath.Join(path, d.Name())
						log.Println("directory found: ", fullpath)
						w.dispather.Add(fullpath)
					}
				}()
			}
		}
	}()
}

// directories from path
func directories(path string) ([]os.FileInfo, error) {
	list, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	dl := []os.FileInfo{}
	for _, fi := range list {
		if fi.Mode() != os.ModeSymlink {
			if fi.IsDir() {
				dl = append(dl, fi)
			}
		}
	}

	return dl, nil
}
