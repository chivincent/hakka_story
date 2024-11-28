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
        "/admin/audio/generate": {
            "post": {
                "description": "Generate an audio from a text",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin audio"
                ],
                "summary": "Generate Audio",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Text prompt",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/haudio.GenerateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/haudio.GenerateResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    }
                }
            }
        },
        "/admin/audio/upload": {
            "post": {
                "description": "Upload Audio",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "admin audio"
                ],
                "summary": "Upload Audio",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "音檔",
                        "name": "audio",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/haudio.UploadAudioRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/haudio.UploadAudioResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    }
                }
            }
        },
        "/admin/auth": {
            "get": {
                "description": "Get the admin status of the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin auth"
                ],
                "summary": "Check if the user is an admin",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/hauth.AuthResponse"
                        }
                    }
                }
            }
        },
        "/admin/image/generate": {
            "post": {
                "description": "Generate an image from a text prompt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin image"
                ],
                "summary": "Generate Image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Text prompt",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/himage.GenerateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/himage.GenerateResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    }
                }
            }
        },
        "/admin/image/upload": {
            "post": {
                "description": "Upload Image",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "admin image"
                ],
                "summary": "Upload Image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "圖片",
                        "name": "image",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/himage.UploadImageRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/himage.UploadImageResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    }
                }
            }
        },
        "/admin/story": {
            "post": {
                "description": "Create Story",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin story"
                ],
                "summary": "Create Story",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "故事內容",
                        "name": "story",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hstory.UpsertStoryRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hstory.StoryResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    }
                }
            }
        },
        "/admin/story/:id": {
            "put": {
                "description": "Update Story by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin story"
                ],
                "summary": "Update Story",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "故事內容",
                        "name": "story",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hstory.UpsertStoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hstory.FullStoryResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    }
                }
            }
        },
        "/admin/story/{id}": {
            "delete": {
                "description": "Delete a story by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin story"
                ],
                "summary": "Delete a story by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Story ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    }
                }
            }
        },
        "/admin/translate/hakka": {
            "post": {
                "description": "Translate a given text to Hakka language",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin translate"
                ],
                "summary": "Translate Text to Hakka",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Text to translate",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/htranslate.TranslateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/htranslate.TranslateResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    }
                }
            }
        },
        "/categories": {
            "post": {
                "description": "Create a new category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin categories"
                ],
                "summary": "Create a new category",
                "parameters": [
                    {
                        "description": "Category to create",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hcategory.UpsertRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/hcategory.CategoryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    }
                }
            }
        },
        "/category": {
            "get": {
                "description": "Get list of categories by name",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "List categories",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/response.Response"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "data": {
                                                "$ref": "#/definitions/hcategory.CategoryResponse"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    }
                }
            }
        },
        "/category/:id": {
            "put": {
                "description": "Update category by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Update category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Category data",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hcategory.UpsertRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hcategory.CategoryResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    }
                }
            }
        },
        "/story": {
            "get": {
                "description": "Get all stories",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stories"
                ],
                "summary": "List stories",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "Category names",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hstory.StoryResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    }
                }
            }
        },
        "/story/:id": {
            "get": {
                "description": "Get story by id with pages",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stories"
                ],
                "summary": "Get story",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Story ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hstory.FullStoryResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseBase"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "haudio.GenerateRequest": {
            "type": "object",
            "properties": {
                "prompt": {
                    "type": "string"
                }
            }
        },
        "haudio.GenerateResponse": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "haudio.UploadAudioRequest": {
            "type": "object",
            "required": [
                "audio"
            ],
            "properties": {
                "audio": {
                    "type": "string"
                }
            }
        },
        "haudio.UploadAudioResponse": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "hauth.AuthResponse": {
            "type": "object",
            "properties": {
                "isAdmin": {
                    "type": "boolean"
                }
            }
        },
        "hcategory.CategoryResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "hcategory.UpsertRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "himage.GenerateRequest": {
            "type": "object",
            "properties": {
                "prompt": {
                    "type": "string"
                }
            }
        },
        "himage.GenerateResponse": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "himage.UploadImageRequest": {
            "type": "object",
            "required": [
                "image"
            ],
            "properties": {
                "image": {
                    "type": "string"
                }
            }
        },
        "himage.UploadImageResponse": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "hstory.AudioResponse": {
            "type": "object",
            "properties": {
                "audio_url": {
                    "type": "string"
                },
                "dialect": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "hstory.CategoryResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "hstory.FullStoryResponse": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hstory.CategoryResponse"
                    }
                },
                "cover_image": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "pages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hstory.PageResponse"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "hstory.PageResponse": {
            "type": "object",
            "properties": {
                "audios": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hstory.AudioResponse"
                    }
                },
                "content_cn": {
                    "type": "string"
                },
                "content_hakka": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "page_number": {
                    "type": "integer"
                }
            }
        },
        "hstory.StoryResponse": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hstory.CategoryResponse"
                    }
                },
                "cover_image": {
                    "type": "string"
                },
                "description": {
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
        "hstory.UpsertAudioRequest": {
            "type": "object",
            "required": [
                "audio_url",
                "dialect"
            ],
            "properties": {
                "audio_url": {
                    "type": "string"
                },
                "dialect": {
                    "type": "string"
                }
            }
        },
        "hstory.UpsertCategoryRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "hstory.UpsertPageRequest": {
            "type": "object",
            "required": [
                "content_cn",
                "content_hakka",
                "page_number"
            ],
            "properties": {
                "audios": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hstory.UpsertAudioRequest"
                    }
                },
                "content_cn": {
                    "type": "string"
                },
                "content_hakka": {
                    "type": "string"
                },
                "page_number": {
                    "type": "integer"
                }
            }
        },
        "hstory.UpsertStoryRequest": {
            "type": "object",
            "required": [
                "description",
                "pages",
                "title"
            ],
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hstory.UpsertCategoryRequest"
                    }
                },
                "cover_image": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "pages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hstory.UpsertPageRequest"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "htranslate.TranslateRequest": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string"
                }
            }
        },
        "htranslate.TranslateResponse": {
            "type": "object",
            "properties": {
                "hakka": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "response.ResponseBase": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
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
