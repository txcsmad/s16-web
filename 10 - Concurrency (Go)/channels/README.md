# channels

Set the env variable `STATTLESHIP_API_TOKEN`. You can get yours from https://www.stattleship.com

Run:

```
$ go run rosters.go
```

Sequentially without channels:

```
real    0m35.834s
user    0m1.713s
sys     0m0.232s
```  
Using channels:

```
real    0m4.895s
user    0m1.690s
sys     0m0.202s
```

