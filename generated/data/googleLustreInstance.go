package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLustreInstance = `{
  "block": {
    "attributes": {
      "capacity_gib": {
        "computed": true,
        "description": "The storage capacity of the instance in gibibytes (GiB). Allowed values\nare from '18000' to '954000', in increments of 9000.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Timestamp when the instance was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A user-readable description of the instance.",
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
      "filesystem": {
        "computed": true,
        "description": "The filesystem name for this instance. This name is used by client-side\ntools, including when mounting the instance. Must be eight characters or\nless and can only contain letters and numbers.",
        "description_kind": "plain",
        "type": "string"
      },
      "gke_support_enabled": {
        "computed": true,
        "description": "Indicates whether you want to enable support for GKE clients. By default,\nGKE clients are not supported.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "description": "The name of the Managed Lustre instance.\n\n* Must contain only lowercase letters, numbers, and hyphens.\n* Must start with a letter.\n* Must be between 1-63 characters.\n* Must end with a number or a letter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "Labels as key value pairs.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "computed": true,
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
        "description_kind": "plain",
        "type": "string"
      },
      "mount_point": {
        "computed": true,
        "description": "Mount point of the instance in the format 'IP_ADDRESS@tcp:/FILESYSTEM'.",
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
        "computed": true,
        "description": "The full name of the VPC network to which the instance is connected.\nMust be in the format\n'projects/{project_id}/global/networks/{network_name}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "per_unit_storage_throughput": {
        "computed": true,
        "description": "The throughput of the instance in MB/s/TiB.\nValid values are 125, 250, 500, 1000.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the instance.\nPossible values:\nSTATE_UNSPECIFIED\nACTIVE\nCREATING\nDELETING\nUPGRADING\nREPAIRING\nSTOPPED",
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
        "description": "Timestamp when the instance was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "zone": {
        "description": "Zone of Lustre instance",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
