package directoryapi
const (
SwaggerJSONDirectory = `
{
  "swagger": "2.0",
  "info": {
    "title": "directory.proto",
    "description": "Package directoryapi defines the directory api.",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/directory/list-services": {
      "get": {
        "summary": "ListServices lists all services and their gateways.",
        "operationId": "ListServices",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/directoryapiListServicesResponse"
            }
          }
        },
        "tags": [
          "Directory"
        ]
      }
    }
  },
  "definitions": {
    "RegisterInwayRequestRegisterService": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "documentation_url": {
          "type": "string"
        }
      }
    },
    "directoryapiListServicesResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "services": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/directoryapiService"
          }
        }
      }
    },
    "directoryapiRegisterInwayResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        }
      }
    },
    "directoryapiService": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "organization_name": {
          "type": "string"
        },
        "inway_addresses": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "documentation_url": {
          "type": "string"
        }
      }
    }
  }
}
`)
