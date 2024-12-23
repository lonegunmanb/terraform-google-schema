package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleParallelstoreInstance = `{
  "block": {
    "attributes": {
      "access_points": {
        "computed": true,
        "description": "Output only. List of access_points.\nContains a list of IPv4 addresses used for client side configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "capacity_gib": {
        "description": "Required. Immutable. Storage capacity of Parallelstore instance in Gibibytes (GiB).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time when the instance was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "daos_version": {
        "computed": true,
        "description": "The version of DAOS software running in the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "The description of the instance. 2048 characters or less.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "directory_stripe_level": {
        "description": "Stripe level for directories.\nMIN when directory has a small number of files.\nMAX when directory has a large number of files.\n  Possible values:\n  DIRECTORY_STRIPE_LEVEL_UNSPECIFIED\n  DIRECTORY_STRIPE_LEVEL_MIN\n  DIRECTORY_STRIPE_LEVEL_BALANCED\n  DIRECTORY_STRIPE_LEVEL_MAX",
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
      "effective_reserved_ip_range": {
        "computed": true,
        "description": "Immutable. Contains the id of the allocated IP address\nrange associated with the private service access connection for example, \\\"test-default\\\"\nassociated with IP range 10.0.0.0/29. This field is populated by the service\nand contains the value currently used by the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "file_stripe_level": {
        "description": "Stripe level for files.\nMIN better suited for small size files.\nMAX higher throughput performance for larger files.\n  Possible values:\n  FILE_STRIPE_LEVEL_UNSPECIFIED\n  FILE_STRIPE_LEVEL_MIN\n  FILE_STRIPE_LEVEL_BALANCED\n  FILE_STRIPE_LEVEL_MAX",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "description": "The logical name of the Parallelstore instance in the user project with the following restrictions:\n  * Must contain only lowercase letters, numbers, and hyphens.\n  * Must start with a letter.\n  * Must be between 1-63 characters.\n  * Must end with a number or a letter.\n  * Must be unique within the customer project/ location",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "labels": {
        "description": "Cloud Labels are a flexible and lightweight mechanism for\norganizing cloud resources into groups that reflect a customer's organizational\nneeds and deployment strategies. Cloud Labels can be used to filter collections\nof resources. They can be used to control how resource metrics are aggregated.\nAnd they can be used as arguments to policy management rules (e.g. route, firewall,\nload balancing, etc.).\n\n* Label keys must be between 1 and 63 characters long and must conform to\n the following regular expression: 'a-z{0,62}'.\n* Label values must be between 0 and 63 characters long and must conform\n to the regular expression '[a-z0-9_-]{0,63}'.\n* No more than 64 labels can be associated with a given resource.\n\nSee https://goo.gl/xmQnxf for more information on and examples of labels.\n\nIf you plan to use labels in your own code, please note that additional\ncharacters may be allowed in the future. Therefore, you are advised to use\nan internal label representation, such as JSON, which doesn't rely upon\nspecific characters being disallowed.  For example, representing labels\nas the string:  'name + \"_\" + value' would prove problematic if we were to\nallow '\"_\"' in a future release. \"\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Part of 'parent'. See documentation of 'projectsId'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name of the instance, in the format\n'projects/{project}/locations/{location}/instances/{instance_id}'",
        "description_kind": "plain",
        "type": "string"
      },
      "network": {
        "description": "Immutable. The name of the Google Compute Engine [VPC network](https://cloud.google.com/vpc/docs/vpc)\nto which the instance is connected.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reserved_ip_range": {
        "description": "Immutable. Contains the id of the allocated IP address range\nassociated with the private service access connection for example, \\\"test-default\\\"\nassociated with IP range 10.0.0.0/29. If no range id is provided all ranges will\nbe considered.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The instance state.\n  Possible values:\n  STATE_UNSPECIFIED\n  CREATING\n  ACTIVE\n  DELETING\n  FAILED\n  UPGRADING",
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
        "description": "The time when the instance was updated.",
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

func GoogleParallelstoreInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleParallelstoreInstance), &result)
	return &result
}
