---
id: tls-between-inway-and-service
title: TLS between the Inway and the service
---

## Problem

The Inway is unable to validate the TLS certificate of a service because the certificate of the service is signed by an unknown certificate-authority.

## Solution Helm

The Inway Helm chart allows users to configure a certificate-authority certificate by adding the content of the certificate file to the value `tls.serviceCA.certificatePEM`.

## Solution Docker Compose

Given that you are using the Docker Compose setup of our [Try Me guide](../try-nlx/docker/introduction.md), you can mount the certificate-authority certificate by replacing the Inway volume section in Try Me Docker Compose file with the following:

```
volumes:
  - ./pki/organization/ca/root.crt:/certs/organization/nlx-pki-root.crt:ro
  - ./pki/organization/certs/org.crt:/certs/organization/nlx-pki-cert.crt:ro
  - ./pki/organization/certs/org.key:/certs/organization/nlx-pki-key.key:ro
  - ./pki/internal/ca/intermediate_ca.pem:/certs/internal/internal-root.crt:ro
  - ./pki/internal/certs/internal-cert.pem:/certs/internal/internal-cert.crt:ro
  - ./pki/internal/certs/internal-cert-key.pem:/certs/internal/internal-cert.key:ro
  - ./pki/internal/certs/internal-cert-key.pem:/certs/internal/internal-cert.key:ro
  - ./service-ca-cert.pem:/etc/ssl/certs/service-ca-cert.pem
```

The Inway configuration should look similar to the following:

```
  inway:
    image: nlxio/inway:v0.138.0
    container_name: nlx-try-me-inway
    restart: always
    ports:
      - 443:8443
      - 8443:8444
    environment:
      <<: *env
      LISTEN_ADDRESS: 0.0.0.0:8443
      LISTEN_ADDRESS_MANAGEMENT_API_PROXY: 0.0.0.0:8444
      INWAY_NAME: Inway-01
      SELF_ADDRESS: "${INWAY_SELF_ADDRESS}"
      MANAGEMENT_API_ADDRESS: management-api.try-me.nlx.local:8443
      MANAGEMENT_API_PROXY_ADDRESS: "${MANAGEMENT_API_PROXY_ADDRESS}"
      POSTGRES_DSN: "postgresql://postgres:postgres@postgres:5432/nlx_txlog?sslmode=disable&connect_timeout=2"
    volumes:
      - ./pki/organization/ca/root.crt:/certs/organization/nlx-pki-root.crt:ro
      - ./pki/organization/certs/org.crt:/certs/organization/nlx-pki-cert.crt:ro
      - ./pki/organization/certs/org.key:/certs/organization/nlx-pki-key.key:ro
      - ./pki/internal/ca/intermediate_ca.pem:/certs/internal/internal-root.crt:ro
      - ./pki/internal/certs/internal-cert.pem:/certs/internal/internal-cert.crt:ro
      - ./pki/internal/certs/internal-cert-key.pem:/certs/internal/internal-cert.key:ro
      - ./pki/internal/certs/internal-cert-key.pem:/certs/internal/internal-cert.key:ro
      - ./service-ca-cert.pem:/etc/ssl/certs/service-ca-cert.pem
```
