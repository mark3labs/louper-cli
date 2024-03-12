.PHONY: all run install clean

run:
	@go rum main.go

install:
	@go install .

clean:
	@rm -fr bin
