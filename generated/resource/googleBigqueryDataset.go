package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryDataset = `{
  "block": {
    "attributes": {
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
        "optional": true,
        "type": "string"
      },
      "default_partition_expiration_ms": {
        "description": "The default partition expiration for all partitioned tables in\nthe dataset, in milliseconds.\n\n\nOnce this property is set, all newly-created partitioned tables in\nthe dataset will have an 'expirationMs' property in the 'timePartitioning'\nsettings set to this value, and changing the value will only\naffect new tables, not existing ones. The storage in a partition will\nhave an expiration time of its partition time plus this value.\nSetting this property overrides the use of 'defaultTableExpirationMs'\nfor partitioned tables: only one of 'defaultTableExpirationMs' and\n'defaultPartitionExpirationMs' will be used for any new partitioned\ntable. If you provide an explicit 'timePartitioning.expirationMs' when\ncreating or updating a partitioned table, that value takes precedence\nover the default partition expiration time indicated by this property.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "default_table_expiration_ms": {
        "description": "The default lifetime of all tables in the dataset, in milliseconds.\nThe minimum value is 3600000 milliseconds (one hour).\n\n\nOnce this property is set, all newly-created tables in the dataset\nwill have an 'expirationTime' property set to the creation time plus\nthe value in this property, and changing the value will only affect\nnew tables, not existing ones. When the 'expirationTime' for a given\ntable is reached, that table will be deleted automatically.\nIf a table's 'expirationTime' is modified or removed before the\ntable expires, or if you provide an explicit 'expirationTime' when\ncreating a table, that value takes precedence over the default\nexpiration time indicated by this property.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "delete_contents_on_destroy": {
        "description": "If set to 'true', delete all the tables in the\ndataset when destroying the resource; otherwise,\ndestroying the resource will fail if tables are present.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "A user-friendly description of the dataset",
        "description_kind": "plain",
        "optional": true,
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
      "friendly_name": {
        "description": "A descriptive name for the dataset",
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
      "is_case_insensitive": {
        "computed": true,
        "description": "TRUE if the dataset and its table names are case-insensitive, otherwise FALSE.\nBy default, this is FALSE, which means the dataset and its table names are\ncase-sensitive. This field does not affect routine references.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "labels": {
        "description": "The labels associated with this dataset. You can use these to\norganize and group your datasets.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The geographic location where the dataset should reside.\nSee [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).\n\n\nThere are two types of locations, regional or multi-regional. A regional\nlocation is a specific geographic place, such as Tokyo, and a multi-regional\nlocation is a large geographic area, such as the United States, that\ncontains at least two geographic places.\n\n\nThe default value is multi-regional location 'US'.\nChanging this forces a new resource to be created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_time_travel_hours": {
        "computed": true,
        "description": "Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
        "optional": true,
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
    "block_types": {
      "access": {
        "block": {
          "attributes": {
            "domain": {
              "description": "A domain to grant access to. Any users signed in with the\ndomain specified will be granted the specified access",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "group_by_email": {
              "description": "An email address of a Google Group to grant access to.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "iam_member": {
              "description": "Some other type of member that appears in the IAM Policy but isn't a user,\ngroup, domain, or special group. For example: 'allUsers'",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role": {
              "description": "Describes the rights granted to the user specified by the other\nmember of the access object. Basic, predefined, and custom roles\nare supported. Predefined roles that have equivalent basic roles\nare swapped by the API to their basic counterparts. See\n[official docs](https://cloud.google.com/bigquery/docs/access-control).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "special_group": {
              "description": "A special group to grant access to. Possible values include:\n\n\n* 'projectOwners': Owners of the enclosing project.\n\n\n* 'projectReaders': Readers of the enclosing project.\n\n\n* 'projectWriters': Writers of the enclosing project.\n\n\n* 'allAuthenticatedUsers': All authenticated BigQuery users.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_by_email": {
              "description": "An email address of a user to grant access to. For example:\nfred@example.com",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "dataset": {
              "block": {
                "attributes": {
                  "target_types": {
                    "description": "Which resources in the dataset this entry applies to. Currently, only views are supported,\nbut additional target types may be added in the future. Possible values: VIEWS",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "dataset": {
                    "block": {
                      "attributes": {
                        "dataset_id": {
                          "description": "The ID of the dataset containing this table.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "project_id": {
                          "description": "The ID of the project containing this table.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The dataset this entry applies to",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Grants all resources of particular types in a particular dataset read access to the current dataset.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "routine": {
              "block": {
                "attributes": {
                  "dataset_id": {
                    "description": "The ID of the dataset containing this table.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "project_id": {
                    "description": "The ID of the project containing this table.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "routine_id": {
                    "description": "The ID of the routine. The ID must contain only letters (a-z,\nA-Z), numbers (0-9), or underscores (_). The maximum length\nis 256 characters.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "A routine from a different dataset to grant access to. Queries\nexecuted against that routine will have read access to tables in\nthis dataset. The role field is not required when this field is\nset. If that routine is updated by any user, access to the routine\nneeds to be granted again via an update operation.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "view": {
              "block": {
                "attributes": {
                  "dataset_id": {
                    "description": "The ID of the dataset containing this table.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "project_id": {
                    "description": "The ID of the project containing this table.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "table_id": {
                    "description": "The ID of the table. The ID must contain only letters (a-z,\nA-Z), numbers (0-9), or underscores (_). The maximum length\nis 1,024 characters.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "A view from a different dataset to grant access to. Queries\nexecuted against that view will have read access to tables in\nthis dataset. The role field is not required when this field is\nset. If that view is updated by any user, access to the view\nneeds to be granted again via an update operation.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "An array of objects that define dataset access for one or more entities.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "default_encryption_configuration": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "Describes the Cloud KMS encryption key that will be used to protect destination\nBigQuery table. The BigQuery Service Account associated with your project requires\naccess to this encryption key.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The default encryption key for all tables in the dataset. Once this property is set,\nall newly-created partitioned tables in the dataset will have encryption key set to\nthis value, unless table creation request (or query) overrides the key.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "external_dataset_reference": {
        "block": {
          "attributes": {
            "connection": {
              "description": "The connection id that is used to access the externalSource.\nFormat: projects/{projectId}/locations/{locationId}/connections/{connectionId}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "external_source": {
              "description": "External source that backs this dataset.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Information about the external metadata storage where the dataset is defined.",
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

func GoogleBigqueryDatasetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryDataset), &result)
	return &result
}
