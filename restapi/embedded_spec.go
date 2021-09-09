// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Kube for Edge Management",
    "title": "Kube4EdgeManagement",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/api/k4e-management/v1",
  "paths": {
    "/control/{device_id}/in": {
      "get": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "GetControlMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    },
    "/control/{device_id}/out": {
      "post": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "PostControlMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          },
          {
            "name": "message",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    },
    "/data/{device_id}/in": {
      "get": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "GetDataMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    },
    "/data/{device_id}/out": {
      "post": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "PostDataMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          },
          {
            "name": "message",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    }
  },
  "definitions": {
    "boot": {
      "type": "object",
      "properties": {
        "current_boot_mode": {
          "type": "string"
        },
        "pxe_interface": {
          "type": "string"
        }
      }
    },
    "cpu": {
      "type": "object",
      "properties": {
        "architecture": {
          "type": "string"
        },
        "count": {
          "type": "integer"
        },
        "flags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "frequency": {
          "type": "number"
        },
        "model_name": {
          "type": "string"
        }
      }
    },
    "data-configuration": {
      "type": "object",
      "properties": {
        "paths": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/data-path"
          }
        }
      }
    },
    "data-path": {
      "description": "Device-to-control plane paths mapping",
      "type": "object",
      "properties": {
        "source": {
          "description": "Path in the workload container",
          "type": "string"
        },
        "target": {
          "description": "Path in the control plane storage",
          "type": "string"
        }
      }
    },
    "device-configuration": {
      "type": "object",
      "properties": {
        "heartbeat": {
          "$ref": "#/definitions/heartbeat-configuration"
        },
        "storage": {
          "$ref": "#/definitions/storage-configuration"
        }
      }
    },
    "device-configuration-message": {
      "type": "object",
      "properties": {
        "configuration": {
          "$ref": "#/definitions/device-configuration"
        },
        "device_id": {
          "description": "Device identifier",
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "workloads": {
          "description": "List of workloads deployed to the device",
          "$ref": "#/definitions/workload-list"
        }
      }
    },
    "disk": {
      "type": "object",
      "properties": {
        "bootable": {
          "type": "boolean"
        },
        "by_id": {
          "description": "by-id is the World Wide Number of the device which guaranteed to be unique for every storage device",
          "type": "string"
        },
        "by_path": {
          "description": "by-path is the shortest physical path to the device",
          "type": "string"
        },
        "drive_type": {
          "type": "string"
        },
        "hctl": {
          "type": "string"
        },
        "id": {
          "description": "Determine the disk's unique identifier which is the by-id field if it exists and fallback to the by-path field otherwise",
          "type": "string"
        },
        "io_perf": {
          "$ref": "#/definitions/io_perf"
        },
        "is_installation_media": {
          "description": "Whether the disk appears to be an installation media or not",
          "type": "boolean"
        },
        "model": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "path": {
          "type": "string"
        },
        "serial": {
          "type": "string"
        },
        "size_bytes": {
          "type": "integer"
        },
        "smart": {
          "type": "string"
        },
        "vendor": {
          "type": "string"
        },
        "wwn": {
          "type": "string"
        }
      }
    },
    "gpu": {
      "type": "object",
      "properties": {
        "address": {
          "description": "Device address (for example \"0000:00:02.0\")",
          "type": "string"
        },
        "device_id": {
          "description": "ID of the device (for example \"3ea0\")",
          "type": "string"
        },
        "name": {
          "description": "Product name of the device (for example \"UHD Graphics 620 (Whiskey Lake)\")",
          "type": "string"
        },
        "vendor": {
          "description": "The name of the device vendor (for example \"Intel Corporation\")",
          "type": "string"
        },
        "vendor_id": {
          "description": "ID of the vendor (for example \"8086\")",
          "type": "string"
        }
      }
    },
    "hardware-info": {
      "type": "object",
      "properties": {
        "boot": {
          "$ref": "#/definitions/boot"
        },
        "cpu": {
          "$ref": "#/definitions/cpu"
        },
        "disks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/disk"
          }
        },
        "gpus": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/gpu"
          }
        },
        "hostname": {
          "type": "string"
        },
        "interfaces": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/interface"
          }
        },
        "memory": {
          "$ref": "#/definitions/memory"
        },
        "system_vendor": {
          "$ref": "#/definitions/system_vendor"
        }
      }
    },
    "hardware-profile-configuration": {
      "type": "object",
      "properties": {
        "include": {
          "type": "boolean"
        },
        "scope": {
          "type": "string",
          "enum": [
            "full",
            "delta"
          ]
        }
      }
    },
    "heartbeat": {
      "type": "object",
      "properties": {
        "hardware": {
          "description": "Hardware information",
          "$ref": "#/definitions/hardware-info"
        },
        "status": {
          "type": "string",
          "enum": [
            "up",
            "degraded"
          ]
        },
        "time": {
          "type": "string",
          "format": "date-time"
        },
        "version": {
          "type": "string"
        },
        "workloads": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/workload-status"
          }
        }
      }
    },
    "heartbeat-configuration": {
      "type": "object",
      "properties": {
        "hardware_profile": {
          "$ref": "#/definitions/hardware-profile-configuration"
        },
        "period_seconds": {
          "type": "integer"
        }
      }
    },
    "interface": {
      "type": "object",
      "properties": {
        "biosdevname": {
          "type": "string"
        },
        "client_id": {
          "type": "string"
        },
        "flags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "has_carrier": {
          "type": "boolean"
        },
        "ipv4_addresses": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "ipv6_addresses": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "mac_address": {
          "type": "string"
        },
        "mtu": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "product": {
          "type": "string"
        },
        "speed_mbps": {
          "type": "integer"
        },
        "vendor": {
          "type": "string"
        }
      }
    },
    "io_perf": {
      "type": "object",
      "properties": {
        "sync_duration": {
          "description": "99th percentile of fsync duration in milliseconds",
          "type": "integer"
        }
      }
    },
    "memory": {
      "type": "object",
      "properties": {
        "physical_bytes": {
          "type": "integer"
        },
        "usable_bytes": {
          "type": "integer"
        }
      }
    },
    "message": {
      "type": "object",
      "properties": {
        "content": {
          "description": "Content"
        },
        "directive": {
          "type": "string"
        },
        "message_id": {
          "type": "string"
        },
        "metadata": {
          "type": "object"
        },
        "response_to": {
          "type": "string"
        },
        "sent": {
          "type": "string",
          "format": "date-time"
        },
        "type": {
          "type": "string",
          "enum": [
            "connection-status",
            "command",
            "event",
            "data"
          ]
        },
        "version": {
          "type": "integer"
        }
      }
    },
    "registration-info": {
      "type": "object",
      "properties": {
        "hardware": {
          "description": "Hardware information",
          "$ref": "#/definitions/hardware-info"
        },
        "os_image_id": {
          "type": "string"
        }
      }
    },
    "s3-storage-configuration": {
      "type": "object",
      "properties": {
        "aws_access_key_id": {
          "type": "string"
        },
        "aws_ca_bundle": {
          "type": "string"
        },
        "aws_secret_access_key": {
          "type": "string"
        },
        "bucket_host": {
          "type": "string"
        },
        "bucket_name": {
          "type": "string"
        },
        "bucket_port": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "storage-configuration": {
      "type": "object",
      "properties": {
        "s3": {
          "$ref": "#/definitions/s3-storage-configuration"
        }
      }
    },
    "system_vendor": {
      "type": "object",
      "properties": {
        "manufacturer": {
          "type": "string"
        },
        "product_name": {
          "type": "string"
        },
        "serial_number": {
          "type": "string"
        },
        "virtual": {
          "description": "Whether the machine appears to be a virtual machine or not",
          "type": "boolean"
        }
      }
    },
    "workload": {
      "type": "object",
      "properties": {
        "data": {
          "description": "Configuration for data transfer",
          "$ref": "#/definitions/data-configuration"
        },
        "name": {
          "description": "Name of the workload",
          "type": "string"
        },
        "specification": {
          "type": "string"
        }
      }
    },
    "workload-list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/workload"
      }
    },
    "workload-status": {
      "type": "object",
      "properties": {
        "last_data_upload": {
          "type": "string",
          "format": "date-time"
        },
        "name": {
          "type": "string"
        },
        "status": {
          "type": "string",
          "enum": [
            "deploying",
            "running",
            "crashed",
            "stopped"
          ]
        }
      }
    }
  },
  "securityDefinitions": {
    "agentAuth": {
      "type": "apiKey",
      "name": "X-Secret-Key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Device management",
      "name": "devices"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Kube for Edge Management",
    "title": "Kube4EdgeManagement",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/api/k4e-management/v1",
  "paths": {
    "/control/{device_id}/in": {
      "get": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "GetControlMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    },
    "/control/{device_id}/out": {
      "post": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "PostControlMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          },
          {
            "name": "message",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    },
    "/data/{device_id}/in": {
      "get": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "GetDataMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    },
    "/data/{device_id}/out": {
      "post": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "PostDataMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          },
          {
            "name": "message",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    }
  },
  "definitions": {
    "boot": {
      "type": "object",
      "properties": {
        "current_boot_mode": {
          "type": "string"
        },
        "pxe_interface": {
          "type": "string"
        }
      }
    },
    "cpu": {
      "type": "object",
      "properties": {
        "architecture": {
          "type": "string"
        },
        "count": {
          "type": "integer"
        },
        "flags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "frequency": {
          "type": "number"
        },
        "model_name": {
          "type": "string"
        }
      }
    },
    "data-configuration": {
      "type": "object",
      "properties": {
        "paths": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/data-path"
          }
        }
      }
    },
    "data-path": {
      "description": "Device-to-control plane paths mapping",
      "type": "object",
      "properties": {
        "source": {
          "description": "Path in the workload container",
          "type": "string"
        },
        "target": {
          "description": "Path in the control plane storage",
          "type": "string"
        }
      }
    },
    "device-configuration": {
      "type": "object",
      "properties": {
        "heartbeat": {
          "$ref": "#/definitions/heartbeat-configuration"
        },
        "storage": {
          "$ref": "#/definitions/storage-configuration"
        }
      }
    },
    "device-configuration-message": {
      "type": "object",
      "properties": {
        "configuration": {
          "$ref": "#/definitions/device-configuration"
        },
        "device_id": {
          "description": "Device identifier",
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "workloads": {
          "description": "List of workloads deployed to the device",
          "$ref": "#/definitions/workload-list"
        }
      }
    },
    "disk": {
      "type": "object",
      "properties": {
        "bootable": {
          "type": "boolean"
        },
        "by_id": {
          "description": "by-id is the World Wide Number of the device which guaranteed to be unique for every storage device",
          "type": "string"
        },
        "by_path": {
          "description": "by-path is the shortest physical path to the device",
          "type": "string"
        },
        "drive_type": {
          "type": "string"
        },
        "hctl": {
          "type": "string"
        },
        "id": {
          "description": "Determine the disk's unique identifier which is the by-id field if it exists and fallback to the by-path field otherwise",
          "type": "string"
        },
        "io_perf": {
          "$ref": "#/definitions/io_perf"
        },
        "is_installation_media": {
          "description": "Whether the disk appears to be an installation media or not",
          "type": "boolean"
        },
        "model": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "path": {
          "type": "string"
        },
        "serial": {
          "type": "string"
        },
        "size_bytes": {
          "type": "integer"
        },
        "smart": {
          "type": "string"
        },
        "vendor": {
          "type": "string"
        },
        "wwn": {
          "type": "string"
        }
      }
    },
    "gpu": {
      "type": "object",
      "properties": {
        "address": {
          "description": "Device address (for example \"0000:00:02.0\")",
          "type": "string"
        },
        "device_id": {
          "description": "ID of the device (for example \"3ea0\")",
          "type": "string"
        },
        "name": {
          "description": "Product name of the device (for example \"UHD Graphics 620 (Whiskey Lake)\")",
          "type": "string"
        },
        "vendor": {
          "description": "The name of the device vendor (for example \"Intel Corporation\")",
          "type": "string"
        },
        "vendor_id": {
          "description": "ID of the vendor (for example \"8086\")",
          "type": "string"
        }
      }
    },
    "hardware-info": {
      "type": "object",
      "properties": {
        "boot": {
          "$ref": "#/definitions/boot"
        },
        "cpu": {
          "$ref": "#/definitions/cpu"
        },
        "disks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/disk"
          }
        },
        "gpus": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/gpu"
          }
        },
        "hostname": {
          "type": "string"
        },
        "interfaces": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/interface"
          }
        },
        "memory": {
          "$ref": "#/definitions/memory"
        },
        "system_vendor": {
          "$ref": "#/definitions/system_vendor"
        }
      }
    },
    "hardware-profile-configuration": {
      "type": "object",
      "properties": {
        "include": {
          "type": "boolean"
        },
        "scope": {
          "type": "string",
          "enum": [
            "full",
            "delta"
          ]
        }
      }
    },
    "heartbeat": {
      "type": "object",
      "properties": {
        "hardware": {
          "description": "Hardware information",
          "$ref": "#/definitions/hardware-info"
        },
        "status": {
          "type": "string",
          "enum": [
            "up",
            "degraded"
          ]
        },
        "time": {
          "type": "string",
          "format": "date-time"
        },
        "version": {
          "type": "string"
        },
        "workloads": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/workload-status"
          }
        }
      }
    },
    "heartbeat-configuration": {
      "type": "object",
      "properties": {
        "hardware_profile": {
          "$ref": "#/definitions/hardware-profile-configuration"
        },
        "period_seconds": {
          "type": "integer"
        }
      }
    },
    "interface": {
      "type": "object",
      "properties": {
        "biosdevname": {
          "type": "string"
        },
        "client_id": {
          "type": "string"
        },
        "flags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "has_carrier": {
          "type": "boolean"
        },
        "ipv4_addresses": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "ipv6_addresses": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "mac_address": {
          "type": "string"
        },
        "mtu": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "product": {
          "type": "string"
        },
        "speed_mbps": {
          "type": "integer"
        },
        "vendor": {
          "type": "string"
        }
      }
    },
    "io_perf": {
      "type": "object",
      "properties": {
        "sync_duration": {
          "description": "99th percentile of fsync duration in milliseconds",
          "type": "integer"
        }
      }
    },
    "memory": {
      "type": "object",
      "properties": {
        "physical_bytes": {
          "type": "integer"
        },
        "usable_bytes": {
          "type": "integer"
        }
      }
    },
    "message": {
      "type": "object",
      "properties": {
        "content": {
          "description": "Content"
        },
        "directive": {
          "type": "string"
        },
        "message_id": {
          "type": "string"
        },
        "metadata": {
          "type": "object"
        },
        "response_to": {
          "type": "string"
        },
        "sent": {
          "type": "string",
          "format": "date-time"
        },
        "type": {
          "type": "string",
          "enum": [
            "connection-status",
            "command",
            "event",
            "data"
          ]
        },
        "version": {
          "type": "integer"
        }
      }
    },
    "registration-info": {
      "type": "object",
      "properties": {
        "hardware": {
          "description": "Hardware information",
          "$ref": "#/definitions/hardware-info"
        },
        "os_image_id": {
          "type": "string"
        }
      }
    },
    "s3-storage-configuration": {
      "type": "object",
      "properties": {
        "aws_access_key_id": {
          "type": "string"
        },
        "aws_ca_bundle": {
          "type": "string"
        },
        "aws_secret_access_key": {
          "type": "string"
        },
        "bucket_host": {
          "type": "string"
        },
        "bucket_name": {
          "type": "string"
        },
        "bucket_port": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "storage-configuration": {
      "type": "object",
      "properties": {
        "s3": {
          "$ref": "#/definitions/s3-storage-configuration"
        }
      }
    },
    "system_vendor": {
      "type": "object",
      "properties": {
        "manufacturer": {
          "type": "string"
        },
        "product_name": {
          "type": "string"
        },
        "serial_number": {
          "type": "string"
        },
        "virtual": {
          "description": "Whether the machine appears to be a virtual machine or not",
          "type": "boolean"
        }
      }
    },
    "workload": {
      "type": "object",
      "properties": {
        "data": {
          "description": "Configuration for data transfer",
          "$ref": "#/definitions/data-configuration"
        },
        "name": {
          "description": "Name of the workload",
          "type": "string"
        },
        "specification": {
          "type": "string"
        }
      }
    },
    "workload-list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/workload"
      }
    },
    "workload-status": {
      "type": "object",
      "properties": {
        "last_data_upload": {
          "type": "string",
          "format": "date-time"
        },
        "name": {
          "type": "string"
        },
        "status": {
          "type": "string",
          "enum": [
            "deploying",
            "running",
            "crashed",
            "stopped"
          ]
        }
      }
    }
  },
  "securityDefinitions": {
    "agentAuth": {
      "type": "apiKey",
      "name": "X-Secret-Key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Device management",
      "name": "devices"
    }
  ]
}`))
}
