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
name                old time/op  new time/op  delta
BboltWrites/10-4     298µs ±15%  8134µs ± 6%  +2626.85%  (p=0.000 n=10+9)
BboltWrites/100-4    280µs ± 1%  8001µs ± 2%  +2758.37%  (p=0.000 n=8+8)
BboltWrites/1000-4   311µs ± 7%  8304µs ± 5%  +2570.40%  (p=0.000 n=10+9)
```
