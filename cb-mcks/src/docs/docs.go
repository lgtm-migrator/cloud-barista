// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://cloud-barista.github.io",
            "email": "contact-to-cloud-barista@googlegroups.com"
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
        "/healthy": {
            "get": {
                "description": "for health check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Default"
                ],
                "summary": "Health Check",
                "operationId": "Healthy",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/mcir/connections/{connection}/specs": {
            "get": {
                "description": "List Specs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mcir"
                ],
                "summary": "List Specs",
                "operationId": "List Spec",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Connection Name",
                        "name": "connection",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "Y",
                            "N"
                        ],
                        "type": "string",
                        "description": "string enums",
                        "name": "control-plane",
                        "in": "query",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "if Control-Plane, \u003e= 2",
                        "name": "cpu-min",
                        "in": "query"
                    },
                    {
                        "maximum": 99999,
                        "minimum": 1,
                        "type": "integer",
                        "description": " \u003c= 99999",
                        "name": "cpu-max",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": " if Control-Plane, \u003e= 2",
                        "name": "memory-min",
                        "in": "query"
                    },
                    {
                        "maximum": 99999,
                        "minimum": 1,
                        "type": "integer",
                        "description": " \u003c= 99999",
                        "name": "memory-max",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.SpecList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    }
                }
            }
        },
        "/ns/{namespace}/clusters": {
            "get": {
                "description": "List all Clusters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cluster"
                ],
                "summary": "List all Clusters",
                "operationId": "ListCluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Namespace ID",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ClusterList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cluster"
                ],
                "summary": "Create Cluster",
                "operationId": "CreateCluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Namespace ID",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "1.18",
                            "1.23"
                        ],
                        "type": "string",
                        "description": "string enums",
                        "name": "minorversion",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Patch version",
                        "name": "patchversion",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request Body to create cluster",
                        "name": "ClusterReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.ClusterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Cluster"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    }
                }
            }
        },
        "/ns/{namespace}/clusters/{cluster}": {
            "get": {
                "description": "Get Cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cluster"
                ],
                "summary": "Get Cluster",
                "operationId": "GetCluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Namespace ID",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Cluster Name",
                        "name": "cluster",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Cluster"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cluster"
                ],
                "summary": "Delete Cluster",
                "operationId": "DeleteCluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Namespace ID",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Cluster Name",
                        "name": "cluster",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    }
                }
            }
        },
        "/ns/{namespace}/clusters/{cluster}/nodes": {
            "get": {
                "description": "List all Nodes in specified Cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Node"
                ],
                "summary": "List all Nodes in specified Cluster",
                "operationId": "ListNode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Namespace ID",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Cluster Name",
                        "name": "cluster",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.NodeList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    }
                }
            },
            "post": {
                "description": "Add Node in specified Cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Node"
                ],
                "summary": "Add Node in specified Cluster",
                "operationId": "AddNode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Namespace ID",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Cluster Name",
                        "name": "cluster",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request Body to add node",
                        "name": "nodeReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.NodeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Node"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    }
                }
            }
        },
        "/ns/{namespace}/clusters/{cluster}/nodes/{node}": {
            "get": {
                "description": "Get Node in specified Cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Node"
                ],
                "summary": "Get Node in specified Cluster",
                "operationId": "GetNode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Namespace ID",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Cluster Name",
                        "name": "cluster",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Node Name",
                        "name": "node",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Node"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove Node in specified Cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Node"
                ],
                "summary": "Remove Node in specified Cluster",
                "operationId": "RemoveNode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Namespace ID",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Cluster Name",
                        "name": "cluster",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Node Name",
                        "name": "node",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Status"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.ClusterConfigKubernetesReq": {
            "type": "object",
            "properties": {
                "networkCni": {
                    "type": "string",
                    "enum": [
                        "canal",
                        "kilo"
                    ],
                    "example": "kilo"
                },
                "podCidr": {
                    "type": "string",
                    "example": "10.244.0.0/16"
                },
                "serviceCidr": {
                    "type": "string",
                    "example": "10.96.0.0/12"
                },
                "serviceDnsDomain": {
                    "type": "string",
                    "example": "cluster.local"
                }
            }
        },
        "app.ClusterConfigReq": {
            "type": "object",
            "properties": {
                "kubernetes": {
                    "$ref": "#/definitions/app.ClusterConfigKubernetesReq"
                }
            }
        },
        "app.ClusterReq": {
            "type": "object",
            "properties": {
                "config": {
                    "$ref": "#/definitions/app.ClusterConfigReq"
                },
                "controlPlane": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app.NodeSetReq"
                    }
                },
                "description": {
                    "type": "string"
                },
                "installMonAgent": {
                    "type": "string",
                    "default": "yes",
                    "example": "no"
                },
                "label": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "example": "cluster-01"
                },
                "worker": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app.NodeSetReq"
                    }
                }
            }
        },
        "app.NodeReq": {
            "type": "object",
            "properties": {
                "controlPlane": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app.NodeSetReq"
                    }
                },
                "worker": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app.NodeSetReq"
                    }
                }
            }
        },
        "app.NodeSetReq": {
            "type": "object",
            "properties": {
                "connection": {
                    "type": "string",
                    "example": "config-aws-ap-northeast-2"
                },
                "count": {
                    "type": "integer",
                    "example": 3
                },
                "spec": {
                    "type": "string",
                    "example": "t2.medium"
                }
            }
        },
        "app.Status": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "kind": {
                    "type": "string"
                },
                "message": {
                    "type": "string",
                    "example": "Any message"
                }
            }
        },
        "model.Cluster": {
            "type": "object",
            "properties": {
                "clusterConfig": {
                    "type": "string"
                },
                "cpLeader": {
                    "type": "string"
                },
                "createdTime": {
                    "type": "string",
                    "example": "2022-01-02T12:00:00Z"
                },
                "description": {
                    "type": "string"
                },
                "installMonAgent": {
                    "type": "string",
                    "default": "yes",
                    "example": "no"
                },
                "k8sVersion": {
                    "type": "string"
                },
                "kind": {
                    "type": "string"
                },
                "label": {
                    "type": "string"
                },
                "mcis": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "networkCni": {
                    "type": "string",
                    "enum": [
                        "canal",
                        "kilo"
                    ]
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Node"
                    }
                },
                "status": {
                    "$ref": "#/definitions/model.ClusterStatus"
                }
            }
        },
        "model.ClusterList": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Cluster"
                    }
                },
                "kind": {
                    "type": "string"
                }
            }
        },
        "model.ClusterStatus": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "phase": {
                    "type": "string",
                    "enum": [
                        "Pending",
                        "Provisioning",
                        "Provisioned",
                        "Failed"
                    ]
                },
                "reason": {
                    "type": "string"
                }
            }
        },
        "model.Node": {
            "type": "object",
            "properties": {
                "createdTime": {
                    "type": "string",
                    "example": "2022-01-02T12:00:00Z"
                },
                "credential": {
                    "type": "string"
                },
                "csp": {
                    "type": "string",
                    "enum": [
                        "aws",
                        "gcp",
                        "azure",
                        "alibaba",
                        "tencent",
                        "openstack",
                        "ibm",
                        "cloudit"
                    ]
                },
                "cspLabel": {
                    "type": "string"
                },
                "kind": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "publicIp": {
                    "type": "string"
                },
                "regionLabel": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "control-plane",
                        "worker"
                    ]
                },
                "spec": {
                    "type": "string"
                },
                "zoneLabel": {
                    "type": "string"
                }
            }
        },
        "model.NodeList": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Node"
                    }
                },
                "kind": {
                    "type": "string"
                }
            }
        },
        "service.SpecList": {
            "type": "object",
            "properties": {
                "connectionName": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.Vmspecs"
                    }
                },
                "kind": {
                    "type": "string"
                }
            }
        },
        "service.Vmspecs": {
            "type": "object",
            "properties": {
                "cpu": {
                    "type": "object",
                    "properties": {
                        "clock": {
                            "description": "output - GHz",
                            "type": "string"
                        },
                        "count": {
                            "description": "output",
                            "type": "string"
                        }
                    }
                },
                "memory": {
                    "description": "output",
                    "type": "string"
                },
                "name": {
                    "description": "output",
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

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "latest",
	Host:        "localhost:1470",
	BasePath:    "/mcks",
	Schemes:     []string{},
	Title:       "CB-MCKS REST API",
	Description: "CB-MCKS REST API",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
