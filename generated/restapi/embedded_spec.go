// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "HackZurich 2019 API\n",
    "title": "HackZurich 2019",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/game": {
      "post": {
        "tags": [
          "game"
        ],
        "summary": "Start the game",
        "operationId": "startGame",
        "responses": {
          "200": {
            "description": "Game started",
            "schema": {
              "$ref": "#/definitions/gameDescription"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/game/{game_id}": {
      "get": {
        "tags": [
          "game"
        ],
        "summary": "Get game description",
        "operationId": "getGameDescription",
        "parameters": [
          {
            "type": "string",
            "name": "game_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Description",
            "schema": {
              "$ref": "#/definitions/gameDescription"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/game/{game_id}/destroy": {
      "post": {
        "tags": [
          "game"
        ],
        "summary": "Destroy disaster on the map",
        "operationId": "destroyDisaster",
        "parameters": [
          {
            "type": "string",
            "name": "game_id",
            "in": "path",
            "required": true
          },
          {
            "name": "goal",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/point"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Destroyed",
            "schema": {
              "$ref": "#/definitions/gameState"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/game/{game_id}/pause": {
      "post": {
        "tags": [
          "game"
        ],
        "summary": "Pause the game",
        "operationId": "pauseGame",
        "parameters": [
          {
            "type": "string",
            "name": "game_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Paused",
            "schema": {
              "$ref": "#/definitions/gameDescription"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/game/{game_id}/resume": {
      "post": {
        "tags": [
          "game"
        ],
        "summary": "Resume the game",
        "operationId": "resumeGame",
        "parameters": [
          {
            "type": "string",
            "name": "game_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Resumed",
            "schema": {
              "$ref": "#/definitions/gameDescription"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/game/{game_id}/state": {
      "get": {
        "tags": [
          "game"
        ],
        "summary": "Get state of the game",
        "operationId": "getGameState",
        "parameters": [
          {
            "type": "string",
            "name": "game_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Game state",
            "schema": {
              "$ref": "#/definitions/gameState"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        },
        "path": {
          "type": "string"
        }
      }
    },
    "gameDescription": {
      "type": "object",
      "properties": {
        "id": {
          "description": "Game ID",
          "type": "string"
        },
        "status": {
          "description": "Status of the game",
          "type": "string",
          "enum": [
            "ACTIVE",
            "PAUSED",
            "FINISHED"
          ]
        }
      }
    },
    "gameState": {
      "type": "object",
      "properties": {
        "count": {
          "description": "Pinguin count",
          "type": "number",
          "format": "integer"
        },
        "countFarm": {
          "description": "Ship count",
          "type": "number",
          "format": "integer"
        },
        "countShip": {
          "description": "Ship count",
          "type": "number",
          "format": "integer"
        },
        "farms": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/point"
          }
        },
        "ships": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/point"
          }
        }
      }
    },
    "point": {
      "type": "object",
      "properties": {
        "lat": {
          "type": "number"
        },
        "lng": {
          "type": "number"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "HackZurich 2019 API\n",
    "title": "HackZurich 2019",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/game": {
      "post": {
        "tags": [
          "game"
        ],
        "summary": "Start the game",
        "operationId": "startGame",
        "responses": {
          "200": {
            "description": "Game started",
            "schema": {
              "$ref": "#/definitions/gameDescription"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/game/{game_id}": {
      "get": {
        "tags": [
          "game"
        ],
        "summary": "Get game description",
        "operationId": "getGameDescription",
        "parameters": [
          {
            "type": "string",
            "name": "game_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Description",
            "schema": {
              "$ref": "#/definitions/gameDescription"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/game/{game_id}/destroy": {
      "post": {
        "tags": [
          "game"
        ],
        "summary": "Destroy disaster on the map",
        "operationId": "destroyDisaster",
        "parameters": [
          {
            "type": "string",
            "name": "game_id",
            "in": "path",
            "required": true
          },
          {
            "name": "goal",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/point"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Destroyed",
            "schema": {
              "$ref": "#/definitions/gameState"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/game/{game_id}/pause": {
      "post": {
        "tags": [
          "game"
        ],
        "summary": "Pause the game",
        "operationId": "pauseGame",
        "parameters": [
          {
            "type": "string",
            "name": "game_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Paused",
            "schema": {
              "$ref": "#/definitions/gameDescription"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/game/{game_id}/resume": {
      "post": {
        "tags": [
          "game"
        ],
        "summary": "Resume the game",
        "operationId": "resumeGame",
        "parameters": [
          {
            "type": "string",
            "name": "game_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Resumed",
            "schema": {
              "$ref": "#/definitions/gameDescription"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/game/{game_id}/state": {
      "get": {
        "tags": [
          "game"
        ],
        "summary": "Get state of the game",
        "operationId": "getGameState",
        "parameters": [
          {
            "type": "string",
            "name": "game_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Game state",
            "schema": {
              "$ref": "#/definitions/gameState"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        },
        "path": {
          "type": "string"
        }
      }
    },
    "gameDescription": {
      "type": "object",
      "properties": {
        "id": {
          "description": "Game ID",
          "type": "string"
        },
        "status": {
          "description": "Status of the game",
          "type": "string",
          "enum": [
            "ACTIVE",
            "PAUSED",
            "FINISHED"
          ]
        }
      }
    },
    "gameState": {
      "type": "object",
      "properties": {
        "count": {
          "description": "Pinguin count",
          "type": "number",
          "format": "integer"
        },
        "countFarm": {
          "description": "Ship count",
          "type": "number",
          "format": "integer"
        },
        "countShip": {
          "description": "Ship count",
          "type": "number",
          "format": "integer"
        },
        "farms": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/point"
          }
        },
        "ships": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/point"
          }
        }
      }
    },
    "point": {
      "type": "object",
      "properties": {
        "lat": {
          "type": "number"
        },
        "lng": {
          "type": "number"
        }
      }
    }
  }
}`))
}
