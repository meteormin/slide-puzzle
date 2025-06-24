PRJ_NAME=Slide Puzzle
AUTHOR="Meteormin \(miniyu97@gmail.com\)"
PRJ_BASE=$(shell pwd)
PRJ_DESC=$(PRJ_NAME) Deployment and Development Makefile.\n Author: $(AUTHOR)

SUPPORTED_OS=linux darwin
SUPPORTED_ARCH=amd64 arm64

mod ?= "cli"

.DEFAULT: help
.SILENT:;

##help: helps (default)
.PHONY: help
help: Makefile
	echo ""
	echo " $(PRJ_DESC)"
	echo ""
	echo " Usage:"
	echo ""
	echo "	make {command}"
	echo ""
	echo " Commands:"
	echo ""
	sed -n 's/^##/	/p' $< | column -t -s ':' |  sed -e 's/^/ /'
	echo ""

# OS와 ARCH가 정의되어 있지 않으면 기본값을 설정합니다.
# uname -s는 OS 이름(예: Linux, Darwin 등)을 반환하고, tr를 통해 소문자로 변환합니다.
OS ?= $(shell uname -s | tr '[:upper:]' '[:lower:]')
# 아키텍처 정보를 반환합니다. (예: amd64, arm64 등)
ARCH := $(shell ./scripts/detect-arch.sh)
LDFLAGS=-ldflags "-linkmode external -extldflags -static"

##build os={os [linux, darwin]} arch={arch [amd64, arm64]} mod={entrypoint [cli]}: build application
.PHONY: build
build: os ?= $(OS)
build: arch ?= $(ARCH)
build: mod ?= "cli"
build:
	@echo "[build] Building $(mod) for $(os)/$(arch)"
ifeq ($(os), linux)
	CC=$(cc) CGO_ENABLED=1 GOOS=$(os) GOARCH=$(arch) go build $(LDFLAGS) -o build/$(mod)-$(os)-$(arch) ./cmd/$(mod)/main.go
else
	CC=$(cc) CGO_ENABLED=1 GOOS=$(os) GOARCH=$(arch) go build -o build/$(mod)-$(os)-$(arch) ./cmd/$(mod)/main.go
endif

##build-wasm: build wasm
.PHONY: build-wasm
build-wasm:
	@echo "[build-wasm] Building wasm"
	GOOS=js GOARCH=wasm go build -o ./webui/main.wasm ./webui/main.go

##build-docker platform={build for platform} mod={entrypoint [cli]} tag={tag [v1.0.0]}: build docker image
.PHONY: build-docker
build-docker: platform ?= "linux/$(ARCH)"
build-docker: tag ?= "latest"
build-docker:
	@echo "[build-docker] Build docker image"
	@echo " *platform: $(platform)"
	@echo " *tag: $(tag)"
	docker build --platform $(platform) --tag "slide-puzzle:$(tag)" .

##clean: clean application
.PHONY: clean
clean:
	@echo "[clean] Cleaning build directory"
	rm -rf build/*

##clean-docker: clean docker
.PHONY: clean-docker
clean-docker:
	rm -rf .docker/*
	./scripts/docker-clean.sh

##run mod={entrypoint [cli]}: run application
.PHONY: run
run: mod ?= "cli"
run:
	@echo "[run] running application"
	@echo "mod: $(mod)"
	@if [ "$(mod)" == "cli" ]; then \
		go run ./cmd/$(mod)/main.go; \
	elif [ "$(mod)" == "web" ]; then \
  		$(MAKE) build-wasm; \
		go run .cmd/$(mod)/main.go; \
	else \
	  		echo "Unknown mod: $(mod). Please use 'cli' or 'web'." && exit 1; \
 	fi