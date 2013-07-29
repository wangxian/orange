package main

import (
	"github.com/howeyc/fsnotify"
	"path/filepath"
	"strings"
	"time"
	"log"
	"os"
)

func ignoresFilter(ignoresList *[]string, path string) bool {
	for _, v := range *ignoresList {
		if strings.Contains(path, v) {
			return true
		}
	}
	return false
}

var eventTime = make(map[string]time.Time)

func Watcher(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	ignoresList := []string{}
	if Config.ignores != "" {
		ignoresList = strings.Split(Config.ignores, ",")
	}
	// log.Println(ignoresList)


	go func() {
		for {
			select {
			case e := <-watcher.Event:

				changed := true
				if t, ok := eventTime[e.String()]; ok {
					if t.Add(time.Millisecond * 800).After(time.Now()) {
						changed = false
					}
				}
				eventTime[e.String()] = time.Now()

				if changed {
					// log.Println(Clients)
					RefreshBrowser()
				}
				// log.Println(e.String())

			case err := <-watcher.Error:
				log.Fatal("watcher.Error:", err)
			}
		}
	}()

	// walk dirs
	walkFn := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !ignoresFilter(&ignoresList, path) {
			log.Println("Watch DIR:", path)

			err = watcher.Watch(path)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	}

	if err := filepath.Walk(path, walkFn); err != nil {
		log.Println(err)
	}
	// os.Exit(-1)

}