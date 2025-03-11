package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkServicesGrpcRoute = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the GrpcRoute was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A free-text description of the resource. Max length 1024 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_labels": {
        "computed": true,
        "description": "All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "gateways": {
        "description": "List of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "hostnames": {
        "description": "Required. Service hostnames with an optional port for which this route describes traffic.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Set of label tags associated with the GrpcRoute resource.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Location (region) of the GRPCRoute resource to be created. Only the value 'global' is currently allowed; defaults to 'global' if omitted.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "meshes": {
        "description": "List of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description": "Name of the GrpcRoute resource.",
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
        "description": "Server-defined URL of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "Time the GrpcRoute was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "rules": {
        "block": {
          "block_types": {
            "action": {
              "block": {
                "attributes": {
                  "timeout": {
                    "description": "Specifies the timeout for selected route.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "destinations": {
                    "block": {
                      "attributes": {
                        "service_name": {
                          "description": "The URL of a BackendService to route traffic to.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "weight": {
                          "description": "Specifies the proportion of requests forwarded to the backend referenced by the serviceName field.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "The destination to which traffic should be forwarded.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "fault_injection_policy": {
                    "block": {
                      "block_types": {
                        "abort": {
                          "block": {
                            "attributes": {
                              "http_status": {
                                "description": "The HTTP status code used to abort the request.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "percentage": {
                                "description": "The percentage of traffic which will be aborted.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "Specification of how client requests are aborted as part of fault injection before being sent to a destination.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "delay": {
                          "block": {
                            "attributes": {
                              "fixed_delay": {
                                "description": "Specify a fixed delay before forwarding the request.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "percentage": {
                                "description": "The percentage of traffic on which delay will be injected.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "Specification of how client requests are delayed as part of fault injection before being sent to a destination.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The specification for fault injection introduced into traffic to test the resiliency of clients to backend service failure.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "retry_policy": {
                    "block": {
                      "attributes": {
                        "num_retries": {
                          "description": "Specifies the allowed number of retries.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "retry_conditions": {
                          "description": "Specifies one or more conditions when this retry policy applies. Possible values: [\"connect-failure\", \"refused-stream\", \"cancelled\", \"deadline-exceeded\", \"resource-exhausted\", \"unavailable\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Specifies the retry policy associated with this route.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Required. A detailed rule defining how to route traffic.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "matches": {
              "block": {
                "block_types": {
                  "headers": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description": "Required. The key of the header.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "type": {
                          "description": "The type of match. Default value: \"EXACT\" Possible values: [\"TYPE_UNSPECIFIED\", \"EXACT\", \"REGULAR_EXPRESSION\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "Required. The value of the header.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Specifies a list of HTTP request headers to match against.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "method": {
                    "block": {
                      "attributes": {
                        "case_sensitive": {
                          "description": "Specifies that matches are case sensitive. The default value is true.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "grpc_method": {
                          "description": "Required. Name of the method to match against.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "grpc_service": {
                          "description": "Required. Name of the service to match against.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "A gRPC method to match against. If this field is empty or omitted, will match all methods.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Matches define conditions used for matching the rule against incoming gRPC requests.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Rules that define how traffic is routed and handled.",
          "description_kind": "plain"
        },
        "min_items": 1,
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
  "version": 1
}`

func GoogleNetworkServicesGrpcRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkServicesGrpcRoute), &result)
	return &result
}
