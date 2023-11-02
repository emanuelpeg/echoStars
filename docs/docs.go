// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/health/check": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/health/info": {
            "get": {
                "description": "get the info of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the info of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sysinfo.SysInfo"
                        }
                    }
                }
            }
        },
        "/info/database/check": {
            "get": {
                "description": "get the status of the database.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of the database.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/info/getFromDb": {
            "get": {
                "description": "It is a test of the database.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "It is a test of the database.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sysinfo.SysInfo"
                        }
                    }
                }
            }
        },
        "/info/save": {
            "get": {
                "description": "It is a test of the database.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "It is a test of the database.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sysinfo.SysInfo"
                        }
                    }
                }
            }
        },
        "/notification/email": {
            "post": {
                "description": "Send an email notification using the EmailNotification factory. Expects JSON data in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "email"
                ],
                "summary": "Send an email notification.",
                "parameters": [
                    {
                        "description": "Email request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/notification.SendEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/notification.SendEmailResponse"
                        }
                    }
                }
            }
        },
        "/servers": {
            "get": {
                "description": "List all registered servers.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "summary": "Lists registered servers.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/server.Server"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates new server for monitoring using a server struct as input.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "summary": "Creates a new server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.Server"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete  server from server id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "summary": "delete a server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "notification.SendEmailRequest": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "recipient": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                }
            }
        },
        "notification.SendEmailResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "server.Server": {
            "type": "object",
            "properties": {
                "frequency": {
                    "type": "integer"
                },
                "hostname": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "mailTo": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "sysinfo.SysInfo": {
            "type": "object",
            "properties": {
                "disk": {
                    "type": "integer"
                },
                "hostname": {
                    "type": "string"
                },
                "platform": {
                    "type": "string"
                },
                "ram": {
                    "type": "integer"
                },
                "ram available": {
                    "type": "integer"
                },
                "ram free": {
                    "type": "integer"
                },
                "uptime": {
                    "type": "integer"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
