package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryTable = `{
  "block": {
    "attributes": {
      "biglake_configuration": {
        "computed": true,
        "description": "Specifies the configuration of a BigLake managed table.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_id": "string",
              "file_format": "string",
              "storage_uri": "string",
              "table_format": "string"
            }
          ]
        ]
      },
      "clustering": {
        "computed": true,
        "description": "Specifies column names to use for data clustering. Up to four top-level columns are allowed, and should be specified in descending priority order.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "creation_time": {
        "computed": true,
        "description": "The time when this table was created, in milliseconds since the epoch.",
        "description_kind": "plain",
        "type": "number"
      },
      "dataset_id": {
        "description": "The dataset ID to create the table in. Changing this forces a new resource to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Whether Terraform will be prevented from destroying the instance. When the field is set to true or unset in Terraform state, a terraform apply or terraform destroy that would delete the table will fail. When the field is set to false, deleting the table is allowed.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "The field description.",
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
      "encryption_configuration": {
        "computed": true,
        "description": "Specifies how the table should be encrypted. If left blank, the table will be encrypted with a Google-managed key; that process is transparent to the user.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key_name": "string",
              "kms_key_version": "string"
            }
          ]
        ]
      },
      "etag": {
        "computed": true,
        "description": "A hash of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "expiration_time": {
        "computed": true,
        "description": "The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed.",
        "description_kind": "plain",
        "type": "number"
      },
      "external_catalog_table_options": {
        "computed": true,
        "description": "Options defining open source compatible table.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_id": "string",
              "parameters": [
                "map",
                "string"
              ],
              "storage_descriptor": [
                "list",
                [
                  "object",
                  {
                    "input_format": "string",
                    "location_uri": "string",
                    "output_format": "string",
                    "serde_info": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "parameters": [
                            "map",
                            "string"
                          ],
                          "serialization_library": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "external_data_configuration": {
        "computed": true,
        "description": "Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "autodetect": "bool",
              "avro_options": [
                "list",
                [
                  "object",
                  {
                    "use_avro_logical_types": "bool"
                  }
                ]
              ],
              "bigtable_options": [
                "list",
                [
                  "object",
                  {
                    "column_family": [
                      "list",
                      [
                        "object",
                        {
                          "column": [
                            "list",
                            [
                              "object",
                              {
                                "encoding": "string",
                                "field_name": "string",
                                "only_read_latest": "bool",
                                "qualifier_encoded": "string",
                                "qualifier_string": "string",
                                "type": "string"
                              }
                            ]
                          ],
                          "encoding": "string",
                          "family_id": "string",
                          "only_read_latest": "bool",
                          "type": "string"
                        }
                      ]
                    ],
                    "ignore_unspecified_column_families": "bool",
                    "output_column_families_as_json": "bool",
                    "read_rowkey_as_string": "bool"
                  }
                ]
              ],
              "compression": "string",
              "connection_id": "string",
              "csv_options": [
                "list",
                [
                  "object",
                  {
                    "allow_jagged_rows": "bool",
                    "allow_quoted_newlines": "bool",
                    "encoding": "string",
                    "field_delimiter": "string",
                    "quote": "string",
                    "skip_leading_rows": "number"
                  }
                ]
              ],
              "file_set_spec_type": "string",
              "google_sheets_options": [
                "list",
                [
                  "object",
                  {
                    "range": "string",
                    "skip_leading_rows": "number"
                  }
                ]
              ],
              "hive_partitioning_options": [
                "list",
                [
                  "object",
                  {
                    "mode": "string",
                    "require_partition_filter": "bool",
                    "source_uri_prefix": "string"
                  }
                ]
              ],
              "ignore_unknown_values": "bool",
              "json_extension": "string",
              "json_options": [
                "list",
                [
                  "object",
                  {
                    "encoding": "string"
                  }
                ]
              ],
              "max_bad_records": "number",
              "metadata_cache_mode": "string",
              "object_metadata": "string",
              "parquet_options": [
                "list",
                [
                  "object",
                  {
                    "enable_list_inference": "bool",
                    "enum_as_string": "bool"
                  }
                ]
              ],
              "reference_file_schema_uri": "string",
              "schema": "string",
              "source_format": "string",
              "source_uris": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "friendly_name": {
        "computed": true,
        "description": "A descriptive name for the table.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ignore_schema_changes": {
        "computed": true,
        "description": "Mention which fields in schema are to be ignored",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "labels": {
        "computed": true,
        "description": "A mapping of labels to assign to the resource.\n\n\t\t\t\t**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\n\t\t\t\tPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "last_modified_time": {
        "computed": true,
        "description": "The time when this table was last modified, in milliseconds since the epoch.",
        "description_kind": "plain",
        "type": "number"
      },
      "location": {
        "computed": true,
        "description": "The geographic location where the table resides. This value is inherited from the dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "materialized_view": {
        "computed": true,
        "description": "If specified, configures this table as a materialized view.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allow_non_incremental_definition": "bool",
              "enable_refresh": "bool",
              "query": "string",
              "refresh_interval_ms": "number"
            }
          ]
        ]
      },
      "max_staleness": {
        "computed": true,
        "description": "The maximum staleness of data that could be returned when the table (or stale MV) is queried. Staleness encoded as a string encoding of [SQL IntervalValue type](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#interval_type).",
        "description_kind": "plain",
        "type": "string"
      },
      "num_bytes": {
        "computed": true,
        "description": "The geographic location where the table resides. This value is inherited from the dataset.",
        "description_kind": "plain",
        "type": "number"
      },
      "num_long_term_bytes": {
        "computed": true,
        "description": "The number of bytes in the table that are considered \"long-term storage\".",
        "description_kind": "plain",
        "type": "number"
      },
      "num_rows": {
        "computed": true,
        "description": "The number of rows of data in this table, excluding any data in the streaming buffer.",
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "description": "The ID of the project in which the resource belongs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "range_partitioning": {
        "computed": true,
        "description": "If specified, configures range-based partitioning for this table.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "field": "string",
              "range": [
                "list",
                [
                  "object",
                  {
                    "end": "number",
                    "interval": "number",
                    "start": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "require_partition_filter": {
        "computed": true,
        "description": "If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.",
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_tags": {
        "computed": true,
        "description": "The tags attached to this table. Tag keys are globally unique. Tag key is expected to be in the namespaced format, for example \"123456789012/environment\" where 123456789012 is the ID of the parent organization or project resource for this tag key. Tag value is expected to be the short name, for example \"Production\".",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "schema": {
        "computed": true,
        "description": "A JSON schema for the table.",
        "description_kind": "plain",
        "type": "string"
      },
      "schema_foreign_type_info": {
        "computed": true,
        "description": "Specifies metadata of the foreign data type definition in field schema.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "type_system": "string"
            }
          ]
        ]
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_constraints": {
        "computed": true,
        "description": "Defines the primary key and foreign keys.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "foreign_keys": [
                "list",
                [
                  "object",
                  {
                    "column_references": [
                      "list",
                      [
                        "object",
                        {
                          "referenced_column": "string",
                          "referencing_column": "string"
                        }
                      ]
                    ],
                    "name": "string",
                    "referenced_table": [
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
              ],
              "primary_key": [
                "list",
                [
                  "object",
                  {
                    "columns": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "table_id": {
        "description": "A unique ID for the resource. Changing this forces a new resource to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "table_metadata_view": {
        "computed": true,
        "description": "View sets the optional parameter \"view\": Specifies the view that determines which table information is returned. By default, basic table information and storage statistics (STORAGE_STATS) are returned. Possible values: TABLE_METADATA_VIEW_UNSPECIFIED, BASIC, STORAGE_STATS, FULL",
        "description_kind": "plain",
        "type": "string"
      },
      "table_replication_info": {
        "computed": true,
        "description": "Replication info of a table created using \"AS REPLICA\" DDL like: \"CREATE MATERIALIZED VIEW mv1 AS REPLICA OF src_mv\".",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "replication_interval_ms": "number",
              "source_dataset_id": "string",
              "source_project_id": "string",
              "source_table_id": "string"
            }
          ]
        ]
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "time_partitioning": {
        "computed": true,
        "description": "If specified, configures time-based partitioning for this table.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "expiration_ms": "number",
              "field": "string",
              "require_partition_filter": "bool",
              "type": "string"
            }
          ]
        ]
      },
      "type": {
        "computed": true,
        "description": "Describes the table type.",
        "description_kind": "plain",
        "type": "string"
      },
      "view": {
        "computed": true,
        "description": "If specified, configures this table as a view.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "query": "string",
              "use_legacy_sql": "bool"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBigqueryTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryTable), &result)
	return &result
}
