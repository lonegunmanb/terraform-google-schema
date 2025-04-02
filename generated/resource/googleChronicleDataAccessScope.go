package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleChronicleDataAccessScope = `{
  "block": {
    "attributes": {
      "allow_all": {
        "description": "Optional. Whether or not the scope allows all labels, allow_all and\nallowed_data_access_labels are mutually exclusive and one of them must be\npresent. denied_data_access_labels can still be used along with allow_all.\nWhen combined with denied_data_access_labels, access will be granted to all\ndata that doesn't have labels mentioned in denied_data_access_labels. E.g.:\nA customer with scope with denied labels A and B and allow_all will be able\nto see all data except data labeled with A and data labeled with B and data\nwith labels A and B.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "author": {
        "computed": true,
        "description": "Output only. The user who created the data access scope.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The time at which the data access scope was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_access_scope_id": {
        "description": "Required. The user provided scope id which will become the last part of the name\nof the scope resource.\nNeeds to be compliant with https://google.aip.dev/122",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "Optional. A description of the data access scope for a human reader.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "Output only. The name to be used for display to customers of the data access scope.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance": {
        "description": "The unique identifier for the Chronicle instance, which is the same as the customer ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_editor": {
        "computed": true,
        "description": "Output only. The user who last updated the data access scope.",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "The location of the resource. This is the geographical region where the Chronicle instance resides, such as \"us\" or \"europe-west2\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique full name of the data access scope. This unique identifier is generated using values provided for the URL parameters.\nFormat:\nprojects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{data_access_scope_id}",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time at which the data access scope was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "allowed_data_access_labels": {
        "block": {
          "attributes": {
            "asset_namespace": {
              "description": "The asset namespace configured in the forwarder\nof the customer's events.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_access_label": {
              "description": "The name of the data access label.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "display_name": {
              "computed": true,
              "description": "Output only. The display name of the label.\nData access label and log types's name\nwill match the display name of the resource.\nThe asset namespace will match the namespace itself.\nThe ingestion key value pair will match the key of the tuple.",
              "description_kind": "plain",
              "type": "string"
            },
            "log_type": {
              "description": "The name of the log type.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "ingestion_label": {
              "block": {
                "attributes": {
                  "ingestion_label_key": {
                    "description": "Required. The key of the ingestion label. Always required.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "ingestion_label_value": {
                    "description": "Optional. The value of the ingestion label. Optional. An object\nwith no provided value and some key provided would match\nagainst the given key and ANY value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Representation of an ingestion label type.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The allowed labels for the scope. There has to be at\nleast one label allowed for the scope to be valid.\nThe logical operator for evaluation of the allowed labels is OR.\nEither allow_all or allowed_data_access_labels needs to be provided.\nE.g.: A customer with scope with allowed labels A and B will be able\nto see data with labeled with A or B or (A and B).",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "denied_data_access_labels": {
        "block": {
          "attributes": {
            "asset_namespace": {
              "description": "The asset namespace configured in the forwarder\nof the customer's events.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_access_label": {
              "description": "The name of the data access label.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "display_name": {
              "computed": true,
              "description": "Output only. The display name of the label.\nData access label and log types's name\nwill match the display name of the resource.\nThe asset namespace will match the namespace itself.\nThe ingestion key value pair will match the key of the tuple.",
              "description_kind": "plain",
              "type": "string"
            },
            "log_type": {
              "description": "The name of the log type.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "ingestion_label": {
              "block": {
                "attributes": {
                  "ingestion_label_key": {
                    "description": "Required. The key of the ingestion label. Always required.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "ingestion_label_value": {
                    "description": "Optional. The value of the ingestion label. Optional. An object\nwith no provided value and some key provided would match\nagainst the given key and ANY value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Representation of an ingestion label type.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Optional. The denied labels for the scope.\nThe logical operator for evaluation of the denied labels is AND.\nE.g.: A customer with scope with denied labels A and B won't be able\nto see data labeled with A and data labeled with B\nand data with labels A and B.",
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

func GoogleChronicleDataAccessScopeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleChronicleDataAccessScope), &result)
	return &result
}
