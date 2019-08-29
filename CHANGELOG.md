# [0.75.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.74.1...v0.75.0) (2019-08-29)


### Bug Fixes

* **ca-certportal:** include common folder while building as we are depending on it ([9382679](https://gitlab.com/commonground/nlx/nlx/commit/9382679))
* **common:** ignore logger sync errors for stdout and stderr ([8aeb22a](https://gitlab.com/commonground/nlx/nlx/commit/8aeb22a))
* **docs:** add newline to make sure list is recognized ([d02877b](https://gitlab.com/commonground/nlx/nlx/commit/d02877b))
* **inway:** content-type in api spec response ([6c7a106](https://gitlab.com/commonground/nlx/nlx/commit/6c7a106))


### Features

* **config-api:** add api for configuring inway and services ([01270dc](https://gitlab.com/commonground/nlx/nlx/commit/01270dc))
* **directory-ui:** let Redoc fetch the API spec ([ad3779b](https://gitlab.com/commonground/nlx/nlx/commit/ad3779b))
* **helm:** add HTTP Strict Transport Security (HSTS) headers to Ingress definitions ([aad13ff](https://gitlab.com/commonground/nlx/nlx/commit/aad13ff))
* **inway:** increase max idle connections to logdb ([b4e24c7](https://gitlab.com/commonground/nlx/nlx/commit/b4e24c7))
* **inway:** set MaxIdleConnsPerHost equal to MaxIdleConns ([7ebce3b](https://gitlab.com/commonground/nlx/nlx/commit/7ebce3b))
* **outway:** create HttpService when needed. log (un)healty information. ([2962eab](https://gitlab.com/commonground/nlx/nlx/commit/2962eab))
* **outway:** grpc health changes backwards compatible with old version ([7543bc0](https://gitlab.com/commonground/nlx/nlx/commit/7543bc0))
* **outway:** grpc health changes backwards compatible with old version ([e5bc050](https://gitlab.com/commonground/nlx/nlx/commit/e5bc050))
* **outway:** increase max idle connections to logdb ([31406aa](https://gitlab.com/commonground/nlx/nlx/commit/31406aa))
* **outway:** now recieving health status from directory ([e3974bc](https://gitlab.com/commonground/nlx/nlx/commit/e3974bc))
* **outway:** now recieving health status from directory ([c1505ff](https://gitlab.com/commonground/nlx/nlx/commit/c1505ff))

## [0.74.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.74.0...v0.74.1) (2019-07-24)


### Bug Fixes

* **directory-ui:** change default port of outway to 80 to be consistent with docs ([1e21bb3](https://gitlab.com/commonground/nlx/nlx/commit/1e21bb3))
* **docs:** do not detach docker run commands ([f6af586](https://gitlab.com/commonground/nlx/nlx/commit/f6af586))

# [0.74.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.73.0...v0.74.0) (2019-07-12)


### Bug Fixes

* **deps:** update dependency ajv-keywords to v3.4.1 ([e59049a](https://gitlab.com/commonground/nlx/nlx/commit/e59049a))
* **deps:** update dependency redux to v4.0.2 ([8b6d3ad](https://gitlab.com/commonground/nlx/nlx/commit/8b6d3ad))
* **deps:** update dependency redux to v4.0.4 ([d842dec](https://gitlab.com/commonground/nlx/nlx/commit/d842dec))
* **directory-registration-api:** fix inway config not being updated in the directory database ([367a851](https://gitlab.com/commonground/nlx/nlx/commit/367a851))


### Features

* **design:** Add sequence diagrams for organization management and keyless and improved organization management config. ([37c7d24](https://gitlab.com/commonground/nlx/nlx/commit/37c7d24))

# [0.71.0-one.zero](https://gitlab.com/commonground/nlx/nlx/compare/v0.71.0...v0.71.0-one.zero) (2019-07-01)

### Bug Fixes

* **deps:** update all non-major dependencies ([447140a](https://gitlab.com/commonground/nlx/nlx/commit/447140a))

### Code Refactoring

* **txlog-ui, txlog-api:** eliminate txlog-ui and txlog-api ([6f9e407](https://gitlab.com/commonground/nlx/nlx/commit/6f9e407))
* **txlog-ui, txlog-api:** txlog-ui and txlog-api are not available any more

# [0.71.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.70.2...v0.71.0) (2019-06-27)


### Bug Fixes

* **directory-ui:** loosen CSP to make sure images work on API documentation page ([e88b3aa](https://gitlab.com/commonground/nlx/nlx/commit/e88b3aa))
* **directory-ui:** make sure OpenAPI3 specification is displayed correctly ([4d3b5de](https://gitlab.com/commonground/nlx/nlx/commit/4d3b5de))
* **docs:** use the correct portnumber in the example link ([61c7d3b](https://gitlab.com/commonground/nlx/nlx/commit/61c7d3b))


### Features

* add api-specification-document-url to example service basisregistratie ([bc0441b](https://gitlab.com/commonground/nlx/nlx/commit/bc0441b))

## [0.70.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.70.1...v0.70.2) (2019-06-27)


### Bug Fixes

* exclude txlog-ui tests from the schedules [#602](https://gitlab.com/commonground/nlx/nlx/issues/602) ([a34661b](https://gitlab.com/commonground/nlx/nlx/commit/a34661b))
* update images-to-scan.txt location ([f4ae4de](https://gitlab.com/commonground/nlx/nlx/commit/f4ae4de))
* update variable substitution for xargs [#602](https://gitlab.com/commonground/nlx/nlx/issues/602) ([81f7fcb](https://gitlab.com/commonground/nlx/nlx/commit/81f7fcb))
* **directory-ui:** adjust CSP to make sure Swagger UI is working as well ([5641a93](https://gitlab.com/commonground/nlx/nlx/commit/5641a93))
* **docs:** add  .js.map files for dependencies ([ce81b18](https://gitlab.com/commonground/nlx/nlx/commit/ce81b18))

## [0.70.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.70.0...v0.70.1) (2019-06-25)


### Bug Fixes

* **docs:** links to folders without slashes now work as expected ([f5a2659](https://gitlab.com/commonground/nlx/nlx/commit/f5a2659))
* add security headers also on error pages ([5b70507](https://gitlab.com/commonground/nlx/nlx/commit/5b70507))

# [0.70.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.69.0...v0.70.0) (2019-06-24)


### Features

* **directory-ui:** show tooltip on API URL and documentation button + show notifier when URL is copied ([af2c7af](https://gitlab.com/commonground/nlx/nlx/commit/af2c7af))

# [0.69.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.68.3...v0.69.0) (2019-06-20)


### Features

* add cache headers ([a44aa58](https://gitlab.com/commonground/nlx/nlx/commit/a44aa58))

## [0.68.3](https://gitlab.com/commonground/nlx/nlx/compare/v0.68.2...v0.68.3) (2019-06-19)

## [0.68.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.68.1...v0.68.2) (2019-06-19)


### Bug Fixes

* loosen CSP to make sure apps work ([3e1666c](https://gitlab.com/commonground/nlx/nlx/commit/3e1666c))

## [0.68.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.68.0...v0.68.1) (2019-06-19)


### Bug Fixes

* fix releaserc config ([b1efafc](https://gitlab.com/commonground/nlx/nlx/commit/b1efafc))
* rework releaserc config to only specify preset overrides ([749caa5](https://gitlab.com/commonground/nlx/nlx/commit/749caa5))

# [0.68.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.67.0...v0.68.0) (2019-06-18)


### Bug Fixes

* specify a form action for all forms to prevent HTTP parameter override attacks [#588](https://gitlab.com/commonground/nlx/nlx/issues/588) ([7403437](https://gitlab.com/commonground/nlx/nlx/commit/7403437))
* **txlog-ui:** install the correct @commonground/design-system version ([313c2ea](https://gitlab.com/commonground/nlx/nlx/commit/313c2ea))
* **txlog-ui:** remove deprecated spacing property ([afbe6e6](https://gitlab.com/commonground/nlx/nlx/commit/afbe6e6))


### Features

* **directory-ui:** replace local VersionLogger with shared component ([860c10d](https://gitlab.com/commonground/nlx/nlx/commit/860c10d))
* **directory-ui:** show the tag field from version.json in the console log ([2046380](https://gitlab.com/commonground/nlx/nlx/commit/2046380))
* **insight-ui:** add VersionLogger component to log the version on the console ([251f872](https://gitlab.com/commonground/nlx/nlx/commit/251f872))
* **txlog-ui:** add VersionLogger component to log the version on the console ([f1e3c73](https://gitlab.com/commonground/nlx/nlx/commit/f1e3c73))

# [0.67.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.66.2...v0.67.0) (2019-06-17)


### Bug Fixes

* **deps:** update dependency react-app-polyfill to v1 ([996031b](https://gitlab.com/commonground/nlx/nlx/commit/996031b))
* **deps:** update dependency react-redux to v7.1.0 ([a095392](https://gitlab.com/commonground/nlx/nlx/commit/a095392))
* **deps:** update dependency redux-saga to v1.0.3 ([fb00941](https://gitlab.com/commonground/nlx/nlx/commit/fb00941))
* **deps:** update mui monorepo ([43a08b1](https://gitlab.com/commonground/nlx/nlx/commit/43a08b1))


### Features

* **ui:** add hover/active styling to organisation, log row and close button ([e286611](https://gitlab.com/commonground/nlx/nlx/commit/e286611))

## [0.66.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.66.1...v0.66.2) (2019-06-13)


### Bug Fixes

* **common:** update default logType for the go-flags [#547](https://gitlab.com/commonground/nlx/nlx/issues/547) ([dd6ae3a](https://gitlab.com/commonground/nlx/nlx/commit/dd6ae3a))
* **deps:** update dependency axios to v0.19.0 ([3fca974](https://gitlab.com/commonground/nlx/nlx/commit/3fca974))
* **deps:** update dependency copy-text-to-clipboard to v2 ([1b98930](https://gitlab.com/commonground/nlx/nlx/commit/1b98930))
* **deps:** update dependency mobx to v4.10.0 ([8085d4d](https://gitlab.com/commonground/nlx/nlx/commit/8085d4d))
* **deps:** update dependency mobx to v5 ([c60b640](https://gitlab.com/commonground/nlx/nlx/commit/c60b640))
* **deps:** update dependency prop-types to v15.7.2 ([52007aa](https://gitlab.com/commonground/nlx/nlx/commit/52007aa))
* **deps:** update dependency react-scripts to v2.1.8 ([0753ea1](https://gitlab.com/commonground/nlx/nlx/commit/0753ea1))
* **deps:** update dependency styled-components to v4.3.1 ([97dd734](https://gitlab.com/commonground/nlx/nlx/commit/97dd734))
* **README.md:** NLX Stelsel ([e3e897f](https://gitlab.com/commonground/nlx/nlx/commit/e3e897f))

## [0.66.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.66.0...v0.66.1) (2019-06-04)


### Bug Fixes

* **ca-certportal:** Fix serving of public assets ([ba7548c](https://gitlab.com/commonground/nlx/nlx/commit/ba7548c))

## [0.61.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.61.1...v0.61.2) (2019-06-03)


### Bug Fixes

* **deps:** update dependency styled-components to v4.2.1 ([4bc9ef0](https://gitlab.com/commonground/nlx/nlx/commit/4bc9ef0))

## [0.61.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.61.0...v0.61.1) (2019-05-28)


### Bug Fixes

* **deps:** update dependency @material-ui/core to v3.9.3 ([f69ec2f](https://gitlab.com/commonground/nlx/nlx/commit/f69ec2f))

# [0.58.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.57.1...v0.58.0) (2019-05-28)


### Bug Fixes

* **insight-ui:** make sure the runtime environment variables are passed ([a1886c7](https://gitlab.com/commonground/nlx/nlx/commit/a1886c7))


### Features

* **gitlab-ci:** Select kube context before deploy ([767da85](https://gitlab.com/commonground/nlx/nlx/commit/767da85))

## [0.57.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.57.0...v0.57.1) (2019-05-28)


### Bug Fixes

* **.releaserc:** removed label to skip the CI for release commits ([5e3ba0c](https://gitlab.com/commonground/nlx/nlx/commit/5e3ba0c))

# [0.57.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.56.0...v0.57.0) (2019-05-28)


### Bug Fixes

* **insight-ui:** convert expected timestamp to UTC to prevent timezone ([e0a22ca](https://gitlab.com/commonground/nlx/nlx/commit/e0a22ca)), closes [#506](https://gitlab.com/commonground/nlx/nlx/issues/506)
* **insight-ui:** extracted proof-fetching for the logs page ([e0d51d3](https://gitlab.com/commonground/nlx/nlx/commit/e0d51d3))
* **insight-ui:** make sure the _env variable is set ([86acd24](https://gitlab.com/commonground/nlx/nlx/commit/86acd24))
* **insight-ui:** make sure the environment variables can be set at ([e5bff21](https://gitlab.com/commonground/nlx/nlx/commit/e5bff21))


### Features

* **insight-ui:** add production value as default for REACT_APP_NAVBAR_DIRECTORY_URL ([7972ecc](https://gitlab.com/commonground/nlx/nlx/commit/7972ecc))
* **insight-ui:** added pagination for LogsPageContainer ([18b421b](https://gitlab.com/commonground/nlx/nlx/commit/18b421b))
* **insight-ui:** increased width of logdetail panel ([2572402](https://gitlab.com/commonground/nlx/nlx/commit/2572402))
* **insight-ui:** restart login flow when viewing a different ([686f3f8](https://gitlab.com/commonground/nlx/nlx/commit/686f3f8))
* **insight-ui:** show detail pane when clicking a log [#506](https://gitlab.com/commonground/nlx/nlx/issues/506) ([56b6102](https://gitlab.com/commonground/nlx/nlx/commit/56b6102))
* **insight-ui:** styled Pagination for LogsPage ([04e80ab](https://gitlab.com/commonground/nlx/nlx/commit/04e80ab))

# [0.56.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.55.1...v0.56.0) (2019-05-21)


### Features

* **ui:** improve header styling ([cc5741f](https://gitlab.com/commonground/nlx/nlx/commit/cc5741f))
* **ui:** remove redoc styling overrides ([f5233e5](https://gitlab.com/commonground/nlx/nlx/commit/f5233e5))
