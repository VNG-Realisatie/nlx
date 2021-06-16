###########
## Chart ##
###########

################
## Sub-charts ##
################
nlx-management:
  config:
    oidc:
      clientSecret: grGSl5W5HcKRETBr3OhmU6Tm
      discoveryURL: https://dex-gemeente-stijns-{{DOMAIN_SUFFIX}}
      redirectURL: https://nlx-management-gemeente-stijns-{{DOMAIN_SUFFIX}}/oidc/callback
      sessionSignKey: 0Xn2DBfb4L4hwN3XosbwoKZalLBU68UU
  ingress:
    hosts:
      - nlx-management-gemeente-stijns-{{DOMAIN_SUFFIX}}

dex:
  config:
    issuer: https://dex-gemeente-stijns-{{DOMAIN_SUFFIX}}
    staticClients:
      - id: nlx-management
        name: NLX Management
        secret: grGSl5W5HcKRETBr3OhmU6Tm
        redirectURIs:
          - https://nlx-management-gemeente-stijns-{{DOMAIN_SUFFIX}}/oidc/callback
  ingress:
    hosts:
      - dex-gemeente-stijns-{{DOMAIN_SUFFIX}}
