# hit-counter-w-lock

Hit counter with lock.

Counts the number of visitors to a website.

## Usage:

```
# go run app.go
```

Then using `curl` or similar:

```
$ curl http://localhost:8080
```

Or to simply check the count without increasing it:

```
$ curl http://localhost:8080/status
```

To reset:

```
$ curl -X POST http://localhost:8080/reset
```

## Why is the mutex important

We generated 1000 requets with 100 running concurrently. The total number of
hits would likely be less than the expected 1000 without the lock.

In the workshop we used [`rakyll/boom`](https://github.com/rakyll/boom) to show this and then checked the status.

```
$ boom -n 1000 -c 100 http://localhost:8080
```
