// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/../../status": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "status"
                ],
                "summary": "Check the status of the server",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/access": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "access"
                ],
                "summary": "Get information about token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domains.Token"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "access"
                ],
                "summary": "Creates new token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domains.Token"
                        }
                    }
                }
            }
        },
        "/access/{token}": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "access"
                ],
                "summary": "Gets header state",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token to delete",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/chain/header/byHeight": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "headers"
                ],
                "summary": "Gets header by height",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Height to start from",
                        "name": "height",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Headers count (optional)",
                        "name": "count",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domains.BlockHeader"
                            }
                        }
                    }
                }
            }
        },
        "/chain/header/commonAncestor": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "headers"
                ],
                "summary": "Gets common ancestors",
                "parameters": [
                    {
                        "description": "JSON",
                        "name": "ancesstors",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domains.BlockHeader"
                        }
                    }
                }
            }
        },
        "/chain/header/state/{hash}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "headers"
                ],
                "summary": "Gets header state",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Requested Header Hash",
                        "name": "hash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/headers.blockHeaderStateResponse"
                        }
                    }
                }
            }
        },
        "/chain/header/{hash}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "headers"
                ],
                "summary": "Gets header by hash",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Requested Header Hash",
                        "name": "hash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domains.BlockHeader"
                        }
                    }
                }
            }
        },
        "/chain/header/{hash}/{ancestorHash}/ancestors": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "headers"
                ],
                "summary": "Gets header ancestors",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Requested Header Hash",
                        "name": "hash",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ancestor Header Hash",
                        "name": "ancestorHash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domains.BlockHeader"
                            }
                        }
                    }
                }
            }
        },
        "/chain/merkleroots/verify": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "merkleroots"
                ],
                "summary": "Verifies Merkle roots inclusion in the longest chain",
                "parameters": [
                    {
                        "description": "JSON",
                        "name": "merkleroots",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/merkleroots.merkleRootsConfirmationsResponse"
                            }
                        }
                    }
                }
            }
        },
        "/chain/tip": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tip"
                ],
                "summary": "Gets tip of longest chain",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tips.tipStateResponse"
                        }
                    }
                }
            }
        },
        "/chain/tips": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tip"
                ],
                "summary": "Gets all tips",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/tips.tipStateResponse"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/chain/tips/prune/{hash}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tip"
                ],
                "summary": "Prune tip",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Requested Header Hash",
                        "name": "hash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/network/peers": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "network"
                ],
                "summary": "Gets all peers",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/network/peers/count": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "network"
                ],
                "summary": "Gets peers count",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/webhook": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "webhooks"
                ],
                "summary": "Get webhook",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Url of webhook to check",
                        "name": "url",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/notification.Webhook"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "webhooks"
                ],
                "summary": "Register new webhook",
                "parameters": [
                    {
                        "description": "Webhook to register",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/webhook.WebhookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/notification.Webhook"
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
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "webhooks"
                ],
                "summary": "Revoke webhook",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Url of webhook to revoke",
                        "name": "url",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "domains.BlockHeader": {
            "type": "object",
            "properties": {
                "creationTimestamp": {
                    "type": "string"
                },
                "hash": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "merkleRoot": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "nonce": {
                    "type": "integer"
                },
                "prevBlockHash": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "version": {
                    "type": "integer"
                },
                "work": {
                    "type": "string"
                }
            }
        },
        "domains.Token": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "isAdmin": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "headers.blockHeaderResponse": {
            "type": "object",
            "properties": {
                "creationTimestamp": {
                    "type": "integer"
                },
                "difficultyTarget": {
                    "type": "integer"
                },
                "hash": {
                    "type": "string"
                },
                "merkleRoot": {
                    "type": "string"
                },
                "nonce": {
                    "type": "integer"
                },
                "prevBlockHash": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                },
                "work": {
                    "type": "string"
                }
            }
        },
        "headers.blockHeaderStateResponse": {
            "type": "object",
            "properties": {
                "chainWork": {
                    "type": "string"
                },
                "header": {
                    "$ref": "#/definitions/headers.blockHeaderResponse"
                },
                "height": {
                    "type": "integer"
                },
                "state": {
                    "type": "string"
                }
            }
        },
        "merkleroots.merkleRootConfirmation": {
            "type": "object",
            "properties": {
                "blockhash": {
                    "type": "string"
                },
                "confirmed": {
                    "type": "boolean"
                },
                "merkleRoot": {
                    "type": "string"
                }
            }
        },
        "merkleroots.merkleRootsConfirmationsResponse": {
            "type": "object",
            "properties": {
                "allConfirmed": {
                    "type": "boolean"
                },
                "confirmations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/merkleroots.merkleRootConfirmation"
                    }
                }
            }
        },
        "notification.Webhook": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "createdAt": {
                    "type": "string"
                },
                "errorsCount": {
                    "type": "integer"
                },
                "lastEmitStatus": {
                    "type": "string"
                },
                "lastEmitTimestamp": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "tips.tipResponse": {
            "type": "object",
            "properties": {
                "creationTimestamp": {
                    "type": "integer"
                },
                "difficultyTarget": {
                    "type": "integer"
                },
                "hash": {
                    "type": "string"
                },
                "merkleRoot": {
                    "type": "string"
                },
                "nonce": {
                    "type": "integer"
                },
                "prevBlockHash": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                },
                "work": {
                    "type": "string"
                }
            }
        },
        "tips.tipStateResponse": {
            "type": "object",
            "properties": {
                "chainWork": {
                    "type": "string"
                },
                "header": {
                    "$ref": "#/definitions/tips.tipResponse"
                },
                "height": {
                    "type": "integer"
                },
                "state": {
                    "type": "string"
                }
            }
        },
        "webhook.RequiredAuth": {
            "type": "object",
            "properties": {
                "header": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "webhook.WebhookRequest": {
            "type": "object",
            "properties": {
                "requiredAuth": {
                    "$ref": "#/definitions/webhook.RequiredAuth"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
