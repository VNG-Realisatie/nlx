###########
## Chart ##
###########

################
## Sub-charts ##
################
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
