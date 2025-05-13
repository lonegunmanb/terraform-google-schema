package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeTargetServer = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A human-readable description of this TargetServer.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "env_id": {
        "description": "The Apigee environment group associated with the Apigee environment,\nin the format 'organizations/{{org_name}}/environments/{{env_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host": {
        "description": "The host name this target connects to. Value must be a valid hostname as described by RFC-1123.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_enabled": {
        "description": "Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description": "The resource id of this reference. Values must match the regular expression [\\w\\s-.]+.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port": {
        "description": "The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "protocol": {
        "computed": true,
        "description": "Immutable. The protocol used by this TargetServer. Possible values: [\"HTTP\", \"HTTP2\", \"GRPC_TARGET\", \"GRPC\", \"EXTERNAL_CALLOUT\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "s_sl_info": {
        "block": {
          "attributes": {
            "ciphers": {
              "description": "The SSL/TLS cipher suites to be used. For programmable proxies, it must be one of the cipher suite names listed in: http://docs.oracle.com/javase/8/docs/technotes/guides/security/StandardNames.html#ciphersuites. For configurable proxies, it must follow the configuration specified in: https://commondatastorage.googleapis.com/chromium-boringssl-docs/ssl.h.html#Cipher-suite-configuration. This setting has no effect for configurable proxies when negotiating TLS 1.3.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "client_auth_enabled": {
              "description": "Enables two-way TLS.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enabled": {
              "description": "Enables TLS. If false, neither one-way nor two-way TLS will be enabled.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "enforce": {
              "description": "If true, TLS is strictly enforced.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ignore_validation_errors": {
              "description": "If true, Edge ignores TLS certificate errors. Valid when configuring TLS for target servers and target endpoints, and when configuring virtual hosts that use 2-way TLS. When used with a target endpoint/target server, if the backend system uses SNI and returns a cert with a subject Distinguished Name (DN) that does not match the hostname, there is no way to ignore the error and the connection fails.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "key_alias": {
              "description": "Required if clientAuthEnabled is true. The resource ID for the alias containing the private key and cert.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_store": {
              "description": "Required if clientAuthEnabled is true. The resource ID of the keystore.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocols": {
              "description": "The TLS versioins to be used.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "trust_store": {
              "description": "The resource ID of the truststore.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "common_name": {
              "block": {
                "attributes": {
                  "value": {
                    "description": "The TLS Common Name string of the certificate.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "wildcard_match": {
                    "description": "Indicates whether the cert should be matched against as a wildcard cert.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "The TLS Common Name of the certificate.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.",
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

func GoogleApigeeTargetServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeTargetServer), &result)
	return &result
}
