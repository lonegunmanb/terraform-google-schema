package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApphubWorkload = `{
  "block": {
    "attributes": {
      "application_id": {
        "description": "Part of 'parent'.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Create time.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "User-defined description of a Workload.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "discovered_workload": {
        "description": "Immutable. The resource name of the original discovered workload.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description": "User-defined name for the Workload.",
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
      "location": {
        "description": "Part of 'parent'.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name of the Workload. Format:\"projects/{host-project-id}/locations/{location}/applications/{application-id}/workloads/{workload-id}\"",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. Workload state. Possible values:  STATE_UNSPECIFIED CREATING ACTIVE DELETING DETACHED",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. A universally unique identifier (UUID) for the 'Workload' in the UUID4 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Update time.",
        "description_kind": "plain",
        "type": "string"
      },
      "workload_id": {
        "description": "The Workload identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "workload_properties": {
        "computed": true,
        "description": "Properties of an underlying compute resource represented by the Workload.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "gcp_project": "string",
              "location": "string",
              "zone": "string"
            }
          ]
        ]
      },
      "workload_reference": {
        "computed": true,
        "description": "Reference of an underlying compute resource represented by the Workload.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "uri": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "attributes": {
        "block": {
          "block_types": {
            "business_owners": {
              "block": {
                "attributes": {
                  "display_name": {
                    "description": "Contact's name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "email": {
                    "description": "Email address of the contacts.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Business team that ensures user needs are met and value is delivered",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "criticality": {
              "block": {
                "attributes": {
                  "type": {
                    "description": "Criticality type. Possible values: [\"MISSION_CRITICAL\", \"HIGH\", \"MEDIUM\", \"LOW\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Criticality of the Application, Service, or Workload",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "developer_owners": {
              "block": {
                "attributes": {
                  "display_name": {
                    "description": "Contact's name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "email": {
                    "description": "Email address of the contacts.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Developer team that owns development and coding.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "environment": {
              "block": {
                "attributes": {
                  "type": {
                    "description": "Environment type. Possible values: [\"PRODUCTION\", \"STAGING\", \"TEST\", \"DEVELOPMENT\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Environment of the Application, Service, or Workload",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "operator_owners": {
              "block": {
                "attributes": {
                  "display_name": {
                    "description": "Contact's name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "email": {
                    "description": "Email address of the contacts.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Operator team that ensures runtime and operations.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Consumer provided attributes.",
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

func GoogleApphubWorkloadSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApphubWorkload), &result)
	return &result
}
