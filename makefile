run: main.go
	go run $^
build: main.go
	go build $^
test:
	go test github.com/kaepa3/mikuji/...
