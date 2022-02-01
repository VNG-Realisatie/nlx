E2E tests
---

## General

The following list contains the environment variables
being used along with their default value:

```
E2E_CONFIG_FILE=default
E2E_PARALLEL_COUNT=1
E2E_BUILD_NAME=local

E2E_SELENIUM_URL=http://localhost:4444/wd/hub
E2E_SELENIUM_USERNAME=<empty>
E2E_SELENIUM_ACCESS_KEY=<empty>

E2E_GEMEENTE_STIJNS_MANAGEMENT_MANAGEMENT_URL: "http://management.organization-a.nlx.local:3011"
E2E_GEMEENTE_STIJNS_MANAGEMENT_BASIC_AUTH: "false"
E2E_GEMEENTE_STIJNS_MANAGEMENT_USERNAME: "admin@nlx.local"
E2E_GEMEENTE_STIJNS_MANAGEMENT_PASSWORD: "development"
E2E_GEMEENTE_STIJNS_DEFAULT_INWAY_NAME: "Inway-01"
E2E_GEMEENTE_STIJNS_DEFAULT_OUTWAY_NAME: "outway-org-a"
E2E_GEMEENTE_STIJNS_DEFAULT_OUTWAY_ADDRESS: "http://127.0.0.1:7917"

E2E_RVRD_MANAGEMENT_MANAGEMENT_URL: "http://management.organization-b.nlx.local:3021
E2E_RVRD_MANAGEMENT_BASIC_AUTH: "true"
E2E_RVRD_MANAGEMENT_USERNAME: "admin@nlx.local"
E2E_RVRD_MANAGEMENT_PASSWORD: "development"
E2E_RVRD_DEFAULT_INWAY_NAME: "Inway-01"
E2E_RVRD_DEFAULT_OUTWAY_NAME: ""
E2E_RVRD_DEFAULT_OUTWAY_ADDRESS: ""

E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_MANAGEMENT_URL: "TODO"
E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_BASIC_AUTH: "true"
E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_USERNAME: "TODO"
E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_PASSWORD: "TODO"
E2E_VERGUNNINGSOFTWARE_BV_DEFAULT_INWAY_NAME: "Inway-01"
E2E_VERGUNNINGSOFTWARE_BV_DEFAULT_OUTWAY_NAME: ""
E2E_VERGUNNINGSOFTWARE_BV_DEFAULT_OUTWAY_ADDRESS: ""
```

## Local setup

### 1. Start development setup from the root of the project

```shell
sh ../scripts/start-development.sh
(cd ../management-ui && npm start)
(cd ../management-ui && npm run start:b)
(cd ../management-ui && npm run start:c)
```

### 2. Start Selenium Hub

**MacOS**

Make sure [Google Chrome](https://www.google.com/chrome/) is installed.

```shell
brew install selenium-server
brew install chromedriver
```

Start Selenium by running:

```shell
selenium-server standalone --port 4444
```

**Linux**

```shell
docker-compose up
```

### 3. Run the E2E tests

```shell
npm install
npm test
```
