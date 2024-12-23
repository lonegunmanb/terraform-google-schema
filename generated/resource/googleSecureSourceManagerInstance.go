package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecureSourceManagerInstance = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the Instance was created in UTC.",
        "description_kind": "plain",
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
      "host_config": {
        "computed": true,
        "description": "A list of hostnames for this instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "api": "string",
              "git_http": "string",
              "git_ssh": "string",
              "html": "string"
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
      "instance_id": {
        "description": "The name for the Instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key": {
        "description": "Customer-managed encryption key name, in the format projects/*/locations/*/keyRings/*/cryptoKeys/*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels as key value pairs.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the Instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name for the Instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the Instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_note": {
        "computed": true,
        "description": "Provides information about the current instance state.",
        "description_kind": "plain",
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
        "description": "Time the Instance was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "private_config": {
        "block": {
          "attributes": {
            "ca_pool": {
              "description": "CA pool resource, resource must in the format of 'projects/{project}/locations/{location}/caPools/{ca_pool}'.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "http_service_attachment": {
              "computed": true,
              "description": "Service Attachment for HTTP, resource is in the format of 'projects/{project}/regions/{region}/serviceAttachments/{service_attachment}'.",
              "description_kind": "plain",
              "type": "string"
            },
            "is_private": {
              "description": "'Indicate if it's private instance.'",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "ssh_service_attachment": {
              "computed": true,
              "description": "Service Attachment for SSH, resource is in the format of 'projects/{project}/regions/{region}/serviceAttachments/{service_attachment}'.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Private settings for private instance.",
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
      "workforce_identity_federation_config": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "'Whether Workforce Identity Federation is enabled.'",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Configuration for Workforce Identity Federation to support third party identity provider.\nIf unset, defaults to the Google OIDC IdP.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSecureSourceManagerInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSecureSourceManagerInstance), &result)
	return &result
}
