###########
## Chart ##
###########

postgresql:
  storageSize: 256Mi

outway:
  ingress:
    enabled: true
    host: nlx-outway-gemeente-stijns-{{DOMAIN_SUFFIX}}

outway-2:
  ingress:
    enabled: true
    host: nlx-outway-2-gemeente-stijns-{{DOMAIN_SUFFIX}}

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

nginx-video-player-ui-proxy:
  organizationName: "Gemeente Stijns"
  outwayProxyUrl: http://gemeente-stijns-nlx-outway/12345678901234567891/voorbeeld-video-stream
  ingress:
    enabled: true
    hosts:
        # abbreviated name, because https://gitlab.com/commonground/nlx/nlx/-/blob/master/technical-docs/notes.md#1215-rename-current-organizations
      - nlx-nginx-vp-p-gemeente-stijns-{{DOMAIN_SUFFIX}}

video-player-ui:
  organizationName: "Gemeente Stijns"
  outwayProxyUrl: http://nlx-nginx-vp-p-gemeente-stijns-{{DOMAIN_SUFFIX}}
  ingress:
    enabled: true
    hosts:
      - nlx-vp-ui-gemeente-stijns-{{DOMAIN_SUFFIX}}

nginx-websockets-proxy:
  organizationName: "Gemeente Stijns"
  outwayServiceBaseUrl: http://gemeente-stijns-nlx-outway/12345678901234567891/voorbeeld-websockets
  ingress:
    enabled: true
    hosts:
        # abbreviated name, because https://gitlab.com/commonground/nlx/nlx/-/blob/master/technical-docs/notes.md#1215-rename-current-organizations
      - nlx-nginx-ws-p-gemeente-stijns-{{DOMAIN_SUFFIX}}

websockets-chat-ui:
  organizationName: "Gemeente Stijns"
  websocketsProxyBaseUrl: wss://nlx-nginx-ws-p-gemeente-stijns-{{DOMAIN_SUFFIX}}
  ingress:
    enabled: true
    hosts:
        # abbreviated name, because https://gitlab.com/commonground/nlx/nlx/-/blob/master/technical-docs/notes.md#1215-rename-current-organizations
      - nlx-ws-chat-ui-gemeente-stijns-{{DOMAIN_SUFFIX}}
