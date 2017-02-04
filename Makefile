# TODO(leo): have the Makefile add copyright statements to every .go file
# grep for // Copyright (c) in *.go
# if present, delete 9 lines from the file
# inject copyright.txt into .go file
# else inject
# perform during build

SOURCES := $(shell find . -name '*.go')

BINARY=bin/focus
VERSION=1.0
BUILD=`git rev-parse --short HEAD`
BUILD_TIME=`date +%F`

PKG=github.com/lfaoro/focus
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD} -X main.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)
$(BINARY): $(SOURCES)
	@echo "# Vetting..."
	go vet ./...
	@echo "# Formatting..."
	go fmt ./...
	@echo "# Building..."
	go build -x ${LDFLAGS} -o ${BINARY} main.go

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	rm ${BINARY}

v := $(shell grep VERSION Makefile | head -1 | cut -d'=' -f2)
v := $(shell echo "$v + 0.1" | bc)
.PHONY: bump
bump:
	sed -i.old "s/^VERSION=.*/VERSION=${v}/" Makefile
	# git tag v${v}
	# git push --tags
