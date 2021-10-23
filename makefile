up: ## Spin up the project
	docker-compose up -d
down: ## Spin down the project
	docker-compose down
build: ## build and start project
	docker-compose up --build --force-recreate -d

.ONESHELL:
gen: ## regen graphql files
	cd golang
	go get github.com/99designs/gqlgen/cmd@v0.14.0  github.com/99designs/gqlgen/internal/imports@v0.14.0 github.com/99designs/gqlgen/internal/code@v0.14.0 github.com/99designs/gqlgen/internal/imports@v0.14.0
	go get ./...
	go run github.com/99designs/gqlgen generate
