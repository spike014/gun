package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if !(event.Op&fsnotify.Chmod == fsnotify.Chmod) {
					kill()
					err := build(base)
					if err != nil {
						return
					}
					go run(base)
					updateFilespath()
					go updateWatchFiles(watcher)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Watch files ERROR:")
				panic(err)
			}
		}
	}()
	updateWatchFiles(watcher)
	<-done
}

func updateWatchFiles(watcher *fsnotify.Watcher) {
	for i := 0; i < len(*filepaths); i++ {
		err := watcher.Add((*filepaths)[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}
