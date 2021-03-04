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
  "tags": [
    {
      "name": "Stats"
    }
  ],
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
        "operationId": "Stats_ListVersionStatistics",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/statsStatsResponse"
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
