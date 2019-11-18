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
        "operationId": "ListStats",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/statsStatsResponse"
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
          "type": "string",
          "format": "uint64"
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
