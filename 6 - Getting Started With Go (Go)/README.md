# 6 - Getting Started With Go

Go is an open-source statically typed programming knowledge. It has excellent
support for:

- concurrency
- server and network programming
- excellent toolchain (we discussed `go fmt`, `go vet`, `go get`)

## Installing Go

- Head over to <https://golang.org/dl/>
- Download the binary release for your platform
- Follow the latest [instructions](https://golang.org/doc/install) to install Go and setup the `GOPATH` environment variable

## Playing around

Take the tour: [tour.golang.org](https://tour.golang.org) or use the [playground](https://play.golang.org)

## Workshop

See the code we wrote in the repo.

To run locally, clone or save the file to a location in your `$GOPATH/src`.
For example, I used: `$GOPATH/src/github.com/nishanths/example/main.go`

Then:

```
$ cd $GOPATH/src/github.com/nishanths/example/
$ go run main.go
```

Or alternatively:

```
$ cd $GOPATH/src/github.com/nishanths/example/
$ go build
$ ./main
```
