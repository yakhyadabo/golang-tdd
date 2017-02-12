default: run

run: 
	go run src/main/calculator.go

test: 
	go test -v src/test/calculator_test.go
