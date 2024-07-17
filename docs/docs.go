// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/login/sign": {
            "post": {
                "tags": [
                    "授权"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户登录",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SignRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.HttpResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/request.SignResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "创建用户",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/result.HttpResult"
                        }
                    }
                }
            }
        },
        "/user/delete/{id}": {
            "delete": {
                "tags": [
                    "用户管理"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/result.HttpResult"
                        }
                    }
                }
            }
        },
        "/user/info/{id}": {
            "get": {
                "tags": [
                    "用户管理"
                ],
                "summary": "用户详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.HttpResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SysUser"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/page": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.HttpResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/response.PageResponse"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "list": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/models.SysUser"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/update": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "description": "更新用户",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/result.HttpResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.SysUser": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "delFlag": {
                    "description": "删除状态 0 正常 1删除",
                    "type": "integer"
                },
                "deptId": {
                    "description": "部门id",
                    "type": "integer"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "gender": {
                    "description": "性别 0未知 1男 2 女",
                    "type": "integer"
                },
                "id": {
                    "description": "主键",
                    "type": "integer"
                },
                "nickName": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "status": {
                    "description": "启用状态 0 启用 1 禁用",
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "request.CreateUserRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "delFlag": {
                    "description": "删除状态 0 正常 1删除",
                    "type": "integer"
                },
                "deptId": {
                    "description": "部门id",
                    "type": "integer"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "gender": {
                    "description": "性别 0未知 1男 2 女",
                    "type": "integer"
                },
                "nickName": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "status": {
                    "description": "启用状态 0 启用 1 禁用",
                    "type": "integer"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "request.SignRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.SignResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "expiresIn": {
                    "type": "integer"
                }
            }
        },
        "request.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "delFlag": {
                    "description": "删除状态 0 正常 1删除",
                    "type": "integer"
                },
                "deptId": {
                    "description": "部门id",
                    "type": "integer"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "gender": {
                    "description": "性别 0未知 1男 2 女",
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "nickName": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "status": {
                    "description": "启用状态 0 启用 1 禁用",
                    "type": "integer"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "response.PageResponse": {
            "type": "object",
            "properties": {
                "list": {},
                "total": {
                    "type": "integer"
                }
            }
        },
        "result.HttpResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "rest_go_api",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
