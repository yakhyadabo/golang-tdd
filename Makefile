default: build

build: 
	go build src/main/calculator/*.go

run: 
	go run src/main/calculator/calculator.go
	
test: 
	go test src/main/calculator/calculator_test.go src/main/calculator/calculator.go
