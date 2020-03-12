# [0.85.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.84.0...v0.85.0) (2020-03-12)


### Bug Fixes

* 'potential file inclusion via variable' ([d7d7e9f](https://gitlab.com/commonground/nlx/nlx/commit/d7d7e9fbfe16953e572329481da3881b141ad59d))
* **deps:** update dependency @commonground/design-system to v8 ([7020c06](https://gitlab.com/commonground/nlx/nlx/commit/7020c06b8dab7e5605c531a6ea15e76516c7ee65))
* **directory-inspection-api:** cast stats amount as integer in the REST api ([8b796a2](https://gitlab.com/commonground/nlx/nlx/commit/8b796a2861d44ac784eebe19a315a70eb0bb79c1))
* **directory-ui:** fix linting [#872](https://gitlab.com/commonground/nlx/nlx/issues/872) ([64730b9](https://gitlab.com/commonground/nlx/nlx/commit/64730b9f8cfd0c60b310177297721bcd448922ef))
* **directory-ui-e2e-tests:** adjust Chrome installation to resolve issue with gnupg ([a4145a0](https://gitlab.com/commonground/nlx/nlx/commit/a4145a0cb803994c582e77b9fd9d809bb18b50e1))
* **docs:** fix broken links according to the link checker [#314](https://gitlab.com/commonground/nlx/nlx/issues/314) ([80948e9](https://gitlab.com/commonground/nlx/nlx/commit/80948e97fab5639913be95ff2277c4641a36971d))
* **management-ui:** fix linting [#872](https://gitlab.com/commonground/nlx/nlx/issues/872) ([0605bf5](https://gitlab.com/commonground/nlx/nlx/commit/0605bf55d2fbba947be22e65b240d17cb9e26c4e))
* **outway:** return error when sonyflake is nil ([a8aaa10](https://gitlab.com/commonground/nlx/nlx/commit/a8aaa1054bd012f6e565ea1b4c5824c88d6591df))


### Features

* **docs:** document to use the PKIoverheid root CA G1 cert ([3cd3059](https://gitlab.com/commonground/nlx/nlx/commit/3cd30591ea1fca79d919ef11101e042b66149889))
* **docs:** remove direct link to Docker download for Windows [#314](https://gitlab.com/commonground/nlx/nlx/issues/314) ([f8c215b](https://gitlab.com/commonground/nlx/nlx/commit/f8c215b3c21cf8d1e48a479df86b3a0e5cc11570))
* **gitlab-ci:** setup link checker for the docs [#314](https://gitlab.com/commonground/nlx/nlx/issues/314) ([03f7e76](https://gitlab.com/commonground/nlx/nlx/commit/03f7e76413a0d15c55e6c3788424c0c4fee5b862))
* **outway:** replace sonyflake with random id ([69755df](https://gitlab.com/commonground/nlx/nlx/commit/69755dfa6265ad97319713f228c76cb242327444))
* add support for certs singed by an intermediate CA ([dafada9](https://gitlab.com/commonground/nlx/nlx/commit/dafada9247d3f6b244418f4513db9a0337bd3cef))

# [0.84.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.83.1...v0.84.0) (2020-03-03)


### Bug Fixes

* **common:** print the issuer of certificate instead of subject ([c789750](https://gitlab.com/commonground/nlx/nlx/commit/c789750533ce5e7c726cfac3663faa4c00f1c90f))
* make sure the nlx version is sent in all calls to the directory ([d9fc13c](https://gitlab.com/commonground/nlx/nlx/commit/d9fc13c640fc522801b808de57969975314e36a1))
* **deps:** update dependency styled-components to v5 ([8676100](https://gitlab.com/commonground/nlx/nlx/commit/86761001586872a5e7c1573f9eb7c4069227c00d))
* **directory-db:** add missing ; ([c82d46a](https://gitlab.com/commonground/nlx/nlx/commit/c82d46a17f11793ae8506d0619cf095f65a08b01))
* **directory-inspection-api:** fix stats endpoint. a version is now required for a inway ([20305ec](https://gitlab.com/commonground/nlx/nlx/commit/20305ecbbec4c575b94582eea3c23aa50ae1997c))
* **docs:** correct instructions for connecting to the production environment ([01b8681](https://gitlab.com/commonground/nlx/nlx/commit/01b868140189dad03841bf2c6e18221534588d97))
* **inway:** do not include non-existing services in service-config.toml ([6444b84](https://gitlab.com/commonground/nlx/nlx/commit/6444b848e236d58cbd02631225f2b2ab19e237cc))
* **management-api:** allow users with the admin role to DELETE ([5bc5a0c](https://gitlab.com/commonground/nlx/nlx/commit/5bc5a0c8b470ffadc9e7258983823865d8abe64e))


### Features

* **docs:** enable custom 404 page for Nginx [#810](https://gitlab.com/commonground/nlx/nlx/issues/810) ([581cce8](https://gitlab.com/commonground/nlx/nlx/commit/581cce8870979f485b2aff4a67dcfc3484132f5c))

## [0.83.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.83.0...v0.83.1) (2020-02-19)


### Bug Fixes

* **directory-registration-api:** allow spaces and dots in organisation name ([47fe829](https://gitlab.com/commonground/nlx/nlx/commit/47fe8298206983efd8bf7e2a9fac20e48193429e))

# [0.83.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.82.2...v0.83.0) (2020-02-12)


### Bug Fixes

* **deps:** do not pin versions of Alpine dependencies ([b0735e4](https://gitlab.com/commonground/nlx/nlx/commit/b0735e43ef87455206848e6f0499e700ea760477))
* **docs:** fix formatting of headings in service-configuration.md ([03a8227](https://gitlab.com/commonground/nlx/nlx/commit/03a8227583d505da0443f8a0879952fdc9518120))


### Features

* **docs:** increase key size to 4096 as 3072 is not supported by all CA's ([3524328](https://gitlab.com/commonground/nlx/nlx/commit/3524328bc610e1877f0d35602fb96ac0afc4542d))
* **management-api:** add login and sessions ([9d26a79](https://gitlab.com/commonground/nlx/nlx/commit/9d26a797f2609c68acf40875d6f6dd9ccae925af))
* **management-ui:** basic logout ([00bba1d](https://gitlab.com/commonground/nlx/nlx/commit/00bba1ddf480fea436b61892bcd14c9a92899222))
* **management-ui:** show login page when not yet authenticated ([9482eec](https://gitlab.com/commonground/nlx/nlx/commit/9482eec4e8ec5107a9ee182898c353b787d002cf))

## [0.82.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.82.1...v0.82.2) (2019-12-20)


### Bug Fixes

* **docs:** link to PKI-O Private Services G1 root cert ([79c0bb7](https://gitlab.com/commonground/nlx/nlx/commit/79c0bb701088ff6c6890ea4a3ea63ae550e4c39e))

## [0.82.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.82.0...v0.82.1) (2019-12-20)

# [0.82.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.81.0...v0.82.0) (2019-12-12)


### Bug Fixes

* expired test certificates ([f5504fd](https://gitlab.com/commonground/nlx/nlx/commit/f5504fdef1ac9c627e80f63c1b93d8f735bcd6a3))


### Features

* **gitlab:** add docker-compose build to renovate testing ([533c9b5](https://gitlab.com/commonground/nlx/nlx/commit/533c9b54de3514692193ca7f0541e144fb35f86d))
* **gitlab:** deploy to acc and demo using the same snippet, including the demo organisations ([8da8ac8](https://gitlab.com/commonground/nlx/nlx/commit/8da8ac8d41c66c15b18118b2070b686482c5671e))

# [0.81.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.80.0...v0.81.0) (2019-12-09)


### Bug Fixes

* **deps:** update all non-major dependencies ([83e3137](https://gitlab.com/commonground/nlx/nlx/commit/83e313718f801599bfc6712dbf78374eaf52acfb))
* **deps:** update all non-major dependencies ([0ca7d69](https://gitlab.com/commonground/nlx/nlx/commit/0ca7d6975544fe623c3f37b254e9b8dbe25e9e74))
* **deps:** update all non-major dependencies ([fcfe2e7](https://gitlab.com/commonground/nlx/nlx/commit/fcfe2e7e255dfcf37a0a66d5eef1cd0da33304a8))
* **docs:** pin docusaurus version to 2.0.0-alpha.36 ([abb8386](https://gitlab.com/commonground/nlx/nlx/commit/abb8386b1b814229610c9268498b9485df5791a2))
* **docs:** sync package-lock.json ([2a52750](https://gitlab.com/commonground/nlx/nlx/commit/2a52750e9db1004c2f99d365ae71c84c8ca4e09f))


### Features

* **directory:** moved inway and outway version to metadata ([c4c8c84](https://gitlab.com/commonground/nlx/nlx/commit/c4c8c841fda48e283566774a37b579fc8049bddc))
* **directory-db:** update the model (+2 squashed commits) ([58cd7c5](https://gitlab.com/commonground/nlx/nlx/commit/58cd7c5e83c61782d19ba456ae159a160132def9))
* **directory-inspection-api:** added stats endpoint ([ab20dc6](https://gitlab.com/commonground/nlx/nlx/commit/ab20dc66362684166db35e4e0d0b559ec2a37b8b))
* **docs:** completed config of docusaurus v2 ([2dac346](https://gitlab.com/commonground/nlx/nlx/commit/2dac34608d74bd6db836420d1d13ab1e0418a8a7))
* **docs:** completed config of docusaurus v2 ([563ecf6](https://gitlab.com/commonground/nlx/nlx/commit/563ecf6cfb76da578cf298bb0ea3f20a9c63a94b))
* **docs:** enable syntax highlighting for bash, json, powershell and toml ([8c2d55b](https://gitlab.com/commonground/nlx/nlx/commit/8c2d55b114349283a775595896f99ae768ab2ecb))
* **docs:** migrate docusaurus from v1 to v2 ([19e2f38](https://gitlab.com/commonground/nlx/nlx/commit/19e2f3887327d64f4def98b8c49c344e6bd0e7a6))
* **docs:** migrate multi language code blocks to docusaurus v2 ([da7d5aa](https://gitlab.com/commonground/nlx/nlx/commit/da7d5aa1a0c4cb5784309dd138899659a49a3141))
* **docs:** translated the docs landing to dutch ([251850f](https://gitlab.com/commonground/nlx/nlx/commit/251850f13d64c30edf2472be619d6bb46eea36a7))
* **inway:** send version to directory-resgistration-api ([2f95f62](https://gitlab.com/commonground/nlx/nlx/commit/2f95f62c744bc8a047aee46360b70cdac1221c93))
* **outway:** send and record outway version with a timestamp ([c7adf64](https://gitlab.com/commonground/nlx/nlx/commit/c7adf6429fa1536710a3d01c430e70d8b1312e5b))
* add index to docusaurus 2 ([2f937d3](https://gitlab.com/commonground/nlx/nlx/commit/2f937d3266e198a1503f1f35f8e2c15c631f3b61))
* directly use helm instead of skaffold ([508f47b](https://gitlab.com/commonground/nlx/nlx/commit/508f47b45b342f717589e4a2a2da1c09aa45cbd9))

# [0.80.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.79.4...v0.80.0) (2019-11-22)


### Bug Fixes

* **directory-inspection-api:** delete availabilities after offline ttl ([6b63eea](https://gitlab.com/commonground/nlx/nlx/commit/6b63eeabe4755562269b6b783b6af7d6da70f5db))
* **directory-inspection-api:** list only services that have an availability ([9ece85f](https://gitlab.com/commonground/nlx/nlx/commit/9ece85f43dda4c5b9aa44b7c61d97b5ae9ffa3cd))
* **directory-monitor:** cleanup ttl to seconds and 24 hours ([2adb7bf](https://gitlab.com/commonground/nlx/nlx/commit/2adb7bf343b5fc658ee7b2f13ff3afc7e4c5def0))
* **directory-monitor:** don't panic when connection to database is lost ([2048bdd](https://gitlab.com/commonground/nlx/nlx/commit/2048bdd944ee10ff3ca017ba7a159403388dacb4))


### Features

* **directory-inspection-api:** return inway state on list services ([0f79dce](https://gitlab.com/commonground/nlx/nlx/commit/0f79dce6f90a5b29b481a17c53a485296715c64e))
* **directory-ui:** display up, down or degraded state of services ([b99f9e4](https://gitlab.com/commonground/nlx/nlx/commit/b99f9e43052d8962f56726ee94de0d1e9a9707f8))

## [0.79.4](https://gitlab.com/commonground/nlx/nlx/compare/v0.79.3...v0.79.4) (2019-11-15)


### Bug Fixes

* **docs:** add redirect from /support to /support/contact/ to fix broken incoming links ([b650eed](https://gitlab.com/commonground/nlx/nlx/commit/b650eed5da5e8147901e6f5c174a14afc01b0859))
* added specific link for common error messages ([6b00886](https://gitlab.com/commonground/nlx/nlx/commit/6b00886fb3823361454ebbe856c0e93e2754f7ae))

## [0.79.3](https://gitlab.com/commonground/nlx/nlx/compare/v0.79.2...v0.79.3) (2019-11-13)


### Bug Fixes

* **deps:** add etcd and redoc to ignoreDeps as the automated upgrades fail ([53ba0d3](https://gitlab.com/commonground/nlx/nlx/commit/53ba0d34e1a2fdca87983a693c3d0231edbc73dc))
* **deps:** update all non-major dependencies ([ca3019f](https://gitlab.com/commonground/nlx/nlx/commit/ca3019fa05de8b6e2055d5ded79f5a60ae5139b7))
* **inway:** change default value of authorization-mode from none to whitelist, as this is a safer default ([2e39aef](https://gitlab.com/commonground/nlx/nlx/commit/2e39aef914d8e125b3e0478785506844556b5a9d))
* **inway:** correctly set service permissions when using config-api, the inway blocked all requests before ([695a894](https://gitlab.com/commonground/nlx/nlx/commit/695a8941b19484c4d48c16c8405ac0b4da29a45f))

## [0.79.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.79.1...v0.79.2) (2019-11-06)

## [0.79.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.79.0...v0.79.1) (2019-11-06)


### Bug Fixes

* Clean user provided path ([7024cb4](https://gitlab.com/commonground/nlx/nlx/commit/7024cb48ef4331cd827379f3bdc84627470de72a))
* **deps:** switch to github.com/coreos/etcd for now due to version issues ([8503a62](https://gitlab.com/commonground/nlx/nlx/commit/8503a62054eab272e514182d2bcbc1eb82b7693f))
* **deps:** update all non-major dependencies ([0dab644](https://gitlab.com/commonground/nlx/nlx/commit/0dab6444515fd9f20782e348ffc360abf01a8566))
* **deps:** update dependency formik to v2 ([d982695](https://gitlab.com/commonground/nlx/nlx/commit/d98269512070ae4fc95d41f9ffd17957844eb162))
* **deps:** update redoc to 2.0.0-rc.15 ([481958e](https://gitlab.com/commonground/nlx/nlx/commit/481958e027b27631470be28363618b21b026ae48))
* **deps:** use only one assert package ([a9391a9](https://gitlab.com/commonground/nlx/nlx/commit/a9391a9f6341d49d9489583625056da26572a839))
* **directory-inspection-api:** Ignore gosec G402 for local connection ([037c515](https://gitlab.com/commonground/nlx/nlx/commit/037c5158575cd8587fa5ab3b8502ec291e4c54ab))
* **inway:** Use SHA264 when no inway name is provided ([0aec4a6](https://gitlab.com/commonground/nlx/nlx/commit/0aec4a679911c1a257b3e6971b31b8b1f763bc2e))
* E2E test acceptance url ([2f1bb07](https://gitlab.com/commonground/nlx/nlx/commit/2f1bb07ec3eee511048ee4ee0feaf47ed7487257))
* More strict security headers ([2ca5d81](https://gitlab.com/commonground/nlx/nlx/commit/2ca5d81ebf9f7b389b14a713de0c7ef028b2f4da))

# [0.79.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.78.0...v0.79.0) (2019-10-29)


### Bug Fixes

* **docs:** CSP settings [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([c1cf973](https://gitlab.com/commonground/nlx/nlx/commit/c1cf973))
* **docs:** fix further reading links [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([9e9293e](https://gitlab.com/commonground/nlx/nlx/commit/9e9293e))
* **docs:** take skaffold build context into account [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([78fb52a](https://gitlab.com/commonground/nlx/nlx/commit/78fb52a))
* **docs:** take skaffold build context into account [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([3a0fbb8](https://gitlab.com/commonground/nlx/nlx/commit/3a0fbb8))
* **docs:** update links [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([c71fdb6](https://gitlab.com/commonground/nlx/nlx/commit/c71fdb6))
* **docs:** update links [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([c4ebf56](https://gitlab.com/commonground/nlx/nlx/commit/c4ebf56))
* **docs:** update paths + change node to not run on alpine because of Docusaurus [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([d10c438](https://gitlab.com/commonground/nlx/nlx/commit/d10c438))


### Features

* **docs:** re-add Dockerfile [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([003caf7](https://gitlab.com/commonground/nlx/nlx/commit/003caf7))
* replace Hugo docs website with basic Docusaurus [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([0a1fd08](https://gitlab.com/commonground/nlx/nlx/commit/0a1fd08))
* **docs:** add NLX docs [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([df4e4b0](https://gitlab.com/commonground/nlx/nlx/commit/df4e4b0))
* **docs:** enable copy code button [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([83b4a4f](https://gitlab.com/commonground/nlx/nlx/commit/83b4a4f))
* **docs:** enable edit button [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([b6cb10c](https://gitlab.com/commonground/nlx/nlx/commit/b6cb10c))
* **docs:** enable search using Algolia [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([4748ac2](https://gitlab.com/commonground/nlx/nlx/commit/4748ac2))
* **docs:** fix internal links [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([3b7f075](https://gitlab.com/commonground/nlx/nlx/commit/3b7f075))
* **docs:** nlx-ize the styling [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([120fedc](https://gitlab.com/commonground/nlx/nlx/commit/120fedc))
* **docs:** re-add contact to Support section [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([4b34112](https://gitlab.com/commonground/nlx/nlx/commit/4b34112))
* **docs:** re-add modification to the transaction log headers [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([5730916](https://gitlab.com/commonground/nlx/nlx/commit/5730916))
* **docs:** re-add support link to navbar [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([b822d63](https://gitlab.com/commonground/nlx/nlx/commit/b822d63))
* **docs:** remove default Docusaurus images [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([8ed0e41](https://gitlab.com/commonground/nlx/nlx/commit/8ed0e41))
* **docs:** replace logo with white variant [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([aeeeb59](https://gitlab.com/commonground/nlx/nlx/commit/aeeeb59))
* **docs:** update existing docs to be present in the sidebar of Docusaurus [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([56350c4](https://gitlab.com/commonground/nlx/nlx/commit/56350c4))
* **docs:** update favicon [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([1236c24](https://gitlab.com/commonground/nlx/nlx/commit/1236c24))
* **docs:** update syntax highlight theme to atom-one-dark [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([60ec7ac](https://gitlab.com/commonground/nlx/nlx/commit/60ec7ac))
* **docs:** use docs instead of the website as the application root URL [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([7c68de6](https://gitlab.com/commonground/nlx/nlx/commit/7c68de6))
* **docs:** use Prism for bash & toml code [#760](https://gitlab.com/commonground/nlx/nlx/issues/760) ([d7f38ec](https://gitlab.com/commonground/nlx/nlx/commit/d7f38ec))

# [0.78.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.77.0...v0.78.0) (2019-10-23)


### Bug Fixes

* add service config for haarlem's demo-api service to the skaffold config [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([2b8f217](https://gitlab.com/commonground/nlx/nlx/commit/2b8f217))
* move service config from parkeervergunning service to the demo-api service [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([9fcd39a](https://gitlab.com/commonground/nlx/nlx/commit/9fcd39a))
* **directory-ui:** clicking a service should show the detail pane [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([6de4485](https://gitlab.com/commonground/nlx/nlx/commit/6de4485))
* **directory-ui:** increase width for detail pane labels so 'email address' fits a single line [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([00f51e0](https://gitlab.com/commonground/nlx/nlx/commit/00f51e0))
* **directory-ui:** properly pass contact email address [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([bde1ed7](https://gitlab.com/commonground/nlx/nlx/commit/bde1ed7))
* **docs:** Add missing pages to 'Further reading' index ([a30150a](https://gitlab.com/commonground/nlx/nlx/commit/a30150a))
* **gitlab-ci:** fix review app URL [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([58ebe94](https://gitlab.com/commonground/nlx/nlx/commit/58ebe94))
* **gitlab-ci:** run e2e-tests in docker image [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([fa74a55](https://gitlab.com/commonground/nlx/nlx/commit/fa74a55))
* **gitlab-ci:** url for e2e tests [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([89c6871](https://gitlab.com/commonground/nlx/nlx/commit/89c6871))
* **nlxctl:** panic on incorrect shorthand flag ([ffd3129](https://gitlab.com/commonground/nlx/nlx/commit/ffd3129))


### Features

* **directory-ui:** add link to support to the header ([8516baa](https://gitlab.com/commonground/nlx/nlx/commit/8516baa))
* **directory-ui:** convert to static Header & Detail pane [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([087fb17](https://gitlab.com/commonground/nlx/nlx/commit/087fb17))
* **directory-ui:** create link from email address [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([43fa210](https://gitlab.com/commonground/nlx/nlx/commit/43fa210))
* **directory-ui:** hide contact section if no email is available [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([90e52cb](https://gitlab.com/commonground/nlx/nlx/commit/90e52cb))
* **directory-ui:** rename contact section to support section [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([98c3e6e](https://gitlab.com/commonground/nlx/nlx/commit/98c3e6e))
* **directory-ui:** replace custom Tooltip with native title [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([2cc755f](https://gitlab.com/commonground/nlx/nlx/commit/2cc755f))
* **directory-ui:** show detail pane with contact details when clicking a service [#449](https://gitlab.com/commonground/nlx/nlx/issues/449) ([36e7e4f](https://gitlab.com/commonground/nlx/nlx/commit/36e7e4f))
* **docs:** add Contacts to Support section and move common errors to a subsection ([a2410cc](https://gitlab.com/commonground/nlx/nlx/commit/a2410cc))
* **docs:** update header documentation ([afcd9f4](https://gitlab.com/commonground/nlx/nlx/commit/afcd9f4))
* **insight-ui:** add link to support to the header ([83004dc](https://gitlab.com/commonground/nlx/nlx/commit/83004dc))
* **outway:** the header X-NLX-Request-Data-Subject is not stripped from the outway ([6fbee5c](https://gitlab.com/commonground/nlx/nlx/commit/6fbee5c))
* add rdw with config-api deployment ([eb7789e](https://gitlab.com/commonground/nlx/nlx/commit/eb7789e))
* Upgrade IRMA server ([46425c1](https://gitlab.com/commonground/nlx/nlx/commit/46425c1))

# [0.77.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.76.0...v0.77.0) (2019-10-03)


### Bug Fixes

* **deps:** update dependency [@tippy](https://gitlab.com/tippy).js/react to v3 ([e9f8c10](https://gitlab.com/commonground/nlx/nlx/commit/e9f8c10))
* **directory-ui:** display documentation on Safari based browsers ([a5616b0](https://gitlab.com/commonground/nlx/nlx/commit/a5616b0))
* **directory-ui:** improve performance of search ([5ac4551](https://gitlab.com/commonground/nlx/nlx/commit/5ac4551))
* **directory-ui:** loosen CSP to allow loading content from external domains ([5ca9bf9](https://gitlab.com/commonground/nlx/nlx/commit/5ca9bf9))
* **outway:** fix only use healthy inways if there are more inway addres options ([150abf3](https://gitlab.com/commonground/nlx/nlx/commit/150abf3))


### Features

* **config-api:** add call to configure insight ([54c0023](https://gitlab.com/commonground/nlx/nlx/commit/54c0023))
* **directory-ui:** add query parameter for search ([a348fbc](https://gitlab.com/commonground/nlx/nlx/commit/a348fbc))
* **management-ui:** add basic auth ability to helm chart ([f898af6](https://gitlab.com/commonground/nlx/nlx/commit/f898af6))

# [0.76.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.75.0...v0.76.0) (2019-09-25)


### Bug Fixes

* **directory:** add inway_version to availabilities table. upgrade to postgres 10 ([2437c7d](https://gitlab.com/commonground/nlx/nlx/commit/2437c7d))
* **directory:** add test to health code ([578a948](https://gitlab.com/commonground/nlx/nlx/commit/578a948))
* **directory:** wip adding health check ([cdd0f86](https://gitlab.com/commonground/nlx/nlx/commit/cdd0f86))
* **docs:** added common errors support page ([ec48a0f](https://gitlab.com/commonground/nlx/nlx/commit/ec48a0f))
* **docs:** error messages now link to support pages ([d3fa282](https://gitlab.com/commonground/nlx/nlx/commit/d3fa282))
* **docs:** wrap / http2 is now supported ([5127691](https://gitlab.com/commonground/nlx/nlx/commit/5127691))
* **helm:** haarlem inway exposes application service ([94bc8ec](https://gitlab.com/commonground/nlx/nlx/commit/94bc8ec))
* **inway:** inway returns 503 statuscode on failed but configured api service ([c39d348](https://gitlab.com/commonground/nlx/nlx/commit/c39d348))
* **inway:** use same root ca for api spec doc ([8cfb3ac](https://gitlab.com/commonground/nlx/nlx/commit/8cfb3ac))
* **outway:** fix blocking / inresponsive outway caused by network changes / issues. ([c2ad2d9](https://gitlab.com/commonground/nlx/nlx/commit/c2ad2d9))
* **outway:** outway returns 503 statuscode on failed but announced service ([fc3ac6c](https://gitlab.com/commonground/nlx/nlx/commit/fc3ac6c))


### Features

* **config-api:** add inway filter to listservices endpoint ([5c3604b](https://gitlab.com/commonground/nlx/nlx/commit/5c3604b))
* **helm:** set api spec url for RDW service ([a3557af](https://gitlab.com/commonground/nlx/nlx/commit/a3557af))
* **inway:** can configure itself using the config-api ([fb3567f](https://gitlab.com/commonground/nlx/nlx/commit/fb3567f))
* **management-api:** add API to manage configuration of components within an organization ([5445450](https://gitlab.com/commonground/nlx/nlx/commit/5445450))
* **management-ui:** initial setup, ability to view inways and manage services ([b7a2e37](https://gitlab.com/commonground/nlx/nlx/commit/b7a2e37))
* **nlxctl:** parse multiple inway/service configs from a single file ([3cb5f9b](https://gitlab.com/commonground/nlx/nlx/commit/3cb5f9b))
* deploy rdw demo setup with the config-api ([e5cd7ef](https://gitlab.com/commonground/nlx/nlx/commit/e5cd7ef))

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
