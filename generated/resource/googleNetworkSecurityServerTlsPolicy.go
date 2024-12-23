package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkSecurityServerTlsPolicy = `{
  "block": {
    "attributes": {
      "allow_open": {
        "description": "This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer policies.\nDetermines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.\nConsider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS and non-TLS traffic reaching port :80.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "create_time": {
        "computed": true,
        "description": "Time the ServerTlsPolicy was created in UTC.",
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
        "description": "Set of label tags associated with the ServerTlsPolicy resource.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the server tls policy.\nThe default value is 'global'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the ServerTlsPolicy resource.",
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
        "description": "Time the ServerTlsPolicy was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "mtls_policy": {
        "block": {
          "attributes": {
            "client_validation_mode": {
              "description": "When the client presents an invalid certificate or no certificate to the load balancer, the clientValidationMode specifies how the client connection is handled.\nRequired if the policy is to be used with the external HTTPS load balancing. For Traffic Director it must be empty. Possible values: [\"CLIENT_VALIDATION_MODE_UNSPECIFIED\", \"ALLOW_INVALID_OR_MISSING_CLIENT_CERT\", \"REJECT_INVALID\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_validation_trust_config": {
              "description": "Reference to the TrustConfig from certificatemanager.googleapis.com namespace.\nIf specified, the chain validation will be performed against certificates configured in the given TrustConfig.\nAllowed only if the policy is to be used with external HTTPS load balancers.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "client_validation_ca": {
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
                      "description": "Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty.\nDefines a mechanism to provision server identity (public and private keys). Cannot be combined with allowOpen as a permissive mode that allows both plain text and TLS is not supported.",
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
                "description": "Required if the policy is to be used with Traffic Director. For external HTTPS load balancers it must be empty.\nDefines the mechanism to obtain the Certificate Authority certificate to validate the client certificate.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic Director.\nDefines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "server_certificate": {
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
                "description": "Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty.\nDefines a mechanism to provision server identity (public and private keys). Cannot be combined with allowOpen as a permissive mode that allows both plain text and TLS is not supported.",
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

func GoogleNetworkSecurityServerTlsPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkSecurityServerTlsPolicy), &result)
	return &result
}
