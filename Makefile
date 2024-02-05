BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean

EXEBUILD=main.go
EXEFILE=main
EXENAME=$(notdir $(CURDIR))

echo:
	@echo "buildpath $(BUILDPATH)"
	@echo "go $(GO)"
	@echo "gobuild $(GOBUILD)"
	@echo "goinstall $(GOINSTALL)"
	@echo "goclean $(GOCLEAN)"
	@echo "exec build file $(EXEBUILD)"
	@echo "exec file $(EXEFILE)"
	@echo "new exec name $(EXENAME)"
	@echo "man name $(MANNAME)"

makedir:
	export GOPATH=$(CURDIR)
	@echo "start building tree..."
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg; fi

build:
	export GOPATH=$(CURDIR)
	@echo "start building executable..."
	GOARCH=amd64 GOOS=linux go build -o ./bin/$(EXENAME)-linux
	@echo "completed..."

clean:
	@echo "cleaning up..."
	@rm -rf $(BUILDPATH)/bin/$(EXENAME)-linux
	#@rm -rf $(BUILDPATH)/pkg

