package main

import (
	"io"
	"log"
	"os"
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

	for _, filename := range filenames {
		print(filename)
	}
}
