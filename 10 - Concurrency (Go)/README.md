# 10. Concurrency (Go)

Go makes writing concurrent programs safe, fun, and quick. The standard
library's `sync` package contains structures used to write concurrent programs.
In addition, Go has built-in language features like channels and goroutines that
can be used to communicate between two or more concurrent processes.

## Wait, what is concurrency?

Normally, programs have a single thread of execution. For example:

```go
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
```

The above program opens a single file and prints its contents to the terminal.

Now consider the case where you may want to open 10 files and print their
contents to the terminal. As you may have heard, opening a file is a blocking
operation since a system call has to be made to the operating system to request
the file. This means that while waiting for a file to open, the remaining files
have to wait their turn doing nothing. 

Using concurrency, you can write programs that open the 10 file simultaneously.
In Go, this can be achieved using the `go` keyword -- to start a goroutine that
runs on its own thread (sort of). Now, other files do not have to wait idly
during the blocking open operation.


```go
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
```

Each invocation of the function followed by the `go` keyword is executed on its
own thread (sort of), so now all the files are opened concurrently without
having to wait for the previous.

But there we had to do some additionally synchronization to wait for all
the spawned goroutines to finish their work before the main program exited.
This is where the `sync` package comes in (in this program, specifically the
WaitGroup).

## Today's workshop

Topics

- Locks (hit-counters examples)
- goroutines, waitgroups (see goroutines examples)
- Channels (time permitting)

## Author

Nishanth Shanmugham

- Email: [nishanths@utexas.edu](mailto:nishanths@utexas.edu)
- GitHub: [nishanths](https://github.com/nishanths)

## License

[MIT](http://nishanths.mit-license.org)
