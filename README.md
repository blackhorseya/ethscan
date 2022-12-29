# ethscan

[![Testing](https://github.com/blackhorseya/ethscan/actions/workflows/test.yaml/badge.svg)](https://github.com/blackhorseya/ethscan/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/blackhorseya/ethscan/branch/main/graph/badge.svg?token=OGLkIzTeqe)](https://codecov.io/gh/blackhorseya/ethscan)

## Documents

[Swagger](https://ethscan.seancheng.space/api/ethscan/docs/index.html)

## Architecture

Design the entire system architecture using `Domain-Driven Design`

### Domain

- block: Represents the timeline on the `blockchain`
- activity: Definition of various behaviors for commercial users

### Project Layout

Take [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
and [monorepo](https://monorepo.tools/) as project structure

- api: OpenAPI/Swagger specs, JSON schema files, protocol definition files.
- cmd: Main applications for this project.
- configs: Configuration file templates or default configs.
- deployments: IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).
- internal: Private application and library code.
    - app: implementation application
    - pkg: internal private library
- pb: `protobuf`
- pkg: Library code that's ok to use by external applications.
- scripts: Scripts to perform various build, install, analysis, etc operations.
- test: Additional external test apps and test data.

### Tech Stack

- [gin](https://github.com/gin-gonic/gin)
- [wire](https://github.com/google/wire)
- [testify](https://github.com/stretchr/testify)
- [mockery](https://github.com/vektra/mockery)
- [migrate](https://github.com/golang-migrate/migrate)
