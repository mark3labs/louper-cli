.PHONY: all build install clean

build:
	@go build -o bin/louper main.go

install:
	@go install .

clean:
	@rm -fr bin
