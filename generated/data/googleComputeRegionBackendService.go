package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionBackendService = `{
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
        "description": "The set of backends that serve this RegionBackendService.",
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
              "failover": "bool",
              "group": "string",
              "max_connections": "number",
              "max_connections_per_endpoint": "number",
              "max_connections_per_instance": "number",
              "max_rate": "number",
              "max_rate_per_endpoint": "number",
              "max_rate_per_instance": "number",
              "max_utilization": "number"
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
              "cache_key_policy": [
                "list",
                [
                  "object",
                  {
                    "include_host": "bool",
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
                    "code": "number"
                  }
                ]
              ],
              "serve_while_stale": "number",
              "signed_url_cache_max_age_sec": "number"
            }
          ]
        ]
      },
      "circuit_breakers": {
        "computed": true,
        "description": "Settings controlling the volume of connections to a backend service. This field\nis applicable only when the 'load_balancing_scheme' is set to INTERNAL_MANAGED\nand the 'protocol' is set to HTTP, HTTPS, HTTP2 or H2C.",
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
      "connection_draining_timeout_sec": {
        "computed": true,
        "description": "Time for which instance will be drained (not accept new\nconnections, but still work to finish started).",
        "description_kind": "plain",
        "type": "number"
      },
      "consistent_hash": {
        "computed": true,
        "description": "Consistent Hash-based load balancing can be used to provide soft session\naffinity based on HTTP headers, cookies or other properties. This load balancing\npolicy is applicable only for HTTP connections. The affinity to a particular\ndestination host will be lost when one or more hosts are added/removed from the\ndestination service. This field specifies parameters that control consistent\nhashing.\nThis field only applies when all of the following are true -\n  * 'load_balancing_scheme' is set to INTERNAL_MANAGED\n  * 'protocol' is set to HTTP, HTTPS, HTTP2 or H2C\n  * 'locality_lb_policy' is set to MAGLEV or RING_HASH",
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
      "description": {
        "computed": true,
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_cdn": {
        "computed": true,
        "description": "If true, enable Cloud CDN for this RegionBackendService.",
        "description_kind": "plain",
        "type": "bool"
      },
      "failover_policy": {
        "computed": true,
        "description": "Policy for failovers.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disable_connection_drain_on_failover": "bool",
              "drop_traffic_if_unhealthy": "bool",
              "failover_ratio": "number"
            }
          ]
        ]
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
      "ha_policy": {
        "computed": true,
        "description": "Configures self-managed High Availability (HA) for External and Internal Protocol Forwarding.\nThe backends of this regional backend service must only specify zonal network endpoint groups\n(NEGs) of type GCE_VM_IP. Note that haPolicy is not for load balancing, and therefore cannot\nbe specified with sessionAffinity, connectionTrackingPolicy, and failoverPolicy. haPolicy\nrequires customers to be responsible for tracking backend endpoint health and electing a\nleader among the healthy endpoints. Therefore, haPolicy cannot be specified with healthChecks.\nhaPolicy can only be specified for External Passthrough Network Load Balancers and Internal\nPassthrough Network Load Balancers.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "fast_ip_move": "string",
              "leader": [
                "list",
                [
                  "object",
                  {
                    "backend_group": "string",
                    "network_endpoint": [
                      "list",
                      [
                        "object",
                        {
                          "instance": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "health_checks": {
        "computed": true,
        "description": "The set of URLs to HealthCheck resources for health checking\nthis RegionBackendService. Currently at most one health\ncheck can be specified.\n\nA health check must be specified unless the backend service uses an internet\nor serverless NEG as a backend.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "iap": {
        "computed": true,
        "description": "Settings for enabling Cloud Identity Aware Proxy.\nIf OAuth client is not set, Google-managed OAuth client is used.",
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
        "description": "Indicates what kind of load balancing this regional backend service\nwill be used for. A backend service created for one type of load\nbalancing cannot be used with the other(s). For more information, refer to\n[Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service). Default value: \"INTERNAL\" Possible values: [\"EXTERNAL\", \"EXTERNAL_MANAGED\", \"INTERNAL\", \"INTERNAL_MANAGED\"]",
        "description_kind": "plain",
        "type": "string"
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
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "The URL of the network to which this backend service belongs.\nThis field can only be specified when the load balancing scheme is set to INTERNAL.",
        "description_kind": "plain",
        "type": "string"
      },
      "outlier_detection": {
        "computed": true,
        "description": "Settings controlling eviction of unhealthy hosts from the load balancing pool.\nThis field is applicable only when the 'load_balancing_scheme' is set\nto INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, HTTP2 or H2C.",
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
        "description": "A named port on a backend instance group representing the port for\ncommunication to the backend VMs in that group. Required when the\nloadBalancingScheme is EXTERNAL, EXTERNAL_MANAGED, INTERNAL_MANAGED, or INTERNAL_SELF_MANAGED\nand the backends are instance groups. The named port must be defined on each\nbackend instance group. This parameter has no meaning if the backends are NEGs. API sets a\ndefault of \"http\" if not given.\nMust be omitted when the loadBalancingScheme is INTERNAL (Internal TCP/UDP Load Balancing).",
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
        "description": "The protocol this BackendService uses to communicate with backends.\nThe default is HTTP. Possible values are HTTP, HTTPS, HTTP2, H2C, TCP, SSL, UDP\nor GRPC. Refer to the documentation for the load balancers or for Traffic Director\nfor more information. Possible values: [\"HTTP\", \"HTTPS\", \"HTTP2\", \"TCP\", \"SSL\", \"UDP\", \"GRPC\", \"UNSPECIFIED\", \"H2C\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "description": "The Region in which the created backend service should reside.\nIf it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "session_affinity": {
        "computed": true,
        "description": "Type of session affinity to use. The default is NONE. Session affinity is\nnot applicable if the protocol is UDP. Possible values: [\"NONE\", \"CLIENT_IP\", \"CLIENT_IP_PORT_PROTO\", \"CLIENT_IP_PROTO\", \"GENERATED_COOKIE\", \"HEADER_FIELD\", \"HTTP_COOKIE\", \"CLIENT_IP_NO_DESTINATION\", \"STRONG_COOKIE_AFFINITY\"]",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeRegionBackendServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionBackendService), &result)
	return &result
}
