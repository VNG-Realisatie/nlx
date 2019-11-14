package inspectionapi
const (
SwaggerJSONDirectoryInspection = `
{
  "swagger": "2.0",
  "info": {
    "title": "inspectionapi.proto",
    "description": "Package inspectionapi defines the directory api.",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/directory/list-organizations": {
      "get": {
        "summary": "ListOrganizations lists all organizations and their details.",
        "operationId": "ListOrganizations",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/inspectionapiListOrganizationsResponse"
            }
          }
        },
        "tags": [
          "DirectoryInspection"
        ]
      }
    },
    "/directory/list-services": {
      "get": {
        "summary": "ListServices lists all services and their gateways.",
        "operationId": "ListServices",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/inspectionapiListServicesResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "outway_version",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "DirectoryInspection"
        ]
      }
    }
  },
  "definitions": {
    "InwayState": {
      "type": "string",
      "enum": [
        "UNKNOWN",
        "UP",
        "DOWN"
      ],
      "default": "UNKNOWN"
    },
    "ListOrganizationsResponseOrganization": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "insight_irma_endpoint": {
          "type": "string"
        },
        "insight_log_endpoint": {
          "type": "string"
        }
      }
    },
    "ListServicesResponseService": {
      "type": "object",
      "properties": {
        "organization_name": {
          "type": "string"
        },
        "service_name": {
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
        },
        "api_specification_type": {
          "type": "string"
        },
        "internal": {
          "type": "boolean",
          "format": "boolean"
        },
        "public_support_contact": {
          "type": "string"
        },
        "healthy_states": {
          "type": "array",
          "items": {
            "type": "boolean",
            "format": "boolean"
          }
        },
        "inways": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/inspectionapiInway"
          }
        }
      }
    },
    "inspectionapiInway": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "state": {
          "$ref": "#/definitions/InwayState"
        }
      }
    },
    "inspectionapiListOrganizationsResponse": {
      "type": "object",
      "properties": {
        "organizations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ListOrganizationsResponseOrganization"
          }
        }
      }
    },
    "inspectionapiListServicesResponse": {
      "type": "object",
      "properties": {
        "services": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ListServicesResponseService"
          }
        }
      }
    }
  }
}
`)
