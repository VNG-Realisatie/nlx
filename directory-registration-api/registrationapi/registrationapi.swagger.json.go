package registrationapi

const (
	SwaggerJSONDirectoryregistration = `
{
  "swagger": "2.0",
  "info": {
    "title": "registrationapi.proto",
    "version": "version not set"
  },
  "tags": [
    {
      "name": "DirectoryRegistration"
    }
  ],
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
        "documentationUrl": {
          "type": "string"
        },
        "apiSpecificationType": {
          "type": "string"
        },
        "apiSpecificationDocumentUrl": {
          "type": "string"
        },
        "insightApiUrl": {
          "type": "string"
        },
        "irmaApiUrl": {
          "type": "string"
        },
        "internal": {
          "type": "boolean"
        },
        "publicSupportContact": {
          "type": "string"
        },
        "techSupportContact": {
          "type": "string"
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
    "RegisterInwayResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
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
