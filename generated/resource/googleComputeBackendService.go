package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeBackendService = `{
  "block": {
    "attributes": {
      "affinity_cookie_ttl_sec": {
        "description": "Lifetime of cookies in seconds if session_affinity is\nGENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts\nonly until the end of the browser session (or equivalent). The\nmaximum allowed value for TTL is one day.\n\nWhen the load balancing scheme is INTERNAL, this field is not used.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "compression_mode": {
        "description": "Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header. Possible values: [\"AUTOMATIC\", \"DISABLED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_draining_timeout_sec": {
        "description": "Time for which instance will be drained (not accept new\nconnections, but still work to finish started).",
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
      "custom_request_headers": {
        "description": "Headers that the HTTP/S load balancer should add to proxied\nrequests.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "custom_response_headers": {
        "description": "Headers that the HTTP/S load balancer should add to proxied\nresponses.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "description": {
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "edge_security_policy": {
        "description": "The resource URL for the edge security policy associated with this backend service.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_cdn": {
        "description": "If true, enable Cloud CDN for this BackendService.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "external_managed_migration_state": {
        "description": "Specifies the canary migration state. Possible values are PREPARE, TEST_BY_PERCENTAGE, and\nTEST_ALL_TRAFFIC.\n\nTo begin the migration from EXTERNAL to EXTERNAL_MANAGED, the state must be changed to\nPREPARE. The state must be changed to TEST_ALL_TRAFFIC before the loadBalancingScheme can be\nchanged to EXTERNAL_MANAGED. Optionally, the TEST_BY_PERCENTAGE state can be used to migrate\ntraffic by percentage using externalManagedMigrationTestingPercentage.\n\nRolling back a migration requires the states to be set in reverse order. So changing the\nscheme from EXTERNAL_MANAGED to EXTERNAL requires the state to be set to TEST_ALL_TRAFFIC at\nthe same time. Optionally, the TEST_BY_PERCENTAGE state can be used to migrate some traffic\nback to EXTERNAL or PREPARE can be used to migrate all traffic back to EXTERNAL. Possible values: [\"PREPARE\", \"TEST_BY_PERCENTAGE\", \"TEST_ALL_TRAFFIC\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "external_managed_migration_testing_percentage": {
        "description": "Determines the fraction of requests that should be processed by the Global external\nApplication Load Balancer.\n\nThe value of this field must be in the range [0, 100].\n\nSession affinity options will slightly affect this routing behavior, for more details,\nsee: Session Affinity.\n\nThis value can only be set if the loadBalancingScheme in the backend service is set to\nEXTERNAL (when using the Classic ALB) and the migration state is TEST_BY_PERCENTAGE.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource. A hash of the contents stored in this\nobject. This field is used in optimistic locking.",
        "description_kind": "plain",
        "type": "string"
      },
      "generated_id": {
        "computed": true,
        "description": "The unique identifier for the resource. This identifier is defined by the server.",
        "description_kind": "plain",
        "type": "number"
      },
      "health_checks": {
        "description": "The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource\nfor health checking this BackendService. Currently at most one health\ncheck can be specified.\n\nA health check must be specified unless the backend service uses an internet\nor serverless NEG as a backend.\n\nFor internal load balancing, a URL to a HealthCheck resource must be specified instead.",
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
      "ip_address_selection_policy": {
        "description": "Specifies preference of traffic to the backend (from the proxy and from the client for proxyless gRPC). Possible values: [\"IPV4_ONLY\", \"PREFER_IPV6\", \"IPV6_ONLY\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancing_scheme": {
        "description": "Indicates whether the backend service will be used with internal or\nexternal load balancing. A backend service created for one type of\nload balancing cannot be used with the other. For more information, refer to\n[Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service). Default value: \"EXTERNAL\" Possible values: [\"EXTERNAL\", \"INTERNAL_SELF_MANAGED\", \"INTERNAL_MANAGED\", \"EXTERNAL_MANAGED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "locality_lb_policy": {
        "description": "The load balancing algorithm used within the scope of the locality.\nThe possible values are:\n\n* 'ROUND_ROBIN': This is a simple policy in which each healthy backend\n                 is selected in round robin order.\n\n* 'LEAST_REQUEST': An O(1) algorithm which selects two random healthy\n                   hosts and picks the host which has fewer active requests.\n\n* 'RING_HASH': The ring/modulo hash load balancer implements consistent\n               hashing to backends. The algorithm has the property that the\n               addition/removal of a host from a set of N hosts only affects\n               1/N of the requests.\n\n* 'RANDOM': The load balancer selects a random healthy host.\n\n* 'ORIGINAL_DESTINATION': Backend host is selected based on the client\n                          connection metadata, i.e., connections are opened\n                          to the same address as the destination address of\n                          the incoming connection before the connection\n                          was redirected to the load balancer.\n\n* 'MAGLEV': used as a drop in replacement for the ring hash load balancer.\n            Maglev is not as stable as ring hash but has faster table lookup\n            build times and host selection times. For more information about\n            Maglev, refer to https://ai.google/research/pubs/pub44824\n\n* 'WEIGHTED_MAGLEV': Per-instance weighted Load Balancing via health check\n                     reported weights. Only applicable to loadBalancingScheme\n                     EXTERNAL. If set, the Backend Service must\n                     configure a non legacy HTTP-based Health Check, and\n                     health check replies are expected to contain\n                     non-standard HTTP response header field\n                     X-Load-Balancing-Endpoint-Weight to specify the\n                     per-instance weights. If set, Load Balancing is weight\n                     based on the per-instance weights reported in the last\n                     processed health check replies, as long as every\n                     instance either reported a valid weight or had\n                     UNAVAILABLE_WEIGHT. Otherwise, Load Balancing remains\n                     equal-weight.\n\n* 'WEIGHTED_ROUND_ROBIN': Per-endpoint weighted round-robin Load Balancing using weights computed\n                          from Backend reported Custom Metrics. If set, the Backend Service\n                          responses are expected to contain non-standard HTTP response header field\n                          X-Endpoint-Load-Metrics. The reported metrics\n                          to use for computing the weights are specified via the\n                          backends[].customMetrics fields.\n\nlocality_lb_policy is applicable to either:\n\n* A regional backend service with the service_protocol set to HTTP, HTTPS, HTTP2 or H2C,\n  and loadBalancingScheme set to INTERNAL_MANAGED.\n* A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.\n* A regional backend service with loadBalancingScheme set to EXTERNAL (External Network\n  Load Balancing). Only MAGLEV and WEIGHTED_MAGLEV values are possible for External\n  Network Load Balancing. The default is MAGLEV.\n\nIf session_affinity is not NONE, and locality_lb_policy is not set to MAGLEV, WEIGHTED_MAGLEV,\nor RING_HASH, session affinity settings will not take effect.\n\nOnly ROUND_ROBIN and RING_HASH are supported when the backend service is referenced\nby a URL map that is bound to target gRPC proxy that has validate_for_proxyless\nfield set to true. Possible values: [\"ROUND_ROBIN\", \"LEAST_REQUEST\", \"RING_HASH\", \"RANDOM\", \"ORIGINAL_DESTINATION\", \"MAGLEV\", \"WEIGHTED_MAGLEV\", \"WEIGHTED_ROUND_ROBIN\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port_name": {
        "computed": true,
        "description": "Name of backend port. The same name should appear in the instance\ngroups referenced by this service. Required when the load balancing\nscheme is EXTERNAL.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
        "computed": true,
        "description": "The protocol this BackendService uses to communicate with backends.\nThe default is HTTP. Possible values are HTTP, HTTPS, HTTP2, H2C, TCP, SSL, UDP\nor GRPC. Refer to the documentation for the load balancers or for Traffic Director\nfor more information. Must be set to GRPC when the backend service is referenced\nby a URL map that is bound to target gRPC proxy. Possible values: [\"HTTP\", \"HTTPS\", \"HTTP2\", \"TCP\", \"SSL\", \"UDP\", \"GRPC\", \"UNSPECIFIED\", \"H2C\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_policy": {
        "description": "The security policy associated with this backend service.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_lb_policy": {
        "description": "URL to networkservices.ServiceLbPolicy resource.\nCan only be set if load balancing scheme is EXTERNAL, EXTERNAL_MANAGED, INTERNAL_MANAGED or INTERNAL_SELF_MANAGED and the scope is global.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "session_affinity": {
        "computed": true,
        "description": "Type of session affinity to use. The default is NONE. Session affinity is\nnot applicable if the protocol is UDP. Possible values: [\"NONE\", \"CLIENT_IP\", \"CLIENT_IP_PORT_PROTO\", \"CLIENT_IP_PROTO\", \"GENERATED_COOKIE\", \"HEADER_FIELD\", \"HTTP_COOKIE\", \"STRONG_COOKIE_AFFINITY\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "timeout_sec": {
        "computed": true,
        "description": "The backend service timeout has a different meaning depending on the type of load balancer.\nFor more information see, [Backend service settings](https://cloud.google.com/compute/docs/reference/rest/v1/backendServices).\nThe default is 30 seconds.\nThe full range of timeout values allowed goes from 1 through 2,147,483,647 seconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "backend": {
        "block": {
          "attributes": {
            "balancing_mode": {
              "description": "Specifies the balancing mode for this backend.\n\nFor global HTTP(S) or TCP/SSL load balancing, the default is\nUTILIZATION. Valid values are UTILIZATION, RATE (for HTTP(S)),\nCUSTOM_METRICS (for HTTP(s)) and CONNECTION (for TCP/SSL).\n\nSee the [Backend Services Overview](https://cloud.google.com/load-balancing/docs/backend-service#balancing-mode)\nfor an explanation of load balancing modes. Default value: \"UTILIZATION\" Possible values: [\"UTILIZATION\", \"RATE\", \"CONNECTION\", \"CUSTOM_METRICS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "capacity_scaler": {
              "description": "A multiplier applied to the group's maximum servicing capacity\n(based on UTILIZATION, RATE or CONNECTION).\n\nDefault value is 1, which means the group will serve up to 100%\nof its configured capacity (depending on balancingMode). A\nsetting of 0 means the group is completely drained, offering\n0% of its available Capacity. Valid range is [0.0,1.0].",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "description": {
              "description": "An optional description of this resource.\nProvide this property when you create the resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "group": {
              "description": "The fully-qualified URL of an Instance Group or Network Endpoint\nGroup resource. In case of instance group this defines the list\nof instances that serve traffic. Member virtual machine\ninstances from each instance group must live in the same zone as\nthe instance group itself. No two backends in a backend service\nare allowed to use same Instance Group resource.\n\nFor Network Endpoint Groups this defines list of endpoints. All\nendpoints of Network Endpoint Group must be hosted on instances\nlocated in the same zone as the Network Endpoint Group.\n\nBackend services cannot mix Instance Group and\nNetwork Endpoint Group backends.\n\nNote that you must specify an Instance Group or Network Endpoint\nGroup resource using the fully-qualified URL, rather than a\npartial URL.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "max_connections": {
              "computed": true,
              "description": "The max number of simultaneous connections for the group. Can\nbe used with either CONNECTION or UTILIZATION balancing modes.\n\nFor CONNECTION mode, either maxConnections or one\nof maxConnectionsPerInstance or maxConnectionsPerEndpoint,\nas appropriate for group type, must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_connections_per_endpoint": {
              "computed": true,
              "description": "The max number of simultaneous connections that a single backend\nnetwork endpoint can handle. This is used to calculate the\ncapacity of the group. Can be used in either CONNECTION or\nUTILIZATION balancing modes.\n\nFor CONNECTION mode, either\nmaxConnections or maxConnectionsPerEndpoint must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_connections_per_instance": {
              "computed": true,
              "description": "The max number of simultaneous connections that a single\nbackend instance can handle. This is used to calculate the\ncapacity of the group. Can be used in either CONNECTION or\nUTILIZATION balancing modes.\n\nFor CONNECTION mode, either maxConnections or\nmaxConnectionsPerInstance must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_rate": {
              "computed": true,
              "description": "The max requests per second (RPS) of the group.\n\nCan be used with either RATE or UTILIZATION balancing modes,\nbut required if RATE mode. For RATE mode, either maxRate or one\nof maxRatePerInstance or maxRatePerEndpoint, as appropriate for\ngroup type, must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_rate_per_endpoint": {
              "computed": true,
              "description": "The max requests per second (RPS) that a single backend network\nendpoint can handle. This is used to calculate the capacity of\nthe group. Can be used in either balancing mode. For RATE mode,\neither maxRate or maxRatePerEndpoint must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_rate_per_instance": {
              "computed": true,
              "description": "The max requests per second (RPS) that a single backend\ninstance can handle. This is used to calculate the capacity of\nthe group. Can be used in either balancing mode. For RATE mode,\neither maxRate or maxRatePerInstance must be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_utilization": {
              "computed": true,
              "description": "Used when balancingMode is UTILIZATION. This ratio defines the\nCPU utilization target for the group. Valid range is [0.0, 1.0].",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "preference": {
              "description": "This field indicates whether this backend should be fully utilized before sending traffic to backends\nwith default preference. This field cannot be set when loadBalancingScheme is set to 'EXTERNAL'. The possible values are:\n  - PREFERRED: Backends with this preference level will be filled up to their capacity limits first,\n    based on RTT.\n  - DEFAULT: If preferred backends don't have enough capacity, backends in this layer would be used and\n    traffic would be assigned based on the load balancing algorithm you use. This is the default Possible values: [\"PREFERRED\", \"DEFAULT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "custom_metrics": {
              "block": {
                "attributes": {
                  "dry_run": {
                    "description": "If true, the metric data is collected and reported to Cloud\nMonitoring, but is not used for load balancing.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "max_utilization": {
                    "description": "Optional parameter to define a target utilization for the Custom Metrics\nbalancing mode. The valid range is \u003ccode\u003e[0.0, 1.0]\u003c/code\u003e.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "name": {
                    "description": "Name of a custom utilization signal. The name must be 1-64 characters\nlong and match the regular expression [a-z]([-_.a-z0-9]*[a-z0-9])? which\nmeans the first character must be a lowercase letter, and all following\ncharacters must be a dash, period, underscore, lowercase letter, or\ndigit, except the last character, which cannot be a dash, period, or\nunderscore. For usage guidelines, see Custom Metrics balancing mode. This\nfield can only be used for a global or regional backend service with the\nloadBalancingScheme set to \u003ccode\u003eEXTERNAL_MANAGED\u003c/code\u003e,\n\u003ccode\u003eINTERNAL_MANAGED\u003c/code\u003e \u003ccode\u003eINTERNAL_SELF_MANAGED\u003c/code\u003e.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The set of custom metrics that are used for \u003ccode\u003eCUSTOM_METRICS\u003c/code\u003e BalancingMode.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "The set of backends that serve this BackendService.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "cdn_policy": {
        "block": {
          "attributes": {
            "cache_mode": {
              "computed": true,
              "description": "Specifies the cache setting for all responses from this backend.\nThe possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC Possible values: [\"USE_ORIGIN_HEADERS\", \"FORCE_CACHE_ALL\", \"CACHE_ALL_STATIC\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_ttl": {
              "computed": true,
              "description": "Specifies the maximum allowed TTL for cached content served by this origin.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "default_ttl": {
              "computed": true,
              "description": "Specifies the default TTL for cached content served by this origin for responses\nthat do not have an existing valid TTL (max-age or s-max-age).",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_ttl": {
              "computed": true,
              "description": "Specifies the maximum allowed TTL for cached content served by this origin.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "negative_caching": {
              "computed": true,
              "description": "Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "request_coalescing": {
              "computed": true,
              "description": "If true then Cloud CDN will combine multiple concurrent cache fill requests into a small number of requests\nto the origin.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "serve_while_stale": {
              "computed": true,
              "description": "Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "signed_url_cache_max_age_sec": {
              "description": "Maximum number of seconds the response to a signed URL request\nwill be considered fresh, defaults to 1hr (3600s). After this\ntime period, the response will be revalidated before\nbeing served.\n\nWhen serving responses to signed URL requests, Cloud CDN will\ninternally behave as though all responses from this backend had a\n\"Cache-Control: public, max-age=[TTL]\" header, regardless of any\nexisting Cache-Control header. The actual headers served in\nresponses will not be altered.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "bypass_cache_on_request_headers": {
              "block": {
                "attributes": {
                  "header_name": {
                    "description": "The header field name to match on when bypassing cache. Values are case-insensitive.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Bypass the cache when the specified request headers are matched - e.g. Pragma or Authorization headers. Up to 5 headers can be specified.\nThe cache is bypassed for all cdnPolicy.cacheMode settings.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "cache_key_policy": {
              "block": {
                "attributes": {
                  "include_host": {
                    "description": "If true requests to different hosts will be cached separately.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "include_http_headers": {
                    "description": "Allows HTTP request headers (by name) to be used in the\ncache key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "include_named_cookies": {
                    "description": "Names of cookies to include in cache keys.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "include_protocol": {
                    "description": "If true, http and https requests will be cached separately.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "include_query_string": {
                    "description": "If true, include query string parameters in the cache key\naccording to query_string_whitelist and\nquery_string_blacklist. If neither is set, the entire query\nstring will be included.\n\nIf false, the query string will be excluded from the cache\nkey entirely.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "query_string_blacklist": {
                    "description": "Names of query string parameters to exclude in cache keys.\n\nAll other parameters will be included. Either specify\nquery_string_whitelist or query_string_blacklist, not both.\n'\u0026' and '=' will be percent encoded and not treated as\ndelimiters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "query_string_whitelist": {
                    "description": "Names of query string parameters to include in cache keys.\n\nAll other parameters will be excluded. Either specify\nquery_string_whitelist or query_string_blacklist, not both.\n'\u0026' and '=' will be percent encoded and not treated as\ndelimiters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description": "The CacheKeyPolicy for this CdnPolicy.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "negative_caching_policy": {
              "block": {
                "attributes": {
                  "code": {
                    "description": "The HTTP status code to define a TTL against. Only HTTP status codes 300, 301, 308, 404, 405, 410, 421, 451 and 501\ncan be specified as values, and you cannot specify a status code more than once.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ttl": {
                    "description": "The TTL (in seconds) for which to cache responses with the corresponding status code. The maximum allowed value is 1800s\n(30 minutes), noting that infrequently accessed objects may be evicted from the cache before the defined TTL.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.\nOmitting the policy and leaving negativeCaching enabled will use Cloud CDN's default cache TTLs.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Cloud CDN configuration for this BackendService.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "circuit_breakers": {
        "block": {
          "attributes": {
            "max_connections": {
              "description": "The maximum number of connections to the backend cluster.\nDefaults to 1024.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_pending_requests": {
              "description": "The maximum number of pending requests to the backend cluster.\nDefaults to 1024.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_requests": {
              "description": "The maximum number of parallel requests to the backend cluster.\nDefaults to 1024.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_requests_per_connection": {
              "description": "Maximum requests for a single backend connection. This parameter\nis respected by both the HTTP/1.1 and HTTP/2 implementations. If\nnot specified, there is no limit. Setting this parameter to 1\nwill effectively disable keep alive.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_retries": {
              "description": "The maximum number of parallel retries to the backend cluster.\nDefaults to 3.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Settings controlling the volume of connections to a backend service. This field\nis applicable only when the load_balancing_scheme is set to INTERNAL_SELF_MANAGED.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "consistent_hash": {
        "block": {
          "attributes": {
            "http_header_name": {
              "description": "The hash based on the value of the specified header field.\nThis field is applicable if the sessionAffinity is set to HEADER_FIELD.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "minimum_ring_size": {
              "description": "The minimum number of virtual nodes to use for the hash ring.\nLarger ring sizes result in more granular load\ndistributions. If the number of hosts in the load balancing pool\nis larger than the ring size, each host will be assigned a single\nvirtual node.\nDefaults to 1024.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "http_cookie": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name of the cookie.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "description": "Path to set for the cookie.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "ttl": {
                    "block": {
                      "attributes": {
                        "nanos": {
                          "description": "Span of time that's a fraction of a second at nanosecond\nresolution. Durations less than one second are represented\nwith a 0 seconds field and a positive nanos field. Must\nbe from 0 to 999,999,999 inclusive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "seconds": {
                          "description": "Span of time at a resolution of a second.\nMust be from 0 to 315,576,000,000 inclusive.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "Lifetime of the cookie.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Hash is based on HTTP Cookie. This field describes a HTTP cookie\nthat will be used as the hash key for the consistent hash load\nbalancer. If the cookie is not present, it will be generated.\nThis field is applicable if the sessionAffinity is set to HTTP_COOKIE.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Consistent Hash-based load balancing can be used to provide soft session\naffinity based on HTTP headers, cookies or other properties. This load balancing\npolicy is applicable only for HTTP connections. The affinity to a particular\ndestination host will be lost when one or more hosts are added/removed from the\ndestination service. This field specifies parameters that control consistent\nhashing. This field only applies if the load_balancing_scheme is set to\nINTERNAL_SELF_MANAGED. This field is only applicable when locality_lb_policy is\nset to MAGLEV or RING_HASH.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "custom_metrics": {
        "block": {
          "attributes": {
            "dry_run": {
              "description": "If true, the metric data is not used for load balancing.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "name": {
              "description": "Name of a custom utilization signal. The name must be 1-64 characters\nlong and match the regular expression [a-z]([-_.a-z0-9]*[a-z0-9])? which\nmeans the first character must be a lowercase letter, and all following\ncharacters must be a dash, period, underscore, lowercase letter, or\ndigit, except the last character, which cannot be a dash, period, or\nunderscore. For usage guidelines, see Custom Metrics balancing mode. This\nfield can only be used for a global or regional backend service with the\nloadBalancingScheme set to \u003ccode\u003eEXTERNAL_MANAGED\u003c/code\u003e,\n\u003ccode\u003eINTERNAL_MANAGED\u003c/code\u003e \u003ccode\u003eINTERNAL_SELF_MANAGED\u003c/code\u003e.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "List of custom metrics that are used for the WEIGHTED_ROUND_ROBIN locality_lb_policy.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "iap": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Whether the serving infrastructure will authenticate and authorize all incoming requests.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "oauth2_client_id": {
              "description": "OAuth2 Client ID for IAP",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "oauth2_client_secret": {
              "description": "OAuth2 Client Secret for IAP",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "oauth2_client_secret_sha256": {
              "computed": true,
              "description": "OAuth2 Client Secret SHA-256 for IAP",
              "description_kind": "plain",
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "Settings for enabling Cloud Identity Aware Proxy.\nIf OAuth client is not set, the Google-managed OAuth client is used.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "locality_lb_policies": {
        "block": {
          "block_types": {
            "custom_policy": {
              "block": {
                "attributes": {
                  "data": {
                    "description": "An optional, arbitrary JSON object with configuration data, understood\nby a locally installed custom policy implementation.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "Identifies the custom policy.\n\nThe value should match the type the custom implementation is registered\nwith on the gRPC clients. It should follow protocol buffer\nmessage naming conventions and include the full path (e.g.\nmyorg.CustomLbPolicy). The maximum length is 256 characters.\n\nNote that specifying the same custom policy more than once for a\nbackend is not a valid configuration and will be rejected.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The configuration for a custom policy implemented by the user and\ndeployed with the client.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "policy": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "The name of a locality load balancer policy to be used. The value\nshould be one of the predefined ones as supported by localityLbPolicy,\nalthough at the moment only ROUND_ROBIN is supported.\n\nThis field should only be populated when the customPolicy field is not\nused.\n\nNote that specifying the same policy more than once for a backend is\nnot a valid configuration and will be rejected.\n\nThe possible values are:\n\n* 'ROUND_ROBIN': This is a simple policy in which each healthy backend\n                is selected in round robin order.\n\n* 'LEAST_REQUEST': An O(1) algorithm which selects two random healthy\n                  hosts and picks the host which has fewer active requests.\n\n* 'RING_HASH': The ring/modulo hash load balancer implements consistent\n              hashing to backends. The algorithm has the property that the\n              addition/removal of a host from a set of N hosts only affects\n              1/N of the requests.\n\n* 'RANDOM': The load balancer selects a random healthy host.\n\n* 'ORIGINAL_DESTINATION': Backend host is selected based on the client\n                          connection metadata, i.e., connections are opened\n                          to the same address as the destination address of\n                          the incoming connection before the connection\n                          was redirected to the load balancer.\n\n* 'MAGLEV': used as a drop in replacement for the ring hash load balancer.\n            Maglev is not as stable as ring hash but has faster table lookup\n            build times and host selection times. For more information about\n            Maglev, refer to https://ai.google/research/pubs/pub44824 Possible values: [\"ROUND_ROBIN\", \"LEAST_REQUEST\", \"RING_HASH\", \"RANDOM\", \"ORIGINAL_DESTINATION\", \"MAGLEV\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The configuration for a built-in load balancing policy.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A list of locality load balancing policies to be used in order of\npreference. Either the policy or the customPolicy field should be set.\nOverrides any value set in the localityLbPolicy field.\n\nlocalityLbPolicies is only supported when the BackendService is referenced\nby a URL Map that is referenced by a target gRPC proxy that has the\nvalidateForProxyless field set to true.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "log_config": {
        "block": {
          "attributes": {
            "enable": {
              "description": "Whether to enable logging for the load balancer traffic served by this backend service.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "optional_fields": {
              "description": "This field can only be specified if logging is enabled for this backend service and \"logConfig.optionalMode\"\nwas set to CUSTOM. Contains a list of optional fields you want to include in the logs.\nFor example: serverInstance, serverGkeDetails.cluster, serverGkeDetails.pod.podNamespace\nFor example: orca_load_report, tls.protocol",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "optional_mode": {
              "computed": true,
              "description": "Specifies the optional logging mode for the load balancer traffic.\nSupported values: INCLUDE_ALL_OPTIONAL, EXCLUDE_ALL_OPTIONAL, CUSTOM. Possible values: [\"INCLUDE_ALL_OPTIONAL\", \"EXCLUDE_ALL_OPTIONAL\", \"CUSTOM\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sample_rate": {
              "description": "This field can only be specified if logging is enabled for this backend service. The value of\nthe field must be in [0, 1]. This configures the sampling rate of requests to the load balancer\nwhere 1.0 means all logged requests are reported and 0.0 means no logged requests are reported.\nThe default value is 1.0.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "This field denotes the logging options for the load balancer traffic served by this backend service.\nIf logging is enabled, logs will be exported to Stackdriver.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "max_stream_duration": {
        "block": {
          "attributes": {
            "nanos": {
              "description": "Span of time that's a fraction of a second at nanosecond resolution.\nDurations less than one second are represented with a 0 seconds field and a positive nanos field.\nMust be from 0 to 999,999,999 inclusive.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "seconds": {
              "description": "Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive. (int64 format)",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the\nbeginning of the stream until the response has been completely processed, including all retries. A stream that\ndoes not complete in this duration is closed.\nIf not specified, there will be no timeout limit, i.e. the maximum duration is infinite.\nThis value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service.\nThis field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "outlier_detection": {
        "block": {
          "attributes": {
            "consecutive_errors": {
              "description": "Number of errors before a host is ejected from the connection pool. When the\nbackend host is accessed over HTTP, a 5xx return code qualifies as an error.\nDefaults to 5.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "consecutive_gateway_failure": {
              "description": "The number of consecutive gateway failures (502, 503, 504 status or connection\nerrors that are mapped to one of those status codes) before a consecutive\ngateway failure ejection occurs. Defaults to 5.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enforcing_consecutive_errors": {
              "description": "The percentage chance that a host will be actually ejected when an outlier\nstatus is detected through consecutive 5xx. This setting can be used to disable\nejection or to ramp it up slowly. Defaults to 100.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enforcing_consecutive_gateway_failure": {
              "description": "The percentage chance that a host will be actually ejected when an outlier\nstatus is detected through consecutive gateway failures. This setting can be\nused to disable ejection or to ramp it up slowly. Defaults to 0.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enforcing_success_rate": {
              "description": "The percentage chance that a host will be actually ejected when an outlier\nstatus is detected through success rate statistics. This setting can be used to\ndisable ejection or to ramp it up slowly. Defaults to 100.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_ejection_percent": {
              "description": "Maximum percentage of hosts in the load balancing pool for the backend service\nthat can be ejected. Defaults to 10%.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "success_rate_minimum_hosts": {
              "description": "The number of hosts in a cluster that must have enough request volume to detect\nsuccess rate outliers. If the number of hosts is less than this setting, outlier\ndetection via success rate statistics is not performed for any host in the\ncluster. Defaults to 5.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "success_rate_request_volume": {
              "description": "The minimum number of total requests that must be collected in one interval (as\ndefined by the interval duration above) to include this host in success rate\nbased outlier detection. If the volume is lower than this setting, outlier\ndetection via success rate statistics is not performed for that host. Defaults\nto 100.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "success_rate_stdev_factor": {
              "description": "This factor is used to determine the ejection threshold for success rate outlier\nejection. The ejection threshold is the difference between the mean success\nrate, and the product of this factor and the standard deviation of the mean\nsuccess rate: mean - (stdev * success_rate_stdev_factor). This factor is divided\nby a thousand to get a double. That is, if the desired factor is 1.9, the\nruntime value should be 1900. Defaults to 1900.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "base_ejection_time": {
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
                    "type": "number"
                  }
                },
                "description": "The base time that a host is ejected for. The real time is equal to the base\ntime multiplied by the number of times the host has been ejected. Defaults to\n30000ms or 30s.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "interval": {
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
                    "type": "number"
                  }
                },
                "description": "Time interval between ejection sweep analysis. This can result in both new\nejections as well as hosts being returned to service. Defaults to 10 seconds.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Settings controlling eviction of unhealthy hosts from the load balancing pool.\nApplicable backend service types can be a global backend service with the\nloadBalancingScheme set to INTERNAL_SELF_MANAGED or EXTERNAL_MANAGED.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "security_settings": {
        "block": {
          "attributes": {
            "client_tls_policy": {
              "description": "ClientTlsPolicy is a resource that specifies how a client should authenticate\nconnections to backends of a service. This resource itself does not affect\nconfiguration unless it is attached to a backend service resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subject_alt_names": {
              "description": "A list of alternate names to verify the subject identity in the certificate.\nIf specified, the client will verify that the server certificate's subject\nalt name matches one of the specified values.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "aws_v4_authentication": {
              "block": {
                "attributes": {
                  "access_key": {
                    "description": "The access key used for s3 bucket authentication.\nRequired for updating or creating a backend that uses AWS v4 signature authentication, but will not be returned as part of the configuration when queried with a REST API GET request.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "access_key_id": {
                    "description": "The identifier of an access key used for s3 bucket authentication.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "access_key_version": {
                    "description": "The optional version identifier for the access key. You can use this to keep track of different iterations of your access key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "origin_region": {
                    "description": "The name of the cloud region of your origin. This is a free-form field with the name of the region your cloud uses to host your origin.\nFor example, \"us-east-1\" for AWS or \"us-ashburn-1\" for OCI.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The configuration needed to generate a signature for access to private storage buckets that support AWS's Signature Version 4 for authentication.\nAllowed only for INTERNET_IP_PORT and INTERNET_FQDN_PORT NEG backends.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The security settings that apply to this backend service. This field is applicable to either\na regional backend service with the service_protocol set to HTTP, HTTPS, HTTP2 or H2C, and\nload_balancing_scheme set to INTERNAL_MANAGED; or a global backend service with the\nload_balancing_scheme set to INTERNAL_SELF_MANAGED.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "strong_session_affinity_cookie": {
        "block": {
          "attributes": {
            "name": {
              "description": "Name of the cookie.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "description": "Path to set for the cookie.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "ttl": {
              "block": {
                "attributes": {
                  "nanos": {
                    "description": "Span of time that's a fraction of a second at nanosecond\nresolution. Durations less than one second are represented\nwith a 0 seconds field and a positive nanos field. Must\nbe from 0 to 999,999,999 inclusive.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "seconds": {
                    "description": "Span of time at a resolution of a second.\nMust be from 0 to 315,576,000,000 inclusive.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Lifetime of the cookie.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Describes the HTTP cookie used for stateful session affinity. This field is applicable and required if the sessionAffinity is set to STRONG_COOKIE_AFFINITY.",
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
      },
      "tls_settings": {
        "block": {
          "attributes": {
            "authentication_config": {
              "description": "Reference to the BackendAuthenticationConfig resource from the networksecurity.googleapis.com namespace.\nCan be used in authenticating TLS connections to the backend, as specified by the authenticationMode field.\nCan only be specified if authenticationMode is not NONE.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sni": {
              "description": "Server Name Indication - see RFC3546 section 3.1. If set, the load balancer sends this string as the SNI hostname in the\nTLS connection to the backend, and requires that this string match a Subject Alternative Name (SAN) in the backend's\nserver certificate. With a Regional Internet NEG backend, if the SNI is specified here, the load balancer uses it\nregardless of whether the Regional Internet NEG is specified with FQDN or IP address and port.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "subject_alt_names": {
              "block": {
                "attributes": {
                  "dns_name": {
                    "description": "The SAN specified as a DNS Name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "uniform_resource_identifier": {
                    "description": "The SAN specified as a URI.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A list of Subject Alternative Names (SANs) that the Load Balancer verifies during a TLS handshake with the backend.\nWhen the server presents its X.509 certificate to the Load Balancer, the Load Balancer inspects the certificate's SAN field,\nand requires that at least one SAN match one of the subjectAltNames in the list. This field is limited to 5 entries.\nWhen both sni and subjectAltNames are specified, the load balancer matches the backend certificate's SAN only to\nsubjectAltNames.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for Backend Authenticated TLS and mTLS. May only be specified when the backend protocol is SSL, HTTPS or HTTP2.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleComputeBackendServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeBackendService), &result)
	return &result
}
