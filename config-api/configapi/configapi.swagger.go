package configapi
const (
SwaggerJSONDirectory = `
{
  "swagger": "2.0",
  "info": {
    "title": "configapi.proto",
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
  "paths": {},
  "definitions": {
    "ListComponentsResponseComponent": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "kind": {
          "type": "string"
        }
      }
    },
    "configapiConfig": {
      "type": "object",
      "properties": {
        "kind": {
          "type": "string"
        },
        "config": {
          "type": "string"
        }
      }
    },
    "configapiListComponentsResponse": {
      "type": "object",
      "properties": {
        "components": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ListComponentsResponseComponent"
          }
        }
      }
    },
    "configapiSetConfigResponse": {
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
