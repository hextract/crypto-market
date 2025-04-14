DOCKER_COMPOSE := $(shell command -v docker-compose >/dev/null 2>&1 && echo docker-compose || echo docker compose)

.PHONY: pull
pull:
	git pull

.PHONY: swagger_generate
swagger_generate:
	./scripts/generate_from_swagger.sh
.PHONY: codegen
codegen: swagger_generate

.PHONY: update
update: pull
	$(DOCKER_COMPOSE) up -d --build --force-recreate -V --remove-orphans

.PHONY: down
down:
	$(DOCKER_COMPOSE) down -v --remove-orphans

.PHONY: build-action
build: down
	$(DOCKER_COMPOSE) up -d --build --force-recreate -V --remove-orphans

.PHONY: up
up:
	$(DOCKER_COMPOSE) up -d --build --force-recreate -V --remove-orphans

.PHONY: ps
ps:
	$(DOCKER_COMPOSE) ps -a


.PHONY: cert
cert:
	./scripts/generate_certificates.sh