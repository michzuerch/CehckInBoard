BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean
GOGET=$(GO) get

EXENAME=main

export GOPATH=$(CURDIR)

hello:
	echo "Hello from Makefile..."

makedir:
	@echo "Start building tree..."
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi

clean:
	@echo "Cleaning"
	@rm -rf $(BUILDPATH)/bin/$(EXENAME)
	@rm -rf $(BUILDPATH)/pkg

get:
	$(GOGET) github.com/gorilla/mux
	

build:
	@echo "Start building..."
	$(GOINSTALL) $(EXENAME)
	@echo  "Building done."

all: makedir get build

run:
	go run main.go

print: $(wildcard *.go2)
	ls -la $?

all: hello build
