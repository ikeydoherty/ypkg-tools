# Ensure the workspace is setup
workspace_deps:
	@ ( \
		test -d $(PROJECT_ROOT) || mkdir -p $(PWD)/$(PROJECT_ROOT); \
		test -e $(PROJECT_ROOT)/$(PROJECT_NAME) || ln -s $(PWD) $(PROJECT_ROOT)/. ; \
	);

# "Normal" static binary
%.statbin: workspace_deps
	GOPATH=$(PWD) go build -o builds/$(subst .statbin,,$@) $(PROJECT_ID)/$(subst .statbin,,$@)

# Dynamic golang binary
%.dynbin: workspace_deps
	GOPATH=$(PWD) go build -linkshared -pkgdir $(PWD)/pkg -o builds/$(subst .dynbin,,$@) $(PROJECT_ID)/$(subst .dynbin,,$@)

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
