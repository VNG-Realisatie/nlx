---
id: ip-addresses
title: IP-addresses
---

## Central NLX components

The central NLX components (directory) are located at the following IP addresses:

### Demo

- directory-registration-api (`20.86.244.209:443`)
- directory-inspection-api (`20.86.243.85:443`)

### Pre-production

- directory-registration-api (`20.86.243.126:443`)
- directory-inspection-api (`20.86.244.123:443`)

### Production

- directory-registration-api (`20.86.244.12:443`)
- directory-inspection-api (`20.76.229.234:443`)

## Your NLX components

> **Note:** you don't need to expose the Management API. It is being proxied by the Management UI. We recommend not to expose the Management UI.

### By default

The following ports are used:

- Inway (`8443`)

If you use this Inway as an organization Inway, port `8444` needs to be exposed too.
This is the Management API proxy via the organization Inway (Inway port + 1)

### When following the [Try NLX Docker Compose guide](../try-nlx/docker/introduction)

The following ports are being used:

- Inway (`8443`)
- Management API proxy via organization Inway (`8444`) (Inway port + 1)

### When following the [Try NLX Helm guide](../try-nlx/helm/introduction)

The following ports are being used:

- Inway (`443`)
- Management API proxy via organization Inway (`444`) (Inway port + 1)
