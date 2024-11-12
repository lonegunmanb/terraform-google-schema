package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetappStoragePool = `{
  "block": {
    "attributes": {
      "active_directory": {
        "description": "Specifies the Active Directory policy to be used. Format: 'projects/{{project}}/locations/{{location}}/activeDirectories/{{name}}'.\nThe policy needs to be in the same location as the storage pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "allow_auto_tiering": {
        "description": "Optional. True if the storage pool supports Auto Tiering enabled volumes. Default is false.\nAuto-tiering can be enabled after storage pool creation but it can't be disabled once enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "capacity_gib": {
        "description": "Capacity of the storage pool (in GiB).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
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
      "encryption_type": {
        "computed": true,
        "description": "Reports if volumes in the pool are encrypted using a Google-managed encryption key or CMEK.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_config": {
        "description": "Specifies the CMEK policy to be used for volume encryption. Format: 'projects/{{project}}/locations/{{location}}/kmsConfigs/{{name}}'.\nThe policy needs to be in the same location as the storage pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels as key value pairs. Example: '{ \"owner\": \"Bob\", \"department\": \"finance\", \"purpose\": \"testing\" }'.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "ldap_enabled": {
        "description": "When enabled, the volumes uses Active Directory as LDAP name service for UID/GID lookups. Required to enable extended group support for NFSv3,\nusing security identifiers for NFSv4.1 or principal names for kerberized NFSv4.1.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "description": "Name of the location. For zonal Flex pools specify a zone name, in all other cases a region name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the storage pool. Needs to be unique per location/region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "VPC network name with format: 'projects/{{project}}/global/networks/{{network}}'",
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
      "replica_zone": {
        "description": "Specifies the replica zone for regional Flex pools. 'zone' and 'replica_zone' values can be swapped to initiate a\n[zone switch](https://cloud.google.com/netapp/volumes/docs/configure-and-use/storage-pools/edit-or-delete-storage-pool#switch_active_and_replica_zones).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_level": {
        "description": "Service level of the storage pool. Possible values: [\"PREMIUM\", \"EXTREME\", \"STANDARD\", \"FLEX\"]",
        "description_kind": "plain",
        "required": true,
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
      "volume_capacity_gib": {
        "computed": true,
        "description": "Size allocated to volumes in the storage pool (in GiB).",
        "description_kind": "plain",
        "type": "string"
      },
      "volume_count": {
        "computed": true,
        "description": "Number of volume in the storage pool.",
        "description_kind": "plain",
        "type": "number"
      },
      "zone": {
        "description": "Specifies the active zone for regional Flex pools. 'zone' and 'replica_zone' values can be swapped to initiate a\n[zone switch](https://cloud.google.com/netapp/volumes/docs/configure-and-use/storage-pools/edit-or-delete-storage-pool#switch_active_and_replica_zones).\nIf you want to create a zonal Flex pool, specify a zone name for 'location' and omit 'zone'.",
        "description_kind": "plain",
        "optional": true,
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

func GoogleNetappStoragePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetappStoragePool), &result)
	return &result
}
