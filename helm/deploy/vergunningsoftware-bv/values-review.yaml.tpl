###########
## Chart ##
###########

postgresql:
  storageSize: 256Mi

outway:
  ingress:
    enabled: true
    host: nlx-outway-vgs-bv-{{DOMAIN_SUFFIX}}

opa:
  enabled: true

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

parkeerrechten-admin:
  enabled: true
  organizationName: "Vergunningsoftware BV"
  organizationLogo: https://gitlab.com/commonground/nlx/demo/-/raw/c75d67e8e0b631efdb0edec4df74b227483744c0/assets/stijns.svg
  organizationColor: #FEBC2D
  servicesList:
    - organization: "Stijns"
      baseUrl: "http://vergunningsoftware-bv-nlx-outway/12345678901234567890/parkeerrechten"
  kentekenApiBaseUrl: http://vergunningsoftware-bv-nlx-outway/12345678901234567891/basisregister-fictieve-kentekens
  personenApiBaseUrl: http://vergunningsoftware-bv-nlx-outway/12345678901234567892/basisregister-fictieve-personen
  ingress:
    enabled: true
    hosts:
      # abbreviated name, because https://gitlab.com/commonground/nlx/nlx/-/blob/master/technical-docs/notes.md#1215-rename-current-organizations
      - nlx-pa-vgs-bv-{{DOMAIN_SUFFIX}}
