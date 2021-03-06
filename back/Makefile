GO_FILES ?= ./...
GO = GO111MODULE=on go

install:
	@echo "[+] install"
	$(GO) get $(GO_FILES)
.PHONY: install

start: install
	@echo "[+] start"
	@docker-compose up
.PHONY: start

db: 
	@echo "[+] set up db"
	@docker-compose up --renew-anon-volumes -d db
.PHONY: db

dev: db
	@echo "[+] start in development mode (docker)"
	@docker-compose up app
.PHONY: dev

watch:
	@echo "[+] start in development mode"
	@watchexec --restart --exts "go" --watch . "go run main.go"
.PHONY: watch

test:
	@echo "[+] test"
	@docker-compose up -d db
	$(GO) test -race -test.parallel 1 $(GO_FILES)
.PHONY: test

watch-test:
	@echo "[+] watch tests"
	@docker-compose up --renew-anon-volumes -d db
	@watchexec --restart --exts "go" --watch . "go test ./..."
.PHONY: watch-test

lint: vet golangci-lint revive sec
.PHONY: lint

scripts/bin/revive: scripts/go.mod
	@cd scripts; \
	$(GO) build -o ./bin/revive github.com/mgechev/revive

scripts/bin/golangci-lint: scripts/go.mod
	@cd scripts; \
	$(GO) build -o ./bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

scripts/bin/gosec: scripts/go.mod
	@cd scripts; \
	$(GO) build -o ./bin/gosec github.com/securego/gosec/cmd/gosec

revive: scripts/bin/revive
	@echo "lint via revive"
	@scripts/bin/revive \
		-formatter stylish \
		-exclude ./vendor/... \
		$(GO_FILES)
.PHONY: revive

golangci-lint: scripts/bin/golangci-lint
	@echo "lint via golangci-lint"
	@scripts/bin/golangci-lint run \
		--config ./scripts/configs/.golangci.yml \
		$(GO_FILES)
.PHONY: golangci-lint

sec: scripts/bin/gosec
	@echo "lint via gosec"
	@scripts/bin/gosec -quiet \
		-exclude=G104,G107,G108,G201,G202,G204,G301,G304,G401,G402,G501 \
		-conf=./scripts/configs/gosec.json \
		$(GO_FILES)
.PHONY: sec

vet:
	@echo "lint via go vet"
	@$(GO) vet $(GO_FILES)
.PHONY: vet