modulesRepositories := `cat .modules`
protoFiles := `find modules -type f -name "*.proto"`

install-modules:
	[ ! -e modules ] || rm -dRf modules
	mkdir -p modules
	@for module in $(modulesRepositories) ; do \
		cd modules && git clone $$module && cd -; \
    done
	@for file in $(protoFiles) ; do \
    	protoc $$file --go_out=plugins=grpc:.; \
    done
run:
	go run .