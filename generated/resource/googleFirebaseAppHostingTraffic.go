package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleFirebaseAppHostingTraffic = `{
  "block": {
    "attributes": {
      "backend": {
        "description": "Id of the backend that this Traffic config applies to",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Time at which the backend was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "current": {
        "computed": true,
        "description": "Current state of traffic allocation for the backend.\nWhen setting 'target', this field may differ for some time until the desired state is reached.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "splits": [
                "list",
                [
                  "object",
                  {
                    "build": "string",
                    "percent": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "delete_time": {
        "computed": true,
        "description": "Time at which the backend was deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Server-computed checksum based on other values; may be sent\non update or delete to ensure operation is done on expected resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location the Backend that this Traffic config applies to",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name of the backend traffic config\n\nFormat:\n\n'projects/{project}/locations/{locationId}/backends/{backendId}/traffic'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "System-assigned, unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time at which the backend was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "rollout_policy": {
        "block": {
          "attributes": {
            "codebase_branch": {
              "description": "Specifies a branch that triggers a new build to be started with this\npolicy. If not set, no automatic rollouts will happen.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disabled": {
              "description": "A flag that, if true, prevents rollouts from being created via this RolloutPolicy.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "disabled_time": {
              "computed": true,
              "description": "If disabled is set, the time at which the rollouts were disabled.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "The policy for how builds and rollouts are triggered and rolled out.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "target": {
        "block": {
          "block_types": {
            "splits": {
              "block": {
                "attributes": {
                  "build": {
                    "description": "The build that traffic is being routed to.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "percent": {
                    "description": "The percentage of traffic to send to the build. Currently must be 100 or 0.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "A list of traffic splits that together represent where traffic is being routed.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Set to manually control the desired traffic for the backend. This will\ncause current to eventually match this value. The percentages must add\nup to 100.",
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

func GoogleFirebaseAppHostingTrafficSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleFirebaseAppHostingTraffic), &result)
	return &result
}
