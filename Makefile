GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=playground
BIN_PATH=bin

DOCKER := docker
DOCKER_SUPPORTED_VERSIONS ?= 17|18|19

REGISTRY_PREFIX ?= docker.io/apodemakeles

# set the version number. you should not need to do this
# for the majority of scenarios.
ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --dirty="-dev" --always --tags | sed 's/-/./2' | sed 's/-/./2' )
endif
export VERSION

GO := go
GIT_SHA=$(shell git rev-parse HEAD)
DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")
GOFMT ?= gofmt "-s"

VERSION ?= $(shell git describe --tags --always --match=v* 2> /dev/null)
MASTER_VERSION=V1.00

.PHONY: build clean run all

all: build

hello:
		@echo "Hello"

build:
	    @[ ! -d $(BIN_PATH) ] && mkdir $(BIN_PATH) || echo bin  dir ok
	    @echo build $(BINARY_NAME)
		$(GOBUILD) \
			-o $(BIN_PATH)/$(BINARY_NAME) \
			cmd/$(BINARY_NAME)/main.go

build-debug:
	    @[ ! -d $(BIN_PATH) ] && mkdir $(BIN_PATH) || echo bin  dir ok
	    @echo build $(BINARY_NAME) with debug flag
	    $(GOBUILD) \
			-gcflags "-N -l" \
			-o $(BIN_PATH)/$(BINARY_NAME) \
			cmd/$(BINARY_NAME)/main.go

clean:
	    @echo cleaning
	    rm -rf bin/$(BINARY_NAME) && $(GOCLEAN)

test:
	    $(GOTEST) ./...



.PHONY: fmt
fmt:
	@$(GOFMT) -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: vet
vet:
	$(GO) vet $(PACKAGES)



.PHONY: go.compress.%
go.compress.%:
	$(eval COMMAND := $(word 2,$(subst ., ,$*)))
	$(eval PLATFORM := $(word 1,$(subst ., ,$*)))
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@echo "===========> Compressing binary $(COMMAND) $(VERSION) for $(OS) $(ARCH)"
	upx bin/$(COMMAND)

.PHONY: go.build.%
go.build.%:
	$(eval COMMAND := $(word 2,$(subst ., ,$*)))
	$(eval PLATFORM := $(word 1,$(subst ., ,$*)))
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@echo "===========> Building binary $(COMMAND) $(VERSION) for $(OS) $(ARCH) $(BUILD_INFO)"
	@mkdir -p bin/$(OS)/$(ARCH)
	@CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) $(GO) build -o bin/$(COMMAND) $(BUILD_INFO) cmd/$(COMMAND)/main.go

.PHONY: image
image:	go.build.linux_amd64.playground	go.compress.linux_amd64.playground
	@echo "===========> Building playground $(VERSION) docker image"
	@$(DOCKER) build --pull -t $(REGISTRY_PREFIX)/playground:$(VERSION) .
	@echo "===========> Pushing playground $(VERSION) image to $(REGISTRY_PREFIX)"
	@$(DOCKER) push $(REGISTRY_PREFIX)/playground:$(VERSION)


dockerimage:
	@$(DOCKER) build --pull -t $(REGISTRY_PREFIX)/playground:$(VERSION) . -f Dockerfile.builder --build-arg GIT_SHA=$(GIT_SHA) --build-arg VERSION=$(VERSION) --build-arg DATE=$(DATE)
