// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PreventiveWorks"
                ],
                "summary": "отображение всех профилактических работ",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.PreventiveWork"
                            }
                        }
                    }
                }
            }
        },
        "/event": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "отображение всех профилактических работ",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Event"
                            }
                        }
                    }
                }
            }
        },
        "/new_work": {
            "post": {
                "tags": [
                    "NewPreventiveWork"
                ],
                "summary": "добавление новой профилактической работы",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Имя сервиса",
                        "name": "name_service",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Дата создания профил. работы",
                        "name": "create_at",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Дата окончания профил. работы",
                        "name": "deadline",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Название профил. работы",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Описание профил. работы",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PreventiveWork"
                ],
                "summary": "отображение профилактической работы по id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "PreventiveWork id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PreventiveWork"
                        }
                    }
                }
            }
        },
        "/{id}/new_event": {
            "put": {
                "tags": [
                    "NewPreventiveWork"
                ],
                "summary": "добавление новой профилактической работы",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id профилактической работы",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Статус события",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Дата создания события",
                        "name": "create_at",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Дата окончания события",
                        "name": "deadline",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Описание события",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Event": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string"
                },
                "deadline": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "id_preventive_work": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.PreventiveWork": {
            "type": "object",
            "properties": {
                "count_event": {
                    "type": "integer"
                },
                "create_at": {
                    "type": "string"
                },
                "deadline": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Event"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "id_service": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "preventive-works",
	Description:      "API для отслеживания профилактических работ",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}