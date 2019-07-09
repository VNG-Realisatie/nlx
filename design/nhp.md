# nhp - nlx http protocol

NHP is the current protocol used by inways and outways.

Outways proxy an HTTP API request to an Inway using HTTP. The outway adds headers with NLX-specific metadata to the request. The inway uses these headers to write a transactionlog and strips them before sending the request to the API endpoint.

## Path

TODO: Describe how the http location header is used to indicate which service the inway should send the request to.

## Metadata

TODO: Describe metadata that is sent for each request

## Transactionlogs

TODO: Describe rules on txlogs
