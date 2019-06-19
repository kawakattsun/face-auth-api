.PHONY: deps clean build

deps:
	dep ensure

lint:
	sh scripts/go-lint.sh

clean: 
	rm -rf ./hello-world/hello-world
	
build:
	sh scripts/build-handllers.sh