package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryTable = `{
  "block": {
    "attributes": {
      "clustering": {
        "description": "Specifies column names to use for data clustering. Up to four top-level columns are allowed, and should be specified in descending priority order.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "Whether Terraform will be prevented from destroying the instance. When the field is set to true or unset in Terraform state, a terraform apply or terraform destroy that would delete the table will fail. When the field is set to false, deleting the table is allowed.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "The field description.",
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
      "expiration_time": {
        "computed": true,
        "description": "The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "friendly_name": {
        "description": "A descriptive name for the table.",
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
      "labels": {
        "description": "A mapping of labels to assign to the resource.\n\n\t\t\t\t**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\n\t\t\t\tPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
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
      "max_staleness": {
        "description": "The maximum staleness of data that could be returned when the table (or stale MV) is queried. Staleness encoded as a string encoding of [SQL IntervalValue type](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#interval_type).",
        "description_kind": "plain",
        "optional": true,
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
        "computed": true,
        "description": "The ID of the project in which the resource belongs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "require_partition_filter": {
        "description": "If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_tags": {
        "description": "The tags attached to this table. Tag keys are globally unique. Tag key is expected to be in the namespaced format, for example \"123456789012/environment\" where 123456789012 is the ID of the parent organization or project resource for this tag key. Tag value is expected to be the short name, for example \"Production\".",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "schema": {
        "computed": true,
        "description": "A JSON schema for the table.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_id": {
        "description": "A unique ID for the resource. Changing this forces a new resource to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "type": {
        "computed": true,
        "description": "Describes the table type.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "biglake_configuration": {
        "block": {
          "attributes": {
            "connection_id": {
              "description": "The connection specifying the credentials to be used to read and write to external storage, such as Cloud Storage. The connection_id can have the form \"\u0026lt;project\\_id\u0026gt;.\u0026lt;location\\_id\u0026gt;.\u0026lt;connection\\_id\u0026gt;\" or \"projects/\u0026lt;project\\_id\u0026gt;/locations/\u0026lt;location\\_id\u0026gt;/connections/\u0026lt;connection\\_id\u0026gt;\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "file_format": {
              "description": "The file format the data is stored in.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "storage_uri": {
              "description": "The fully qualified location prefix of the external folder where table data is stored. The '*' wildcard character is not allowed. The URI should be in the format \"gs://bucket/path_to_table/\"",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "table_format": {
              "description": "The table format the metadata only snapshots are stored in.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Specifies the configuration of a BigLake managed table.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "encryption_configuration": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "The self link or full name of a key which should be used to encrypt this table. Note that the default bigquery service account will need to have encrypt/decrypt permissions on this key - you may want to see the google_bigquery_default_service_account datasource and the google_kms_crypto_key_iam_binding resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "kms_key_version": {
              "computed": true,
              "description": "The self link or full name of the kms key version used to encrypt this table.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Specifies how the table should be encrypted. If left blank, the table will be encrypted with a Google-managed key; that process is transparent to the user.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "external_data_configuration": {
        "block": {
          "attributes": {
            "autodetect": {
              "description": "Let BigQuery try to autodetect the schema and format of the table.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "compression": {
              "description": "The compression type of the data source. Valid values are \"NONE\" or \"GZIP\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connection_id": {
              "description": "The connection specifying the credentials to be used to read external storage, such as Azure Blob, Cloud Storage, or S3. The connectionId can have the form \"\u003cproject\u003e.\u003clocation\u003e.\u003cconnection_id\u003e\" or \"projects/\u003cproject\u003e/locations/\u003clocation\u003e/connections/\u003cconnection_id\u003e\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "file_set_spec_type": {
              "description": "Specifies how source URIs are interpreted for constructing the file set to load.  By default source URIs are expanded against the underlying storage.  Other options include specifying manifest files. Only applicable to object storage systems.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ignore_unknown_values": {
              "description": "Indicates if BigQuery should allow extra values that are not represented in the table schema. If true, the extra values are ignored. If false, records with extra columns are treated as bad records, and if there are too many bad records, an invalid error is returned in the job result. The default value is false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "json_extension": {
              "description": "Load option to be used together with sourceFormat newline-delimited JSON to indicate that a variant of JSON is being loaded. To load newline-delimited GeoJSON, specify GEOJSON (and sourceFormat must be set to NEWLINE_DELIMITED_JSON).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_bad_records": {
              "description": "The maximum number of bad records that BigQuery can ignore when reading data.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "metadata_cache_mode": {
              "description": "Metadata Cache Mode for the table. Set this to enable caching of metadata from external data source.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "object_metadata": {
              "description": "Object Metadata is used to create Object Tables. Object Tables contain a listing of objects (with their metadata) found at the sourceUris. If ObjectMetadata is set, sourceFormat should be omitted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "reference_file_schema_uri": {
              "description": "When creating an external table, the user can provide a reference file with the table schema. This is enabled for the following formats: AVRO, PARQUET, ORC.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schema": {
              "computed": true,
              "description": "A JSON schema for the external table. Schema is required for CSV and JSON formats and is disallowed for Google Cloud Bigtable, Cloud Datastore backups, and Avro formats when using external tables.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_format": {
              "description": "Please see sourceFormat under ExternalDataConfiguration in Bigquery's public API documentation (https://cloud.google.com/bigquery/docs/reference/rest/v2/tables#externaldataconfiguration) for supported formats. To use \"GOOGLE_SHEETS\" the scopes must include \"googleapis.com/auth/drive.readonly\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_uris": {
              "description": "A list of the fully-qualified URIs that point to your data in Google Cloud.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "avro_options": {
              "block": {
                "attributes": {
                  "use_avro_logical_types": {
                    "description": "If sourceFormat is set to \"AVRO\", indicates whether to interpret logical types as the corresponding BigQuery data type (for example, TIMESTAMP), instead of using the raw type (for example, INTEGER).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Additional options if source_format is set to \"AVRO\"",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "bigtable_options": {
              "block": {
                "attributes": {
                  "ignore_unspecified_column_families": {
                    "description": "If field is true, then the column families that are not specified in columnFamilies list are not exposed in the table schema. Otherwise, they are read with BYTES type values. The default value is false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "output_column_families_as_json": {
                    "description": "If field is true, then each column family will be read as a single JSON column. Otherwise they are read as a repeated cell structure containing timestamp/value tuples. The default value is false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "read_rowkey_as_string": {
                    "description": "If field is true, then the rowkey column families will be read and converted to string. Otherwise they are read with BYTES type values and users need to manually cast them with CAST if necessary. The default value is false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "column_family": {
                    "block": {
                      "attributes": {
                        "encoding": {
                          "description": "The encoding of the values when the type is not STRING. Acceptable encoding values are: TEXT - indicates values are alphanumeric text strings. BINARY - indicates values are encoded using HBase Bytes.toBytes family of functions. This can be overridden for a specific column by listing that column in 'columns' and specifying an encoding for it.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "family_id": {
                          "description": "Identifier of the column family.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "only_read_latest": {
                          "description": "If this is set only the latest version of value are exposed for all columns in this column family. This can be overridden for a specific column by listing that column in 'columns' and specifying a different setting for that column.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "type": {
                          "description": "The type to convert the value in cells of this column family. The values are expected to be encoded using HBase Bytes.toBytes function when using the BINARY encoding value. Following BigQuery types are allowed (case-sensitive): \"BYTES\", \"STRING\", \"INTEGER\", \"FLOAT\", \"BOOLEAN\", \"JSON\". Default type is BYTES. This can be overridden for a specific column by listing that column in 'columns' and specifying a type for it.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "column": {
                          "block": {
                            "attributes": {
                              "encoding": {
                                "description": "The encoding of the values when the type is not STRING. Acceptable encoding values are: TEXT - indicates values are alphanumeric text strings. BINARY - indicates values are encoded using HBase Bytes.toBytes family of functions. 'encoding' can also be set at the column family level. However, the setting at this level takes precedence if 'encoding' is set at both levels.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "field_name": {
                                "description": "If the qualifier is not a valid BigQuery field identifier i.e. does not match [a-zA-Z][a-zA-Z0-9_]*, a valid identifier must be provided as the column field name and is used as field name in queries.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "only_read_latest": {
                                "description": "If this is set, only the latest version of value in this column are exposed. 'onlyReadLatest' can also be set at the column family level. However, the setting at this level takes precedence if 'onlyReadLatest' is set at both levels.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "qualifier_encoded": {
                                "description": "Qualifier of the column. Columns in the parent column family that has this exact qualifier are exposed as . field. If the qualifier is valid UTF-8 string, it can be specified in the qualifierString field. Otherwise, a base-64 encoded value must be set to qualifierEncoded. The column field name is the same as the column qualifier. However, if the qualifier is not a valid BigQuery field identifier i.e. does not match [a-zA-Z][a-zA-Z0-9_]*, a valid identifier must be provided as fieldName.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "qualifier_string": {
                                "description": "Qualifier string.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "The type to convert the value in cells of this column. The values are expected to be encoded using HBase Bytes.toBytes function when using the BINARY encoding value. Following BigQuery types are allowed (case-sensitive): \"BYTES\", \"STRING\", \"INTEGER\", \"FLOAT\", \"BOOLEAN\", \"JSON\", Default type is \"BYTES\". 'type' can also be set at the column family level. However, the setting at this level takes precedence if 'type' is set at both levels.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "A List of columns that should be exposed as individual fields as opposed to a list of (column name, value) pairs. All columns whose qualifier matches a qualifier in this list can be accessed as Other columns can be accessed as a list through column field",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "A list of column families to expose in the table schema along with their types. This list restricts the column families that can be referenced in queries and specifies their value types. You can use this list to do type conversions - see the 'type' field for more details. If you leave this list empty, all column families are present in the table schema and their values are read as BYTES. During a query only the column families referenced in that query are read from Bigtable.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Additional options if sourceFormat is set to BIGTABLE.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "csv_options": {
              "block": {
                "attributes": {
                  "allow_jagged_rows": {
                    "description": "Indicates if BigQuery should accept rows that are missing trailing optional columns.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "allow_quoted_newlines": {
                    "description": "Indicates if BigQuery should allow quoted data sections that contain newline characters in a CSV file. The default value is false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "encoding": {
                    "description": "The character encoding of the data. The supported values are UTF-8 or ISO-8859-1.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "field_delimiter": {
                    "description": "The separator for fields in a CSV file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "quote": {
                    "description": "The value that is used to quote data sections in a CSV file. If your data does not contain quoted sections, set the property value to an empty string. If your data contains quoted newline characters, you must also set the allow_quoted_newlines property to true. The API-side default is \", specified in Terraform escaped as \\\". Due to limitations with Terraform default values, this value is required to be explicitly set.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "skip_leading_rows": {
                    "description": "The number of rows at the top of a CSV file that BigQuery will skip when reading the data.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Additional properties to set if source_format is set to \"CSV\".",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "google_sheets_options": {
              "block": {
                "attributes": {
                  "range": {
                    "description": "Range of a sheet to query from. Only used when non-empty. At least one of range or skip_leading_rows must be set. Typical format: \"sheet_name!top_left_cell_id:bottom_right_cell_id\" For example: \"sheet1!A1:B20",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "skip_leading_rows": {
                    "description": "The number of rows at the top of the sheet that BigQuery will skip when reading the data. At least one of range or skip_leading_rows must be set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Additional options if source_format is set to \"GOOGLE_SHEETS\".",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "hive_partitioning_options": {
              "block": {
                "attributes": {
                  "mode": {
                    "description": "When set, what mode of hive partitioning to use when reading data.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "require_partition_filter": {
                    "description": "If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "source_uri_prefix": {
                    "description": "When hive partition detection is requested, a common for all source uris must be required. The prefix must end immediately before the partition key encoding begins.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "When set, configures hive partitioning support. Not all storage formats support hive partitioning -- requesting hive partitioning on an unsupported format will lead to an error, as will providing an invalid specification.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "json_options": {
              "block": {
                "attributes": {
                  "encoding": {
                    "description": "The character encoding of the data. The supported values are UTF-8, UTF-16BE, UTF-16LE, UTF-32BE, and UTF-32LE. The default value is UTF-8.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Additional properties to set if sourceFormat is set to JSON.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "parquet_options": {
              "block": {
                "attributes": {
                  "enable_list_inference": {
                    "description": "Indicates whether to use schema inference specifically for Parquet LIST logical type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enum_as_string": {
                    "description": "Indicates whether to infer Parquet ENUM logical type as STRING instead of BYTES by default.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Additional properties to set if sourceFormat is set to PARQUET.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "materialized_view": {
        "block": {
          "attributes": {
            "allow_non_incremental_definition": {
              "description": "Allow non incremental materialized view definition. The default value is false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_refresh": {
              "description": "Specifies if BigQuery should automatically refresh materialized view when the base table is updated. The default is true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "query": {
              "description": "A query whose result is persisted.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "refresh_interval_ms": {
              "description": "Specifies maximum frequency at which this materialized view will be refreshed. The default is 1800000.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "If specified, configures this table as a materialized view.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "range_partitioning": {
        "block": {
          "attributes": {
            "field": {
              "description": "The field used to determine how to create a range-based partition.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "range": {
              "block": {
                "attributes": {
                  "end": {
                    "description": "End of the range partitioning, exclusive.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "interval": {
                    "description": "The width of each range within the partition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "start": {
                    "description": "Start of the range partitioning, inclusive.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description": "Information required to partition based on ranges. Structure is documented below.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "If specified, configures range-based partitioning for this table.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "table_constraints": {
        "block": {
          "block_types": {
            "foreign_keys": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Set only if the foreign key constraint is named.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "column_references": {
                    "block": {
                      "attributes": {
                        "referenced_column": {
                          "description": "The column in the primary key that are referenced by the referencingColumn.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "referencing_column": {
                          "description": "The column that composes the foreign key.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The pair of the foreign key column and primary key column.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "referenced_table": {
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
                          "description": "The ID of the table. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 1,024 characters. Certain operations allow suffixing of the table ID with a partition decorator, such as sample_table$20190123.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The table that holds the primary key and is referenced by this foreign key.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Present only if the table has a foreign key. The foreign key is not enforced.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "primary_key": {
              "block": {
                "attributes": {
                  "columns": {
                    "description": "The columns that are composed of the primary key constraint.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Represents a primary key constraint on a table's columns. Present only if the table has a primary key. The primary key is not enforced.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Defines the primary key and foreign keys.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "table_replication_info": {
        "block": {
          "attributes": {
            "replication_interval_ms": {
              "description": "The interval at which the source materialized view is polled for updates. The default is 300000.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "source_dataset_id": {
              "description": "The ID of the source dataset.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_project_id": {
              "description": "The ID of the source project.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_table_id": {
              "description": "The ID of the source materialized view.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Replication info of a table created using \"AS REPLICA\" DDL like: \"CREATE MATERIALIZED VIEW mv1 AS REPLICA OF src_mv\".",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "time_partitioning": {
        "block": {
          "attributes": {
            "expiration_ms": {
              "computed": true,
              "description": "Number of milliseconds for which to keep the storage for a partition.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "field": {
              "description": "The field used to determine how to create a time-based partition. If time-based partitioning is enabled without this value, the table is partitioned based on the load time.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "require_partition_filter": {
              "deprecated": true,
              "description": "If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "type": {
              "description": "The supported types are DAY, HOUR, MONTH, and YEAR, which will generate one partition per day, hour, month, and year, respectively.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "If specified, configures time-based partitioning for this table.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "view": {
        "block": {
          "attributes": {
            "query": {
              "description": "A query that BigQuery executes when the view is referenced.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "use_legacy_sql": {
              "description": "Specifies whether to use BigQuery's legacy SQL for this view. The default value is true. If set to false, the view will use BigQuery's standard SQL",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "If specified, configures this table as a view.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
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
