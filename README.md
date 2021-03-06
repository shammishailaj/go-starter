# go-starter

`go-starter` is an opinionated [golang](https://golang.org/) backend development template by [allaboutapps](https://allaboutapps.at/).

## Overview

## Table of Contents

- [Features](#features)
- [Usage](#usage)
  - [Requirements](#requirements)
  - [Quickstart](#quickstart)
- [Contributing](#contributing)
- [Maintainers](#maintainers)
- [License](#license)

## Features

- Full local golang service development environment using [Docker Compose](https://docs.docker.com/compose/install/) and [VSCode devcontainers](https://code.visualstudio.com/docs/remote/containers) that just works with Linux, MacOS and Windows.
- Adheres to the project layout defined in [golang-standard/project-layout](https://github.com/golang-standards/project-layout)
- Provides database migration ([sql-migrate](https://github.com/rubenv/sql-migrate)) and models generation ([SQLboiler](https://github.com/volatiletech/sqlboiler)) workflows for [PostgreSQL](https://www.postgresql.org/) databases
- Integrates [IntegreSQL](https://github.com/allaboutapps/integresql) for fast, concurrent and isolated integration testing with real PostgreSQL databases
- Autoinstalls our recommended VSCode extensions for golang development
- Integrates [go-swagger](https://github.com/go-swagger/go-swagger) for compile-time autogeneration of `swagger.json` and request/response validation functions.
- Integrates [MailHog](https://github.com/mailhog/MailHog) for easy SMTP-based email testing
- Integrates [pgFormatter](https://github.com/darold/pgFormatter) and [vscode-pgFormatter](https://marketplace.visualstudio.com/items?itemName=bradymholt.pgformatter) for SQL formatting
- Comes with fully implemented `auth` package, an OAuth2 RESTful JSON API ready to be extended according to your requirements. [The latest `swagger.json` at ReDoc](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/allaboutapps/go-starter/master/api/swagger.json).
- Implements [OAuth 2.0 Bearer Tokens](https://tools.ietf.org/html/rfc6750) and password authentication using [argon2id](https://godoc.org/github.com/alexedwards/argon2id) hashes.
- Parallel jobs optimized `Makefile` and various convenience scripts, a full rebuild via `make build` only takes seconds
- Multi-staged `Dockerfile` (`development` -> `builder` -> `builder-apiserver` -> `apiserver`)

## Usage

### Requirements

Requires the following local setup for development:

- [Docker CE](https://docs.docker.com/install/) (19.03 or above)
- [Docker Compose](https://docs.docker.com/compose/install/) (1.25 or above)

The project makes use of the [VSCode devcontainers functionality](https://code.visualstudio.com/docs/remote/containers) provided by [Visual Studio Code](https://code.visualstudio.com/), thus a local installation of a Go compiler is *no longer* required when using this IDE.

### Quickstart

> GitHub: Click on **[Use this template](https://github.com/allaboutapps/go-starter/generate)** to create your own project.   
> Contributions and others: You will need to fork this repository.

After your `git clone` you may do the following:

```bash

# $local
# Easily start the docker-compose dev environment through our helper
./docker-helper.sh --up

# ---

# development@XXXXXXXXX:/app$
# You should now have a docker shell...

# If you have forked this project:
# change the go project module name and create a new README
# module allaboutapps.dev/<GIT_PROJECT>/<GIT_REPO>
make set-module-name
# e.g. example.com/my-awesome-service
mv README.md README-go-starter.md
make get-module-name > README.md

# Shortcut for make init, make build, make info and make test
make all

# Init install/cache dependencies and install tools to bin
make init

# Rebuild only after changes to files (generate, format, build, vet)
make

# Execute all tests
make test

# Migrate up the development database
sql-migrate up

# Start the local-built server
apiserver

# ---

# $local

# you may attach to the development container through multiple shells, it's always the same command
./docker-helper.sh --up

# if you ever need to halt the docker-compose env (without deleting your projects' images & volumes)
./docker-helper.sh --halt

# if you ever change something in the Dockerfile and require a rebuild of the service image only
./docker-helper.sh --rebuild

# if you ever need to wipe ALL traces (will delete your projects' images & volumes)
./docker-helper.sh --destroy

```

Regarding [Visual Studio Code](https://code.visualstudio.com/): Always develop *inside* the running `development` docker container, by attaching to this container.

Run CMD+SHIFT+P `Go: Install/Update Tools` after starting vscode to autoinstall all golang vscode dependencies, then **reload your window**.

## Contributing

Pull requests are welcome. For major changes, please [open an issue](https://github.com/allaboutapps/go-starter/issues/new) first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Maintainers

- [Michael Farkas - @farkmi](https://github.com/farkmi)
- [Nick Müller - @MorpheusXAUT](https://github.com/MorpheusXAUT)
- [Mario Ranftl - @majodev](https://github.com/majodev)

## License

[MIT](LICENSE) © 2020 aaa – all about apps GmbH | Michael Farkas | Nick Müller | Mario Ranftl and the `go-starter` project contributors
