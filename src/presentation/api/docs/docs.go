// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://speedia.net/tos/",
        "contact": {
            "name": "Speedia Engineering",
            "url": "https://speedia.net/",
            "email": "eng+swagger@speedia.net"
        },
        "license": {
            "name": "SPEEDIA WEB SERVICES © 2023. All Rights Reserved.",
            "url": "https://speedia.net/tos/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "List accs.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "GetAccounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Account"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update an account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "UpdateAccount",
                "parameters": [
                    {
                        "description": "UpdateAccount",
                        "name": "updateAccountDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "AccountUpdated message or NewKeyString",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Add a new account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "AddNewAccount",
                "parameters": [
                    {
                        "description": "NewAccount",
                        "name": "addAccountDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddAccount"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "AccountCreated",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/account/{accountId}/": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete an account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "DeleteAccount",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccountId",
                        "name": "accountId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "AccountDeleted",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/auth/login/": {
            "post": {
                "description": "Generate JWT with credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "GenerateJwtWithCredentials",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "loginDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.AccessToken"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/o11y/overview/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Show system information and resource usage.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "o11y"
                ],
                "summary": "O11yOverview",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.O11yOverview"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AddAccount": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "quota": {
                    "$ref": "#/definitions/valueObject.AccountQuota"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.Login": {
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
        "dto.UpdateAccount": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "quota": {
                    "$ref": "#/definitions/valueObject.AccountQuota"
                },
                "shouldUpdateApiKey": {
                    "type": "boolean"
                }
            }
        },
        "entity.AccessToken": {
            "type": "object",
            "properties": {
                "expiresIn": {
                    "type": "integer"
                },
                "tokenStr": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/valueObject.AccessTokenType"
                }
            }
        },
        "entity.Account": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "integer"
                },
                "groupId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "quota": {
                    "$ref": "#/definitions/valueObject.AccountQuota"
                },
                "quotaUsage": {
                    "$ref": "#/definitions/valueObject.AccountQuota"
                },
                "updatedAt": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.O11yOverview": {
            "type": "object",
            "properties": {
                "currentUsage": {
                    "$ref": "#/definitions/valueObject.CurrentResourceUsage"
                },
                "hostname": {
                    "type": "string"
                },
                "publicIp": {
                    "type": "string"
                },
                "specs": {
                    "$ref": "#/definitions/valueObject.HardwareSpecs"
                },
                "uptimeSecs": {
                    "type": "integer"
                }
            }
        },
        "valueObject.AccessTokenType": {
            "type": "string",
            "enum": [
                "sessionToken",
                "accountApiKey"
            ],
            "x-enum-varnames": [
                "sessionToken",
                "accountApiKey"
            ]
        },
        "valueObject.AccountQuota": {
            "type": "object",
            "properties": {
                "cpuCores": {
                    "type": "number"
                },
                "diskBytes": {
                    "type": "integer"
                },
                "inodes": {
                    "type": "integer"
                },
                "memoryBytes": {
                    "type": "integer"
                }
            }
        },
        "valueObject.CurrentResourceUsage": {
            "type": "object",
            "properties": {
                "cpuUsagePercent": {
                    "type": "number"
                },
                "memUsagePercent": {
                    "type": "number"
                },
                "netUsage": {
                    "$ref": "#/definitions/valueObject.NetUsage"
                },
                "storageUsage": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/valueObject.DiskInfo"
                    }
                }
            }
        },
        "valueObject.DiskInfo": {
            "type": "object",
            "properties": {
                "availableBytes": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "totalBytes": {
                    "type": "integer"
                },
                "usedBytes": {
                    "type": "integer"
                }
            }
        },
        "valueObject.HardwareSpecs": {
            "type": "object",
            "properties": {
                "cpuCores": {
                    "type": "integer"
                },
                "cpuFrequency": {
                    "type": "number"
                },
                "cpuModel": {
                    "type": "string"
                },
                "memoryTotal": {
                    "type": "integer"
                },
                "storageTotal": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/valueObject.DiskInfo"
                    }
                }
            }
        },
        "valueObject.NetUsage": {
            "type": "object",
            "properties": {
                "receivedBytes": {
                    "type": "integer"
                },
                "sentBytes": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" + JWT token or API key.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:10001",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "SfmApi",
	Description:      "Speedia FleetManager API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
