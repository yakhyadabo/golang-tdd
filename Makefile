default: build

build: 
	go build src/main/calculator.go

run: 
	go run src/main/calculator.go
	
test: 
	go test -v src/test/calculator_test.go
