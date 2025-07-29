package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeApiProduct = `{
  "block": {
    "attributes": {
      "api_resources": {
        "description": "Comma-separated list of API resources to be bundled in the API product. By default, the resource paths are mapped from the proxy.pathsuffix variable.\nThe proxy path suffix is defined as the URI fragment following the ProxyEndpoint base path. For example, if the apiResources element is defined to be /forecastrss and the base path defined for the API proxy is /weather, then only requests to /weather/forecastrss are permitted by the API product.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "approval_type": {
        "description": "Flag that specifies how API keys are approved to access the APIs defined by the API product.\nValid values are 'auto' or 'manual'. Possible values: [\"auto\", \"manual\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Response only. Creation time of this environment as milliseconds since epoch.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the API product. Include key information about the API product that is not captured by other fields.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Name displayed in the UI or developer portal to developers registering for API access.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "environments": {
        "description": "Comma-separated list of environment names to which the API product is bound. Requests to environments that are not listed are rejected.\nBy specifying one or more environments, you can bind the resources listed in the API product to a specific environment, preventing developers from accessing those resources through API proxies deployed in another environment.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified_at": {
        "computed": true,
        "description": "Response only. Modified time of this environment as milliseconds since epoch.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Internal name of the API product.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization associated with the Apigee API product,\nin the format 'organizations/{{org_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "proxies": {
        "description": "Comma-separated list of API proxy names to which this API product is bound. By specifying API proxies, you can associate resources in the API product with specific API proxies, preventing developers from accessing those resources through other API proxies.\nApigee rejects requests to API proxies that are not listed.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "quota": {
        "description": "Number of request messages permitted per app by this API product for the specified quotaInterval and quotaTimeUnit.\nFor example, a quota of 50, for a quotaInterval of 12 and a quotaTimeUnit of hours means 50 requests are allowed every 12 hours.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "quota_counter_scope": {
        "description": "Scope of the quota decides how the quota counter gets applied and evaluate for quota violation. If the Scope is set as PROXY, then all the operations defined for the APIproduct that are associated with the same proxy will share the same quota counter set at the APIproduct level, making it a global counter at a proxy level. If the Scope is set as OPERATION, then each operations get the counter set at the API product dedicated, making it a local counter. Note that, the QuotaCounterScope applies only when an operation does not have dedicated quota set for itself. Possible values: [\"QUOTA_COUNTER_SCOPE_UNSPECIFIED\", \"PROXY\", \"OPERATION\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "quota_interval": {
        "description": "Time interval over which the number of request messages is calculated.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "quota_time_unit": {
        "description": "Time unit defined for the quotaInterval. Valid values include second, minute, hour, day, month or year.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scopes": {
        "description": "Comma-separated list of OAuth scopes that are validated at runtime. Apigee validates that the scopes in any access token presented match the scopes defined in the OAuth policy associated with the API product.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "space": {
        "description": "Optional. The resource ID of the parent Space. If not set, the parent resource will be the Organization.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "attributes": {
        "block": {
          "attributes": {
            "name": {
              "description": "Key of the attribute.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description": "Value of the attribute.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Array of attributes that may be used to extend the default API product profile with customer-specific metadata. You can specify a maximum of 18 attributes.\nUse this property to specify the access level of the API product as either public, private, or internal.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "graphql_operation_group": {
        "block": {
          "attributes": {
            "operation_config_type": {
              "description": "Flag that specifes whether the configuration is for Apigee API proxy or a remote service. Valid values include proxy or remoteservice. Defaults to proxy. Set to proxy when Apigee API proxies are associated with the API product. Set to remoteservice when non-Apigee proxies like Istio-Envoy are associated with the API product. Possible values: [\"proxy\", \"remoteservice\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "operation_configs": {
              "block": {
                "attributes": {
                  "api_source": {
                    "description": "Required. Name of the API proxy endpoint or remote service with which the GraphQL operation and quota are associated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "attributes": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Key of the attribute.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "Value of the attribute.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Custom attributes associated with the operation.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "operations": {
                    "block": {
                      "attributes": {
                        "operation": {
                          "description": "GraphQL operation name. The name and operation type will be used to apply quotas. If no name is specified, the quota will be applied to all GraphQL operations irrespective of their operation names in the payload.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "operation_types": {
                          "description": "Required. GraphQL operation types. Valid values include query or mutation.\nNote: Apigee does not currently support subscription types.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "description": "Required. List of GraphQL name/operation type pairs for the proxy or remote service to which quota will be applied. If only operation types are specified, the quota will be applied to all GraphQL requests irrespective of the GraphQL name.\n\nNote: Currently, you can specify only a single GraphQLOperation. Specifying more than one will cause the operation to fail.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "quota": {
                    "block": {
                      "attributes": {
                        "interval": {
                          "description": "Required. Time interval over which the number of request messages is calculated.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "limit": {
                          "description": "Required. Upper limit allowed for the time interval and time unit specified. Requests exceeding this limit will be rejected.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "time_unit": {
                          "description": "Time unit defined for the interval. Valid values include second, minute, hour, day, month or year. If limit and interval are valid, the default value is hour; otherwise, the default is null.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Quota parameters to be enforced for the resources, methods, and API source combination. If none are specified, quota enforcement will not be done.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of graphQL operation configuration details associated with Apigee API proxies or remote services. Remote services are non-Apigee proxies, such as Istio-Envoy.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "Configuration used to group Apigee proxies or remote services with graphQL operation name, graphQL operation type and quotas. This grouping allows us to precisely set quota for a particular combination of graphQL name and operation type for a particular proxy request. If graphQL name is not set, this would imply quota will be applied on all graphQL requests matching the operation type.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "grpc_operation_group": {
        "block": {
          "block_types": {
            "operation_configs": {
              "block": {
                "attributes": {
                  "api_source": {
                    "description": "Required. Name of the API proxy with which the gRPC operation and quota are associated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "methods": {
                    "description": "List of unqualified gRPC method names for the proxy to which quota will be applied. If this field is empty, the Quota will apply to all operations on the gRPC service defined on the proxy.\n\nExample: Given a proxy that is configured to serve com.petstore.PetService, the methods com.petstore.PetService.ListPets and com.petstore.PetService.GetPet would be specified here as simply [\"ListPets\", \"GetPet\"].\n\nNote: Currently, you can specify only a single GraphQLOperation. Specifying more than one will cause the operation to fail.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "service": {
                    "description": "Required. gRPC Service name associated to be associated with the API proxy, on which quota rules can be applied upon.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "attributes": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Key of the attribute.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "Value of the attribute.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Custom attributes associated with the operation.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "quota": {
                    "block": {
                      "attributes": {
                        "interval": {
                          "description": "Required. Time interval over which the number of request messages is calculated.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "limit": {
                          "description": "Required. Upper limit allowed for the time interval and time unit specified. Requests exceeding this limit will be rejected.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "time_unit": {
                          "description": "Time unit defined for the interval. Valid values include second, minute, hour, day, month or year. If limit and interval are valid, the default value is hour; otherwise, the default is null.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Quota parameters to be enforced for the resources, methods, and API source combination. If none are specified, quota enforcement will not be done.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Required. List of operation configurations for either Apigee API proxies that are associated with this API product.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "Optional. Configuration used to group Apigee proxies with gRPC services and method names. This grouping allows us to set quota for a particular proxy with the gRPC service name and method. If a method name is not set, this implies quota and authorization are applied to all gRPC methods implemented by that proxy for that particular gRPC service.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "operation_group": {
        "block": {
          "attributes": {
            "operation_config_type": {
              "description": "Flag that specifes whether the configuration is for Apigee API proxy or a remote service. Valid values include proxy or remoteservice. Defaults to proxy. Set to proxy when Apigee API proxies are associated with the API product. Set to remoteservice when non-Apigee proxies like Istio-Envoy are associated with the API product. Possible values: [\"proxy\", \"remoteservice\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "operation_configs": {
              "block": {
                "attributes": {
                  "api_source": {
                    "description": "Required. Name of the API proxy or remote service with which the resources, methods, and quota are associated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "attributes": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Key of the attribute.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "Value of the attribute.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Custom attributes associated with the operation.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "operations": {
                    "block": {
                      "attributes": {
                        "methods": {
                          "description": "Methods refers to the REST verbs, when none specified, all verb types are allowed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "resource": {
                          "description": "Required. REST resource path associated with the API proxy or remote service.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "List of resource/method pairs for the API proxy or remote service to which quota will applied.\nNote: Currently, you can specify only a single resource/method pair. The call will fail if more than one resource/method pair is provided.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "quota": {
                    "block": {
                      "attributes": {
                        "interval": {
                          "description": "Required. Time interval over which the number of request messages is calculated.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "limit": {
                          "description": "Required. Upper limit allowed for the time interval and time unit specified. Requests exceeding this limit will be rejected.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "time_unit": {
                          "description": "Time unit defined for the interval. Valid values include second, minute, hour, day, month or year. If limit and interval are valid, the default value is hour; otherwise, the default is null.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Quota parameters to be enforced for the resources, methods, and API source combination. If none are specified, quota enforcement will not be done.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Required. List of operation configurations for either Apigee API proxies or other remote services that are associated with this API product.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "Configuration used to group Apigee proxies or remote services with resources, method types, and quotas. The resource refers to the resource URI (excluding the base path). With this grouping, the API product creator is able to fine-tune and give precise control over which REST methods have access to specific resources and how many calls can be made (using the quota setting).\nNote: The apiResources setting cannot be specified for both the API product and operation group; otherwise the call will fail.",
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

func GoogleApigeeApiProductSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeApiProduct), &result)
	return &result
}
