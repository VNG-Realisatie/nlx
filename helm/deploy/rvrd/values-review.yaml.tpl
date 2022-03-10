###########
## Chart ##
###########

postgresql:
  storageSize: 256Mi

################
## Sub-charts ##
################
basisregister-fictieve-kentekens:
  image:
    registry:  registry.nlx.reviews
    repository: commonground/nlx/nlx/basisregister-fictieve-kentekens-hackathon@sha256
    tag: ac25076a78f67ef9ac3a4f996d37a461a4ed9303694b10971e3cc1fcec46b9a7

nlx-inway:
  config:
    verwerkingenloggingAPIBaseUrl: http://vergunningsoftware-bv-nlx-outway/12345678901234567890/verwerkingenlogging

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
