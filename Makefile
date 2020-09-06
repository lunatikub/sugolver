all: sugolver

sugolver: 
	go build -o sugovler

test:
	./test/test_sugolver.sh

unit-test:
	go test -v ./...

clean:
	rm sugovler

.PHONY: test unit-test sugolver clean
