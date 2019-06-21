.PHONY: deps clean build deploy creat-bucket delete-stack describe-stacks

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

create-bucket:
	@sh scripts/create-bucket.sh

delete-stack:
	@sh scripts/delete-stack.sh

describe-stacks:
	@sh scripts/describe-stacks.sh