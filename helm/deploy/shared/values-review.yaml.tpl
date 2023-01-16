###########
## Chart ##
###########

postgresql:
  storageSize: 256Mi

################
## Sub-charts ##
################
nlx-directory:
  ui:
    environment: "demo"
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

apps-overview:
  config:
    environmentSubdomain: "review"
    reviewSlugWithDomain: "{{DOMAIN_SUFFIX}}"
  ingress:
    hosts:
      - "{{DOMAIN_SUFFIX}}"
