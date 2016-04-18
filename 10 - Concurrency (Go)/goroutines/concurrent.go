package main

import (
	"io"
	"log"
	"os"
	"sync"
)

func print(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, f)
}

func main() {
	filenames := []string{"foo.txt", "bar.txt", "goo.txt"}
	var wg sync.WaitGroup
	wg.Add(len(filenames))

	for _, filename := range filenames {
		go func(filename string) {
			print(filename)
			wg.Done()
		}(filename)
	}

	wg.Wait()
}
