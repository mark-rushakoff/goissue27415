# goissue27415

This is a reduced test case to help with [Go issue #27415](https://github.com/golang/go/issues/27415).

By running the benchmarks 10 times:

```
go test -bench=. -count=10 .
```

And comparing go versions:

```
go version go1.11 darwin/amd64
go version devel +be10ad7622 Wed Aug 22 22:26:40 2018 +0000 darwin/amd64
```

I observed the following with benchstat:

```
name           old time/op  new time/op  delta
BboltWrites-4   281µs ± 1%  8283µs ±11%  +2843.16%  (p=0.000 n=8+10)
```
