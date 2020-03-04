package stats

const (
	SwaggerJSONDirectoryInspection = `
{
  "swagger": "2.0",
  "info": {
    "title": "stats.proto",
    "description": "Package stats defines the stats api.",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/stats": {
      "get": {
        "summary": "ListStats lists all versions for inways and outways.",
        "operationId": "ListVersionStatistics",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/statsStatsResponse"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/runtimeError"
            }
          }
        },
        "tags": [
          "Stats"
        ]
      }
    }
  },
  "definitions": {
    "StatsResponseVersionStat": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string"
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
    },
    "statsStatsResponse": {
      "type": "object",
      "properties": {
        "versions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/StatsResponseVersionStat"
          }
        }
      }
    }
  }
}
`
)
