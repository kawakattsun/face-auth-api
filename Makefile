.PHONY: deps clean build deploy

deps:
	@dep ensure

lint:
	@sh scripts/go-lint.sh

clean: 
	@find handlers -name main -type f | xargs rm -f

build:
	@make deps
	@make clean
	@sh scripts/build-handllers.sh

deploy:
	@sh scripts/deploy.sh
