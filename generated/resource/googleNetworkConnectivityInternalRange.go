package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkConnectivityInternalRange = `{
  "block": {
    "attributes": {
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
      "exclude_cidr_ranges": {
        "description": "Optional. List of IP CIDR ranges to be excluded. Resulting reserved Internal Range will not overlap with any CIDR blocks mentioned in this list.\nOnly IPv4 CIDR ranges are supported.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "immutable": {
        "description": "Immutable ranges cannot have their fields modified, except for labels and description.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ip_cidr_range": {
        "computed": true,
        "description": "The IP range that this internal range defines.\nNOTE: IPv6 ranges are limited to usage=EXTERNAL_TO_VPC and peering=FOR_SELF\nNOTE: For IPv6 Ranges this field is compulsory, i.e. the address range must be specified explicitly.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "User-defined labels.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The name of the policy based route.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "Fully-qualified URL of the network that this route applies to, for example: projects/my-project/global/networks/my-network.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "overlaps": {
        "description": "Optional. Types of resources that are allowed to overlap with the current internal range. Possible values: [\"OVERLAP_ROUTE_RANGE\", \"OVERLAP_EXISTING_SUBNET_RANGE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "peering": {
        "description": "The type of peering set for this internal range. Possible values: [\"FOR_SELF\", \"FOR_PEER\", \"NOT_SHARED\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "prefix_length": {
        "description": "An alternate to ipCidrRange. Can be set when trying to create a reservation that automatically finds a free range of the given size.\nIf both ipCidrRange and prefixLength are set, there is an error if the range sizes do not match. Can also be used during updates to change the range size.\nNOTE: For IPv6 this field only works if ip_cidr_range is set as well, and both fields must match. In other words, with IPv6 this field only works as\na redundant parameter.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_cidr_range": {
        "description": "Optional. Can be set to narrow down or pick a different address space while searching for a free range.\nIf not set, defaults to the \"10.0.0.0/8\" address space. This can be used to search in other rfc-1918 address spaces like \"172.16.0.0/12\" and \"192.168.0.0/16\" or non-rfc-1918 address spaces used in the VPC.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
      "usage": {
        "description": "The type of usage set for this InternalRange. Possible values: [\"FOR_VPC\", \"EXTERNAL_TO_VPC\", \"FOR_MIGRATION\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "users": {
        "computed": true,
        "description": "Output only. The list of resources that refer to this internal range.\nResources that use the internal range for their range allocation are referred to as users of the range.\nOther resources mark themselves as users while doing so by creating a reference to this internal range. Having a user, based on this reference, prevents deletion of the internal range referred to. Can be empty.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "allocation_options": {
        "block": {
          "attributes": {
            "allocation_strategy": {
              "description": "Optional. Sets the strategy used to automatically find a free range of a size given by prefixLength. Can be set only when trying to create a reservation that automatically finds the free range to reserve. Possible values: [\"RANDOM\", \"FIRST_AVAILABLE\", \"RANDOM_FIRST_N_AVAILABLE\", \"FIRST_SMALLEST_FITTING\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "first_available_ranges_lookup_size": {
              "description": "Must be set when allocation_strategy is RANDOM_FIRST_N_AVAILABLE, otherwise must remain unset. Defines the size of the set of free ranges from which RANDOM_FIRST_N_AVAILABLE strategy randomy selects one,\nin other words it sets the N in the RANDOM_FIRST_N_AVAILABLE.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Options for automatically allocating a free range with a size given by prefixLength.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "migration": {
        "block": {
          "attributes": {
            "source": {
              "description": "Resource path as an URI of the source resource, for example a subnet.\nThe project for the source resource should match the project for the\nInternalRange.\nAn example /projects/{project}/regions/{region}/subnetworks/{subnet}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target": {
              "description": "Resource path of the target resource. The target project can be\ndifferent, as in the cases when migrating to peer networks. The resource\nmay not exist yet.\nFor example /projects/{project}/regions/{region}/subnetworks/{subnet}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Specification for migration with source and target resource names.",
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

func GoogleNetworkConnectivityInternalRangeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkConnectivityInternalRange), &result)
	return &result
}
