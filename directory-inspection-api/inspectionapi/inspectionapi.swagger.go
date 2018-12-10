package inspectionapi
const (
SwaggerJSONDirectory = `
{
  "swagger": "2.0",
  "info": {
    "title": "inspectionapi.proto",
    "description": "Package inspectionapi defines the directory api.",
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
    "/directory/get-service-api-spec/{organization_name}/{service_name}": {
      "get": {
        "operationId": "GetServiceAPISpec",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/inspectionapiGetServiceAPISpecResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "organization_name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "service_name",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Directory"
        ]
      }
    },
    "/directory/list-organizations": {
      "get": {
        "summary": "ListOrganizations lists all organizations and their details.",
        "operationId": "ListOrganizations",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/inspectionapiListOrganizationsResponse"
            }
          }
        },
        "tags": [
          "Directory"
        ]
      }
    },
    "/directory/list-services": {
      "get": {
        "summary": "ListServices lists all services and their gateways.",
        "operationId": "ListServices",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/inspectionapiListServicesResponse"
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
        }
      }
    },
    "RegisterInwayRequestRegisterService": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "documentation_url": {
          "type": "string"
        },
        "api_specification_type": {
          "type": "string"
        },
        "api_specification_document_url": {
          "type": "string"
        },
        "insight_api_url": {
          "type": "string"
        },
        "irma_api_url": {
          "type": "string"
        },
        "internal": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "inspectionapiGetServiceAPISpecResponse": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string"
        },
        "document": {
          "type": "string",
          "format": "byte"
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
    },
    "inspectionapiRegisterInwayResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        }
      }
    }
  }
}
`)
