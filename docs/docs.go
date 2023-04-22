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
        "/equip": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "create equip for personal user",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Equip"
                ],
                "summary": "create equip",
                "parameters": [
                    {
                        "type": "string",
                        "description": "equip name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "note for equip",
                        "name": "note",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{'message': 'create success'}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "{'message': 'jwt token error', 'error': 'error message'}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "{'message': '缺少必要欄位', 'error': 'error message'}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/equip/{id}": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "delete equip for personal user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Equip"
                ],
                "summary": "delete equip",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "equip id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{'message': 'deleted success'}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "{'message': 'jwt token error', 'error': 'error message'}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "{'message': '缺少必要欄位', 'error': 'error message'}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "update equip for personal user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Equip"
                ],
                "summary": "update equip",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "equip id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "note for equip",
                        "name": "weights",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/equip.updateFrom"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{'message': 'create success'}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "{'message': 'jwt token error', 'error': 'error message'}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "{'message': '缺少必要欄位', 'error': 'error message'}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/equip/{id}/weight": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "update equip weight for personal user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Equip"
                ],
                "summary": "update equip weight",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "equip id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "note for equip",
                        "name": "weights",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/equip.weightForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{'message': 'create success'}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "{'message': 'jwt token error', 'error': 'error message'}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "{'message': '缺少必要欄位', 'error': 'error message'}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Logs in a user with the provided email and password, and generates a JWT token for the user",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{'message': 'login success', 'token': 'JWT token'}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "{'message': '帳號或密碼錯誤', 'error': 'error message'}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "user register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "register",
                "parameters": [
                    {
                        "description": "registerForm",
                        "name": "registerForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.registerForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"register success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "{\"message\": \"Email 重複\", \"error\": \"duplicate email\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"message\": \"register failed\", \"error\": \"server error\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.registerForm": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "equip.updateFrom": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                }
            }
        },
        "equip.weightForm": {
            "type": "object",
            "required": [
                "weights"
            ],
            "properties": {
                "weights": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type Bearer followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header",
            "scopes": {
                "write": " Grants write access"
            }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
