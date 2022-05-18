###########
## Chart ##
###########

postgresql:
  storageSize: 256Mi


################
## Sub-charts ##
################
nlx-management:
  config:
    oidc:
      clientSecret: grGSl5W5HcKRETBr3OhmU6Tm
      discoveryURL: https://dex-gemeente-riemer-{{DOMAIN_SUFFIX}}
      redirectURL: https://nlx-management-gemeente-riemer-{{DOMAIN_SUFFIX}}/oidc/callback
      sessionSignKey: 0Xn2DBfb4L4hwN3XosbwoKZalLBU68UU
  ingress:
    hosts:
      - nlx-management-gemeente-riemer-{{DOMAIN_SUFFIX}}

dex:
  config:
    issuer: https://dex-gemeente-riemer-{{DOMAIN_SUFFIX}}
    staticClients:
      - id: nlx-management
        name: NLX Management
        secret: grGSl5W5HcKRETBr3OhmU6Tm
        redirectURIs:
          - https://nlx-management-gemeente-riemer-{{DOMAIN_SUFFIX}}/oidc/callback
  ingress:
    hosts:
      - dex-gemeente-riemer-{{DOMAIN_SUFFIX}}

manage-citizens-ui:
  enabled: true
  ingress:
    enabled: true
    hosts:
        # abbreviated name, because https://gitlab.com/commonground/nlx/nlx/-/blob/master/technical-docs/notes.md#1215-rename-current-organizations
      - nlx-mc-ui-gemeente-riemer-{{DOMAIN_SUFFIX}}
