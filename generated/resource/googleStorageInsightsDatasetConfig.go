package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageInsightsDatasetConfig = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The UTC time at which the DatasetConfig was created. This is auto-populated.",
        "description_kind": "plain",
        "type": "string"
      },
      "dataset_config_id": {
        "description": "The user-defined ID of the DatasetConfig",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dataset_config_state": {
        "computed": true,
        "description": "State of the DatasetConfig.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional user-provided description for the dataset configuration with a maximum length of 256 characters.",
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
      "include_newly_created_buckets": {
        "description": "If set to true, the request includes all the newly created buckets in the dataset that meet the inclusion and exclusion rules.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "link": {
        "computed": true,
        "description": "Details of the linked DatasetConfig.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dataset": "string",
              "linked": "bool"
            }
          ]
        ]
      },
      "link_dataset": {
        "description": "A boolean terraform only flag to link/unlink dataset.\n\nSetting this field to true while creation will automatically link the created dataset as an additional functionality.\n-\u003e **Note** A dataset config resource can only be destroyed once it is unlinked,\nso users must set this field to false to unlink the dataset and destroy the dataset config resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "description": "The location of the DatasetConfig.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The full canonical resource name of the DatasetConfig (e.g., projects/P/locations/L/datasetConfigs/ID).",
        "description_kind": "plain",
        "type": "string"
      },
      "organization_number": {
        "computed": true,
        "description": "Organization resource ID that the source projects should belong to.\nProjects that do not belong to the provided organization are not considered when creating the dataset.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "organization_scope": {
        "description": "Defines the options for providing a source organization for the DatasetConfig.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retention_period_days": {
        "description": "Number of days of history that must be retained.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "uid": {
        "computed": true,
        "description": "System generated unique identifier for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The UTC time at which the DatasetConfig was updated. This is auto-populated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "exclude_cloud_storage_buckets": {
        "block": {
          "block_types": {
            "cloud_storage_buckets": {
              "block": {
                "attributes": {
                  "bucket_name": {
                    "description": "The list of cloud storage bucket names to exclude in the DatasetConfig.\nExactly one of the bucket_name and bucket_prefix_regex should be specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "bucket_prefix_regex": {
                    "description": "The list of regex patterns for bucket names matching the regex.\nRegex should follow the syntax specified in google/re2 on GitHub.\nExactly one of the bucket_name and bucket_prefix_regex should be specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The list of cloud storage buckets/bucket prefix regexes to exclude in the DatasetConfig.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Defined the options for excluding cloud storage buckets for the DatasetConfig.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "exclude_cloud_storage_locations": {
        "block": {
          "attributes": {
            "locations": {
              "description": "The list of cloud storage locations to exclude in the DatasetConfig.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Defines the options for excluding cloud storage locations for the DatasetConfig.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "identity": {
        "block": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "Name of the identity.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "description": "Type of identity to use for the DatasetConfig. Possible values: [\"IDENTITY_TYPE_PER_CONFIG\", \"IDENTITY_TYPE_PER_PROJECT\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Identity used by DatasetConfig.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "include_cloud_storage_buckets": {
        "block": {
          "block_types": {
            "cloud_storage_buckets": {
              "block": {
                "attributes": {
                  "bucket_name": {
                    "description": "The list of cloud storage bucket names to include in the DatasetConfig.\nExactly one of the bucket_name and bucket_prefix_regex should be specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "bucket_prefix_regex": {
                    "description": "The list of regex patterns for bucket names matching the regex.\nRegex should follow the syntax specified in google/re2 on GitHub.\nExactly one of the bucket_name and bucket_prefix_regex should be specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The list of cloud storage buckets/bucket prefix regexes to include in the DatasetConfig.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Defines the options for including cloud storage buckets for the DatasetConfig.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "include_cloud_storage_locations": {
        "block": {
          "attributes": {
            "locations": {
              "description": "The list of cloud storage locations to include in the DatasetConfig.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Defines the options for including cloud storage locations for the DatasetConfig.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_folders": {
        "block": {
          "attributes": {
            "folder_numbers": {
              "description": "The list of folder numbers to include in the DatasetConfig.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Defines the options for providing source folders for the DatasetConfig.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_projects": {
        "block": {
          "attributes": {
            "project_numbers": {
              "description": "The list of project numbers to include in the DatasetConfig.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Defines the options for providing source projects for the DatasetConfig.",
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

func GoogleStorageInsightsDatasetConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageInsightsDatasetConfig), &result)
	return &result
}
