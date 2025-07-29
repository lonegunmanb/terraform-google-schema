package data

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
        "computed": true,
        "description": "An optional user-provided description for the dataset configuration with a maximum length of 256 characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "exclude_cloud_storage_buckets": {
        "computed": true,
        "description": "Defined the options for excluding cloud storage buckets for the DatasetConfig.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cloud_storage_buckets": [
                "list",
                [
                  "object",
                  {
                    "bucket_name": "string",
                    "bucket_prefix_regex": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "exclude_cloud_storage_locations": {
        "computed": true,
        "description": "Defines the options for excluding cloud storage locations for the DatasetConfig.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "locations": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identity": {
        "computed": true,
        "description": "Identity used by DatasetConfig.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "type": "string"
            }
          ]
        ]
      },
      "include_cloud_storage_buckets": {
        "computed": true,
        "description": "Defines the options for including cloud storage buckets for the DatasetConfig.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cloud_storage_buckets": [
                "list",
                [
                  "object",
                  {
                    "bucket_name": "string",
                    "bucket_prefix_regex": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "include_cloud_storage_locations": {
        "computed": true,
        "description": "Defines the options for including cloud storage locations for the DatasetConfig.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "locations": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "include_newly_created_buckets": {
        "computed": true,
        "description": "If set to true, the request includes all the newly created buckets in the dataset that meet the inclusion and exclusion rules.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "A boolean terraform only flag to link/unlink dataset.\n\nSetting this field to true while creation will automatically link the created dataset as an additional functionality.\n-\u003e **Note** A dataset config resource can only be destroyed once it is unlinked,\nso users must set this field to false to unlink the dataset and destroy the dataset config resource.",
        "description_kind": "plain",
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
        "type": "string"
      },
      "organization_scope": {
        "computed": true,
        "description": "Defines the options for providing a source organization for the DatasetConfig.",
        "description_kind": "plain",
        "type": "bool"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retention_period_days": {
        "computed": true,
        "description": "Number of days of history that must be retained.",
        "description_kind": "plain",
        "type": "number"
      },
      "source_folders": {
        "computed": true,
        "description": "Defines the options for providing source folders for the DatasetConfig.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "folder_numbers": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "source_projects": {
        "computed": true,
        "description": "Defines the options for providing source projects for the DatasetConfig.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "project_numbers": [
                "list",
                "string"
              ]
            }
          ]
        ]
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
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleStorageInsightsDatasetConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageInsightsDatasetConfig), &result)
	return &result
}
