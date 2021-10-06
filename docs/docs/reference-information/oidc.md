---
id: oidc
title: OpenID Connect
---

## Usage

NLX Management uses the OpenID Connect (OIDC) protocol to authenticate users, which is
configured through a number of environment variables:

* `OIDC_CLIENT_ID`
* `OIDC_CLIENT_SECRET`
* `OIDC_DISCOVERY_URL`
* `OIDC_REDIRECT_URL`

When the management API starts, it will fetch the OIDC configuration from
`${OIDC_DISCOVERY_URL}/.well-known/openid-configuration` and use this metadata to
configure the OIDC client.

## Configuration hints for providers

There are multiple applications and cloud providers that offer OIDC.
Below are configuration hints for OIDC providers that have been used successfully with NLX.

### Dex

[Dex](https://github.com/dexidp/dex) is a federated OIDC provider.

Assuming Dex is deployed on `https://dex.example.com`, the discovery url is:

```
OIDC_DISCOVERY_URL=https://dex.example.com
```

### Azure Active Directory

Azure Active Directory is a cloud-hosted identity provider from Microsoft, part of Azure
webservices.

To use AAD as OIDC provider, you must obtain:

* Tenant ID, which is usually a UUID4
* Application (client) ID
* Application secret value

And the correct callback URL must be provided:
`https://nlx-management.example.com/oidc/callback`.

The tenant ID is used in the discovery URL:

```
OIDC_DISCOVERY_URL=https://login.microsoftonline.com/${tenantId}/v2.0
```

You can inspect the metadata document in your browser at
`https://login.microsoftonline.com/${tenantId}/v2.0/.well-known/openid-configuration`

Known issues:

* The `v2.0` is essential, since the URL without the suffix has an invalid `issuer`
  which doesn't match the tenant-specific URL.
* Make sure there's no trailing slash - the `issuer` does not have one.
