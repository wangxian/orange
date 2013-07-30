package main

import (
	"github.com/howeyc/fsnotify"
	"path/filepath"
	"os/exec"
	"strings"
	"bytes"
	"time"
	"log"
	"os"
)

func BeforePreload(filename string) {
	if len(Config.precmd) == 0 {
		return
	}

	precmd := strings.Replace(Config.precmd, "{0}", filename, -1)
	PrecmdList := strings.Split(precmd, " ")
	// log.Println(PrecmdList[1:])

	cmd := exec.Command(PrecmdList[0], PrecmdList[1:]...)
	// cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
	    log.Fatal(err)
	}
	println(out.String())
}

func ignoresFilter(ignoresList *[]string, path string) bool {
	for _, v := range *ignoresList {
		if strings.Contains(path, v) {
			return true
		}
	}
	return false
}

// fixed: File change soon
var eventTime = make(map[string]time.Time)

func Watcher(path string) {
	ignoresList := []string{}
	if Config.ignores != "" {
		if Config.ignores == "." {
			log.Println("watch ignore all")
			return
		}
		ignoresList = strings.Split(Config.ignores, ",")
	}
	ignoresList = append(ignoresList, ".git")
	ignoresList = append(ignoresList, ".svn")
	log.Println("ignoresList:", ignoresList)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

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
					log.Println(e.String())
					RefreshBrowser()

					BeforePreload(e.Name)
					//@todo: do someting here, eg: precommand
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