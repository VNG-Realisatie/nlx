package registrationapi

const (
	SwaggerJSONDirectory = `
{
  "swagger": "2.0",
  "info": {
    "title": "registrationapi.proto",
    "description": "Package directoryapi defines the directory api.",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {},
  "definitions": {
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
        },
        "public_support_contact": {
          "type": "string"
        },
        "tech_support_contact": {
          "type": "string"
        },
        "one_time_costs": {
          "type": "integer",
          "format": "int32"
        },
        "monthly_costs": {
          "type": "integer",
          "format": "int32"
        },
        "request_costs": {
          "type": "integer",
          "format": "int32"
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
    "registrationapiRegisterInwayResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        }
      }
    },
    "runtimeError": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
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
