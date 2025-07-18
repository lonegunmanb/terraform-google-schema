package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionUrlMap = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_service": {
        "description": "The full or partial URL of the defaultService resource to which traffic is directed if\nnone of the hostRules match. If defaultRouteAction is additionally specified, advanced\nrouting actions like URL Rewrites, etc. take effect prior to sending the request to the\nbackend. However, if defaultService is specified, defaultRouteAction cannot contain any\nweightedBackendServices. Conversely, if routeAction specifies any\nweightedBackendServices, service must not be specified.  Only one of defaultService,\ndefaultUrlRedirect or defaultRouteAction.weightedBackendService must be set.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource. This field is used internally during\nupdates of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "map_id": {
        "computed": true,
        "description": "The unique identifier for the resource.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
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
      "region": {
        "computed": true,
        "description": "The Region in which the url map should reside.\nIf it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "default_route_action": {
        "block": {
          "block_types": {
            "cors_policy": {
              "block": {
                "attributes": {
                  "allow_credentials": {
                    "description": "In response to a preflight request, setting this to true indicates that the actual request can include user credentials. This field translates to the Access-Control-Allow-Credentials header.\nDefault is false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "allow_headers": {
                    "description": "Specifies the content for the Access-Control-Allow-Headers header.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allow_methods": {
                    "description": "Specifies the content for the Access-Control-Allow-Methods header.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allow_origin_regexes": {
                    "description": "Specifies the regualar expression patterns that match allowed origins. For regular expression grammar\nplease see en.cppreference.com/w/cpp/regex/ecmascript\nAn origin is allowed if it matches either an item in allowOrigins or an item in allowOriginRegexes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allow_origins": {
                    "description": "Specifies the list of origins that will be allowed to do CORS requests.\nAn origin is allowed if it matches either an item in allowOrigins or an item in allowOriginRegexes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "disabled": {
                    "description": "If true, the setting specifies the CORS policy is disabled. The default value of false, which indicates that the CORS policy is in effect.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "expose_headers": {
                    "description": "Specifies the content for the Access-Control-Expose-Headers header.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "max_age": {
                    "description": "Specifies how long results of a preflight request can be cached in seconds.\nThis translates to the Access-Control-Max-Age header.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "The specification for allowing client side cross-origin requests. Please see\n[W3C Recommendation for Cross Origin Resource Sharing](https://www.w3.org/TR/cors/)",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "fault_injection_policy": {
              "block": {
                "block_types": {
                  "abort": {
                    "block": {
                      "attributes": {
                        "http_status": {
                          "description": "The HTTP status code used to abort the request.\nThe value must be between 200 and 599 inclusive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "percentage": {
                          "description": "The percentage of traffic (connections/operations/requests) which will be aborted as part of fault injection.\nThe value must be between 0.0 and 100.0 inclusive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "The specification for how client requests are aborted as part of fault injection.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "delay": {
                    "block": {
                      "attributes": {
                        "percentage": {
                          "description": "The percentage of traffic (connections/operations/requests) on which delay will be introduced as part of fault injection.\nThe value must be between 0.0 and 100.0 inclusive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "fixed_delay": {
                          "block": {
                            "attributes": {
                              "nanos": {
                                "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations less than one second are\nrepresented with a 0 seconds field and a positive nanos field. Must be from 0 to 999,999,999 inclusive.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "seconds": {
                                "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive.\nNote: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Specifies the value of the fixed delay interval.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The specification for how client requests are delayed as part of fault injection, before being sent to a backend service.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The specification for fault injection introduced into traffic to test the resiliency of clients to backend service failure.\nAs part of fault injection, when clients send requests to a backend service, delays can be introduced by a load balancer on a percentage of requests before sending those requests to the backend service.\nSimilarly requests from clients can be aborted by the load balancer for a percentage of requests.\ntimeout and retryPolicy is ignored by clients that are configured with a faultInjectionPolicy if: 1. The traffic is generated by fault injection AND 2. The fault injection is not a delay fault injection.\nFault injection is not supported with the global external HTTP(S) load balancer (classic). To see which load balancers support fault injection, see Load balancing: [Routing and traffic management features](https://cloud.google.com/load-balancing/docs/features#routing-traffic-management).",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "request_mirror_policy": {
              "block": {
                "attributes": {
                  "backend_service": {
                    "description": "The full or partial URL to the RegionBackendService resource being mirrored to.\nThe backend service configured for a mirroring policy must reference backends that are of the same type as the original backend service matched in the URL map.\nServerless NEG backends are not currently supported as a mirrored backend service.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Specifies the policy on how requests intended for the route's backends are shadowed to a separate mirrored backend service.\nThe load balancer does not wait for responses from the shadow service. Before sending traffic to the shadow service, the host / authority header is suffixed with -shadow.\nNot supported when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "retry_policy": {
              "block": {
                "attributes": {
                  "num_retries": {
                    "description": "Specifies the allowed number retries. This number must be \u003e 0. If not specified, defaults to 1.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "retry_conditions": {
                    "description": "Specifies one or more conditions when this retry policy applies.\nValid values are listed below. Only the following codes are supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true: cancelled, deadline-exceeded, internal, resource-exhausted, unavailable.\n  - 5xx : retry is attempted if the instance or endpoint responds with any 5xx response code, or if the instance or endpoint does not respond at all. For example, disconnects, reset, read timeout, connection failure, and refused streams.\n  - gateway-error : Similar to 5xx, but only applies to response codes 502, 503 or 504.\n  - connect-failure : a retry is attempted on failures connecting to the instance or endpoint. For example, connection timeouts.\n  - retriable-4xx : a retry is attempted if the instance or endpoint responds with a 4xx response code. The only error that you can retry is error code 409.\n  - refused-stream : a retry is attempted if the instance or endpoint resets the stream with a REFUSED_STREAM error code. This reset type indicates that it is safe to retry.\n  - cancelled : a retry is attempted if the gRPC status code in the response header is set to cancelled.\n  - deadline-exceeded : a retry is attempted if the gRPC status code in the response header is set to deadline-exceeded.\n  - internal :  a retry is attempted if the gRPC status code in the response header is set to internal.\n  - resource-exhausted : a retry is attempted if the gRPC status code in the response header is set to resource-exhausted.\n  - unavailable : a retry is attempted if the gRPC status code in the response header is set to unavailable.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "per_try_timeout": {
                    "block": {
                      "attributes": {
                        "nanos": {
                          "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations less than one second are\nrepresented with a 0 seconds field and a positive nanos field. Must be from 0 to 999,999,999 inclusive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "seconds": {
                          "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive.\nNote: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Specifies a non-zero timeout per retry attempt.\n\nIf not specified, will use the timeout set in HttpRouteAction. If timeout in HttpRouteAction is not set,\nwill use the largest timeout among all backend services associated with the route.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specifies the retry policy associated with this route.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "timeout": {
              "block": {
                "attributes": {
                  "nanos": {
                    "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations less than one second are represented with a 0 seconds field and a positive nanos field. Must be from 0 to 999,999,999 inclusive.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "seconds": {
                    "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive. Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Specifies the timeout for the selected route. Timeout is computed from the time the request has been fully processed (known as end-of-stream) up until the response has been processed. Timeout includes all retries.\nIf not specified, this field uses the largest timeout among all backend services associated with the route.\nNot supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "url_rewrite": {
              "block": {
                "attributes": {
                  "host_rewrite": {
                    "description": "Before forwarding the request to the selected service, the request's host header is replaced with contents of hostRewrite.\nThe value must be from 1 to 255 characters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path_prefix_rewrite": {
                    "description": "Before forwarding the request to the selected backend service, the matching portion of the request's path is replaced by pathPrefixRewrite.\nThe value must be from 1 to 1024 characters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The spec to modify the URL of the request, before forwarding the request to the matched service.\nurlRewrite is the only action supported in UrlMaps for external HTTP(S) load balancers.\nNot supported when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "weighted_backend_services": {
              "block": {
                "attributes": {
                  "backend_service": {
                    "description": "The full or partial URL to the default BackendService resource. Before forwarding the request to backendService, the load balancer applies any relevant headerActions specified as part of this backendServiceWeight.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "weight": {
                    "description": "Specifies the fraction of traffic sent to a backend service, computed as weight / (sum of all weightedBackendService weights in routeAction) .\nThe selection of a backend service is determined only for new traffic. Once a user's request has been directed to a backend service, subsequent requests are sent to the same backend service as determined by the backend service's session affinity policy.\nThe value must be from 0 to 1000.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "header_action": {
                    "block": {
                      "attributes": {
                        "request_headers_to_remove": {
                          "description": "A list of header names for headers that need to be removed from the request before forwarding the request to the backendService.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "response_headers_to_remove": {
                          "description": "A list of header names for headers that need to be removed from the response before sending the response back to the client.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "request_headers_to_add": {
                          "block": {
                            "attributes": {
                              "header_name": {
                                "description": "The name of the header.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "header_value": {
                                "description": "The value of the header to add.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "replace": {
                                "description": "If false, headerValue is appended to any values that already exist for the header. If true, headerValue is set for the header, discarding any values that were set for that header.\nThe default value is false.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "Headers to add to a matching request before forwarding the request to the backendService.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "response_headers_to_add": {
                          "block": {
                            "attributes": {
                              "header_name": {
                                "description": "The name of the header.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "header_value": {
                                "description": "The value of the header to add.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "replace": {
                                "description": "If false, headerValue is appended to any values that already exist for the header. If true, headerValue is set for the header, discarding any values that were set for that header.\nThe default value is false.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "Headers to add the response before sending the response back to the client.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Specifies changes to request and response headers that need to take effect for the selected backendService.\nheaderAction specified here take effect before headerAction in the enclosing HttpRouteRule, PathMatcher and UrlMap.\nheaderAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL.\nNot supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A list of weighted backend services to send traffic to when a route match occurs. The weights determine the fraction of traffic that flows to their corresponding backend service. If all traffic needs to go to a single backend service, there must be one weightedBackendService with weight set to a non-zero number.\nAfter a backend service is identified and before forwarding the request to the backend service, advanced routing actions such as URL rewrites and header transformations are applied depending on additional settings specified in this HttpRouteAction.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.\nOnly one of defaultRouteAction or defaultUrlRedirect must be set.\nURL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction.\ndefaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "default_url_redirect": {
        "block": {
          "attributes": {
            "host_redirect": {
              "description": "The host that will be used in the redirect response instead of the one that was\nsupplied in the request. The value must be between 1 and 255 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "https_redirect": {
              "description": "If set to true, the URL scheme in the redirected request is set to https. If set to\nfalse, the URL scheme of the redirected request will remain the same as that of the\nrequest. This must only be set for UrlMaps used in TargetHttpProxys. Setting this\ntrue for TargetHttpsProxy is not permitted. The default is set to false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "path_redirect": {
              "description": "The path that will be used in the redirect response instead of the one that was\nsupplied in the request. pathRedirect cannot be supplied together with\nprefixRedirect. Supply one alone or neither. If neither is supplied, the path of the\noriginal request will be used for the redirect. The value must be between 1 and 1024\ncharacters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "prefix_redirect": {
              "description": "The prefix that replaces the prefixMatch specified in the HttpRouteRuleMatch,\nretaining the remaining portion of the URL before redirecting the request.\nprefixRedirect cannot be supplied together with pathRedirect. Supply one alone or\nneither. If neither is supplied, the path of the original request will be used for\nthe redirect. The value must be between 1 and 1024 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "redirect_response_code": {
              "description": "The HTTP Status code to use for this RedirectAction. Supported values are:\n\n* MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301.\n\n* FOUND, which corresponds to 302.\n\n* SEE_OTHER which corresponds to 303.\n\n* TEMPORARY_REDIRECT, which corresponds to 307. In this case, the request method\nwill be retained.\n\n* PERMANENT_REDIRECT, which corresponds to 308. In this case,\nthe request method will be retained. Possible values: [\"FOUND\", \"MOVED_PERMANENTLY_DEFAULT\", \"PERMANENT_REDIRECT\", \"SEE_OTHER\", \"TEMPORARY_REDIRECT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "strip_query": {
              "description": "If set to true, any accompanying query portion of the original URL is removed prior\nto redirecting the request. If set to false, the query portion of the original URL is\nretained.\n This field is required to ensure an empty block is not set. The normal default value is false.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "When none of the specified hostRules match, the request is redirected to a URL specified\nby defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or\ndefaultRouteAction must not be set.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "host_rule": {
        "block": {
          "attributes": {
            "description": {
              "description": "An optional description of this HostRule. Provide this property\nwhen you create the resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hosts": {
              "description": "The list of host patterns to match. They must be valid\nhostnames, except * will match any string of ([a-z0-9-.]*). In\nthat case, * must be the first character and must be followed in\nthe pattern by either - or ..",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "path_matcher": {
              "description": "The name of the PathMatcher to use to match the path portion of\nthe URL if the hostRule matches the URL's host portion.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The list of HostRules to use against the URL.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "path_matcher": {
        "block": {
          "attributes": {
            "default_service": {
              "description": "A reference to a RegionBackendService resource. This will be used if\nnone of the pathRules defined by this PathMatcher is matched by\nthe URL's path portion.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "description": "An optional description of this resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description": "The name to which this PathMatcher is referred by the HostRule.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "default_route_action": {
              "block": {
                "block_types": {
                  "cors_policy": {
                    "block": {
                      "attributes": {
                        "allow_credentials": {
                          "description": "In response to a preflight request, setting this to true indicates that the actual request can include user credentials.\nThis translates to the Access-Control-Allow-Credentials header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "allow_headers": {
                          "description": "Specifies the content for the Access-Control-Allow-Headers header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "allow_methods": {
                          "description": "Specifies the content for the Access-Control-Allow-Methods header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "allow_origin_regexes": {
                          "description": "Specifies the regular expression patterns that match allowed origins. For regular expression grammar\nplease see en.cppreference.com/w/cpp/regex/ecmascript\nAn origin is allowed if it matches either an item in allowOrigins or an item in allowOriginRegexes.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "allow_origins": {
                          "description": "Specifies the list of origins that will be allowed to do CORS requests.\nAn origin is allowed if it matches either an item in allowOrigins or an item in allowOriginRegexes.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "disabled": {
                          "description": "If true, specifies the CORS policy is disabled. The default value is false, which indicates that the CORS policy is in effect.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "expose_headers": {
                          "description": "Specifies the content for the Access-Control-Expose-Headers header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "max_age": {
                          "description": "Specifies how long results of a preflight request can be cached in seconds.\nThis translates to the Access-Control-Max-Age header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "The specification for allowing client side cross-origin requests. Please see\n[W3C Recommendation for Cross Origin Resource Sharing](https://www.w3.org/TR/cors/)",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "fault_injection_policy": {
                    "block": {
                      "block_types": {
                        "abort": {
                          "block": {
                            "attributes": {
                              "http_status": {
                                "description": "The HTTP status code used to abort the request.\nThe value must be between 200 and 599 inclusive.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "percentage": {
                                "description": "The percentage of traffic (connections/operations/requests) which will be aborted as part of fault injection.\nThe value must be between 0.0 and 100.0 inclusive.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "The specification for how client requests are aborted as part of fault injection.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "delay": {
                          "block": {
                            "attributes": {
                              "percentage": {
                                "description": "The percentage of traffic (connections/operations/requests) on which delay will be introduced as part of fault injection.\nThe value must be between 0.0 and 100.0 inclusive.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "fixed_delay": {
                                "block": {
                                  "attributes": {
                                    "nanos": {
                                      "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations less than one second are\nrepresented with a 0 seconds field and a positive nanos field. Must be from 0 to 999,999,999 inclusive.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "seconds": {
                                      "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive.\nNote: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Specifies the value of the fixed delay interval.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "The specification for how client requests are delayed as part of fault injection, before being sent to a backend service.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The specification for fault injection introduced into traffic to test the resiliency of clients to backend service failure.\nAs part of fault injection, when clients send requests to a backend service, delays can be introduced by Loadbalancer on a\npercentage of requests before sending those request to the backend service. Similarly requests from clients can be aborted\nby the Loadbalancer for a percentage of requests.\n\ntimeout and retryPolicy will be ignored by clients that are configured with a faultInjectionPolicy.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "max_stream_duration": {
                    "block": {
                      "attributes": {
                        "nanos": {
                          "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations less than one second are represented\nwith a 0 seconds field and a positive nanos field. Must be from 0 to 999,999,999 inclusive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "seconds": {
                          "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive.\nNote: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Specifies the maximum duration (timeout) for streams on the selected route.\nUnlike the 'Timeout' field where the timeout duration starts from the time the request\nhas been fully processed (known as end-of-stream), the duration in this field\nis computed from the beginning of the stream until the response has been processed,\nincluding all retries. A stream that does not complete in this duration is closed.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "request_mirror_policy": {
                    "block": {
                      "attributes": {
                        "backend_service": {
                          "description": "The full or partial URL to the BackendService resource being mirrored to.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Specifies the policy on how requests intended for the route's backends are shadowed to a separate mirrored backend service.\nLoadbalancer does not wait for responses from the shadow service. Prior to sending traffic to the shadow service,\nthe host / authority header is suffixed with -shadow.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "retry_policy": {
                    "block": {
                      "attributes": {
                        "num_retries": {
                          "description": "Specifies the allowed number retries. This number must be \u003e 0. If not specified, defaults to 1.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "retry_conditions": {
                          "description": "Specfies one or more conditions when this retry rule applies. Valid values are:\n\n* 5xx: Loadbalancer will attempt a retry if the backend service responds with any 5xx response code,\n  or if the backend service does not respond at all, example: disconnects, reset, read timeout,\n* connection failure, and refused streams.\n* gateway-error: Similar to 5xx, but only applies to response codes 502, 503 or 504.\n* connect-failure: Loadbalancer will retry on failures connecting to backend services,\n  for example due to connection timeouts.\n* retriable-4xx: Loadbalancer will retry for retriable 4xx response codes.\n  Currently the only retriable error supported is 409.\n* refused-stream:Loadbalancer will retry if the backend service resets the stream with a REFUSED_STREAM error code.\n  This reset type indicates that it is safe to retry.\n* cancelled: Loadbalancer will retry if the gRPC status code in the response header is set to cancelled\n* deadline-exceeded: Loadbalancer will retry if the gRPC status code in the response header is set to deadline-exceeded\n* resource-exhausted: Loadbalancer will retry if the gRPC status code in the response header is set to resource-exhausted\n* unavailable: Loadbalancer will retry if the gRPC status code in the response header is set to unavailable",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "per_try_timeout": {
                          "block": {
                            "attributes": {
                              "nanos": {
                                "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations less than one second are\nrepresented with a 0 seconds field and a positive nanos field. Must be from 0 to 999,999,999 inclusive.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "seconds": {
                                "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive.\nNote: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Specifies a non-zero timeout per retry attempt.\n\nIf not specified, will use the timeout set in HttpRouteAction. If timeout in HttpRouteAction is not set,\nwill use the largest timeout among all backend services associated with the route.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Specifies the retry policy associated with this route.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "timeout": {
                    "block": {
                      "attributes": {
                        "nanos": {
                          "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations less than one second are represented\nwith a 0 seconds field and a positive nanos field. Must be from 0 to 999,999,999 inclusive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "seconds": {
                          "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive.\nNote: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Specifies the timeout for the selected route. Timeout is computed from the time the request has been\nfully processed (i.e. end-of-stream) up until the response has been completely processed. Timeout includes all retries.\n\nIf not specified, will use the largest timeout among all backend services associated with the route.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "url_rewrite": {
                    "block": {
                      "attributes": {
                        "host_rewrite": {
                          "description": "Prior to forwarding the request to the selected service, the request's host header is replaced\nwith contents of hostRewrite.\n\nThe value must be between 1 and 255 characters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "path_prefix_rewrite": {
                          "description": "Prior to forwarding the request to the selected backend service, the matching portion of the\nrequest's path is replaced by pathPrefixRewrite.\n\nThe value must be between 1 and 1024 characters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "path_template_rewrite": {
                          "description": "If specified, the pattern rewrites the URL path (based on the :path header) using the HTTP template syntax.\n\nA corresponding pathTemplateMatch must be specified. Any template variables must exist in the pathTemplateMatch field.\n\n* At least one variable must be specified in the pathTemplateMatch field\n* You can omit variables from the rewritten URL\n* The * and ** operators cannot be matched unless they have a corresponding variable name - e.g. {format=*} or {var=**}.\n\nFor example, a pathTemplateMatch of /static/{format=**} could be rewritten as /static/content/{format} to prefix\n/content to the URL. Variables can also be re-ordered in a rewrite, so that /{country}/{format}/{suffix=**} can be\nrewritten as /content/{format}/{country}/{suffix}.\n\nAt least one non-empty routeRules[].matchRules[].path_template_match is required.\n\nOnly one of pathPrefixRewrite or pathTemplateRewrite may be specified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The spec to modify the URL of the request, prior to forwarding the request to the matched service.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "weighted_backend_services": {
                    "block": {
                      "attributes": {
                        "backend_service": {
                          "description": "The full or partial URL to the default BackendService resource. Before forwarding the\nrequest to backendService, the loadbalancer applies any relevant headerActions\nspecified as part of this backendServiceWeight.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "weight": {
                          "description": "Specifies the fraction of traffic sent to backendService, computed as\nweight / (sum of all weightedBackendService weights in routeAction) .\n\nThe selection of a backend service is determined only for new traffic. Once a user's request\nhas been directed to a backendService, subsequent requests will be sent to the same backendService\nas determined by the BackendService's session affinity policy.\n\nThe value must be between 0 and 1000",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "header_action": {
                          "block": {
                            "attributes": {
                              "request_headers_to_remove": {
                                "description": "A list of header names for headers that need to be removed from the request prior to\nforwarding the request to the backendService.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "response_headers_to_remove": {
                                "description": "A list of header names for headers that need to be removed from the response prior to sending the\nresponse back to the client.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "request_headers_to_add": {
                                "block": {
                                  "attributes": {
                                    "header_name": {
                                      "description": "The name of the header to add.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "header_value": {
                                      "description": "The value of the header to add.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "replace": {
                                      "description": "If false, headerValue is appended to any values that already exist for the header.\nIf true, headerValue is set for the header, discarding any values that were set for that header.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description": "Headers to add to a matching request prior to forwarding the request to the backendService.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "response_headers_to_add": {
                                "block": {
                                  "attributes": {
                                    "header_name": {
                                      "description": "The name of the header to add.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "header_value": {
                                      "description": "The value of the header to add.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "replace": {
                                      "description": "If false, headerValue is appended to any values that already exist for the header.\nIf true, headerValue is set for the header, discarding any values that were set for that header.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description": "Headers to add the response prior to sending the response back to the client.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Specifies changes to request and response headers that need to take effect for\nthe selected backendService.\n\nheaderAction specified here take effect before headerAction in the enclosing\nHttpRouteRule, PathMatcher and UrlMap.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "A list of weighted backend services to send traffic to when a route match occurs.\nThe weights determine the fraction of traffic that flows to their corresponding backend service.\nIf all traffic needs to go to a single backend service, there must be one weightedBackendService\nwith weight set to a non-zero number.\n\nOnce a backendService is identified and before forwarding the request to the backend service,\nadvanced routing actions like Url rewrites and header transformations are applied depending on\nadditional settings specified in this HttpRouteAction.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs\nadvanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request\nto the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.\nConversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.\n\nOnly one of defaultRouteAction or defaultUrlRedirect must be set.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "default_url_redirect": {
              "block": {
                "attributes": {
                  "host_redirect": {
                    "description": "The host that will be used in the redirect response instead of the one that was\nsupplied in the request. The value must be between 1 and 255 characters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "https_redirect": {
                    "description": "If set to true, the URL scheme in the redirected request is set to https. If set to\nfalse, the URL scheme of the redirected request will remain the same as that of the\nrequest. This must only be set for UrlMaps used in TargetHttpProxys. Setting this\ntrue for TargetHttpsProxy is not permitted. The default is set to false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "path_redirect": {
                    "description": "The path that will be used in the redirect response instead of the one that was\nsupplied in the request. pathRedirect cannot be supplied together with\nprefixRedirect. Supply one alone or neither. If neither is supplied, the path of the\noriginal request will be used for the redirect. The value must be between 1 and 1024\ncharacters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "prefix_redirect": {
                    "description": "The prefix that replaces the prefixMatch specified in the HttpRouteRuleMatch,\nretaining the remaining portion of the URL before redirecting the request.\nprefixRedirect cannot be supplied together with pathRedirect. Supply one alone or\nneither. If neither is supplied, the path of the original request will be used for\nthe redirect. The value must be between 1 and 1024 characters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "redirect_response_code": {
                    "description": "The HTTP Status code to use for this RedirectAction. Supported values are:\n\n* MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301.\n\n* FOUND, which corresponds to 302.\n\n* SEE_OTHER which corresponds to 303.\n\n* TEMPORARY_REDIRECT, which corresponds to 307. In this case, the request method\nwill be retained.\n\n* PERMANENT_REDIRECT, which corresponds to 308. In this case,\nthe request method will be retained. Possible values: [\"FOUND\", \"MOVED_PERMANENTLY_DEFAULT\", \"PERMANENT_REDIRECT\", \"SEE_OTHER\", \"TEMPORARY_REDIRECT\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "strip_query": {
                    "description": "If set to true, any accompanying query portion of the original URL is removed prior\nto redirecting the request. If set to false, the query portion of the original URL is\nretained.\n This field is required to ensure an empty block is not set. The normal default value is false.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "When none of the specified hostRules match, the request is redirected to a URL specified\nby defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or\ndefaultRouteAction must not be set.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "path_rule": {
              "block": {
                "attributes": {
                  "paths": {
                    "description": "The list of path patterns to match. Each must start with / and the only place a\n\\* is allowed is at the end following a /. The string fed to the path matcher\ndoes not include any text after the first ? or #, and those chars are not\nallowed here.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "service": {
                    "description": "The region backend service resource to which traffic is\ndirected if this rule is matched. If routeAction is additionally specified,\nadvanced routing actions like URL Rewrites, etc. take effect prior to sending\nthe request to the backend. However, if service is specified, routeAction cannot\ncontain any weightedBackendService s. Conversely, if routeAction specifies any\nweightedBackendServices, service must not be specified. Only one of urlRedirect,\nservice or routeAction.weightedBackendService must be set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "route_action": {
                    "block": {
                      "block_types": {
                        "cors_policy": {
                          "block": {
                            "attributes": {
                              "allow_credentials": {
                                "description": "In response to a preflight request, setting this to true indicates that the\nactual request can include user credentials. This translates to the Access-\nControl-Allow-Credentials header. Defaults to false.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "allow_headers": {
                                "description": "Specifies the content for the Access-Control-Allow-Headers header.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "allow_methods": {
                                "description": "Specifies the content for the Access-Control-Allow-Methods header.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "allow_origin_regexes": {
                                "description": "Specifies the regular expression patterns that match allowed origins. For\nregular expression grammar please see en.cppreference.com/w/cpp/regex/ecmascript\nAn origin is allowed if it matches either allow_origins or allow_origin_regex.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "allow_origins": {
                                "description": "Specifies the list of origins that will be allowed to do CORS requests. An\norigin is allowed if it matches either allow_origins or allow_origin_regex.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "disabled": {
                                "description": "If true, specifies the CORS policy is disabled.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              },
                              "expose_headers": {
                                "description": "Specifies the content for the Access-Control-Expose-Headers header.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "max_age": {
                                "description": "Specifies how long the results of a preflight request can be cached. This\ntranslates to the content for the Access-Control-Max-Age header.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "The specification for allowing client side cross-origin requests. Please see W3C\nRecommendation for Cross Origin Resource Sharing",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "fault_injection_policy": {
                          "block": {
                            "block_types": {
                              "abort": {
                                "block": {
                                  "attributes": {
                                    "http_status": {
                                      "description": "The HTTP status code used to abort the request. The value must be between 200\nand 599 inclusive.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "percentage": {
                                      "description": "The percentage of traffic (connections/operations/requests) which will be\naborted as part of fault injection. The value must be between 0.0 and 100.0\ninclusive.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "description": "The specification for how client requests are aborted as part of fault\ninjection.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "delay": {
                                "block": {
                                  "attributes": {
                                    "percentage": {
                                      "description": "The percentage of traffic (connections/operations/requests) on which delay will\nbe introduced as part of fault injection. The value must be between 0.0 and\n100.0 inclusive.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "fixed_delay": {
                                      "block": {
                                        "attributes": {
                                          "nanos": {
                                            "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations\nless than one second are represented with a 0 'seconds' field and a positive\n'nanos' field. Must be from 0 to 999,999,999 inclusive.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "seconds": {
                                            "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000\ninclusive.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Specifies the value of the fixed delay interval.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "The specification for how client requests are delayed as part of fault\ninjection, before being sent to a backend service.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "The specification for fault injection introduced into traffic to test the\nresiliency of clients to backend service failure. As part of fault injection,\nwhen clients send requests to a backend service, delays can be introduced by\nLoadbalancer on a percentage of requests before sending those request to the\nbackend service. Similarly requests from clients can be aborted by the\nLoadbalancer for a percentage of requests. timeout and retry_policy will be\nignored by clients that are configured with a fault_injection_policy.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "request_mirror_policy": {
                          "block": {
                            "attributes": {
                              "backend_service": {
                                "description": "The RegionBackendService resource being mirrored to.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Specifies the policy on how requests intended for the route's backends are\nshadowed to a separate mirrored backend service. Loadbalancer does not wait for\nresponses from the shadow service. Prior to sending traffic to the shadow\nservice, the host / authority header is suffixed with -shadow.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "retry_policy": {
                          "block": {
                            "attributes": {
                              "num_retries": {
                                "description": "Specifies the allowed number retries. This number must be \u003e 0.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "retry_conditions": {
                                "description": "Specifies one or more conditions when this retry rule applies. Valid values are:\n\n- 5xx: Loadbalancer will attempt a retry if the backend service responds with\nany 5xx response code, or if the backend service does not respond at all,\nfor example: disconnects, reset, read timeout, connection failure, and refused\nstreams.\n- gateway-error: Similar to 5xx, but only applies to response codes\n502, 503 or 504.\n- connect-failure: Loadbalancer will retry on failures\nconnecting to backend services, for example due to connection timeouts.\n- retriable-4xx: Loadbalancer will retry for retriable 4xx response codes.\nCurrently the only retriable error supported is 409.\n- refused-stream: Loadbalancer will retry if the backend service resets the stream with a\nREFUSED_STREAM error code. This reset type indicates that it is safe to retry.\n- cancelled: Loadbalancer will retry if the gRPC status code in the response\nheader is set to cancelled\n- deadline-exceeded: Loadbalancer will retry if the\ngRPC status code in the response header is set to deadline-exceeded\n- resource-exhausted: Loadbalancer will retry if the gRPC status code in the response\nheader is set to resource-exhausted\n- unavailable: Loadbalancer will retry if\nthe gRPC status code in the response header is set to unavailable",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "per_try_timeout": {
                                "block": {
                                  "attributes": {
                                    "nanos": {
                                      "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations\nless than one second are represented with a 0 'seconds' field and a positive\n'nanos' field. Must be from 0 to 999,999,999 inclusive.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "seconds": {
                                      "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000\ninclusive.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Specifies a non-zero timeout per retry attempt.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Specifies the retry policy associated with this route.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "timeout": {
                          "block": {
                            "attributes": {
                              "nanos": {
                                "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations\nless than one second are represented with a 0 'seconds' field and a positive\n'nanos' field. Must be from 0 to 999,999,999 inclusive.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "seconds": {
                                "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000\ninclusive.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Specifies the timeout for the selected route. Timeout is computed from the time\nthe request is has been fully processed (i.e. end-of-stream) up until the\nresponse has been completely processed. Timeout includes all retries. If not\nspecified, the default value is 15 seconds.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "url_rewrite": {
                          "block": {
                            "attributes": {
                              "host_rewrite": {
                                "description": "Prior to forwarding the request to the selected service, the request's host\nheader is replaced with contents of hostRewrite. The value must be between 1 and\n255 characters.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path_prefix_rewrite": {
                                "description": "Prior to forwarding the request to the selected backend service, the matching\nportion of the request's path is replaced by pathPrefixRewrite. The value must\nbe between 1 and 1024 characters.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The spec to modify the URL of the request, prior to forwarding the request to\nthe matched service",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "weighted_backend_services": {
                          "block": {
                            "attributes": {
                              "backend_service": {
                                "description": "The default RegionBackendService resource. Before\nforwarding the request to backendService, the loadbalancer applies any relevant\nheaderActions specified as part of this backendServiceWeight.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "weight": {
                                "description": "Specifies the fraction of traffic sent to backendService, computed as weight /\n(sum of all weightedBackendService weights in routeAction) . The selection of a\nbackend service is determined only for new traffic. Once a user's request has\nbeen directed to a backendService, subsequent requests will be sent to the same\nbackendService as determined by the BackendService's session affinity policy.\nThe value must be between 0 and 1000",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "header_action": {
                                "block": {
                                  "attributes": {
                                    "request_headers_to_remove": {
                                      "description": "A list of header names for headers that need to be removed from the request\nprior to forwarding the request to the backendService.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "response_headers_to_remove": {
                                      "description": "A list of header names for headers that need to be removed from the response\nprior to sending the response back to the client.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "request_headers_to_add": {
                                      "block": {
                                        "attributes": {
                                          "header_name": {
                                            "description": "The name of the header.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "header_value": {
                                            "description": "The value of the header to add.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "replace": {
                                            "description": "If false, headerValue is appended to any values that already exist for the\nheader. If true, headerValue is set for the header, discarding any values that\nwere set for that header.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "bool"
                                          }
                                        },
                                        "description": "Headers to add to a matching request prior to forwarding the request to the\nbackendService.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response_headers_to_add": {
                                      "block": {
                                        "attributes": {
                                          "header_name": {
                                            "description": "The name of the header.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "header_value": {
                                            "description": "The value of the header to add.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "replace": {
                                            "description": "If false, headerValue is appended to any values that already exist for the\nheader. If true, headerValue is set for the header, discarding any values that\nwere set for that header.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "bool"
                                          }
                                        },
                                        "description": "Headers to add the response prior to sending the response back to the client.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Specifies changes to request and response headers that need to take effect for\nthe selected backendService. headerAction specified here take effect before\nheaderAction in the enclosing HttpRouteRule, PathMatcher and UrlMap.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A list of weighted backend services to send traffic to when a route match\noccurs. The weights determine the fraction of traffic that flows to their\ncorresponding backend service. If all traffic needs to go to a single backend\nservice, there must be one  weightedBackendService with weight set to a non 0\nnumber. Once a backendService is identified and before forwarding the request to\nthe backend service, advanced routing actions like Url rewrites and header\ntransformations are applied depending on additional settings specified in this\nHttpRouteAction.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "In response to a matching path, the load balancer performs advanced routing\nactions like URL rewrites, header transformations, etc. prior to forwarding the\nrequest to the selected backend. If routeAction specifies any\nweightedBackendServices, service must not be set. Conversely if service is set,\nrouteAction cannot contain any  weightedBackendServices. Only one of routeAction\nor urlRedirect must be set.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "url_redirect": {
                    "block": {
                      "attributes": {
                        "host_redirect": {
                          "description": "The host that will be used in the redirect response instead of the one\nthat was supplied in the request. The value must be between 1 and 255\ncharacters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "https_redirect": {
                          "description": "If set to true, the URL scheme in the redirected request is set to https.\nIf set to false, the URL scheme of the redirected request will remain the\nsame as that of the request. This must only be set for UrlMaps used in\nTargetHttpProxys. Setting this true for TargetHttpsProxy is not\npermitted. The default is set to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "path_redirect": {
                          "description": "The path that will be used in the redirect response instead of the one\nthat was supplied in the request. pathRedirect cannot be supplied\ntogether with prefixRedirect. Supply one alone or neither. If neither is\nsupplied, the path of the original request will be used for the redirect.\nThe value must be between 1 and 1024 characters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "prefix_redirect": {
                          "description": "The prefix that replaces the prefixMatch specified in the\nHttpRouteRuleMatch, retaining the remaining portion of the URL before\nredirecting the request. prefixRedirect cannot be supplied together with\npathRedirect. Supply one alone or neither. If neither is supplied, the\npath of the original request will be used for the redirect. The value\nmust be between 1 and 1024 characters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "redirect_response_code": {
                          "description": "The HTTP Status code to use for this RedirectAction. Supported values are:\n\n* MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301.\n\n* FOUND, which corresponds to 302.\n\n* SEE_OTHER which corresponds to 303.\n\n* TEMPORARY_REDIRECT, which corresponds to 307. In this case, the request method\nwill be retained.\n\n* PERMANENT_REDIRECT, which corresponds to 308. In this case,\nthe request method will be retained. Possible values: [\"FOUND\", \"MOVED_PERMANENTLY_DEFAULT\", \"PERMANENT_REDIRECT\", \"SEE_OTHER\", \"TEMPORARY_REDIRECT\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "strip_query": {
                          "description": "If set to true, any accompanying query portion of the original URL is removed\nprior to redirecting the request. If set to false, the query portion of the\noriginal URL is retained.\n This field is required to ensure an empty block is not set. The normal default value is false.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "description": "When a path pattern is matched, the request is redirected to a URL specified\nby urlRedirect. If urlRedirect is specified, service or routeAction must not\nbe set.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The list of path rules. Use this list instead of routeRules when routing based\non simple path matching is all that's required. The order by which path rules\nare specified does not matter. Matches are always done on the longest-path-first\nbasis. For example: a pathRule with a path /a/b/c/* will match before /a/b/*\nirrespective of the order in which those paths appear in this list. Within a\ngiven pathMatcher, only one of pathRules or routeRules must be set.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "route_rules": {
              "block": {
                "attributes": {
                  "priority": {
                    "description": "For routeRules within a given pathMatcher, priority determines the order\nin which load balancer will interpret routeRules. RouteRules are evaluated\nin order of priority, from the lowest to highest number. The priority of\na rule decreases as its number increases (1, 2, 3, N+1). The first rule\nthat matches the request is applied.\n\nYou cannot configure two or more routeRules with the same priority.\nPriority for each rule must be set to a number between 0 and\n2147483647 inclusive.\n\nPriority numbers can have gaps, which enable you to add or remove rules\nin the future without affecting the rest of the rules. For example,\n1, 2, 3, 4, 5, 9, 12, 16 is a valid series of priority numbers to which\nyou could add rules numbered from 6 to 8, 10 to 11, and 13 to 15 in the\nfuture without any impact on existing rules.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "service": {
                    "description": "The region backend service resource to which traffic is\ndirected if this rule is matched. If routeAction is additionally specified,\nadvanced routing actions like URL Rewrites, etc. take effect prior to sending\nthe request to the backend. However, if service is specified, routeAction cannot\ncontain any weightedBackendService s. Conversely, if routeAction specifies any\nweightedBackendServices, service must not be specified. Only one of urlRedirect,\nservice or routeAction.weightedBackendService must be set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "header_action": {
                    "block": {
                      "attributes": {
                        "request_headers_to_remove": {
                          "description": "A list of header names for headers that need to be removed from the request\nprior to forwarding the request to the backendService.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "response_headers_to_remove": {
                          "description": "A list of header names for headers that need to be removed from the response\nprior to sending the response back to the client.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "request_headers_to_add": {
                          "block": {
                            "attributes": {
                              "header_name": {
                                "description": "The name of the header.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "header_value": {
                                "description": "The value of the header to add.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "replace": {
                                "description": "If false, headerValue is appended to any values that already exist for the\nheader. If true, headerValue is set for the header, discarding any values that\nwere set for that header.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              }
                            },
                            "description": "Headers to add to a matching request prior to forwarding the request to the\nbackendService.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "response_headers_to_add": {
                          "block": {
                            "attributes": {
                              "header_name": {
                                "description": "The name of the header.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "header_value": {
                                "description": "The value of the header to add.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "replace": {
                                "description": "If false, headerValue is appended to any values that already exist for the\nheader. If true, headerValue is set for the header, discarding any values that\nwere set for that header.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              }
                            },
                            "description": "Headers to add the response prior to sending the response back to the client.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Specifies changes to request and response headers that need to take effect for\nthe selected backendService. The headerAction specified here are applied before\nthe matching pathMatchers[].headerAction and after pathMatchers[].routeRules[].r\nouteAction.weightedBackendService.backendServiceWeightAction[].headerAction",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "match_rules": {
                    "block": {
                      "attributes": {
                        "full_path_match": {
                          "description": "For satisfying the matchRule condition, the path of the request must exactly\nmatch the value specified in fullPathMatch after removing any query parameters\nand anchor that may be part of the original URL. FullPathMatch must be between 1\nand 1024 characters. Only one of prefixMatch, fullPathMatch or regexMatch must\nbe specified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "ignore_case": {
                          "description": "Specifies that prefixMatch and fullPathMatch matches are case sensitive.\nDefaults to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "path_template_match": {
                          "description": "For satisfying the matchRule condition, the path of the request\nmust match the wildcard pattern specified in pathTemplateMatch\nafter removing any query parameters and anchor that may be part\nof the original URL.\n\npathTemplateMatch must be between 1 and 255 characters\n(inclusive).  The pattern specified by pathTemplateMatch may\nhave at most 5 wildcard operators and at most 5 variable\ncaptures in total.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "prefix_match": {
                          "description": "For satisfying the matchRule condition, the request's path must begin with the\nspecified prefixMatch. prefixMatch must begin with a /. The value must be\nbetween 1 and 1024 characters. Only one of prefixMatch, fullPathMatch or\nregexMatch must be specified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "regex_match": {
                          "description": "For satisfying the matchRule condition, the path of the request must satisfy the\nregular expression specified in regexMatch after removing any query parameters\nand anchor supplied with the original URL. For regular expression grammar please\nsee en.cppreference.com/w/cpp/regex/ecmascript  Only one of prefixMatch,\nfullPathMatch or regexMatch must be specified.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "header_matches": {
                          "block": {
                            "attributes": {
                              "exact_match": {
                                "description": "The value should exactly match contents of exactMatch. Only one of exactMatch,\nprefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "header_name": {
                                "description": "The name of the HTTP header to match. For matching against the HTTP request's\nauthority, use a headerMatch with the header name \":authority\". For matching a\nrequest's method, use the headerName \":method\".",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "invert_match": {
                                "description": "If set to false, the headerMatch is considered a match if the match criteria\nabove are met. If set to true, the headerMatch is considered a match if the\nmatch criteria above are NOT met. Defaults to false.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "prefix_match": {
                                "description": "The value of the header must start with the contents of prefixMatch. Only one of\nexactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch\nmust be set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "present_match": {
                                "description": "A header with the contents of headerName must exist. The match takes place\nwhether or not the request's header has a value or not. Only one of exactMatch,\nprefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "regex_match": {
                                "description": "The value of the header must match the regular expression specified in\nregexMatch. For regular expression grammar, please see:\nen.cppreference.com/w/cpp/regex/ecmascript  For matching against a port\nspecified in the HTTP request, use a headerMatch with headerName set to PORT and\na regular expression that satisfies the RFC2616 Host header's port specifier.\nOnly one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or\nrangeMatch must be set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "suffix_match": {
                                "description": "The value of the header must end with the contents of suffixMatch. Only one of\nexactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch\nmust be set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "range_match": {
                                "block": {
                                  "attributes": {
                                    "range_end": {
                                      "description": "The end of the range (exclusive).",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "range_start": {
                                      "description": "The start of the range (inclusive).",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "description": "The header value must be an integer and its value must be in the range specified\nin rangeMatch. If the header does not contain an integer, number or is empty,\nthe match fails. For example for a range [-5, 0]\n\n* -3 will match\n* 0 will not match\n* 0.25 will not match\n* -3someString will not match.\n\nOnly one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or\nrangeMatch must be set.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Specifies a list of header match criteria, all of which must match corresponding\nheaders in the request.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "metadata_filters": {
                          "block": {
                            "attributes": {
                              "filter_match_criteria": {
                                "description": "Specifies how individual filterLabel matches within the list of filterLabels\ncontribute towards the overall metadataFilter match. Supported values are:\n\n* MATCH_ANY: At least one of the filterLabels must have a matching label in the\nprovided metadata.\n* MATCH_ALL: All filterLabels must have matching labels in\nthe provided metadata. Possible values: [\"MATCH_ALL\", \"MATCH_ANY\"]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "filter_labels": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "Name of metadata label. The name can have a maximum length of 1024 characters\nand must be at least 1 character long.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "The value of the label must match the specified value. value can have a maximum\nlength of 1024 characters.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "The list of label value pairs that must match labels in the provided metadata\nbased on filterMatchCriteria  This list must not be empty and can have at the\nmost 64 entries.",
                                  "description_kind": "plain"
                                },
                                "max_items": 64,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Opaque filter criteria used by Loadbalancer to restrict routing configuration to\na limited set xDS compliant clients. In their xDS requests to Loadbalancer, xDS\nclients present node metadata. If a match takes place, the relevant routing\nconfiguration is made available to those proxies. For each metadataFilter in\nthis list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the\nfilterLabels must match the corresponding label provided in the metadata. If its\nfilterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match\nwith corresponding labels in the provided metadata. metadataFilters specified\nhere can be overrides those specified in ForwardingRule that refers to this\nUrlMap. metadataFilters only applies to Loadbalancers that have their\nloadBalancingScheme set to INTERNAL_SELF_MANAGED.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "query_parameter_matches": {
                          "block": {
                            "attributes": {
                              "exact_match": {
                                "description": "The queryParameterMatch matches if the value of the parameter exactly matches\nthe contents of exactMatch. Only one of presentMatch, exactMatch and regexMatch\nmust be set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "description": "The name of the query parameter to match. The query parameter must exist in the\nrequest, in the absence of which the request match fails.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "present_match": {
                                "description": "Specifies that the queryParameterMatch matches if the request contains the query\nparameter, irrespective of whether the parameter has a value or not. Only one of\npresentMatch, exactMatch and regexMatch must be set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "regex_match": {
                                "description": "The queryParameterMatch matches if the value of the parameter matches the\nregular expression specified by regexMatch. For the regular expression grammar,\nplease see en.cppreference.com/w/cpp/regex/ecmascript  Only one of presentMatch,\nexactMatch and regexMatch must be set.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Specifies a list of query parameter match criteria, all of which must match\ncorresponding query parameters in the request.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The rules for determining a match.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "route_action": {
                    "block": {
                      "block_types": {
                        "cors_policy": {
                          "block": {
                            "attributes": {
                              "allow_credentials": {
                                "description": "In response to a preflight request, setting this to true indicates that the\nactual request can include user credentials. This translates to the Access-\nControl-Allow-Credentials header. Defaults to false.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "allow_headers": {
                                "description": "Specifies the content for the Access-Control-Allow-Headers header.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "allow_methods": {
                                "description": "Specifies the content for the Access-Control-Allow-Methods header.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "allow_origin_regexes": {
                                "description": "Specifies the regular expression patterns that match allowed origins. For\nregular expression grammar please see en.cppreference.com/w/cpp/regex/ecmascript\nAn origin is allowed if it matches either allow_origins or allow_origin_regex.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "allow_origins": {
                                "description": "Specifies the list of origins that will be allowed to do CORS requests. An\norigin is allowed if it matches either allow_origins or allow_origin_regex.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "disabled": {
                                "description": "If true, specifies the CORS policy is disabled.\nwhich indicates that the CORS policy is in effect. Defaults to false.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "expose_headers": {
                                "description": "Specifies the content for the Access-Control-Expose-Headers header.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "max_age": {
                                "description": "Specifies how long the results of a preflight request can be cached. This\ntranslates to the content for the Access-Control-Max-Age header.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "The specification for allowing client side cross-origin requests. Please see W3C\nRecommendation for Cross Origin Resource Sharing",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "fault_injection_policy": {
                          "block": {
                            "block_types": {
                              "abort": {
                                "block": {
                                  "attributes": {
                                    "http_status": {
                                      "description": "The HTTP status code used to abort the request. The value must be between 200\nand 599 inclusive.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "percentage": {
                                      "description": "The percentage of traffic (connections/operations/requests) which will be\naborted as part of fault injection. The value must be between 0.0 and 100.0\ninclusive.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "description": "The specification for how client requests are aborted as part of fault\ninjection.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "delay": {
                                "block": {
                                  "attributes": {
                                    "percentage": {
                                      "description": "The percentage of traffic (connections/operations/requests) on which delay will\nbe introduced as part of fault injection. The value must be between 0.0 and\n100.0 inclusive.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "fixed_delay": {
                                      "block": {
                                        "attributes": {
                                          "nanos": {
                                            "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations\nless than one second are represented with a 0 'seconds' field and a positive\n'nanos' field. Must be from 0 to 999,999,999 inclusive.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "seconds": {
                                            "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000\ninclusive.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Specifies the value of the fixed delay interval.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "The specification for how client requests are delayed as part of fault\ninjection, before being sent to a backend service.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "The specification for fault injection introduced into traffic to test the\nresiliency of clients to backend service failure. As part of fault injection,\nwhen clients send requests to a backend service, delays can be introduced by\nLoadbalancer on a percentage of requests before sending those request to the\nbackend service. Similarly requests from clients can be aborted by the\nLoadbalancer for a percentage of requests. timeout and retry_policy will be\nignored by clients that are configured with a fault_injection_policy.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "request_mirror_policy": {
                          "block": {
                            "attributes": {
                              "backend_service": {
                                "description": "The RegionBackendService resource being mirrored to.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Specifies the policy on how requests intended for the route's backends are\nshadowed to a separate mirrored backend service. Loadbalancer does not wait for\nresponses from the shadow service. Prior to sending traffic to the shadow\nservice, the host / authority header is suffixed with -shadow.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "retry_policy": {
                          "block": {
                            "attributes": {
                              "num_retries": {
                                "description": "Specifies the allowed number retries. This number must be \u003e 0.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "retry_conditions": {
                                "description": "Specifies one or more conditions when this retry rule applies. Valid values are:\n\n* 5xx: Loadbalancer will attempt a retry if the backend service responds with\n  any 5xx response code, or if the backend service does not respond at all,\n  for example: disconnects, reset, read timeout, connection failure, and refused\n  streams.\n* gateway-error: Similar to 5xx, but only applies to response codes\n  502, 503 or 504.\n* connect-failure: Loadbalancer will retry on failures\n  connecting to backend services, for example due to connection timeouts.\n* retriable-4xx: Loadbalancer will retry for retriable 4xx response codes.\n  Currently the only retriable error supported is 409.\n* refused-stream: Loadbalancer will retry if the backend service resets the stream with a\n  REFUSED_STREAM error code. This reset type indicates that it is safe to retry.\n* cancelled: Loadbalancer will retry if the gRPC status code in the response\n  header is set to cancelled\n* deadline-exceeded: Loadbalancer will retry if the\n  gRPC status code in the response header is set to deadline-exceeded\n* resource-exhausted: Loadbalancer will retry if the gRPC status code in the response\n  header is set to resource-exhausted\n* unavailable: Loadbalancer will retry if the gRPC status code in\n  the response header is set to unavailable",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "per_try_timeout": {
                                "block": {
                                  "attributes": {
                                    "nanos": {
                                      "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations\nless than one second are represented with a 0 'seconds' field and a positive\n'nanos' field. Must be from 0 to 999,999,999 inclusive.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "seconds": {
                                      "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000\ninclusive.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Specifies a non-zero timeout per retry attempt.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Specifies the retry policy associated with this route.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "timeout": {
                          "block": {
                            "attributes": {
                              "nanos": {
                                "description": "Span of time that's a fraction of a second at nanosecond resolution. Durations\nless than one second are represented with a 0 'seconds' field and a positive\n'nanos' field. Must be from 0 to 999,999,999 inclusive.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "seconds": {
                                "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000\ninclusive.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Specifies the timeout for the selected route. Timeout is computed from the time\nthe request is has been fully processed (i.e. end-of-stream) up until the\nresponse has been completely processed. Timeout includes all retries. If not\nspecified, the default value is 15 seconds.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "url_rewrite": {
                          "block": {
                            "attributes": {
                              "host_rewrite": {
                                "description": "Prior to forwarding the request to the selected service, the request's host\nheader is replaced with contents of hostRewrite. The value must be between 1 and\n255 characters.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path_prefix_rewrite": {
                                "description": "Prior to forwarding the request to the selected backend service, the matching\nportion of the request's path is replaced by pathPrefixRewrite. The value must\nbe between 1 and 1024 characters.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path_template_rewrite": {
                                "description": "Prior to forwarding the request to the selected origin, if the\nrequest matched a pathTemplateMatch, the matching portion of the\nrequest's path is replaced re-written using the pattern specified\nby pathTemplateRewrite.\n\npathTemplateRewrite must be between 1 and 255 characters\n(inclusive), must start with a '/', and must only use variables\ncaptured by the route's pathTemplate matchers.\n\npathTemplateRewrite may only be used when all of a route's\nMatchRules specify pathTemplate.\n\nOnly one of pathPrefixRewrite and pathTemplateRewrite may be\nspecified.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The spec to modify the URL of the request, prior to forwarding the request to\nthe matched service",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "weighted_backend_services": {
                          "block": {
                            "attributes": {
                              "backend_service": {
                                "description": "The default RegionBackendService resource. Before\nforwarding the request to backendService, the loadbalancer applies any relevant\nheaderActions specified as part of this backendServiceWeight.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "weight": {
                                "description": "Specifies the fraction of traffic sent to backendService, computed as weight /\n(sum of all weightedBackendService weights in routeAction) . The selection of a\nbackend service is determined only for new traffic. Once a user's request has\nbeen directed to a backendService, subsequent requests will be sent to the same\nbackendService as determined by the BackendService's session affinity policy.\nThe value must be between 0 and 1000",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "header_action": {
                                "block": {
                                  "attributes": {
                                    "request_headers_to_remove": {
                                      "description": "A list of header names for headers that need to be removed from the request\nprior to forwarding the request to the backendService.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "response_headers_to_remove": {
                                      "description": "A list of header names for headers that need to be removed from the response\nprior to sending the response back to the client.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "request_headers_to_add": {
                                      "block": {
                                        "attributes": {
                                          "header_name": {
                                            "description": "The name of the header.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "header_value": {
                                            "description": "The value of the header to add.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "replace": {
                                            "description": "If false, headerValue is appended to any values that already exist for the\nheader. If true, headerValue is set for the header, discarding any values that\nwere set for that header.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "bool"
                                          }
                                        },
                                        "description": "Headers to add to a matching request prior to forwarding the request to the\nbackendService.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response_headers_to_add": {
                                      "block": {
                                        "attributes": {
                                          "header_name": {
                                            "description": "The name of the header.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "header_value": {
                                            "description": "The value of the header to add.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "replace": {
                                            "description": "If false, headerValue is appended to any values that already exist for the\nheader. If true, headerValue is set for the header, discarding any values that\nwere set for that header.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "bool"
                                          }
                                        },
                                        "description": "Headers to add the response prior to sending the response back to the client.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Specifies changes to request and response headers that need to take effect for\nthe selected backendService. headerAction specified here take effect before\nheaderAction in the enclosing HttpRouteRule, PathMatcher and UrlMap.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A list of weighted backend services to send traffic to when a route match\noccurs. The weights determine the fraction of traffic that flows to their\ncorresponding backend service. If all traffic needs to go to a single backend\nservice, there must be one  weightedBackendService with weight set to a non 0\nnumber. Once a backendService is identified and before forwarding the request to\nthe backend service, advanced routing actions like Url rewrites and header\ntransformations are applied depending on additional settings specified in this\nHttpRouteAction.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "In response to a matching matchRule, the load balancer performs advanced routing\nactions like URL rewrites, header transformations, etc. prior to forwarding the\nrequest to the selected backend. If  routeAction specifies any\nweightedBackendServices, service must not be set. Conversely if service is set,\nrouteAction cannot contain any  weightedBackendServices. Only one of routeAction\nor urlRedirect must be set.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "url_redirect": {
                    "block": {
                      "attributes": {
                        "host_redirect": {
                          "description": "The host that will be used in the redirect response instead of the one\nthat was supplied in the request. The value must be between 1 and 255\ncharacters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "https_redirect": {
                          "description": "If set to true, the URL scheme in the redirected request is set to https.\nIf set to false, the URL scheme of the redirected request will remain the\nsame as that of the request. This must only be set for UrlMaps used in\nTargetHttpProxys. Setting this true for TargetHttpsProxy is not\npermitted. The default is set to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "path_redirect": {
                          "description": "The path that will be used in the redirect response instead of the one\nthat was supplied in the request. pathRedirect cannot be supplied\ntogether with prefixRedirect. Supply one alone or neither. If neither is\nsupplied, the path of the original request will be used for the redirect.\nThe value must be between 1 and 1024 characters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "prefix_redirect": {
                          "description": "The prefix that replaces the prefixMatch specified in the\nHttpRouteRuleMatch, retaining the remaining portion of the URL before\nredirecting the request. prefixRedirect cannot be supplied together with\npathRedirect. Supply one alone or neither. If neither is supplied, the\npath of the original request will be used for the redirect. The value\nmust be between 1 and 1024 characters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "redirect_response_code": {
                          "description": "The HTTP Status code to use for this RedirectAction. Supported values are:\n\n* MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301.\n\n* FOUND, which corresponds to 302.\n\n* SEE_OTHER which corresponds to 303.\n\n* TEMPORARY_REDIRECT, which corresponds to 307. In this case, the request method\nwill be retained.\n\n* PERMANENT_REDIRECT, which corresponds to 308. In this case,\nthe request method will be retained. Possible values: [\"FOUND\", \"MOVED_PERMANENTLY_DEFAULT\", \"PERMANENT_REDIRECT\", \"SEE_OTHER\", \"TEMPORARY_REDIRECT\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "strip_query": {
                          "description": "If set to true, any accompanying query portion of the original URL is\nremoved prior to redirecting the request. If set to false, the query\nportion of the original URL is retained. The default value is false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "When this rule is matched, the request is redirected to a URL specified by\nurlRedirect. If urlRedirect is specified, service or routeAction must not be\nset.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The list of ordered HTTP route rules. Use this list instead of pathRules when\nadvanced route matching and routing actions are desired. The order of specifying\nrouteRules matters: the first rule that matches will cause its specified routing\naction to take effect. Within a given pathMatcher, only one of pathRules or\nrouteRules must be set. routeRules are not supported in UrlMaps intended for\nExternal load balancers.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "The list of named PathMatchers to use against the URL.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "test": {
        "block": {
          "attributes": {
            "description": {
              "description": "Description of this test case.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host": {
              "description": "Host portion of the URL.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description": "Path portion of the URL.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service": {
              "description": "A reference to expected RegionBackendService resource the given URL should be mapped to.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The list of expected URL mappings. Requests to update this UrlMap will\nsucceed only if all of the test cases pass.",
          "description_kind": "plain"
        },
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

func GoogleComputeRegionUrlMapSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionUrlMap), &result)
	return &result
}
