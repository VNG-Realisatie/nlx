---
id: new-releases
title: New releases
---

## How we release

The NLX team releases new versions roughly every two weeks, to deliver security updates, fix known issues and introduce new features. Our releases follow [SemVer](https://semver.org/) and have the structure MAJOR.MINOR.PATCH. We increment:

- MAJOR, when we make a breaking change,
- MINOR, when we add a new functionality in a backwards compatible manner, and
- PATCH, when we make backwards compatible bug fixes.

 Organizations are adviced to update their NLX components every two weeks to ensure compatibility with the central components and components of other organizations.

## Announcement

New versions are announced on [the releases page](https://gitlab.com/commonground/nlx/nlx/-/releases). It is also possible to subscribe to new version tags through [RSS](https://gitlab.com/commonground/nlx/nlx/-/tags).

## Deprecation policy

We try to minimize the effort for administrators to keep up with new NLX versions. Therefore we follow a deprecation policy for external interfaces of the NLX system. When we are not able to make the change in a backwards-compatible manner, we release a new version of the config structure or API and we keep supporting the old version for at least:

- Stable: 6 months or 4 minor releases (whichever is longer)
- Beta: 3 months or 2 minor release (whichever is longer)
- Alpha: 0 releases

In this way administrators are able to gracefuly adopt new config structures and API's. To be informed about deprecated API's, please subscribe to our mailinglist by sending a message to [support@nlx.io](mailto:support@nlx.io).
