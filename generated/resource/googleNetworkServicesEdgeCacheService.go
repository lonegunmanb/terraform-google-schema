package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkServicesEdgeCacheService = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A human-readable description of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disable_http2": {
        "description": "Disables HTTP/2.\n\nHTTP/2 (h2) is enabled by default and recommended for performance. HTTP/2 improves connection re-use and reduces connection setup overhead by sending multiple streams over the same connection.\n\nSome legacy HTTP clients may have issues with HTTP/2 connections due to broken HTTP/2 implementations. Setting this to true will prevent HTTP/2 from being advertised and negotiated.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_quic": {
        "computed": true,
        "description": "HTTP/3 (IETF QUIC) and Google QUIC are enabled by default.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "edge_security_policy": {
        "description": "Resource URL that points at the Cloud Armor edge security policy that is applied on each request against the EdgeCacheService.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "edge_ssl_certificates": {
        "description": "URLs to sslCertificate resources that are used to authenticate connections between users and the EdgeCacheService.\n\nNote that only \"global\" certificates with a \"scope\" of \"EDGE_CACHE\" can be attached to an EdgeCacheService.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv4_addresses": {
        "computed": true,
        "description": "The IPv4 addresses associated with this service. Addresses are static for the lifetime of the service.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "ipv6_addresses": {
        "computed": true,
        "description": "The IPv6 addresses associated with this service. Addresses are static for the lifetime of the service.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "labels": {
        "description": "Set of label tags associated with the EdgeCache resource.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "Name of the resource; provided by the client when the resource is created.\nThe name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,\nand all following characters must be a dash, underscore, letter or digit.",
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
      "require_tls": {
        "computed": true,
        "description": "Require TLS (HTTPS) for all clients connecting to this service.\n\nClients who connect over HTTP (port 80) will receive a HTTP 301 to the same URL over HTTPS (port 443).\nYou must have at least one (1) edgeSslCertificate specified to enable this.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ssl_policy": {
        "description": "URL of the SslPolicy resource that will be associated with the EdgeCacheService.\n\nIf not set, the EdgeCacheService has no SSL policy configured, and will default to the \"COMPATIBLE\" policy.",
        "description_kind": "plain",
        "optional": true,
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
      }
    },
    "block_types": {
      "log_config": {
        "block": {
          "attributes": {
            "enable": {
              "computed": true,
              "description": "Specifies whether to enable logging for traffic served by this service.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "sample_rate": {
              "description": "Configures the sampling rate of requests, where 1.0 means all logged requests are reported and 0.0 means no logged requests are reported. The default value is 1.0, and the value of the field must be in [0, 1].\n\nThis field can only be specified if logging is enabled for this service.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Specifies the logging options for the traffic served by this service. If logging is enabled, logs will be exported to Cloud Logging.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "routing": {
        "block": {
          "block_types": {
            "host_rule": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "A human-readable description of the hostRule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "hosts": {
                    "description": "The list of host patterns to match.\n\nHost patterns must be valid hostnames. Ports are not allowed. Wildcard hosts are supported in the suffix or prefix form. * matches any string of ([a-z0-9-.]*). It does not match the empty string.\n\nWhen multiple hosts are specified, hosts are matched in the following priority:\n\n  1. Exact domain names: ''www.foo.com''.\n  2. Suffix domain wildcards: ''*.foo.com'' or ''*-bar.foo.com''.\n  3. Prefix domain wildcards: ''foo.*'' or ''foo-*''.\n  4. Special wildcard ''*'' matching any domain.\n\n  Notes:\n\n    The wildcard will not match the empty string. e.g. ''*-bar.foo.com'' will match ''baz-bar.foo.com'' but not ''-bar.foo.com''. The longest wildcards match first. Only a single host in the entire service can match on ''*''. A domain must be unique across all configured hosts within a service.\n\n    Hosts are matched against the HTTP Host header, or for HTTP/2 and HTTP/3, the \":authority\" header, from the incoming request.\n\n    You may specify up to 10 hosts.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "path_matcher": {
                    "description": "The name of the pathMatcher associated with this hostRule.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The list of hostRules to match against. These rules define which hostnames the EdgeCacheService will match against, and which route configurations apply.",
                "description_kind": "plain"
              },
              "max_items": 50,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "path_matcher": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "A human-readable description of the resource.",
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
                  "route_rule": {
                    "block": {
                      "attributes": {
                        "description": {
                          "description": "A human-readable description of the routeRule.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "origin": {
                          "description": "The Origin resource that requests to this route should fetch from when a matching response is not in cache. Origins can be defined as short names (\"my-origin\") or fully-qualified resource URLs - e.g. \"networkservices.googleapis.com/projects/my-project/global/edgecacheorigins/my-origin\"\n\nOnly one of origin or urlRedirect can be set.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "priority": {
                          "description": "The priority of this route rule, where 1 is the highest priority.\n\nYou cannot configure two or more routeRules with the same priority. Priority for each rule must be set to a number between 1 and 999 inclusive.\n\nPriority numbers can have gaps, which enable you to add or remove rules in the future without affecting the rest of the rules. For example, 1, 2, 3, 4, 5, 9, 12, 16 is a valid series of priority numbers\nto which you could add rules numbered from 6 to 8, 10 to 11, and 13 to 15 in the future without any impact on existing rules.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "header_action": {
                          "block": {
                            "block_types": {
                              "request_header_to_add": {
                                "block": {
                                  "attributes": {
                                    "header_name": {
                                      "description": "The name of the header to add.",
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
                                      "computed": true,
                                      "description": "Whether to replace all existing headers with the same name.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description": "Describes a header to add.",
                                  "description_kind": "plain"
                                },
                                "max_items": 25,
                                "nesting_mode": "list"
                              },
                              "request_header_to_remove": {
                                "block": {
                                  "attributes": {
                                    "header_name": {
                                      "description": "The name of the header to remove.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "A list of header names for headers that need to be removed from the request prior to forwarding the request to the origin.",
                                  "description_kind": "plain"
                                },
                                "max_items": 25,
                                "nesting_mode": "list"
                              },
                              "response_header_to_add": {
                                "block": {
                                  "attributes": {
                                    "header_name": {
                                      "description": "The name of the header to add.",
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
                                      "computed": true,
                                      "description": "Whether to replace all existing headers with the same name.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description": "Headers to add to the response prior to sending it back to the client.\n\nResponse headers are only sent to the client, and do not have an effect on the cache serving the response.",
                                  "description_kind": "plain"
                                },
                                "max_items": 25,
                                "nesting_mode": "list"
                              },
                              "response_header_to_remove": {
                                "block": {
                                  "attributes": {
                                    "header_name": {
                                      "description": "Headers to remove from the response prior to sending it back to the client.\n\nResponse headers are only sent to the client, and do not have an effect on the cache serving the response.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "A list of header names for headers that need to be removed from the request prior to forwarding the request to the origin.",
                                  "description_kind": "plain"
                                },
                                "max_items": 25,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "The header actions, including adding \u0026 removing headers, for requests that match this route.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "match_rule": {
                          "block": {
                            "attributes": {
                              "full_path_match": {
                                "description": "For satisfying the matchRule condition, the path of the request must exactly match the value specified in fullPathMatch after removing any query parameters and anchor that may be part of the original URL.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "ignore_case": {
                                "computed": true,
                                "description": "Specifies that prefixMatch and fullPathMatch matches are case sensitive.",
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
                                "description": "For satisfying the matchRule condition, the request's path must begin with the specified prefixMatch. prefixMatch must begin with a /.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "header_match": {
                                "block": {
                                  "attributes": {
                                    "exact_match": {
                                      "description": "The value of the header should exactly match contents of exactMatch.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "header_name": {
                                      "description": "The header name to match on.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "invert_match": {
                                      "computed": true,
                                      "description": "If set to false (default), the headerMatch is considered a match if the match criteria above are met.\nIf set to true, the headerMatch is considered a match if the match criteria above are NOT met.",
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
                                      "description": "A header with the contents of headerName must exist. The match takes place whether or not the request's header has a value.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "suffix_match": {
                                      "description": "The value of the header must end with the contents of suffixMatch.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Specifies a list of header match criteria, all of which must match corresponding headers in the request.",
                                  "description_kind": "plain"
                                },
                                "max_items": 3,
                                "nesting_mode": "list"
                              },
                              "query_parameter_match": {
                                "block": {
                                  "attributes": {
                                    "exact_match": {
                                      "description": "The queryParameterMatch matches if the value of the parameter exactly matches the contents of exactMatch.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "name": {
                                      "description": "The name of the query parameter to match. The query parameter must exist in the request, in the absence of which the request match fails.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "present_match": {
                                      "description": "Specifies that the queryParameterMatch matches if the request contains the query parameter, irrespective of whether the parameter has a value or not.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description": "Specifies a list of query parameter match criteria, all of which must match corresponding query parameters in the request.",
                                  "description_kind": "plain"
                                },
                                "max_items": 5,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "The list of criteria for matching attributes of a request to this routeRule. This list has OR semantics: the request matches this routeRule when any of the matchRules are satisfied. However predicates\nwithin a given matchRule have AND semantics. All predicates within a matchRule must match for the request to match the rule.",
                            "description_kind": "plain"
                          },
                          "max_items": 5,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "route_action": {
                          "block": {
                            "attributes": {
                              "compression_mode": {
                                "description": "Setting the compression mode to automatic enables dynamic compression for every eligible response.\n\nWhen dynamic compression is enabled, it is recommended to also set a cache policy to maximize efficiency. Possible values: [\"DISABLED\", \"AUTOMATIC\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "cdn_policy": {
                                "block": {
                                  "attributes": {
                                    "cache_mode": {
                                      "computed": true,
                                      "description": "Cache modes allow users to control the behaviour of the cache, what content it should cache automatically, whether to respect origin headers, or whether to unconditionally cache all responses.\n\nFor all cache modes, Cache-Control headers will be passed to the client. Use clientTtl to override what is sent to the client. Possible values: [\"CACHE_ALL_STATIC\", \"USE_ORIGIN_HEADERS\", \"FORCE_CACHE_ALL\", \"BYPASS_CACHE\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "client_ttl": {
                                      "description": "Specifies a separate client (e.g. browser client) TTL, separate from the TTL used by the edge caches. Leaving this empty will use the same cache TTL for both the CDN and the client-facing response.\n\n- The TTL must be \u003e 0 and \u003c= 86400s (1 day)\n- The clientTtl cannot be larger than the defaultTtl (if set)\n- Fractions of a second are not allowed.\n\nOmit this field to use the defaultTtl, or the max-age set by the origin, as the client-facing TTL.\n\nWhen the cache mode is set to \"USE_ORIGIN_HEADERS\" or \"BYPASS_CACHE\", you must omit this field.\nA duration in seconds terminated by 's'. Example: \"3s\".",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "default_ttl": {
                                      "computed": true,
                                      "description": "Specifies the default TTL for cached content served by this origin for responses that do not have an existing valid TTL (max-age or s-max-age).\n\nDefaults to 3600s (1 hour).\n\n- The TTL must be \u003e= 0 and \u003c= 31,536,000 seconds (1 year)\n- Setting a TTL of \"0\" means \"always revalidate\" (equivalent to must-revalidate)\n- The value of defaultTTL cannot be set to a value greater than that of maxTTL.\n- Fractions of a second are not allowed.\n- When the cacheMode is set to FORCE_CACHE_ALL, the defaultTTL will overwrite the TTL set in all responses.\n\nNote that infrequently accessed objects may be evicted from the cache before the defined TTL. Objects that expire will be revalidated with the origin.\n\nWhen the cache mode is set to \"USE_ORIGIN_HEADERS\" or \"BYPASS_CACHE\", you must omit this field.\n\nA duration in seconds terminated by 's'. Example: \"3s\".",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "max_ttl": {
                                      "computed": true,
                                      "description": "Specifies the maximum allowed TTL for cached content served by this origin.\n\nDefaults to 86400s (1 day).\n\nCache directives that attempt to set a max-age or s-maxage higher than this, or an Expires header more than maxTtl seconds in the future will be capped at the value of maxTTL, as if it were the value of an s-maxage Cache-Control directive.\n\n- The TTL must be \u003e= 0 and \u003c= 31,536,000 seconds (1 year)\n- Setting a TTL of \"0\" means \"always revalidate\"\n- The value of maxTtl must be equal to or greater than defaultTtl.\n- Fractions of a second are not allowed.\n\nWhen the cache mode is set to \"USE_ORIGIN_HEADERS\", \"FORCE_CACHE_ALL\", or \"BYPASS_CACHE\", you must omit this field.\n\nA duration in seconds terminated by 's'. Example: \"3s\".",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "negative_caching": {
                                      "description": "Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects. This can reduce the load on your origin and improve end-user experience by reducing response latency.\n\nBy default, the CDNPolicy will apply the following default TTLs to these status codes:\n\n- HTTP 300 (Multiple Choice), 301, 308 (Permanent Redirects): 10m\n- HTTP 404 (Not Found), 410 (Gone), 451 (Unavailable For Legal Reasons): 120s\n- HTTP 405 (Method Not Found), 414 (URI Too Long), 501 (Not Implemented): 60s\n\nThese defaults can be overridden in negativeCachingPolicy",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "negative_caching_policy": {
                                      "description": "Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.\n\n- Omitting the policy and leaving negativeCaching enabled will use the default TTLs for each status code, defined in negativeCaching.\n- TTLs must be \u003e= 0 (where 0 is \"always revalidate\") and \u003c= 86400s (1 day)\n\nNote that when specifying an explicit negativeCachingPolicy, you should take care to specify a cache TTL for all response codes that you wish to cache. The CDNPolicy will not apply any default negative caching when a policy exists.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    },
                                    "signed_request_keyset": {
                                      "computed": true,
                                      "description": "The EdgeCacheKeyset containing the set of public keys used to validate signed requests at the edge.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "signed_request_maximum_expiration_ttl": {
                                      "description": "Limit how far into the future the expiration time of a signed request may be.\n\nWhen set, a signed request is rejected if its expiration time is later than now + signedRequestMaximumExpirationTtl, where now is the time at which the signed request is first handled by the CDN.\n\n- The TTL must be \u003e 0.\n- Fractions of a second are not allowed.\n\nBy default, signedRequestMaximumExpirationTtl is not set and the expiration time of a signed request may be arbitrarily far into future.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "signed_request_mode": {
                                      "computed": true,
                                      "description": "Whether to enforce signed requests. The default value is DISABLED, which means all content is public, and does not authorize access.\n\nYou must also set a signedRequestKeyset to enable signed requests.\n\nWhen set to REQUIRE_SIGNATURES, all matching requests will have their signature validated. Requests that were not signed with the corresponding private key, or that are otherwise invalid (expired, do not match the signature, IP address, or header) will be rejected with a HTTP 403 and (if enabled) logged. Possible values: [\"DISABLED\", \"REQUIRE_SIGNATURES\", \"REQUIRE_TOKENS\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "add_signatures": {
                                      "block": {
                                        "attributes": {
                                          "actions": {
                                            "description": "The actions to take to add signatures to responses. Possible values: [\"GENERATE_COOKIE\", \"GENERATE_TOKEN_HLS_COOKIELESS\", \"PROPAGATE_TOKEN_HLS_COOKIELESS\"]",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "copied_parameters": {
                                            "description": "The parameters to copy from the verified token to the generated token.\n\nOnly the following parameters may be copied:\n\n  * 'PathGlobs'\n  * 'paths'\n  * 'acl'\n  * 'URLPrefix'\n  * 'IPRanges'\n  * 'SessionID'\n  * 'id'\n  * 'Data'\n  * 'data'\n  * 'payload'\n  * 'Headers'\n\nYou may specify up to 6 parameters to copy.  A given parameter is be copied only if the parameter exists in the verified token.  Parameter names are matched exactly as specified.  The order of the parameters does not matter.  Duplicates are not allowed.\n\nThis field may only be specified when the GENERATE_COOKIE or GENERATE_TOKEN_HLS_COOKIELESS actions are specified.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "keyset": {
                                            "description": "The keyset to use for signature generation.\n\nThe following are both valid paths to an EdgeCacheKeyset resource:\n\n  * 'projects/project/locations/global/edgeCacheKeysets/yourKeyset'\n  * 'yourKeyset'\n\nThis must be specified when the GENERATE_COOKIE or GENERATE_TOKEN_HLS_COOKIELESS actions are specified.  This field may not be specified otherwise.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "token_query_parameter": {
                                            "description": "The query parameter in which to put the generated token.\n\nIf not specified, defaults to 'edge-cache-token'.\n\nIf specified, the name must be 1-64 characters long and match the regular expression '[a-zA-Z]([a-zA-Z0-9_-])*' which means the first character must be a letter, and all following characters must be a dash, underscore, letter or digit.\n\nThis field may only be set when the GENERATE_TOKEN_HLS_COOKIELESS or PROPAGATE_TOKEN_HLS_COOKIELESS actions are specified.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "token_ttl": {
                                            "description": "The duration the token is valid starting from the moment the token is first generated.\n\nDefaults to '86400s' (1 day).\n\nThe TTL must be \u003e= 0 and \u003c= 604,800 seconds (1 week).\n\nThis field may only be specified when the GENERATE_COOKIE or GENERATE_TOKEN_HLS_COOKIELESS actions are specified.\n\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Enable signature generation or propagation on this route.\n\nThis field may only be specified when signedRequestMode is set to REQUIRE_TOKENS.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "cache_key_policy": {
                                      "block": {
                                        "attributes": {
                                          "exclude_host": {
                                            "computed": true,
                                            "description": "If true, requests to different hosts will be cached separately.\n\nNote: this should only be enabled if hosts share the same origin and content. Removing the host from the cache key may inadvertently result in different objects being cached than intended, depending on which route the first user matched.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "exclude_query_string": {
                                            "description": "If true, exclude query string parameters from the cache key\n\nIf false (the default), include the query string parameters in\nthe cache key according to includeQueryParameters and\nexcludeQueryParameters. If neither includeQueryParameters nor\nexcludeQueryParameters is set, the entire query string will be\nincluded.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "excluded_query_parameters": {
                                            "description": "Names of query string parameters to exclude from cache keys. All other parameters will be included.\n\nEither specify includedQueryParameters or excludedQueryParameters, not both. '\u0026' and '=' will be percent encoded and not treated as delimiters.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "include_protocol": {
                                            "computed": true,
                                            "description": "If true, http and https requests will be cached separately.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "included_cookie_names": {
                                            "description": "Names of Cookies to include in cache keys.  The cookie name and cookie value of each cookie named will be used as part of the cache key.\n\nCookie names:\n  - must be valid RFC 6265 \"cookie-name\" tokens\n  - are case sensitive\n  - cannot start with \"Edge-Cache-\" (case insensitive)\n\n  Note that specifying several cookies, and/or cookies that have a large range of values (e.g., per-user) will dramatically impact the cache hit rate, and may result in a higher eviction rate and reduced performance.\n\n  You may specify up to three cookie names.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "included_header_names": {
                                            "description": "Names of HTTP request headers to include in cache keys. The value of the header field will be used as part of the cache key.\n\n- Header names must be valid HTTP RFC 7230 header field values.\n- Header field names are case insensitive\n- To include the HTTP method, use \":method\"\n\nNote that specifying several headers, and/or headers that have a large range of values (e.g. per-user) will dramatically impact the cache hit rate, and may result in a higher eviction rate and reduced performance.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "included_query_parameters": {
                                            "description": "Names of query string parameters to include in cache keys. All other parameters will be excluded.\n\nEither specify includedQueryParameters or excludedQueryParameters, not both. '\u0026' and '=' will be percent encoded and not treated as delimiters.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "Defines the request parameters that contribute to the cache key.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "signed_token_options": {
                                      "block": {
                                        "attributes": {
                                          "allowed_signature_algorithms": {
                                            "description": "The allowed signature algorithms to use.\n\nDefaults to using only ED25519.\n\nYou may specify up to 3 signature algorithms to use. Possible values: [\"ED25519\", \"HMAC_SHA_256\", \"HMAC_SHA1\"]",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "token_query_parameter": {
                                            "description": "The query parameter in which to find the token.\n\nThe name must be 1-64 characters long and match the regular expression '[a-zA-Z]([a-zA-Z0-9_-])*' which means the first character must be a letter, and all following characters must be a dash, underscore, letter or digit.\n\nDefaults to 'edge-cache-token'.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Additional options for signed tokens.\n\nsignedTokenOptions may only be specified when signedRequestMode is REQUIRE_TOKENS.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "The policy to use for defining caching and signed request behaviour for requests that match this route.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "cors_policy": {
                                "block": {
                                  "attributes": {
                                    "allow_credentials": {
                                      "description": "In response to a preflight request, setting this to true indicates that the actual request can include user credentials.\n\nThis translates to the Access-Control-Allow-Credentials response header.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "allow_headers": {
                                      "description": "Specifies the content for the Access-Control-Allow-Headers response header.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "allow_methods": {
                                      "description": "Specifies the content for the Access-Control-Allow-Methods response header.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "allow_origins": {
                                      "description": "Specifies the list of origins that will be allowed to do CORS requests.\n\nThis translates to the Access-Control-Allow-Origin response header.",
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
                                      "description": "Specifies the content for the Access-Control-Allow-Headers response header.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "max_age": {
                                      "description": "Specifies how long results of a preflight request can be cached by a client in seconds. Note that many browser clients enforce a maximum TTL of 600s (10 minutes).\n\n- Setting the value to -1 forces a pre-flight check for all requests (not recommended)\n- A maximum TTL of 86400s can be set, but note that (as above) some clients may force pre-flight checks at a more regular interval.\n- This translates to the Access-Control-Max-Age header.\n\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "CORSPolicy defines Cross-Origin-Resource-Sharing configuration, including which CORS response headers will be set.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "url_rewrite": {
                                "block": {
                                  "attributes": {
                                    "host_rewrite": {
                                      "description": "Prior to forwarding the request to the selected origin, the request's host header is replaced with contents of hostRewrite.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "path_prefix_rewrite": {
                                      "description": "Prior to forwarding the request to the selected origin, the matching portion of the request's path is replaced by pathPrefixRewrite.",
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
                                  "description": "The URL rewrite configuration for requests that match this route.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "In response to a matching path, the routeAction performs advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request to the selected origin.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "route_methods": {
                          "block": {
                            "attributes": {
                              "allowed_methods": {
                                "description": "The non-empty set of HTTP methods that are allowed for this route.\n\nAny combination of \"GET\", \"HEAD\", \"OPTIONS\", \"PUT\", \"POST\", \"DELETE\", and \"PATCH\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "Allow overriding the set of methods that are allowed for this route.\nWhen not set, Media CDN allows only \"GET\", \"HEAD\", and \"OPTIONS\".",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "url_redirect": {
                          "block": {
                            "attributes": {
                              "host_redirect": {
                                "description": "The host that will be used in the redirect response instead of the one that was supplied in the request.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "https_redirect": {
                                "computed": true,
                                "description": "If set to true, the URL scheme in the redirected request is set to https. If set to false, the URL scheme of the redirected request will remain the same as that of the request.\n\nThis can only be set if there is at least one (1) edgeSslCertificate set on the service.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "path_redirect": {
                                "description": "The path that will be used in the redirect response instead of the one that was supplied in the request.\n\npathRedirect cannot be supplied together with prefixRedirect. Supply one alone or neither. If neither is supplied, the path of the original request will be used for the redirect.\n\nThe path value must be between 1 and 1024 characters.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "prefix_redirect": {
                                "description": "The prefix that replaces the prefixMatch specified in the routeRule, retaining the remaining portion of the URL before redirecting the request.\n\nprefixRedirect cannot be supplied together with pathRedirect. Supply one alone or neither. If neither is supplied, the path of the original request will be used for the redirect.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_response_code": {
                                "computed": true,
                                "description": "The HTTP Status code to use for this RedirectAction.\n\nThe supported values are:\n\n- 'MOVED_PERMANENTLY_DEFAULT', which is the default value and corresponds to 301.\n- 'FOUND', which corresponds to 302.\n- 'SEE_OTHER' which corresponds to 303.\n- 'TEMPORARY_REDIRECT', which corresponds to 307. in this case, the request method will be retained.\n- 'PERMANENT_REDIRECT', which corresponds to 308. in this case, the request method will be retained. Possible values: [\"MOVED_PERMANENTLY_DEFAULT\", \"FOUND\", \"SEE_OTHER\", \"TEMPORARY_REDIRECT\", \"PERMANENT_REDIRECT\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "strip_query": {
                                "computed": true,
                                "description": "If set to true, any accompanying query portion of the original URL is removed prior to redirecting the request. If set to false, the query portion of the original URL is retained.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "The URL redirect configuration for requests that match this route.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The routeRules to match against. routeRules support advanced routing behaviour, and can match on paths, headers and query parameters, as well as status codes and HTTP methods.",
                      "description_kind": "plain"
                    },
                    "max_items": 200,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The list of pathMatchers referenced via name by hostRules. PathMatcher is used to match the path portion of the URL when a HostRule matches the URL's host portion.",
                "description_kind": "plain"
              },
              "max_items": 50,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Defines how requests are routed, modified, cached and/or which origin content is filled from.",
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GoogleNetworkServicesEdgeCacheServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkServicesEdgeCacheService), &result)
	return &result
}
