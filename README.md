# Portto

[![Testing](https://github.com/blackhorseya/portto/actions/workflows/test.yaml/badge.svg)](https://github.com/blackhorseya/portto/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/blackhorseya/portto/branch/main/graph/badge.svg?token=OGLkIzTeqe)](https://codecov.io/gh/blackhorseya/portto)

## Architecture

使用 `Domain-Driven Design` 設計整個系統架構

### Domain

- block: 代表了 `blockchain` 上時間軸
- activity: 面向商業上使用者各種行為的定義

### Project Layout

以 [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
和 [monorepo](https://monorepo.tools/) 作為專案結構

- api: OpenAPI/Swagger 規格、JSON schema 檔案、各種協定定義檔
- cmd: 所有 service 的進入點
- configs: 組態設定的檔案範本或預設設定
- deployments: IaaS、PaaS、系統和容器編配部署的組態設定與範本
- internal: 私有應用程式和函式庫的程式碼，是你不希望其他人在其應用程式或函式庫中匯入的程式碼
    - app: 實際 application 實作的地方
    - pkg: 內部函式庫實作的地方
- pb: 存放所有 `protobuf` 的地方
- pkg: 函式庫的程式碼當然可以讓外部應用程式來使用
- scripts: 放置要執行各種建置、安裝、分析等操作的命令腳本
- test: 額外的外部測試應用程式和測試資料

### Tech Stack

- [gin](https://github.com/gin-gonic/gin)
- [wire](https://github.com/google/wire)
- [testify](https://github.com/stretchr/testify)
- [mockery](https://github.com/vektra/mockery)
- [migrate](https://github.com/golang-migrate/migrate)
