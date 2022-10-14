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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/rulesheets": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create Rulesheet description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rulesheet"
                ],
                "summary": "Create Rulesheet",
                "parameters": [
                    {
                        "description": "Rulesheet body",
                        "name": "Rulesheet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.Rulesheet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.Rulesheet"
                        },
                        "headers": {
                            "Authorization": {
                                "type": "string",
                                "description": "token access"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Format",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    }
                }
            }
        },
        "/rulesheets/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List Rulesheet description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rulesheet"
                ],
                "summary": "List Rulesheets",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "Total of results",
                        "name": "count",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Max length of the array returned",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number that is multiplied by 'limit' to calculate the offset",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/v1.Rulesheet"
                            }
                        },
                        "headers": {
                            "Authorization": {
                                "type": "string",
                                "description": "token access"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Format",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    }
                }
            }
        },
        "/rulesheets/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get Rulesheet by ID description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rulesheet"
                ],
                "summary": "Get Rulesheet by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Rulesheet ID",
                        "name": "id",
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
                                "$ref": "#/definitions/v1.Rulesheet"
                            }
                        },
                        "headers": {
                            "Authorization": {
                                "type": "string",
                                "description": "token access"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Format",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update Rulesheet by ID description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rulesheet"
                ],
                "summary": "Update Rulesheet by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Rulesheet ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Rulesheet body",
                        "name": "rulesheet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.Rulesheet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/v1.Rulesheet"
                            }
                        },
                        "headers": {
                            "Authorization": {
                                "type": "string",
                                "description": "token access"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Format",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete Rulesheet by ID description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rulesheet"
                ],
                "summary": "Delete Rulesheet by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Rulesheet ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "Authorization": {
                                "type": "string",
                                "description": "token access"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Format",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "v1.Error": {
            "type": "object",
            "properties": {
                "error": {},
                "validation_errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.ValidationError"
                    }
                }
            }
        },
        "v1.Rulesheet": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "features": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": true
                    }
                },
                "hasStringRule": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "parameters": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": true
                    }
                },
                "rules": {
                    "type": "object",
                    "additionalProperties": true
                },
                "slug": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "v1.ValidationError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "field": {
                    "type": "string"
                },
                "tag": {
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
    },
    "x-extension-openapi": {
        "example": "value on a json format"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9007",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "FeatWS API",
	Description:      "API Project to provide operations to manage FeatWS knowledge repositories rules",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
