# nqp - NLX Quic Protocol

NQP is an improved protocol for communication between inways and outways. It is based on Quic and h2quic.

The design of NQP has not been completed yet. The purpose of this document is to gather notes and thoughts on how NQP would work.

## Quic as a foundation for communication between inways and outways

TODO: List obvious reasons for Quic

### Quick & TLS

TODO: Describe challenges when using our (keyless) TLS setup with Quic.

### Proxy streams

http requests are proxied onto streams using h2quic. A non-h2quic leader transaction can be used to exchange information about the upcoming http requiest so that, unlike with [nhp](./nhp.md), the proxied http request is unmodified from the original.

### Gossip streams

Separate streams between gateways can be used to gossip about the proxy streams. This can be used to perform time-travel; a proxied HTTP request can be sent to an inway, and kept in a buffer, until both sides have written the txlog to db. This improves handling speed as the database transaction and request proxying can be performed in parallel.

## Metadata

TODO: Describe metadata that is sent for each request

## Transactionlogs

TODO: Describe rules on txlogs
