.PHONY: all run install clean release

run:
	@go rum main.go

install:
	@go install .

clean:
	@rm -fr dist

release:
	@goreleaser release --clean
	@make clean
