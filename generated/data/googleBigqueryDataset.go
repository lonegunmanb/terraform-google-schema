package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryDataset = `{
  "block": {
    "attributes": {
      "access": {
        "computed": true,
        "description": "An array of objects that define dataset access for one or more entities.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "condition": [
                "list",
                [
                  "object",
                  {
                    "description": "string",
                    "expression": "string",
                    "location": "string",
                    "title": "string"
                  }
                ]
              ],
              "dataset": [
                "list",
                [
                  "object",
                  {
                    "dataset": [
                      "list",
                      [
                        "object",
                        {
                          "dataset_id": "string",
                          "project_id": "string"
                        }
                      ]
                    ],
                    "target_types": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "domain": "string",
              "group_by_email": "string",
              "iam_member": "string",
              "role": "string",
              "routine": [
                "list",
                [
                  "object",
                  {
                    "dataset_id": "string",
                    "project_id": "string",
                    "routine_id": "string"
                  }
                ]
              ],
              "special_group": "string",
              "user_by_email": "string",
              "view": [
                "list",
                [
                  "object",
                  {
                    "dataset_id": "string",
                    "project_id": "string",
                    "table_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "creation_time": {
        "computed": true,
        "description": "The time when this dataset was created, in milliseconds since the\nepoch.",
        "description_kind": "plain",
        "type": "number"
      },
      "dataset_id": {
        "description": "A unique ID for this dataset, without the project name. The ID\nmust contain only letters (a-z, A-Z), numbers (0-9), or\nunderscores (_). The maximum length is 1,024 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_collation": {
        "computed": true,
        "description": "Defines the default collation specification of future tables created\nin the dataset. If a table is created in this dataset without table-level\ndefault collation, then the table inherits the dataset default collation,\nwhich is applied to the string fields that do not have explicit collation\nspecified. A change to this field affects only tables created afterwards,\nand does not alter the existing tables.\n\nThe following values are supported:\n- 'und:ci': undetermined locale, case insensitive.\n- '': empty string. Default to case-sensitive behavior.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_encryption_configuration": {
        "computed": true,
        "description": "The default encryption key for all tables in the dataset. Once this property is set,\nall newly-created partitioned tables in the dataset will have encryption key set to\nthis value, unless table creation request (or query) overrides the key.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key_name": "string"
            }
          ]
        ]
      },
      "default_partition_expiration_ms": {
        "computed": true,
        "description": "The default partition expiration for all partitioned tables in\nthe dataset, in milliseconds.\nOnce this property is set, all newly-created partitioned tables in\nthe dataset will have an 'expirationMs' property in the 'timePartitioning'\nsettings set to this value, and changing the value will only\naffect new tables, not existing ones. The storage in a partition will\nhave an expiration time of its partition time plus this value.\nSetting this property overrides the use of 'defaultTableExpirationMs'\nfor partitioned tables: only one of 'defaultTableExpirationMs' and\n'defaultPartitionExpirationMs' will be used for any new partitioned\ntable. If you provide an explicit 'timePartitioning.expirationMs' when\ncreating or updating a partitioned table, that value takes precedence\nover the default partition expiration time indicated by this property.",
        "description_kind": "plain",
        "type": "number"
      },
      "default_table_expiration_ms": {
        "computed": true,
        "description": "The default lifetime of all tables in the dataset, in milliseconds.\nThe minimum value is 3600000 milliseconds (one hour).\nOnce this property is set, all newly-created tables in the dataset\nwill have an 'expirationTime' property set to the creation time plus\nthe value in this property, and changing the value will only affect\nnew tables, not existing ones. When the 'expirationTime' for a given\ntable is reached, that table will be deleted automatically.\nIf a table's 'expirationTime' is modified or removed before the\ntable expires, or if you provide an explicit 'expirationTime' when\ncreating a table, that value takes precedence over the default\nexpiration time indicated by this property.",
        "description_kind": "plain",
        "type": "number"
      },
      "delete_contents_on_destroy": {
        "computed": true,
        "description": "If set to 'true', delete all the tables in the\ndataset when destroying the resource; otherwise,\ndestroying the resource will fail if tables are present.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "A user-friendly description of the dataset",
        "description_kind": "plain",
        "type": "string"
      },
      "effective_labels": {
        "computed": true,
        "description": "All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "etag": {
        "computed": true,
        "description": "A hash of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "external_catalog_dataset_options": {
        "computed": true,
        "description": "Options defining open source compatible datasets living in the BigQuery catalog. Contains\nmetadata of open source database, schema or namespace represented by the current dataset.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "default_storage_location_uri": "string",
              "parameters": [
                "map",
                "string"
              ]
            }
          ]
        ]
      },
      "external_dataset_reference": {
        "computed": true,
        "description": "Information about the external metadata storage where the dataset is defined.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection": "string",
              "external_source": "string"
            }
          ]
        ]
      },
      "friendly_name": {
        "computed": true,
        "description": "A descriptive name for the dataset",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_case_insensitive": {
        "computed": true,
        "description": "TRUE if the dataset and its table names are case-insensitive, otherwise FALSE.\nBy default, this is FALSE, which means the dataset and its table names are\ncase-sensitive. This field does not affect routine references.",
        "description_kind": "plain",
        "type": "bool"
      },
      "labels": {
        "computed": true,
        "description": "The labels associated with this dataset. You can use these to\norganize and group your datasets.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "last_modified_time": {
        "computed": true,
        "description": "The date when this dataset or any of its tables was last modified, in\nmilliseconds since the epoch.",
        "description_kind": "plain",
        "type": "number"
      },
      "location": {
        "computed": true,
        "description": "The geographic location where the dataset should reside.\nSee [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).\nThere are two types of locations, regional or multi-regional. A regional\nlocation is a specific geographic place, such as Tokyo, and a multi-regional\nlocation is a large geographic area, such as the United States, that\ncontains at least two geographic places.\nThe default value is multi-regional location 'US'.\nChanging this forces a new resource to be created.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_time_travel_hours": {
        "computed": true,
        "description": "Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_tags": {
        "computed": true,
        "description": "The tags attached to this table. Tag keys are globally unique. Tag key is expected to be\nin the namespaced format, for example \"123456789012/environment\" where 123456789012 is the\nID of the parent organization or project resource for this tag key. Tag value is expected\nto be the short name, for example \"Production\". See [Tag definitions](https://cloud.google.com/iam/docs/tags-access-control#definitions)\nfor more details.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_billing_model": {
        "computed": true,
        "description": "Specifies the storage billing model for the dataset.\nSet this flag value to LOGICAL to use logical bytes for storage billing,\nor to PHYSICAL to use physical bytes instead.\n\nLOGICAL is the default if this flag isn't specified.",
        "description_kind": "plain",
        "type": "string"
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBigqueryDatasetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryDataset), &result)
	return &result
}
