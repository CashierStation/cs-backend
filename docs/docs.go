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
        "/auth/callback": {
            "get": {
                "description": "Callback",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Endpoint the user is redirected to after logging in.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.Callback.response"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "get": {
                "description": "Login",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Redirect user to third party login",
                "responses": {}
            }
        },
        "/auth/logout": {
            "get": {
                "description": "Logout",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Log user out",
                "responses": {}
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
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Example"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
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
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Example": {
            "type": "object"
        },
        "api.User": {
            "type": "object"
        },
        "auth.Callback.response": {
            "type": "object"
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