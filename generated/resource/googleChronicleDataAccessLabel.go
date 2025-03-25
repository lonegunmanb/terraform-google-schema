package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleChronicleDataAccessLabel = `{
  "block": {
    "attributes": {
      "author": {
        "computed": true,
        "description": "Output only. The user who created the data access label.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The time at which the data access label was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_access_label_id": {
        "description": "Required. The ID to use for the data access label, which will become the label's\ndisplay name and the final component of the label's resource name. The\nmaximum number of characters should be 63. Regex pattern is as per AIP:\nhttps://google.aip.dev/122#resource-id-segments",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "Optional. A description of the data access label for a human reader.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "Output only. The short name displayed for the label as it appears on event data. This is same as data access label id.",
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
        "description": "Output only. The user who last updated the data access label.",
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
        "description": "The unique resource name of the data access label. This unique identifier is generated using values provided for the URL parameters.\nFormat:\nprojects/{project}/locations/{location}/instances/{instance}/dataAccessLabels/{data_access_label_id}",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "udm_query": {
        "description": "A UDM query over event data.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time at which the data access label was last updated.",
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

func GoogleChronicleDataAccessLabelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleChronicleDataAccessLabel), &result)
	return &result
}
