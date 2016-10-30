# Change for your particular project, this is the _root_ level name.
PROJECT_NAME := ypkg-tools

# For example, github.com/ikeydoherty/go-testing
PROJECT_PREFIX = github.com/ikeydoherty
PROJECT_ID := $(PROJECT_PREFIX)/$(PROJECT_NAME)

# Root of your project. For simple libraries/binaries this should just be
# src. For reusable libraries and such this should likely be src/$(PROJECT_ID)
PROJECT_ROOT := src/$(PROJECT_PREFIX)

.DEFAULT_GOAL := all

# Ensure the workspace is setup
workspace_deps:
	@ ( \
		test -d $(PROJECT_ROOT) || mkdir -p $(PWD)/$(PROJECT_ROOT); \
		test -e $(PROJECT_ROOT)/$(PROJECT_NAME) || ln -s $(PWD) $(PROJECT_ROOT)/. ; \
	);

# "Normal" static binary
%.statbin: workspace_deps
	GOPATH=$(PWD) go build -o builds/$(subst .statbin,,$@) $(PROJECT_ID)/$(subst .statbin,,$@)

clean:
	test ! -e $(PROJECT_ROOT) || rm -rvf $(PROJECT_ROOT); \
	test ! -d $(PWD)/pkg || rm -rvf $(PWD)/pkg; \
	test ! -d $(PWD)/builds || rm -rvf $(PWD)/builds

%.compliant:
	@ ( \
		pushd "$(subst .compliant,,$@)" >/dev/null || exit 1; \
		go fmt || exit 1; \
		GOPATH=$(PWD)/ golint || exit 1; \
		GOPATH=$(PWD)/ go vet || exit 1; \
	);

BINARIES = \
	yauto

COMPLIANCE = $(addsuffix .compliant,$(BINARIES))
# Build yauto as static for now
BINS = $(addsuffix .statbin,$(BINARIES))
	
# Ensure our own code is compliant..
compliant: $(COMPLIANCE)
install: $(BINS)
	test -d $(DESTDIR)/usr/bin || install -D -d -m 00755 $(DESTDIR)/usr/bin; \
	install -m 00755 builds/* $(DESTDIR)/usr/bin/.

all: $(BINS)
