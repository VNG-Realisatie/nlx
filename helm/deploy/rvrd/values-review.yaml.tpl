###########
## Chart ##
###########
postgresql:
  storageType: Ephemeral

################
## Sub-charts ##
################
managementAPI:
  insight:
    insightAPIURL: https://insight-api-rvrd-{{DOMAIN_SUFFIX}}
    irmaServerURL: https://irma-rvrd-{{DOMAIN_SUFFIX}}

nlx-management:
  config:
    oidc:
      clientSecret: 99DbIk7FqlUYqbyD3qSX4Wmf
      discoveryURL: https://dex-rvrd-{{DOMAIN_SUFFIX}}
      redirectURL: https://nlx-management-rvrd-{{DOMAIN_SUFFIX}}/oidc/callback
      sessionSignKey: 0Xn2DBfb4L4hwN3XosbwoKZalLBU68UU
  ingress:
    hosts:
      - nlx-management-rvrd-{{DOMAIN_SUFFIX}}

dex:
  config:
    issuer: https://dex-rvrd-{{DOMAIN_SUFFIX}}
    staticClients:
      - id: nlx-management
        name: NLX Management
        secret: 99DbIk7FqlUYqbyD3qSX4Wmf
        redirectURIs:
          - https://nlx-management-rvrd-{{DOMAIN_SUFFIX}}/oidc/callback
  ingress:
    hosts:
      - dex-rvrd-{{DOMAIN_SUFFIX}}

insight-api:
  ingress:
    hosts:
      - insight-api-rvrd-{{DOMAIN_SUFFIX}}

irma-server:
  ingress:
    annotations:
      ingress.kubernetes.io/custom-response-headers: "Access-Control-Allow-Origin: https://insight-{{DOMAIN_SUFFIX}}"
    host: irma-rvrd-{{DOMAIN_SUFFIX}}
