package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCertificateManagerTrustConfig = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The creation timestamp of a TrustConfig.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "One or more paragraphs of text description of a trust config.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "labels": {
        "description": "Set of label tags associated with the trust config.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The trust config location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "A user-defined name of the trust config. Trust config names must be unique globally.",
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
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "The last update timestamp of a TrustConfig.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "allowlisted_certificates": {
        "block": {
          "attributes": {
            "pem_certificate": {
              "description": "PEM certificate that is allowlisted. The certificate can be up to 5k bytes, and must be a parseable X.509 certificate.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Allowlisted PEM-encoded certificates. A certificate matching an allowlisted certificate is always considered valid as long as\nthe certificate is parseable, proof of private key possession is established, and constraints on the certificate's SAN field are met.",
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
      },
      "trust_stores": {
        "block": {
          "block_types": {
            "intermediate_cas": {
              "block": {
                "attributes": {
                  "pem_certificate": {
                    "description": "PEM intermediate certificate used for building up paths for validation.\nEach certificate provided in PEM format may occupy up to 5kB.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  }
                },
                "description": "Set of intermediate CA certificates used for the path building phase of chain validation.\nThe field is currently not supported if trust config is used for the workload certificate feature.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "trust_anchors": {
              "block": {
                "attributes": {
                  "pem_certificate": {
                    "description": "PEM root certificate of the PKI used for validation.\nEach certificate provided in PEM format may occupy up to 5kB.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  }
                },
                "description": "List of Trust Anchors to be used while performing validation against a given TrustStore.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Set of trust stores to perform validation against.\nThis field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCertificateManagerTrustConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCertificateManagerTrustConfig), &result)
	return &result
}
