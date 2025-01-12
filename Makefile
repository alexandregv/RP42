NAME = RP42
MODULE = github.com/alexandregv/$(NAME)

# Compiler & Preprocessor flags
LDFLAGS		+= -s -w
MAKEFLAGS	+= --no-print-directory

VERSION ?= $(shell git describe --tags --dirty --broken)

.DEFAULT_GOAL = help


##@ Make standards
$(NAME): all

re: clean build  ## Clean and Build the Go binary

clean:           ## Delete the Go binary, if any
	rm -f $(NAME)


##@ Utilities
help:     ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z_\-\.]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

get-tag:      ## Get the next SemVer tag based on commits
	go tool github.com/caarlos0/svu

get-version:  ## Get the next version based on repo status (commit hash, dirty, broken)
	# git describe --tags --dirty --broken
	@echo $(VERSION)

commit:       ## Make a commit following the Conventional Commits convention
	go tool github.com/stefanlogue/meteor

tag:          ## Make a SemVer tag based on commits (make get-tag)
	git tag $(shell go tool github.com/caarlos0/svu)


##@ Build
all: linux macos windows ## Build for all OS

linux:       ## Build the Go binary
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/linux/$(NAME) -ldflags="$(LDFLAGS)" .

macos:       ## Build the Go binary
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o build/macos/$(NAME) -ldflags="$(LDFLAGS)" .

windows:     ## Build the Go binary
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o build/windows/$(NAME).exe -ldflags="$(LDFLAGS)" .


##@ Checks (tests, linters, etc)
test:        ## Run Go tests
	go test -v ./...
	go test -v ./... -json | go tool github.com/mfridman/tparse -all

cover:       ## Run Go tests with coverage
	go test -cover -v ./...

lint:        ## Run linters and fix code, when possible (golangci-lint)
	go tool github.com/golangci/golangci-lint/cmd/golangci-lint run --show-stats --fix

check-lint:  ## Run linters in read-only (golangci-lint)
	go tool github.com/golangci/golangci-lint/cmd/golangci-lint run --show-stats

check: check-lint pre-commit-run test cover  ## Run all checks


##@ Pre-commit hooks
pre-commit-install:                           ## Install pre-commit hooks locally
	pre-commit install

pre-commit-update:                            ## Update pre-commit hooks to the latest version
	pre-commit autoupdate

pre-commit-run: pre-commit-install            ## Run pre-commit hooks on all files
	pre-commit run --all-files

pre-commit: pre-commit-install pre-commit-run ## Install and run pre-commit hooks on all files


##@ Make utils
genphony:  ## Generate .PHONY target with all Makefile targets
	echo .PHONY: $$(grep -E '^[A-Za-z0-9\-]+:' Makefile | rev | cut -d: -f2- | rev | grep -v phony) >> Makefile
	# don't forget to remove the old .PHONY line

.PHONY: re clean help get-tag get-version commit tag all linux macos windows test cover lint check-lint check pre-commit-install pre-commit-update pre-commit-run pre-commit
