#Package info
NAME := medaka
VERSION := v1.0.0
REVISION := $(shell git rev-parse --short HEAD || echo "unsupported")

LDFLAGS = -ldflags="-s -w -X \"github.com/beito123/medaka.Version=$(VERSION)\" -X \"github.com/beito123/medaka.Revision=$(REVISION)\" -extldflags \"-static\""
# LDFLAGS := -ldflags=" -extldflags \"-static\""

#Build path
BUILDPATH := ./app/medaka/
ASSETPATH := ./data

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

GOASSETBUILDER=go-assets-builder

# For compatibility
ifeq ($(OS),Windows_NT)
	BINARYNAME = $(NAME).exe
	RM = cmd.exe /C del /Q
	RMDIR = $(RM)
	SRCS = $(subst $(CURDIR)/,,$(subst \,/,$(shell cmd.exe /C "dir *.go /b /a-d /s")))
else
	BINARYNAME = $(NAME)
	RM = rm -f
	RMDIR = rm -rf
	SRCS := $(shell find . -type f -name '*.go')
endif

# Commands
.PHONY: all medaka

all: medaka

medaka: app/medaka

app/medaka: $(SRCS)
	@echo "Ready assets..."
	@cd $(ASSETPATH); \
		$(GOASSETBUILDER) --package=data ./static/ > assets.go
	@echo "Building..."
	@$(GOBUILD) -a -tags netgo -installsuffix netgo $(LDFLAGS) -o $(BINARYNAME) $(BUILDPATH)

.PHONY: install clean test deps cross-build

install:
	@$(GOINSTALL) $(LDFLAGS) $(BUILDPATH)

clean:
	$(GOCLEAN)
	@$(RM) $(NAME)
	@$(RM) $(NAME).exe
	@$(RMDIR) ./dist

test:
	$(GOTEST) -race $(BUILDPATH)

deps:
	dep ensure
	@echo "Installing go-assets-builder..."
	@cd ./vendor/github.com/jessevdk/go-assets-builder
	@$(GOINSTALL) .

cross-build: src
	@echo "Ready..."
	@$(GOGET) github.com/mitchellh/gox

	@echo "Building..."
	@gox \
		-os="darwin linux windows" \
		-arch="386 amd64" \
		-output "dist/{{.OS}}_{{.Arch}}/{{.Dir}}" \
		$(BUILDPATH)
