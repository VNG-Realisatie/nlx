# Technical notes as reference

## [#1215 - Rename current organizations](https://gitlab.com/commonground/nlx/nlx/-/issues/1215)

When renaming the demo services as part of [#1215](https://gitlab.com/commonground/nlx/nlx/-/issues/1215)
we hit the limit related to domain names.

The idea was to use `https://dex-vergunningsoftware-bv.{{DOMAIN_SUFFIX}}` instead of
`https://dex-vergunningsoftware-bv-{{DOMAIN_SUFFIX}}`,
because labels (parts of domain name separated by dots) may
not exceed the length of 63 characters.

Source https://tools.ietf.org/html/rfc1034#page-11

While working on this, we tried to request a wildcard certificate from Let's Encrypt,
which applied to `.{{DOMAIN_SUFFIX}}`. However, Let's Encrypt does not support issuing wildcard certificates with HTTP-01 challenges. To issue wildcard certificates, you must use the DNS-01 challenge.
This is not straightforward to automate, and it would take too much time until the created DNS record would resolve.
