### -----------------------
# --- Building
### -----------------------

# go module name (as in go.mod)
# only evaluated if required by a recipe
# http://make.mad-scientist.net/deferred-simple-variable-expansion/
GO_MODULE_NAME = $(eval GO_MODULE_NAME := $$(shell \
	(mkdir -p tmp 2> /dev/null && cat tmp/.modulename 2> /dev/null) \
	|| (go run scripts/modulename/modulename.go | tee tmp/.modulename) \
))$(GO_MODULE_NAME)

# first is default task when running "make" without args
build:
	@$(MAKE) build-pre
	@$(MAKE) go-format
	@$(MAKE) go-build
	@$(MAKE) go-lint

# useful to ensure that everything gets resetuped from scratch
all:
	@$(MAKE) clean
	@$(MAKE) init
	@$(MAKE) build
	@$(MAKE) info
	@$(MAKE) test

info:
	@echo "database:"
	@cat scripts/sql/info.sql | psql -q -d "${PSQL_DBNAME}"
	@echo "handlers:"
	@go run scripts/handlers/gen_handlers.go --print-only
	@$(MAKE) info-module-name
	@go version

# these recipies may execute in parallel
build-pre: sql-generate swagger go-generate 

go-format:
	go fmt

go-build: 
	go build -o bin/apiserver ./cmd/api

go-lint:
	golangci-lint run --fast

go-generate:
	go run scripts/handlers/gen_handlers.go

# https://github.com/golang/go/issues/24573
# w/o cache - see "go help testflag"
# use https://github.com/kyoh86/richgo to color
# note that these tests should not run verbose by default (e.g. use your IDE for this)
# TODO: add test shuffling/seeding when landed in go v1.15 (https://github.com/golang/go/issues/28592)
test:
	richgo test -cover -race -count=1 ./...

### -----------------------
# --- Initializing
### -----------------------

init:
	@$(MAKE) modules
	@$(MAKE) tools
	@$(MAKE) tidy
	@go version

# cache go modules (locally into .pkg)
modules:
	go mod download

# https://marcofranssen.nl/manage-go-tools-via-go-modules/
tools:
	cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

tidy:
	go mod tidy

### -----------------------
# --- SQL
### -----------------------

sql-reset:
	@echo "DROP & CREATE database:"
	@echo "  PGHOST=${PGHOST} PGDATABASE=${PGDATABASE}" PGUSER=${PGUSER}
	@echo -n "Are you sure? [y/N] " && read ans && [ $${ans:-N} = y ]
	psql -d postgres -c 'DROP DATABASE IF EXISTS "${PGDATABASE}";'
	psql -d postgres -c 'CREATE DATABASE "${PGDATABASE}" WITH OWNER ${PGUSER} TEMPLATE "template0";'

# This step is only required to be executed when the "migrations" folder has changed!
# MIGRATION_FILES = $(find ./migrations/ -type f -iname '*.sql')
sql-generate: # ./migrations $(MIGRATION_FILES)
	@$(MAKE) sql-format
	@$(MAKE) sql-check-files
	@$(MAKE) sql-spec-reset
	@$(MAKE) sql-spec-migrate
	@$(MAKE) sql-check-structure
	sqlboiler psql

sql-format:
	@echo "make sql-format"
	@find ${PWD} -name ".*" -prune -o -type f -iname "*.sql" -print \
		| xargs -i pg_format {} -o {}

sql-check-files: sql-check-syntax sql-check-migrations-unnecessary-null

# check syntax via the real database
# https://stackoverflow.com/questions/8271606/postgresql-syntax-check-without-running-the-query
sql-check-syntax:
	@echo "make sql-check-syntax"
	@find ${PWD} -name ".*" -prune -o -type f -iname "*.sql" -print \
		| xargs -i sed '1s#^#DO $$SYNTAX_CHECK$$ BEGIN RETURN;#; $$aEND; $$SYNTAX_CHECK$$;' {} \
		| psql -d postgres --quiet -v ON_ERROR_STOP=1

sql-check-migrations-unnecessary-null:
	@echo "make sql-check-migrations-unnecessary-null"
	@(grep -R "NULL" ./migrations/ | grep --invert "DEFAULT NULL" | grep --invert "NOT NULL" | grep --invert "WITH NULL" | grep --invert "NULL, " | grep --invert ", NULL") \
		&& exit 1 || exit 0

sql-spec-reset:
	@echo "make sql-spec-reset"
	@psql --quiet -d postgres -c 'DROP DATABASE IF EXISTS "${PSQL_DBNAME}";'
	@psql --quiet -d postgres -c 'CREATE DATABASE "${PSQL_DBNAME}" WITH OWNER ${PSQL_USER} TEMPLATE "template0";'

sql-spec-migrate:
	@echo "make sql-spec-migrate"
	@sql-migrate up -env spec

sql-check-structure: 
	@$(MAKE) sql-check-structure-fk-missing-index
	@$(MAKE) sql-check-structure-default-zero-values

sql-check-structure-fk-missing-index:
	@echo "make sql-check-structure-fk-missing-index"
	@cat scripts/sql/fk_missing_index.sql | psql -qtz0 --no-align -d  "${PSQL_DBNAME}" -v ON_ERROR_STOP=1

sql-check-structure-default-zero-values:
	@echo "make sql-check-structure-default-zero-values"
	@cat scripts/sql/default_zero_values.sql | psql -qtz0 --no-align -d "${PSQL_DBNAME}" -v ON_ERROR_STOP=1

### -----------------------
# --- Swagger
### -----------------------

swagger-gen-spec: 
	@echo "make swagger-gen-spec"
	@swagger generate spec \
		-i api/swagger.yml \
		-o api/swagger.json \
		--scan-models \
		-q

swagger-models:
	@echo "make swagger-models"
	@rm -rf tmp/types 2> /dev/null
	@swagger generate model \
		--allow-template-override \
		--template-dir=internal/types/swagger \
		--spec=api/swagger.json \
		--existing-models=${GO_MODULE_NAME}/internal/types \
		--model-package=tmp/types \
		--all-definitions \
		-q
	go run scripts/gocat/gocat.go -p types tmp/types/* > internal/types/validations.go
	@rm -rf tmp/types

swagger-validate:
	@echo "make swagger-validate"
	@swagger validate api/swagger.json \
		--stop-on-error \
		-q

swagger-gen-server: swagger-validate swagger-models

swagger: 
	@$(MAKE) swagger-gen-spec
	@$(MAKE) swagger-gen-server

# accessable from outside via:
# mac: http://docker.for.mac.localhost:8080/docs
swagger-serve:
	swagger serve --no-open -p 8080 api/swagger.json

### -----------------------
# --- Helpers
### -----------------------

clean:
	rm -rf tmp

get-module-name:
	@echo "${GO_MODULE_NAME}"

info-module-name:
	@echo "current go module-name: ${GO_MODULE_NAME}"

set-module-name:
	@rm -f tmp/.modulename
	@$(MAKE) info-module-name
	@echo "Enter new go module-name:" \
		&& read ans \
		&& find . -not -path '*/\.*' -type f -exec sed -i "s|${GO_MODULE_NAME}|$${ans}|g" {} \; \
		&& echo "new go module-name: $${ans}"
	@rm -f tmp/.modulename

### -----------------------
# --- Special targets
### -----------------------

# https://www.gnu.org/software/make/manual/html_node/Special-Targets.html
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
# ignore matching file/make rule combinations in working-dir
.PHONY: test

# https://unix.stackexchange.com/questions/153763/dont-stop-makeing-if-a-command-fails-but-check-exit-status
# https://www.gnu.org/software/make/manual/html_node/One-Shell.html
# required to ensure make fails if one recipe fails (even on parallel jobs)
.ONESHELL:
SHELL = /bin/bash
.SHELLFLAGS = -ec
