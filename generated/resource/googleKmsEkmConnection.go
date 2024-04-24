package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsEkmConnection = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The time at which the EkmConnection was created.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "crypto_space_path": {
        "computed": true,
        "description": "Optional. Identifies the EKM Crypto Space that this EkmConnection maps to. Note: This field is required if KeyManagementMode is CLOUD_KMS.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Optional. Etag of the currently stored EkmConnection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_management_mode": {
        "description": "Optional. Describes who can perform control plane operations on the EKM. If unset, this defaults to MANUAL Default value: \"MANUAL\" Possible values: [\"MANUAL\", \"CLOUD_KMS\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location for the EkmConnection.\nA full list of valid locations can be found by running 'gcloud kms locations list'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name for the EkmConnection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "service_resolvers": {
        "block": {
          "attributes": {
            "endpoint_filter": {
              "computed": true,
              "description": "Optional. The filter applied to the endpoints of the resolved service. If no filter is specified, all endpoints will be considered. An endpoint will be chosen arbitrarily from the filtered list for each request. For endpoint filter syntax and examples, see https://cloud.google.com/service-directory/docs/reference/rpc/google.cloud.servicedirectory.v1#resolveservicerequest.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hostname": {
              "description": "Required. The hostname of the EKM replica used at TLS and HTTP layers.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service_directory_service": {
              "description": "Required. The resource name of the Service Directory service pointing to an EKM replica, in the format projects/*/locations/*/namespaces/*/services/*",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "server_certificates": {
              "block": {
                "attributes": {
                  "issuer": {
                    "computed": true,
                    "description": "Output only. The issuer distinguished name in RFC 2253 format. Only present if parsed is true.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "not_after_time": {
                    "computed": true,
                    "description": "Output only. The certificate is not valid after this time. Only present if parsed is true.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "not_before_time": {
                    "computed": true,
                    "description": "Output only. The certificate is not valid before this time. Only present if parsed is true.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "parsed": {
                    "computed": true,
                    "description": "Output only. True if the certificate was parsed successfully.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "raw_der": {
                    "description": "Required. The raw certificate bytes in DER format. A base64-encoded string.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "serial_number": {
                    "computed": true,
                    "description": "Output only. The certificate serial number as a hex string. Only present if parsed is true.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "sha256_fingerprint": {
                    "computed": true,
                    "description": "Output only. The SHA-256 certificate fingerprint as a hex string. Only present if parsed is true.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "subject": {
                    "computed": true,
                    "description": "Output only. The subject distinguished name in RFC 2253 format. Only present if parsed is true.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "subject_alternative_dns_names": {
                    "computed": true,
                    "description": "Output only. The subject Alternative DNS names. Only present if parsed is true.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Required. A list of leaf server certificates used to authenticate HTTPS connections to the EKM replica. Currently, a maximum of 10 Certificate is supported.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently, only a single ServiceResolver is supported",
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

func GoogleKmsEkmConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsEkmConnection), &result)
	return &result
}
