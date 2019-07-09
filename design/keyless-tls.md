
https://blog.cloudflare.com/keyless-ssl-the-nitty-gritty-technical-details/

https://github.com/cloudflare/gokeyless

Our challenge is making it mutual, and based on our own PKI structure.

As a part of this, the keypair that is used by the keyserver to sign the handshake of inway and outway connections, can be a keypair that is uniquely generated for the given inway/outway, and signed by a PKI-O organization certificate.

root > service > intermediate (csp) > organization > instance

The instance cert is generated and signed for each instance that is spawned, so that on both sides of the connection, that unique instance can be logged.

Also need to think about session resumption between different outways/inways. Can they do hand-over of session tickets via a centralized server?
