package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLustreInstance = `{
  "block": {
    "attributes": {
      "capacity_gib": {
        "description": "Required. The storage capacity of the instance in gibibytes (GiB). Allowed values\nare from 18000 to 954000, in increments of 9000.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Timestamp when the instance was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Optional. A user-readable description of the instance.",
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
      "filesystem": {
        "description": "Required. Immutable. The filesystem name for this instance. This name is used by client-side\ntools, including when mounting the instance. Must be 8 characters or less\nand may only contain letters and numbers.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "gke_support_enabled": {
        "description": "Optional. Indicates whether you want to enable support for GKE clients. By default,\nGKE clients are not supported.",
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
      "instance_id": {
        "description": "Required. The name of the Managed Lustre instance.\n\n* Must contain only lowercase letters, numbers, and hyphens.\n* Must start with a letter.\n* Must be between 1-63 characters.\n* Must end with a number or a letter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "labels": {
        "description": "Optional. Labels as key value pairs.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mount_point": {
        "computed": true,
        "description": "Output only. Mount point of the instance in the format 'IP_ADDRESS@tcp:/FILESYSTEM'.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The name of the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "network": {
        "description": "Required. Immutable. The full name of the VPC network to which the instance is connected.\nMust be in the format\n'projects/{project_id}/global/networks/{network_name}'.",
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
      "state": {
        "computed": true,
        "description": "Output only. The state of the instance.\nPossible values:\nSTATE_UNSPECIFIED\nACTIVE\nCREATING\nDELETING\nUPGRADING\nREPAIRING\nSTOPPED",
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
        "description": "Output only. Timestamp when the instance was last updated.",
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

func GoogleLustreInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLustreInstance), &result)
	return &result
}
