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
              "description": "string",
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
      "enable_cdn": {
        "computed": true,
        "description": "If true, enable Cloud CDN for this BackendService.",
        "description_kind": "plain",
        "type": "bool"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource. A hash of the contents stored in this\nobject. This field is used in optimistic locking.",
        "description_kind": "plain",
        "type": "string"
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
        "description": "Settings for enabling Cloud Identity Aware Proxy",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
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
      "load_balancing_scheme": {
        "computed": true,
        "description": "Indicates whether the backend service will be used with internal or\nexternal load balancing. A backend service created for one type of\nload balancing cannot be used with the other. Default value: \"EXTERNAL\" Possible values: [\"EXTERNAL\", \"INTERNAL_SELF_MANAGED\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "locality_lb_policy": {
        "computed": true,
        "description": "The load balancing algorithm used within the scope of the locality.\nThe possible values are -\n\n* ROUND_ROBIN - This is a simple policy in which each healthy backend\n                is selected in round robin order.\n\n* LEAST_REQUEST - An O(1) algorithm which selects two random healthy\n                  hosts and picks the host which has fewer active requests.\n\n* RING_HASH - The ring/modulo hash load balancer implements consistent\n              hashing to backends. The algorithm has the property that the\n              addition/removal of a host from a set of N hosts only affects\n              1/N of the requests.\n\n* RANDOM - The load balancer selects a random healthy host.\n\n* ORIGINAL_DESTINATION - Backend host is selected based on the client\n                         connection metadata, i.e., connections are opened\n                         to the same address as the destination address of\n                         the incoming connection before the connection\n                         was redirected to the load balancer.\n\n* MAGLEV - used as a drop in replacement for the ring hash load balancer.\n           Maglev is not as stable as ring hash but has faster table lookup\n           build times and host selection times. For more information about\n           Maglev, refer to https://ai.google/research/pubs/pub44824\n\nThis field is applicable only when the load_balancing_scheme is set to\nINTERNAL_SELF_MANAGED. Possible values: [\"ROUND_ROBIN\", \"LEAST_REQUEST\", \"RING_HASH\", \"RANDOM\", \"ORIGINAL_DESTINATION\", \"MAGLEV\"]",
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
      "outlier_detection": {
        "computed": true,
        "description": "Settings controlling eviction of unhealthy hosts from the load balancing pool.\nThis field is applicable only when the load_balancing_scheme is set\nto INTERNAL_SELF_MANAGED.",
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
        "description": "The protocol this BackendService uses to communicate with backends.\nThe default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer\ntypes and may result in errors if used with the GA API. Possible values: [\"HTTP\", \"HTTPS\", \"HTTP2\", \"TCP\", \"SSL\", \"GRPC\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "security_policy": {
        "computed": true,
        "description": "The security policy associated with this backend service.",
        "description_kind": "plain",
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "session_affinity": {
        "computed": true,
        "description": "Type of session affinity to use. The default is NONE. Session affinity is\nnot applicable if the protocol is UDP. Possible values: [\"NONE\", \"CLIENT_IP\", \"CLIENT_IP_PORT_PROTO\", \"CLIENT_IP_PROTO\", \"GENERATED_COOKIE\", \"HEADER_FIELD\", \"HTTP_COOKIE\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "timeout_sec": {
        "computed": true,
        "description": "How many seconds to wait for the backend before considering it a\nfailed request. Default is 30 seconds. Valid range is [1, 86400].",
        "description_kind": "plain",
        "type": "number"
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
