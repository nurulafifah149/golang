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
        "/book": {
            "get": {
                "description": "Get all data Book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Book",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.BookResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Get all data Book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Book",
                "parameters": [
                    {
                        "description": "Input Data Buku",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BookCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.BookResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/book/{id}": {
            "get": {
                "description": "Get data Book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "book search by id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sukses Mengambil data Buku",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.BookResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "Get data Book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "book update by id",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "description": "Input Data Buku",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BookCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sukses Mengambil data Buku",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.BookResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete data Book By Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Delete book by id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/model.WebResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.BookCreateRequest": {
            "type": "object",
            "required": [
                "author",
                "desc",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.BookResponse": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.WebResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
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
