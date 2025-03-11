package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkServicesHttpRoute = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the HttpRoute was created in UTC.",
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
        "description": "Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests served by the gateway.\nEach gateway reference should match the pattern: projects/*/locations/global/gateways/\u003cgateway_name\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "hostnames": {
        "description": "Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.",
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
        "description": "Set of label tags associated with the HttpRoute resource.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "meshes": {
        "description": "Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served by the mesh.\nEach mesh reference should match the pattern: projects/*/locations/global/meshes/\u003cmesh_name\u003e.\nThe attached Mesh should be of a type SIDECAR.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description": "Name of the HttpRoute resource.",
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
        "description": "Time the HttpRoute was updated in UTC.",
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
                  "cors_policy": {
                    "block": {
                      "attributes": {
                        "allow_credentials": {
                          "description": "In response to a preflight request, setting this to true indicates that the actual request can include user credentials.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "allow_headers": {
                          "description": "Specifies the content for Access-Control-Allow-Headers header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "allow_methods": {
                          "description": "Specifies the content for Access-Control-Allow-Methods header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "allow_origin_regexes": {
                          "description": "Specifies the regular expression patterns that match allowed origins.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "allow_origins": {
                          "description": "Specifies the list of origins that will be allowed to do CORS requests.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "disabled": {
                          "description": "If true, the CORS policy is disabled. The default value is false, which indicates that the CORS policy is in effect.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "expose_headers": {
                          "description": "Specifies the content for Access-Control-Expose-Headers header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "max_age": {
                          "description": "Specifies how long result of a preflight request can be cached in seconds.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The specification for allowing client side cross-origin requests.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
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
                          "description": "Specifies the proportion of requests forwarded to the backend referenced by the serviceName field. This is computed as: weight/Sum(weights in this destination list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports.\nIf only one serviceName is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend.\nIf weights are specified for any one service name, they need to be specified for all of them.\nIf weights are unspecified for all services, then, traffic is distributed in equal proportions to all of them.",
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
                  "redirect": {
                    "block": {
                      "attributes": {
                        "host_redirect": {
                          "description": "The host that will be used in the redirect response instead of the one that was supplied in the request.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "https_redirect": {
                          "description": "If set to true, the URL scheme in the redirected request is set to https.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "path_redirect": {
                          "description": "The path that will be used in the redirect response instead of the one that was supplied in the request. pathRedirect can not be supplied together with prefixRedirect. Supply one alone or neither. If neither is supplied, the path of the original request will be used for the redirect.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "port_redirect": {
                          "description": "The port that will be used in the redirected request instead of the one that was supplied in the request.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "prefix_rewrite": {
                          "description": "Indicates that during redirection, the matched prefix (or path) should be swapped with this value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "response_code": {
                          "description": "The HTTP Status code to use for the redirect.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "strip_query": {
                          "description": "If set to true, any accompanying query portion of the original URL is removed prior to redirecting the request.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "If set, the request is directed as configured by this field.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "request_header_modifier": {
                    "block": {
                      "attributes": {
                        "add": {
                          "description": "Add the headers with given map where key is the name of the header, value is the value of the header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "remove": {
                          "description": "Remove headers (matching by header names) specified in the list.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "set": {
                          "description": "Completely overwrite/replace the headers with given map where key is the name of the header, value is the value of the header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "The specification for modifying the headers of a matching request prior to delivery of the request to the destination.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "request_mirror_policy": {
                    "block": {
                      "block_types": {
                        "destination": {
                          "block": {
                            "attributes": {
                              "service_name": {
                                "description": "The URL of a BackendService to route traffic to.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "weight": {
                                "description": "Specifies the proportion of requests forwarded to the backend referenced by the serviceName field. This is computed as: weight/Sum(weights in this destination list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports.\nIf only one serviceName is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend.\nIf weights are specified for any one service name, they need to be specified for all of them.\nIf weights are unspecified for all services, then, traffic is distributed in equal proportions to all of them.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "The destination the requests will be mirrored to.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Specifies the policy on how requests intended for the routes destination are shadowed to a separate mirrored destination.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "response_header_modifier": {
                    "block": {
                      "attributes": {
                        "add": {
                          "description": "Add the headers with given map where key is the name of the header, value is the value of the header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "remove": {
                          "description": "Remove headers (matching by header names) specified in the list.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "set": {
                          "description": "Completely overwrite/replace the headers with given map where key is the name of the header, value is the value of the header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "The specification for modifying the headers of a response prior to sending the response back to the client.",
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
                        "per_try_timeout": {
                          "description": "Specifies a non-zero timeout per retry attempt. A duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.5s\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "retry_conditions": {
                          "description": "Specifies one or more conditions when this retry policy applies.",
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
                  },
                  "url_rewrite": {
                    "block": {
                      "attributes": {
                        "host_rewrite": {
                          "description": "Prior to forwarding the request to the selected destination, the requests host header is replaced by this value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "path_prefix_rewrite": {
                          "description": "Prior to forwarding the request to the selected destination, the matching portion of the requests path is replaced by this value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The specification for rewrite URL before forwarding requests to the destination.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The detailed rule defining how to route matched traffic.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "matches": {
              "block": {
                "attributes": {
                  "full_path_match": {
                    "description": "The HTTP request path value should exactly match this value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ignore_case": {
                    "description": "Specifies if prefixMatch and fullPathMatch matches are case sensitive. The default value is false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "prefix_match": {
                    "description": "The HTTP request path value must begin with specified prefixMatch. prefixMatch must begin with a /.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "regex_match": {
                    "description": "The HTTP request path value must satisfy the regular expression specified by regexMatch after removing any query parameters and anchor supplied with the original URL. For regular expression grammar, please see https://github.com/google/re2/wiki/Syntax",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "headers": {
                    "block": {
                      "attributes": {
                        "exact_match": {
                          "description": "The value of the header should match exactly the content of exactMatch.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "header": {
                          "description": "The name of the HTTP header to match against.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "invert_match": {
                          "description": "If specified, the match result will be inverted before checking. Default value is set to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "prefix_match": {
                          "description": "The value of the header must start with the contents of prefixMatch.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "present_match": {
                          "description": "A header with headerName must exist. The match takes place whether or not the header has a value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "regex_match": {
                          "description": "The value of the header must match the regular expression specified in regexMatch.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "suffix_match": {
                          "description": "The value of the header must end with the contents of suffixMatch.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "range_match": {
                          "block": {
                            "attributes": {
                              "end": {
                                "description": "End of the range (exclusive).",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "start": {
                                "description": "Start of the range (inclusive).",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description": "If specified, the rule will match if the request header value is within the range.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Specifies a list of HTTP request headers to match against.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "query_parameters": {
                    "block": {
                      "attributes": {
                        "exact_match": {
                          "description": "The value of the query parameter must exactly match the contents of exactMatch.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "present_match": {
                          "description": "Specifies that the QueryParameterMatcher matches if request contains query parameter, irrespective of whether the parameter has a value or not.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "query_parameter": {
                          "description": "The name of the query parameter to match.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "regex_match": {
                          "description": "The value of the query parameter must match the regular expression specified by regexMatch.For regular expression grammar, please see https://github.com/google/re2/wiki/Syntax",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Specifies a list of query parameters to match against.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "A list of matches define conditions used for matching the rule against incoming HTTP requests. Each match is independent, i.e. this rule will be matched if ANY one of the matches is satisfied.\nIf no matches field is specified, this rule will unconditionally match traffic.\nIf a default rule is desired to be configured, add a rule with no matches specified to the end of the rules list.",
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
  "version": 0
}`

func GoogleNetworkServicesHttpRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkServicesHttpRoute), &result)
	return &result
}
