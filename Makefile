BIN = ./node_modules/.bin
.PHONY: all run install clean release release-patch release-minor release-major go-release

run:
	@go rum main.go

install:
	@go install .

clean:
	@rm -fr dist

define release
    VERSION=`node -pe "require('./package.json').version"` && \
    NEXT_VERSION=`node -pe "require('semver').inc(\"$$VERSION\", '$(1)')"` && \
    node -e "\
        var j = require('./package.json');\
        j.version = \"$$NEXT_VERSION\";\
        var s = JSON.stringify(j, null, 2);\
        require('fs').writeFileSync('./package.json', s);" && \
    git commit -m "Version $$NEXT_VERSION" -- package.json && \
    git tag v"$$NEXT_VERSION" -m "Version $$NEXT_VERSION"
endef

release-patch:
	@$(call release,patch)

release-minor:
	@$(call release,minor)

release-major:
	@$(call release,major)

go-release:
	@goreleaser release --clean
	@make clean
