//nolint
package configapi
const (
SwaggerJSONDirectory = `
{
  "swagger": "2.0",
  "info": {
    "title": "configapi.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v1/insight-configuration": {
      "get": {
        "operationId": "GetInsightConfiguration",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiInsightConfiguration"
            }
          }
        },
        "tags": [
          "ConfigApi"
        ]
      },
      "put": {
        "operationId": "PutInsightConfiguration",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiInsightConfiguration"
            }
          }
        },
        "tags": [
          "ConfigApi"
        ]
      }
    },
    "/api/v1/inways": {
      "get": {
        "operationId": "ListInways",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiListInwaysResponse"
            }
          }
        },
        "tags": [
          "ConfigApi"
        ]
      },
      "post": {
        "operationId": "CreateInway",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiInway"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/configapiInway"
            }
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      }
    },
    "/api/v1/inways/{name}": {
      "get": {
        "operationId": "GetInway",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiInway"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      },
      "delete": {
        "operationId": "DeleteInway",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      },
      "put": {
        "operationId": "UpdateInway",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiInway"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/configapiInway"
            }
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      }
    },
    "/api/v1/services": {
      "get": {
        "operationId": "ListServices",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiListServicesResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "inwayName",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      },
      "post": {
        "operationId": "CreateService",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiService"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/configapiService"
            }
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      }
    },
    "/api/v1/services/{name}": {
      "get": {
        "operationId": "GetService",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiService"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      },
      "delete": {
        "operationId": "DeleteService",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      },
      "put": {
        "operationId": "UpdateService",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/configapiService"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/configapiService"
            }
          }
        ],
        "tags": [
          "ConfigApi"
        ]
      }
    }
  },
  "definitions": {
    "AuthorizationSettingsAuthorization": {
      "type": "object",
      "properties": {
        "organizationName": {
          "type": "string"
        },
        "publicKeyHash": {
          "type": "string"
        }
      }
    },
    "ServiceAuthorizationSettings": {
      "type": "object",
      "properties": {
        "mode": {
          "type": "string"
        },
        "authorizations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/AuthorizationSettingsAuthorization"
          }
        }
      }
    },
    "configapiEmpty": {
      "type": "object"
    },
    "configapiInsightConfiguration": {
      "type": "object",
      "properties": {
        "irmaServerURL": {
          "type": "string"
        },
        "insightAPIURL": {
          "type": "string"
        }
      }
    },
    "configapiInway": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "configapiListInwaysResponse": {
      "type": "object",
      "properties": {
        "inways": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/configapiInway"
          }
        }
      }
    },
    "configapiListServicesResponse": {
      "type": "object",
      "properties": {
        "services": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/configapiService"
          }
        }
      }
    },
    "configapiService": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "endpointURL": {
          "type": "string"
        },
        "documentationURL": {
          "type": "string"
        },
        "apiSpecificationURL": {
          "type": "string"
        },
        "internal": {
          "type": "boolean",
          "format": "boolean"
        },
        "techSupportContact": {
          "type": "string"
        },
        "publicSupportContact": {
          "type": "string"
        },
        "authorizationSettings": {
          "$ref": "#/definitions/ServiceAuthorizationSettings"
        },
        "inways": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  }
}
`)
