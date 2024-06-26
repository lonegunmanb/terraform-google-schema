package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigtableGcPolicy = `{
  "block": {
    "attributes": {
      "column_family": {
        "description": "The name of the column family.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "deletion_policy": {
        "description": "The deletion policy for the GC policy. Setting ABANDON allows the resource\n\t\t\t\tto be abandoned rather than deleted. This is useful for GC policy as it cannot be deleted\n\t\t\t\tin a replicated instance. Possible values are: \"ABANDON\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gc_rules": {
        "description": "Serialized JSON string for garbage collection policy. Conflicts with \"mode\", \"max_age\" and \"max_version\".",
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
      "ignore_warnings": {
        "description": "Allows ignoring warnings when updating the GC policy. This can be used\n\t\t\t\tto increase the gc policy on replicated clusters. Doing this may make clusters be\n\t\t\t\tinconsistent for a longer period of time, before using this make sure you understand\n\t\t\t\tthe risks listed at https://cloud.google.com/bigtable/docs/garbage-collection#increasing",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "instance_name": {
        "description": "The name of the Bigtable instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mode": {
        "description": "NOTE: 'gc_rules' is more flexible, and should be preferred over this field for new resources. This field may be deprecated in the future. If multiple policies are set, you should choose between UNION OR INTERSECTION.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "table": {
        "description": "The name of the table.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "max_age": {
        "block": {
          "attributes": {
            "days": {
              "computed": true,
              "deprecated": true,
              "description": "Number of days before applying GC policy.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "duration": {
              "computed": true,
              "description": "Duration before applying GC policy",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "NOTE: 'gc_rules' is more flexible, and should be preferred over this field for new resources. This field may be deprecated in the future. GC policy that applies to all cells older than the given age.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "max_version": {
        "block": {
          "attributes": {
            "number": {
              "description": "Number of version before applying the GC policy.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "NOTE: 'gc_rules' is more flexible, and should be preferred over this field for new resources. This field may be deprecated in the future. GC policy that applies to all versions of a cell except for the most recent.",
          "description_kind": "plain"
        },
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

func GoogleBigtableGcPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigtableGcPolicy), &result)
	return &result
}
