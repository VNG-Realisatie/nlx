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
