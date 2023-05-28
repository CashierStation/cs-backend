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
        "/api/snack": {
            "get": {
                "security": [
                    {
                        "SessionToken": []
                    }
                ],
                "description": "Snack",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "snack"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/snack.GetSnackResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "SessionToken": []
                    }
                ],
                "description": "Snack",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "snack"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Snack name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Snack price",
                        "name": "price",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/snack.PostSnackResponse"
                        }
                    }
                }
            }
        },
        "/api/snack/{id}": {
            "put": {
                "security": [
                    {
                        "SessionToken": []
                    }
                ],
                "description": "Snack",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "snack"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Snack ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Snack name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Snack price",
                        "name": "price",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/snack.PutSnackResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "SessionToken": []
                    }
                ],
                "description": "Snack",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "snack"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Snack ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/snack.DeleteSnackResponse"
                        }
                    }
                }
            }
        },
        "/api/unit": {
            "get": {
                "security": [
                    {
                        "SessionToken": []
                    }
                ],
                "description": "Unit",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "unit"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/unit.GetUnitResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "SessionToken": []
                    }
                ],
                "description": "Unit",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "unit"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Unit name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Unit hourly price",
                        "name": "hourly_price",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/unit.PostUnitResponse"
                        }
                    }
                }
            }
        },
        "/api/unit/{id}": {
            "put": {
                "security": [
                    {
                        "SessionToken": []
                    }
                ],
                "description": "Unit",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "unit"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Unit ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Unit name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Unit hourly price",
                        "name": "hourly_price",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/unit.PutUnitResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "SessionToken": []
                    }
                ],
                "description": "Unit",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "unit"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Unit ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/unit.DeleteUnitResponse"
                        }
                    }
                }
            }
        },
        "/api/user": {
            "get": {
                "security": [
                    {
                        "SessionToken": []
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
        "/auth/login": {
            "post": {
                "description": "dev: http://localhost:8080/auth/login\nprod: https://csbackend.fly.dev/auth/login",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login a new employee/owner account",
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
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/login.LoginPostResponse"
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
        "login.LoginPostResponse": {
            "type": "object",
            "properties": {
                "session_token": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "register.RegisterPostResponse": {
            "type": "object",
            "properties": {
                "role": {
                    "type": "string"
                },
                "session_token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "snack.DeleteSnackResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "snack.GetSnackResponse": {
            "type": "object",
            "properties": {
                "snacks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/snack.SnackResponse"
                    }
                }
            }
        },
        "snack.PostSnackResponse": {
            "type": "object",
            "properties": {
                "snack": {
                    "$ref": "#/definitions/snack.SnackResponse"
                }
            }
        },
        "snack.PutSnackResponse": {
            "type": "object",
            "properties": {
                "snack": {
                    "$ref": "#/definitions/snack.SnackResponse"
                }
            }
        },
        "snack.SnackResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "rental_id": {
                    "type": "string"
                }
            }
        },
        "unit.DeleteUnitResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "unit.GetUnitResponse": {
            "type": "object",
            "properties": {
                "units": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/unit.UnitResponse"
                    }
                }
            }
        },
        "unit.PostUnitResponse": {
            "type": "object",
            "properties": {
                "unit": {
                    "$ref": "#/definitions/unit.UnitResponse"
                }
            }
        },
        "unit.PutUnitResponse": {
            "type": "object",
            "properties": {
                "unit": {
                    "$ref": "#/definitions/unit.UnitResponse"
                }
            }
        },
        "unit.UnitResponse": {
            "type": "object",
            "properties": {
                "hourly_price": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "rental_id": {
                    "type": "string"
                }
            }
        },
        "user.GET.response": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "rental_id": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "SessionToken": {
            "type": "apiKey",
            "name": "X-Session",
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
