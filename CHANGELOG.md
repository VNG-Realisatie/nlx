## [0.146.4](https://gitlab.com/commonground/nlx/nlx/compare/v0.146.3...v0.146.4) (2023-02-27)

## [0.146.3](https://gitlab.com/commonground/nlx/nlx/compare/v0.146.2...v0.146.3) (2023-02-23)

## [0.146.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.146.1...v0.146.2) (2023-02-23)

## [0.146.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.146.0...v0.146.1) (2023-02-22)

# [0.146.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.145.0...v0.146.0) (2023-02-13)


### Bug Fixes

* **helm:** directory chart uses correct existingSecret value ([a6c5039](https://gitlab.com/commonground/nlx/nlx/commit/a6c5039fc2f24650d5b8a02b86e5f8c84edbd96e)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)


### Features

* **management:** display feedback when removing an Inway fails ([129afbd](https://gitlab.com/commonground/nlx/nlx/commit/129afbdf8c96c96afa020303633e789db4dd0a9f)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)

# [0.145.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.144.1...v0.145.0) (2023-01-02)


### Features

* remove LoadBalancer IPs for Inways ([096897c](https://gitlab.com/commonground/nlx/nlx/commit/096897cbbae69d80722377cabcf7f680776d2e87)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)

## [0.144.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.144.0...v0.144.1) (2022-12-29)


### Bug Fixes

* **management:** use correct translation for sync error ([13ef255](https://gitlab.com/commonground/nlx/nlx/commit/13ef25554d424913111f1c35a3bfcf1e657adf84)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)

# [0.144.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.143.0...v0.144.0) (2022-11-29)


### Features

* **management:** replace toasters with an alert when synchronization of access requests fails ([64820f6](https://gitlab.com/commonground/nlx/nlx/commit/64820f6db49eb4e2ab02c1f626b44728659fda44)), closes [nlx#1635](https://gitlab.com/nlx/issues/1635)
* **management:** save OIDC sessions in database instead of cookie ([1dae1f5](https://gitlab.com/commonground/nlx/nlx/commit/1dae1f5b5acdb13a2078d7df840685ace4769632)), closes [nlx#1648](https://gitlab.com/nlx/issues/1648)

# [0.143.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.142.0...v0.143.0) (2022-11-08)


### Bug Fixes

* **management:** delete OutgoingAccessRequest when the remote proof can not be found ([c2e4c9c](https://gitlab.com/commonground/nlx/nlx/commit/c2e4c9c2ff57cfe29656e0660a114cf6e97c172d)), closes [nlx#1653](https://gitlab.com/nlx/issues/1653)


### Features

* **management:** add withdraw and terminate access as a consumer ([3adb068](https://gitlab.com/commonground/nlx/nlx/commit/3adb068c3b94aef60ab5c640033caa8d1084be4a)), closes [nlx#1420](https://gitlab.com/nlx/issues/1420)

# [0.142.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.141.0...v0.142.0) (2022-10-31)


### Bug Fixes

* **management:** display only the 100 latest auditlogs ([c6375bf](https://gitlab.com/commonground/nlx/nlx/commit/c6375bffd82d5a71f6ffe9bffc12a364c3ac642c)), closes [nlx#1633](https://gitlab.com/nlx/issues/1633)
* **management:** ensure syncing access requests on the Directory Overview page properly pauzes ([d24d56c](https://gitlab.com/commonground/nlx/nlx/commit/d24d56c6f44a80327d5b262ab202b03646038dd5)), closes [nlx#1632](https://gitlab.com/nlx/issues/1632)
* **management:** prevent request access model from flickering ([4581126](https://gitlab.com/commonground/nlx/nlx/commit/4581126613e1ea8b1d83b5b3a36c5585a8a03e17)), closes [nlx#1632](https://gitlab.com/nlx/issues/1632)
* **management:** show correct destination in transaction log when order is empty ([39bf749](https://gitlab.com/commonground/nlx/nlx/commit/39bf7493b9c509773ce5500158de93a2fcbf659e)), closes [nlx#1645](https://gitlab.com/nlx/issues/1645)


### Features

* **management:** show delegator in transaction log page when direction is OUT ([b0bcd4f](https://gitlab.com/commonground/nlx/nlx/commit/b0bcd4ff737eb338cda928401d435bf82203de5e)), closes [nlx#1636](https://gitlab.com/nlx/issues/1636)

# [0.141.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.140.0...v0.141.0) (2022-10-06)


### Bug Fixes

* add correct casing for creating services via nlxctl ([88ba50d](https://gitlab.com/commonground/nlx/nlx/commit/88ba50d6bd15be66dabaca8359097a0ffaee82e9)), closes [nlx#1612](https://gitlab.com/nlx/issues/1612)
* **inway:** use correct transaction log direction in postgres txlogger ([4cb73f3](https://gitlab.com/commonground/nlx/nlx/commit/4cb73f35914b5fefe03646147b936df72039ef28)), closes [nlx#1630](https://gitlab.com/nlx/issues/1630)
* let protoc openapiv2 output correct snake_case fieldnames ([b6dd8ba](https://gitlab.com/commonground/nlx/nlx/commit/b6dd8ba6330f8219bf2862e43c10cb6757ada407)), closes [nlx#1612](https://gitlab.com/nlx/issues/1612)


### Features

* **management:** configure OIDC client secret using an existing secret ([f738e3c](https://gitlab.com/commonground/nlx/nlx/commit/f738e3c677978c7401581e893cfd4c2d7c3819ab)), closes [nlx#1628](https://gitlab.com/nlx/issues/1628)

# [0.140.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.139.0...v0.140.0) (2022-09-28)


### Bug Fixes

* **management:** add missing organization name to toast when rejecting access ([1bbb9e7](https://gitlab.com/commonground/nlx/nlx/commit/1bbb9e7d6c368964d16e5b369187a1e4d8b91207)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)
* **management:** ensure Access Proof is synced when an AR gets approved ([9abd9aa](https://gitlab.com/commonground/nlx/nlx/commit/9abd9aa4b28abe65d3d01e9c698c0ed4505fad50)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)


### Features

* **management:** add sync & sync_all permissions to the readonly role ([d78d2eb](https://gitlab.com/commonground/nlx/nlx/commit/d78d2eb4c9085bfe3977ebc48cf02706317aeb6c)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)
* **management:** add sync of all outgoing access requests when adding an order ([757049c](https://gitlab.com/commonground/nlx/nlx/commit/757049c05a1c7c2047151829945c32190384da79)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)
* **management:** add sync of all outgoing access requests when opening the Directory overview ([da3e580](https://gitlab.com/commonground/nlx/nlx/commit/da3e580cc9d7b4120fd70330bb559defa1ceb2e0)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)
* **management:** add SyncAllOutgoingAccessRequests API endpoint ([b9a8254](https://gitlab.com/commonground/nlx/nlx/commit/b9a82549856b3605f176dbb1b30a1d00068d15ed)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)
* **management:** disable the scheduler job for Outgoing Access Requests ([a930af5](https://gitlab.com/commonground/nlx/nlx/commit/a930af5f6badd60a8828464fef229b963f86dc8e)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)
* **management:** display the organizations for which access syncing has failed ([a82047e](https://gitlab.com/commonground/nlx/nlx/commit/a82047ec6103d22f1d8da6161be8fedb01b206b4)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)
* **management:** display toaster if sync of Access Request fails for a directory service ([ea13f3b](https://gitlab.com/commonground/nlx/nlx/commit/ea13f3b0c79167330d8323babb049108dfaaba4a)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)
* **management:** introduce SyncOutgoingAccessRequests endpoint to enable syncing manually ([7754923](https://gitlab.com/commonground/nlx/nlx/commit/7754923242e7ecf64cf7babb1365add213bd3e87)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)
* **management:** pause syncing access requests if the user is visiting another tab ([2eee10f](https://gitlab.com/commonground/nlx/nlx/commit/2eee10f0b237822f9c0ebc39f9611d2acfad77ff)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)
* **management:** show spinner on Directory page until the sync has completed ([ddb5ef1](https://gitlab.com/commonground/nlx/nlx/commit/ddb5ef1680dab06d96226793ca24693b1b7b86e6)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)
* **management:** synchronize access requests every 3 seconds for the Directory overview page ([0727739](https://gitlab.com/commonground/nlx/nlx/commit/0727739f4257fb7b3b3d4b16e1c4069603182edf)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)

# [0.139.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.138.0...v0.139.0) (2022-09-21)


### Bug Fixes

* add missing copyright notice headers ([19cba1e](https://gitlab.com/commonground/nlx/nlx/commit/19cba1e16f048d24f9a5046b957c37527ddbeb82)), closes [nlx#1422](https://gitlab.com/nlx/issues/1422)
* add missing newlines ([bdc5015](https://gitlab.com/commonground/nlx/nlx/commit/bdc501539132dc2a650bfe54ec59d49efa9adb8b)), closes [nlx#1422](https://gitlab.com/nlx/issues/1422)
* **directory:** add 'norefferer' to documentation link to resolve vulnerability ([b4e5666](https://gitlab.com/commonground/nlx/nlx/commit/b4e5666e127ecdb530b485b80b5971a2b8c6f018)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)
* **directory:** enable clicking a service row to view the details of a service ([9941054](https://gitlab.com/commonground/nlx/nlx/commit/99410543f6848ecda59f6d7211b627eef61ab0e2)), closes [nlx#1617](https://gitlab.com/nlx/issues/1617)
* **management:** add 'norefferer' to documentation link to resolve vulnerability ([ab1055f](https://gitlab.com/commonground/nlx/nlx/commit/ab1055f78ed875fcc23cea499467846dc66ed9cc)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)
* **management:** correctly update Inway fields during registration ([283570c](https://gitlab.com/commonground/nlx/nlx/commit/283570ced6f11e86b13411c54c9d838f663728a7)), closes [nlx#1603](https://gitlab.com/nlx/issues/1603)
* **txlog:** allow datasubject keys with a minimal length of 1 character ([86640e5](https://gitlab.com/commonground/nlx/nlx/commit/86640e5f02ce64f354f071a8a70091fc87e55062)), closes [nlx#866](https://gitlab.com/nlx/issues/866)
* **txlog:** fix creating transactionlog without order ([f28bcf5](https://gitlab.com/commonground/nlx/nlx/commit/f28bcf5d8e4801dcbca955dd9fe9565e86dca37b)), closes [nlx#866](https://gitlab.com/nlx/issues/866)
* **txlog:** remove default tag from txlog-api helm values.yaml ([910bf63](https://gitlab.com/commonground/nlx/nlx/commit/910bf63eaaf3df109d53ba713e3cef242ca0fb3e)), closes [nlx#1619](https://gitlab.com/nlx/issues/1619)


### Features

* **directory:** add version endpoint to the directory api ([0889561](https://gitlab.com/commonground/nlx/nlx/commit/0889561b511c3e5ac140f95997d348f423201605)), closes [nlx#1598](https://gitlab.com/nlx/issues/1598)
* **docs:** add directory monitor egress ip ([d073106](https://gitlab.com/commonground/nlx/nlx/commit/d07310663e49a7a31e1eef995a4aad7801180161)), closes [nlx#1605](https://gitlab.com/nlx/issues/1605)
* **inway,outway:** include error in output message when setting up the transaction logger fails ([18687fd](https://gitlab.com/commonground/nlx/nlx/commit/18687fd9e22d61a2f66d52c6700c9e46b34359c8)), closes [nlx#866](https://gitlab.com/nlx/issues/866)
* **inway:** add option to helm chart to configure ca certificate ([003183f](https://gitlab.com/commonground/nlx/nlx/commit/003183fe9b49394de513af9e3bb5a4281548a52b)), closes [nlx#1601](https://gitlab.com/nlx/issues/1601)
* **inway:** add support for grpc transactionlog api ([eaf9c91](https://gitlab.com/commonground/nlx/nlx/commit/eaf9c9170f249d1f215babc9aeff0cb8f08bcdb7)), closes [nlx#866](https://gitlab.com/nlx/issues/866)
* **management:** keep track of errors using Outgoing Access Request when requesting access fails ([d64705b](https://gitlab.com/commonground/nlx/nlx/commit/d64705b25431be309f4f310873a0afdbf6d3bd5b)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)
* **management:** rework SendAccessRequest to directly communicate with the other organization ([b7c1db8](https://gitlab.com/commonground/nlx/nlx/commit/b7c1db8229a25e46b52168f5cd15fd18e5d66a8f)), closes [nlx#1602](https://gitlab.com/nlx/issues/1602)
* **outway:** add support for grpc transactionlog api ([b245312](https://gitlab.com/commonground/nlx/nlx/commit/b245312db5e5a934704544cf2854faf673555453)), closes [nlx#866](https://gitlab.com/nlx/issues/866)
* **txlog:** implement add record endpoint ([743d6f0](https://gitlab.com/commonground/nlx/nlx/commit/743d6f00a2f6ac84b7af03b8bc4e15150e27f3ea)), closes [nlx#866](https://gitlab.com/nlx/issues/866)

# [0.138.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.137.3...v0.138.0) (2022-08-23)


### Bug Fixes

* **management:** enable parsing audience as array in oidc access-token ([3a57a51](https://gitlab.com/commonground/nlx/nlx/commit/3a57a51e87154105d73df4d79964f7141474d412)), closes [nlx#1595](https://gitlab.com/nlx/issues/1595)


### Features

* **txlog:** add index to created column of records table ([7eec5c2](https://gitlab.com/commonground/nlx/nlx/commit/7eec5c2ab778b230f46f4509f580c8b9dda5a21d)), closes [nlx#785](https://gitlab.com/nlx/issues/785)

## [0.137.3](https://gitlab.com/commonground/nlx/nlx/compare/v0.137.2...v0.137.3) (2022-08-16)

## [0.137.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.137.1...v0.137.2) (2022-08-16)

## [0.137.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.137.0...v0.137.1) (2022-08-16)

# [0.137.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.136.0...v0.137.0) (2022-08-15)


### Features

* **management:** log errors of the authenticator ([20082e7](https://gitlab.com/commonground/nlx/nlx/commit/20082e7121883d4b4df04422147a3239a2241468)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)

# [0.136.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.135.1...v0.136.0) (2022-08-03)


### Bug Fixes

* **directory:** ensure a service is registered for the correct Inway ([f97b578](https://gitlab.com/commonground/nlx/nlx/commit/f97b578ac40c66814d3312014d9c5ca6c8c62436)), closes [nlx#1584](https://gitlab.com/nlx/issues/1584)
* **directory:** ignore transaction done error when registering a service ([4f43626](https://gitlab.com/commonground/nlx/nlx/commit/4f436268a6b4c20828a75d65ba420839ce33a52e)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)
* **inway,outway:** import Postgres driver ([4550323](https://gitlab.com/commonground/nlx/nlx/commit/4550323d9ca1dc57ad431057673156667c332fc2)), closes [nlx#1579](https://gitlab.com/nlx/issues/1579)


### Features

* **inway:** log error when offered service cannot be reached ([c07cbdf](https://gitlab.com/commonground/nlx/nlx/commit/c07cbdf04fa44b3ae54dad2639fc4eccc9736f96)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)
* **management:** add auditlog record for deleting an Outway ([b63b012](https://gitlab.com/commonground/nlx/nlx/commit/b63b0127c6eb165407b4611d96f79927f4cae2cb)), closes [nlx#1581](https://gitlab.com/nlx/issues/1581)
* **management:** add remove outway feature ([5478565](https://gitlab.com/commonground/nlx/nlx/commit/547856599bdd3fd68ee8c1ae46dc8f1ca329a838)), closes [nlx#1581](https://gitlab.com/nlx/issues/1581)
* **management:** show error when user has no permission in management UI ([a18271e](https://gitlab.com/commonground/nlx/nlx/commit/a18271ecedaf12b0c7ebecffee71f7052dea7b22)), closes [nlx#1547](https://gitlab.com/nlx/issues/1547)
* **management:** show feedback if accepting the ToS fails ([e8c2f2f](https://gitlab.com/commonground/nlx/nlx/commit/e8c2f2f336ecf4cc5773e60f3caf1f98d8847cd7)), closes [nlx#1547](https://gitlab.com/nlx/issues/1547)
* **management:** show feedback if adding an Inway fails ([2dcb0c4](https://gitlab.com/commonground/nlx/nlx/commit/2dcb0c4c7a0196f45e1d7ecfd002ec562e04ef09)), closes [nlx#1547](https://gitlab.com/nlx/issues/1547)
* **management:** show feedback if removing an Inway fails ([63a5d0e](https://gitlab.com/commonground/nlx/nlx/commit/63a5d0e43da1c60285380cbdc1d3299d46bf26f4)), closes [nlx#1547](https://gitlab.com/nlx/issues/1547)
* **management:** show feedback if updating the settings fails ([1841efd](https://gitlab.com/commonground/nlx/nlx/commit/1841efdf184b12afe819a696f21f87460c5e6867)), closes [nlx#1547](https://gitlab.com/nlx/issues/1547)
* **outway:** keep announcing to the management api after startup ([4b8f483](https://gitlab.com/commonground/nlx/nlx/commit/4b8f4837d4f44550df6802f5cff3248f7551ed56)), closes [nlx#1581](https://gitlab.com/nlx/issues/1581)
* return structured errors in json format when calling outway ([e99ad57](https://gitlab.com/commonground/nlx/nlx/commit/e99ad57baff48052f9ebf70e26f88d5fbe41a949)), closes [nlx#1557](https://gitlab.com/nlx/issues/1557)
* **txlog:** enable 'migrate status' command for the txlog-api + copy migrations from txlog-db ([2de8538](https://gitlab.com/commonground/nlx/nlx/commit/2de8538ea4733f0afd5f085c04c6b42713dd936a)), closes [nlx#1579](https://gitlab.com/nlx/issues/1579)
* **txlog:** enable 'migrate up' command for the txlog-api ([914797f](https://gitlab.com/commonground/nlx/nlx/commit/914797f33c07500c5facf214b29c36cd05751d59)), closes [nlx#1579](https://gitlab.com/nlx/issues/1579)
* **txlog:** introduce initContainer for txlog-api migrations + remove job-transaction-log.yaml ([d510ffb](https://gitlab.com/commonground/nlx/nlx/commit/d510ffbd062baee1c4297d7379121f4786ae97da)), closes [nlx#1579](https://gitlab.com/nlx/issues/1579)


### Reverts

* **ca-cfssl-unsafe:** use go 1.18.5 instead of go 1.19.0 ([25711a3](https://gitlab.com/commonground/nlx/nlx/commit/25711a32e05e58bc2ff74effe882eaa4ed86ab97)), closes [nlx#1581](https://gitlab.com/nlx/issues/1581)

## [0.135.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.135.0...v0.135.1) (2022-07-19)

# [0.135.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.134.1...v0.135.0) (2022-07-18)


### Bug Fixes

* **directory:** handle closing pg conn error ([d605e71](https://gitlab.com/commonground/nlx/nlx/commit/d605e71fbf0e75e0e382f1c128e465a370053644)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)
* enable opa for rvrd demo environment ([91cc0dc](https://gitlab.com/commonground/nlx/nlx/commit/91cc0dc640605833226680fff0f889db9933bd57)), closes [nlx#1577](https://gitlab.com/nlx/issues/1577)
* **outway:** handle closing log db connection error ([b32d81a](https://gitlab.com/commonground/nlx/nlx/commit/b32d81a79499bda1747ad204ab32eeb12ac5620d)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)
* update demo video application url ([47483fb](https://gitlab.com/commonground/nlx/nlx/commit/47483fb50b1203f4658c3d782fee0de2b3dc6f50)), closes [nlx#1577](https://gitlab.com/nlx/issues/1577)
* use correct service endpoint url for demo video api ([7ef9d7a](https://gitlab.com/commonground/nlx/nlx/commit/7ef9d7aaf1319cf1c3ba2d0280aee1e0a0eaef00)), closes [nlx#1577](https://gitlab.com/nlx/issues/1577)


### Features

* **management:** introduce readonly role ([ca0c638](https://gitlab.com/commonground/nlx/nlx/commit/ca0c638bacb5e21c2daa345b72868b5e55b4dc7c)), closes [nlx#1547](https://gitlab.com/nlx/issues/1547)
* **management:** move contact section to the top + convert to collapsible ([f7cb253](https://gitlab.com/commonground/nlx/nlx/commit/f7cb253a0c0d1005efeb14e0ed90428f3333675b)), closes [nlx#1576](https://gitlab.com/nlx/issues/1576)
* **management:** reintroduce access icon for the Directory page ([d38d2ee](https://gitlab.com/commonground/nlx/nlx/commit/d38d2ee45116b3630a4d0b3e58169474eb759253)), closes [nlx#1576](https://gitlab.com/nlx/issues/1576)

## [0.134.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.134.0...v0.134.1) (2022-06-30)

# [0.134.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.133.2...v0.134.0) (2022-06-30)


### Features

* **inway:** add split service ability to the helm chart ([e18a442](https://gitlab.com/commonground/nlx/nlx/commit/e18a442430ca5825be56a9df2ec251eac5408a6d)), closes [nlx#1563](https://gitlab.com/nlx/issues/1563)

## [0.133.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.133.1...v0.133.2) (2022-06-28)

## [0.133.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.133.0...v0.133.1) (2022-06-28)

# [0.133.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.132.2...v0.133.0) (2022-06-28)


### Features

* **management:** add rbac permission system ([90f3049](https://gitlab.com/commonground/nlx/nlx/commit/90f304918cd7c3fa4f9d21de2b4901a82377c56f)), closes [nlx#1545](https://gitlab.com/nlx/issues/1545)

## [0.132.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.132.1...v0.132.2) (2022-06-27)

## [0.132.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.132.0...v0.132.1) (2022-06-27)

# [0.132.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.131.1...v0.132.0) (2022-06-27)


### Bug Fixes

* **management:** show destination organization for outgoing transactions ([79a5905](https://gitlab.com/commonground/nlx/nlx/commit/79a59058f8287633f9791540297eb7df91e9c961)), closes [nlx#1555](https://gitlab.com/nlx/issues/1555)


### Features

* **management:** display organization name instead of OIN for the audit log page ([1bb1be0](https://gitlab.com/commonground/nlx/nlx/commit/1bb1be0879cc68b4b9292b9c03344b7bcdbc9ff9)), closes [nlx#1555](https://gitlab.com/nlx/issues/1555)
* **management:** display organization name instead of OIN for the orders page ([51f9b7b](https://gitlab.com/commonground/nlx/nlx/commit/51f9b7b89466ff22dd1f0b2960741cf7337672ce)), closes [nlx#1555](https://gitlab.com/nlx/issues/1555)
* **management:** display organization name instead of OIN for the transaction log page ([665a13c](https://gitlab.com/commonground/nlx/nlx/commit/665a13c0a0bce0f1d771fab40c31130a0871c4f1)), closes [nlx#1555](https://gitlab.com/nlx/issues/1555)
* **management:** return organization name with serial number for the transaction log service ([4e171f8](https://gitlab.com/commonground/nlx/nlx/commit/4e171f8ccffa3aa504e624d488f72730641d89fd)), closes [nlx#1555](https://gitlab.com/nlx/issues/1555)
* **management:** show OIN instead of name if the organization name is not available ([a9fb881](https://gitlab.com/commonground/nlx/nlx/commit/a9fb8813e7bf3a32720b9e1925d7dcfe41a3f944)), closes [nlx#1555](https://gitlab.com/nlx/issues/1555)

## [0.131.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.131.0...v0.131.1) (2022-06-10)

# [0.131.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.130.3...v0.131.0) (2022-06-07)


### Bug Fixes

* **management:** fix unsetting org inway when deleting inway ([5312610](https://gitlab.com/commonground/nlx/nlx/commit/531261063b12eb18a998887e71e65186f1f636d4)), closes [nlx#1458](https://gitlab.com/nlx/issues/1458)
* **management:** fix wrong title for inway delete button ([a477092](https://gitlab.com/commonground/nlx/nlx/commit/a4770924aabf365a4b42789feb6231973c28f639)), closes [nlx#1458](https://gitlab.com/nlx/issues/1458)


### Features

* **management:** display transaction id in the transaction log ([ef8f56a](https://gitlab.com/commonground/nlx/nlx/commit/ef8f56a15d017479d74a764192be1b2c25dc4b28)), closes [nlx#1554](https://gitlab.com/nlx/issues/1554)
* **management:** show error message when requesting access fails ([935b25f](https://gitlab.com/commonground/nlx/nlx/commit/935b25f3a71dbd67d6e95d44530165ecd5ac5523)), closes [nlx#1545](https://gitlab.com/nlx/issues/1545)
* **management:** show toaster when succesfully deleted inway in management ui ([01f9071](https://gitlab.com/commonground/nlx/nlx/commit/01f907128d4fa741893521c25a30881ecf7392d9)), closes [nlx#1458](https://gitlab.com/nlx/issues/1458)

## [0.130.3](https://gitlab.com/commonground/nlx/nlx/compare/v0.130.2...v0.130.3) (2022-05-30)


### Reverts

* update helm release parkeerrechten-admin to v0.2.0 ([a0eac3b](https://gitlab.com/commonground/nlx/nlx/commit/a0eac3b94b2c3674e1af1cd6687041053cc8ffeb)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)

## [0.130.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.130.1...v0.130.2) (2022-05-25)

## [0.130.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.130.0...v0.130.1) (2022-05-19)

# [0.130.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.129.0...v0.130.0) (2022-05-18)


### Features

* **helm:** use pull policy 'Always' as default for all NLX Helm charts ([3ca215e](https://gitlab.com/commonground/nlx/nlx/commit/3ca215ece513f9f04764f9522ba2304b2beb9275)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)

# [0.129.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.128.2...v0.129.0) (2022-05-17)


### Bug Fixes

* **directory:** return ManagementAPIProxyAddress instead of the regular address ([d71e063](https://gitlab.com/commonground/nlx/nlx/commit/d71e063762b5b82769696ca674665a9225318c60)), closes [nlx#1518](https://gitlab.com/nlx/issues/1518)
* **docs:** fix link to try-nlx docker introduction page ([58d378d](https://gitlab.com/commonground/nlx/nlx/commit/58d378d36c96c59737087aae8ab3cbc313bd0950)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)
* **docs:** re-enable tabs for OS specific commands for Getting up and running ([254bee3](https://gitlab.com/commonground/nlx/nlx/commit/254bee30f18d3866dcc4b0732561178baa0b44d2)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)


### Features

* add structured errors for grpc endpoints in outway and management api ([0277122](https://gitlab.com/commonground/nlx/nlx/commit/0277122e693722df3bed6496bd76a886a9576ee7)), closes [nlx#1522](https://gitlab.com/nlx/issues/1522)
* add the Manage Citizens UI to review, ACC and demo environment ([02c03c1](https://gitlab.com/commonground/nlx/nlx/commit/02c03c1e4e530728a44f21eed815f77a747a6eae)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)
* **outway:** rename authorization header to X-Nlx-Autorization ([0e030b3](https://gitlab.com/commonground/nlx/nlx/commit/0e030b32b7e62347fba78075305e4a1bb1880bc3)), closes [nlx#1527](https://gitlab.com/nlx/issues/1527)
* **outway:** strip authorization headers ([bd7a6b6](https://gitlab.com/commonground/nlx/nlx/commit/bd7a6b6d85bfce147a9cf58fc2d459065321d911)), closes [nlx#1527](https://gitlab.com/nlx/issues/1527)
* remove custom implementation of the auth-service ([21ab9fe](https://gitlab.com/commonground/nlx/nlx/commit/21ab9fee6b53d463f6c83696aa0dc4e1069048c7)), closes [nlx#1527](https://gitlab.com/nlx/issues/1527)


### Reverts

* fix security jobs from GitLab without a stage ([96cea58](https://gitlab.com/commonground/nlx/nlx/commit/96cea58e6941e9877f7973573c6aed88b747ffdd)), closes [nlx#1171](https://gitlab.com/nlx/issues/1171)

## [0.128.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.128.1...v0.128.2) (2022-04-26)

## [0.128.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.128.0...v0.128.1) (2022-04-25)

# [0.128.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.127.0...v0.128.0) (2022-04-25)


### Bug Fixes

* **management:** handle access request being revoked before access proof is synced ([1850f9e](https://gitlab.com/commonground/nlx/nlx/commit/1850f9ed30cc5346e43f9e8e2e264e4d26200606)), closes [nlx#1526](https://gitlab.com/nlx/issues/1526)
* **management:** remove duplicate 'One time cost' field from the request access dialog ([78ce73d](https://gitlab.com/commonground/nlx/nlx/commit/78ce73de0af232c450461e963d869c2e8151fc01)), closes [nlx#1525](https://gitlab.com/nlx/issues/1525)


### Features

* **inway:** created separation between nlx errors and other errors ([8c3b406](https://gitlab.com/commonground/nlx/nlx/commit/8c3b40667b2c3045b2b51d5f6cc4b6f9f047051d)), closes [nlx#1521](https://gitlab.com/nlx/issues/1521)
* **inway:** enable specifying the Management API Proxy address ([5ba951c](https://gitlab.com/commonground/nlx/nlx/commit/5ba951cc92f4c80db37ff7b16981ffb5a56a5b6c)), closes [nlx#1503](https://gitlab.com/nlx/issues/1503)
* **management:** auto sync organization email address to directory ([d38f57f](https://gitlab.com/commonground/nlx/nlx/commit/d38f57fd56eb264e0c1d73f8cb52f4adec17ef37)), closes [#1444](https://gitlab.com/commonground/nlx/nlx/issues/1444)
* **outway:** return 540 as HTTP status code for all NLX related errors ([7dbfbe9](https://gitlab.com/commonground/nlx/nlx/commit/7dbfbe97864ccb42a721cb48747941641a8f4f4a)), closes [nlx#1521](https://gitlab.com/nlx/issues/1521)


### Reverts

* remove authorization server from local setup ([ca81620](https://gitlab.com/commonground/nlx/nlx/commit/ca8162064575d36c60f5093053ba0859e2d53601)), closes [nlx#1527](https://gitlab.com/nlx/issues/1527) [nlx#1466](https://gitlab.com/nlx/issues/1466)

# [0.127.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.126.0...v0.127.0) (2022-04-04)


### Features

* inway and outway allow and preserve trailing slashes ([98158bc](https://gitlab.com/commonground/nlx/nlx/commit/98158bcb928e821d72261e88aacd60569971fc7f))
* **inway:** add authorization plugin to inway ([c4dc4fe](https://gitlab.com/commonground/nlx/nlx/commit/c4dc4fe3bb05cca6e8edc8d5a877bfdfd36a336f))
* **management:** replace Serial Number with OIN in the Management UI ([4b504b5](https://gitlab.com/commonground/nlx/nlx/commit/4b504b5eb0c36345082931178a8c43354d77263c)), closes [#1498](https://gitlab.com/commonground/nlx/nlx/issues/1498)
* **management:** separate OIN from organization name in the confirmation modals ([3c50548](https://gitlab.com/commonground/nlx/nlx/commit/3c5054813a5bc22a48959b0f5f4da5aa6275c594)), closes [#1498](https://gitlab.com/commonground/nlx/nlx/issues/1498)
* **outway:** change auth plugin to work with OPA servers ([0e97b43](https://gitlab.com/commonground/nlx/nlx/commit/0e97b43b612d98a1a0cee502c77a627f71fc0932))

# [0.126.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.125.0...v0.126.0) (2022-03-22)


### Bug Fixes

* **management:** show 404 view for directory service when it is removed while polling ([fe82b8b](https://gitlab.com/commonground/nlx/nlx/commit/fe82b8bd4b28bc4a2f19d44d9b49c156c37dee9a)), closes [#1499](https://gitlab.com/commonground/nlx/nlx/issues/1499)


### Features

* **outway:** do not require trailing slash in service call ([b702fa8](https://gitlab.com/commonground/nlx/nlx/commit/b702fa813d36546f4df24d7dba91379b306d4506))

# [0.125.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.124.0...v0.125.0) (2022-03-17)


### Bug Fixes

* add delegatee public key fingerprint to request claim call ([51ad428](https://gitlab.com/commonground/nlx/nlx/commit/51ad4284b4b5633589bedc6275c07d8b3046c586))
* **docs:** update Algolia settings to re-enable search ([1c050fd](https://gitlab.com/commonground/nlx/nlx/commit/1c050fddb39cd56a8f5710e5ef46f36dce0aaf5b)), closes [#1491](https://gitlab.com/commonground/nlx/nlx/issues/1491)
* **management:** add default value outway self address to migration ([235ed73](https://gitlab.com/commonground/nlx/nlx/commit/235ed736758d9a38791c324f16d96d9ba994c711))
* **management:** always resume polling the directory service ([73081fc](https://gitlab.com/commonground/nlx/nlx/commit/73081fc7e62c56eddcbf6653f628e1d05b205068)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** correct revoke access translation and add finger print to access grant ([83d20dc](https://gitlab.com/commonground/nlx/nlx/commit/83d20dcfff830e09e33cdee82bbedf189e69281e))
* **management:** enable displaying the incoming order detail drawer ([ab92fd5](https://gitlab.com/commonground/nlx/nlx/commit/ab92fd566837d6169e073f4f848c84088e3d952e)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** go import'ed ([3c66034](https://gitlab.com/commonground/nlx/nlx/commit/3c660346598aaf25effb7711c921ac264198cecf))
* **management:** make retry requesting access work again ([2638804](https://gitlab.com/commonground/nlx/nlx/commit/2638804c9abc44e81894ac183551f265cab5e275)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** reword error message ([e66e77c](https://gitlab.com/commonground/nlx/nlx/commit/e66e77cf21c3ad6d3c3b8e9457d0de9b3e575b4c)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** rework broken Edit Order page ([4211e71](https://gitlab.com/commonground/nlx/nlx/commit/4211e718ca4a482e714365b68cf487ad0a2d0052)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** send the public key PEM when requesting access to a service ([8f083b0](https://gitlab.com/commonground/nlx/nlx/commit/8f083b0d7980e60bf4191dfbb9f9daae6bb7b8c3)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** show access request as received when its approved but has no prove yet ([81e2c2a](https://gitlab.com/commonground/nlx/nlx/commit/81e2c2a0ede6d33081ac77d63cd77e6f8524c670)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** use correct organization serial number to update a DirectoryService model ([3e2e6e5](https://gitlab.com/commonground/nlx/nlx/commit/3e2e6e54fea4e12f92d43839c1a529973e4b09df)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** using an order works ([41f04bd](https://gitlab.com/commonground/nlx/nlx/commit/41f04bd0b74a0d2caf99ef1205a4d4bc0a1ced13))


### Features

* add outway self addres to modd and helm chart ([de7f872](https://gitlab.com/commonground/nlx/nlx/commit/de7f8728a38217e88ea203f71fdea49f47f589a7))
* **directory:** delegation claims use access proofs ([fe46215](https://gitlab.com/commonground/nlx/nlx/commit/fe46215dc18d8d3806a25a185d7a4cef9139b9e1))
* **inway:** improved delegation error messages ([b88c4bd](https://gitlab.com/commonground/nlx/nlx/commit/b88c4bd9c5f382e19602d378d190f307c21fcb3f))
* **management:** add get outways by public key fingerprint call and register outway self address ([ff94f64](https://gitlab.com/commonground/nlx/nlx/commit/ff94f64035bc28d797537fcfe454546858ca939e))
* **management:** add migration for outgoing_orders_access_proofs table ([448c108](https://gitlab.com/commonground/nlx/nlx/commit/448c1085056b8703bc34ba3b0b2fa2cdc11c7d17))
* **management:** add pub key fingerprint to outway ([3d434f6](https://gitlab.com/commonground/nlx/nlx/commit/3d434f656bfe647ea032efab1380b45d1c3a4e7f))
* **management:** add public fingerprint to list incoming access request ([c876d3b](https://gitlab.com/commonground/nlx/nlx/commit/c876d3b188bf33fb2d58855f2461f1a66d9135e3))
* **management:** add public key fingerprint to access request detail page ([6cad222](https://gitlab.com/commonground/nlx/nlx/commit/6cad222ff2a362508004c7eeba9a89c500cc8dd9))
* **management:** add public key fingerprint to create orders ([3eb5410](https://gitlab.com/commonground/nlx/nlx/commit/3eb54106dfeeb14c20dcbc678ba0c33dd973c4ab))
* **management:** add public key PEM to create outgoing access request and duplicate service check ([46e16c9](https://gitlab.com/commonground/nlx/nlx/commit/46e16c93fb433c77f67e6ed85964caf8b406930b))
* **management:** disable selecting an access proof of an order for an already selected service ([805c2b2](https://gitlab.com/commonground/nlx/nlx/commit/805c2b2c2c256220cf4ba5035980b4b1d0f71171)), closes [#1491](https://gitlab.com/commonground/nlx/nlx/issues/1491)
* **management:** display Outways & public key in the request access modal ([9ad5f9c](https://gitlab.com/commonground/nlx/nlx/commit/9ad5f9ca3d4e558a0f34ebab1e14710aceb6f4f4)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** display warning when one of the access requests failed for a service ([2f60a48](https://gitlab.com/commonground/nlx/nlx/commit/2f60a481edf0933fcede42c74dfa17010f756c27)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** enable listing incoming and outgoing orders with public key fingerprints ([439ebe1](https://gitlab.com/commonground/nlx/nlx/commit/439ebe1feccd7baecacd775a0f36aa1b62406816)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** include access proof models for ListOutgoingOrders ([e3d3be4](https://gitlab.com/commonground/nlx/nlx/commit/e3d3be4eb59049fd786bc0fc76b931d0f6e5f34b)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** include access proof models for ListOutgoingOrders (wip) ([a4cee4e](https://gitlab.com/commonground/nlx/nlx/commit/a4cee4e639584d236e973a9f569cc0037868ceef)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** include ListenAddress when listing the Outways ([7f7b3b3](https://gitlab.com/commonground/nlx/nlx/commit/7f7b3b3557d0ef3dda765e7b1b0f28463f7b01f9)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** make access detail modal more organized ([0eec9d0](https://gitlab.com/commonground/nlx/nlx/commit/0eec9d09d0440848b99d351a0b19c3297a25d460))
* **management:** only show the subject serial number for the incoming orders ([0a0c183](https://gitlab.com/commonground/nlx/nlx/commit/0a0c183f752a47d2552f3f15dfba1df5e59335f9)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** orders use accessproofs ([29ae245](https://gitlab.com/commonground/nlx/nlx/commit/29ae245d8c9e7e0c677098c380d4a5107b254c37))
* **management:** provide organization serialnumber when retrieving a claim ([fabc95b](https://gitlab.com/commonground/nlx/nlx/commit/fabc95b36ab330f9257c713738807df79b9ed15e))
* **management:** regenerate management api client ([89f790e](https://gitlab.com/commonground/nlx/nlx/commit/89f790ee76ac50b2f5227f77f1c7e2fb662c571e)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** regenerate management api client ([dacfbe9](https://gitlab.com/commonground/nlx/nlx/commit/dacfbe9279ff8698c850b2bbd8f1ddda44bae4c5)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** remove column title for access ([80978fe](https://gitlab.com/commonground/nlx/nlx/commit/80978fe21eafb1149653a0dd49e404434af45a98))
* **management:** remove the cancelled state for the access requests ([59589d3](https://gitlab.com/commonground/nlx/nlx/commit/59589d35ba46a28d593e6949af48edea35bb8011)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** return access state per public key fingerprint in list services ([cc9d2ed](https://gitlab.com/commonground/nlx/nlx/commit/cc9d2ed1d8db1aa9773b17577dc8e48144c1defb))
* **management:** rework code & ui of access state for Outways ([9c5830d](https://gitlab.com/commonground/nlx/nlx/commit/9c5830d137de44ca7322be077da58f54f46c0d36)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** rework determining the access state based on the revoked at property ([22d8f8c](https://gitlab.com/commonground/nlx/nlx/commit/22d8f8cbe22896b8a8749fe20ec93c690d82f019)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** rework the AddOrderPage to use Public Key Fingerprints for access ([8ba71b1](https://gitlab.com/commonground/nlx/nlx/commit/8ba71b15fdbb9cb3bd815e59da607bef597c2c1c)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** show status of access requests per certificate ([34eeedf](https://gitlab.com/commonground/nlx/nlx/commit/34eeedf282fa7dcc58bfb7faf755745a3dd57609)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** store access proof models when creating an outgoing order model ([883de52](https://gitlab.com/commonground/nlx/nlx/commit/883de524a3583f83908e944cef37240936d90c21)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** sync access requests using the public key fingerprint ([4a921f7](https://gitlab.com/commonground/nlx/nlx/commit/4a921f781cab2a2b7dc14bee55187b559f71d675)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** use full with for alert with access request error ([121d9cc](https://gitlab.com/commonground/nlx/nlx/commit/121d9cc46152c85e06b051514cedf7f6e294fdcb)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **outway:** add grpc server and sign order claim call ([62d193f](https://gitlab.com/commonground/nlx/nlx/commit/62d193f1699f37555dd42bd3ec9c6e66528c194a))
* **outway:** improved error messages delegation ([ea7572c](https://gitlab.com/commonground/nlx/nlx/commit/ea7572cd3833d7866654e4abb6575d9ebeeedc51))
* **outway:** pass listen address when registering to the Management API ([13ae3f5](https://gitlab.com/commonground/nlx/nlx/commit/13ae3f5359e1be86fbbfb399426ce1f7407aa387)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **outway:** register self address in management api ([0991d7f](https://gitlab.com/commonground/nlx/nlx/commit/0991d7fa27a1cf218cff1b2b5c381f1b8a904d94))
* remove passing listen address when registering an Outway ([23190c4](https://gitlab.com/commonground/nlx/nlx/commit/23190c4b7bbd4ecaabdf542a6b447c5bc51cc180)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)


### Reverts

* **management:** include ListenAddress when listing the Outways ([3f26d76](https://gitlab.com/commonground/nlx/nlx/commit/3f26d76886195daa9d7dea82b5060ca1fe7617e1))
* **outway:** pass listen address when registering to the Management API ([4284a87](https://gitlab.com/commonground/nlx/nlx/commit/4284a87366c0952f3b926a5696b0434124794db0))
* update dependency @commonground/design-system to v21.1.1 ([4bd2e2c](https://gitlab.com/commonground/nlx/nlx/commit/4bd2e2c46df39323de745ef634e3eaffdf702312))

# [0.124.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.123.0...v0.124.0) (2022-02-28)


### Bug Fixes

* **directory:** ensure external links of the Navigation are opened correctly ([ca02b92](https://gitlab.com/commonground/nlx/nlx/commit/ca02b92241028ab55d6ea7a5ec75841285278199))
* **docs:** update filename in command to fix file permissions ([daa5433](https://gitlab.com/commonground/nlx/nlx/commit/daa5433e06fb8a4858f85cabea2fbb555e54562d))
* **docs:** upgrade the Algolia credentials to use the team core mailbox account ([91596c1](https://gitlab.com/commonground/nlx/nlx/commit/91596c1a83d701ea2f253f26bed57ba1168ba905))


### Features

* **directory:** made database connection parameters configurable for directory-monitor ([45aea7e](https://gitlab.com/commonground/nlx/nlx/commit/45aea7e896c3d4fa1511fd2c0edd0928881cad51))

# [0.123.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.122.0...v0.123.0) (2022-02-11)


### Bug Fixes

* **directory:** increase health check interval of the monitor ([78b08ca](https://gitlab.com/commonground/nlx/nlx/commit/78b08ca8a05f2d3b360f57d2de0cdecadc2914de))
* **management:** ensure only unlocked access requests are being synced ([5e586d9](https://gitlab.com/commonground/nlx/nlx/commit/5e586d98dad04fd870bce275c0f018d9f22d3ad8))
* **management:** typo in transaction logs of requests made in behalf of an other organization ([facb26a](https://gitlab.com/commonground/nlx/nlx/commit/facb26a8f4194d91d3f03da2b08537039c67c4b2))


### Features

* **management:** remove stacktrace drawer when showing a failed outgoing access request ([cc43af9](https://gitlab.com/commonground/nlx/nlx/commit/cc43af901e6c7a50e2c8cdd0fd40c79142894e42)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)
* **management:** remove the quick access button from the Directory page ([a0bee7a](https://gitlab.com/commonground/nlx/nlx/commit/a0bee7a391fa632dbd27e6e95e2c11f9b4e94bfe)), closes [#1427](https://gitlab.com/commonground/nlx/nlx/issues/1427)

# [0.122.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.121.1...v0.122.0) (2022-02-07)


### Bug Fixes

* enable ToS for production only ([323131c](https://gitlab.com/commonground/nlx/nlx/commit/323131ca3beb84170dd805fde3659930ff737e63)), closes [#1438](https://gitlab.com/commonground/nlx/nlx/issues/1438)
* **management:** add specific error message for empty inway address ([299495f](https://gitlab.com/commonground/nlx/nlx/commit/299495f09635a03f5a8ddc3dada4bc031ca9d368))
* **management:** ensure audit log for accepting ToS can be handled ([60d3e91](https://gitlab.com/commonground/nlx/nlx/commit/60d3e91ee24fd711ac838d849c9c9f5c8722e84a)), closes [#1438](https://gitlab.com/commonground/nlx/nlx/issues/1438)
* **management:** ensure ToS can only be accepted once ([411413a](https://gitlab.com/commonground/nlx/nlx/commit/411413a3ada7ba0dccc723d7ac4c90def98210f7))
* **management:** prioritize OutgoingAccessRequests with state 'created' when syncing access requests ([1bd2eeb](https://gitlab.com/commonground/nlx/nlx/commit/1bd2eeb1410d9d31ba69f0ab149c03b811670d37)), closes [#1429](https://gitlab.com/commonground/nlx/nlx/issues/1429)
* **management:** removed confusing horizontal lines on login page ([bbebd14](https://gitlab.com/commonground/nlx/nlx/commit/bbebd14adb83d49afcd24dff4149760aa0650481))
* **management:** show org name in access grant toaster ([5c73924](https://gitlab.com/commonground/nlx/nlx/commit/5c73924f805ba3df27f2719f6f1aa5e287eb45dd))
* **management:** sync access proofs even if the outgoing access request has the state failed ([4696e1d](https://gitlab.com/commonground/nlx/nlx/commit/4696e1d3b97b704d07c7ada7086edcd0cb09f0d3))


### Features

* **directory:** add terms of service flag to helm chart ([5c10fa9](https://gitlab.com/commonground/nlx/nlx/commit/5c10fa9c6cafc930a35941fe9c97ffe4ccab6934))
* **directory:** add terms of service to the directory api ([7a44390](https://gitlab.com/commonground/nlx/nlx/commit/7a44390ed28f0da2d64b04d760d8774fa0a40061))
* **management:** add terms of service status and accept routes in api ([fa6babe](https://gitlab.com/commonground/nlx/nlx/commit/fa6babe0e61ff10abe49a66ccab541d755fe688c))
* **management:** add the Terms and Condition route ([af971d6](https://gitlab.com/commonground/nlx/nlx/commit/af971d62c62d803db6be66e59000e7b43324e52f)), closes [#1438](https://gitlab.com/commonground/nlx/nlx/issues/1438)
* **management:** do not enforce accepting ToS when the directory is unavailable ([bb0aa52](https://gitlab.com/commonground/nlx/nlx/commit/bb0aa5204376ce97cbd442ce23dbec845c8b531d)), closes [#1438](https://gitlab.com/commonground/nlx/nlx/issues/1438)
* **management:** enable accepting the ToS ([3b66d17](https://gitlab.com/commonground/nlx/nlx/commit/3b66d1792480d181d210298a0681b6c70015195f)), closes [#1438](https://gitlab.com/commonground/nlx/nlx/issues/1438)
* **management:** enable retrieving the ToS using the ApplicationStore ([92b2c49](https://gitlab.com/commonground/nlx/nlx/commit/92b2c492fced7c04e8d41dae172719b5054f0ee0)), closes [#1438](https://gitlab.com/commonground/nlx/nlx/issues/1438)
* **management:** fetch the terms of service status as part of the ToS context ([7b145a5](https://gitlab.com/commonground/nlx/nlx/commit/7b145a5e3569d1c91de999e21da2cb48ceee91c2)), closes [#1438](https://gitlab.com/commonground/nlx/nlx/issues/1438)
* **management:** implement AuditLog for accepting Terms of Service ([333430e](https://gitlab.com/commonground/nlx/nlx/commit/333430e2ccd94570d9b40dd77052cb5b0c5df6bb)), closes [#1438](https://gitlab.com/commonground/nlx/nlx/issues/1438)
* **management:** implement GetTermsOfService API call for the Management API ([6e5e8f4](https://gitlab.com/commonground/nlx/nlx/commit/6e5e8f45b6d22962e1580b8970060901314a339a)), closes [#1438](https://gitlab.com/commonground/nlx/nlx/issues/1438)
* **management:** implement Terms of Service context provider ([dfef6a1](https://gitlab.com/commonground/nlx/nlx/commit/dfef6a111da5e84913c6346360ac1809d39c2baf)), closes [#1438](https://gitlab.com/commonground/nlx/nlx/issues/1438)
* **management:** lower the interval for synchronizing access requests ([0ecd15e](https://gitlab.com/commonground/nlx/nlx/commit/0ecd15e08f381ae2391db9b16cf64bd63869b3ff))
* **management:** only redirect to ToS when ToS is enabled for the directory ([4875fdd](https://gitlab.com/commonground/nlx/nlx/commit/4875fdd8e109cd1488312e46125188abab46ebcd)), closes [#1438](https://gitlab.com/commonground/nlx/nlx/issues/1438)
* **management:** scheduler processes all outdated access requests ([b1d3c23](https://gitlab.com/commonground/nlx/nlx/commit/b1d3c2319311f2867c021742a4fd8f4c65138ec8))
* **management:** update ToS link text ([154998c](https://gitlab.com/commonground/nlx/nlx/commit/154998c6e407556cb26c536c398a07ac0aac47c8))


### Reverts

* strip nlxio prefix from Docker images when not using Docker Hub ([75c4e14](https://gitlab.com/commonground/nlx/nlx/commit/75c4e1493a7384de8b11493d52af5a879f6fe946))
* update dependency axe-core to v4.4.0 ([5820288](https://gitlab.com/commonground/nlx/nlx/commit/58202880e4f86b4f183d734c26692dfca5ab46e2))

## [0.121.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.121.0...v0.121.1) (2022-01-12)


### Bug Fixes

* **management:** set txlog basic auth headers ([3591a05](https://gitlab.com/commonground/nlx/nlx/commit/3591a052ad8d0a6c1f133df42b3a14fca61c66e8))

# [0.121.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.120.0...v0.121.0) (2022-01-12)


### Bug Fixes

* downgrade Node v17.x to v16.x to prevent OpenSSL issues ([dbb046e](https://gitlab.com/commonground/nlx/nlx/commit/dbb046eba15ba22d6c84e6f7ca7b592203eec2ee)), closes [/github.com/reactjs/reactjs.org/issues/4073#issuecomment-1010758303](https://gitlab.com//github.com/reactjs/reactjs.org/issues/4073/issues/issuecomment-1010758303)
* **outway:** correctly add authorization plugin ([763bab4](https://gitlab.com/commonground/nlx/nlx/commit/763bab4ce8691d0e812d38d88a3875ea30f276f1))


### Features

* **management:** add transactionlog to the user interface ([805269d](https://gitlab.com/commonground/nlx/nlx/commit/805269d660aad1f557ec82e9d66b89571fd8780b))
* **management:** show message when transaction log api is disabled ([74f9993](https://gitlab.com/commonground/nlx/nlx/commit/74f999389ab61c2c7e20e301e4fd93a8e4be6934))
* **txlog:** add txlog api and add txlog client to management api ([2941495](https://gitlab.com/commonground/nlx/nlx/commit/2941495b5069d6096ee89140b271e5c99a83a439))


### Performance Improvements

* **management:** prevent useless rerender of the Avatar component ([e6f7c1f](https://gitlab.com/commonground/nlx/nlx/commit/e6f7c1f84e2f386965907224c6b7001cbd512f32))

# [0.120.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.119.0...v0.120.0) (2021-12-22)


### Bug Fixes

* add frontend validation to organization email field and remove required field ([ec5fa95](https://gitlab.com/commonground/nlx/nlx/commit/ec5fa955d895c5f9f7b2b22aa102009d24443a57))
* **directory:** add directory.nlx.io as Demo environment to environments menu ([010febf](https://gitlab.com/commonground/nlx/nlx/commit/010febf2da60620360c67e1e17b5c1aaa2b543c2))
* **management:** corrected typo in settings page ([47e905c](https://gitlab.com/commonground/nlx/nlx/commit/47e905cfafa15a3d8a99fd354d9a3436967abc19))
* **management:** preselect selected services when editing an order ([c4dbfbc](https://gitlab.com/commonground/nlx/nlx/commit/c4dbfbc1f1f4c463fd912cb111c17a14ab8aea35))
* **management:** show non-editable fields in edit order form ([25f4925](https://gitlab.com/commonground/nlx/nlx/commit/25f4925e5dd635fa8db51121bfd4e208f18674ae))
* **txlog-db:** update migrate version in txlog-db image ([4ec4b96](https://gitlab.com/commonground/nlx/nlx/commit/4ec4b96ebbb671dfa554a7888f18893361ea5c6f))


### Features

* **ca-certportal:** use existing serial number when csr contains the serial number ([3ffea7d](https://gitlab.com/commonground/nlx/nlx/commit/3ffea7d1bc50caa5bcf60ee6205f240e64ca3df9))
* **directory:** add participant data to the directory ui ([227559f](https://gitlab.com/commonground/nlx/nlx/commit/227559ff71384a09199398288bb5c10a528c282f))
* **directory:** register outway ([14c21ef](https://gitlab.com/commonground/nlx/nlx/commit/14c21eff6637eaf0c13ddcd56afff00fa55ffcce))
* **directory:** register outway ([6ee4289](https://gitlab.com/commonground/nlx/nlx/commit/6ee4289f6ad1a73591e1c1adb6a7881e1e356a03))
* **management:** add organization email to settings ([6e4bffd](https://gitlab.com/commonground/nlx/nlx/commit/6e4bffde6a131088781b7562d58d93f69a5292dc))
* **management:** allow changing existing orders in management API ([fd7baf9](https://gitlab.com/commonground/nlx/nlx/commit/fd7baf925e2204c2392465edbb131c73b4d8baaf))
* **management:** make it possible to edit orders in ui ([e575f1f](https://gitlab.com/commonground/nlx/nlx/commit/e575f1f114f6f3b72d32c8469ad2dd0189cb18f9))
* **management:** show order update in audit log ([41c4132](https://gitlab.com/commonground/nlx/nlx/commit/41c4132f6e53b3bd082b37bd94b607a281e3c7e0))

# [0.119.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.118.1...v0.119.0) (2021-11-24)


### Bug Fixes

* **directory:** add correct service names in the grpc files for backward compatibility ([fbc7672](https://gitlab.com/commonground/nlx/nlx/commit/fbc7672ca14046c0b4df52b6818caabe09a9bae0))
* **inway:** correctly set inway as ready ([340055f](https://gitlab.com/commonground/nlx/nlx/commit/340055f225b15b8a77adb9897dc85f1604bc6f38))
* **management:** eslint issues from new eslint config ([08851ac](https://gitlab.com/commonground/nlx/nlx/commit/08851ac5437ed334e6902f300c54697f273eafa5))
* **management:** update delete inway popup copy ([99a5979](https://gitlab.com/commonground/nlx/nlx/commit/99a597935f7a0d81e992415a613f2643b24bd789))
* **management:** use stam+t for third person singular verb in Dutch ([9f876d6](https://gitlab.com/commonground/nlx/nlx/commit/9f876d6efd090b98892d829ee9b67d43b6c64650))


### Features

* **management:** add inway delete to auditlog ([880f977](https://gitlab.com/commonground/nlx/nlx/commit/880f977f2f43bd95d8f49de08c07579c43ed1b0f))
* **management:** add remove inway button to management ui ([4533fc5](https://gitlab.com/commonground/nlx/nlx/commit/4533fc5f307a8bf7b7f9ce65ac1ebd4ea59b55a8))
* **management:** add remove inway button to management ui ([f965d2e](https://gitlab.com/commonground/nlx/nlx/commit/f965d2eaebcf0a291c6f89116afb3f5bb2566b86))
* **management:** implement deletion of inway in management UI and handle inway reregister ([8d170c3](https://gitlab.com/commonground/nlx/nlx/commit/8d170c3db5bc2b95e2b13a9b672164b7d1977aac))
* **management:** implement deletion of inway in management UI and handle inway reregister ([9557e7e](https://gitlab.com/commonground/nlx/nlx/commit/9557e7ec13ed9f1872e23ab8615132a7b4e0fec0))
* **management:** show clearly which services are offered by own organization ([f3f6546](https://gitlab.com/commonground/nlx/nlx/commit/f3f654695942b977c909c8e4d422ab1796a18a84))

## [0.118.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.118.0...v0.118.1) (2021-11-09)


### Bug Fixes

* **management:** if an access request already exists it is reused ([571913c](https://gitlab.com/commonground/nlx/nlx/commit/571913cacb7eed7a09b37f0a6e3065539f12bda5))

# [0.118.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.117.1...v0.118.0) (2021-11-08)


### Bug Fixes

* **management:** remove broken specification button ([f68bbea](https://gitlab.com/commonground/nlx/nlx/commit/f68bbead8498a9fa2f87182e61aea11c07661743))


### Features

* **directory:** open specifiation links in external redoc website ([970a523](https://gitlab.com/commonground/nlx/nlx/commit/970a52320b51bca2e4399efce93295edb38895ae))
* **management:** add open specification button for service details ([435f1d3](https://gitlab.com/commonground/nlx/nlx/commit/435f1d3a3bd2e22b4654af86caf23dfb8a0f89af))

## [0.117.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.117.0...v0.117.1) (2021-10-31)


### Bug Fixes

* **management:** update public key pem in database when registering an outway ([e0ef0e0](https://gitlab.com/commonground/nlx/nlx/commit/e0ef0e04c98ccb1a49478863dcb41417d5809daf))
* **outway:** send organization certificate public key to management api ([d7bcaee](https://gitlab.com/commonground/nlx/nlx/commit/d7bcaee61e044729bbf4e8df400bc8126271610e))

# [0.117.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.116.0...v0.117.0) (2021-10-31)


### Bug Fixes

* **directory:** add idle connection timeout to the directory monitor ([d9c72db](https://gitlab.com/commonground/nlx/nlx/commit/d9c72dba11b61f556066b6ef3be19afc1e63cebe))
* **directory:** add missing organization struct ([a67e06c](https://gitlab.com/commonground/nlx/nlx/commit/a67e06c6b826d28efc2e0ddcc0c6553fcab645a3)), closes [#1270](https://gitlab.com/commonground/nlx/nlx/issues/1270)
* **directory:** enable filtering by serial number ([5b417a9](https://gitlab.com/commonground/nlx/nlx/commit/5b417a9131ce3977833a549c41da1b7c994fb094)), closes [#1270](https://gitlab.com/commonground/nlx/nlx/issues/1270)
* **docs:** regenerate packe-lock.json ([09e106a](https://gitlab.com/commonground/nlx/nlx/commit/09e106ac6efbc2899b90624b44a70f33fd322a76))
* **helm:** update Outway endpoints to use organization serial number instead of organization name ([896f0e2](https://gitlab.com/commonground/nlx/nlx/commit/896f0e2184aaafa765a6cc7e28772e259624b061)), closes [#1376](https://gitlab.com/commonground/nlx/nlx/issues/1376)
* **helm:** update Outway endpoints to use organization serial number instead of organization name ([83dd545](https://gitlab.com/commonground/nlx/nlx/commit/83dd54590e6a7d6e6723f2981479b8ff5a58c557)), closes [#1376](https://gitlab.com/commonground/nlx/nlx/issues/1376)
* **inway:** use organization serial number for authentication ([ae0ac80](https://gitlab.com/commonground/nlx/nlx/commit/ae0ac8059335c3bc1e5677600645269b8cd8dc5a))
* **management:** ensure height of PrimaryNavigation is always at least 100vh ([07bb39a](https://gitlab.com/commonground/nlx/nlx/commit/07bb39a571203563a7fba047179fff96b23c45c7)), closes [#1376](https://gitlab.com/commonground/nlx/nlx/issues/1376)
* **management:** fix formatting of new lines when displaying Public Key PEM of Outway ([fe42fab](https://gitlab.com/commonground/nlx/nlx/commit/fe42fab9341f6b61302729183fd1e04171cdaf3f)), closes [#917](https://gitlab.com/commonground/nlx/nlx/issues/917)
* **management:** fix incorrect translation ([f68849d](https://gitlab.com/commonground/nlx/nlx/commit/f68849d8c2179ffe7909db9704f31f70ceb7aa14)), closes [#917](https://gitlab.com/commonground/nlx/nlx/issues/917)
* **management:** linting & tests ([312ed0f](https://gitlab.com/commonground/nlx/nlx/commit/312ed0f8beb379bfb2a2976d9b84f67356316d74))
* **management:** remove existing audit logs ([48790d5](https://gitlab.com/commonground/nlx/nlx/commit/48790d58599604bd7569ddd7cf2a8dc6331a3b63)), closes [#1376](https://gitlab.com/commonground/nlx/nlx/issues/1376)
* **management:** remove existing incoming/outgoing access requests ([706f511](https://gitlab.com/commonground/nlx/nlx/commit/706f5110591ea73c81573ee58e3133d5477ccb07)), closes [#1376](https://gitlab.com/commonground/nlx/nlx/issues/1376)
* **management:** rework incorrect migrations ([cf2e7eb](https://gitlab.com/commonground/nlx/nlx/commit/cf2e7eb856ea77aa19cbbe2ea57b48ea6e7b77ea))
* **management:** update create access request audit log ([a3131ca](https://gitlab.com/commonground/nlx/nlx/commit/a3131ca941a3f0f90585112a70e982ce026475cd)), closes [#1376](https://gitlab.com/commonground/nlx/nlx/issues/1376)
* **management:** update reject access request audit log ([65265ba](https://gitlab.com/commonground/nlx/nlx/commit/65265bacb1d94ab80176656ed2f2c034e56e3306)), closes [#1376](https://gitlab.com/commonground/nlx/nlx/issues/1376)
* **management:** use correct organization name and serial number in audit log service record ([7198964](https://gitlab.com/commonground/nlx/nlx/commit/7198964b984a5ab1c73bd80be42d0f2e7314db76))
* **management:** use correct plural keys in translation files ([fae7408](https://gitlab.com/commonground/nlx/nlx/commit/fae7408155da946674d4065ea3f291cbc2ac39db))
* **management:** use max length of 100 for the outway name in the DB ([f966156](https://gitlab.com/commonground/nlx/nlx/commit/f966156f2b5a0e3a346c195edc73415576c11f20)), closes [#917](https://gitlab.com/commonground/nlx/nlx/issues/917)
* **management:** use organization serial number to get inway proxy address ([a6c5b04](https://gitlab.com/commonground/nlx/nlx/commit/a6c5b04fbcf8bd13b2d312161bc54b09fb52b17b))
* **management:** use varchar instead of char for directory organization serial numbers ([f441a3d](https://gitlab.com/commonground/nlx/nlx/commit/f441a3d8776ba1d169cf6b3c1adddfe6d0657f4c)), closes [#1376](https://gitlab.com/commonground/nlx/nlx/issues/1376)


### Features

* **directory:** add serial number to directory inspection api ([62aefde](https://gitlab.com/commonground/nlx/nlx/commit/62aefde08cbcb3885341fdf1e0e299911e71d2e4))
* **directory:** connect ui to inspection api ([719fdbe](https://gitlab.com/commonground/nlx/nlx/commit/719fdbe9f26d9be6a66f3b875c0569414bd08efd))
* **directory:** expanded logging capability of the directory monitor ([7621e5e](https://gitlab.com/commonground/nlx/nlx/commit/7621e5ea2438ce00de2b538c11baaa3a9e84bf5f))
* **directory:** use serial number in inspection api ([431187b](https://gitlab.com/commonground/nlx/nlx/commit/431187b2838f0d74a7516c8e4f63965653996b0e))
* **management:** add filter buttons to the Inways page ([21860b5](https://gitlab.com/commonground/nlx/nlx/commit/21860b550544bbe4b6b2ad0f1678cdbde6c13cc3)), closes [#917](https://gitlab.com/commonground/nlx/nlx/issues/917)
* **management:** add icons for Inways and Outway tabs ([8f9c6d6](https://gitlab.com/commonground/nlx/nlx/commit/8f9c6d69f80875b3a69f8bf7dff28ab7529913a5)), closes [#917](https://gitlab.com/commonground/nlx/nlx/issues/917)
* **management:** add organization serial number when listing incoming access requests ([80b55cf](https://gitlab.com/commonground/nlx/nlx/commit/80b55cf99be4330ecbb5cec4e40518106bf7aaaa)), closes [#1270](https://gitlab.com/commonground/nlx/nlx/issues/1270)
* **management:** add organizationSerialNumber for the ListServicesResponse ([9dfcb2d](https://gitlab.com/commonground/nlx/nlx/commit/9dfcb2d8ea20c015e8a13ed074f814f6b9d1c9a4)), closes [#1270](https://gitlab.com/commonground/nlx/nlx/issues/1270)
* **management:** add outway registering itself to management api and listing outways ([a88f38d](https://gitlab.com/commonground/nlx/nlx/commit/a88f38d8839aaa2f9b5a2889f3a1a7b5db98dc3c))
* **management:** add serial number support to access grants ([ba2e0d5](https://gitlab.com/commonground/nlx/nlx/commit/ba2e0d527b7af3a453ee7bb3a3ca2101f1a61a6c))
* **management:** generate new client for ui from api protobuf ([db2095b](https://gitlab.com/commonground/nlx/nlx/commit/db2095b1998024693cfa3f964dda1f3f64ed30ae))
* **management:** introduce Outway detail page ([d2ffca9](https://gitlab.com/commonground/nlx/nlx/commit/d2ffca9f8700ac856ce514e64caa7f7c544b6f47)), closes [#917](https://gitlab.com/commonground/nlx/nlx/issues/917)
* **management:** introduce Outway model + store ([ddb994c](https://gitlab.com/commonground/nlx/nlx/commit/ddb994c30b2819d6810f55a49bb4a44d9401348d)), closes [#917](https://gitlab.com/commonground/nlx/nlx/issues/917)
* **management:** introduce subpage /inways-and-outways/inways to display inways ([67fe881](https://gitlab.com/commonground/nlx/nlx/commit/67fe88114e9ac2d3c6dfdec45e1e7e0b7a965f14)), closes [#917](https://gitlab.com/commonground/nlx/nlx/issues/917)
* **management:** introduce the Outways component ([52c90e3](https://gitlab.com/commonground/nlx/nlx/commit/52c90e335d1a03cb26b66a1eebacc99e0f5c23eb)), closes [#917](https://gitlab.com/commonground/nlx/nlx/issues/917)
* **management:** rename /inways route to /inways-and-outways ([e3bb10c](https://gitlab.com/commonground/nlx/nlx/commit/e3bb10c70bd368ce2a8046934a1e3f0486870b8b)), closes [#917](https://gitlab.com/commonground/nlx/nlx/issues/917)
* **management:** rename access request method name ([bf9b102](https://gitlab.com/commonground/nlx/nlx/commit/bf9b10261b0ab297706ef7185d563b457296730d)), closes [#1270](https://gitlab.com/commonground/nlx/nlx/issues/1270)
* **management:** rework revoke Access Grant to work with serial number ([ac8b1e9](https://gitlab.com/commonground/nlx/nlx/commit/ac8b1e981a9bbfc5f678a235610f2f7b3b1d9011)), closes [#1370](https://gitlab.com/commonground/nlx/nlx/issues/1370)
* **management:** show outways in outway tab list ([5790ca6](https://gitlab.com/commonground/nlx/nlx/commit/5790ca649489f5e4da94dfaf433fd8757d4fc420))
* **management:** style the Public Key PEM information of an Outway ([8eae9fa](https://gitlab.com/commonground/nlx/nlx/commit/8eae9fa67e955e0db0aecbff705414a1817bb939)), closes [#917](https://gitlab.com/commonground/nlx/nlx/issues/917)
* **management:** update Inway to pass the certificate serial number when proxying to Management ([6bf00b6](https://gitlab.com/commonground/nlx/nlx/commit/6bf00b616a8388eaad4ae417e3cf920e52199799)), closes [#1370](https://gitlab.com/commonground/nlx/nlx/issues/1370)
* **management:** update label for the delegatee of the Add order form ([b870d70](https://gitlab.com/commonground/nlx/nlx/commit/b870d70f3204eb5214babb860535a08dea5ee847)), closes [#1376](https://gitlab.com/commonground/nlx/nlx/issues/1376)
* **management:** update SendAccessRequest to only require the outgoing access request ID ([73c3bb5](https://gitlab.com/commonground/nlx/nlx/commit/73c3bb55a2cccc16a6a0056a8168f362809852a8)), closes [#1376](https://gitlab.com/commonground/nlx/nlx/issues/1376)
* **management:** use organization serial number for GetLatestAccessGrantForService ([1637645](https://gitlab.com/commonground/nlx/nlx/commit/16376459a73743401ceb7d72cf017eb1ad2d1d24)), closes [#1376](https://gitlab.com/commonground/nlx/nlx/issues/1376)
* **management:** use organization serial number for orders ([8f31ce7](https://gitlab.com/commonground/nlx/nlx/commit/8f31ce79aeb8e6c8b579511514c05ce61290f183))
* **management:** use organization serial number instead of name for CreateAccessRequest ([bcd5524](https://gitlab.com/commonground/nlx/nlx/commit/bcd5524a09d5388dc980240e82aa412272c79146)), closes [#1376](https://gitlab.com/commonground/nlx/nlx/issues/1376)
* **management:** use organization serialnumber for get latest incoming access requests ([3be8a7e](https://gitlab.com/commonground/nlx/nlx/commit/3be8a7e17394d78da8e9d2397ff6334830982c65))
* **management:** use serial number in service request ([11a2606](https://gitlab.com/commonground/nlx/nlx/commit/11a2606598596fc0b0907502363ba2f839f3d451))
* **management:** use varchar(20) for `organization_serial_number` instead of char(20) ([8349c5f](https://gitlab.com/commonground/nlx/nlx/commit/8349c5fd6edb94649a9743aaeb38225672309c74)), closes [#1270](https://gitlab.com/commonground/nlx/nlx/issues/1270)
* **outway:** use serial number instead of organization name ([20e65e6](https://gitlab.com/commonground/nlx/nlx/commit/20e65e6a69889647452b226d22aaaab68a459872))
* register organization inway in directory via inways instead of management api ([f0ac6ce](https://gitlab.com/commonground/nlx/nlx/commit/f0ac6cea123fedcde253eaadeebc8e88756b47c0)), closes [#1362](https://gitlab.com/commonground/nlx/nlx/issues/1362)
* use subject serialnumber in the directory registration API ([fa1162c](https://gitlab.com/commonground/nlx/nlx/commit/fa1162cda141d1cd1aa5690a2714ad4a36516cc7))

# [0.116.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.115.0...v0.116.0) (2021-10-12)


### Bug Fixes

* **docs:** resolve broken link to Try NLX ([df0f75c](https://gitlab.com/commonground/nlx/nlx/commit/df0f75c5c358b02c280a2f3ecd3ae2d453938a5d))


### Features

* register organization inway in directory via inways instead of management api ([c905d77](https://gitlab.com/commonground/nlx/nlx/commit/c905d77aafe2dba76e263902dee26aefe7269cfb))


### Reverts

* register organization inway in directory via inways instead of management api ([b6159aa](https://gitlab.com/commonground/nlx/nlx/commit/b6159aa1b451aac3c295d54595eab5bd5c123d36))

# [0.115.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.114.1...v0.115.0) (2021-09-21)


### Bug Fixes

* add listen to `syscall.SIGQUIT` to enable graceful shutdown ([61bf651](https://gitlab.com/commonground/nlx/nlx/commit/61bf65158f514354fedb26db61d65875526d64fb)), closes [#1349](https://gitlab.com/commonground/nlx/nlx/issues/1349)
* **inway:** add missing signal to enable graceful shutdown ([37b4ad1](https://gitlab.com/commonground/nlx/nlx/commit/37b4ad1dec14a487a2a210d036973e93ba10093e)), closes [#1349](https://gitlab.com/commonground/nlx/nlx/issues/1349)
* wait for sigterm in go routine instead of defer ([da88f78](https://gitlab.com/commonground/nlx/nlx/commit/da88f78e40183194aa50b1bbdfd0875741597bd8))


### Features

* **ca-certportal:** add subject serialnumber when generating a certificate ([ddaafc2](https://gitlab.com/commonground/nlx/nlx/commit/ddaafc25f58ee10bc2c6da3b6154442d47010ccd)), closes [#1364](https://gitlab.com/commonground/nlx/nlx/issues/1364)
* **directory:** display the organization serial number when present ([e8f975a](https://gitlab.com/commonground/nlx/nlx/commit/e8f975a484acca9773120533f900cfc4386379e5)), closes [#1356](https://gitlab.com/commonground/nlx/nlx/issues/1356)
* **helm:** improve error message when creating a user which already exists ([96a0931](https://gitlab.com/commonground/nlx/nlx/commit/96a0931e8eace43614e8e8867fa3610142867d3e))
* **management:** display the organization serial number when present ([88513d4](https://gitlab.com/commonground/nlx/nlx/commit/88513d434a32d38f5012bdbaf90661c523e76902)), closes [#1366](https://gitlab.com/commonground/nlx/nlx/issues/1366)

## [0.114.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.114.0...v0.114.1) (2021-09-16)

# [0.114.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.113.0...v0.114.0) (2021-09-16)


### Bug Fixes

* **helm:** oidc sessionSignKey not required with enableBasicAuth ([0d174d5](https://gitlab.com/commonground/nlx/nlx/commit/0d174d5b359f98b00f6e2c1de98777f887c99615))
* **management:** properly synchronize incoming orders ([fb8dd6c](https://gitlab.com/commonground/nlx/nlx/commit/fb8dd6c5383d250ad7d0ab2ae27675573712e3d7))
* **management:** use error.is instead of == when checking errors ([450071a](https://gitlab.com/commonground/nlx/nlx/commit/450071a9560bae078f9c3ec61f3f33d618e1b909))


### Features

* **management:** add tests for service deletion endpoint ([ef0c465](https://gitlab.com/commonground/nlx/nlx/commit/ef0c465e8f808e9f04c66789059fa4ee74ee5ad6))
* **management:** delete service with active requests and/or orders ([c988c9f](https://gitlab.com/commonground/nlx/nlx/commit/c988c9ff829ffe37411edce0ef68b2aeaafedc86))
* **management:** remove outgoing access request when a service is removed ([f82f546](https://gitlab.com/commonground/nlx/nlx/commit/f82f546472a397468f242ec762458a53981009dd))
* **management:** remove outgoing access request when a service is removed ([3bce171](https://gitlab.com/commonground/nlx/nlx/commit/3bce171f105cfa62fc15902f55944a1f5ba955b7))
* **management:** remove service, access requests and order service entries from db ([01fb250](https://gitlab.com/commonground/nlx/nlx/commit/01fb250507e0ead7c74e7088ce5c5aee3b4c09f3))
* **management:** scheduler removes outgoing accessrequests when a service is deleted ([d7d25c4](https://gitlab.com/commonground/nlx/nlx/commit/d7d25c443877a90640e0d30e02b4121b269b5761))

# [0.113.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.112.0...v0.113.0) (2021-09-07)


### Bug Fixes

* **helm:** convert PostgreSQL connectTimeout value from string to integer ([6022a69](https://gitlab.com/commonground/nlx/nlx/commit/6022a696621bf7636c34d8bb9962721f79b2125e)), closes [#1272](https://gitlab.com/commonground/nlx/nlx/issues/1272)
* **helm:** rename schema.values.json -> values.schema.json ([a1a2c50](https://gitlab.com/commonground/nlx/nlx/commit/a1a2c506a6dcdf3658dfddb13d28492118cb4387)), closes [#1272](https://gitlab.com/commonground/nlx/nlx/issues/1272)


### Features

* **management:** update revoke button color to red ([545c4dc](https://gitlab.com/commonground/nlx/nlx/commit/545c4dc99525a547be69b54abe9624530da4f842)), closes [#1346](https://gitlab.com/commonground/nlx/nlx/issues/1346)

# [0.112.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.111.0...v0.112.0) (2021-09-06)


### Bug Fixes

* **directory:** fix wrong key name ([a0b97f8](https://gitlab.com/commonground/nlx/nlx/commit/a0b97f8fb67551670b7446272f414c6b0d1de109))
* start development script adds user with password for organisation b ([5966356](https://gitlab.com/commonground/nlx/nlx/commit/59663562a2eb688aee5afaba3bfe6da02f9a64db))


### Features

* **directory:** add activity timestamps to inway ([f30b907](https://gitlab.com/commonground/nlx/nlx/commit/f30b9074a481c092560d733cf2af2d1ca3bdd58e))
* **directory:** add migrations for inway activity timestamps ([69087ce](https://gitlab.com/commonground/nlx/nlx/commit/69087ce848003608230ccc9b7cb945ff3b0cb2c3))

# [0.111.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.110.0...v0.111.0) (2021-08-30)


### Bug Fixes

* **directory:** rework E2E test so it works with the new responsive layout ([b12b978](https://gitlab.com/commonground/nlx/nlx/commit/b12b978233ce683b702800fb05abce1a30e770ff))
* **management:** grammar ([ebbacf9](https://gitlab.com/commonground/nlx/nlx/commit/ebbacf91194638d8b40e98bb67bd4cd966661dd3)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** prevent redirecting to /orders when closing an order detail view ([ba54688](https://gitlab.com/commonground/nlx/nlx/commit/ba5468884305d7e85d7b4aa595fa7fd5cf19431c)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** restore content of migration files ([5d3f4b6](https://gitlab.com/commonground/nlx/nlx/commit/5d3f4b672440429983b57d8a6fd015aec015553e)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** return correct auditlog actiontypes ([8f82265](https://gitlab.com/commonground/nlx/nlx/commit/8f822656a689a60db3d38f7407d772df3dc3e180))


### Features

* **directory:** implement environment select component ([75c6dbc](https://gitlab.com/commonground/nlx/nlx/commit/75c6dbc5182c6c986e6e7cb082a8b7ae53f06376))
* **management:** add detail view for incoming orders ([5ee5fdb](https://gitlab.com/commonground/nlx/nlx/commit/5ee5fdb6eacf8e2e8fd62dcc31c72cca75568245)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** add order detail drawer for outgoing orders ([5fa382e](https://gitlab.com/commonground/nlx/nlx/commit/5fa382ebbbf75e943d46dbf373f734b9cf60497d))
* **management:** add status to order drawer ([1bf13d4](https://gitlab.com/commonground/nlx/nlx/commit/1bf13d46542ebf13905c55d6845f2e497bc7c638))
* **management:** add subtitle to outgoing order drawer ([f3725db](https://gitlab.com/commonground/nlx/nlx/commit/f3725dbe851d3d43e6338396bd1241d6df9ddde7))
* **management:** convert order services to links ([d101d9c](https://gitlab.com/commonground/nlx/nlx/commit/d101d9cd71c8f85b3bd4bfe5f7168c3466457c0c)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** display proper error when creating an order with duplicate reference & delegator ([6290c12](https://gitlab.com/commonground/nlx/nlx/commit/6290c129fed71f33473edd83a21efbb23bfa9748)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** display status icon for outgoing orders ([220b29b](https://gitlab.com/commonground/nlx/nlx/commit/220b29b26058d33b6ce49189bb9af8b557ede9d3)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** display valid until and from dates for outgoing orders ([6e3a058](https://gitlab.com/commonground/nlx/nlx/commit/6e3a0584e0eb43ed24f55bee6ec2fb4890e17eb0)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** enable getting an specific outgoing order ([b910e4b](https://gitlab.com/commonground/nlx/nlx/commit/b910e4b2118ce5a739cab02916c3cc02e52c5cdb))
* **management:** enable revoking outgoing order ([3c0fed7](https://gitlab.com/commonground/nlx/nlx/commit/3c0fed76eac027315903ff15587cb3f5dbfd1cce)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** enable storing outgoing orders with same reference but different delegatee ([2155b5a](https://gitlab.com/commonground/nlx/nlx/commit/2155b5ad748bfb042682c2a6da7be4cef1564079))
* **management:** enforce outgoing orders have a unique combination of delegatee and reference ([0175939](https://gitlab.com/commonground/nlx/nlx/commit/0175939e6d42b7fd5358c1e01726221d4cba1b4b)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** implement incoming order auditlog ([bebe102](https://gitlab.com/commonground/nlx/nlx/commit/bebe1021130d2af077ff1f32ea5514039d03efe2))
* **management:** implement Order status icon ([45de81d](https://gitlab.com/commonground/nlx/nlx/commit/45de81d1c3dade82f5b27306a03d13993060dd00)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** make orders in the outgoing orders list clickable ([b9f72aa](https://gitlab.com/commonground/nlx/nlx/commit/b9f72aa11bf3766f8ac7f539a285b88b415ba4b8)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** remove audit log for updating insight settings ([c5d87a3](https://gitlab.com/commonground/nlx/nlx/commit/c5d87a36dfbb9be2eb0addb3e16ef866d6b26272))
* **management:** show correct text for order revoke in auditlog ([40c2d26](https://gitlab.com/commonground/nlx/nlx/commit/40c2d268b0faaf45ad9e5109aaf1ce5aa125d05e))
* **management:** show reference enddate in order drawer ([b8f8aed](https://gitlab.com/commonground/nlx/nlx/commit/b8f8aed8f3c64945dcc55c3656e79f124fae3668))
* **management:** show reference number for outgoing orders ([bf1e37f](https://gitlab.com/commonground/nlx/nlx/commit/bf1e37f282ca5bdde94f59393fd5faaa27a0209b)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** show services for outgoing order ([5749cc3](https://gitlab.com/commonground/nlx/nlx/commit/5749cc36b7fb87aadbd400e4df4b547d4b0256aa)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)
* **management:** update order status when order gets revoked ([7f556aa](https://gitlab.com/commonground/nlx/nlx/commit/7f556aafe3d2dc7b1fac7000389fcc0d9e2f9a92))


### Reverts

* **management:** remove functionality to revoke incoming orders ([05b9771](https://gitlab.com/commonground/nlx/nlx/commit/05b9771a8f399a2088771d8dd68b3f229770d3ca)), closes [#1181](https://gitlab.com/commonground/nlx/nlx/issues/1181)

# [0.110.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.109.2...v0.110.0) (2021-08-18)


### Bug Fixes

* **management:** display service error on form submit ([beb5f15](https://gitlab.com/commonground/nlx/nlx/commit/beb5f154bba10b4a3f0f43d11bd7a6aff7414c12))
* **management:** resolve order form input issues ([1695318](https://gitlab.com/commonground/nlx/nlx/commit/16953185f3d746591998ba003bda6fa1cebb9fcd))
* remove leftovers of service-config.toml ([b60eef8](https://gitlab.com/commonground/nlx/nlx/commit/b60eef82c79008ac8abdfdfe84ab4a2d1a027aa4))
* **helm:** fix postgres secret key names ([3f06a4c](https://gitlab.com/commonground/nlx/nlx/commit/3f06a4cab56bc3eb87d074022339bce0acf8cf20))


### Features

* **helm:** enable configuring basic auth for NLX Management ([4d5c8f0](https://gitlab.com/commonground/nlx/nlx/commit/4d5c8f03e6352930ee81ae742318bd76a375c057)), closes [#1332](https://gitlab.com/commonground/nlx/nlx/issues/1332)
* **management:** implement revoke order in the management-api ([837cae1](https://gitlab.com/commonground/nlx/nlx/commit/837cae152ec2b702e2bab93bac2ec03da47d3584))
* **management:** warn when adding order without accessible services ([7ae6076](https://gitlab.com/commonground/nlx/nlx/commit/7ae6076626b34f07670a5977798dd4eb5c08b2cb))

## [0.109.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.109.1...v0.109.2) (2021-08-10)


### Bug Fixes

* **management:** do not require secret key to be set when using basic auth ([90138d5](https://gitlab.com/commonground/nlx/nlx/commit/90138d5f8ca43d06162e8a6294f154df3f268d2a)), closes [#1330](https://gitlab.com/commonground/nlx/nlx/issues/1330)
* **management:** re-enable detection for basic auth ([80398ef](https://gitlab.com/commonground/nlx/nlx/commit/80398ef7a3fb9ab28616e1db43cc3f18d6050da1)), closes [#1330](https://gitlab.com/commonground/nlx/nlx/issues/1330)

## [0.109.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.109.0...v0.109.1) (2021-08-04)


### Bug Fixes

* **helm:** removed trailing dashes from if statements in ingresses ([8676b80](https://gitlab.com/commonground/nlx/nlx/commit/8676b807e2ee8cab02d521d20a7237ab463d219f))

# [0.109.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.108.0...v0.109.0) (2021-08-03)


### Bug Fixes

* **docs:** resolve breaking change for docusaurus ([fe14644](https://gitlab.com/commonground/nlx/nlx/commit/fe14644cbe74a8bc5c06ca07c2fe33aacf32ae99))
* **management:** remove restricting url validation ([a6dc056](https://gitlab.com/commonground/nlx/nlx/commit/a6dc0568d78511fd0f2e47d4292c2ca20dd3c0fc))


### Features

* **helm:** add support for both the networking.k8s.io/v1beta1 and networking.k8s.io/v1 API ([e8900dd](https://gitlab.com/commonground/nlx/nlx/commit/e8900dd9e5b4ea182bc06a022d63faceabab3469))
* **helm:** reinstate semverCompare ([150ce96](https://gitlab.com/commonground/nlx/nlx/commit/150ce9613a459f33bfa9791899a1504295056937))
* **helm:** update deprecated version ([f902e7e](https://gitlab.com/commonground/nlx/nlx/commit/f902e7e9b83b0c8efa317ab2ac8d11d160a5cd89))
* **inway:** add annotations and loadbalancerip ([a32767b](https://gitlab.com/commonground/nlx/nlx/commit/a32767be3a6b6d806411f49174ee0b3c419771f8))
* **management:** implement support for basic authentication ([19c59ab](https://gitlab.com/commonground/nlx/nlx/commit/19c59ab079d7249478c0ba4982af0159f992d3c1))

# [0.108.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.107.2...v0.108.0) (2021-07-13)


### Bug Fixes

* **management:** added missing distiction between incoming and outgoing orders ([38c65eb](https://gitlab.com/commonground/nlx/nlx/commit/38c65eba6640673858dd42317ce4f0abaeb382fa))
* **management:** update color of the degraded service icon ([2d33704](https://gitlab.com/commonground/nlx/nlx/commit/2d337049b236aa4d82aa066b8943f63a4653fc7f))


### Features

* **management:** display warning if outgoing orders are present but no organization inway is set ([f88e3be](https://gitlab.com/commonground/nlx/nlx/commit/f88e3be2bf100e90973f3723fc1f90e993cd3bd9)), closes [#1304](https://gitlab.com/commonground/nlx/nlx/issues/1304)
* **management:** update description of the organization inway to include syncing orders ([070fea2](https://gitlab.com/commonground/nlx/nlx/commit/070fea2d98a73ad1c9291cd42aaccf98b928098b)), closes [#1304](https://gitlab.com/commonground/nlx/nlx/issues/1304)

## [0.107.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.107.1...v0.107.2) (2021-07-08)

# [0.107.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.106.0...v0.107.0) (2021-07-05)


### Bug Fixes

* revert easter eggs ([ee8cc68](https://gitlab.com/commonground/nlx/nlx/commit/ee8cc68d6f929c637a5f976ad6352b0522023217))


### Features

* **directory:** enable inways to register without a name ([f6fa28a](https://gitlab.com/commonground/nlx/nlx/commit/f6fa28a3fbad50fb17691625a6edd52e78e27c56))
* move rebuild script from acc to demo ([8e76fa6](https://gitlab.com/commonground/nlx/nlx/commit/8e76fa6c3dfe9dc35187fc9383112608cfffcfb9)), closes [#1270](https://gitlab.com/commonground/nlx/nlx/issues/1270)
* move stijns api & viewer from acc to demo ([782a3b9](https://gitlab.com/commonground/nlx/nlx/commit/782a3b9ef99a617fa1f1c8047b00b070a72adc36)), closes [#1270](https://gitlab.com/commonground/nlx/nlx/issues/1270)
* move vergunningsoftware from acc to demo ([48737db](https://gitlab.com/commonground/nlx/nlx/commit/48737dbc831c37e900366777989e06763070588f)), closes [#1270](https://gitlab.com/commonground/nlx/nlx/issues/1270)

# [0.106.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.105.0...v0.106.0) (2021-07-02)


### Bug Fixes

* **management:** rename invalid issued order message ([b43f514](https://gitlab.com/commonground/nlx/nlx/commit/b43f514f707793dce5fd78b18eefd9d7a04110f6))
* **management:** translate missing translations English -> Dutch ([e19944c](https://gitlab.com/commonground/nlx/nlx/commit/e19944c3088ba36be9cac3c54ad34d90433cb19d)), closes [#1281](https://gitlab.com/commonground/nlx/nlx/issues/1281)
* actually show goodbye present ([75ef4bb](https://gitlab.com/commonground/nlx/nlx/commit/75ef4bbdb30c541ad1e2af60aece3cf876f51aad))
* add license statement to order-synchronize ([373cda9](https://gitlab.com/commonground/nlx/nlx/commit/373cda91e68e51d318bf99fc158deff81869f23e))
* make sure upsert only updates or inserts and not accidentally both ([91e2529](https://gitlab.com/commonground/nlx/nlx/commit/91e252924d42a57a15661bf980931068db605665))
* rename incoming order fields ([3c7806b](https://gitlab.com/commonground/nlx/nlx/commit/3c7806b2e46a1fce68361a29a6c0416df929957c))
* rename Order to OutgoingOrder ([b7bed73](https://gitlab.com/commonground/nlx/nlx/commit/b7bed738d474c22c086dc82e64717c8ecfaafa48))
* skip order in synchronization if it hasn't been changed ([db97ecc](https://gitlab.com/commonground/nlx/nlx/commit/db97ecc0fdf484a7276bc3d49c61bdc6ca87b14a))
* title for button with only icon ([9e804b1](https://gitlab.com/commonground/nlx/nlx/commit/9e804b10b092d3b04c9553b6c175a83e7d409ecb)), closes [#1166](https://gitlab.com/commonground/nlx/nlx/issues/1166)
* **docs:** links ([0775c76](https://gitlab.com/commonground/nlx/nlx/commit/0775c76c740189e3368197404aa73bcf770d114e))
* **docs:** minor text changes ([7528d0e](https://gitlab.com/commonground/nlx/nlx/commit/7528d0ef8f57f46fc4771ff007feb36aeaa5e4a7))
* **management:** service validation ([4f8fa2e](https://gitlab.com/commonground/nlx/nlx/commit/4f8fa2e5a38646c888949d8b85c5bfa301d185dc))


### Features

* **ca-certportal:** improve logging when failed to decode CSR ([abd231a](https://gitlab.com/commonground/nlx/nlx/commit/abd231ac65298d38c39e04ed594026601c677a26)), closes [#1249](https://gitlab.com/commonground/nlx/nlx/issues/1249)
* **ca-certportal:** log signing errors ([4b24df1](https://gitlab.com/commonground/nlx/nlx/commit/4b24df1aef747cdaca43227428e391869f8f67ca)), closes [#1249](https://gitlab.com/commonground/nlx/nlx/issues/1249)
* **ca-certportal:** return proper feedback message when failing to parse CSR ([f499c42](https://gitlab.com/commonground/nlx/nlx/commit/f499c4299aab7e5f2a95ee16094417ba95df0a87)), closes [#1249](https://gitlab.com/commonground/nlx/nlx/issues/1249)
* **directory:** extend error log for registering inway ([0561134](https://gitlab.com/commonground/nlx/nlx/commit/056113444e6447e0c930f763107aeafe7a6a02dd)), closes [#1285](https://gitlab.com/commonground/nlx/nlx/issues/1285)
* **directory:** extend error log for registering inway ([471a8ad](https://gitlab.com/commonground/nlx/nlx/commit/471a8ad63fc3e31a7dfa13b2f35923f968085944)), closes [#1285](https://gitlab.com/commonground/nlx/nlx/issues/1285)
* **directory:** remove existing inways & availabilities ([f3a5994](https://gitlab.com/commonground/nlx/nlx/commit/f3a59944d5126632ef75d8d55cac2d1b78612609)), closes [#1285](https://gitlab.com/commonground/nlx/nlx/issues/1285)
* **directory:** require name to be provided when registering an inway ([8bf740d](https://gitlab.com/commonground/nlx/nlx/commit/8bf740d4f42735c255247a6a1fbcac27ba1fb404)), closes [#1285](https://gitlab.com/commonground/nlx/nlx/issues/1285)
* **directory:** show clear error when attempting to register inway with a duplicate address ([67c4a4d](https://gitlab.com/commonground/nlx/nlx/commit/67c4a4df0f546f26b0e3f389d28eb2fb2c36a93f)), closes [#1285](https://gitlab.com/commonground/nlx/nlx/issues/1285)
* **directory:** update inway name when registering an inway ([8ac236a](https://gitlab.com/commonground/nlx/nlx/commit/8ac236ac3d02c906de1128dcff3d27c59c596d4d)), closes [#1285](https://gitlab.com/commonground/nlx/nlx/issues/1285)
* **management:** received order overview ([b622dc5](https://gitlab.com/commonground/nlx/nlx/commit/b622dc55306fd1522db6408f17cc10a0ae4d7c3a))
* nav humorability ([95a81de](https://gitlab.com/commonground/nlx/nlx/commit/95a81de59ed44b31515ffcedc04c8260eaecc730))
* synchronize orders ([c60dee3](https://gitlab.com/commonground/nlx/nlx/commit/c60dee3a6f719968a82960433a04b38b73945ee7))
* **inway:** provide inway name when announcing inway to the directory ([abf1404](https://gitlab.com/commonground/nlx/nlx/commit/abf140484af9db67a25453c48445a41c6d2dad0f)), closes [#1285](https://gitlab.com/commonground/nlx/nlx/issues/1285)
* **inway:** require a valid name on startup ([0defe5d](https://gitlab.com/commonground/nlx/nlx/commit/0defe5d682784414f4531268a4c4cd941eeaab6a)), closes [#1285](https://gitlab.com/commonground/nlx/nlx/issues/1285)
* **management:** add endpoint to list orders for my organization ([163394f](https://gitlab.com/commonground/nlx/nlx/commit/163394f003bd69b7b211a57a9c384846bf42e398))

# [0.105.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.104.0...v0.105.0) (2021-06-16)


### Bug Fixes

* **helm:** update outdated kubedb secret for parkeerrechten api ([03c9cec](https://gitlab.com/commonground/nlx/nlx/commit/03c9cec321e2db794311b051395e72cc79ffe504))


### Features

* replace kubedb with zalando operator ([60244be](https://gitlab.com/commonground/nlx/nlx/commit/60244bec03b55885bd39a8b45325fa9aaa50aec9))

# [0.104.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.103.1...v0.104.0) (2021-06-15)


### Bug Fixes

* **ca-certportal:** prevent CSP errors for inline JS ([6596b92](https://gitlab.com/commonground/nlx/nlx/commit/6596b925d365661a6e7dee33ee581004f12e18af)), closes [#1279](https://gitlab.com/commonground/nlx/nlx/issues/1279)
* **ca-certportal:** submit Certificate Signing Request correctly ([1c8eac4](https://gitlab.com/commonground/nlx/nlx/commit/1c8eac4cc8b1627e502f4bff065aec0f57180fd2)), closes [#1249](https://gitlab.com/commonground/nlx/nlx/issues/1249)


### Features

* **directory:** enable clearing the organization inway for an organization that is not yet present ([bcf3f53](https://gitlab.com/commonground/nlx/nlx/commit/bcf3f53332afff6dec89e1d346b641154e748638)), closes [#1200](https://gitlab.com/commonground/nlx/nlx/issues/1200)
* **directory:** enable registering inway even when the amount of services exceeds the limit of 250 ([c314a20](https://gitlab.com/commonground/nlx/nlx/commit/c314a20d071a4343fbeb547248094853a276d3d6)), closes [#1200](https://gitlab.com/commonground/nlx/nlx/issues/1200)
* **directory:** enable registering inways which do not have a service ([5984e44](https://gitlab.com/commonground/nlx/nlx/commit/5984e4479a100c71000d4bb49e1702605e3442b0)), closes [#1200](https://gitlab.com/commonground/nlx/nlx/issues/1200)
* **helm:** configure the nodeport when the inway service is of type nodeport ([630cae5](https://gitlab.com/commonground/nlx/nlx/commit/630cae5ec4d4375f09482164745c5d6c3ce030e3))

## [0.103.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.103.0...v0.103.1) (2021-06-11)

# [0.103.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.102.4...v0.103.0) (2021-06-09)


### Bug Fixes

* **management:** prefer scheduling outgoing access requests updated least recently ([63ae989](https://gitlab.com/commonground/nlx/nlx/commit/63ae9895b7f9ba0617abbf54fbc3c0b271f30f11)), closes [#1249](https://gitlab.com/commonground/nlx/nlx/issues/1249)
* update helm chart repository in chart readme ([45fa1be](https://gitlab.com/commonground/nlx/nlx/commit/45fa1be8820e1fdb3de7f14fd7d9052c57526340))


### Features

* **ca-certportal:** replace font 'Muli' with system font ([8a80c9d](https://gitlab.com/commonground/nlx/nlx/commit/8a80c9dade88e3153c72af113755c6bbe0c87cf2))

## [0.102.4](https://gitlab.com/commonground/nlx/nlx/compare/v0.102.3...v0.102.4) (2021-06-03)

## [0.102.3](https://gitlab.com/commonground/nlx/nlx/compare/v0.102.2...v0.102.3) (2021-06-02)

## [0.102.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.102.1...v0.102.2) (2021-06-02)


### Bug Fixes

* add stable repo before updating stijns helm dependency ([1cc79cc](https://gitlab.com/commonground/nlx/nlx/commit/1cc79cc5404bedfe9ca8ea9f60f76707b57529e1))

## [0.102.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.102.0...v0.102.1) (2021-06-02)


### Bug Fixes

* install helm from alpine testing repository ([5409900](https://gitlab.com/commonground/nlx/nlx/commit/540990065535adb125b52eea71adb329aa9381ac))

# [0.102.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.101.1...v0.102.0) (2021-06-01)


### Bug Fixes

* fix linting issues after updating prettier to 2.3.0 ([4db6dfc](https://gitlab.com/commonground/nlx/nlx/commit/4db6dfc141daca8aa858cd713d8b16c4504c1f10))


### Features

* **docs:** upgrade Docusaurus v2 alpha to beta + remove custom template ([9e9bc9b](https://gitlab.com/commonground/nlx/nlx/commit/9e9bc9bb7af6a34388dabf1aecea886b7040f2ce)), closes [#1171](https://gitlab.com/commonground/nlx/nlx/issues/1171)


### Reverts

* remove demo application 'Parkeervergunningen' ([1f5a9e1](https://gitlab.com/commonground/nlx/nlx/commit/1f5a9e1ff0a45f2129dc5812eb52484d6217780f)), closes [#1240](https://gitlab.com/commonground/nlx/nlx/issues/1240)

## [0.101.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.101.0...v0.101.1) (2021-05-11)

# [0.101.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.100.2...v0.101.0) (2021-05-05)


### Features

* set static IP address for NLX Directory Inspection API in pre-prod/prod ([8a386de](https://gitlab.com/commonground/nlx/nlx/commit/8a386de4b0d2a18ffb794d6dbfa562158187aad7))

## [0.100.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.100.1...v0.100.2) (2021-05-03)

## [0.100.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.100.0...v0.100.1) (2021-05-01)

# [0.100.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.99.0...v0.100.0) (2021-04-30)


### Bug Fixes

* **management:** update service name for directory E2E tests ([f74ea0c](https://gitlab.com/commonground/nlx/nlx/commit/f74ea0c9d37446b2d1008c43ea853e2a4ea8c7f0)), closes [#1215](https://gitlab.com/commonground/nlx/nlx/issues/1215)


### Features

* **management:** rename haarlem -> gemeente-stijns ([8a742af](https://gitlab.com/commonground/nlx/nlx/commit/8a742af1ef2430a646781d48201cfd35b4c47a23))
* **management:** rename Saas org X -> Vergunningsoftware BV ([7924c63](https://gitlab.com/commonground/nlx/nlx/commit/7924c6343e7a9c1e0c4e24a944aaefb1e17593ed)), closes [#1215](https://gitlab.com/commonground/nlx/nlx/issues/1215)

# [0.99.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.98.0...v0.99.0) (2021-04-28)


### Bug Fixes

* **management:** audit log mobx warning ([2e8cee6](https://gitlab.com/commonground/nlx/nlx/commit/2e8cee6e8fa41c4e059b8dec802e90d3d48c37a0))
* **management:** dayjs locale loading order ([bf18b6c](https://gitlab.com/commonground/nlx/nlx/commit/bf18b6c4afd178e97cf955c6c8eb9d7b465cfbf8))


### Features

* **management:** add description to the organization field when adding an order ([c6a9ddc](https://gitlab.com/commonground/nlx/nlx/commit/c6a9ddcd7b1325387f216a6de091742b735a2463))
* **management:** add order overview page ([4f772da](https://gitlab.com/commonground/nlx/nlx/commit/4f772dacdcb1c9e806545c073ba857e5dfba40ba)), closes [#1182](https://gitlab.com/commonground/nlx/nlx/issues/1182)
* **management:** implement ListOutgoingOrders API endpoint ([09e943f](https://gitlab.com/commonground/nlx/nlx/commit/09e943fc3f537ba1829e7df0a70d3caeceb04c21)), closes [#1182](https://gitlab.com/commonground/nlx/nlx/issues/1182)
* **management:** introduce ListOutgoingOrders API endpoint ([ce67a9b](https://gitlab.com/commonground/nlx/nlx/commit/ce67a9ba728b69051d4f71ff5be52ba062b9d4a4)), closes [#1182](https://gitlab.com/commonground/nlx/nlx/issues/1182)
* **management:** sort issued orders descending on valid until property ([2985564](https://gitlab.com/commonground/nlx/nlx/commit/29855641b2102bc844c8f9e44af305217fed78ff)), closes [#1182](https://gitlab.com/commonground/nlx/nlx/issues/1182)

# [0.98.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.97.1...v0.98.0) (2021-04-19)


### Bug Fixes

* **directory:** correctly retrieve stats from database ([524d5cb](https://gitlab.com/commonground/nlx/nlx/commit/524d5cbbb0cb84c4e26399c6bccc66b1ef5688be)), closes [#1213](https://gitlab.com/commonground/nlx/nlx/issues/1213)
* **directory:** enable /stats endpoint ([2386fa7](https://gitlab.com/commonground/nlx/nlx/commit/2386fa77c0c75d2aceca2ce4dc83f11523d48450))


### Features

* move support/contact page to nlx.io ([5bff2ca](https://gitlab.com/commonground/nlx/nlx/commit/5bff2ca78e931cf3caa881fa11583c93af2c7df8))

## [0.97.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.97.0...v0.97.1) (2021-04-15)


### Bug Fixes

* **helm:** remove invalid override of managementAPI ([b1719a3](https://gitlab.com/commonground/nlx/nlx/commit/b1719a3d453cfc749fa167127312ed92548fad5d))

# [0.97.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.96.0...v0.97.0) (2021-04-15)


### Bug Fixes

* **directory:** remove tools file in favor of Earthly ([cfbec32](https://gitlab.com/commonground/nlx/nlx/commit/cfbec32d136e7f25646318c66a544a7dc17916b8))
* **directory:** remove unused Makefile (replaced by Earthly) ([ad16c7d](https://gitlab.com/commonground/nlx/nlx/commit/ad16c7db8a58cbac8412687b8f23682feb5e2ba4))
* **directory:** restore property name casing for backwards compatibility ([b897fd1](https://gitlab.com/commonground/nlx/nlx/commit/b897fd15dd8977acf28744e84ad08d2a835da711))
* **directory:** stats endpoint path ([6dabed3](https://gitlab.com/commonground/nlx/nlx/commit/6dabed39744b7dcb1a6fa6cdbff7c516b96af4c2))
* **docs:** update docs to use Earthly for protobuf compilation ([d834460](https://gitlab.com/commonground/nlx/nlx/commit/d834460bdfbc017cebf71c05238bea6a92555a3d))
* **inway:** fix missing directory inspection API in inway ([e501905](https://gitlab.com/commonground/nlx/nlx/commit/e50190529411f6b10c1a781a52f31ec8fdcf1e40))
* **inway:** set delegator properly ([f8b9f0b](https://gitlab.com/commonground/nlx/nlx/commit/f8b9f0b59c06ce5150a35352c2a5fb3178a9d338))
* **inway:** store the actual organization properly when using delegation ([0fd9fb3](https://gitlab.com/commonground/nlx/nlx/commit/0fd9fb32cdeba36cd08ce60aaa1c42f613cff7ce))
* **inway:** use management client instead of the removed delegation client ([17dc447](https://gitlab.com/commonground/nlx/nlx/commit/17dc447bede9aff5be9701bbc45ab200c8296575))
* **inway:** validate service and organization name in claims ([ce47f77](https://gitlab.com/commonground/nlx/nlx/commit/ce47f7733a004161bf7bc9264b078ca09ac49647))
* **management:** correctly create related services for audit logs ([4330026](https://gitlab.com/commonground/nlx/nlx/commit/4330026fb5dbb5b1d71db1a12a2153af6fd99f73)), closes [#1180](https://gitlab.com/commonground/nlx/nlx/issues/1180)
* **management:** create auditlog without services when no service was specified ([1a0b0a3](https://gitlab.com/commonground/nlx/nlx/commit/1a0b0a3ce56d6fca1e9b7ff8082b441db33327d5))
* **management:** don't put access request in failed state when polling fails ([438e2d3](https://gitlab.com/commonground/nlx/nlx/commit/438e2d3939c46b7bc825b122abca1c2282be72a8))
* **management:** error message should include correct expected instance ([7000a95](https://gitlab.com/commonground/nlx/nlx/commit/7000a9504857d48212414a63669f102373e9d1a6)), closes [#1180](https://gitlab.com/commonground/nlx/nlx/issues/1180)
* **management:** fix imports and type conversions as much as possible ([4a6e3d2](https://gitlab.com/commonground/nlx/nlx/commit/4a6e3d24e8fa8fd91bc700ea2c52d1192fc870e5))
* **management:** fix linting errors ([69ffe35](https://gitlab.com/commonground/nlx/nlx/commit/69ffe357285d11ecf2a7c31e9128f9ff2f0423db))
* **management:** fix UI of audit logs ([3efb045](https://gitlab.com/commonground/nlx/nlx/commit/3efb0453cd42c570618b8792217b3342b616cb6f)), closes [#1180](https://gitlab.com/commonground/nlx/nlx/issues/1180)
* **management:** implemented correct method for RetrieveClaimForOrder ([f6975ec](https://gitlab.com/commonground/nlx/nlx/commit/f6975ec5e5a3477be3c0650eb2afb1653e6e2218))
* **management:** initialize gRPC client services in management.NewClient ([3d741d4](https://gitlab.com/commonground/nlx/nlx/commit/3d741d450e2fc816c5fae65783f4ffc04bfd6310))
* **management:** migrations ([63790f5](https://gitlab.com/commonground/nlx/nlx/commit/63790f513d2c0007d85c415813c3318ab8f5ee2a)), closes [#1180](https://gitlab.com/commonground/nlx/nlx/issues/1180)
* **management:** nav icon spacing ([dd6a1e5](https://gitlab.com/commonground/nlx/nlx/commit/dd6a1e55962cccad4ae595bff0bdeab3bc89c67c))
* **management:** register delegation service server in the management API ([d682d00](https://gitlab.com/commonground/nlx/nlx/commit/d682d005a35b109a5a55eec2dfb2f9f6481369e3))
* **management:** remove unused proxy metadata in retrieve claims ([d8cfdd6](https://gitlab.com/commonground/nlx/nlx/commit/d8cfdd6cfb873fd85703826299ee00c36f5a0471))
* **management:** rename AuditLogRecord to AuditLog ([07f2f26](https://gitlab.com/commonground/nlx/nlx/commit/07f2f2660866f99b95e5d0fc73c236ad353a72fc)), closes [#1180](https://gitlab.com/commonground/nlx/nlx/issues/1180)
* **management:** set timeout to 4 minutes for all jobs ([2709462](https://gitlab.com/commonground/nlx/nlx/commit/270946226c7ed696bb37b5391d705d8b6ef87449))
* **management:** translation ([f3bc9c8](https://gitlab.com/commonground/nlx/nlx/commit/f3bc9c84b3cb8f0f53775f489ccd62a3f6e440b4))
* add missing translations for order creation ([82f9686](https://gitlab.com/commonground/nlx/nlx/commit/82f96864b99ee15cad52f9001392258401dc222f))
* disable go-modules for outdated projects ([1020905](https://gitlab.com/commonground/nlx/nlx/commit/102090577e5b63e43b8ed0b55c60a8e90af60b81))
* use PKIX format for RSA public keys instead of PKCS1 ([88163a9](https://gitlab.com/commonground/nlx/nlx/commit/88163a93a943be511b9d03fc6659ffd6b349cee4))
* **management:** integrate with UI ([eca98f5](https://gitlab.com/commonground/nlx/nlx/commit/eca98f5f53b8fa07e12918627940a009ce5c45d5)), closes [#1180](https://gitlab.com/commonground/nlx/nlx/issues/1180)
* hotfix cypress ([599c407](https://gitlab.com/commonground/nlx/nlx/commit/599c40769c6f2a0f1411bfb22ce9dd22d167ab14))
* **management:** rename incorrect property name ([069fac7](https://gitlab.com/commonground/nlx/nlx/commit/069fac7475ba536bec044c90b9f3a7ea0344949e)), closes [#1180](https://gitlab.com/commonground/nlx/nlx/issues/1180)
* remove skaffold comment from ignore files as we no longer use it ([50b3222](https://gitlab.com/commonground/nlx/nlx/commit/50b32226217732ba46c945dc4b7f8a184fa1a8b2))
* remove storybook file ([ad76a0d](https://gitlab.com/commonground/nlx/nlx/commit/ad76a0df5e4d09f776c59cca6f428caae4b40beb))
* resolve linting issues for Golang code ([1e34601](https://gitlab.com/commonground/nlx/nlx/commit/1e34601cde92338973cb9cde9b4b3697421f54ed))
* start to replace Makefile ([f54606d](https://gitlab.com/commonground/nlx/nlx/commit/f54606dac3c09585bc523380da0ba81c22b8cae2))
* support contact page ([fa46845](https://gitlab.com/commonground/nlx/nlx/commit/fa46845e24c2c08a3eb8760b02b5f0ec57d9bf7d))
* test typing service cost values ([e07f2b5](https://gitlab.com/commonground/nlx/nlx/commit/e07f2b599a17cbaa758b4d9ac2ff0f029d51ef86))
* use html code for apostrophe ([69174cc](https://gitlab.com/commonground/nlx/nlx/commit/69174cc877fbbe7d70d81124269117fd93810a65))
* use the dex command to serve the app as required in the new version ([95f9b3d](https://gitlab.com/commonground/nlx/nlx/commit/95f9b3d72e1ff4fb4217187a2c7b3763dab306f5))
* **management:** temp not used strings ([229e6bf](https://gitlab.com/commonground/nlx/nlx/commit/229e6bf042362c6956968a1bf92b2185885cc72d))
* **outway:** fix several issues with log-data not being properly set in delegation plugin ([01fc068](https://gitlab.com/commonground/nlx/nlx/commit/01fc0689d2506db3bb128ec795341082b3afc1c7))
* **outway:** mitigate race condition in handleOnNlx tests ([1172c13](https://gitlab.com/commonground/nlx/nlx/commit/1172c13aa5ded77a9afb5979bf44aff91c569cb1))


### Features

* **directory:** add stats service to the Directory Inspection API ([34a2177](https://gitlab.com/commonground/nlx/nlx/commit/34a2177aa209facb5de6dc2782cfa63aa3137d3c))
* **helm:** add nlx management to brp deployment ([bd5ac98](https://gitlab.com/commonground/nlx/nlx/commit/bd5ac98cf3d7ddf35223f37cba720df8e9be5df1))
* **helm:** add nlx management to brp deployment ([d27125a](https://gitlab.com/commonground/nlx/nlx/commit/d27125a2fb8c4ca908fecd10e0797f810c3720d6))
* **inway:** refactor inway and remove config.toml ([c00ad6a](https://gitlab.com/commonground/nlx/nlx/commit/c00ad6ab87e665d5765b20a2ca01fe73ad0b1f23))
* **inway:** validate delegatee in JWT claims when using delegation ([27d0ba0](https://gitlab.com/commonground/nlx/nlx/commit/27d0ba004db87f26d217ba8496c1f3a6e2d0edac))
* **inway:** verify claim if it is present ([84bd528](https://gitlab.com/commonground/nlx/nlx/commit/84bd528b1fb0c769f0c16104691e25f09ac07802)), closes [#1179](https://gitlab.com/commonground/nlx/nlx/issues/1179)
* **management:** add CreateOutgoingOrder endpoint ([4ae2532](https://gitlab.com/commonground/nlx/nlx/commit/4ae25323c3cf9becc927e7db0b28f4985d2ede8e)), closes [#1180](https://gitlab.com/commonground/nlx/nlx/issues/1180)
* **management:** add public key pem to ListServices response ([f5603ba](https://gitlab.com/commonground/nlx/nlx/commit/f5603bad9fda4e1ae9187e754476c5f73de3827f))
* **management:** component DateInput ([2b989b2](https://gitlab.com/commonground/nlx/nlx/commit/2b989b2eb8ef2e9c82f1ab5231ccb2ee0740ec81))
* **management:** implement design feedback ([4645f50](https://gitlab.com/commonground/nlx/nlx/commit/4645f502f462537b4df8ef185bb16a648ba175e5)), closes [#1180](https://gitlab.com/commonground/nlx/nlx/issues/1180)
* **management:** implement input validation for CreateOutgoingOrder endpoint ([66c6586](https://gitlab.com/commonground/nlx/nlx/commit/66c6586ab7128781222b21a097a8e5971d7e17da)), closes [#1180](https://gitlab.com/commonground/nlx/nlx/issues/1180)
* **management:** implement input validation for CreateOutgoingOrder endpoint + add database migrations ([b902b65](https://gitlab.com/commonground/nlx/nlx/commit/b902b65fcf95cbb25007f54aab3526fcadfd5403)), closes [#1180](https://gitlab.com/commonground/nlx/nlx/issues/1180)
* **management:** implement public key pinning and other checks for request claim ([ad84704](https://gitlab.com/commonground/nlx/nlx/commit/ad847042b819d2023c6b96c6292b5d41a736ae0a))
* **management:** preload services when retrieving auditlogs ([542933f](https://gitlab.com/commonground/nlx/nlx/commit/542933f6c1bf5ccd84d8e80124e32590727a5af2))
* add inway to haarlem deploy chart ([725f342](https://gitlab.com/commonground/nlx/nlx/commit/725f34255eb9cfc9d4f146b7ffd54afa677144ba))
* added auditlogs on order creation and added services and organizations ([8bd7dcf](https://gitlab.com/commonground/nlx/nlx/commit/8bd7dcf636311c22423c770e5479db29c9c9fba5))
* create claims based on orders and validate services within a claim ([ffd51e5](https://gitlab.com/commonground/nlx/nlx/commit/ffd51e5dd49daaeb3372b2ebe6057e261e4afe2e))
* **management:** implement storing orders in the database ([821bcbd](https://gitlab.com/commonground/nlx/nlx/commit/821bcbdcbd87e5f9ba37d3554020ffdd0c19c6e4))
* **management:** implement VerifyClaim endpoint ([185386c](https://gitlab.com/commonground/nlx/nlx/commit/185386c0c584dc8d09c111b8273568b248e0b26c)), closes [#1179](https://gitlab.com/commonground/nlx/nlx/issues/1179)
* **management:** introduce CreateOutgoingOrderPage ([ac839e4](https://gitlab.com/commonground/nlx/nlx/commit/ac839e44f94c418c8d9a76b2a448b7b41b151918)), closes [#1180](https://gitlab.com/commonground/nlx/nlx/issues/1180)
* enable transaction logging in SaaS Party X helm chart ([849bcea](https://gitlab.com/commonground/nlx/nlx/commit/849bcea3931063f50fe02f7585ed2c0bb18bcada))
* **directory:** remove items from navbar for new nlx.io ([8d8ee2f](https://gitlab.com/commonground/nlx/nlx/commit/8d8ee2fe0ab353f1a720e633fa3f4597780ce161))
* **docs:** update navbar for new nlx.io ([fa6c960](https://gitlab.com/commonground/nlx/nlx/commit/fa6c960d01862ad15749e82e0ca51631c15f564a))
* **management:** add endpoint to Request the claim for an order ([32c59d9](https://gitlab.com/commonground/nlx/nlx/commit/32c59d9cbb0347d8813dbd3ebb4125a3bbd3bd31)), closes [#1179](https://gitlab.com/commonground/nlx/nlx/issues/1179)
* **management:** implement RetrieveClaim endpoint ([35a1e36](https://gitlab.com/commonground/nlx/nlx/commit/35a1e360b5964217868763f8a0cb57c0a40f5834)), closes [#1179](https://gitlab.com/commonground/nlx/nlx/issues/1179)
* **management:** use FieldLabel for inway select ([bd2d65a](https://gitlab.com/commonground/nlx/nlx/commit/bd2d65a0c279edbe0cd7f0251b1f8a342d6d6e6e))
* **outway:** assert delegation plugin errors in unittests ([18d056c](https://gitlab.com/commonground/nlx/nlx/commit/18d056c9b49906ca6e49b5b191481488c1ede36a))
* **outway:** implement a new plugin-system and add delegation ([a0c0517](https://gitlab.com/commonground/nlx/nlx/commit/a0c0517b41c0fd0563027ece9e3ff7eba0efecb5))
* **outway:** retrieve and validate claims in the outway when delegation is used ([fa7f4fe](https://gitlab.com/commonground/nlx/nlx/commit/fa7f4fe8d60caa0b9a7a8588ff105f5b1f66cf31))
* **outway:** return more detailed errors when retrieving claim fails ([00137d0](https://gitlab.com/commonground/nlx/nlx/commit/00137d06ce035e01e1ef9f13bb3d4579f7b90c1d))
* add support for outway with management API for development ([42b511e](https://gitlab.com/commonground/nlx/nlx/commit/42b511e9165507f4486ccf8d9723a281bfd25495))
* set management API address in outway ([4d3faaa](https://gitlab.com/commonground/nlx/nlx/commit/4d3faaa25c6b5b2c7731332f041bda43cb250b9a))
* store the full public key in access-requests ([4a609f8](https://gitlab.com/commonground/nlx/nlx/commit/4a609f8d068f00aa75e74e8e4973342d0b69228e))

# [0.96.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.95.0...v0.96.0) (2021-03-03)


### Bug Fixes

* **docs:** update docusaurus to 2.0.0-alpha.70 ([a559856](https://gitlab.com/commonground/nlx/nlx/commit/a559856f1a98f5466a2d551a8fa7957fa330f4d1))
* **management:** month is now exported correctly to finance export csv ([55074c9](https://gitlab.com/commonground/nlx/nlx/commit/55074c9caa5e559d4172212658f521b5d733d5d0))
* **management:** set inway properly including creation ([31bf4b7](https://gitlab.com/commonground/nlx/nlx/commit/31bf4b7979a1d51fe47167d7a05b601f368f4c6f))
* include new security templates ([f27de30](https://gitlab.com/commonground/nlx/nlx/commit/f27de30e92e41998d0392471146fe3a79cc6f138))
* partially mitigate CVE-2020-28477 ([48040da](https://gitlab.com/commonground/nlx/nlx/commit/48040da94e876e77b6afbb2a3693b6a6862d0c68))
* switch jwt-go library because the former one wasn't actively developed anymore ([e728044](https://gitlab.com/commonground/nlx/nlx/commit/e72804487c00e4f7ef65c85e3640b3ec88d44713))
* update cobra to the latest version ([4095621](https://gitlab.com/commonground/nlx/nlx/commit/4095621a6b80b2b1175e8445fce62af4e8aeb26d))
* **directory:** correctly format go code ([18d92e5](https://gitlab.com/commonground/nlx/nlx/commit/18d92e59447faf253aefbc7e486a0d7296beb9bb)), closes [#1158](https://gitlab.com/commonground/nlx/nlx/issues/1158)
* **directory:** typo request costs column name ([27bca47](https://gitlab.com/commonground/nlx/nlx/commit/27bca478d98284847e1586cdefe9821bd1f03b1c)), closes [#1158](https://gitlab.com/commonground/nlx/nlx/issues/1158)
* **management:** add license ([b2aae1d](https://gitlab.com/commonground/nlx/nlx/commit/b2aae1d5657b68bc233d9e99d07fac3be6334c44))
* **management:** fix tests after rebase ([5a076e1](https://gitlab.com/commonground/nlx/nlx/commit/5a076e1134483daa539370d296c75bd8ba37bfbf)), closes [#1158](https://gitlab.com/commonground/nlx/nlx/issues/1158)
* **management:** re-structure finance is enabled unittest ([40f673e](https://gitlab.com/commonground/nlx/nlx/commit/40f673efdd42d6a1254c6ab307bce2c309088448))
* **management:** remove duplicated test-case ([07e2225](https://gitlab.com/commonground/nlx/nlx/commit/07e22256de30403350681ade281e16cf12c97e5f))
* **management:** rename billing to finance ([682682e](https://gitlab.com/commonground/nlx/nlx/commit/682682e97949968d03cc677ca5e87957c5f53228))
* **management:** translation file ([01f6763](https://gitlab.com/commonground/nlx/nlx/commit/01f6763090061c34a3dbd094dc6e510770c58104))
* **management:** translation file ([ebd8090](https://gitlab.com/commonground/nlx/nlx/commit/ebd8090fa829d605bb2d77d6322bb5c71da2dbbe))
* **management:** update CostsSection to display Free if the costs are undefined ([77b164c](https://gitlab.com/commonground/nlx/nlx/commit/77b164c2ccd354bf991e39ef55846ac9126faac7)), closes [#1158](https://gitlab.com/commonground/nlx/nlx/issues/1158)
* **management:** use organization name in finance export instead of common-name ([4ea45b4](https://gitlab.com/commonground/nlx/nlx/commit/4ea45b4c1c6a8b4e4e0c7d6b3510fde60191bf07))


### Features

* add transaction log support to helm chart ([f9ef5c1](https://gitlab.com/commonground/nlx/nlx/commit/f9ef5c16e12b09fc4b59df5cef20fa2725953176))
* **directory:** add pricing ([9b6b95c](https://gitlab.com/commonground/nlx/nlx/commit/9b6b95c986cb50c22c4906b907a505eb258f01ed))
* **directory:** add service costs when registering an inway ([be66ade](https://gitlab.com/commonground/nlx/nlx/commit/be66ade1e2667ca51756e3d6492f3e5a1d832f2b)), closes [#1158](https://gitlab.com/commonground/nlx/nlx/issues/1158)
* **directory:** expose service costs in ListServices endpoint ([465bcaa](https://gitlab.com/commonground/nlx/nlx/commit/465bcaac346f87359fcd7a898f15650d1e333a60)), closes [#1158](https://gitlab.com/commonground/nlx/nlx/issues/1158)
* **docs:** document billing ([241d4bf](https://gitlab.com/commonground/nlx/nlx/commit/241d4bf3c68cdc7a3dc0489b9f17de07d233a497))
* **inway:** pass service costs when registering an inway ([a29dfd0](https://gitlab.com/commonground/nlx/nlx/commit/a29dfd0aab4366fbbdeddd59e9b113bdec385d88)), closes [#1158](https://gitlab.com/commonground/nlx/nlx/issues/1158)
* **management:** add unittests for finance export ([9d170d7](https://gitlab.com/commonground/nlx/nlx/commit/9d170d78d901979a2541c649317a218ff1a6ae48))
* **management:** configure finance page ([5695086](https://gitlab.com/commonground/nlx/nlx/commit/569508686487a0194824de89ceeb0d983b01c038))
* **management:** detect if finance is enabled ([b38a623](https://gitlab.com/commonground/nlx/nlx/commit/b38a6232a70eb4a530ec6adb9a0dd5d94dbf162b))
* **management:** display costs before requesting access to a service ([f6258da](https://gitlab.com/commonground/nlx/nlx/commit/f6258da5cb871be6fcda18eb750c6252bf4d06ec)), closes [#1158](https://gitlab.com/commonground/nlx/nlx/issues/1158)
* **management:** display costs for Directory services ([0d31b39](https://gitlab.com/commonground/nlx/nlx/commit/0d31b39b0613bddc813c1b0c13c45bf59074d528)), closes [#1158](https://gitlab.com/commonground/nlx/nlx/issues/1158)
* **management:** display service costs for directory ([b8df0a9](https://gitlab.com/commonground/nlx/nlx/commit/b8df0a902a7cc26cbd474033812d489f822b732a)), closes [#1158](https://gitlab.com/commonground/nlx/nlx/issues/1158)
* **management:** display service costs for directory ([a81fd89](https://gitlab.com/commonground/nlx/nlx/commit/a81fd8928a5e2e611169e8135cb646d856ced08e)), closes [#1158](https://gitlab.com/commonground/nlx/nlx/issues/1158)
* **management:** export billing overview based on transaction logs ([606d544](https://gitlab.com/commonground/nlx/nlx/commit/606d54436b291bedf06926c4965a79d75c182e02))
* **management:** expose service costs in ListServices endpoint ([a703430](https://gitlab.com/commonground/nlx/nlx/commit/a703430309dda5e3d0693e6da6ef3a0cd3a4bc37)), closes [#1158](https://gitlab.com/commonground/nlx/nlx/issues/1158)
* **management:** finance page UI ([6b9cee8](https://gitlab.com/commonground/nlx/nlx/commit/6b9cee8d7aee91b3c1eb142522bcfddde9322c22))

# [0.95.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.94.2...v0.95.0) (2021-02-24)


### Bug Fixes

* **docs:** explain how to create admin user ([e323982](https://gitlab.com/commonground/nlx/nlx/commit/e323982bec5d56270360023c739b5f141357633a))
* **management:** add missing translations ([e2161b9](https://gitlab.com/commonground/nlx/nlx/commit/e2161b9c8ff2ba972f88d6e03491ef87df031a1a)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** bug in polling ([4db3e14](https://gitlab.com/commonground/nlx/nlx/commit/4db3e14ea8c2e40968f833c08c2f4a3400d1d38e))
* **management:** check if service exists ([d7f9036](https://gitlab.com/commonground/nlx/nlx/commit/d7f9036e7324cc69910630186a720202bb5ed288))
* **management:** directory Switch no default ([d18dfea](https://gitlab.com/commonground/nlx/nlx/commit/d18dfeae65421e9bae0c445f7b4052aa9d57a1e2))
* **management:** fix color for Logout icon ([64cfc89](https://gitlab.com/commonground/nlx/nlx/commit/64cfc89333a51890d47540f0fce557256146d8fd)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** fix linting error ([953ab90](https://gitlab.com/commonground/nlx/nlx/commit/953ab907553bbacf8f55ab951b95513b86404c66)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** fix typo in ActionType name ([d404207](https://gitlab.com/commonground/nlx/nlx/commit/d4042076b8b44ccd2955f4dee79ea6ce037092ee)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** generic text for polling updates ([5e44edf](https://gitlab.com/commonground/nlx/nlx/commit/5e44edff964f0479e93d8c3712d073fdb6f9c518)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** mock css transitions in modal ([78972d7](https://gitlab.com/commonground/nlx/nlx/commit/78972d7888e1de6b1c039d3036875d2b2bfc87c6))
* **management:** name of management API URL argument ([7e96823](https://gitlab.com/commonground/nlx/nlx/commit/7e9682375538443397ccc262234b0b5a4122e201))
* **management:** only emit error when loading key ([f37c1fb](https://gitlab.com/commonground/nlx/nlx/commit/f37c1fb76e2561470b10d8266bec1cffe2b20853))
* **management:** prevent Collapsible from rerendering when fetching incoming access grants ([b1991cc](https://gitlab.com/commonground/nlx/nlx/commit/b1991cc4c8ee4fa68a12b6bfa5701d32d9727707)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** prevent Collapsible from rerendering when fetching incoming access requests ([6afd3c0](https://gitlab.com/commonground/nlx/nlx/commit/6afd3c0d4e672ead8fae706e3c6e6405e1caa9c8)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** regenerate migrations and api client from proto ([11c9828](https://gitlab.com/commonground/nlx/nlx/commit/11c9828733672df34c058695654fab4fefefe3c8)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** rename and regenerate migrations after rebase ([bc21bcf](https://gitlab.com/commonground/nlx/nlx/commit/bc21bcf85083a675191a3d2b35e8796922f1a010)), closes [#1157](https://gitlab.com/commonground/nlx/nlx/issues/1157)
* **management:** rename facturatie to kosten in ServiceForm ([5aee086](https://gitlab.com/commonground/nlx/nlx/commit/5aee086964bb7b9a998c536090e449e1d0f32225))
* **management:** resolve most linting errors ([931d445](https://gitlab.com/commonground/nlx/nlx/commit/931d44560cc345261e310df656ed7ab71cb49cb8)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** several unit test errors & warnings ([041fcf5](https://gitlab.com/commonground/nlx/nlx/commit/041fcf579bfb3ea9a62663a57c360a1fbf576a23))
* **management:** sync AuditLog action types of client with API ([0a06fa0](https://gitlab.com/commonground/nlx/nlx/commit/0a06fa0ee9336bb4ab8113e7aefeb8ce3bcba152)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** test for general settings form ([3bf6ce8](https://gitlab.com/commonground/nlx/nlx/commit/3bf6ce8eac956498985d88dd7ef3564ee3778058))
* **management:** update audit log translations ([ff058f3](https://gitlab.com/commonground/nlx/nlx/commit/ff058f3b0a63b8b9547c61cf4e95469c137e8eac)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** update incomingAccessRequestsCount when opening service details ([94cb3da](https://gitlab.com/commonground/nlx/nlx/commit/94cb3dabd49be07817323a5324be41ac7ce32d18)), closes [#1133](https://gitlab.com/commonground/nlx/nlx/issues/1133)
* **management:** update settings icon to have dynamic size ([b837bc1](https://gitlab.com/commonground/nlx/nlx/commit/b837bc1a0aa932a3584deeb889e552fba8fb340b)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** update tests for AuditLogStore ([f8218b5](https://gitlab.com/commonground/nlx/nlx/commit/f8218b5cf5c12d2a174bc357f810316ae3746770)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** update translation files ([cfe0df5](https://gitlab.com/commonground/nlx/nlx/commit/cfe0df505bdb9c1a55d22f75a97886d04f6f49d5)), closes [#1157](https://gitlab.com/commonground/nlx/nlx/issues/1157)
* **management:** update translations for audit log date ([64a2535](https://gitlab.com/commonground/nlx/nlx/commit/64a2535990d675505324b236afeed63c0b39faee)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* add pre-commit config to gitignore ([4f932d5](https://gitlab.com/commonground/nlx/nlx/commit/4f932d5ac8a6d3f81758b5b227fed916fa2d4c95))
* disable secure cookies locally ([38fbe5d](https://gitlab.com/commonground/nlx/nlx/commit/38fbe5d92a51470272dc9084371f4e82b4829371))
* explain what managementAPI.address means in the Helm charts ([1ff7471](https://gitlab.com/commonground/nlx/nlx/commit/1ff7471986c9b46cc1474a705898a67e0eb3ffac))
* make sure inway can be used without an existing TLS secret but with NLX Management ([b80ae6e](https://gitlab.com/commonground/nlx/nlx/commit/b80ae6edd4f7bac0348b4b69c5b423bea2be74c0))
* rename API property in service ([2d45a04](https://gitlab.com/commonground/nlx/nlx/commit/2d45a0442a9488fe1c9c4cd008549aee8b64c765))
* use Helm v3 syntax in readme and renamed my-NAME to NAME ([8134fe8](https://gitlab.com/commonground/nlx/nlx/commit/8134fe8b5d7af7b36fa9cfb011ed27c70cfb574f))
* **outway:** panic on --help ([2794d28](https://gitlab.com/commonground/nlx/nlx/commit/2794d284212a7ccbd5fbe9b72a224e68e0f30e6a))


### Features

* **helm:** disable management-api by default ([b3562ec](https://gitlab.com/commonground/nlx/nlx/commit/b3562ec03102a4f0442a9d9c1cfafbf043d0ae20))
* **helm:** enable management-api by default ([2916faa](https://gitlab.com/commonground/nlx/nlx/commit/2916faa2781b8fe4c5e9345c8cf6520fdaec3abf))
* **helm:** remove comments from valyes.yaml ([2cd5913](https://gitlab.com/commonground/nlx/nlx/commit/2cd5913eda492b33927c1e4b82283a38300c6ed1))
* **management:** add audit log ([e1cdce9](https://gitlab.com/commonground/nlx/nlx/commit/e1cdce9c51e4af009dda7fdaa49b6dee001ed05a))
* **management:** add audit log as interface ([90f64ef](https://gitlab.com/commonground/nlx/nlx/commit/90f64efbabe85cd2629bd10ad18178cc6408d446))
* **management:** add AuditLog record for accepting an incoming access request ([daf6a8c](https://gitlab.com/commonground/nlx/nlx/commit/daf6a8c2606b3208c563e5bf29eaeaf78de1f80f))
* **management:** add AuditLog record for login fail ([6216727](https://gitlab.com/commonground/nlx/nlx/commit/62167273c4ed77a30934b83bfb71335e4d1fad15))
* **management:** add AuditLog record for logout success ([32042cd](https://gitlab.com/commonground/nlx/nlx/commit/32042cd225d18da3cbef031d29d987375c506664))
* **management:** add AuditLog record for rejecting an incoming access request ([4b94c52](https://gitlab.com/commonground/nlx/nlx/commit/4b94c52d792e6afaae14b16981846dcaa2cf1e1e))
* **management:** add basic version of AuditLog page ([8900f11](https://gitlab.com/commonground/nlx/nlx/commit/8900f11b0d4bd378b6e7d518e34b04aee996a2a3)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** add error messages for costs ([d6b560b](https://gitlab.com/commonground/nlx/nlx/commit/d6b560bc2ffc022a04ba68eb8e8218ee1815f49c))
* **management:** add link to Audit log page in primary navigation ([f6effe6](https://gitlab.com/commonground/nlx/nlx/commit/f6effe62c2d494542a6dc8ffa4ef5e58b9bf75ff)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** add service name to AuditLog records ([0af9c2a](https://gitlab.com/commonground/nlx/nlx/commit/0af9c2ae5f922d7ad29398be36f0a3cf82f552b8))
* **management:** add ServiceForm tests and clear costs when billing is disabled ([afee812](https://gitlab.com/commonground/nlx/nlx/commit/afee812a866028c9d6100570c1e877c1546c1e93))
* **management:** add sorting by createdAt date for audit logs ([92cb803](https://gitlab.com/commonground/nlx/nlx/commit/92cb803255f754e98329058855c17182f40ec5f1)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** add support for AuditLog request access failed ([c20e1f1](https://gitlab.com/commonground/nlx/nlx/commit/c20e1f15165ec4cfa0f219bef7a29aa2415cf36b)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** allow setting path to static web files ([febe86d](https://gitlab.com/commonground/nlx/nlx/commit/febe86d2ab97da5f405ef7f83325d56ad4a6998d))
* **management:** create AuditLog when creating a service ([164e0b6](https://gitlab.com/commonground/nlx/nlx/commit/164e0b69eda18eb03a777146bc3378c259c7a270))
* **management:** create AuditLog when deleting a service ([04273c8](https://gitlab.com/commonground/nlx/nlx/commit/04273c8d4365b911452394899f5aa85a7e87516b))
* **management:** create AuditLog when updating a service ([e44c97a](https://gitlab.com/commonground/nlx/nlx/commit/e44c97ab5278faf122247485a0437c2fb1c6d68e))
* **management:** create AuditLog when updating organization settings ([4c666c6](https://gitlab.com/commonground/nlx/nlx/commit/4c666c6c1d50468455d983f3db3ea87c0264fe05)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** display costs for a service on the service overview page ([b273134](https://gitlab.com/commonground/nlx/nlx/commit/b27313482435c938df8dfc8314ada21ff3a64ee5)), closes [#1157](https://gitlab.com/commonground/nlx/nlx/issues/1157)
* **management:** display list of actions from the AuditLogs ([4d666f2](https://gitlab.com/commonground/nlx/nlx/commit/4d666f2a6e57ebc3a19f47e3b6bc02d683c64017)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** display operating system, browser and client for logs ([7d81465](https://gitlab.com/commonground/nlx/nlx/commit/7d8146587ff2f2634d39511238c18db4f8e4b8f9)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** display toast when removing a service fails ([ee5044d](https://gitlab.com/commonground/nlx/nlx/commit/ee5044df59f851900a1889a659927acbe80fbaee))
* **management:** display user agent for audit logs ([c2291b1](https://gitlab.com/commonground/nlx/nlx/commit/c2291b1935bc196c664b4f8b1ce822d96ac323cc)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** empty state should take up as much height as single row ([e9e2d5c](https://gitlab.com/commonground/nlx/nlx/commit/e9e2d5cec689fdfb2c775e5f560e95ab56a12c6c)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** fix AuditLog tests ([2867dd8](https://gitlab.com/commonground/nlx/nlx/commit/2867dd8d7b02178f0b4e412ddc33075bc185d437)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement /audit-logs API endpoint with dummy response data ([d03801b](https://gitlab.com/commonground/nlx/nlx/commit/d03801b556914701f625440a2ff5cab1194f1bae)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement audit logs for login, logout and unknown actions ([95b4c87](https://gitlab.com/commonground/nlx/nlx/commit/95b4c87ca55c476221043c776048ce44b269ecd5)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement AuditLog record for approved incoming access request ([39d9476](https://gitlab.com/commonground/nlx/nlx/commit/39d947682e7713f61a7c137aab4614dd33920681)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement AuditLog record for creating a service ([5c65210](https://gitlab.com/commonground/nlx/nlx/commit/5c65210247cdafd3d418e3eaf350464362c2edd6)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement AuditLog record for failed login attempt ([dccc305](https://gitlab.com/commonground/nlx/nlx/commit/dccc305b60e53d5b18e7f70188b94c4825486590)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement AuditLog record for rejected incoming access request ([910e653](https://gitlab.com/commonground/nlx/nlx/commit/910e653a1558d68e36328db55ea1c711a1dc3f49)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement AuditLog record for removing a service ([15bf1e4](https://gitlab.com/commonground/nlx/nlx/commit/15bf1e4f02dbf7a9e0250561fe6700b492228ca5)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement AuditLog record for requesting access ([783fbe6](https://gitlab.com/commonground/nlx/nlx/commit/783fbe6ae88e7273e795f50399b0c277684f2e1c)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement AuditLog record for revoked access grant ([a5ff361](https://gitlab.com/commonground/nlx/nlx/commit/a5ff3618419eeb46c139127af23980c457e34ade)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement AuditLog record for updating a service ([6c9c44a](https://gitlab.com/commonground/nlx/nlx/commit/6c9c44a66eeb60c7851fd20e04d40b6d31aa0f26)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement AuditLog record for updating the insight configuration settings ([21c3303](https://gitlab.com/commonground/nlx/nlx/commit/21c3303db0cf1c3b33f1b298138f438b89c8dd7e)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement AuditLog record for updating the organization settins ([014ce65](https://gitlab.com/commonground/nlx/nlx/commit/014ce6531e055bdc1cbfde374e4bb32afeafa483)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement AuditLogging for creating access request ([61df9cd](https://gitlab.com/commonground/nlx/nlx/commit/61df9cd4cfcce6defbd442d89005be28128f6ae6)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement AuditLogging for revoking access grants ([425d567](https://gitlab.com/commonground/nlx/nlx/commit/425d56764961bf69e7266c4e3153fc432cb8c158)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement AuditLogging when updating the organization insight settings ([ca09725](https://gitlab.com/commonground/nlx/nlx/commit/ca097254f4956fb5d2eec9cccc86fa6b95c6e212)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** implement authorization and check for admin role ([2c77533](https://gitlab.com/commonground/nlx/nlx/commit/2c7753373bf5cb7ec6eedeec97fa237d91a49140))
* **management:** implement billing costs for service ([15d5f8f](https://gitlab.com/commonground/nlx/nlx/commit/15d5f8fecaf9212a8974ddce72fc09aea2b852c0))
* **management:** implement polling for the AccessGrants section ([37d9938](https://gitlab.com/commonground/nlx/nlx/commit/37d993831b573b55b5982fc6af41834451dbfeaf)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** improve styling of AuditLogs ([91882f6](https://gitlab.com/commonground/nlx/nlx/commit/91882f63cff74a052e722dddfb43806a677c0b5e)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** move fetching application settings from PageTemplate to the App component ([1d41e07](https://gitlab.com/commonground/nlx/nlx/commit/1d41e07d0d583271a08166856b20d35989cc3eea))
* **management:** one cell for warnings service overview ([fb90916](https://gitlab.com/commonground/nlx/nlx/commit/fb90916c6f2ed6b40dc5f5db513bc246666bbdd1))
* **management:** parse useragent into operating system, browser and client ([aae2da4](https://gitlab.com/commonground/nlx/nlx/commit/aae2da4b39528f92d1582ed0170539a416dd6887)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** poll service statistics ([dd63e62](https://gitlab.com/commonground/nlx/nlx/commit/dd63e62ff485cc699e2898911c087b759e0a86d5)), closes [#1113](https://gitlab.com/commonground/nlx/nlx/issues/1113)
* **management:** remove dot from page description ([8b9bca0](https://gitlab.com/commonground/nlx/nlx/commit/8b9bca0bae7617cb9dcc1e1964e1b5d1e9805633)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** remove resolved TODO ([ab95711](https://gitlab.com/commonground/nlx/nlx/commit/ab95711aba4aed57c65751409c8dfb347a6b2998)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** replace userAgent with operating system, browser and client ([7788e13](https://gitlab.com/commonground/nlx/nlx/commit/7788e137322b64bdee55c7fb957a0f45bd10d281)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** show proper error message for creating user with unknown role ([67cd238](https://gitlab.com/commonground/nlx/nlx/commit/67cd23889f76a3b2937c05dbd7a81517d1bcc655))
* **management:** show proper error message when login fails ([479ccd2](https://gitlab.com/commonground/nlx/nlx/commit/479ccd27cb6df9a232a00526b9690fa3bd656df3))
* **management:** style AuditLog meta information + add test ([4661244](https://gitlab.com/commonground/nlx/nlx/commit/4661244fe629955cdc71a89d78acc35764fab37f)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** update audit log action names ([becb7cc](https://gitlab.com/commonground/nlx/nlx/commit/becb7cc5c4bc80a962b3d9534aa298443bd86fad)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** update Audit log page icon ([7b7886f](https://gitlab.com/commonground/nlx/nlx/commit/7b7886f8057cca3ddf0d93742375536a23285418)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* **management:** update AuditLogStore to use sorting from API response ([d27f950](https://gitlab.com/commonground/nlx/nlx/commit/d27f950f1728b278d365ecf0fcc660c417241352)), closes [#1016](https://gitlab.com/commonground/nlx/nlx/issues/1016)
* fix several issues with helm config for create-admin job ([848141b](https://gitlab.com/commonground/nlx/nlx/commit/848141bd3dce90e7733da38cd1947621eb366b13))

## [0.94.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.94.1...v0.94.2) (2021-01-22)


### Bug Fixes

* **management:** remove debug code that leads to logging OIDC options and disabling secure cookies ([5cf2c08](https://gitlab.com/commonground/nlx/nlx/commit/5cf2c08bc1237b314b50d58921139eef0078df73))

## [0.94.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.94.0...v0.94.1) (2021-01-21)


### Bug Fixes

* **docs:** resolve typos ([1e769e1](https://gitlab.com/commonground/nlx/nlx/commit/1e769e1fd9d7d96679e4298ebf55f967fe5a9d81))

# [0.94.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.93.0...v0.94.0) (2021-01-21)


### Bug Fixes

* replace postgres init script with job ([84aa60c](https://gitlab.com/commonground/nlx/nlx/commit/84aa60c9cb634e1208fe427ec0450e5af953ba59))
* **docs:** fix docs for TOML config ([5fa244d](https://gitlab.com/commonground/nlx/nlx/commit/5fa244df66eb44691867085041dfcade57bcb59d))
* **docs:** replace external image with local version to prevent CSP issues ([ddabd7d](https://gitlab.com/commonground/nlx/nlx/commit/ddabd7da534f94aa4fc123b01de8be08ad753a8c))
* **docs:** update image to use PostgreSQL instead of ETCD ([a6e5eb2](https://gitlab.com/commonground/nlx/nlx/commit/a6e5eb2e0be0e830cde63d0518d40c35c7f2ef87))
* **helm:** changed listen address of auth-service so it can be reached ([d1e887d](https://gitlab.com/commonground/nlx/nlx/commit/d1e887d62dd5d0fece1526c6ae1e73ea4f5f4d47))
* **management:** confirmationModal improvements ([add1183](https://gitlab.com/commonground/nlx/nlx/commit/add118393cc45a6032f1378cf822446655c3fb45))
* **management:** convert value too boolean when setting the property isOrganizationInway ([85906cc](https://gitlab.com/commonground/nlx/nlx/commit/85906cce5b46c6b3991efc5a8e037df90817836b))
* **management:** decrease table row height for incoming access requests ([f493db9](https://gitlab.com/commonground/nlx/nlx/commit/f493db966dc71a25de623c835c8500d51759f899)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** fix scheduler and api tests after moving to PostgreSQL ([dd5bcaa](https://gitlab.com/commonground/nlx/nlx/commit/dd5bcaa512bf73d4084f3aeb4b129e5283f979c8))
* **management:** make sure clearing organization inway works ([2410c87](https://gitlab.com/commonground/nlx/nlx/commit/2410c876faed89b0de3febb303c478e73d81e109))
* **management:** make sure valid pending requests are picked up by the scheduler ([8107178](https://gitlab.com/commonground/nlx/nlx/commit/810717894dbed7a4b0da638b7f7612ad59d01c92))
* **management:** make sure we can delete services without accessrequests ([82b0b9f](https://gitlab.com/commonground/nlx/nlx/commit/82b0b9fbc80a4185efd929d5cbc51529d81961ee))
* **management:** modal & confirm vertical positioning ([980c805](https://gitlab.com/commonground/nlx/nlx/commit/980c805c8f22f04ef8bf5cbcadee0c4d682049e4))
* **management:** modal close scenarios & aria support ([03ccdb1](https://gitlab.com/commonground/nlx/nlx/commit/03ccdb1fe1c5de571867063e203b3017868dad3e))
* **management:** modal fixes and feedback ([4cd8b39](https://gitlab.com/commonground/nlx/nlx/commit/4cd8b3916d8ff9c01a1dfa41413c5deaf036b2dd))
* **management:** modal improvements ([47b63d2](https://gitlab.com/commonground/nlx/nlx/commit/47b63d213174fd907be05f70eec5d945c0f16085))
* **management:** only create service in deployment if does not exists ([9a0bddf](https://gitlab.com/commonground/nlx/nlx/commit/9a0bddf26fd23a674eb7e02f83a8f09565ff2bcc))
* **management:** prevent collapsible content from jumping when showing updated content button ([e6a7cb7](https://gitlab.com/commonground/nlx/nlx/commit/e6a7cb7f95fff96aa51e918ad5c92b504e4cb789)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** processed feedback ([8489ba5](https://gitlab.com/commonground/nlx/nlx/commit/8489ba5a83a475f20b9860f7baea30163c8a7484))
* **management:** return exit 0 when migrations are up-to-date ([b156298](https://gitlab.com/commonground/nlx/nlx/commit/b156298b958ea6f24f13ab2f3158fdad09c973f8))
* **management:** services can be removed again ([0ba6aad](https://gitlab.com/commonground/nlx/nlx/commit/0ba6aad4c494a91dd833b1e4f96d30e25e8eb917))
* **management:** store the stacktrace when an access request fails ([0591789](https://gitlab.com/commonground/nlx/nlx/commit/05917893c687b649d37dfcfb25fe6ad5ee425e2c))
* **management:** unlock outgoing access requests after schedule loop is done ([2837260](https://gitlab.com/commonground/nlx/nlx/commit/2837260477bc21d2fe98f642a8cf232d632a9733))
* **management:** useInterval -> usePolling with pause & continue ([814fd55](https://gitlab.com/commonground/nlx/nlx/commit/814fd55ea061d4a498ca6a4edcc3ba21a4b3a648))
* deprecate privileged runners ([e9c06e1](https://gitlab.com/commonground/nlx/nlx/commit/e9c06e1d9db2303c135cd24f235c9d029f0e66ea))


### Features

* **directory:** limit max number of services per inway ([a949e41](https://gitlab.com/commonground/nlx/nlx/commit/a949e413387962bc8a0636a64b315d092736b816))
* **management:** add polling for IncomingAccessRequests when section is expanded ([b0d14ff](https://gitlab.com/commonground/nlx/nlx/commit/b0d14ff06b00029aef3488b0d1d4e16a63902372)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** add statistics of services endpoint ([0a6cf9d](https://gitlab.com/commonground/nlx/nlx/commit/0a6cf9d7a7134ee3d972df763eb9bbe1fca34b28))
* **management:** add update incoming access requests button ([5ec6bac](https://gitlab.com/commonground/nlx/nlx/commit/5ec6bac6151d90406ad7416aa679f9efd1185ed3)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** automatically poll changes for the DirectoryDetail page ([c6693c3](https://gitlab.com/commonground/nlx/nlx/commit/c6693c3393e461048526eeb75be54df16b802007)), closes [#1131](https://gitlab.com/commonground/nlx/nlx/issues/1131)
* **management:** automatically update count incoming access requests for service detail ([a1dd9f8](https://gitlab.com/commonground/nlx/nlx/commit/a1dd9f8f66475d308cba252de5d1dc3362d8390c)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** confirmationModal @ directory detail ([d1a556c](https://gitlab.com/commonground/nlx/nlx/commit/d1a556c6f9f25701db6df1ffbed3d5fb43fced93))
* **management:** confirmationModal @ incoming access requests ([80fa84a](https://gitlab.com/commonground/nlx/nlx/commit/80fa84a0a17d307fc84e4354afef737303e87ee0))
* **management:** confirmationModal @ remove service ([c841e91](https://gitlab.com/commonground/nlx/nlx/commit/c841e91c087621f57846fa4e186e9ef84d2dba0a))
* **management:** confirmationModal @ revoke grants ([84a77ac](https://gitlab.com/commonground/nlx/nlx/commit/84a77acd1a134e94fd744b55532d81ce8dfb0b5e))
* **management:** confirmModal @ directory list ([57d7a84](https://gitlab.com/commonground/nlx/nlx/commit/57d7a8437725adf77518a817ea865d64b43ad87e))
* **management:** fetching a single service should add the model to the store ([6fb0596](https://gitlab.com/commonground/nlx/nlx/commit/6fb0596e61c77fb46e63cf1444c1a2cbb4623b16)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** implement button to trigger updating the UI ([4341392](https://gitlab.com/commonground/nlx/nlx/commit/4341392007097df8ebd78a13d67947073900ecb4)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** modal component ([1a9b547](https://gitlab.com/commonground/nlx/nlx/commit/1a9b54760f44389b8532b192d8a8caba63249ac0))
* **management:** modalConfirm component in general settings ([6c6c236](https://gitlab.com/commonground/nlx/nlx/commit/6c6c2366a1172ed80ba05b1ca614ae9080a3a70f))
* **management:** position update UI button when currently no access requests are listed ([515c11f](https://gitlab.com/commonground/nlx/nlx/commit/515c11f2e0c41d7660a9344c826c4506ee1a8a48)), closes [#1114](https://gitlab.com/commonground/nlx/nlx/issues/1114)
* **management:** replace ETCD with PostgreSQL ([7b282a6](https://gitlab.com/commonground/nlx/nlx/commit/7b282a6483351825d4c2a290b47e31c5b449432f))
* **management:** validate OIDC claims in authentication middleware ([5d2cd69](https://gitlab.com/commonground/nlx/nlx/commit/5d2cd69d4f4284e6d2bafafda7b0c15cfe8a1f54))
* add Permission-Policy header to front ends ([d8ce7a9](https://gitlab.com/commonground/nlx/nlx/commit/d8ce7a9039ed5451430f33d1829e627474a9495e))

# [0.93.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.92.0...v0.93.0) (2021-01-06)


### Bug Fixes

* **management:** eslint-config version mismatch ([04c4d5c](https://gitlab.com/commonground/nlx/nlx/commit/04c4d5c08f79ec4cb2b296987ad2c606a2764782))
* set explicit Gitlab environment in deployments ([bc09726](https://gitlab.com/commonground/nlx/nlx/commit/bc097262197cf7702d99775b38c81549beb00745))
* **docs:** remove creating env file since that doesn't always work ([563515b](https://gitlab.com/commonground/nlx/nlx/commit/563515b91be71e525790bf7d4313e845ba5e598d))
* **docs:** update docs to point to the new nlx-try-me repository ([f2c5755](https://gitlab.com/commonground/nlx/nlx/commit/f2c5755ae9a997745e63c2a7429882d059dc5ebd))
* **inway:** log error on service config warning ([1dbe883](https://gitlab.com/commonground/nlx/nlx/commit/1dbe883c99c79a58c99ccd8d24b0731b0f720a62))
* **management:** a11y warnings ([b1cfe8b](https://gitlab.com/commonground/nlx/nlx/commit/b1cfe8b2328843ea83644e3b5d6c60f6180b01f4))
* **management:** collapsible animation ([8dc119f](https://gitlab.com/commonground/nlx/nlx/commit/8dc119f6153642bff5de4bb112a029a153ec9dc3))
* **management:** proper handling non-existent service ([1f945f0](https://gitlab.com/commonground/nlx/nlx/commit/1f945f02a5dfb78a2da98573838e95b142798c96))
* move try-me code to it's own repository ([1ff2702](https://gitlab.com/commonground/nlx/nlx/commit/1ff2702a0be588146c6212dca58a55f3b417ce0a))


### Features

* **common:** implement safe permissions for private key in tls package ([012696c](https://gitlab.com/commonground/nlx/nlx/commit/012696cdd346446e662b989559070c9de0a43161))
* **directory:** add example documentation URL to kentekenregister ([1d5af6c](https://gitlab.com/commonground/nlx/nlx/commit/1d5af6c6f5248ebef8b1004a8540bcae2a3cc257)), closes [#627](https://gitlab.com/commonground/nlx/nlx/issues/627)
* **directory:** rework detail pane UI ([22068ac](https://gitlab.com/commonground/nlx/nlx/commit/22068ac745accae550727347722ae01ea09b4432)), closes [#627](https://gitlab.com/commonground/nlx/nlx/issues/627)
* **directory:** show documentation URL if available for service ([042a1fd](https://gitlab.com/commonground/nlx/nlx/commit/042a1fd0ae826cd98abe3bd94ea60f6712a1b18d)), closes [#627](https://gitlab.com/commonground/nlx/nlx/issues/627)
* **management:** add ApplicationStore and update on change setting inway ([2fde4c8](https://gitlab.com/commonground/nlx/nlx/commit/2fde4c8a0f92b5025101b41d7fdf340d4e6b3a10))
* **management:** add GlobalAlert component ([52b16f4](https://gitlab.com/commonground/nlx/nlx/commit/52b16f424b184b09ac4bdb82cb85464d6984e246))
* **management:** add logout icon ([e0a52b4](https://gitlab.com/commonground/nlx/nlx/commit/e0a52b467443ed51339dbb406a2eb404ee588c08)), closes [#1049](https://gitlab.com/commonground/nlx/nlx/issues/1049)
* **management:** add migrations ([0c6d4df](https://gitlab.com/commonground/nlx/nlx/commit/0c6d4df7a40e216c8d0eacfdb69f290a238bf09b))
* **management:** add warning when removing organization inway ([db99a62](https://gitlab.com/commonground/nlx/nlx/commit/db99a6206db685c4b3b5141cd97250e944e813ff))
* **management:** deprecated browser page ([796742a](https://gitlab.com/commonground/nlx/nlx/commit/796742aeada75ba7eb6b133eccc42789404784b8))
* **management:** highlight selected Directory Service on the Directory page ([bafac73](https://gitlab.com/commonground/nlx/nlx/commit/bafac73593f94a139566b0926c16ffd7e8194f35)), closes [#941](https://gitlab.com/commonground/nlx/nlx/issues/941)
* **management:** highlight selected Inway on the Inways overview page ([0106c43](https://gitlab.com/commonground/nlx/nlx/commit/0106c436bd83fa1af70dcf854cee07965f4298e1)), closes [#941](https://gitlab.com/commonground/nlx/nlx/issues/941)
* **management:** highlight selected Service on the Services overview page ([35ff6d0](https://gitlab.com/commonground/nlx/nlx/commit/35ff6d06255c07f82231b9d8f326c943223f5cb6)), closes [#941](https://gitlab.com/commonground/nlx/nlx/issues/941)
* **management:** replace native select with Select from Design System ([c80d180](https://gitlab.com/commonground/nlx/nlx/commit/c80d1800f7c3f24b4646d471140affea2edceee8)), closes [#1071](https://gitlab.com/commonground/nlx/nlx/issues/1071)
* **management:** show access request button directory ([c0727a9](https://gitlab.com/commonground/nlx/nlx/commit/c0727a92adc7ee5c687449e6ab250d65dd491e91)), closes [#1087](https://gitlab.com/commonground/nlx/nlx/issues/1087)
* **management:** show warning when no organization inway ([d847da0](https://gitlab.com/commonground/nlx/nlx/commit/d847da0390ea30bd84ed7de6a9f273f51e6820c9)), closes [#1045](https://gitlab.com/commonground/nlx/nlx/issues/1045)
* improved input validation in directory and management ([d455854](https://gitlab.com/commonground/nlx/nlx/commit/d45585439a8d5a1ab95257ff3239245baf132438))

# [0.92.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.91.0...v0.92.0) (2020-11-25)


### Bug Fixes

* **directory:** back to styled-components@5.1 ([8ae37eb](https://gitlab.com/commonground/nlx/nlx/commit/8ae37eb680c5f3d4d2a80c7efb6bd7da42aade8c))
* **directory:** prevent hard fail on versionLogger ([730deae](https://gitlab.com/commonground/nlx/nlx/commit/730deaefbef0c2ab05bfed3692cd99437f87ff6c))
* **directory:** redoc ([6c93652](https://gitlab.com/commonground/nlx/nlx/commit/6c936523934e4e4cdee47acff6cb14f0fbbc4a9c))
* **directory:** return documentation url ([b118359](https://gitlab.com/commonground/nlx/nlx/commit/b118359f202143daefb6f2858af26850be4e28e3))
* **helm:** certportal listen address ([8bf727c](https://gitlab.com/commonground/nlx/nlx/commit/8bf727c20e4b9cc1c31559ab55239461b36a6154))
* **management:** align detailview header icons ([cb77a2d](https://gitlab.com/commonground/nlx/nlx/commit/cb77a2d252333aa17537914724b77451e43d41d4))
* **management:** align state & api type ([133b5bf](https://gitlab.com/commonground/nlx/nlx/commit/133b5bfb9306414996600e99a253f9d6a8d30616))
* **management:** buttons visible in sub-drawer ([77ada96](https://gitlab.com/commonground/nlx/nlx/commit/77ada968875526d49d1a887c13a9cd89dd3af20f))
* **management:** make toaster animation visible ([0870240](https://gitlab.com/commonground/nlx/nlx/commit/087024017dc3dd62e4b3b6fd58e08e144e925dd9))


### Features

* **management:** add contact section and documentation links ([c4a9dbc](https://gitlab.com/commonground/nlx/nlx/commit/c4a9dbc450c42e825fe2aab5c92992ce36c43efe))
* **management:** add details to directory detail drawer (wip) ([58a9194](https://gitlab.com/commonground/nlx/nlx/commit/58a91940e1a72a5abefb409e49beb9f5f68442f8))
* **management:** return documentation url in directory service ([97c616f](https://gitlab.com/commonground/nlx/nlx/commit/97c616f6ee5e6bb4d68cc966a855148939f7f16e))
* **management:** return public contact email in directory service ([19a6fe1](https://gitlab.com/commonground/nlx/nlx/commit/19a6fe160904792fe755e6970a85252b303b854d))

# [0.91.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.90.0...v0.91.0) (2020-11-24)


### Bug Fixes

* **insight:** replace unsafe methods with safe methods ([b65bc2b](https://gitlab.com/commonground/nlx/nlx/commit/b65bc2b6b2a965047ced8c05581f6f63af58d4ab))
* **management:** allow retries after revoked requests ([1c5892a](https://gitlab.com/commonground/nlx/nlx/commit/1c5892ac27f9f9829a2db035e2c208b8d52179c6))
* **management:** better empty collapsible rendering ([5398ab0](https://gitlab.com/commonground/nlx/nlx/commit/5398ab064e51bd42a3de59d605f5f452f3273635))
* **management:** cleanup ([3646bb6](https://gitlab.com/commonground/nlx/nlx/commit/3646bb64ef7950e978d922aa2268610d695ab18d))
* **management:** directory list view ([c13a86d](https://gitlab.com/commonground/nlx/nlx/commit/c13a86d30333dde0f1aa6fc1fadce7110594845c))
* **management:** display loader when re-sending access request ([39aef55](https://gitlab.com/commonground/nlx/nlx/commit/39aef552729915b0e21977b54338fedbd3ae3d3b)), closes [#1057](https://gitlab.com/commonground/nlx/nlx/issues/1057)
* **management:** don't listen on loopback address as it doesn't work in our setup ([e5c9075](https://gitlab.com/commonground/nlx/nlx/commit/e5c9075b4a461227e0c5ee8acbc9f6ce8e576439))
* **management:** don't overwrite accessProof argument in directoryServiceAccessState] ([ed9077a](https://gitlab.com/commonground/nlx/nlx/commit/ed9077acc34439ee3d5433c46a8bce29968c1f70))
* **management:** fix connection leak in scheduler ([f358504](https://gitlab.com/commonground/nlx/nlx/commit/f3585048ecc83cb43d4b9f25a7d0119a0dc8a321))
* **management:** fix invalid organization name bug and added tests ([5d2cc96](https://gitlab.com/commonground/nlx/nlx/commit/5d2cc96eb344c04bcf95feb4c6698fdb5c28fa0e))
* **management:** fmt package missing ([605391c](https://gitlab.com/commonground/nlx/nlx/commit/605391c133ca8eeed8b2771e3bc169cc306680d4))
* **management:** ignore .eslintcache ([b4d0501](https://gitlab.com/commonground/nlx/nlx/commit/b4d05012ef5e18d5c13ee28d94f43b3e6f609415))
* **management:** implemented tests for creating and retrying AccessRequests ([830a849](https://gitlab.com/commonground/nlx/nlx/commit/830a849d1fe7bb18760a097e6d92744b4e4dc633))
* **management:** little UI tweaks ([a7ecdb7](https://gitlab.com/commonground/nlx/nlx/commit/a7ecdb7cb3811dcf28754403d1e19a6a8debbaf6))
* **management:** make sure nlxctl currectly uses the management api ([59d12c8](https://gitlab.com/commonground/nlx/nlx/commit/59d12c8674c581538bc384f30be196000a31ecbe)), closes [#1007](https://gitlab.com/commonground/nlx/nlx/issues/1007)
* **management:** nl language strings [#1080](https://gitlab.com/commonground/nlx/nlx/issues/1080) ([cd6e82f](https://gitlab.com/commonground/nlx/nlx/commit/cd6e82f2e5ead208cd60920df9f1a0bd6f6bc111))
* **management:** no new accessRequest when revoked ([fd23b50](https://gitlab.com/commonground/nlx/nlx/commit/fd23b50caffa9b918245398079ec1612583aa98e))
* **management:** normalised/fixed icon files ([6028c16](https://gitlab.com/commonground/nlx/nlx/commit/6028c1660e9074c462e38559087e01399fd3907c))
* **management:** normalized all custom icon styles ([d803ed3](https://gitlab.com/commonground/nlx/nlx/commit/d803ed3c962aedd69c5902e6d1e2bbdb65dcf3ff))
* **management:** re-enable animation when toggling the user navigation ([4e6a5fe](https://gitlab.com/commonground/nlx/nlx/commit/4e6a5fea376fa343ef86ab09fff3974815608be0)), closes [#1081](https://gitlab.com/commonground/nlx/nlx/issues/1081)
* **management:** set minimum height of the state indicator of an access request ([d0f3070](https://gitlab.com/commonground/nlx/nlx/commit/d0f3070ee8e9dfcfa3408c291c11e735868c9750))
* **management:** split primarynav items in two parts ([c548dd5](https://gitlab.com/commonground/nlx/nlx/commit/c548dd5f570d778504753488149c35710c663e4f))
* **management:** test suddenly throwing errors ([534d3bd](https://gitlab.com/commonground/nlx/nlx/commit/534d3bdddcc46dfcfbe9bc54c299f280f9c043f5))
* **management:** update incoming access request label color for hover state ([6d3408e](https://gitlab.com/commonground/nlx/nlx/commit/6d3408e3351ad8b5450fd3e20e592c59bb77eb11)), closes [#1007](https://gitlab.com/commonground/nlx/nlx/issues/1007)
* **management:** use consistent dutch text in button title and edit heading ([67fb801](https://gitlab.com/commonground/nlx/nlx/commit/67fb801fcd97cb26f2c0029ccb2958f97a0003e2))


### Features

* **directory:** set timeout in directory monitor ([3cb62db](https://gitlab.com/commonground/nlx/nlx/commit/3cb62db26d09bcd69d32fc08751eac7f20c24123))
* **management:** add IncomingAccessRequestStore ([8cde0ec](https://gitlab.com/commonground/nlx/nlx/commit/8cde0ecfa536baebe431563c6f26371729b2ed54)), closes [#1020](https://gitlab.com/commonground/nlx/nlx/issues/1020)
* **management:** add title attributes for approving, rejecting and revoking access ([a8fab5b](https://gitlab.com/commonground/nlx/nlx/commit/a8fab5b40b8fba8a39b80969c0eb7b42d831246b)), closes [#1057](https://gitlab.com/commonground/nlx/nlx/issues/1057)
* **management:** display access state per directoryService ([d40eb88](https://gitlab.com/commonground/nlx/nlx/commit/d40eb88eff187a4885010a0fb04928f6ca7edbbe))
* **management:** display amount of incoming access requests on the service overview page ([0a59c9b](https://gitlab.com/commonground/nlx/nlx/commit/0a59c9b27015b77e10e520344b49b18aac875c8f)), closes [#1007](https://gitlab.com/commonground/nlx/nlx/issues/1007)
* **management:** enable creating a new OutgoingAccessRequest when the previous one was rejected ([86f5cce](https://gitlab.com/commonground/nlx/nlx/commit/86f5ccefabbccdb237bd92f089f941c3dba2e34e)), closes [#1080](https://gitlab.com/commonground/nlx/nlx/issues/1080)
* **management:** enable requesting access when a previous request was revoked ([e8b6478](https://gitlab.com/commonground/nlx/nlx/commit/e8b6478dc6e29383741c68ce17587eaeffe5e4c8)), closes [#1043](https://gitlab.com/commonground/nlx/nlx/issues/1043)
* **management:** implement polling for access proof ([b4402e8](https://gitlab.com/commonground/nlx/nlx/commit/b4402e8570a5775b4ecc7e8cc5494a60b0e8d2a9))
* **management:** implement polling for access requests ([c285ffd](https://gitlab.com/commonground/nlx/nlx/commit/c285ffd4f31df03c3c8d2b26d4aca2f27d11785b))
* **management:** implement reference ID in AccessRequest and verify that it belongs to a proof ([3ba994d](https://gitlab.com/commonground/nlx/nlx/commit/3ba994d0a59505f52e67d91c96995aa4eede1d8b))
* **management:** implement reference to AccessRequest in an AccessProof ([2fd6f8e](https://gitlab.com/commonground/nlx/nlx/commit/2fd6f8ec35ee16df3b0767416bec8ca6fd911846))
* **management:** implement UI for AccessRequest errors in failure mode ([ba48df7](https://gitlab.com/commonground/nlx/nlx/commit/ba48df72ff9ea45a63347df7eb316452a1acb1d6))
* **management:** improve alignment of actions for the IncomingAccessRequest section ([aeacdf0](https://gitlab.com/commonground/nlx/nlx/commit/aeacdf07216300c2205429fd19e0a2b896c57fc0))
* **management:** introduce AccessProof in mobx ([1c268e1](https://gitlab.com/commonground/nlx/nlx/commit/1c268e16d86c27b6cd78ce447c7d4a732c994344))
* **management:** introduce Switch component ([f947ecc](https://gitlab.com/commonground/nlx/nlx/commit/f947ecca3ab84339c95508d47c435183e91887b8))
* **management:** parse and store AccessRequest error details ([b0671f6](https://gitlab.com/commonground/nlx/nlx/commit/b0671f694b31ad539b9d6e89c54d6414340009a4))
* **management:** parse and store AccessRequest error details ([279e51b](https://gitlab.com/commonground/nlx/nlx/commit/279e51ba86d786e17a137fe8fb789d449136e917))
* **management:** parse gRPC errors properly ([56797ad](https://gitlab.com/commonground/nlx/nlx/commit/56797ade88bbf9f739f08812372222fc48f07075))
* **management:** prevent Service row from being too high ([3b97243](https://gitlab.com/commonground/nlx/nlx/commit/3b972438d614c754eadc6fc99683f5e3a56aeb8c)), closes [#1007](https://gitlab.com/commonground/nlx/nlx/issues/1007)
* **management:** re-fetch service when accepting/rejecting access ([3239042](https://gitlab.com/commonground/nlx/nlx/commit/3239042e24a8a01455dced4ab076739f2f1ecc4d)), closes [#1057](https://gitlab.com/commonground/nlx/nlx/issues/1057)
* **management:** reject access requests ([1514e54](https://gitlab.com/commonground/nlx/nlx/commit/1514e54c0b7b66e0e62b5ed70cc6dde04c52ac98))
* **management:** reload services list when approving/rejecting an access request ([3a6109a](https://gitlab.com/commonground/nlx/nlx/commit/3a6109a5e2431797f69e1a09ce34f02b58b88e08)), closes [#1007](https://gitlab.com/commonground/nlx/nlx/issues/1007)
* **management:** remove redundant properties when creating a service ([53900d9](https://gitlab.com/commonground/nlx/nlx/commit/53900d9980dab1f4f088e2add634afa79090f290)), closes [#1007](https://gitlab.com/commonground/nlx/nlx/issues/1007)
* **management:** remove redundant properties when retrieving a single service ([67ac63e](https://gitlab.com/commonground/nlx/nlx/commit/67ac63e08ab1ef0eead9626a172b5459f44b064d)), closes [#1007](https://gitlab.com/commonground/nlx/nlx/issues/1007)
* **management:** remove redundant properties when updating a service ([f1209fa](https://gitlab.com/commonground/nlx/nlx/commit/f1209fa883e43c2c2c20715ff9ee3e99bce41f92)), closes [#1007](https://gitlab.com/commonground/nlx/nlx/issues/1007)
* **management:** return amount of incoming access requests when listing all services ([b8952b9](https://gitlab.com/commonground/nlx/nlx/commit/b8952b9dbcda36d8b763251fae9a8273b4b7b1ad)), closes [#1007](https://gitlab.com/commonground/nlx/nlx/issues/1007)
* **management:** show different counter when incoming accessrequest count is 0 ([c7b1e5d](https://gitlab.com/commonground/nlx/nlx/commit/c7b1e5db15c2808403131904dcf9db00ef87451d))
* **management:** show request access button when rejected ([074d0a7](https://gitlab.com/commonground/nlx/nlx/commit/074d0a7b90a273a883330c610d50700738076186))
* **management:** subtract incoming access request counter client side after approving/rejecting ([85bcb37](https://gitlab.com/commonground/nlx/nlx/commit/85bcb37e52807c0e3ebe8cc1cb5795c2cf32d2ab)), closes [#1007](https://gitlab.com/commonground/nlx/nlx/issues/1007)
* **management:** update font weight for the incoming requests label ([0e9cc25](https://gitlab.com/commonground/nlx/nlx/commit/0e9cc25456cdc2d6ce33d64e4865b9c498835974)), closes [#1007](https://gitlab.com/commonground/nlx/nlx/issues/1007)
* default listen on loopback address ([7945faf](https://gitlab.com/commonground/nlx/nlx/commit/7945faf3b05148343ab3a97e44047f892c3991a1))

# [0.90.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.89.2...v0.90.0) (2020-10-22)


### Bug Fixes

* **directory:** panic retieving API spec ([7bcbadf](https://gitlab.com/commonground/nlx/nlx/commit/7bcbadfae11fd92034edb549dec8e97727644ba2))
* **directory:** rename remaining instances of handlers ([7db4d8f](https://gitlab.com/commonground/nlx/nlx/commit/7db4d8f05697972198c4dccbaca470a82313f848))
* **management:** check not found error GetLastOutgoingAccesRequest ([795f572](https://gitlab.com/commonground/nlx/nlx/commit/795f572b5e849ab9212fb5328bdb1cf5dde2dd36))
* **management:** correctly pass required stores when setting up stores in the rootstore ([6650fe8](https://gitlab.com/commonground/nlx/nlx/commit/6650fe8e982dc62c40168a4278862e3ebc1d3b87)), closes [#1023](https://gitlab.com/commonground/nlx/nlx/issues/1023)
* **management:** fix english text strings ([dc3f7e2](https://gitlab.com/commonground/nlx/nlx/commit/dc3f7e2f0f7075a3abd5ef34d80e0ba416ed2e6c))
* **management:** fix tests after upgrade mobx ([04c3dbe](https://gitlab.com/commonground/nlx/nlx/commit/04c3dbeb1d0fb9795cc959dadc1cdc9b1fa36522))
* **management:** ip address in inway detail view ([4956372](https://gitlab.com/commonground/nlx/nlx/commit/49563722afb6d4406d1221dcb553c255005ebbe1))
* **management:** linting errors ([17febd6](https://gitlab.com/commonground/nlx/nlx/commit/17febd60e7312de1fa723a817519f1c7dc364529))
* **management:** mobx this bindings ([10b8045](https://gitlab.com/commonground/nlx/nlx/commit/10b8045094bf78536c9eb1e8560c4d11eb845a3a))
* **management:** read request body when updating insight configuration over HTTP ([8f2d0c1](https://gitlab.com/commonground/nlx/nlx/commit/8f2d0c106867ed3c92f09b17d2f01206a4510a1a)), closes [#1026](https://gitlab.com/commonground/nlx/nlx/issues/1026)
* **management:** reference rootStore directly instead of via 'this' ([e8f2988](https://gitlab.com/commonground/nlx/nlx/commit/e8f29885247349f0a736323d88f4f564115128aa)), closes [#1023](https://gitlab.com/commonground/nlx/nlx/issues/1023)
* **management:** resolve issues with fetching directoryServices ([b488970](https://gitlab.com/commonground/nlx/nlx/commit/b488970b4124a74abe6a539ac4674c3c2f97badb)), closes [#1023](https://gitlab.com/commonground/nlx/nlx/issues/1023)
* **management:** use latest image of the  management ui in docker compose ([3ccbdcd](https://gitlab.com/commonground/nlx/nlx/commit/3ccbdcd6ad47f5136a41fe4914d7c73eef1bfad9))
* add translations ([9fa4b15](https://gitlab.com/commonground/nlx/nlx/commit/9fa4b15adc0d1f74cc8185767cefa945a36bc3cf))
* **management:** resolve security issues reported by ESLint ([93b3f63](https://gitlab.com/commonground/nlx/nlx/commit/93b3f63da2cb2c757cf5d362e4ba8e7ecbc66009))


### Features

* inter-NLX components TLS 1.3 only ([c11b172](https://gitlab.com/commonground/nlx/nlx/commit/c11b172dacce31da2f5bbeabcaaf9f19a793e888))
* **directory:** no default value for Postgres DSN ([41b2123](https://gitlab.com/commonground/nlx/nlx/commit/41b21232b5a24751794323edec914702ddcd4cbe))
* **docs:** add try nlx management docs ([9ba0ff5](https://gitlab.com/commonground/nlx/nlx/commit/9ba0ff5c7321030d29ecef2cd223e73527765983))
* **insight:** no default value for Postgres DSN ([3c5645b](https://gitlab.com/commonground/nlx/nlx/commit/3c5645b5e4bb1e27b2e1c66fee0a3d00156aa8cc))
* **inway:** change config retrieval interval to 10 seconds ([07bdb1a](https://gitlab.com/commonground/nlx/nlx/commit/07bdb1a4e8e796c76ee9fd9c435b6cdb86cb6f38))
* **inway:** no default value for Postgres DSN ([cbde5d5](https://gitlab.com/commonground/nlx/nlx/commit/cbde5d5f91c545d928d4453b0cd750cb74d978c6))
* **management:** add OutgoingAccessRequest store with create method ([b957cc9](https://gitlab.com/commonground/nlx/nlx/commit/b957cc9ec823b3be2d8239da818013aa6097ffec)), closes [#1023](https://gitlab.com/commonground/nlx/nlx/issues/1023)
* **management:** add retry button for a failing Access Request ([3ecd391](https://gitlab.com/commonground/nlx/nlx/commit/3ecd391f76cb9704683abe466a8c2179433799a9)), closes [#1023](https://gitlab.com/commonground/nlx/nlx/issues/1023)
* **management:** add reverse proxy for the management-api in nginx ([e3afda1](https://gitlab.com/commonground/nlx/nlx/commit/e3afda19f52d7af90f45f8f9ee11b3517444ba35))
* **management:** add RPC to send access request ([1cb85c0](https://gitlab.com/commonground/nlx/nlx/commit/1cb85c0736d0b299a948747d26bd5aa9d16194e1))
* **management:** add sub navigation to the settings page ([fc0887b](https://gitlab.com/commonground/nlx/nlx/commit/fc0887b8404780d535ec235f4ac80946be5b26a4)), closes [#1026](https://gitlab.com/commonground/nlx/nlx/issues/1026)
* **management:** display error toast when revoking access fails ([98f875b](https://gitlab.com/commonground/nlx/nlx/commit/98f875bb350d300943282d38323bb39c3dca7f8a))
* **management:** display loading status when retrying to request access ([ab45ebe](https://gitlab.com/commonground/nlx/nlx/commit/ab45ebe8cd9c638c1f29744f0e7cb2bee15dbc5e)), closes [#1023](https://gitlab.com/commonground/nlx/nlx/issues/1023)
* **management:** enable updating insight settings via NLX Management ([1547ac8](https://gitlab.com/commonground/nlx/nlx/commit/1547ac8deeeefdeac83f08b09ca0475b0a0025f0))
* **management:** implement create access request for the DirectoryStore ([3ef4aa5](https://gitlab.com/commonground/nlx/nlx/commit/3ef4aa5af06a24b4dad32fdbf436115629f2e8e9)), closes [#1023](https://gitlab.com/commonground/nlx/nlx/issues/1023)
* **management:** implement revoke access action in UI ([e165965](https://gitlab.com/commonground/nlx/nlx/commit/e165965a62feaa2967e9b1d1899230a25f503853)), closes [#1020](https://gitlab.com/commonground/nlx/nlx/issues/1020)
* **management:** rename settings navigation items ([46ed704](https://gitlab.com/commonground/nlx/nlx/commit/46ed70408367b6445e60d711911fb7313b4cb550))
* **management:** set padding right for settings navigation items ([898f1d2](https://gitlab.com/commonground/nlx/nlx/commit/898f1d2a3e82f12167624ab06cce372fd68a6a39)), closes [#1026](https://gitlab.com/commonground/nlx/nlx/issues/1026)
* **management:** show dates for access request state ([4473359](https://gitlab.com/commonground/nlx/nlx/commit/447335983c560502d9086a6df960c6affc3ea851))
* show outgoing access request states ([c2f7486](https://gitlab.com/commonground/nlx/nlx/commit/c2f748681a7e44088e587cdd7e572eb10f404d4b))
* **management:** move re-fetching DirectoryModel from DetailPage to the router ([5224d4c](https://gitlab.com/commonground/nlx/nlx/commit/5224d4c2f8643f12038b3028da555aab2a94e7c6)), closes [#1023](https://gitlab.com/commonground/nlx/nlx/issues/1023)
* **management:** return latest access proof in directory service list ([5497719](https://gitlab.com/commonground/nlx/nlx/commit/5497719212961791522ebe24bf33806da38cc1d4))
* add access grants to service detail page ([ddc75d2](https://gitlab.com/commonground/nlx/nlx/commit/ddc75d25b64d1292e7435ca3caa1909c42966e19))
* add docker-compose setup for NLX with management ([6113b0a](https://gitlab.com/commonground/nlx/nlx/commit/6113b0a471d55b243c150be4803543fb71ad4f17))
* implement access request tasks in nlxctl ([7b3f6b6](https://gitlab.com/commonground/nlx/nlx/commit/7b3f6b68abff0ea664450443ba9f514a92ccd778))
* **management:** replace inline Alert with Toaster when updating the insight settings fails ([6aea13d](https://gitlab.com/commonground/nlx/nlx/commit/6aea13d00305c436df4a4a2f54264add32bd1c77)), closes [#1026](https://gitlab.com/commonground/nlx/nlx/issues/1026)
* **management:** replace inline Alert with Toaster when updating the settings fails ([3b576fd](https://gitlab.com/commonground/nlx/nlx/commit/3b576fd8a7bb2f1afbcc8d0c4abb8f6ae86aa334)), closes [#1026](https://gitlab.com/commonground/nlx/nlx/issues/1026)
* **management:** return not found error ([7cc7c42](https://gitlab.com/commonground/nlx/nlx/commit/7cc7c42c4b0fdda62b7df37131812af400323743))
* **management:** revoke access request ([a4a6abe](https://gitlab.com/commonground/nlx/nlx/commit/a4a6abedd287f1f83fbb19c100b43b741cd93495))
* **management:** view access grants for a service ([0f88a21](https://gitlab.com/commonground/nlx/nlx/commit/0f88a217a67f8daddf534322ecf37b0347fca6ea))
* **management:** when getting accesgrants return a 404 if the service does not exits ([8653571](https://gitlab.com/commonground/nlx/nlx/commit/86535712eb47407f04a7edb1f6113a7f265b233e))
* **outway:** no default value for Postgres DSN ([a369bff](https://gitlab.com/commonground/nlx/nlx/commit/a369bff74603335b08d86c88eba19499696eedf4))

## [0.89.2](https://gitlab.com/commonground/nlx/nlx/compare/v0.89.1...v0.89.2) (2020-10-05)


### Bug Fixes

* **ca-certportal:** caHost config value ([91a98c5](https://gitlab.com/commonground/nlx/nlx/commit/91a98c5f3216b3b0ba002a918061bfcc10e29475))

## [0.89.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.89.0...v0.89.1) (2020-10-02)


### Bug Fixes

* **helm:** sync ca-certportal chart with the rest ([7b21e6b](https://gitlab.com/commonground/nlx/nlx/commit/7b21e6b358c659ff161f8eb46bb3d3f8edb0ec4c))

# [0.89.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.88.0...v0.89.0) (2020-10-01)


### Bug Fixes

* **common:** gosec G304 file inclusion ([ec63674](https://gitlab.com/commonground/nlx/nlx/commit/ec636741bc9d4ab2f0f4919e5cee03d8f555facf))
* **directory:** gosec G402 TLS MinVersion ([16d7754](https://gitlab.com/commonground/nlx/nlx/commit/16d7754cc0c7a1a721d3bba1e6805a77d3bf3c42))
* **directory:** resolve vulnerable YAML parsing dependency ([9268d56](https://gitlab.com/commonground/nlx/nlx/commit/9268d56d951d2fb000775d56068bf853c58ca468))
* **docs:** fix build failure for docusaurus after upgrade ([7cd3552](https://gitlab.com/commonground/nlx/nlx/commit/7cd355258751591c900b2f86be61308d5dda5878))
* **inway:** gosec G402 TLS MinVersion ([f37cfde](https://gitlab.com/commonground/nlx/nlx/commit/f37cfdecfc6d89a15666595a64f4803211434230))
* **management:** add unspecified state to fix default value issue ([1a2a590](https://gitlab.com/commonground/nlx/nlx/commit/1a2a59087d53e79a2c69303086c39bf2b0b5f246))
* **management:** apply receive state for incoming access requests ([9074cde](https://gitlab.com/commonground/nlx/nlx/commit/9074cde4a7c36dfa12e844344b282217cddd14aa))
* **management:** fix issue with updating settings in the management UI ([180a883](https://gitlab.com/commonground/nlx/nlx/commit/180a883916b661a57883b1e8719fa680834dcf7e))
* **management:** fix issues with locking, etcd watch and concurrency ([3894e5d](https://gitlab.com/commonground/nlx/nlx/commit/3894e5d4393df89f1e883ce639f38b82bc1e561b))
* **management:** gosec G402 TLS MinVersion ([6a603bb](https://gitlab.com/commonground/nlx/nlx/commit/6a603bb24ff701645facb4f7bf3f3c62dcf0d2f3))
* **management:** implement a fallback for dead-locked AccessRequests ([2f33ce7](https://gitlab.com/commonground/nlx/nlx/commit/2f33ce7fb30c5cbb5cdf32c2174d3383520cab01))
* **management:** resolve potential memory leak for usePromise hook ([ccf7fc2](https://gitlab.com/commonground/nlx/nlx/commit/ccf7fc2679ec4b5a9bc1221deffe2b5cebbda6f2)), closes [core/team#75](https://gitlab.com/core/team/issues/75)
* **outway:** gosec G402 TLS MinVersion ([8fa420e](https://gitlab.com/commonground/nlx/nlx/commit/8fa420e4b789e3c3733fac99993f0449cb72a5c7))
* temporary ignore CommonName deprecation in Go 1.15 ([b6339a3](https://gitlab.com/commonground/nlx/nlx/commit/b6339a38990731ddc782ffa3d6a6e2c0d84797d1))


### Features

* **ca-certportal:** add SAN to csr when missing ([d2dfb6a](https://gitlab.com/commonground/nlx/nlx/commit/d2dfb6ade661f96e7cc0d7ead53df8549d4b59d3))
* **inway:** proxy management-api requests ([9c3de64](https://gitlab.com/commonground/nlx/nlx/commit/9c3de64d9c13649ae31abbfef24ffa8f5c751b4c))
* **management:** add access grants ([d5344f6](https://gitlab.com/commonground/nlx/nlx/commit/d5344f6c06e885c74025b09162d35deeca13e6cc))
* **management:** add public key fingerprint to outgoing access-request ([c465414](https://gitlab.com/commonground/nlx/nlx/commit/c46541457ed27c70a0114da8f47694905fbf4cdf))
* **management:** add update access request to the database ([918ad7b](https://gitlab.com/commonground/nlx/nlx/commit/918ad7bca8558fadb204060a4e4c3a290116bf94))
* **management:** emit empty fields in JSON responses ([24ab50e](https://gitlab.com/commonground/nlx/nlx/commit/24ab50efbceef9746f5daa72eb8b317b630040ec))
* **management:** implement access request status loop ([8343712](https://gitlab.com/commonground/nlx/nlx/commit/834371240ebb2502ad3546f7143860d0a200f17b))
* **management:** implement locking for access requests ([6147c52](https://gitlab.com/commonground/nlx/nlx/commit/6147c5252af318d70d7f0987a6e22bde2af873be))
* **management:** implement management external API ([339be14](https://gitlab.com/commonground/nlx/nlx/commit/339be14013341b0d14a42d34fb3bf7a3f3ab302a))
* **management:** implement management gRPC client ([410c001](https://gitlab.com/commonground/nlx/nlx/commit/410c00146af3dbf6a7099cb49ba7a385042d8b71))
* **management:** re-implement unique constraint for access-request ([b07da55](https://gitlab.com/commonground/nlx/nlx/commit/b07da553b8a625b3b29a6cc5046947393541dd38))
* **management:** remove whitelist configuration ([662008a](https://gitlab.com/commonground/nlx/nlx/commit/662008ac1dda87f46f01988cc08edc5b663ab747)), closes [#1008](https://gitlab.com/commonground/nlx/nlx/issues/1008)
* **management:** return incomming access requests ([c3bab34](https://gitlab.com/commonground/nlx/nlx/commit/c3bab34b07f1ffe11886dc5b16557d401e3388a6))
* **management:** save organization inway address in the directory ([6775c62](https://gitlab.com/commonground/nlx/nlx/commit/6775c622f851c0464ef6b1c48369fad0d4a8abe7))
* **management:** services show if there are no access requests ([5a5679e](https://gitlab.com/commonground/nlx/nlx/commit/5a5679e0cfa07042ac29ff678f6d712fcc7aea17))


### Reverts

* update dependency styled-components to v5.2.0 ([e749b37](https://gitlab.com/commonground/nlx/nlx/commit/e749b372c2207e458c1545b57aec04afd6a83d71))
* **management:** enable specifying OIDC & Management API base URLs ([a37f0a1](https://gitlab.com/commonground/nlx/nlx/commit/a37f0a1dfcd49d7f5ceb0960a2d4263f01e2bc88))

# [0.88.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.87.1...v0.88.0) (2020-09-17)


### Bug Fixes

* **insight:** close HTML-elements correctly ([40f1424](https://gitlab.com/commonground/nlx/nlx/commit/40f1424e48a66d579b96f9ae00fbcab3dd0fbaf7))
* **management:** linting ([5811a71](https://gitlab.com/commonground/nlx/nlx/commit/5811a71205bebb77f2d18ec70aa8a412751a5e54)), closes [#1002](https://gitlab.com/commonground/nlx/nlx/issues/1002)
* console warning mobx-react-lite ([c48143a](https://gitlab.com/commonground/nlx/nlx/commit/c48143ace47d4337b2bdb9669edb8eb5f4ebaee7))
* **directory:** linting issues ([2dd865f](https://gitlab.com/commonground/nlx/nlx/commit/2dd865fe1e8e449bd4cdff52d2174a2e417bdee5)), closes [#973](https://gitlab.com/commonground/nlx/nlx/issues/973)
* **docs:** update broken link for BAG API ([a679887](https://gitlab.com/commonground/nlx/nlx/commit/a679887d2fede1c472c705a700db0d6ded88a8b6))
* **helm:** regenerate Chart.lock ([09dab6b](https://gitlab.com/commonground/nlx/nlx/commit/09dab6b53b2806eec50edb9f3e3136070aa65924))
* **helm:** regenerate invalid digest for Chart.lock ([08c5aeb](https://gitlab.com/commonground/nlx/nlx/commit/08c5aeb9c5ba5d2ac8b37d8f7d5ab4b710018a2e))
* **management:** local e2e tests need to work with modd again ([750e5ad](https://gitlab.com/commonground/nlx/nlx/commit/750e5ad1c3acebe34dbef94fd81aaf6625e4775b))
* **management:** provide service model defaults ([cab1397](https://gitlab.com/commonground/nlx/nlx/commit/cab139711c8c92b1856352751df969cecccb25ef))
* **management:** resolve typo ([f82ea69](https://gitlab.com/commonground/nlx/nlx/commit/f82ea6951e82b5e1c1cc061d180bd80e66605fb4)), closes [#973](https://gitlab.com/commonground/nlx/nlx/issues/973)
* **management:** set nlxversion on directory client ([ce96f09](https://gitlab.com/commonground/nlx/nlx/commit/ce96f0947e570caf94ba85d13f88243d84cf04b2))
* **management:** type-o in settings ([eb7ecab](https://gitlab.com/commonground/nlx/nlx/commit/eb7ecab421bfeaeb63eafd0c617a9b1222761059))


### Features

* **common:** remove non-Linux OS warning ([657635e](https://gitlab.com/commonground/nlx/nlx/commit/657635e8aedcdeb22e4355b2f96fa600667866f3))
* **directory:** set inway for an organization ([eb5ac76](https://gitlab.com/commonground/nlx/nlx/commit/eb5ac769048bddf950f406b4496d1dcfd1498b3a))
* **helm:** remove option to generate certificates ([366a019](https://gitlab.com/commonground/nlx/nlx/commit/366a019f4f0b236a27cd7e349f7e141b47ca3a42))
* **helm:** use prebuild Docker image for the IRMA server ([b54eb15](https://gitlab.com/commonground/nlx/nlx/commit/b54eb150f593c25e2293e61c7f78b122263f4c87))
* **management:** add all inway properties to store ([e2e4c38](https://gitlab.com/commonground/nlx/nlx/commit/e2e4c38d82a0112847b9e7e2dda06ccb71d0de35)), closes [#1002](https://gitlab.com/commonground/nlx/nlx/issues/1002)
* **management:** add description for the organisation outway field ([1dde860](https://gitlab.com/commonground/nlx/nlx/commit/1dde860ef8f27517b1ce121a18562f876118177d))
* **management:** add mobx store to services ([b3e2377](https://gitlab.com/commonground/nlx/nlx/commit/b3e2377969401a8a00beb7c4d398910c70ae6cc2))
* **management:** enable setting the inway for the management api traffic ([329f6d4](https://gitlab.com/commonground/nlx/nlx/commit/329f6d45f289479ad2e57e99ec3b310c28bcaf49)), closes [#973](https://gitlab.com/commonground/nlx/nlx/issues/973)
* **management:** enable specifying OIDC & Management API base URLs ([26f1f90](https://gitlab.com/commonground/nlx/nlx/commit/26f1f9083b9aed138fc54bd0b5526a915fa44300))
* **management:** external API for access requests ([c349b1e](https://gitlab.com/commonground/nlx/nlx/commit/c349b1e81cd92c9ab6a3b3ccf27c6c4baf7908d5))
* **management:** implement inway model ([d24f919](https://gitlab.com/commonground/nlx/nlx/commit/d24f9193b6922ff444532c0d5add74487f766329))
* **management:** introduce InwaysStore ([d9c7f3c](https://gitlab.com/commonground/nlx/nlx/commit/d9c7f3c16a255fc1330164adba1ee0aad67f913d)), closes [#1002](https://gitlab.com/commonground/nlx/nlx/issues/1002)
* **management:** replace http client with gRPC client ([3c2e1fe](https://gitlab.com/commonground/nlx/nlx/commit/3c2e1fef05c5c2b084654ecd20e44c13a55fc545))
* rename ConfigAPI service to Management service ([8d3f264](https://gitlab.com/commonground/nlx/nlx/commit/8d3f2640eeadb6e4a86fb9a75a20b58a0a8a1e02))
* use different certificate for mangement-api ([582e4d2](https://gitlab.com/commonground/nlx/nlx/commit/582e4d2fae89f536099aca9f738e41fbf31eebc8))

## [0.87.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.87.0...v0.87.1) (2020-08-10)


### Bug Fixes

* **helm:** directory nil pointer on terminationPolicy value ([6c95e86](https://gitlab.com/commonground/nlx/nlx/commit/6c95e86be631a80ca4fa13365030ac07b8cc9537))

# [0.87.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.86.0...v0.87.0) (2020-08-10)


### Bug Fixes

* **directory-inspection-api:** add /api to routes ([18891df](https://gitlab.com/commonground/nlx/nlx/commit/18891df5e72f5584d1b3803ef148f61b6e92ad4c))
* **directory-inspection-api:** expose under /api ([47c9e68](https://gitlab.com/commonground/nlx/nlx/commit/47c9e68047dd8b97a210e6ce9413cd63618b1066))
* **directory-ui:** enable IE11 support ([60ab3ec](https://gitlab.com/commonground/nlx/nlx/commit/60ab3ec817daec1e551cfd291750580c4030cbee))
* **docs:** contact page ([0106685](https://gitlab.com/commonground/nlx/nlx/commit/01066857eb076c9e776ae8eef89d7e3e727b3a71))
* **docs:** update broken docusaurus config ([7937933](https://gitlab.com/commonground/nlx/nlx/commit/79379336759b450b3f3e3060f0575577f73d3ae8))
* **docs:** update swizzled components ([6c568af](https://gitlab.com/commonground/nlx/nlx/commit/6c568afd9a05be64d12120f5fc5085c8eb619438))
* **helm:** delete job before install,upgrade ([edff902](https://gitlab.com/commonground/nlx/nlx/commit/edff902c76d9a35ee7213a62b01ddb3581cddd98))
* **helm:** linting errors ([4eeaa94](https://gitlab.com/commonground/nlx/nlx/commit/4eeaa94bddf7d4270141aded5f97e9a2707e3f6d))
* **helm:** nlx-directory: use existingSecret value ([d61c5f3](https://gitlab.com/commonground/nlx/nlx/commit/d61c5f3bbded02bdf4b81e18f1476d5fde65a417))
* **helm:** nlx-management: allow setting tls hostnames ([83727e4](https://gitlab.com/commonground/nlx/nlx/commit/83727e44c3cb53dcb763259239fc904cfa39d83d))
* **helm:** run nlxctl job after database migration ([9e3dc52](https://gitlab.com/commonground/nlx/nlx/commit/9e3dc52f493a25042e811188a765486ca1b787fc))
* **helm:** use replicaCount value ([769671e](https://gitlab.com/commonground/nlx/nlx/commit/769671e8d457dbabca23f6b336449fefb626a8fb))
* **helm:** use root certificate value ([8e65cd2](https://gitlab.com/commonground/nlx/nlx/commit/8e65cd2ba1118988cdb205ee2a159c210e68e1de))
* **management-ui:** adjust test to assert if Verzoek button is visible ([712d5d1](https://gitlab.com/commonground/nlx/nlx/commit/712d5d1201c0247574d6e07a3b7d20b58a6311ca)), closes [#918](https://gitlab.com/commonground/nlx/nlx/issues/918)
* **management-ui:** api call and statuses ([b979c19](https://gitlab.com/commonground/nlx/nlx/commit/b979c19ed8e6b683a9adab6fdfa19792f8be56df))
* app not working in IE11 ([8c36212](https://gitlab.com/commonground/nlx/nlx/commit/8c362121e14c694edf30d76870c4b5a70dd8278a))
* aria attributes in Collapsible ([a920412](https://gitlab.com/commonground/nlx/nlx/commit/a92041261279590e1c44b5741c431af2d3998da2))
* removed e2e test in unit tests that gave error ([91b5ecb](https://gitlab.com/commonground/nlx/nlx/commit/91b5ecbf41c4ae935c96974520f689dbab15888f))
* **insight-ui:** don't strip /api ([d633c37](https://gitlab.com/commonground/nlx/nlx/commit/d633c37898fd87d918223ee56c9ab0cad69ccc87))
* **management-ui:** added translations ([f75bb1c](https://gitlab.com/commonground/nlx/nlx/commit/f75bb1c06ae64ef2d3dde568812009d870076f57))
* **management-ui:** allow to create service with an empty whitelist ([29e196a](https://gitlab.com/commonground/nlx/nlx/commit/29e196a801b13c2b5447cd378725ecbd70fa89c2))
* **management-ui:** authorization -> access ([dea52ac](https://gitlab.com/commonground/nlx/nlx/commit/dea52acbdeeae3c82bdee2009853bcaa0175cf00))
* **management-ui:** correctly submit authorization mode [#896](https://gitlab.com/commonground/nlx/nlx/issues/896) ([ca570d1](https://gitlab.com/commonground/nlx/nlx/commit/ca570d1f5963c3af1f743df1b924b82b22214d0c))
* **management-ui:** correctly submit authorizations [#896](https://gitlab.com/commonground/nlx/nlx/issues/896) ([32e8e49](https://gitlab.com/commonground/nlx/nlx/commit/32e8e4961950a8e2589cb302ad56bc88a2eecf17))
* **management-ui:** e2e ie11 ([7a74bf0](https://gitlab.com/commonground/nlx/nlx/commit/7a74bf077dacddf7e307a45f955f80dea68affc9))
* **management-ui:** include Array.entries polyfill for IE11 ([68f4afb](https://gitlab.com/commonground/nlx/nlx/commit/68f4afb52d4d953689b5b7c1da27f0d6debf07f5))
* **management-ui:** include Object.entries polyfill for IE11 ([0235341](https://gitlab.com/commonground/nlx/nlx/commit/02353417358f6fe1f11bdd521145c850513cfdef)), closes [#942](https://gitlab.com/commonground/nlx/nlx/issues/942)
* possible fix for e2e test ([833d6d4](https://gitlab.com/commonground/nlx/nlx/commit/833d6d4ea059f45ec9f4e5ffbfbcc827dd3cb7b8))
* service icon preview in finder ([34161f2](https://gitlab.com/commonground/nlx/nlx/commit/34161f23e7b39c6c9d42319a25410dc3e41b4907))
* tweaks and processed feedback ([fd89ca9](https://gitlab.com/commonground/nlx/nlx/commit/fd89ca9604cd9e982a535bf868f8968244eb31b0))
* undo e2e test changes ([d9cb608](https://gitlab.com/commonground/nlx/nlx/commit/d9cb608253b586f8ac92ee25b7637580dd8d5cf9))
* **helm:** postgres storage type default to 'Durable' ([1c2668b](https://gitlab.com/commonground/nlx/nlx/commit/1c2668bad848d9617f59a4b3a44f53dea1f06405))
* **helm:** run database migrations post-upgrade ([0189234](https://gitlab.com/commonground/nlx/nlx/commit/0189234eadbd7f26097242e99172777a80857cdb))
* **insight-ui:** navigation bar urls ([862fb2b](https://gitlab.com/commonground/nlx/nlx/commit/862fb2bf69f02c2106c82a34d5df0205cd21d145))
* **inway:** replace services when deleted from management-api ([8957524](https://gitlab.com/commonground/nlx/nlx/commit/8957524f690066c6197e797745b806d9a189c7b1))
* **management-ui:** increase top padding of page ([826d0d2](https://gitlab.com/commonground/nlx/nlx/commit/826d0d250119a650777fc33625efcfbd48c774c6))
* **management-ui:** logout should work ([964be20](https://gitlab.com/commonground/nlx/nlx/commit/964be209b2fa9206ba54931797b053ef92eb5967))
* **management-ui:** logout works in all browsers ([3c52c5e](https://gitlab.com/commonground/nlx/nlx/commit/3c52c5e877b326e09c3a129652e0d90e1728bc90))
* **management-ui:** make sure close button is always visible ([e5b1caf](https://gitlab.com/commonground/nlx/nlx/commit/e5b1caf4603c8285d0257ffd9e34e46ce173492a))
* **management-ui:** prevent caching of current user api call ([abb79e2](https://gitlab.com/commonground/nlx/nlx/commit/abb79e2fbd257313365ef68f677088a730786314))
* **management-ui:** remove dot from login button text ([6e6edda](https://gitlab.com/commonground/nlx/nlx/commit/6e6edda9d3df5412d0bf56067fae995f06635db4))
* **management-ui:** several fixes related to inway details ([82cd8f6](https://gitlab.com/commonground/nlx/nlx/commit/82cd8f626e399fb09263b42efb2b6b299fc8cbb2))
* **management-ui:** update feedback when loading inways ([f5564b6](https://gitlab.com/commonground/nlx/nlx/commit/f5564b6b71e5db84c16d48ee04b55981a3b71d2d)), closes [#923](https://gitlab.com/commonground/nlx/nlx/issues/923)
* **management-ui:** update service detail link ([2986e42](https://gitlab.com/commonground/nlx/nlx/commit/2986e4274370d45df97f27c3250a11e3f95a826b)), closes [#923](https://gitlab.com/commonground/nlx/nlx/issues/923)
* **management-ui:** use correct spacing around user menu ([f3c2260](https://gitlab.com/commonground/nlx/nlx/commit/f3c22601191a103d1ddaccdfa807e38a2fcbc940))
* **management-ui:** usePromise hook re-calling didn't reset error ([b97d142](https://gitlab.com/commonground/nlx/nlx/commit/b97d142f447efdb31bd7e6cad797ec45e7680b26))
* icon color for non-active menu items [#896](https://gitlab.com/commonground/nlx/nlx/issues/896) ([b5a110a](https://gitlab.com/commonground/nlx/nlx/commit/b5a110a82e86bb077c1f6d630fd5e04baa5432e4))
* translate validation messages [#896](https://gitlab.com/commonground/nlx/nlx/issues/896) ([6ce2d15](https://gitlab.com/commonground/nlx/nlx/commit/6ce2d157438def2fc83d6b4e6b42db6d24226a55))


### Features

* added mobx versions, removed ie from e2e tests because of Proxy ([1f31271](https://gitlab.com/commonground/nlx/nlx/commit/1f31271ffea0cbfd36266adb19e631300ba3840e))
* directory detail uses mobx store ([155e270](https://gitlab.com/commonground/nlx/nlx/commit/155e270aba72566dfccd50f1cf622d081bf4cddc))
* directory list now uses mobx ([5c7a580](https://gitlab.com/commonground/nlx/nlx/commit/5c7a5801df21d4697b9b36b73b300790589bdc7a))
* **common:** use base64 for SPKI fingerprint ([df8256a](https://gitlab.com/commonground/nlx/nlx/commit/df8256a609c4a820713f3f171022a90d31ac400a))
* **config-api:** add validation for authorization settings in service ([5dab61d](https://gitlab.com/commonground/nlx/nlx/commit/5dab61d509144ebc113d2c608024f7f83b283046))
* **config-api:** remove config-api component ([7d75dda](https://gitlab.com/commonground/nlx/nlx/commit/7d75dda8c0af7abff7b12e617a7a4b8eb510d6f3))
* **helm:** add directory endpoint url value ([6b4425d](https://gitlab.com/commonground/nlx/nlx/commit/6b4425da62fe5e10280460f6a8cbac405fc4ffc1))
* **helm:** add Helm v3 charts ([d90c230](https://gitlab.com/commonground/nlx/nlx/commit/d90c230ac25de52e9902e9976f9c654b525f72d1))
* **helm:** config API -> management API ([39c3cd1](https://gitlab.com/commonground/nlx/nlx/commit/39c3cd1afad64adc43f6bc4c8e9a4618e6b17a16))
* **helm:** outway: https support ([8bdb304](https://gitlab.com/commonground/nlx/nlx/commit/8bdb304aa746b29a4cbe9d7c38434189e7ae30bf))
* **helm:** remove config-api component ([60fea94](https://gitlab.com/commonground/nlx/nlx/commit/60fea94067ba6ec356df12a2ff2990ea54fbe50d))
* **inway:** selfAddress must be present in the certificate ([6388118](https://gitlab.com/commonground/nlx/nlx/commit/638811856db89248676373bec7d7aceb873862c4))
* **inway:** use the management API instead of the config API ([6d79ad7](https://gitlab.com/commonground/nlx/nlx/commit/6d79ad77fdd7384e523a4ac3cb95a4d5120c562a))
* **management-api:** add directory proto ([00c8f6b](https://gitlab.com/commonground/nlx/nlx/commit/00c8f6bed4be297553b78863059e91731cef8e12))
* **management-api:** add directory services endpoint ([ce1440e](https://gitlab.com/commonground/nlx/nlx/commit/ce1440ec65db0f627b27b479c6ed6d09c592fa1d))
* **management-api:** add environment endpoint ([41fa8a1](https://gitlab.com/commonground/nlx/nlx/commit/41fa8a135415c931f1e9227197ad6ef9650ab850))
* **management-api:** add ip address to inway data in etcd ([a27b670](https://gitlab.com/commonground/nlx/nlx/commit/a27b6704bfc5d248bc8c5bb169d5b2253171bbbe))
* **management-api:** add services to getinway ([d786d60](https://gitlab.com/commonground/nlx/nlx/commit/d786d6084a91af0c6b8e3850f0f49a39a260ec85))
* **management-api:** bring in the config api ([089da4b](https://gitlab.com/commonground/nlx/nlx/commit/089da4bf6d5581f4073a16c50e05403253250c45))
* **management-api:** connect directory service to grpc ([51b8a32](https://gitlab.com/commonground/nlx/nlx/commit/51b8a32210cffabbfded9207318baf6f04ec46b2))
* **management-api:** persist access requests ([8505384](https://gitlab.com/commonground/nlx/nlx/commit/8505384666446828d9aff55f8288b8bf80a86993))
* **management-ui:** add directory page ([91e0342](https://gitlab.com/commonground/nlx/nlx/commit/91e03427c15e8b10cbdc4f45f4446028ecc7949a))
* **management-ui:** add edit service ([03d5607](https://gitlab.com/commonground/nlx/nlx/commit/03d56074cb53d828625d3cb02805df1a53b86bee))
* **management-ui:** add empty state on collapsible ([aa9fbf5](https://gitlab.com/commonground/nlx/nlx/commit/aa9fbf5e9bd45e4f6abbda26f5704a7fdf8fcadd))
* **management-ui:** add inway selection to service form ([6320f19](https://gitlab.com/commonground/nlx/nlx/commit/6320f19b4ac0963fdd0a6e4d8627009e3e57105d))
* **management-ui:** add organization name to the page header ([03b4bf5](https://gitlab.com/commonground/nlx/nlx/commit/03b4bf54586b1d21f9ee19dd817eb691e9c7789a))
* **management-ui:** add overview of Inways ([6604667](https://gitlab.com/commonground/nlx/nlx/commit/6604667562ec30ac5b497f316af13a349a117087))
* **management-ui:** add service details page ([c72204a](https://gitlab.com/commonground/nlx/nlx/commit/c72204a11e78bd0d705b585064ea43a7424797e8))
* **management-ui:** directory detail has access request status ([4572c62](https://gitlab.com/commonground/nlx/nlx/commit/4572c626a966d2f12368ac5cbdecfef56090f4d5))
* **management-ui:** directory detail with actual data and mock request button ([35e10f9](https://gitlab.com/commonground/nlx/nlx/commit/35e10f9e278638f839a78bd427284984be57d394))
* **management-ui:** display inways after login ([b88601d](https://gitlab.com/commonground/nlx/nlx/commit/b88601d32dc3f1adc80cb3b340c1b8148907bd97))
* **management-ui:** display services count for inways ([55d3a9d](https://gitlab.com/commonground/nlx/nlx/commit/55d3a9d006d0a289b0e23401aaa12f8c6a2acb1b))
* **management-ui:** display toaster on adding service ([aa986db](https://gitlab.com/commonground/nlx/nlx/commit/aa986db644443c6870afe11f028629bd82bb821b))
* **management-ui:** display toaster when editing a service ([671ff24](https://gitlab.com/commonground/nlx/nlx/commit/671ff24cdccb9f6e9e9fa7b7f695bb4bb3a461f6)), closes [#942](https://gitlab.com/commonground/nlx/nlx/issues/942)
* **management-ui:** display toaster when removing a service ([40072b9](https://gitlab.com/commonground/nlx/nlx/commit/40072b9e85f82fa5b85efc687e26cc4e7a1a471e)), closes [#942](https://gitlab.com/commonground/nlx/nlx/issues/942)
* **management-ui:** implement AddServiceForm [#896](https://gitlab.com/commonground/nlx/nlx/issues/896) ([56add2d](https://gitlab.com/commonground/nlx/nlx/commit/56add2d37dd4d84ab019eb8fb44980a7ae07903a))
* **management-ui:** implement AddServicePage [#896](https://gitlab.com/commonground/nlx/nlx/issues/896) ([7390cad](https://gitlab.com/commonground/nlx/nlx/commit/7390cad43d5bbac28b7b5a9f0111d2ef683f7716))
* **management-ui:** proper failed message in directory detail view ([f7e7579](https://gitlab.com/commonground/nlx/nlx/commit/f7e75795fc2b3a41115fe3243bb17db70fe636f4))
* **management-ui:** remove asserting the modals as part of the E2E-tests ([fe50fb0](https://gitlab.com/commonground/nlx/nlx/commit/fe50fb03cdd56f670adeefc7c6feb31d6c90710d)), closes [#942](https://gitlab.com/commonground/nlx/nlx/issues/942)
* **management-ui:** request access on directory page ([a855a49](https://gitlab.com/commonground/nlx/nlx/commit/a855a4959dc4f6e0df4826b39321324901472e79))
* add access request endpoints ([b949daf](https://gitlab.com/commonground/nlx/nlx/commit/b949daf8ae9b874417af0b6889098ab930838986))
* **management-ui:** disable mask for edit service drawer ([958745b](https://gitlab.com/commonground/nlx/nlx/commit/958745b8140f4c5639147fb2d119309eecdb994e))
* **management-ui:** improve ui ([8044822](https://gitlab.com/commonground/nlx/nlx/commit/80448229874f64b4b84cef37085fc3611f02d603))
* **management-ui:** initial routes and mock detailview directory ([5e70ccc](https://gitlab.com/commonground/nlx/nlx/commit/5e70ccce7c35021860efa8626c9537c94985f3cb))
* **management-ui:** inway details page v1 ([9d401d8](https://gitlab.com/commonground/nlx/nlx/commit/9d401d8151290043f48f178bda393059e0fbb225))
* **management-ui:** return empty list when no services are present [#896](https://gitlab.com/commonground/nlx/nlx/issues/896) ([597904a](https://gitlab.com/commonground/nlx/nlx/commit/597904aa101adced64717abaabc3bc6f19dea8ba))
* **management-ui:** scroll the first error message into view ([11a45f1](https://gitlab.com/commonground/nlx/nlx/commit/11a45f19acb1123afcf9e7a6147c6bd6f8d903cc))
* **management-ui:** services in inway details and linking between each other ([40faa53](https://gitlab.com/commonground/nlx/nlx/commit/40faa534833b92249826c9a7b6280914dfd73025))
* **management-ui:** show confirm box when requesting access ([79ca23d](https://gitlab.com/commonground/nlx/nlx/commit/79ca23d7a05085fd442d9242d39d2d380964f7b6))
* **management-ui:** show warning about missing connection between service and inway ([3b41466](https://gitlab.com/commonground/nlx/nlx/commit/3b41466e424eade964a604b51209209a5f1aceda))
* **management-ui:** status indicator optionally shows status text ([b0e00cc](https://gitlab.com/commonground/nlx/nlx/commit/b0e00cc80d5e4ea5be6a87b1764e709594db96b7))
* **management-ui:** ui tweaks [#896](https://gitlab.com/commonground/nlx/nlx/issues/896) ([ddfcfa1](https://gitlab.com/commonground/nlx/nlx/commit/ddfcfa1556d487f14d5daacef3e17cf2efac5a0e))
* **management-ui:** use Drawer.Header from Design System ([7c3105a](https://gitlab.com/commonground/nlx/nlx/commit/7c3105ad5f48fdde11b70ba2eb92c3dc6ebe9039))
* **management-ui:** user navigation slide open animation ([94d3996](https://gitlab.com/commonground/nlx/nlx/commit/94d399652471a94cdebedd78b03156da76288c07))
* **outway:** use different certificate for https endpoint ([192c1d3](https://gitlab.com/commonground/nlx/nlx/commit/192c1d3e682af0c272dd8ae054b693e56eb01f0d))
* display user info in user navigation ([b0158d7](https://gitlab.com/commonground/nlx/nlx/commit/b0158d71ddb3b30ff6a9debbb968054e3640cb78))
* **outway:** add option to function as a proxy ([7b60f4b](https://gitlab.com/commonground/nlx/nlx/commit/7b60f4b914d980e8fe00d762bf68d9a2da730e2b))


### Performance Improvements

* small improvements service form ([bed44d1](https://gitlab.com/commonground/nlx/nlx/commit/bed44d11b684ed9c1a6c4013b1e904c0c7ea3150))

# [0.86.0](https://gitlab.com/commonground/nlx/nlx/compare/v0.85.1...v0.86.0) (2020-04-02)


### Bug Fixes

* **deps:** update dependency @testing-library/react to v10 ([955d0fd](https://gitlab.com/commonground/nlx/nlx/commit/955d0fd9687d4d08f5a9afc63bd7519935c8969f))
* **docs:** remove link to community ([05884f0](https://gitlab.com/commonground/nlx/nlx/commit/05884f095b890370a0d0d13c08bfb1b55bbaa6d6))
* **helm:** indentation of metadata labels ([65b2e1e](https://gitlab.com/commonground/nlx/nlx/commit/65b2e1e66352ce49c1226376191efbb5e272dc53))
* **management-ui:** added translations for 404 page ([1492d23](https://gitlab.com/commonground/nlx/nlx/commit/1492d23d498f6c2f887e8d4318882e4a78ba433c))
* **management-ui:** processed feedback ([30546c3](https://gitlab.com/commonground/nlx/nlx/commit/30546c3396721d613083ef3fc381ba76ee3d39f6))
* **management-ui:** rebase fixes and additions ([e7ea3da](https://gitlab.com/commonground/nlx/nlx/commit/e7ea3da378751e90208eb882ac5b2fe2eebc46d9))
* **management-ui:** review feedback ([2e6d1aa](https://gitlab.com/commonground/nlx/nlx/commit/2e6d1aa6ac332960dd2abe7362fcb61999fb5817))
* **management-ui:** unit test ([6a18a7a](https://gitlab.com/commonground/nlx/nlx/commit/6a18a7a21f012814d5d4702fa8c101468b8e3060))
* **nlx-management:** i18next dependency [#873](https://gitlab.com/commonground/nlx/nlx/issues/873) ([2817bac](https://gitlab.com/commonground/nlx/nlx/commit/2817bac6ba18c8ca00b188fd6405de5e411edcde))


### Features

* **config-api:** extend whitelist with publicKeyHash ([a888cc7](https://gitlab.com/commonground/nlx/nlx/commit/a888cc703e54a548cb8e04f931db0278bf8539f0))
* **inway:** create updated version of serviceconfig.toml with backward compatibility ([f2a4765](https://gitlab.com/commonground/nlx/nlx/commit/f2a476501936923132385969f8a3d86ece6e809b))
* **inway:** create updated version of serviceconfig.toml with backward compatibility ([b140fa1](https://gitlab.com/commonground/nlx/nlx/commit/b140fa1366e06e79456c4e93e028614cdd8aa5c2))
* **inway:** use whitelisted subjectPublicKeyInfo to allow outways through ([6e452de](https://gitlab.com/commonground/nlx/nlx/commit/6e452de9e75b86993e9ec74c2d5b3bf975f15845))
* **inway:** use whitelisted subjectPublicKeyInfo to allow outways through ([b711e51](https://gitlab.com/commonground/nlx/nlx/commit/b711e510e13a1edb01c2113fe8f457d6da2f4e3b))
* **management-api:** add OpenID Connect support ([b7fc893](https://gitlab.com/commonground/nlx/nlx/commit/b7fc893a65aa2442fe2542bf94f8afcd0baf5bca))
* **management-ui:** add OpenID Connect support ([2470a01](https://gitlab.com/commonground/nlx/nlx/commit/2470a01d8e83f485ce1a85cab66fb3f0c7d53fd9))
* **management-ui:** added 404 page ([9b2ad3a](https://gitlab.com/commonground/nlx/nlx/commit/9b2ad3aaa4c9cade3e613894bb73fae09eca3476))
* **management-ui:** added translation files ([042f77c](https://gitlab.com/commonground/nlx/nlx/commit/042f77c8169f7b3fb53bce065dc7d1576bc4d4f0))
* **management-ui:** adjust outline for PrimaryNavigation buttons [#876](https://gitlab.com/commonground/nlx/nlx/issues/876) ([37536ed](https://gitlab.com/commonground/nlx/nlx/commit/37536ed6795e70a9c6dfed1c2794c49777361f6b))
* **management-ui:** display no results message [#881](https://gitlab.com/commonground/nlx/nlx/issues/881) ([dade50c](https://gitlab.com/commonground/nlx/nlx/commit/dade50c34164e9ec0129f7985e026dc38bdcdac5))
* **management-ui:** e2e test NotFound page ([a0d3f40](https://gitlab.com/commonground/nlx/nlx/commit/a0d3f4074d99c4473385b91fcea56a55b490e197))
* **management-ui:** e2e tests render i18n keys instead of translations ([d97d96b](https://gitlab.com/commonground/nlx/nlx/commit/d97d96be0944294acba02c66279cfe5c2bf4b75c))
* **management-ui:** implement Services page [#876](https://gitlab.com/commonground/nlx/nlx/issues/876) ([5ff96ad](https://gitlab.com/commonground/nlx/nlx/commit/5ff96ad93ab6ce0ecb21e3dd1655e68dd97c743d))
* **management-ui:** implement Services page [#881](https://gitlab.com/commonground/nlx/nlx/issues/881) ([2c9760e](https://gitlab.com/commonground/nlx/nlx/commit/2c9760efbd13dedb24ce9c2d1d4b45bb86ee21cb))
* **management-ui:** translate aria labels [#876](https://gitlab.com/commonground/nlx/nlx/issues/876) ([6eaa6d0](https://gitlab.com/commonground/nlx/nlx/commit/6eaa6d0ecf2350df3c8c06e21cb2d629ccfcaed5))
* **nlx-management:** add background pattern to Welcome page [#872](https://gitlab.com/commonground/nlx/nlx/issues/872) ([549e7fe](https://gitlab.com/commonground/nlx/nlx/commit/549e7feabee4e9036789df29907004b34d7c6403))
* **nlx-management:** add initial setup using Create React App [#873](https://gitlab.com/commonground/nlx/nlx/issues/873) ([96b97ed](https://gitlab.com/commonground/nlx/nlx/commit/96b97ed9fdb1320afa0959c9d64984d7b12506da))
* **nlx-management:** implement global styles using Styled Components [#873](https://gitlab.com/commonground/nlx/nlx/issues/873) ([06d2a8d](https://gitlab.com/commonground/nlx/nlx/commit/06d2a8d644e4b4ec0d57c380c830eff0b7f34988))
* **nlx-management:** implement routing + redirect logic from / to /login [#873](https://gitlab.com/commonground/nlx/nlx/issues/873) ([7757b93](https://gitlab.com/commonground/nlx/nlx/commit/7757b933916b29c5711e85bb641ea0e9882a0b85))
* **nlx-management:** implement Welcome page [#873](https://gitlab.com/commonground/nlx/nlx/issues/873) ([b290616](https://gitlab.com/commonground/nlx/nlx/commit/b2906165bfddf491c707c5956d17d508619c6425))
* **nlx-management:** setup internationalization using react-i18n [#873](https://gitlab.com/commonground/nlx/nlx/issues/873) ([637843b](https://gitlab.com/commonground/nlx/nlx/commit/637843bf3c2eae8769fe4daffbd91f1592d34a98))

## [0.85.1](https://gitlab.com/commonground/nlx/nlx/compare/v0.85.0...v0.85.1) (2020-03-13)


### Bug Fixes

* **directory-db, txlog-db:** remove pgModeler, restructure database migrations, remove database roles from migrations ([d22e244](https://gitlab.com/commonground/nlx/nlx/commit/d22e24452ac7381e46398a21bb5dc37f54974dbd))

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
