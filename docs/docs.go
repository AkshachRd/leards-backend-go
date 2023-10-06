// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "returns user id of an existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Login an existing user",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/httputils.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "creates new user and returns a token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "User register data",
                        "name": "createUserData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httputils.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/httputils.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    }
                }
            }
        },
        "/auth/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "when token is expired you need to refresh it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh user's token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputils.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "when user signs out token needs to be revoked",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Revokes user's token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputils.BasicResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    }
                }
            }
        },
        "/decks/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "fetches the deck from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "decks"
                ],
                "summary": "Get single deck by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Deck ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputils.DeckResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    }
                }
            }
        },
        "/folders/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "fetches the folder from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "folders"
                ],
                "summary": "Get single folder by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Folder ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputils.FolderResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httputils.BasicResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Successfully"
                }
            }
        },
        "httputils.Card": {
            "type": "object",
            "properties": {
                "backSide": {
                    "type": "string"
                },
                "cardId": {
                    "type": "string"
                },
                "frontSide": {
                    "type": "string"
                }
            }
        },
        "httputils.Content": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "httputils.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "bob@leards.space"
                },
                "password": {
                    "type": "string",
                    "example": "123"
                },
                "username": {
                    "type": "string",
                    "example": "Bob"
                }
            }
        },
        "httputils.Deck": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/httputils.Card"
                    }
                },
                "deckId": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "httputils.DeckResponse": {
            "type": "object",
            "properties": {
                "deck": {
                    "$ref": "#/definitions/httputils.Deck"
                },
                "message": {
                    "type": "string",
                    "example": "Successfully"
                }
            }
        },
        "httputils.Folder": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/httputils.Content"
                    }
                },
                "folderId": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/httputils.Path"
                    }
                }
            }
        },
        "httputils.FolderResponse": {
            "type": "object",
            "properties": {
                "folder": {
                    "$ref": "#/definitions/httputils.Folder"
                },
                "message": {
                    "type": "string",
                    "example": "Successfully"
                }
            }
        },
        "httputils.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "httputils.Path": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "httputils.TokenResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Successfully"
                },
                "token": {
                    "type": "string",
                    "example": "\u003ctoken\u003e"
                }
            }
        },
        "httputils.UserResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Successfully"
                },
                "token": {
                    "type": "string",
                    "example": "\u003ctoken\u003e"
                },
                "userId": {
                    "type": "string",
                    "example": "53f4cf69-9da6-49e4-8651-450b74abdf9e"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        },
        "Bearer": {
            "description": "ATTENTION! HOW TO USE: Type \"Bearer\" followed by a space and a token. Example: \"Bearer \\\u003ctoken\\\u003e\".",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Leards Backend API",
	Description:      "This is a leards language learning app api.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
