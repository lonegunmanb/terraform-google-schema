package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApphubApplication = `{
  "block": {
    "attributes": {
      "application_id": {
        "description": "Required. The Application identifier.",
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
        "description": "Optional. User-defined description of an Application.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Optional. User-defined name for the Application.",
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
        "description": "Part of 'parent'. See documentation of 'projectsId'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name of an Application. Format:\n\"projects/{host-project-id}/locations/{location}/applications/{application-id}\"",
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
        "description": "Output only. Application state. \n Possible values:\n STATE_UNSPECIFIED\nCREATING\nACTIVE\nDELETING",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. A universally unique identifier (in UUID4 format) for the 'Application'.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Update time.",
        "description_kind": "plain",
        "type": "string"
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
                    "description": "Optional. Contact's name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "email": {
                    "description": "Required. Email address of the contacts.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Optional. Business team that ensures user needs are met and value is delivered",
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
                    "description": "Optional. Contact's name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "email": {
                    "description": "Required. Email address of the contacts.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Optional. Developer team that owns development and coding.",
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
                    "description": "Optional. Contact's name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "email": {
                    "description": "Required. Email address of the contacts.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Optional. Operator team that ensures runtime and operations.",
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
      "scope": {
        "block": {
          "attributes": {
            "type": {
              "description": "Required. Scope Type. \n Possible values:\nREGIONAL\nGLOBAL Possible values: [\"REGIONAL\", \"GLOBAL\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Scope of an application.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleApphubApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApphubApplication), &result)
	return &result
}
