// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "joe.nghiem",
            "email": "joe.nghiem.goldenowl@gmail.com"
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
        "/api/v1/user": {
            "get": {
                "description": "GetByUsername by username",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "username param",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/gorms.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseFailed"
                        }
                    }
                }
            },
            "post": {
                "description": "Create by username, password, email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "body params",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UserCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/gorms.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseFailed"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseFailed"
                        }
                    }
                }
            }
        },
        "/api/v1/user/login": {
            "post": {
                "description": "login by username \u0026 password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login user to system",
                "parameters": [
                    {
                        "description": "body params",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/responses.UserLoginResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseFailed"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseFailed"
                        }
                    }
                }
            }
        },
        "/api/v1/user/logout": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Logout user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseFailed"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorms.User": {
            "type": "object",
            "properties": {
                "approved_status": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "requests.UserCreateRequest": {
            "type": "object",
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
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
        "requests.UserLoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "responses.UserLoginResponse": {
            "type": "object",
            "properties": {
                "expires": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "utils.ResponseFailed": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "string"
                }
            }
        },
        "utils.ResponseSuccess": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "GoldenOwl Gin API",
	Description:      "This API is for GoldenOwl API application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}