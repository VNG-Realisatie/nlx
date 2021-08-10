###########
## Chart ##
###########

################
## Sub-charts ##
################
nlx-management:
  config:
    enableBasicAuth: true
  ingress:
    hosts:
        # abbreviated organization name, because https://gitlab.com/commonground/nlx/nlx/-/blob/master/technical-docs/notes.md#1215-rename-current-organizations
      - nlx-management-vgs-bv-{{DOMAIN_SUFFIX}}

dex:
  config:
    issuer: https://dex-vgs-bv-{{DOMAIN_SUFFIX}}
    staticClients:
      - id: nlx-management
        name: NLX Management
        secret: N6Gr4v8wZKlLuKrMSV1I
        redirectURIs:
          # abbreviated organization name, because https://gitlab.com/commonground/nlx/nlx/-/blob/master/technical-docs/notes.md#1215-rename-current-organizations
          - https://nlx-management-vgs-bv-{{DOMAIN_SUFFIX}}/oidc/callback
  ingress:
    hosts:
        # abbreviated organization name, because https://gitlab.com/commonground/nlx/nlx/-/blob/master/technical-docs/notes.md#1215-rename-current-organizations
      - dex-vgs-bv-{{DOMAIN_SUFFIX}}
