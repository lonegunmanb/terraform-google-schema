package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeStoragePoolTypes = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "deprecated": {
        "computed": true,
        "description": "The deprecation status associated with this storage pool type.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "deleted": "string",
              "deprecated": "string",
              "obsolete": "string",
              "replacement": "string",
              "state": "string"
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "The unique identifier for the resource. This identifier is defined by the server.",
        "description_kind": "plain",
        "type": "number"
      },
      "kind": {
        "computed": true,
        "description": "Type of the resource. Always compute#storagePoolType for storage pool types.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_pool_provisioned_capacity_gb": {
        "computed": true,
        "description": "Maximum storage pool size in GB.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_pool_provisioned_iops": {
        "computed": true,
        "description": "Maximum provisioned IOPS.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_pool_provisioned_throughput": {
        "computed": true,
        "description": "Maximum provisioned throughput.",
        "description_kind": "plain",
        "type": "number"
      },
      "min_pool_provisioned_capacity_gb": {
        "computed": true,
        "description": "Minimum storage pool size in GB.",
        "description_kind": "plain",
        "type": "number"
      },
      "min_pool_provisioned_iops": {
        "computed": true,
        "description": "Minimum provisioned IOPS.",
        "description_kind": "plain",
        "type": "number"
      },
      "min_pool_provisioned_throughput": {
        "computed": true,
        "description": "Minimum provisioned throughput.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "Name of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "Server-defined URL for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "self_link_with_id": {
        "computed": true,
        "description": "Server-defined URL for this resource with the resource id.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_pool_type": {
        "description": "Name of the storage pool type.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "supported_disk_types": {
        "computed": true,
        "description": "The list of disk types supported in this storage pool type.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "zone": {
        "description": "The name of the zone.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeStoragePoolTypesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeStoragePoolTypes), &result)
	return &result
}
