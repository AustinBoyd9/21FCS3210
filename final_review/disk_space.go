package main

import (
	"fmt"
	"io/ioutil"
)

type Dir struct {
	path string
	size int64
}

func scan(dir Dir, rpt chan<- Dir) {
	files, _ := ioutil.ReadDir(dir.path)

	// counts the number of subdirectories in the current folder (excluding "." and "..")
	sub := 0
	for _, f := range files {
		if f.IsDir() {
			sub += 1
		} else {
			dir.size += f.Size()
		}
	}

	// TODOd #1 (2 points): create a channel of "Dir" type named "rpt_sub"
	rpt_sub := make(chan Dir)

	for _, f := range files {
		if f.IsDir() {
			dir_sub := Dir{dir.path + "/" + f.Name(), 0}

			// TODOd #2 (2 points): start a "scan" goroutine for each "dir_sub", using "rpt_sub" as the channel
			go scan(dir_sub, rpt_sub)

		}
	}

	// TODOd #3 (5 points): read "rpt_sub" "sub" times, updating "dir.size" with the reports
	for i := 0; i < sub; i++ {
		dir.size += (<-rpt_sub).size
	}

	// TODOd #4 (2 points): use "fmt.Printf" to display the current's folder path and size
	fmt.Printf("%s: %d\n", dir.path, dir.size)

	// TODOd #5 (1 point): write "dir" to "rpt"
	rpt <- dir
}

func main() {
	// creates a channel of "Dir" type named "rpt"
	rpt := make(chan Dir)

	// creates a "Dir" object named "dir" with the folder to be scanned
	dir := Dir{"..", 0}

	// starts a single "scan" goroutine
	go scan(dir, rpt)

	// reads the report
	dir = <-rpt
}
