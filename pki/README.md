# Development/testing PKIs

PKIs for use during development. Do not use in production.


## Scripts

- `init.sh`:  (re)create new root and intermediate certificates.
- `issue.sh`:  issue new certificates. Pass `-f` to force existing certificates to be reissued.
- `fix-permissions.sh`: removes global read and write permissions from key files 
