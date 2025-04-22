package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeStoragePool = `{
  "block": {
    "attributes": {
      "capacity_provisioning_type": {
        "computed": true,
        "description": "Provisioning type of the byte capacity of the pool. Possible values: [\"STANDARD\", \"ADVANCED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "description": "Whether Terraform will be prevented from destroying the StoragePool.\nWhen the field is set to true or unset in Terraform state, a 'terraform apply'\nor 'terraform destroy' that would delete the StoragePool will fail.\nWhen the field is set to false, deleting the StoragePool is allowed.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "A description of this resource. Provide this property when you create the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "The unique identifier for the resource. This identifier is defined by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "kind": {
        "computed": true,
        "description": "Type of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "label_fingerprint": {
        "computed": true,
        "description": "The fingerprint used for optimistic locking of this resource.\nUsed internally during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is created.\nThe name must be 1-63 characters long, and comply with RFC1035.\nSpecifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'\nwhich means the first character must be a lowercase letter,\nand all following characters must be a dash, lowercase letter, or digit,\nexcept the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "performance_provisioning_type": {
        "computed": true,
        "description": "Provisioning type of the performance-related parameters of the pool, such as throughput and IOPS. Possible values: [\"STANDARD\", \"ADVANCED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pool_provisioned_capacity_gb": {
        "description": "Size, in GiB, of the storage pool. For more information about the size limits,\nsee https://cloud.google.com/compute/docs/disks/storage-pools.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pool_provisioned_iops": {
        "description": "Provisioned IOPS of the storage pool.\nOnly relevant if the storage pool type is 'hyperdisk-balanced'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pool_provisioned_throughput": {
        "description": "Provisioned throughput, in MB/s, of the storage pool.\nOnly relevant if the storage pool type is 'hyperdisk-balanced' or 'hyperdisk-throughput'.",
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
      "resource_status": {
        "computed": true,
        "description": "Status information for the storage pool resource.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disk_count": "string",
              "last_resize_timestamp": "string",
              "max_total_provisioned_disk_capacity_gb": "string",
              "pool_used_capacity_bytes": "string",
              "pool_used_iops": "string",
              "pool_used_throughput": "string",
              "pool_user_written_bytes": "string",
              "total_provisioned_disk_capacity_gb": "string",
              "total_provisioned_disk_iops": "string",
              "total_provisioned_disk_throughput": "string"
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description": "Status information for the storage pool resource.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disk_count": "string",
              "last_resize_timestamp": "string",
              "max_total_provisioned_disk_capacity_gb": "string",
              "pool_used_capacity_bytes": "string",
              "pool_used_iops": "string",
              "pool_used_throughput": "string",
              "pool_user_written_bytes": "string",
              "total_provisioned_disk_capacity_gb": "string",
              "total_provisioned_disk_iops": "string",
              "total_provisioned_disk_throughput": "string"
            }
          ]
        ]
      },
      "storage_pool_type": {
        "description": "Type of the storage pool. For example, the\nfollowing are valid values:\n\n* 'https://www.googleapis.com/compute/v1/projects/{project_id}/zones/{zone}/storagePoolTypes/hyperdisk-balanced'\n* 'hyperdisk-throughput'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone": {
        "computed": true,
        "description": "A reference to the zone where the storage pool resides.",
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

func GoogleComputeStoragePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeStoragePool), &result)
	return &result
}
