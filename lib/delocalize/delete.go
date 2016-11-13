package delocalize

import (
	"fmt"
	"log"
	"os"
)

// @see: https://gist.github.com/kaneshin/69bd13c7b57ba8bac84fb4de0098b5fc

type (
	// DeleteDispatcher management worker
	DeleteDispatcher struct {
		pool    chan *deleteWorker
		queue   chan string
		workers []*deleteWorker
		quit    chan struct{}
	}

	deleteWorker struct {
		dispatcher *DeleteDispatcher
		data       chan string
		quit       chan struct{}
	}
)

var (
	// DryRun will not delete files
	DryRun = true
)

// NewDeleteDispatcher .
func NewDeleteDispatcher(maxQueues, maxWorkers int) *DeleteDispatcher {
	d := &DeleteDispatcher{
		pool:  make(chan *deleteWorker, maxWorkers),
		queue: make(chan string, maxQueues),
		quit:  make(chan struct{}),
	}

	// worker の初期化
	d.workers = make([]*deleteWorker, cap(d.pool))
	for i := 0; i < cap(d.pool); i++ {
		w := deleteWorker{
			dispatcher: d,
			data:       make(chan string),
			quit:       make(chan struct{}),
		}
		d.workers[i] = &w
	}
	return d
}

// Add value to queue for worker
func (d *DeleteDispatcher) Add(path string) {
	d.queue <- path
}

// Start dispather
func (d *DeleteDispatcher) Start() {
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

// start worker
func (w *deleteWorker) start() {
	go func() {
		for {
			w.dispatcher.pool <- w

			select {
			case path := <-w.data:
				if DryRun {
					fmt.Println(path)
				} else {
					err := delete(path)
					if err != nil {
						log.Print(err)
					}
				}
			case <-w.quit:
				return
			}
		}
	}()
}

// delete path name file
func delete(path string) error {
	if !IsLocalizedFile(path) {
		return ErrThisFileIsNotLocalizedFile
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}
