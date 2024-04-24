package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFilestoreInstance = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the instance.",
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
      "etag": {
        "computed": true,
        "description": "Server-specified ETag for the instance resource to prevent\nsimultaneous updates from overwriting each other.",
        "description_kind": "plain",
        "type": "string"
      },
      "file_shares": {
        "computed": true,
        "description": "File system shares on the instance. For this version, only a\nsingle file share is supported.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "capacity_gb": "number",
              "name": "string",
              "nfs_export_options": [
                "list",
                [
                  "object",
                  {
                    "access_mode": "string",
                    "anon_gid": "number",
                    "anon_uid": "number",
                    "ip_ranges": [
                      "list",
                      "string"
                    ],
                    "squash_mode": "string"
                  }
                ]
              ],
              "source_backup": "string"
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
      "kms_key_name": {
        "computed": true,
        "description": "KMS key name used for data encryption.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "Resource labels to represent user-provided metadata.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The name of the location of the instance. This can be a region for ENTERPRISE tier instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "networks": {
        "computed": true,
        "description": "VPC networks to which the instance is connected. For this version,\nonly a single network is supported.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connect_mode": "string",
              "ip_addresses": [
                "list",
                "string"
              ],
              "modes": [
                "list",
                "string"
              ],
              "network": "string",
              "reserved_ip_range": "string"
            }
          ]
        ]
      },
      "project": {
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
      "tier": {
        "computed": true,
        "description": "The service tier of the instance.\nPossible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD, ZONAL, REGIONAL and ENTERPRISE",
        "description_kind": "plain",
        "type": "string"
      },
      "zone": {
        "computed": true,
        "description": "The name of the Filestore zone of the instance.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleFilestoreInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFilestoreInstance), &result)
	return &result
}
