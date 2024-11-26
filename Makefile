GOFILE = misskey-cli.go
DIR 		= ./build
NAME  	= misskey-cli
GHNAME = github.com/mikuta0407/misskey-cli
TAGVER = $(shell git describe --tags)
LDFLAGS = -ldflags "-X cmd.name=$(GHNAME) -X cmd.version=$(TAGVER)"

LOCALBUILD = go build -o $(DIR)/$(NAME) $(LDFLAGS) $(GOFILE)
BUILDCMD = GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(DIR)/$(NAME)_$(GOOS)_$(GOARCH) $(GOFILE)
WINBUILDCMD = GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(DIR)/$(NAME)_$(GOOS)_$(GOARCH).exe $(GOFILE)

.PHONY: all
all: linux-arm64 linux-arm linux-riscv64 linux-386 linux-amd64 darwin-amd64 darwin-arm64 windows-386 windows-amd64 windows-arm64

.PHONY: clean
clean:
	rm -rf $(DIR)/*

.PHONY: build
build:
	$(LOCALBUILD)

.PHONY: linux-arm64
linux-arm64:
	$(eval GOOS := linux)
	$(eval GOARCH := arm64)
	$(BUILDCMD)

.PHONY: linux-arm
linux-arm:
	$(eval GOOS := linux)
	$(eval GOARCH := arm)
	$(BUILDCMD)

.PHONY: linux-riscv64
linux-riscv64:
	$(eval GOOS := linux)
	$(eval GOARCH := riscv64)
	$(BUILDCMD)

.PHONY: linux-386
linux-386:
	$(eval GOOS := linux)
	$(eval GOARCH := 386)
	$(BUILDCMD)

.PHONY: linux-amd64
linux-amd64:
	$(eval GOOS := linux)
	$(eval GOARCH := amd64)
	$(BUILDCMD)

.PHONY: darwin-amd64
darwin-amd64:
	$(eval GOOS := darwin)
	$(eval GOARCH := amd64)
	$(BUILDCMD)

.PHONY: darwin-arm64
darwin-arm64:
	$(eval GOOS := darwin)
	$(eval GOARCH := arm64)
	$(BUILDCMD)

.PHONY: windows-386
windows-386:
	$(eval GOOS := windows)
	$(eval GOARCH := 386)
	$(WINBUILDCMD)

.PHONY: windows-amd64
windows-amd64:
	$(eval GOOS := windows)
	$(eval GOARCH := amd64)
	$(WINBUILDCMD)

.PHONY: windows-arm64
windows-arm64:
	$(eval GOOS := windows)
	$(eval GOARCH := arm64)
	$(WINBUILDCMD)
