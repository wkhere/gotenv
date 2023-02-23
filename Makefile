lib: *.go match.go
	go test
	go build

match.go: match.rl
	ragel -Z -G1 match.rl

bench:
	go test -bench=$(sel) -count=5  -benchmem

sel=.

graph:
	ragel -Vp match.rl -o match.dot
	dot -Tsvg match.dot >| match.svg

.PHONY: bench graph
