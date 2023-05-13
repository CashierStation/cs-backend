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
        "/api/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "User",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID Token",
                        "name": "id_token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.GET.response"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "dev: http://localhost:8080/auth/register\nprod: https://csbackend.fly.dev/auth/register",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new employee/owner account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Access token from Auth0",
                        "name": "access_token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 32,
                        "minLength": 3,
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 32,
                        "minLength": 6,
                        "type": "string",
                        "description": "Password (Numeric)",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "owner",
                            "karyawan"
                        ],
                        "type": "string",
                        "description": "Role",
                        "name": "role",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/register.RegisterPostResponse"
                        }
                    }
                }
            }
        },
        "/example": {
            "get": {
                "description": "Example",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "responses": {}
            }
        },
        "/metrics": {
            "get": {
                "description": "dev: http://localhost:8080/metrics\nprod: https://csbackend.fly.dev/metrics",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "metrics"
                ],
                "responses": {}
            }
        },
        "/oauth/callback": {
            "get": {
                "description": "Callback",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "oauth"
                ],
                "summary": "Endpoint the user is redirected to after logging in.",
                "responses": {}
            }
        },
        "/oauth/login": {
            "get": {
                "description": "dev: http://localhost:8080/oauth/login\nprod: https://csbackend.fly.dev/oauth/login",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "oauth"
                ],
                "summary": "Redirect user to third party login",
                "responses": {}
            }
        },
        "/oauth/logout": {
            "get": {
                "description": "dev: http://localhost:8080/oauth/logout\nprod: https://csbackend.fly.dev/oauth/logout",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "oauth"
                ],
                "summary": "Log user out",
                "responses": {}
            }
        }
    },
    "definitions": {
        "register.RegisterPostResponse": {
            "type": "object",
            "properties": {
                "role": {
                    "type": "string"
                },
                "session_token": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.GET.response": {
            "type": "object",
            "properties": {
                "aud": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "email_verified": {
                    "type": "boolean"
                },
                "exp": {
                    "type": "integer"
                },
                "family_name": {
                    "type": "string"
                },
                "given_name": {
                    "type": "string"
                },
                "iat": {
                    "type": "integer"
                },
                "iss": {
                    "type": "string"
                },
                "locale": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "sid": {
                    "type": "string"
                },
                "sub": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "CashierStation Backend Server API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
