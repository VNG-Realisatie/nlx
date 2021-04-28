###########
## Chart ##
###########
postgresql:
  storageType: Ephemeral

managementAPI:
  insight:
    insightAPIURL: https://insight-api-gemeente-stijns-{{DOMAIN_SUFFIX}}
    irmaServerURL: https://irma-gemeente-stijns-{{DOMAIN_SUFFIX}}

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

insight-api:
  ingress:
    hosts:
      - insight-api-gemeente-stijns-{{DOMAIN_SUFFIX}}

irma-server:
  ingress:
    annotations:
      ingress.kubernetes.io/custom-response-headers: "Access-Control-Allow-Origin: https://insight-{{DOMAIN_SUFFIX}}"
    host: irma-gemeente-stijns-{{DOMAIN_SUFFIX}}

parkeervergunning-application:
  ingress:
    hosts:
      - parkeren-gemeente-stijns-{{DOMAIN_SUFFIX}}
