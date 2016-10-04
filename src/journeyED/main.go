package main

import (
	"bufio"
	"encoding/json"
	"event"
	"eventhandler"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	var wg sync.WaitGroup
	handleLine := eventhandler.HandleLine{Handlers: make(map[string]func(*event.Event, bool) string)}
	eventhandler.HandleEventInit(&handleLine)

	done := make(chan bool)
	go func() {
		var f *os.File = nil
		var r *bufio.Reader = nil
		var e event.Event
		wg.Add(1)
		for {
			select {
			// if we get an event
			case event := <-watcher.Events:
				log.Println("event:", event)
				// and if it is writing to a file
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					// 1st run file exists reader doesnt
					if r == nil {
						// open the file
						f, err = os.Open(event.Name)
						if err != nil {
							log.Println("Could not open file, ", err)
							// quit
							wg.Done()
							done <- true
							continue
						}
						// create the reader
						r = bufio.NewReader(f)
					}
					// read the stuff from the file
					b, err := r.ReadBytes('\n')
					if err != nil && err != io.EOF {
						log.Println("Could not read a line, ", err)
						wg.Done()
						done <- true
						continue
					}
					// convert it to json and pass it to a handler
					fmt.Println(string(b))
					err = json.Unmarshal(b, &e)
					if err != nil {
						log.Println(err)
						wg.Done()
						done <- true
						continue
					} else {
						handleLine.HandleEvent(&e)
					}
				}
				// otherwise if a file was created then
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("created file:", event.Name)
					// if file is open close it cuz we now have a new file to read from
					if f != nil {
						f.Close()
					}
					// open new file
					f, err := os.Open(event.Name)
					if err != nil {
						log.Println("Could not open file, ", err)
						wg.Done()
						// quit
						done <- true
						continue
					}
					// make reader read from file
					if r == nil {
						r = bufio.NewReader(f)
					} else {
						r.Reset(f)
					}
				}
			// errors! oh no!
			case er := <-watcher.Errors:
				log.Println("error:", er)
			case <-done:
				return
			}
		}
	}()

	err = watcher.Add(os.Getenv("USERPROFILE") + "\\Documents\\Saved Games\\Frontier Developments\\Elite Dangerous\\")
	if err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
