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
        "operationId": "DirectoryInspection_ListOrganizations",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/inspectionapiListOrganizationsResponse"
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
        "operationId": "DirectoryInspection_ListServices",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/inspectionapiListServicesResponse"
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
    "/api/stats": {
      "get": {
        "operationId": "DirectoryInspection_ListInOutwayStatistics",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/inspectionapiListInOutwayStatisticsResponse"
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
    "InwayState": {
      "type": "string",
      "enum": [
        "UNKNOWN",
        "UP",
        "DOWN"
      ],
      "default": "UNKNOWN"
    },
    "ListInOutwayStatisticsResponseStatistics": {
      "type": "object",
      "properties": {
        "type": {
          "$ref": "#/definitions/ListInOutwayStatisticsResponseStatisticsType"
        },
        "version": {
          "type": "string"
        },
        "amount": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "ListInOutwayStatisticsResponseStatisticsType": {
      "type": "string",
      "enum": [
        "INWAY",
        "OUTWAY"
      ],
      "default": "INWAY"
    },
    "ListServicesResponseCosts": {
      "type": "object",
      "properties": {
        "one_time": {
          "type": "integer",
          "format": "int32"
        },
        "monthly": {
          "type": "integer",
          "format": "int32"
        },
        "request": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "ListServicesResponseService": {
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
        "internal": {
          "type": "boolean"
        },
        "public_support_contact": {
          "type": "string"
        },
        "inways": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/inspectionapiInway"
          }
        },
        "costs": {
          "$ref": "#/definitions/ListServicesResponseCosts"
        },
        "organization": {
          "$ref": "#/definitions/inspectionapiOrganization"
        }
      }
    },
    "inspectionapiGetOrganizationInwayResponse": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
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
    "inspectionapiListInOutwayStatisticsResponse": {
      "type": "object",
      "properties": {
        "versions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ListInOutwayStatisticsResponseStatistics"
          }
        }
      }
    },
    "inspectionapiListOrganizationsResponse": {
      "type": "object",
      "properties": {
        "organizations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/inspectionapiOrganization"
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
    "inspectionapiOrganization": {
      "type": "object",
      "properties": {
        "serial_number": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "protobufAny": {
      "type": "object",
      "properties": {
        "type_url": {
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
