package inspectionapi

const (
	SwaggerJSONDirectoryInspection = `
{
  "swagger": "2.0",
  "info": {
    "title": "inspectionapi.proto",
    "version": "version not set"
  },
  "tags": [
    {
      "name": "DirectoryInspection"
    }
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/directory/list-organizations": {
      "get": {
        "summary": "ListOrganizations lists all organizations and their details.",
        "operationId": "DirectoryInspection_ListOrganizations",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/ListOrganizationsResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "tags": [
          "DirectoryInspection"
        ]
      }
    },
    "/api/directory/list-services": {
      "get": {
        "summary": "ListServices lists all services and their gateways.",
        "operationId": "DirectoryInspection_ListServices",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/ListServicesResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "tags": [
          "DirectoryInspection"
        ]
      }
    }
  },
  "definitions": {
    "GetOrganizationInwayResponse": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        }
      }
    },
    "Inway": {
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
    "InwayState": {
      "type": "string",
      "enum": [
        "UNKNOWN",
        "UP",
        "DOWN"
      ],
      "default": "UNKNOWN"
    },
    "ListOrganizationsResponse": {
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
    "ListOrganizationsResponseOrganization": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "insightIrmaEndpoint": {
          "type": "string"
        },
        "insightLogEndpoint": {
          "type": "string"
        }
      }
    },
    "ListServicesResponse": {
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
    "ListServicesResponseService": {
      "type": "object",
      "properties": {
        "organizationName": {
          "type": "string"
        },
        "serviceName": {
          "type": "string"
        },
        "inwayAddresses": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "documentationUrl": {
          "type": "string"
        },
        "apiSpecificationType": {
          "type": "string"
        },
        "internal": {
          "type": "boolean"
        },
        "publicSupportContact": {
          "type": "string"
        },
        "healthyStates": {
          "type": "array",
          "items": {
            "type": "boolean"
          }
        },
        "inways": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Inway"
          }
        },
        "oneTimeCosts": {
          "type": "integer",
          "format": "int32"
        },
        "monthlyCosts": {
          "type": "integer",
          "format": "int32"
        },
        "requestCosts": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "protobufAny": {
      "type": "object",
      "properties": {
        "typeUrl": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "rpcStatus": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/protobufAny"
          }
        }
      }
    }
  }
}
`
)
