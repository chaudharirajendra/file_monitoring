package file

import (
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// WatcherInterface defines the behavior of a file watcher.
type WatcherInterface interface {
	Start() error
	FileChan() <-chan string
	Close() error
}

// Watcher implements the WatcherInterface for watching a directory for file events.
type Watcher struct {
	watcher   *fsnotify.Watcher
	fileCh    chan string
	targetDir string
}

// NewWatcher creates a new file watcher.
func NewWatcher(targetDir string) (WatcherInterface, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &Watcher{
		watcher:   watcher,
		targetDir: targetDir,
		fileCh:    make(chan string),
	}, nil
}

// Start starts the file watcher.
func (w *Watcher) Start() error {
	// Start watching the directory
	err := w.watcher.Add(w.targetDir)
	if err != nil {
		return err
	}

	go func() {
		defer w.watcher.Close()

		for {
			select {
			case event, ok := <-w.watcher.Events:
				if !ok {
					return
				}
				if (event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create) && !isSwapFile(event.Name) {
					w.fileCh <- event.Name
				}
			case err, ok := <-w.watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	return nil
}

// FileChan returns the channel for receiving file events.
func (w *Watcher) FileChan() <-chan string {
	return w.fileCh
}

// Close closes the file watcher.
func (w *Watcher) Close() error {
	return w.watcher.Close()
}

func isSwapFile(filename string) bool {
	ext := filepath.Ext(filename)
	return ext == ".swp"
}
