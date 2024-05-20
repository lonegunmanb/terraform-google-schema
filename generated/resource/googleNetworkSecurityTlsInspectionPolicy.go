package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkSecurityTlsInspectionPolicy = `{
  "block": {
    "attributes": {
      "ca_pool": {
        "description": "A CA pool resource used to issue interception certificates.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp when the resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_tls_features": {
        "description": "List of custom TLS cipher suites selected. This field is valid only if the selected tls_feature_profile is CUSTOM. The compute.SslPoliciesService.ListAvailableFeatures method returns the set of features that can be specified in this list. Note that Secure Web Proxy does not yet honor this field.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "description": "Free-text description of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclude_public_ca_set": {
        "description": "If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the tls inspection policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "min_tls_version": {
        "description": "Minimum TLS version that the firewall should use when negotiating connections with both clients and servers. If this is not set, then the default value is to allow the broadest set of clients and servers (TLS 1.0 or higher). Setting this to more restrictive values may improve security, but may also prevent the firewall from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field. Possible values: [\"TLS_VERSION_UNSPECIFIED\", \"TLS_1_0\", \"TLS_1_1\", \"TLS_1_2\", \"TLS_1_3\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Short name of the TlsInspectionPolicy resource to be created.",
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
      "tls_feature_profile": {
        "description": "The selected Profile. If this is not set, then the default value is to allow the broadest set of clients and servers (\\\"PROFILE_COMPATIBLE\\\"). Setting this to more restrictive values may improve security, but may also prevent the TLS inspection proxy from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field. Possible values: [\"PROFILE_UNSPECIFIED\", \"PROFILE_COMPATIBLE\", \"PROFILE_MODERN\", \"PROFILE_RESTRICTED\", \"PROFILE_CUSTOM\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "trust_config": {
        "description": "A TrustConfig resource used when making a connection to the TLS server. This is a relative resource path following the form \\\"projects/{project}/locations/{location}/trustConfigs/{trust_config}\\\". This is necessary to intercept TLS connections to servers with certificates signed by a private CA or self-signed certificates. Trust config and the TLS inspection policy must be in the same region. Note that Secure Web Proxy does not yet honor this field.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp when the resource was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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

func GoogleNetworkSecurityTlsInspectionPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkSecurityTlsInspectionPolicy), &result)
	return &result
}
