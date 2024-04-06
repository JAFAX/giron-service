// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Gary Greene",
            "url": "https://github.com/JAFAX/giron-service"
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
        "/building": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Create a new building",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Create a new building",
                "parameters": [
                    {
                        "description": "Building data",
                        "name": "building",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ProposedBuilding"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/building/{id}": {
            "get": {
                "description": "Retrieve building by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Retrieve building by Id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Building"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update building information",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Update building information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/buildings": {
            "get": {
                "description": "Retrieve list of all panels",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Retrieve list of all panels",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BuildingList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/panel": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Create a new panel event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "panels"
                ],
                "summary": "Create a new panel event",
                "parameters": [
                    {
                        "description": "Panel data",
                        "name": "panel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Panel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/panel/{id}": {
            "get": {
                "description": "Retrieve panel by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "panels"
                ],
                "summary": "Retrieve panel by Id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Panel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/panel/{id}/location": {
            "get": {
                "description": "Retrieve panel location by the panel Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "panels"
                ],
                "summary": "Retrieve panel location by the panel Id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Location"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            },
            "post": {
                "description": "Set panel location",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "panels"
                ],
                "summary": "Set panel location",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/panel/{id}/schedule": {
            "get": {
                "description": "Retrieve panel schedule by the panel Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "panels"
                ],
                "summary": "Retrieve panel schedule by the panel Id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Schedule"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/panels": {
            "get": {
                "description": "Retrieve list of all approved panels",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "panels"
                ],
                "summary": "Retrieve list of all approved panels",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PanelList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/panels/all": {
            "get": {
                "description": "Retrieve list of all panels",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "panels"
                ],
                "summary": "Retrieve list of all panels",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PanelList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Add a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ProposedUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/user/id/{id}": {
            "get": {
                "description": "Retrieve a user by their Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Retrieve a user by their Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SafeUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/user/name/{name}": {
            "get": {
                "description": "Retrieve a user by their UserName",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Retrieve a user by their UserName",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SafeUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/user/{name}": {
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Delete a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Change password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Change password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Password data",
                        "name": "changePassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PasswordChange"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/user/{name}/status": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Retrieve a user's active status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Retrieve a user's active status. Can be either 'enabled' or 'locked'",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserStatusMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Set a user's active status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Set a user's active status. Can be either 'enabled' or 'locked'",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserStatusMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Retrieve list of all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Retrieve list of all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UsersList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.FailureMsg"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.SafeUser": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer"
                },
                "creationDate": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "model.Building": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer"
                },
                "city": {
                    "type": "string"
                },
                "creationDateTime": {
                    "type": "string"
                },
                "creatorId": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                }
            }
        },
        "model.BuildingList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Building"
                    }
                }
            }
        },
        "model.FailureMsg": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.Location": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer"
                },
                "buildingId": {
                    "type": "integer"
                },
                "creationDateTime": {
                    "type": "string"
                },
                "creatorId": {
                    "type": "integer"
                },
                "floorId": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                }
            }
        },
        "model.Panel": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer"
                },
                "approvalDateTime": {
                    "type": "string"
                },
                "approvalStatus": {
                    "type": "boolean"
                },
                "approvedById": {
                    "type": "integer"
                },
                "creationDateTime": {
                    "type": "string"
                },
                "creatorId": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "durationInMinutes": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "panelRequestorEmail": {
                    "type": "string"
                },
                "scheduledTime": {
                    "type": "string"
                },
                "topic": {
                    "type": "string"
                }
            }
        },
        "model.PanelList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Panel"
                    }
                }
            }
        },
        "model.PasswordChange": {
            "type": "object",
            "properties": {
                "newPassword": {
                    "type": "string"
                },
                "oldPassword": {
                    "type": "string"
                }
            }
        },
        "model.ProposedBuilding": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "creatorId": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                }
            }
        },
        "model.ProposedUser": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer"
                },
                "creationDate": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "model.Schedule": {
            "type": "object",
            "properties": {
                "durationInMinutes": {
                    "type": "integer"
                },
                "startTime": {
                    "type": "string"
                }
            }
        },
        "model.SuccessMsg": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer"
                },
                "creationDate": {
                    "type": "string"
                },
                "lastChangedDate": {
                    "type": "string"
                },
                "passwordHash": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "model.UserStatusMsg": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "userStatus": {
                    "type": "string"
                }
            }
        },
        "model.UsersList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
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
	Version:          "0.0.15",
	Host:             "localhost:5000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Giron-Service",
	Description:      "An API for managing panel events",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
