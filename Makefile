LATEST-GIT-SHA := $(shell git rev-parse HEAD)
VERSION := $(shell cat VERSION)
FLAGS := -X main.commit=$(LATEST-GIT-SHA) -X main.version=$(VERSION)
BIN := packer-plugin-veertu-anka
ARCH := amd64
OS_TYPE ?= $(shell uname -s | tr '[:upper:]' '[:lower:]')

.PHONY: lint.golang lint.packer go.test anka.test anka.clean anka.clean-images

.DEFAULT_GOAL := help

#help:	@ List available tasks on this project
help:
	@grep -h -E '[a-zA-Z\.\-]+:.*?@ .*$$' $(MAKEFILE_LIST) | sort | tr -d '#' | awk 'BEGIN {FS = ":.*?@ "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#lint.golang:		@ Run `golangci-lint run` against the current code
lint.golang:
	golangci-lint run --fast

#lint.packer:  @ Run `packer validate` against packer definitions
lint.packer:
	packer validate examples/create-from-installer.pkr.hcl
	packer validate examples/create-from-installer-with-post-processing.pkr.hcl
	packer validate examples/clone-existing.pkr.hcl
	packer validate examples/clone-existing-with-post-processing.pkr.hcl
	packer validate examples/clone-existing-with-port-forwarding-rules.pkr.hcl
	packer validate examples/clone-existing-with-hwuuid.pkr.hcl
	packer validate examples/clone-existing-with-expect-disconnect.pkr.hcl
	packer validate examples/clone-existing-with-use-anka-cp.pkr.hcl

#go.test:		@ Run `go test` against the current tests
go.test:
	go test -v builder/anka/*.go
	go test -v post-processor/ankaregistry/*.go

#go.hcl2spec:		@ Run `go generate` to generate hcl2 config specs
go.hcl2spec:
	GOOS=$(OS_TYPE) go install github.com/hashicorp/packer/cmd/mapstructure-to-hcl2
	GOOS=$(OS_TYPE) PATH="$(shell pwd):${PATH}" go generate builder/anka/config.go
	GOOS=$(OS_TYPE) PATH="$(shell pwd):${PATH}" go generate post-processor/ankaregistry/post-processor.go

build-linux:
	GOOS=linux OS_TYPE=linux $(MAKE) go.build

build-mac:
	GOOS=darwin OS_TYPE=darwin $(MAKE) go.build

#go.build:		@ Run `go build` to generate the binary
go.build: $(BIN)
$(BIN): go.hcl2spec
	GOARCH=$(ARCH) go build $(RACE) -ldflags "$(FLAGS)" -o bin/$(BIN)_$(OS_TYPE)_$(ARCH)
	chmod +x bin/$(BIN)_$(OS_TYPE)_$(ARCH)

#packer.install:		@ Copy the binary to the packer plugins folder
packer.install: $(BIN)
	mkdir -p ~/.packer.d/plugins/
	cp $(BIN) ~/.packer.d/plugins/

#anka.build-and-install:		@ Run make targets to setup the initialize the binary
anka.build-and-install: $(BIN)
	$(MAKE) anka.clean
	$(MAKE) go.build
	$(MAKE) packer.install

#anka.packer-test:		@ Run `packer build` with the default .pkr.hcl file
anka.test: lint.packer packer.install
	PACKER_LOG=1 packer build examples/create-from-installer.pkr.hcl

#anka.clean:		@ Remove the binary
anka.clean:
	rm -f $(BIN)

#anka.clean-images:		@ Remove all anka images with `anka delete`
anka.clean-images:
	anka --machine-readable list | jq -r '.body[].name' | grep anka-packer | xargs -n1 anka delete --yes

#anka.clean-clones:		@ Remove all anka clones with `anka delete`
anka.clean-clones:
	anka --machine-readable list | jq -r '.body[].name' | grep anka-packer | grep -v base | xargs -n1 anka delete --yes

#anka.wipe-anka:		@ Remove all anka images from the local library
anka.wipe-anka:
	-rm -rf ~/Library/Application\ Support/Veertu
	-rm -rf ~/.anka