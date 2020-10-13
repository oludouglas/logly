short:
	go test -v -short ./... -cover

bench:
	go test -v -bench=. ./...