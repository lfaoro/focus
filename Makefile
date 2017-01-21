SOURCES := $(shell find . -name '*.go')

BINARY=focus
VERSION=1.0
BUILD=`git rev-parse --short HEAD`
BUILD_TIME=`date +%F`

PKG=github.com/lfaoro/focus
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD} -X main.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)
$(BINARY): $(SOURCES)
	go build ${LDFLAGS} -o ${BINARY} main.go

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

v := $(shell grep VERSION Makefile | head -1 | cut -d'=' -f2)
v := $(shell echo "$v + 0.1" | bc)
.PHONY: bump
bump:
	sed -i.old "s/^VERSION=.*/VERSION=${v}/" Makefile
	# git tag v${v}
	# git push --tags
