package delocalize

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type (

	// DirectoryDispatcher management workers
	DirectoryDispatcher struct {
		pool    chan *directoryWorker
		queue   chan string
		workers []*directoryWorker
		wg      sync.WaitGroup
		quit    chan struct{}
		deleter FileDeleter
	}

	directoryWorker struct {
		dispather *DirectoryDispatcher
		data      chan string
		quit      chan struct{}
	}

	FileDeleter interface {
		Add(string)
	}
)

// NewDirectoryDispatcher .
func NewDirectoryDispatcher(maxQueues, maxWorkers int, del FileDeleter) *DirectoryDispatcher {
	d := &DirectoryDispatcher{
		pool:    make(chan *directoryWorker, maxWorkers),
		queue:   make(chan string, maxQueues),
		quit:    make(chan struct{}),
		deleter: del,
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

					list, err := ioutil.ReadDir(path)
					if err != nil {
						panic(err)
					}

					for _, fi := range list {
						if fi.Mode() == os.ModeSymlink {
							continue
						}

						fullpath := filepath.Join(path, fi.Name())
						if fi.IsDir() {
							w.dispather.Add(fullpath)
							continue
						} else {
							if IsLocalizedFile(fullpath) {
								w.dispather.deleter.Add(fullpath)
							}
						}
					}
				}()
			}
		}
	}()
}
