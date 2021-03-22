###########
## Chart ##
###########
postgresql:
  storageType: Ephemeral

################
## Sub-charts ##
################
nlx-management:
  config:
    oidc:
      clientSecret: N6Gr4v8wZKlLuKrMSV1I
      discoveryURL: https://dex-saas-org-x-{{DOMAIN_SUFFIX}}
      redirectURL: https://nlx-management-saas-org-x-{{DOMAIN_SUFFIX}}/oidc/callback
      sessionSignKey: 0Xn2DBfb4L4hwN3XosbwoKZalLBU68UU
  ingress:
    hosts:
      - nlx-management-saas-org-x-{{DOMAIN_SUFFIX}}

dex:
  config:
    issuer: https://dex-saas-org-x-{{DOMAIN_SUFFIX}}
    staticClients:
      - id: nlx-management
        name: NLX Management
        secret: N6Gr4v8wZKlLuKrMSV1I
        redirectURIs:
          - https://nlx-management-saas-org-x-{{DOMAIN_SUFFIX}}/oidc/callback
  ingress:
    hosts:
      - dex-saas-org-x-{{DOMAIN_SUFFIX}}
