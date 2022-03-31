---
id: ip-addresses
title: IP-addresses
---

## Central NLX components

The central NLX components (directory) are located at the following IP addresses:

### Demo

- directory-api (`20.86.243.85:443`)

### Pre-production

- directory-api (`20.86.244.123:443`)

### Production

- directory-api (`20.76.229.234:443`)

## Your NLX components

> **Note:** you don't need to expose the Management API. It is being proxied by the Management UI. We recommend not to expose the Management UI.

### By default

The Inway must accessible on either port `443` or `8443`, the same applies to the Management API proxy.
The ports are restricted so organizations know beforehand which ports that need to be exposed on their network.

### For versions before v0.128.0

By default following ports are used:

- Inway (`443`)
- Management API Proxy (Inway port + 1)
