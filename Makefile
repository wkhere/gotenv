lib: *.go match.go
	go vet
	go test
	go build

match.go: match.rl
	ragel -Z -G1 match.rl

bench:
	go test -bench=. -count=5  -benchmem

graph:
	ragel -Vp match.rl -o match.dot
	dot -Tsvg match.dot >| match.svg

.PHONY: bench graph
