package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryRowAccessPolicy = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time when this row access policy was created, in milliseconds since\nthe epoch.",
        "description_kind": "plain",
        "type": "string"
      },
      "dataset_id": {
        "description": "The ID of the dataset containing this row access policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filter_predicate": {
        "description": "A SQL boolean expression that represents the rows defined by this row\naccess policy, similar to the boolean expression in a WHERE clause of a\nSELECT query on a table.\nReferences to other tables, routines, and temporary functions are not\nsupported.\n\nExamples: region=\"EU\"\ndate_field = CAST('2019-9-27' as DATE)\nnullable_field is not NULL\nnumeric_field BETWEEN 1.0 AND 5.0",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "grantees": {
        "description": "Input only. The optional list of iam_member users or groups that specifies the initial\nmembers that the row-level access policy should be created with.\n\ngrantees types:\n- \"user:alice@example.com\": An email address that represents a specific\nGoogle account.\n- \"serviceAccount:my-other-app@appspot.gserviceaccount.com\": An email\naddress that represents a service account.\n- \"group:admins@example.com\": An email address that represents a Google\ngroup.\n- \"domain:example.com\":The Google Workspace domain (primary) that\nrepresents all the users of that domain.\n- \"allAuthenticatedUsers\": A special identifier that represents all service\naccounts and all users on the internet who have authenticated with a Google\nAccount. This identifier includes accounts that aren't connected to a\nGoogle Workspace or Cloud Identity domain, such as personal Gmail accounts.\nUsers who aren't authenticated, such as anonymous visitors, aren't\nincluded.\n- \"allUsers\":A special identifier that represents anyone who is on\nthe internet, including authenticated and unauthenticated users. Because\nBigQuery requires authentication before a user can access the service,\nallUsers includes only authenticated users.",
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
      "last_modified_time": {
        "computed": true,
        "description": "The time when this row access policy was last modified, in milliseconds\nsince the epoch.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_id": {
        "description": "The ID of the row access policy. The ID must contain only\nletters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum\nlength is 256 characters.",
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
      "table_id": {
        "description": "The ID of the table containing this row access policy.",
        "description_kind": "plain",
        "required": true,
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

func GoogleBigqueryRowAccessPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryRowAccessPolicy), &result)
	return &result
}
