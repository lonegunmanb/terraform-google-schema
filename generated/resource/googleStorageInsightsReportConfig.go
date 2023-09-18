package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageInsightsReportConfig = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "The editable display name of the inventory report configuration. Has a limit of 256 characters. Can be empty.",
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
        "description": "The location of the ReportConfig. The source and destination buckets specified in the ReportConfig\nmust be in the same location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The UUID of the inventory report configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "csv_options": {
        "block": {
          "attributes": {
            "delimiter": {
              "description": "The delimiter used to separate the fields in the inventory report CSV file.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "header_required": {
              "description": "The boolean that indicates whether or not headers are included in the inventory report CSV file.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "record_separator": {
              "description": "The character used to separate the records in the inventory report CSV file.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Options for configuring the format of the inventory report CSV file.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "frequency_options": {
        "block": {
          "attributes": {
            "frequency": {
              "description": "The frequency in which inventory reports are generated. Values are DAILY or WEEKLY. Possible values: [\"DAILY\", \"WEEKLY\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "end_date": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "The day of the month to stop generating inventory reports.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "month": {
                    "description": "The month to stop generating inventory reports.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "year": {
                    "description": "The year to stop generating inventory reports",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "The date to stop generating inventory reports. For example, {\"day\": 15, \"month\": 9, \"year\": 2022}.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "start_date": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "The day of the month to start generating inventory reports.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "month": {
                    "description": "The month to start generating inventory reports.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "year": {
                    "description": "The year to start generating inventory reports",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "The date to start generating inventory reports. For example, {\"day\": 15, \"month\": 8, \"year\": 2022}.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Options for configuring how inventory reports are generated.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "object_metadata_report_options": {
        "block": {
          "attributes": {
            "metadata_fields": {
              "description": "The metadata fields included in an inventory report.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "storage_destination_options": {
              "block": {
                "attributes": {
                  "bucket": {
                    "description": "The destination bucket that stores the generated inventory reports.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "destination_path": {
                    "description": "The path within the destination bucket to store generated inventory reports.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Options for where the inventory reports are stored.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "storage_filters": {
              "block": {
                "attributes": {
                  "bucket": {
                    "description": "The filter to use when specifying which bucket to generate inventory reports for.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A nested object resource",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Options for including metadata in an inventory report.",
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

func GoogleStorageInsightsReportConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageInsightsReportConfig), &result)
	return &result
}
