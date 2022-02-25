---
id: organization-identification
title: Organization identification
---

## Certificate Subject Serial Number

Each certificate that is used for establishing mutual-TLS connections using NLX should have a Subject SerialNumber set. This Subject Serial Number field will be used as a primary identifier for organizations and establishing connections.

The reason why the Subject SerialNumber is chosen as the primary identifier, is because organization names can be duplicate and sometimes there can be multiple certificates within the same organization.

NLX follows the [x509 v2/v3 CRL Serial Number standards](https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.2), meaning the Subject Serial Number field has maximum of 20 octets, or else the certificate will be rejected.

The NLX applications check whether the Subject Serial Number field is not empty and it validates the length. There is no restriction built into NLX for the contents of the field, like only numbers or alphanumerical. This is to keep the possible implementation of NLX as wide as possible.

## Demo certificates

When [requesting a demo certificate](/try-nlx/docker/retrieve-a-demo-certificate), the server will add a unique Subject Serial Number based on the UNIX timestamp. The field added will be a string with 20 number characters (regex: `^\d{20}$`).

## OIN for production

An organization's [OIN (Organisatie-identificatienummer)](https://www.logius.nl/diensten/oin) will be used as the Subject Serial Number in the production environment. The production environment uses PKIoverheid certificates, this means that the Subject Serial Number field in certificates will be filled with a string containing the organization's OIN (20 number characters, in regex: `^\d{20}$`).

### Compatibility with non-Dutch standards

The first objective of NLX is to be used as a standard in The Netherlands, but the ambitions go beyond these borders. This is why the Subject Serial Number field is not limited to be only numbers (which would be a check whether it is a valid OIN), but only a length check is in place to check compliancy with the x509 standards.

Another reason why NLX does not force the use of OIN-like Subject Serial Numbers is because this is not the responsibility of NLX to check; this is the responsibility of the certificate authority. The goal of NLX is to be as open as possible. Therefore, NLX leaves specific implementation details up to the certificate authorities.
