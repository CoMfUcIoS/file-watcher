package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	var source string
	var action string
	var version bool

	flag.StringVar(&source, "source", "", "Source file or directory to watch")
	flag.StringVar(&action, "action", "", "CLI command to execute on change")
	flag.BoolVar(&version, "version", false, "Print the version and exit")
	flag.Parse()

	if version {
		fmt.Println("Version: v0.5.0")
		os.Exit(0)
	}
	if source == "" || action == "" {
		flag.Usage()
		os.Exit(1)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		var debounceTimer *time.Timer
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove) != 0 {
					if debounceTimer != nil {
						debounceTimer.Stop()
					}
					debounceTimer = time.AfterFunc(500*time.Millisecond, func() {
						fmt.Println("Executing action:", action)
						cmd := exec.Command("sh", "-c", action)
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						err := cmd.Run()
						if err != nil {
							log.Println("Error executing action:", err)
						}
					})
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	<-done
}
