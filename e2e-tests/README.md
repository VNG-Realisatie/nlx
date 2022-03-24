E2E tests
---

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

The following tags can be used:

- `@ignore`: this Feature / Scenario won't be executed
- `@execution:serial`: this Feature / Scenario will run serially
- `@unauthenticated`: logout for every organization before executing the Feature / Scenario. Should only be used for serial tests.

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
E2E_GEMEENTE_STIJNS_OUTWAY_1_NAME: "gemeente-stijns-nlx-outway"
E2E_GEMEENTE_STIJNS_OUTWAY_1_ADDRESS: "http://127.0.0.1:7917"
E2E_GEMEENTE_STIJNS_OUTWAY_2_NAME: "gemeente-stijns-nlx-outway-2"
E2E_GEMEENTE_STIJNS_OUTWAY_2_ADDRESS: "http://127.0.0.1:7947"

E2E_RVRD_MANAGEMENT_MANAGEMENT_URL: "http://management.organization-b.nlx.local:3021
E2E_RVRD_MANAGEMENT_BASIC_AUTH: "true"
E2E_RVRD_MANAGEMENT_USERNAME: "admin@nlx.local"
E2E_RVRD_MANAGEMENT_PASSWORD: "development"
E2E_RVRD_DEFAULT_INWAY_NAME: "Inway-01"

E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_MANAGEMENT_URL: "http://management.organization-c.nlx.local:3031"
E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_BASIC_AUTH: "true"
E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_USERNAME: "admin@nlx.local"
E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_PASSWORD: "development"
E2E_VERGUNNINGSOFTWARE_BV_DEFAULT_INWAY_NAME: ""
E2E_VERGUNNINGSOFTWARE_BV_OUTWAY_1_NAME: "vergunningsoftware-bv-nlx-outway"
E2E_VERGUNNINGSOFTWARE_BV_OUTWAY_1_ADDRESS: "http://127.0.0.1:7937"
```

## Useful links

- https://cucumber.io/docs/cucumber/api/#tag-expressions
