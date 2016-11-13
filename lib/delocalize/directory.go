package delocalize

import (
	"fmt"
	"log"
	"path/filepath"
)

type (
	// DirectoryDispatcher management workers
	DirectoryDispatcher struct {
		pool    chan *directoryWorker
		queue   chan string
		workers []*directoryWorker
		quit    chan struct{}
	}

	directoryWorker struct {
		dispather *DirectoryDispatcher
		data      chan string
		quit      chan struct{}
	}
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
	d.queue <- path
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
				dl, err := directories(path)
				fmt.Println(dl)
				if err != nil {
					panic(err)
				}

				for _, d := range dl {
					fullpath := filepath.Join(path, d.Name())
					log.Println(fullpath)
					w.dispather.Add(fullpath)
				}

			}
		}
	}()
}
