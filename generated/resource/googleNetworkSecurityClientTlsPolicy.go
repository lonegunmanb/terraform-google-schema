package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkSecurityClientTlsPolicy = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the ClientTlsPolicy was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A free-text description of the resource. Max length 1024 characters.",
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
        "description": "Set of label tags associated with the ClientTlsPolicy resource.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the client tls policy.\nThe default value is 'global'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the ClientTlsPolicy resource.",
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
      "sni": {
        "description": "Server Name Indication string to present to the server during TLS handshake. E.g: \"secure.example.com\".",
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
        "description": "Time the ClientTlsPolicy was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "client_certificate": {
        "block": {
          "block_types": {
            "certificate_provider_instance": {
              "block": {
                "attributes": {
                  "plugin_instance": {
                    "description": "Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to \"google_cloud_private_spiffe\" to use Certificate Authority Service certificate provider instance.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The certificate provider instance specification that will be passed to the data plane, which will be used to load necessary credential information.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "grpc_endpoint": {
              "block": {
                "attributes": {
                  "target_uri": {
                    "description": "The target URI of the gRPC endpoint. Only UDS path is supported, and should start with \"unix:\".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "gRPC specific configuration to access the gRPC server to obtain the cert and private key.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "server_validation_ca": {
        "block": {
          "block_types": {
            "certificate_provider_instance": {
              "block": {
                "attributes": {
                  "plugin_instance": {
                    "description": "Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to \"google_cloud_private_spiffe\" to use Certificate Authority Service certificate provider instance.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The certificate provider instance specification that will be passed to the data plane, which will be used to load necessary credential information.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "grpc_endpoint": {
              "block": {
                "attributes": {
                  "target_uri": {
                    "description": "The target URI of the gRPC endpoint. Only UDS path is supported, and should start with \"unix:\".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "gRPC specific configuration to access the gRPC server to obtain the cert and private key.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.",
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

func GoogleNetworkSecurityClientTlsPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkSecurityClientTlsPolicy), &result)
	return &result
}
