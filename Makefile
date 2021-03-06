# Change for your particular project, this is the _root_ level name.
PROJECT_NAME := ypkg-tools

# For example, github.com/ikeydoherty/go-testing
PROJECT_PREFIX = github.com/ikeydoherty
PROJECT_ID := $(PROJECT_PREFIX)/$(PROJECT_NAME)

# Root of your project. For simple libraries/binaries this should just be
# src. For reusable libraries and such this should likely be src/$(PROJECT_ID)
PROJECT_ROOT := src/$(PROJECT_PREFIX)

.DEFAULT_GOAL := all

include Makefile.go

# The resulting binaries map to the subproject names
BINARIES = \
	yauto

# Just for compliance
LIBRARIES = \
	ylib

# We want to add compliance for all built binaries
_CHECK_COMPLIANCE = $(addsuffix .compliant,$(BINARIES)) $(addsuffix .compliant,$(LIBRARIES))

# Build all binaries as static
BINS = $(addsuffix .dynbin,$(BINARIES))
	
# Ensure our own code is compliant..
compliant: $(_CHECK_COMPLIANCE)
install: $(BINS)
	test -d $(DESTDIR)/usr/bin || install -D -d -m 00755 $(DESTDIR)/usr/bin; \
	install -m 00755 builds/* $(DESTDIR)/usr/bin/.

all: $(BINS)
