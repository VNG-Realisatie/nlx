###########
## Chart ##
###########
postgresql:
  storageType: Ephemeral

################
## Sub-charts ##
################
nlx-inway:
  config:
    serviceConfig:
      services:
        demo-api:
          insight-api-url: https://insight-api-haarlem-{{DOMAIN_SUFFIX}}
          irma-api-url: https://irma-haarlem-{{DOMAIN_SUFFIX}}

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
