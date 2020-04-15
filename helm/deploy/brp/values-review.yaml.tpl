################
## Sub-charts ##
################
nlx-inway:
  config:
    serviceConfig:
      services:
        basisregistratie:
          insight-api-url: https://insight-api-brp-{{DOMAIN_SUFFIX}}
          irma-api-url: https://irma-brp-{{DOMAIN_SUFFIX}}

insight-api:
  ingress:
    hosts:
      - insight-api-brp-{{DOMAIN_SUFFIX}}

irma-server:
  ingress:
    annotations:
      ingress.kubernetes.io/custom-response-headers: "Access-Control-Allow-Origin: https://insight-{{DOMAIN_SUFFIX}}"
    host: irma-brp-{{DOMAIN_SUFFIX}}
