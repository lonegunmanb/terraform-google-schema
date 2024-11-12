package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeHealthCheck = `{
  "block": {
    "attributes": {
      "check_interval_sec": {
        "description": "How often (in seconds) to send a health check. The default value is 5\nseconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "healthy_threshold": {
        "description": "A so-far unhealthy instance will be marked healthy after this many\nconsecutive successes. The default value is 2.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035.  Specifically, the name must be 1-63 characters long and\nmatch the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means\nthe first character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the\nlast character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_regions": {
        "description": "The list of cloud regions from which health checks are performed. If\nany regions are specified, then exactly 3 regions should be specified.\nThe region names must be valid names of Google Cloud regions. This can\nonly be set for global health check. If this list is non-empty, then\nthere are restrictions on what other health check fields are supported\nand what other resources can use this health check:\n\n* SSL, HTTP2, and GRPC protocols are not supported.\n\n* The TCP request field is not supported.\n\n* The proxyHeader field for HTTP, HTTPS, and TCP is not supported.\n\n* The checkIntervalSec field must be at least 30.\n\n* The health check cannot be used with BackendService nor with managed\ninstance group auto-healing.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "timeout_sec": {
        "description": "How long (in seconds) to wait before claiming failure.\nThe default value is 5 seconds.  It is invalid for timeoutSec to have\ngreater value than checkIntervalSec.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "type": {
        "computed": true,
        "description": "The type of the health check. One of HTTP, HTTPS, TCP, or SSL.",
        "description_kind": "plain",
        "type": "string"
      },
      "unhealthy_threshold": {
        "description": "A so-far healthy instance will be marked unhealthy after this many\nconsecutive failures. The default value is 2.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "grpc_health_check": {
        "block": {
          "attributes": {
            "grpc_service_name": {
              "description": "The gRPC service name for the health check.\nThe value of grpcServiceName has the following meanings by convention:\n  - Empty serviceName means the overall status of all services at the backend.\n  - Non-empty serviceName means the health of that gRPC service, as defined by the owner of the service.\nThe grpcServiceName can only be ASCII.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description": "The port number for the health check request.\nMust be specified if portName and portSpecification are not set\nor if port_specification is USE_FIXED_PORT. Valid values are 1 through 65535.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "port_name": {
              "description": "Port name as defined in InstanceGroup#NamedPort#name. If both port and\nport_name are defined, port takes precedence.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port_specification": {
              "description": "Specifies how port is selected for health checking, can be one of the\nfollowing values:\n\n  * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.\n\n  * 'USE_NAMED_PORT': The 'portName' is used for health checking.\n\n  * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each\n  network endpoint is used for health checking. For other backends, the\n  port or named port specified in the Backend Service is used for health\n  checking.\n\nIf not specified, gRPC health check follows behavior specified in 'port' and\n'portName' fields. Possible values: [\"USE_FIXED_PORT\", \"USE_NAMED_PORT\", \"USE_SERVING_PORT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A nested object resource.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "http2_health_check": {
        "block": {
          "attributes": {
            "host": {
              "description": "The value of the host header in the HTTP2 health check request.\nIf left empty (default value), the public IP on behalf of which this health\ncheck is performed will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description": "The TCP port number for the HTTP2 health check request.\nThe default value is 443.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "port_name": {
              "description": "Port name as defined in InstanceGroup#NamedPort#name. If both port and\nport_name are defined, port takes precedence.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port_specification": {
              "description": "Specifies how port is selected for health checking, can be one of the\nfollowing values:\n\n  * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.\n\n  * 'USE_NAMED_PORT': The 'portName' is used for health checking.\n\n  * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each\n  network endpoint is used for health checking. For other backends, the\n  port or named port specified in the Backend Service is used for health\n  checking.\n\nIf not specified, HTTP2 health check follows behavior specified in 'port' and\n'portName' fields. Possible values: [\"USE_FIXED_PORT\", \"USE_NAMED_PORT\", \"USE_SERVING_PORT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "proxy_header": {
              "description": "Specifies the type of proxy header to append before sending data to the\nbackend. Default value: \"NONE\" Possible values: [\"NONE\", \"PROXY_V1\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "request_path": {
              "description": "The request path of the HTTP2 health check request.\nThe default value is /.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "response": {
              "description": "The bytes to match against the beginning of the response data. If left empty\n(the default value), any response will indicate health. The response data\ncan only be ASCII.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A nested object resource.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "http_health_check": {
        "block": {
          "attributes": {
            "host": {
              "description": "The value of the host header in the HTTP health check request.\nIf left empty (default value), the public IP on behalf of which this health\ncheck is performed will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description": "The TCP port number for the HTTP health check request.\nThe default value is 80.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "port_name": {
              "description": "Port name as defined in InstanceGroup#NamedPort#name. If both port and\nport_name are defined, port takes precedence.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port_specification": {
              "description": "Specifies how port is selected for health checking, can be one of the\nfollowing values:\n\n  * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.\n\n  * 'USE_NAMED_PORT': The 'portName' is used for health checking.\n\n  * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each\n  network endpoint is used for health checking. For other backends, the\n  port or named port specified in the Backend Service is used for health\n  checking.\n\nIf not specified, HTTP health check follows behavior specified in 'port' and\n'portName' fields. Possible values: [\"USE_FIXED_PORT\", \"USE_NAMED_PORT\", \"USE_SERVING_PORT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "proxy_header": {
              "description": "Specifies the type of proxy header to append before sending data to the\nbackend. Default value: \"NONE\" Possible values: [\"NONE\", \"PROXY_V1\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "request_path": {
              "description": "The request path of the HTTP health check request.\nThe default value is /.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "response": {
              "description": "The bytes to match against the beginning of the response data. If left empty\n(the default value), any response will indicate health. The response data\ncan only be ASCII.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A nested object resource.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "https_health_check": {
        "block": {
          "attributes": {
            "host": {
              "description": "The value of the host header in the HTTPS health check request.\nIf left empty (default value), the public IP on behalf of which this health\ncheck is performed will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description": "The TCP port number for the HTTPS health check request.\nThe default value is 443.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "port_name": {
              "description": "Port name as defined in InstanceGroup#NamedPort#name. If both port and\nport_name are defined, port takes precedence.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port_specification": {
              "description": "Specifies how port is selected for health checking, can be one of the\nfollowing values:\n\n  * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.\n\n  * 'USE_NAMED_PORT': The 'portName' is used for health checking.\n\n  * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each\n  network endpoint is used for health checking. For other backends, the\n  port or named port specified in the Backend Service is used for health\n  checking.\n\nIf not specified, HTTPS health check follows behavior specified in 'port' and\n'portName' fields. Possible values: [\"USE_FIXED_PORT\", \"USE_NAMED_PORT\", \"USE_SERVING_PORT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "proxy_header": {
              "description": "Specifies the type of proxy header to append before sending data to the\nbackend. Default value: \"NONE\" Possible values: [\"NONE\", \"PROXY_V1\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "request_path": {
              "description": "The request path of the HTTPS health check request.\nThe default value is /.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "response": {
              "description": "The bytes to match against the beginning of the response data. If left empty\n(the default value), any response will indicate health. The response data\ncan only be ASCII.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A nested object resource.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "log_config": {
        "block": {
          "attributes": {
            "enable": {
              "description": "Indicates whether or not to export logs. This is false by default,\nwhich means no health check logging will be done.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Configure logging on this health check.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ssl_health_check": {
        "block": {
          "attributes": {
            "port": {
              "description": "The TCP port number for the SSL health check request.\nThe default value is 443.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "port_name": {
              "description": "Port name as defined in InstanceGroup#NamedPort#name. If both port and\nport_name are defined, port takes precedence.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port_specification": {
              "description": "Specifies how port is selected for health checking, can be one of the\nfollowing values:\n\n  * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.\n\n  * 'USE_NAMED_PORT': The 'portName' is used for health checking.\n\n  * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each\n  network endpoint is used for health checking. For other backends, the\n  port or named port specified in the Backend Service is used for health\n  checking.\n\nIf not specified, SSL health check follows behavior specified in 'port' and\n'portName' fields. Possible values: [\"USE_FIXED_PORT\", \"USE_NAMED_PORT\", \"USE_SERVING_PORT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "proxy_header": {
              "description": "Specifies the type of proxy header to append before sending data to the\nbackend. Default value: \"NONE\" Possible values: [\"NONE\", \"PROXY_V1\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "request": {
              "description": "The application data to send once the SSL connection has been\nestablished (default value is empty). If both request and response are\nempty, the connection establishment alone will indicate health. The request\ndata can only be ASCII.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "response": {
              "description": "The bytes to match against the beginning of the response data. If left empty\n(the default value), any response will indicate health. The response data\ncan only be ASCII.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A nested object resource.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "tcp_health_check": {
        "block": {
          "attributes": {
            "port": {
              "description": "The TCP port number for the TCP health check request.\nThe default value is 443.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "port_name": {
              "description": "Port name as defined in InstanceGroup#NamedPort#name. If both port and\nport_name are defined, port takes precedence.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port_specification": {
              "description": "Specifies how port is selected for health checking, can be one of the\nfollowing values:\n\n  * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.\n\n  * 'USE_NAMED_PORT': The 'portName' is used for health checking.\n\n  * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each\n  network endpoint is used for health checking. For other backends, the\n  port or named port specified in the Backend Service is used for health\n  checking.\n\nIf not specified, TCP health check follows behavior specified in 'port' and\n'portName' fields. Possible values: [\"USE_FIXED_PORT\", \"USE_NAMED_PORT\", \"USE_SERVING_PORT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "proxy_header": {
              "description": "Specifies the type of proxy header to append before sending data to the\nbackend. Default value: \"NONE\" Possible values: [\"NONE\", \"PROXY_V1\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "request": {
              "description": "The application data to send once the TCP connection has been\nestablished (default value is empty). If both request and response are\nempty, the connection establishment alone will indicate health. The request\ndata can only be ASCII.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "response": {
              "description": "The bytes to match against the beginning of the response data. If left empty\n(the default value), any response will indicate health. The response data\ncan only be ASCII.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A nested object resource.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeHealthCheckSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeHealthCheck), &result)
	return &result
}
