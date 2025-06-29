package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeUrlMap = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_service": {
        "description": "The backend service or backend bucket to use when none of the given rules match.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when you create\nthe resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource. A hash of the contents stored in this object. This\nfield is used in optimistic locking.",
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
        "description": "Name of the resource. Provided by the client when the resource is created. The\nname must be 1-63 characters long, and comply with RFC1035. Specifically, the\nname must be 1-63 characters long and match the regular expression\n'[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase\nletter, and all following characters must be a dash, lowercase letter, or digit,\nexcept the last character, which cannot be a dash.",
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
      }
    },
    "block_types": {
      "default_custom_error_response_policy": {
        "block": {
          "attributes": {
            "error_service": {
              "description": "The full or partial URL to the BackendBucket resource that contains the custom error content. Examples are:\n\nhttps://www.googleapis.com/compute/v1/projects/project/global/backendBuckets/myBackendBucket\ncompute/v1/projects/project/global/backendBuckets/myBackendBucket\nglobal/backendBuckets/myBackendBucket\n\nIf errorService is not specified at lower levels like pathMatcher, pathRule and routeRule, an errorService specified at a higher level in the UrlMap will be used. If UrlMap.defaultCustomErrorResponsePolicy contains one or more errorResponseRules[], it must specify errorService.\nIf load balancer cannot reach the backendBucket, a simple Not Found Error will be returned, with the original response code (or overrideResponseCode if configured).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "error_response_rule": {
              "block": {
                "attributes": {
                  "match_response_codes": {
                    "description": "Valid values include:\n- A number between 400 and 599: For example 401 or 503, in which case the load balancer applies the policy if the error code exactly matches this value.\n- 5xx: Load Balancer will apply the policy if the backend service responds with any response code in the range of 500 to 599.\n- 4xx: Load Balancer will apply the policy if the backend service responds with any response code in the range of 400 to 499.\nValues must be unique within matchResponseCodes and across all errorResponseRules of CustomErrorResponsePolicy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "override_response_code": {
                    "description": "The HTTP status code returned with the response containing the custom error content.\nIf overrideResponseCode is not supplied, the same response code returned by the original backend bucket or backend service is returned to the client.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "path": {
                    "description": "The full path to a file within backendBucket. For example: /errors/defaultError.html\npath must start with a leading slash. path cannot have trailing slashes.\nIf the file is not available in backendBucket or the load balancer cannot reach the BackendBucket, a simple Not Found Error is returned to the client.\nThe value must be from 1 to 1024 characters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Specifies rules for returning error responses.\nIn a given policy, if you specify rules for both a range of error codes as well as rules for specific error codes then rules with specific error codes have a higher priority.\nFor example, assume that you configure a rule for 401 (Un-authorized) code, and another for all 4 series error codes (4XX).\nIf the backend service returns a 401, then the rule for 401 will be applied. However if the backend service returns a 403, the rule for 4xx takes effect.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendService or BackendBucket responds with an error.\n\nThis policy takes effect at the PathMatcher level and applies only when no policy has been defined for the error code at lower levels like RouteRule and PathRule within this PathMatcher. If an error code does not have a policy defined in defaultCustomErrorResponsePolicy, then a policy defined for the error code in UrlMap.defaultCustomErrorResponsePolicy takes effect.\n\nFor example, consider a UrlMap with the following configuration:\n\nUrlMap.defaultCustomErrorResponsePolicy is configured with policies for 5xx and 4xx errors\nA RouteRule for /coming_soon/ is configured for the error code 404.\nIf the request is for www.myotherdomain.com and a 404 is encountered, the policy under UrlMap.defaultCustomErrorResponsePolicy takes effect. If a 404 response is encountered for the request www.example.com/current_events/, the pathMatcher's policy takes effect. If however, the request for www.example.com/coming_soon/ encounters a 404, the policy in RouteRule.customErrorResponsePolicy takes effect. If any of the requests in this example encounter a 500 error code, the policy at UrlMap.defaultCustomErrorResponsePolicy takes effect.\n\nWhen used in conjunction with pathMatcher.defaultRouteAction.retryPolicy, retries take precedence. Only once all retries are exhausted, the defaultCustomErrorResponsePolicy is applied. While attempting a retry, if load balancer is successful in reaching the service, the defaultCustomErrorResponsePolicy is ignored and the response from the service is returned to the client.\n\ndefaultCustomErrorResponsePolicy is supported only for global external Application Load Balancers.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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
                "description": "A list of weighted backend services to send traffic to when a route match occurs.\nThe weights determine the fraction of traffic that flows to their corresponding backend service.\nIf all traffic needs to go to a single backend service, there must be one weightedBackendService\nwith weight set to a non 0 number.\n\nOnce a backendService is identified and before forwarding the request to the backend service,\nadvanced routing actions like Url rewrites and header transformations are applied depending on\nadditional settings specified in this HttpRouteAction.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions\nlike URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend.\nIf defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService\nis set, defaultRouteAction cannot contain any weightedBackendServices.\n\nOnly one of defaultRouteAction or defaultUrlRedirect must be set.",
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
              "description": "If set to true, any accompanying query portion of the original URL is removed prior\nto redirecting the request. If set to false, the query portion of the original URL is\nretained. The default is set to false.\n This field is required to ensure an empty block is not set. The normal default value is false.",
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
          "description": "Specifies changes to request and response headers that need to take effect for\nthe selected backendService. The headerAction specified here take effect after\nheaderAction specified under pathMatcher.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "host_rule": {
        "block": {
          "attributes": {
            "description": {
              "description": "An optional description of this resource. Provide this property when you create\nthe resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hosts": {
              "description": "The list of host patterns to match. They must be valid hostnames, except * will\nmatch any string of ([a-z0-9-.]*). In that case, * must be the first character\nand must be followed in the pattern by either - or ..",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "path_matcher": {
              "description": "The name of the PathMatcher to use to match the path portion of the URL if the\nhostRule matches the URL's host portion.",
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
              "description": "The backend service or backend bucket to use when none of the given paths match.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "description": "An optional description of this resource. Provide this property when you create\nthe resource.",
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
            "default_custom_error_response_policy": {
              "block": {
                "attributes": {
                  "error_service": {
                    "description": "The full or partial URL to the BackendBucket resource that contains the custom error content. Examples are:\nhttps://www.googleapis.com/compute/v1/projects/project/global/backendBuckets/myBackendBucket\ncompute/v1/projects/project/global/backendBuckets/myBackendBucket\nglobal/backendBuckets/myBackendBucket\nIf errorService is not specified at lower levels like pathMatcher, pathRule and routeRule, an errorService specified at a higher level in the UrlMap will be used. If UrlMap.defaultCustomErrorResponsePolicy contains one or more errorResponseRules[], it must specify errorService.\nIf load balancer cannot reach the backendBucket, a simple Not Found Error will be returned, with the original response code (or overrideResponseCode if configured).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "error_response_rule": {
                    "block": {
                      "attributes": {
                        "match_response_codes": {
                          "description": "Valid values include:\n- A number between 400 and 599: For example 401 or 503, in which case the load balancer applies the policy if the error code exactly matches this value.\n- 5xx: Load Balancer will apply the policy if the backend service responds with any response code in the range of 500 to 599.\n- 4xx: Load Balancer will apply the policy if the backend service responds with any response code in the range of 400 to 499.\nValues must be unique within matchResponseCodes and across all errorResponseRules of CustomErrorResponsePolicy.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "override_response_code": {
                          "description": "The HTTP status code returned with the response containing the custom error content.\nIf overrideResponseCode is not supplied, the same response code returned by the original backend bucket or backend service is returned to the client.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "path": {
                          "description": "The full path to a file within backendBucket . For example: /errors/defaultError.html\npath must start with a leading slash. path cannot have trailing slashes.\nIf the file is not available in backendBucket or the load balancer cannot reach the BackendBucket, a simple Not Found Error is returned to the client.\nThe value must be from 1 to 1024 characters",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Specifies rules for returning error responses.\nIn a given policy, if you specify rules for both a range of error codes as well as rules for specific error codes then rules with specific error codes have a higher priority.\nFor example, assume that you configure a rule for 401 (Un-authorized) code, and another for all 4 series error codes (4XX).\nIf the backend service returns a 401, then the rule for 401 will be applied. However if the backend service returns a 403, the rule for 4xx takes effect.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendService or BackendBucket responds with an error.\n\nThis policy takes effect at the PathMatcher level and applies only when no policy has been defined for the error code at lower levels like RouteRule and PathRule within this PathMatcher. If an error code does not have a policy defined in defaultCustomErrorResponsePolicy, then a policy defined for the error code in UrlMap.defaultCustomErrorResponsePolicy takes effect.\n\nFor example, consider a UrlMap with the following configuration:\n\nUrlMap.defaultCustomErrorResponsePolicy is configured with policies for 5xx and 4xx errors\nA RouteRule for /coming_soon/ is configured for the error code 404.\nIf the request is for www.myotherdomain.com and a 404 is encountered, the policy under UrlMap.defaultCustomErrorResponsePolicy takes effect. If a 404 response is encountered for the request www.example.com/current_events/, the pathMatcher's policy takes effect. If however, the request for www.example.com/coming_soon/ encounters a 404, the policy in RouteRule.customErrorResponsePolicy takes effect. If any of the requests in this example encounter a 500 error code, the policy at UrlMap.defaultCustomErrorResponsePolicy takes effect.\n\nWhen used in conjunction with pathMatcher.defaultRouteAction.retryPolicy, retries take precedence. Only once all retries are exhausted, the defaultCustomErrorResponsePolicy is applied. While attempting a retry, if load balancer is successful in reaching the service, the defaultCustomErrorResponsePolicy is ignored and the response from the service is returned to the client.\n\ndefaultCustomErrorResponsePolicy is supported only for global external Application Load Balancers.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
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
                      "description": "A list of weighted backend services to send traffic to when a route match occurs.\nThe weights determine the fraction of traffic that flows to their corresponding backend service.\nIf all traffic needs to go to a single backend service, there must be one weightedBackendService\nwith weight set to a non 0 number.\n\nOnce a backendService is identified and before forwarding the request to the backend service,\nadvanced routing actions like Url rewrites and header transformations are applied depending on\nadditional settings specified in this HttpRouteAction.",
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
                "description": "Specifies changes to request and response headers that need to take effect for\nthe selected backendService. HeaderAction specified here are applied after the\nmatching HttpRouteRule HeaderAction and before the HeaderAction in the UrlMap",
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
                    "description": "The backend service or backend bucket to use if any of the given paths match.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "custom_error_response_policy": {
                    "block": {
                      "attributes": {
                        "error_service": {
                          "description": "The full or partial URL to the BackendBucket resource that contains the custom error content. Examples are:\n\nhttps://www.googleapis.com/compute/v1/projects/project/global/backendBuckets/myBackendBucket\ncompute/v1/projects/project/global/backendBuckets/myBackendBucket\nglobal/backendBuckets/myBackendBucket\n\nIf errorService is not specified at lower levels like pathMatcher, pathRule and routeRule, an errorService specified at a higher level in the UrlMap will be used. If UrlMap.defaultCustomErrorResponsePolicy contains one or more errorResponseRules[], it must specify errorService.\nIf load balancer cannot reach the backendBucket, a simple Not Found Error will be returned, with the original response code (or overrideResponseCode if configured).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "error_response_rule": {
                          "block": {
                            "attributes": {
                              "match_response_codes": {
                                "description": "Valid values include:\n\n- A number between 400 and 599: For example 401 or 503, in which case the load balancer applies the policy if the error code exactly matches this value.\n- 5xx: Load Balancer will apply the policy if the backend service responds with any response code in the range of 500 to 599.\n- 4xx: Load Balancer will apply the policy if the backend service responds with any response code in the range of 400 to 499.\n\nValues must be unique within matchResponseCodes and across all errorResponseRules of CustomErrorResponsePolicy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "override_response_code": {
                                "description": "The HTTP status code returned with the response containing the custom error content.\nIf overrideResponseCode is not supplied, the same response code returned by the original backend bucket or backend service is returned to the client.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "path": {
                                "description": "The full path to a file within backendBucket . For example: /errors/defaultError.html\npath must start with a leading slash. path cannot have trailing slashes.\nIf the file is not available in backendBucket or the load balancer cannot reach the BackendBucket, a simple Not Found Error is returned to the client.\nThe value must be from 1 to 1024 characters",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Specifies rules for returning error responses.\nIn a given policy, if you specify rules for both a range of error codes as well as rules for specific error codes then rules with specific error codes have a higher priority.\nFor example, assume that you configure a rule for 401 (Un-authorized) code, and another for all 4 series error codes (4XX).\nIf the backend service returns a 401, then the rule for 401 will be applied. However if the backend service returns a 403, the rule for 4xx takes effect.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "customErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendService or BackendBucket responds with an error.\nIf a policy for an error code is not configured for the PathRule, a policy for the error code configured in pathMatcher.defaultCustomErrorResponsePolicy is applied. If one is not specified in pathMatcher.defaultCustomErrorResponsePolicy, the policy configured in UrlMap.defaultCustomErrorResponsePolicy takes effect.\nFor example, consider a UrlMap with the following configuration:\nUrlMap.defaultCustomErrorResponsePolicy are configured with policies for 5xx and 4xx errors\nA PathRule for /coming_soon/ is configured for the error code 404.\nIf the request is for www.myotherdomain.com and a 404 is encountered, the policy under UrlMap.defaultCustomErrorResponsePolicy takes effect. If a 404 response is encountered for the request www.example.com/current_events/, the pathMatcher's policy takes effect. If however, the request for www.example.com/coming_soon/ encounters a 404, the policy in PathRule.customErrorResponsePolicy takes effect. If any of the requests in this example encounter a 500 error code, the policy at UrlMap.defaultCustomErrorResponsePolicy takes effect.\ncustomErrorResponsePolicy is supported only for global external Application Load Balancers.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
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
                        "max_stream_duration": {
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
                                "description": "The BackendService resource being mirrored to.",
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
                                "description": "Specifies one or more conditions when this retry rule applies. Valid values are:\n\n* 5xx: Loadbalancer will attempt a retry if the backend service responds with\nany 5xx response code, or if the backend service does not respond at all,\nfor example: disconnects, reset, read timeout, connection failure, and refused\nstreams.\n* gateway-error: Similar to 5xx, but only applies to response codes\n502, 503 or 504.\n* connect-failure: Loadbalancer will retry on failures\nconnecting to backend services, for example due to connection timeouts.\n* retriable-4xx: Loadbalancer will retry for retriable 4xx response codes.\nCurrently the only retriable error supported is 409.\n* refused-stream: Loadbalancer will retry if the backend service resets the stream with a\nREFUSED_STREAM error code. This reset type indicates that it is safe to retry.\n* cancelled: Loadbalancer will retry if the gRPC status code in the response\nheader is set to cancelled\n* deadline-exceeded: Loadbalancer will retry if the\ngRPC status code in the response header is set to deadline-exceeded\n* resource-exhausted: Loadbalancer will retry if the gRPC status code in the response\nheader is set to resource-exhausted\n* unavailable: Loadbalancer will retry if\nthe gRPC status code in the response header is set to unavailable",
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
                                "description": "The default BackendService resource. Before\nforwarding the request to backendService, the loadbalancer applies any relevant\nheaderActions specified as part of this backendServiceWeight.",
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
                          "description": "If set to true, any accompanying query portion of the original URL is\nremoved prior to redirecting the request. If set to false, the query\nportion of the original URL is retained.\n This field is required to ensure an empty block is not set. The normal default value is false.",
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
                    "description": "The backend service resource to which traffic is\ndirected if this rule is matched. If routeAction is additionally specified,\nadvanced routing actions like URL Rewrites, etc. take effect prior to sending\nthe request to the backend. However, if service is specified, routeAction cannot\ncontain any weightedBackendService s. Conversely, if routeAction specifies any\nweightedBackendServices, service must not be specified. Only one of urlRedirect,\nservice or routeAction.weightedBackendService must be set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "custom_error_response_policy": {
                    "block": {
                      "attributes": {
                        "error_service": {
                          "description": "The full or partial URL to the BackendBucket resource that contains the custom error content. Examples are:\n\nhttps://www.googleapis.com/compute/v1/projects/project/global/backendBuckets/myBackendBucket\ncompute/v1/projects/project/global/backendBuckets/myBackendBucket\nglobal/backendBuckets/myBackendBucket\n\nIf errorService is not specified at lower levels like pathMatcher, pathRule and routeRule, an errorService specified at a higher level in the UrlMap will be used. If UrlMap.defaultCustomErrorResponsePolicy contains one or more errorResponseRules[], it must specify errorService.\nIf load balancer cannot reach the backendBucket, a simple Not Found Error will be returned, with the original response code (or overrideResponseCode if configured).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "error_response_rule": {
                          "block": {
                            "attributes": {
                              "match_response_codes": {
                                "description": "Valid values include:\n\n- A number between 400 and 599: For example 401 or 503, in which case the load balancer applies the policy if the error code exactly matches this value.\n- 5xx: Load Balancer will apply the policy if the backend service responds with any response code in the range of 500 to 599.\n- 4xx: Load Balancer will apply the policy if the backend service responds with any response code in the range of 400 to 499.\n\nValues must be unique within matchResponseCodes and across all errorResponseRules of CustomErrorResponsePolicy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "override_response_code": {
                                "description": "The HTTP status code returned with the response containing the custom error content.\nIf overrideResponseCode is not supplied, the same response code returned by the original backend bucket or backend service is returned to the client.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "path": {
                                "description": "The full path to a file within backendBucket . For example: /errors/defaultError.html\npath must start with a leading slash. path cannot have trailing slashes.\nIf the file is not available in backendBucket or the load balancer cannot reach the BackendBucket, a simple Not Found Error is returned to the client.\nThe value must be from 1 to 1024 characters",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Specifies rules for returning error responses.\nIn a given policy, if you specify rules for both a range of error codes as well as rules for specific error codes then rules with specific error codes have a higher priority.\nFor example, assume that you configure a rule for 401 (Un-authorized) code, and another for all 4 series error codes (4XX).\nIf the backend service returns a 401, then the rule for 401 will be applied. However if the backend service returns a 403, the rule for 4xx takes effect.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "customErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendService or BackendBucket responds with an error.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
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
                                  "description": "The header value must be an integer and its value must be in the range specified\nin rangeMatch. If the header does not contain an integer, number or is empty,\nthe match fails. For example for a range [-5, 0]   - -3 will match.  - 0 will\nnot match.  - 0.25 will not match.  - -3someString will not match.   Only one of\nexactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch\nmust be set.",
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
                                "description": "Specifies how individual filterLabel matches within the list of filterLabels\ncontribute towards the overall metadataFilter match. Supported values are:\n  - MATCH_ANY: At least one of the filterLabels must have a matching label in the\nprovided metadata.\n  - MATCH_ALL: All filterLabels must have matching labels in\nthe provided metadata. Possible values: [\"MATCH_ALL\", \"MATCH_ANY\"]",
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
                        "max_stream_duration": {
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
                                "description": "The BackendService resource being mirrored to.",
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
                                "description": "Specfies one or more conditions when this retry rule applies. Valid values are:\n\n* 5xx: Loadbalancer will attempt a retry if the backend service responds with\n  any 5xx response code, or if the backend service does not respond at all,\n  for example: disconnects, reset, read timeout, connection failure, and refused\n  streams.\n* gateway-error: Similar to 5xx, but only applies to response codes\n  502, 503 or 504.\n* connect-failure: Loadbalancer will retry on failures\n  connecting to backend services, for example due to connection timeouts.\n* retriable-4xx: Loadbalancer will retry for retriable 4xx response codes.\n  Currently the only retriable error supported is 409.\n* refused-stream: Loadbalancer will retry if the backend service resets the stream with a\n  REFUSED_STREAM error code. This reset type indicates that it is safe to retry.\n* cancelled: Loadbalancer will retry if the gRPC status code in the response\n  header is set to cancelled\n* deadline-exceeded: Loadbalancer will retry if the\n  gRPC status code in the response header is set to deadline-exceeded\n* resource-exhausted: Loadbalancer will retry if the gRPC status code in the response\n  header is set to resource-exhausted\n* unavailable: Loadbalancer will retry if the gRPC status code in\n  the response header is set to unavailable",
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
                                  "description": "Specifies a non-zero timeout per retry attempt.\nIf not specified, will use the timeout set in HttpRouteAction. If timeout in HttpRouteAction\nis not set, will use the largest timeout among all backend services associated with the route.",
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
                                "description": "The default BackendService resource. Before\nforwarding the request to backendService, the loadbalancer applies any relevant\nheaderActions specified as part of this backendServiceWeight.",
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
                          "description": "The host that will be used in the redirect response instead of the one that was\nsupplied in the request. The value must be between 1 and 255 characters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "https_redirect": {
                          "description": "If set to true, the URL scheme in the redirected request is set to https. If set\nto false, the URL scheme of the redirected request will remain the same as that\nof the request. This must only be set for UrlMaps used in TargetHttpProxys.\nSetting this true for TargetHttpsProxy is not permitted. Defaults to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "path_redirect": {
                          "description": "The path that will be used in the redirect response instead of the one that was\nsupplied in the request. Only one of pathRedirect or prefixRedirect must be\nspecified. The value must be between 1 and 1024 characters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "prefix_redirect": {
                          "description": "The prefix that replaces the prefixMatch specified in the HttpRouteRuleMatch,\nretaining the remaining portion of the URL before redirecting the request.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "redirect_response_code": {
                          "description": "The HTTP Status code to use for this RedirectAction. Supported values are:\n\n* MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301.\n\n* FOUND, which corresponds to 302.\n\n* SEE_OTHER which corresponds to 303.\n\n* TEMPORARY_REDIRECT, which corresponds to 307. In this case, the request method will be retained.\n\n* PERMANENT_REDIRECT, which corresponds to 308. In this case, the request method will be retained. Possible values: [\"FOUND\", \"MOVED_PERMANENTLY_DEFAULT\", \"PERMANENT_REDIRECT\", \"SEE_OTHER\", \"TEMPORARY_REDIRECT\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "strip_query": {
                          "description": "If set to true, any accompanying query portion of the original URL is removed\nprior to redirecting the request. If set to false, the query portion of the\noriginal URL is retained. Defaults to false.",
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
            "expected_output_url": {
              "description": "The expected output URL evaluated by the load balancer containing the scheme, host, path and query parameters.\n\nFor rules that forward requests to backends, the test passes only when expectedOutputUrl matches the request forwarded by the load balancer to backends. For rules with urlRewrite, the test verifies that the forwarded request matches hostRewrite and pathPrefixRewrite in the urlRewrite action. When service is specified, expectedOutputUrl's scheme is ignored.\n\nFor rules with urlRedirect, the test passes only if expectedOutputUrl matches the URL in the load balancer's redirect response. If urlRedirect specifies httpsRedirect, the test passes only if the scheme in expectedOutputUrl is also set to HTTPS. If urlRedirect specifies stripQuery, the test passes only if expectedOutputUrl does not contain any query parameters.\n\nexpectedOutputUrl is optional when service is specified.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expected_redirect_response_code": {
              "description": "For rules with urlRedirect, the test passes only if expectedRedirectResponseCode matches the HTTP status code in load balancer's redirect response.\n\nexpectedRedirectResponseCode cannot be set when service is set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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
              "description": "The backend service or backend bucket link that should be matched by this test.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "headers": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Header name.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Header value.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "HTTP headers for this request.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "The list of expected URL mapping tests. Request to update this UrlMap will\nsucceed only if all of the test cases pass. You can specify a maximum of 100\ntests per UrlMap.",
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

func GoogleComputeUrlMapSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeUrlMap), &result)
	return &result
}
