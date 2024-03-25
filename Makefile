OSVC_CONTEXT =

GOCMD ?= go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGEN := $(GOCMD) generate
GOVET := $(GOCMD) vet

STRIP := /usr/bin/strip
MKDIR := /usr/bin/mkdir
INSTALL := /usr/bin/install
PREFIX ?= /usr

DIST := dist
OM := bin/om
OX := bin/ox
COMPOBJ := bin/compobj
COMPOBJ_D := share/opensvc/compliance

VERSION := $(shell git describe --tags --abbrev)

.PHONY: version dist

all: clean vet test race build dist

build: version api om ox compobj

api:
	$(GOGEN) ./daemon/api

clean:
	$(GOCLEAN)
	$(GOCLEAN) -testcache
	rm -f $(OM) $(OX)

om:
	$(GOBUILD) -o $(OM) -v ./cmd/om/

ox:
	$(GOBUILD) -o $(OX) -v ./cmd/ox/

compobj:
	$(GOBUILD) -o $(COMPOBJ) -v ./util/compobj/

test:
	$(GOTEST) -p 1 -timeout 60s ./...

testinfo:
	TEST_LOG_LEVEL=info $(GOTEST) -p 1 -timeout 60s ./...

race:
	$(GOTEST) -p 1 -timeout 240s ./... -race

vet:
	$(GOVET) ./...

install:
	$(MKDIR) -p $(PREFIX)/bin
	$(MKDIR) -p $(PREFIX)/$(COMPOBJ_D)
	$(INSTALL) -m 755 $(OM) $(PREFIX)/$(OM)
	$(INSTALL) -m 755 $(OX) $(PREFIX)/$(OX)
	$(INSTALL) -m 755 $(COMPOBJ) $(PREFIX)/$(COMPOBJ)
	$(PREFIX)/$(COMPOBJ) -i $(PREFIX)/$(COMPOBJ_D)

version:
	echo $(VERSION) >util/version/text/VERSION

dist:
	$(MKDIR) -p $(DIST)/bin
	$(MKDIR) -p $(DIST)/$(COMPOBJ_D)
	$(INSTALL) -m 755 $(OM) $(DIST)/$(OM)
	$(INSTALL) -m 755 $(OX) $(DIST)/$(OX)
	$(INSTALL) -m 755 $(COMPOBJ) $(DIST)/$(COMPOBJ)
	$(DIST)/$(COMPOBJ) -r -i $(DIST)/$(COMPOBJ_D)
	$(STRIP) --strip-all $(DIST)/$(OM) $(DIST)/$(OX) $(DIST)/$(COMPOBJ)
	cd $(DIST) && tar czvf opensvc-$(VERSION).tar.gz $(OM) $(OX) $(COMPOBJ) $(COMPOBJ_D) && cd -


