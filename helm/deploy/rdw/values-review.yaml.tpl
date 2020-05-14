###########
## Chart ##
###########
managementAPI:
  insight:
    insightAPIURL: https://insight-api-rdw-{{DOMAIN_SUFFIX}}
    irmaServerURL: https://irma-rdw-{{DOMAIN_SUFFIX}}

################
## Sub-charts ##
################
nlx-management:
  config:
    directoryEndpointURL: https://directory-{{DOMAIN_SUFFIX}}/api
    oidc:
      clientSecret: ZXhhbXBsZS1hcHAtc2VjcmV0
      discoveryURL: https://dex-rdw-{{DOMAIN_SUFFIX}}
      redirectURL: https://nlx-management-rdw-{{DOMAIN_SUFFIX}}/oidc/callback
      sessionSignKey: 0Xn2DBfb4L4hwN3XosbwoKZalLBU68UU
  ingress:
    uiHostname: nlx-management-rdw-{{DOMAIN_SUFFIX}}

dex:
  config:
    issuer: https://dex-rdw-{{DOMAIN_SUFFIX}}
    staticClients:
      - id: nlx-management
        name: NLX Management
        secret: ZXhhbXBsZS1hcHAtc2VjcmV0
        redirectURIs:
          - https://nlx-management-rdw-{{DOMAIN_SUFFIX}}/oidc/callback
  ingress:
    hosts:
      - dex-rdw-{{DOMAIN_SUFFIX}}

insight-api:
  ingress:
    hosts:
      - insight-api-rdw-{{DOMAIN_SUFFIX}}

irma-server:
  ingress:
    annotations:
      ingress.kubernetes.io/custom-response-headers: "Access-Control-Allow-Origin: https://insight-{{DOMAIN_SUFFIX}}"
    host: irma-rdw-{{DOMAIN_SUFFIX}}
