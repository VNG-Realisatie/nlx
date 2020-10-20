###########
## Chart ##
###########
postgresql:
  storageType: Ephemeral

managementAPI:
  insight:
    insightAPIURL: https://insight-api-haarlem-{{DOMAIN_SUFFIX}}
    irmaServerURL: https://irma-haarlem-{{DOMAIN_SUFFIX}}

################
## Sub-charts ##
################
nlx-management:
  config:
    oidc:
      clientSecret: grGSl5W5HcKRETBr3OhmU6Tm
      discoveryURL: https://dex-haarlem-{{DOMAIN_SUFFIX}}
      redirectURL: https://nlx-management-haarlem-{{DOMAIN_SUFFIX}}/oidc/callback
      sessionSignKey: 0Xn2DBfb4L4hwN3XosbwoKZalLBU68UU
  ingress:
    hosts:
      - nlx-management-haarlem-{{DOMAIN_SUFFIX}}

dex:
  config:
    issuer: https://dex-haarlem-{{DOMAIN_SUFFIX}}
    staticClients:
      - id: nlx-management
        name: NLX Management
        secret: grGSl5W5HcKRETBr3OhmU6Tm
        redirectURIs:
          - https://nlx-management-haarlem-{{DOMAIN_SUFFIX}}/oidc/callback
  ingress:
    hosts:
      - dex-haarlem-{{DOMAIN_SUFFIX}}

insight-api:
  ingress:
    hosts:
      - insight-api-haarlem-{{DOMAIN_SUFFIX}}

irma-server:
  ingress:
    annotations:
      ingress.kubernetes.io/custom-response-headers: "Access-Control-Allow-Origin: https://insight-{{DOMAIN_SUFFIX}}"
    host: irma-haarlem-{{DOMAIN_SUFFIX}}

parkeervergunning-application:
  ingress:
    hosts:
      - parkeren-haarlem-{{DOMAIN_SUFFIX}}
