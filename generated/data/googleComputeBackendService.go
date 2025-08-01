package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeBackendService = `{
  "block": {
    "attributes": {
      "affinity_cookie_ttl_sec": {
        "computed": true,
        "description": "Lifetime of cookies in seconds if session_affinity is\nGENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts\nonly until the end of the browser session (or equivalent). The\nmaximum allowed value for TTL is one day.\n\nWhen the load balancing scheme is INTERNAL, this field is not used.",
        "description_kind": "plain",
        "type": "number"
      },
      "backend": {
        "computed": true,
        "description": "The set of backends that serve this BackendService.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "balancing_mode": "string",
              "capacity_scaler": "number",
              "custom_metrics": [
                "list",
                [
                  "object",
                  {
                    "dry_run": "bool",
                    "max_utilization": "number",
                    "name": "string"
                  }
                ]
              ],
              "description": "string",
              "group": "string",
              "max_connections": "number",
              "max_connections_per_endpoint": "number",
              "max_connections_per_instance": "number",
              "max_rate": "number",
              "max_rate_per_endpoint": "number",
              "max_rate_per_instance": "number",
              "max_utilization": "number",
              "preference": "string"
            }
          ]
        ]
      },
      "cdn_policy": {
        "computed": true,
        "description": "Cloud CDN configuration for this BackendService.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bypass_cache_on_request_headers": [
                "list",
                [
                  "object",
                  {
                    "header_name": "string"
                  }
                ]
              ],
              "cache_key_policy": [
                "list",
                [
                  "object",
                  {
                    "include_host": "bool",
                    "include_http_headers": [
                      "list",
                      "string"
                    ],
                    "include_named_cookies": [
                      "list",
                      "string"
                    ],
                    "include_protocol": "bool",
                    "include_query_string": "bool",
                    "query_string_blacklist": [
                      "set",
                      "string"
                    ],
                    "query_string_whitelist": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "cache_mode": "string",
              "client_ttl": "number",
              "default_ttl": "number",
              "max_ttl": "number",
              "negative_caching": "bool",
              "negative_caching_policy": [
                "list",
                [
                  "object",
                  {
                    "code": "number",
                    "ttl": "number"
                  }
                ]
              ],
              "request_coalescing": "bool",
              "serve_while_stale": "number",
              "signed_url_cache_max_age_sec": "number"
            }
          ]
        ]
      },
      "circuit_breakers": {
        "computed": true,
        "description": "Settings controlling the volume of connections to a backend service. This field\nis applicable only when the load_balancing_scheme is set to INTERNAL_SELF_MANAGED.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "max_connections": "number",
              "max_pending_requests": "number",
              "max_requests": "number",
              "max_requests_per_connection": "number",
              "max_retries": "number"
            }
          ]
        ]
      },
      "compression_mode": {
        "computed": true,
        "description": "Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header. Possible values: [\"AUTOMATIC\", \"DISABLED\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "connection_draining_timeout_sec": {
        "computed": true,
        "description": "Time for which instance will be drained (not accept new\nconnections, but still work to finish started).",
        "description_kind": "plain",
        "type": "number"
      },
      "consistent_hash": {
        "computed": true,
        "description": "Consistent Hash-based load balancing can be used to provide soft session\naffinity based on HTTP headers, cookies or other properties. This load balancing\npolicy is applicable only for HTTP connections. The affinity to a particular\ndestination host will be lost when one or more hosts are added/removed from the\ndestination service. This field specifies parameters that control consistent\nhashing. This field only applies if the load_balancing_scheme is set to\nINTERNAL_SELF_MANAGED. This field is only applicable when locality_lb_policy is\nset to MAGLEV or RING_HASH.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "http_cookie": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "path": "string",
                    "ttl": [
                      "list",
                      [
                        "object",
                        {
                          "nanos": "number",
                          "seconds": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "http_header_name": "string",
              "minimum_ring_size": "number"
            }
          ]
        ]
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_metrics": {
        "computed": true,
        "description": "List of custom metrics that are used for the WEIGHTED_ROUND_ROBIN locality_lb_policy.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dry_run": "bool",
              "name": "string"
            }
          ]
        ]
      },
      "custom_request_headers": {
        "computed": true,
        "description": "Headers that the HTTP/S load balancer should add to proxied\nrequests.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "custom_response_headers": {
        "computed": true,
        "description": "Headers that the HTTP/S load balancer should add to proxied\nresponses.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "edge_security_policy": {
        "computed": true,
        "description": "The resource URL for the edge security policy associated with this backend service.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_cdn": {
        "computed": true,
        "description": "If true, enable Cloud CDN for this BackendService.",
        "description_kind": "plain",
        "type": "bool"
      },
      "external_managed_migration_state": {
        "computed": true,
        "description": "Specifies the canary migration state. Possible values are PREPARE, TEST_BY_PERCENTAGE, and\nTEST_ALL_TRAFFIC.\n\nTo begin the migration from EXTERNAL to EXTERNAL_MANAGED, the state must be changed to\nPREPARE. The state must be changed to TEST_ALL_TRAFFIC before the loadBalancingScheme can be\nchanged to EXTERNAL_MANAGED. Optionally, the TEST_BY_PERCENTAGE state can be used to migrate\ntraffic by percentage using externalManagedMigrationTestingPercentage.\n\nRolling back a migration requires the states to be set in reverse order. So changing the\nscheme from EXTERNAL_MANAGED to EXTERNAL requires the state to be set to TEST_ALL_TRAFFIC at\nthe same time. Optionally, the TEST_BY_PERCENTAGE state can be used to migrate some traffic\nback to EXTERNAL or PREPARE can be used to migrate all traffic back to EXTERNAL. Possible values: [\"PREPARE\", \"TEST_BY_PERCENTAGE\", \"TEST_ALL_TRAFFIC\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "external_managed_migration_testing_percentage": {
        "computed": true,
        "description": "Determines the fraction of requests that should be processed by the Global external\nApplication Load Balancer.\n\nThe value of this field must be in the range [0, 100].\n\nSession affinity options will slightly affect this routing behavior, for more details,\nsee: Session Affinity.\n\nThis value can only be set if the loadBalancingScheme in the backend service is set to\nEXTERNAL (when using the Classic ALB) and the migration state is TEST_BY_PERCENTAGE.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource\nfor health checking this BackendService. Currently at most one health\ncheck can be specified.\n\nA health check must be specified unless the backend service uses an internet\nor serverless NEG as a backend.\n\nFor internal load balancing, a URL to a HealthCheck resource must be specified instead.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "iap": {
        "computed": true,
        "description": "Settings for enabling Cloud Identity Aware Proxy.\nIf OAuth client is not set, the Google-managed OAuth client is used.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "oauth2_client_id": "string",
              "oauth2_client_secret": "string",
              "oauth2_client_secret_sha256": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_address_selection_policy": {
        "computed": true,
        "description": "Specifies preference of traffic to the backend (from the proxy and from the client for proxyless gRPC). Possible values: [\"IPV4_ONLY\", \"PREFER_IPV6\", \"IPV6_ONLY\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancing_scheme": {
        "computed": true,
        "description": "Indicates whether the backend service will be used with internal or\nexternal load balancing. A backend service created for one type of\nload balancing cannot be used with the other. For more information, refer to\n[Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service). Default value: \"EXTERNAL\" Possible values: [\"EXTERNAL\", \"INTERNAL_SELF_MANAGED\", \"INTERNAL_MANAGED\", \"EXTERNAL_MANAGED\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "locality_lb_policies": {
        "computed": true,
        "description": "A list of locality load balancing policies to be used in order of\npreference. Either the policy or the customPolicy field should be set.\nOverrides any value set in the localityLbPolicy field.\n\nlocalityLbPolicies is only supported when the BackendService is referenced\nby a URL Map that is referenced by a target gRPC proxy that has the\nvalidateForProxyless field set to true.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "custom_policy": [
                "list",
                [
                  "object",
                  {
                    "data": "string",
                    "name": "string"
                  }
                ]
              ],
              "policy": [
                "list",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "locality_lb_policy": {
        "computed": true,
        "description": "The load balancing algorithm used within the scope of the locality.\nThe possible values are:\n\n* 'ROUND_ROBIN': This is a simple policy in which each healthy backend\n                 is selected in round robin order.\n\n* 'LEAST_REQUEST': An O(1) algorithm which selects two random healthy\n                   hosts and picks the host which has fewer active requests.\n\n* 'RING_HASH': The ring/modulo hash load balancer implements consistent\n               hashing to backends. The algorithm has the property that the\n               addition/removal of a host from a set of N hosts only affects\n               1/N of the requests.\n\n* 'RANDOM': The load balancer selects a random healthy host.\n\n* 'ORIGINAL_DESTINATION': Backend host is selected based on the client\n                          connection metadata, i.e., connections are opened\n                          to the same address as the destination address of\n                          the incoming connection before the connection\n                          was redirected to the load balancer.\n\n* 'MAGLEV': used as a drop in replacement for the ring hash load balancer.\n            Maglev is not as stable as ring hash but has faster table lookup\n            build times and host selection times. For more information about\n            Maglev, refer to https://ai.google/research/pubs/pub44824\n\n* 'WEIGHTED_MAGLEV': Per-instance weighted Load Balancing via health check\n                     reported weights. Only applicable to loadBalancingScheme\n                     EXTERNAL. If set, the Backend Service must\n                     configure a non legacy HTTP-based Health Check, and\n                     health check replies are expected to contain\n                     non-standard HTTP response header field\n                     X-Load-Balancing-Endpoint-Weight to specify the\n                     per-instance weights. If set, Load Balancing is weight\n                     based on the per-instance weights reported in the last\n                     processed health check replies, as long as every\n                     instance either reported a valid weight or had\n                     UNAVAILABLE_WEIGHT. Otherwise, Load Balancing remains\n                     equal-weight.\n\n* 'WEIGHTED_ROUND_ROBIN': Per-endpoint weighted round-robin Load Balancing using weights computed\n                          from Backend reported Custom Metrics. If set, the Backend Service\n                          responses are expected to contain non-standard HTTP response header field\n                          X-Endpoint-Load-Metrics. The reported metrics\n                          to use for computing the weights are specified via the\n                          backends[].customMetrics fields.\n\nlocality_lb_policy is applicable to either:\n\n* A regional backend service with the service_protocol set to HTTP, HTTPS, HTTP2 or H2C,\n  and loadBalancingScheme set to INTERNAL_MANAGED.\n* A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.\n* A regional backend service with loadBalancingScheme set to EXTERNAL (External Network\n  Load Balancing). Only MAGLEV and WEIGHTED_MAGLEV values are possible for External\n  Network Load Balancing. The default is MAGLEV.\n\nIf session_affinity is not NONE, and locality_lb_policy is not set to MAGLEV, WEIGHTED_MAGLEV,\nor RING_HASH, session affinity settings will not take effect.\n\nOnly ROUND_ROBIN and RING_HASH are supported when the backend service is referenced\nby a URL map that is bound to target gRPC proxy that has validate_for_proxyless\nfield set to true. Possible values: [\"ROUND_ROBIN\", \"LEAST_REQUEST\", \"RING_HASH\", \"RANDOM\", \"ORIGINAL_DESTINATION\", \"MAGLEV\", \"WEIGHTED_MAGLEV\", \"WEIGHTED_ROUND_ROBIN\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "log_config": {
        "computed": true,
        "description": "This field denotes the logging options for the load balancer traffic served by this backend service.\nIf logging is enabled, logs will be exported to Stackdriver.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable": "bool",
              "optional_fields": [
                "list",
                "string"
              ],
              "optional_mode": "string",
              "sample_rate": "number"
            }
          ]
        ]
      },
      "max_stream_duration": {
        "computed": true,
        "description": "Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the\nbeginning of the stream until the response has been completely processed, including all retries. A stream that\ndoes not complete in this duration is closed.\nIf not specified, there will be no timeout limit, i.e. the maximum duration is infinite.\nThis value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service.\nThis field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "nanos": "number",
              "seconds": "string"
            }
          ]
        ]
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "outlier_detection": {
        "computed": true,
        "description": "Settings controlling eviction of unhealthy hosts from the load balancing pool.\nApplicable backend service types can be a global backend service with the\nloadBalancingScheme set to INTERNAL_SELF_MANAGED or EXTERNAL_MANAGED.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "base_ejection_time": [
                "list",
                [
                  "object",
                  {
                    "nanos": "number",
                    "seconds": "number"
                  }
                ]
              ],
              "consecutive_errors": "number",
              "consecutive_gateway_failure": "number",
              "enforcing_consecutive_errors": "number",
              "enforcing_consecutive_gateway_failure": "number",
              "enforcing_success_rate": "number",
              "interval": [
                "list",
                [
                  "object",
                  {
                    "nanos": "number",
                    "seconds": "number"
                  }
                ]
              ],
              "max_ejection_percent": "number",
              "success_rate_minimum_hosts": "number",
              "success_rate_request_volume": "number",
              "success_rate_stdev_factor": "number"
            }
          ]
        ]
      },
      "port_name": {
        "computed": true,
        "description": "Name of backend port. The same name should appear in the instance\ngroups referenced by this service. Required when the load balancing\nscheme is EXTERNAL.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
        "computed": true,
        "description": "The protocol this BackendService uses to communicate with backends.\nThe default is HTTP. Possible values are HTTP, HTTPS, HTTP2, H2C, TCP, SSL, UDP\nor GRPC. Refer to the documentation for the load balancers or for Traffic Director\nfor more information. Must be set to GRPC when the backend service is referenced\nby a URL map that is bound to target gRPC proxy. Possible values: [\"HTTP\", \"HTTPS\", \"HTTP2\", \"TCP\", \"SSL\", \"UDP\", \"GRPC\", \"UNSPECIFIED\", \"H2C\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "security_policy": {
        "computed": true,
        "description": "The security policy associated with this backend service.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_settings": {
        "computed": true,
        "description": "The security settings that apply to this backend service. This field is applicable to either\na regional backend service with the service_protocol set to HTTP, HTTPS, HTTP2 or H2C, and\nload_balancing_scheme set to INTERNAL_MANAGED; or a global backend service with the\nload_balancing_scheme set to INTERNAL_SELF_MANAGED.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "aws_v4_authentication": [
                "list",
                [
                  "object",
                  {
                    "access_key": "string",
                    "access_key_id": "string",
                    "access_key_version": "string",
                    "origin_region": "string"
                  }
                ]
              ],
              "client_tls_policy": "string",
              "subject_alt_names": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_lb_policy": {
        "computed": true,
        "description": "URL to networkservices.ServiceLbPolicy resource.\nCan only be set if load balancing scheme is EXTERNAL, EXTERNAL_MANAGED, INTERNAL_MANAGED or INTERNAL_SELF_MANAGED and the scope is global.",
        "description_kind": "plain",
        "type": "string"
      },
      "session_affinity": {
        "computed": true,
        "description": "Type of session affinity to use. The default is NONE. Session affinity is\nnot applicable if the protocol is UDP. Possible values: [\"NONE\", \"CLIENT_IP\", \"CLIENT_IP_PORT_PROTO\", \"CLIENT_IP_PROTO\", \"GENERATED_COOKIE\", \"HEADER_FIELD\", \"HTTP_COOKIE\", \"STRONG_COOKIE_AFFINITY\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "strong_session_affinity_cookie": {
        "computed": true,
        "description": "Describes the HTTP cookie used for stateful session affinity. This field is applicable and required if the sessionAffinity is set to STRONG_COOKIE_AFFINITY.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "path": "string",
              "ttl": [
                "list",
                [
                  "object",
                  {
                    "nanos": "number",
                    "seconds": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "timeout_sec": {
        "computed": true,
        "description": "The backend service timeout has a different meaning depending on the type of load balancer.\nFor more information see, [Backend service settings](https://cloud.google.com/compute/docs/reference/rest/v1/backendServices).\nThe default is 30 seconds.\nThe full range of timeout values allowed goes from 1 through 2,147,483,647 seconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "tls_settings": {
        "computed": true,
        "description": "Configuration for Backend Authenticated TLS and mTLS. May only be specified when the backend protocol is SSL, HTTPS or HTTP2.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "authentication_config": "string",
              "sni": "string",
              "subject_alt_names": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "uniform_resource_identifier": "string"
                  }
                ]
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeBackendServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeBackendService), &result)
	return &result
}
