DOCKER_SERVICE_GO ?= go
DOCKER_SERVICE_SAM ?= sam

DOCKER_COMPOSE_RUN = docker-compose run --rm

.PHONY: clean build deploy creat-bucket delete-stack describe-stacks create-collection

all:
	@echo "make all not defined."

lint:
	$(DOCKER_COMPOSE_RUN) $(DOCKER_SERVICE_GO) sh scripts/go-lint.sh

clean: 
	find build/deploy/cmd -name main -type f | xargs rm -f

build: lint clean
	$(DOCKER_COMPOSE_RUN) $(DOCKER_SERVICE_GO) sh scripts/build-handllers.sh

deploy: build
	$(DOCKER_COMPOSE_RUN) $(DOCKER_SERVICE_SAM) sh scripts/deploy.sh

create-bucket:
	$(DOCKER_COMPOSE_RUN) $(DOCKER_SERVICE_SAM) sh scripts/create-bucket.sh

delete-stack:
	$(DOCKER_COMPOSE_RUN) $(DOCKER_SERVICE_SAM) sh scripts/delete-stack.sh

describe-stacks:
	$(DOCKER_COMPOSE_RUN) $(DOCKER_SERVICE_SAM) sh scripts/describe-stacks.sh

create-collection:
	$(DOCKER_COMPOSE_RUN) $(DOCKER_SERVICE_SAM) sh scripts/create-collection.sh

generate-event:
	$(DOCKER_COMPOSE_RUN) $(DOCKER_SERVICE_SAM) sh scripts/generate-event.sh
