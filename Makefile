
# put binary in /bin
# create .tar archive
# tag based on version
# update homebrew

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
#BUILDPATH=$(CURDIR)
#GO=$(shell which go)
#GOINSTALL=$(GO) install
#GOCLEAN=$(GO) clean
#GOGET=$(GO) get
#
#EXENAME=focus
#
#export GOPATH=$(CURDIR)
#
#myname:
#	@echo "I am a makefile"
#
#makedir:
#	@echo "start building tree..."
#	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
#	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi
#
#get:
#	#@$(GOGET) github.com/Sirupsen/logrus
#
#build:
#	@echo "start building..."
#	$(GOINSTALL) $(EXENAME)
#	@echo "Yay! all DONE!"
#
#clean:
#	@echo "cleanning"
#	@rm -rf $(BUILDPATH)/bin/$(EXENAME)
#	@rm -rf $(BUILDPATH)/pkg
#	@rm -rf $(BUILDPATH)/src/github.com
#
#all: makedir get build
