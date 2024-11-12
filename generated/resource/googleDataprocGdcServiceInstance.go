package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataprocGdcServiceInstance = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The timestamp when the resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "User-provided human-readable name to be used in user interfaces.",
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
      "effective_service_account": {
        "computed": true,
        "description": "Effective service account associated with ServiceInstance. This will be the service_account if specified. Otherwise, it will be an automatically created per-resource P4SA that also automatically has Fleet Workload. Identity bindings applied.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "The labels to associate with this service instance. Labels may be used for filtering and billing tracking. \n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Location of the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The name of the service instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "Whether the service instance is currently reconciling. True if the current state of the resource does not match the intended state, and the system is working to reconcile them, whether or not the change was user initiated.",
        "description_kind": "plain",
        "type": "bool"
      },
      "requested_state": {
        "computed": true,
        "description": "The intended state to which the service instance is reconciling. Possible values:\n* 'CREATING'\n* 'ACTIVE'\n* 'DISCONNECTED'\n* 'DELETING'\n* 'STOPPING'\n* 'STOPPED'\n* 'STARTING'\n* 'UPDATING'\n* 'FAILED'",
        "description_kind": "plain",
        "type": "string"
      },
      "service_account": {
        "description": "Requested service account to associate with ServiceInstance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_instance_id": {
        "description": "Id of the service instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state. Possible values:\n* 'CREATING'\n* 'ACTIVE'\n* 'DISCONNECTED'\n* 'DELETING'\n* 'STOPPING'\n* 'STOPPED'\n* 'STARTING'\n* 'UPDATING'\n* 'FAILED'",
        "description_kind": "plain",
        "type": "string"
      },
      "state_message": {
        "computed": true,
        "description": "A message explaining the current state.",
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
      "uid": {
        "computed": true,
        "description": "System generated unique identifier for this service instance, formatted as UUID4.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp when the resource was most recently updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "gdce_cluster": {
        "block": {
          "attributes": {
            "gdce_cluster": {
              "description": "Gdce cluster resource id.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Gdce cluster information.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "spark_service_instance_config": {
        "block": {
          "description": "Spark-specific service instance configuration.",
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

func GoogleDataprocGdcServiceInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataprocGdcServiceInstance), &result)
	return &result
}
