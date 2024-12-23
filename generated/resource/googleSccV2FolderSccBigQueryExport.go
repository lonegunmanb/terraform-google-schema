package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSccV2FolderSccBigQueryExport = `{
  "block": {
    "attributes": {
      "big_query_export_id": {
        "description": "This must be unique within the organization.  It must consist of only lowercase letters,\nnumbers, and hyphens, must start with a letter, must end with either a letter or a number,\nand must be 63 characters or less.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time at which the BigQuery export was created. This field is set by the server and will be ignored if provided on export on creation.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "dataset": {
        "description": "The dataset to write findings' updates to.\nIts format is \"projects/[projectId]/datasets/[bigquery_dataset_id]\".\nBigQuery Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "The description of the notification config (max of 1024 characters).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter": {
        "description": "Expression that defines the filter to apply across create/update\nevents of findings. The\nexpression is a list of zero or more restrictions combined via\nlogical operators AND and OR. Parentheses are supported, and OR\nhas higher precedence than AND.\n\nRestrictions have the form \u003cfield\u003e \u003coperator\u003e \u003cvalue\u003e and may have\na - character in front of them to indicate negation. The fields\nmap to those defined in the corresponding resource.\n\nThe supported operators are:\n\n* = for all value types.\n* \u003e, \u003c, \u003e=, \u003c= for integer values.\n* :, meaning substring matching, for strings.\n\nThe supported value types are:\n\n* string literals in quotes.\n* integer literals without quotes.\n* boolean literals true and false without quotes.\n\nSee\n[Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications)\nfor information on how to write a filter.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "folder": {
        "description": "The folder where Cloud Security Command Center Big Query Export\nConfig lives in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The BigQuery export configuration is stored in this location. If not provided, Use global as default.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "most_recent_editor": {
        "computed": true,
        "description": "Email address of the user who last edited the BigQuery export.\nThis field is set by the server and will be ignored if provided on export creation or update.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of this export, in the format\n'folders/{{folder}}/locations/{{location}}/bigQueryExports/{{big_query_export_id}}'.\nThis field is provided in responses, and is ignored when provided in create requests.",
        "description_kind": "plain",
        "type": "string"
      },
      "principal": {
        "computed": true,
        "description": "The service account that needs permission to create table and upload data to the BigQuery dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The most recent time at which the BigQuery export was updated. This field is set by the server and will be ignored if provided on export creation or update.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
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

func GoogleSccV2FolderSccBigQueryExportSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSccV2FolderSccBigQueryExport), &result)
	return &result
}
