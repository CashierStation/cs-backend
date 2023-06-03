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
        "/api/employee/list": {
            "get": {
                "security": [
                    {
                        "SessionToken": []
                    }
                ],
                "description": "Get list of employees from access token",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/employee"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Access token from Auth0",
                        "name": "access_token",
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
                    "api/snack"
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
                    "api/snack"
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
                    },
                    {
                        "type": "string",
                        "description": "Snack category",
                        "name": "category",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Snack stock",
                        "name": "stock",
                        "in": "query"
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
        "/api/snack/restock": {
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
                    "api/snack"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Snack ID",
                        "name": "snack_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Quantity",
                        "name": "quantity",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Total price",
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
        "/api/snack/transaction": {
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
                    "api/snack"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Unit ID",
                        "name": "unit_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Snack ID",
                        "name": "snack_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Quantity",
                        "name": "quantity",
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
                    "api/snack"
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
                        "type": "string",
                        "description": "Snack category",
                        "name": "category",
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
                    "api/snack"
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
                    "api/unit"
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
                    "api/unit"
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
                    },
                    {
                        "type": "string",
                        "description": "Unit category",
                        "name": "category",
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
                    "api/unit"
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
                    "api/unit"
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
        "/api/unit_session": {
            "get": {
                "security": [
                    {
                        "SessionToken": []
                    }
                ],
                "description": "Sesi pemakaian unit",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/unit_session"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "unit id",
                        "name": "unit_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "default": "desc",
                        "description": "order",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "id",
                            "unit_id",
                            "start_time",
                            "finish_time",
                            "tarif"
                        ],
                        "type": "string",
                        "default": "start_time",
                        "description": "sort_by",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "default": false,
                        "description": "select only latest session for each unit",
                        "name": "latest",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/unitsession.GetUnitSessionsResponse"
                        }
                    }
                }
            }
        },
        "/api/unit_session/start/{unit_id}": {
            "put": {
                "security": [
                    {
                        "SessionToken": []
                    }
                ],
                "description": "Sesi pemakaian unit",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/unit_session"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Unit ID",
                        "name": "unit_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/unitsession.StartUnitSessionsResponse"
                        }
                    }
                }
            }
        },
        "/api/unit_session/stop/{unit_id}": {
            "put": {
                "security": [
                    {
                        "SessionToken": []
                    }
                ],
                "description": "Sesi pemakaian unit",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/unit_session"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Unit ID",
                        "name": "unit_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/unitsession.StopUnitSessionsResponse"
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
                    "api/user"
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
        "/jobs": {
            "get": {
                "description": "dev: http://localhost:8080/jobs\nprod: https://csbackend.fly.dev/jobs",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
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
        "/metrics/db": {
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
        "enum.UnitStatus": {
            "type": "string",
            "enum": [
                "idle",
                "in_use",
                "booked",
                "booked_while_in_use"
            ],
            "x-enum-varnames": [
                "Idle",
                "InUse",
                "Booked",
                "BookedWhileInUse"
            ]
        },
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
                "category": {
                    "type": "string"
                },
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
                },
                "stock": {
                    "type": "integer"
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
        "unit.GetUnit": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
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
                },
                "status": {
                    "$ref": "#/definitions/unit.GetUnitStatus"
                }
            }
        },
        "unit.GetUnitResponse": {
            "type": "object",
            "properties": {
                "units": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/unit.GetUnit"
                    }
                }
            }
        },
        "unit.GetUnitStatus": {
            "type": "object",
            "properties": {
                "latest_finish_time": {
                    "type": "string"
                },
                "latest_start_time": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/enum.UnitStatus"
                },
                "tarif": {
                    "type": "integer"
                }
            }
        },
        "unit.PostUnit": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
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
        "unit.PostUnitResponse": {
            "type": "object",
            "properties": {
                "unit": {
                    "$ref": "#/definitions/unit.PostUnit"
                }
            }
        },
        "unit.PutUnit": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
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
        "unit.PutUnitResponse": {
            "type": "object",
            "properties": {
                "unit": {
                    "$ref": "#/definitions/unit.PutUnit"
                }
            }
        },
        "unitsession.GetUnitSessionsResponse": {
            "type": "object",
            "properties": {
                "total": {
                    "type": "integer"
                },
                "unit_sessions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/unitsession.UnitSessionResponse"
                    }
                }
            }
        },
        "unitsession.SnackTransactionResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "snack_name": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "unitsession.StartUnitSessionsResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "$ref": "#/definitions/unitsession.startUnitSessionsStatus"
                },
                "unit_id": {
                    "type": "integer"
                },
                "unit_name": {
                    "type": "string"
                }
            }
        },
        "unitsession.StopUnitSessionsResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "$ref": "#/definitions/unitsession.stopUnitSessionsStatus"
                },
                "unit_id": {
                    "type": "integer"
                },
                "unit_name": {
                    "type": "string"
                }
            }
        },
        "unitsession.UnitSessionResponse": {
            "type": "object",
            "properties": {
                "finish_time": {
                    "type": "string"
                },
                "grand_total": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "snack_transactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/unitsession.SnackTransactionResponse"
                    }
                },
                "start_time": {
                    "type": "string"
                },
                "tarif": {
                    "type": "integer"
                },
                "unit_id": {
                    "type": "integer"
                }
            }
        },
        "unitsession.startUnitSessionsStatus": {
            "type": "object",
            "properties": {
                "finish_time": {
                    "type": "string"
                },
                "start_time": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/enum.UnitStatus"
                }
            }
        },
        "unitsession.stopUnitSessionsStatus": {
            "type": "object",
            "properties": {
                "finish_time": {
                    "type": "string"
                },
                "start_time": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/enum.UnitStatus"
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
