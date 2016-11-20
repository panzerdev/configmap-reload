package main

import (
	"flag"
	"log"
	"os"

	fsnotify "gopkg.in/fsnotify.v1"
	"os/exec"
	"path/filepath"
	"strings"
)

var folder = flag.String("folder", "", "Folder to watch for changes in it")
var shCommand = flag.String("command", "", "Command beeing executed on trigger. Arguments need to be seperated by ',' like 'nginx,-s,reload'")
var containerName = flag.String("container", "", "Container name in pod")

func main() {
	flag.Parse()
	if *folder == "" {
		log.Println("Missing folder")
		flag.Usage()
		os.Exit(1)
	}
	if *shCommand == "" {
		log.Println("Missing shell command to execute")
		flag.Usage()
		os.Exit(1)
	}
	if *containerName == "" {
		log.Println("Missing container name in the pod")
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
		for {
			select {
			case event := <-watcher.Events:
				log.Println("Something happend:", event)
				if event.Op == fsnotify.Create && filepath.Base(event.Name) == "..data" {
					log.Println("Watched file changed")
					// command exec, pod name aka HOSTNAME, container name
					args := []string{"exec", os.Getenv("HOSTNAME"), "-c", *containerName}
					// optional NAMESPACE if present
					nameSpace := os.Getenv("NAMESPACE")
					log.Println("Namespace found", nameSpace != "", "with value", nameSpace)
					if nameSpace != "" {
						args = append(args, "-n", nameSpace)
					}
					// command splitted by commas
					splitted := strings.Split(*shCommand, ",")
					args = append(args, "--")
					args = append(args, splitted...)

					log.Println("Command executed:", "kubectl", args)
					// execute
					b, err := exec.Command("kubectl", args...).CombinedOutput()
					if err != nil {
						log.Println(err)
					}
					log.Println("Command result", string(b))
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(*folder)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
