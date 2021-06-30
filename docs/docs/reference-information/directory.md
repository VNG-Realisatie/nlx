---
id: directory
title: Directory
---

## Version Statistics

To keep track of what versions are in use within the NLX landscape the version of inways and outways are recorded.

### Inways

Every time the inway registers itself, its version is sent using the `NLX-Version` header and recorded in the `inways` table in the directory. 

### Outways

Every time the outway calls `ListServices` on the directory it sends along its version in the `NLX-Version` header.
Currently this is every 30 seconds, a live outway will generate up to 2880 records per day.
In the directory database this is recorded anonymously in the `outways` table
Only the version and timestamp are stored, no other identifying information is recorded

### Directory

The directory-inspection-api has a `/stats` endpoint that lists all known versions:
- Inways are registered and the numbers represent the real amount of inways and their versions
- Outways are unregistered and the amount represent the versions that called ListServices over the past 24 hours.
