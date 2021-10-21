---
id: data-validation
title: Data validation
---

## Organization name

The organization name in the application (extracted from the certificate) is checked to be at most 100 characters, which can be alphanumerical and whitespace characters.

The regex used is: `^[a-zA-Z0-9-._\s]{1,100}$`.

## Organization serial number

The organization serial number (extracted from the certificate), which in production is an OIN, has a maximum length of 20 bytes. No further validation on the data is implemented.

For details about organization serial numbers and organization identification, see [Organization identification](./organization-identification).

## Service name

Names given to services can have at most 100 characters, which can be alphanumerical and whitespace characters.

The regex used is: `^[a-zA-Z0-9-._\s]{1,100}$`.
