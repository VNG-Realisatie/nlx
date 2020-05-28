###########
## Chart ##
###########
postgresql:
  storageType: Ephemeral

################
## Sub-charts ##
################
nlx-directory:
  ui:
    ingress:
      hosts:
        - "directory-{{DOMAIN_SUFFIX}}"

nlx-docs:
  ingress:
    hosts:
      - "docs-{{DOMAIN_SUFFIX}}"

ca-certportal:
  ingress:
    hosts:
      - "certportal-{{DOMAIN_SUFFIX}}"

insight-ui:
  config:
    navbarDocsPageURL:  "https://docs-{{DOMAIN_SUFFIX}}"
    navbarDirectoryURL: "https://directory-{{DOMAIN_SUFFIX}}"
  ingress:
    hosts:
      - "insight-{{DOMAIN_SUFFIX}}"
