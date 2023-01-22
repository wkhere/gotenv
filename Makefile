lib: *.go match.go
	go vet
	go test
	go build

match.go: match.rl
	ragel -Z -G1 match.rl
