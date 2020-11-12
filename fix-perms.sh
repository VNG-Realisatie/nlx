#!/bin/bash

chmod 0600 pki/organization-a/ca/intermediate-key.pem \
    pki/organization-a/ca/root-key.pem \
    pki/organization-b/ca/intermediate-key.pem \
    pki/organization-b/ca/root-key.pem \
    pki/shared/ca/intermediate-key.pem \
    pki/shared/ca/root-key.pem \
    pki/organization-a/certs/inway/key.pem \
    pki/organization-a/certs/management-api/key.pem \
    pki/organization-b/certs/inway/key.pem \
    pki/organization-b/certs/management-api/key.pem \
    pki/shared/certs/directory.shared.nlx.local/key.pem \
    pki/shared/certs/inway.organization-a.nlx.local/key.pem \
    pki/shared/certs/inway.organization-b.nlx.local/key.pem \
    testing/pki/ca-intermediate-key.pem \
    testing/pki/ca-root-key.pem \
    testing/pki/ca-root-second-key.pem \
    testing/pki/org-nlx-test-b-key.pem \
    testing/pki/org-nlx-test-key.pem \
    testing/pki/org-without-name-key.pem && \
    chmod 0644 testing/pki/org-nlx-test-key-invalid-perms.pem
