package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDatastreamStream = `{
  "block": {
    "attributes": {
      "create_without_validation": {
        "description": "Create the stream without validating it.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "customer_managed_encryption_key": {
        "description": "A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data\nwill be encrypted using an internal Stream-specific encryption key provisioned through KMS.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "desired_state": {
        "description": "Desired state of the Stream. Set this field to 'RUNNING' to start the stream,\n'NOT_STARTED' to create the stream without starting and 'PAUSED' to pause\nthe stream from a 'RUNNING' state.\nPossible values: NOT_STARTED, RUNNING, PAUSED. Default: NOT_STARTED",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Display name.",
        "description_kind": "plain",
        "required": true,
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The name of the location this stream is located in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The stream's name.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "stream_id": {
        "description": "The stream identifier.",
        "description_kind": "plain",
        "required": true,
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
      "backfill_all": {
        "block": {
          "block_types": {
            "mysql_excluded_objects": {
              "block": {
                "block_types": {
                  "mysql_databases": {
                    "block": {
                      "attributes": {
                        "database": {
                          "description": "Database name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "mysql_tables": {
                          "block": {
                            "attributes": {
                              "table": {
                                "description": "Table name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "mysql_columns": {
                                "block": {
                                  "attributes": {
                                    "collation": {
                                      "description": "Column collation.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "column": {
                                      "description": "Column name.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "data_type": {
                                      "description": "The MySQL data type. Full data types list can be found here:\nhttps://dev.mysql.com/doc/refman/8.0/en/data-types.html",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "length": {
                                      "computed": true,
                                      "description": "Column length.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "nullable": {
                                      "description": "Whether or not the column can accept a null value.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "ordinal_position": {
                                      "description": "The ordinal position of the column in the table.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "primary_key": {
                                      "description": "Whether or not the column represents a primary key.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description": "MySQL columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Tables in the database.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "MySQL databases on the server",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "MySQL data source objects to avoid backfilling.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oracle_excluded_objects": {
              "block": {
                "block_types": {
                  "oracle_schemas": {
                    "block": {
                      "attributes": {
                        "schema": {
                          "description": "Schema name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "oracle_tables": {
                          "block": {
                            "attributes": {
                              "table": {
                                "description": "Table name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "oracle_columns": {
                                "block": {
                                  "attributes": {
                                    "column": {
                                      "description": "Column name.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "data_type": {
                                      "description": "The Oracle data type. Full data types list can be found here:\nhttps://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/Data-Types.html",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "encoding": {
                                      "computed": true,
                                      "description": "Column encoding.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "length": {
                                      "computed": true,
                                      "description": "Column length.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "nullable": {
                                      "computed": true,
                                      "description": "Whether or not the column can accept a null value.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "ordinal_position": {
                                      "computed": true,
                                      "description": "The ordinal position of the column in the table.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "precision": {
                                      "computed": true,
                                      "description": "Column precision.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "primary_key": {
                                      "computed": true,
                                      "description": "Whether or not the column represents a primary key.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "scale": {
                                      "computed": true,
                                      "description": "Column scale.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "description": "Oracle columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Tables in the database.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Oracle schemas/databases in the database server",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "PostgreSQL data source objects to avoid backfilling.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "postgresql_excluded_objects": {
              "block": {
                "block_types": {
                  "postgresql_schemas": {
                    "block": {
                      "attributes": {
                        "schema": {
                          "description": "Database name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "postgresql_tables": {
                          "block": {
                            "attributes": {
                              "table": {
                                "description": "Table name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "postgresql_columns": {
                                "block": {
                                  "attributes": {
                                    "column": {
                                      "description": "Column name.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "data_type": {
                                      "description": "The PostgreSQL data type. Full data types list can be found here:\nhttps://www.postgresql.org/docs/current/datatype.html",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "length": {
                                      "computed": true,
                                      "description": "Column length.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "nullable": {
                                      "description": "Whether or not the column can accept a null value.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "ordinal_position": {
                                      "description": "The ordinal position of the column in the table.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "precision": {
                                      "computed": true,
                                      "description": "Column precision.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "primary_key": {
                                      "description": "Whether or not the column represents a primary key.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "scale": {
                                      "computed": true,
                                      "description": "Column scale.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "description": "PostgreSQL columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Tables in the schema.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "PostgreSQL schemas on the server",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "PostgreSQL data source objects to avoid backfilling.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "salesforce_excluded_objects": {
              "block": {
                "block_types": {
                  "objects": {
                    "block": {
                      "attributes": {
                        "object_name": {
                          "description": "Name of object in Salesforce Org.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "fields": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Field name.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Fields in the Salesforce object. When unspecified as part of include/exclude objects, includes/excludes everything/nothing.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Salesforce objects in Salesforce Org.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Salesforce objects to avoid backfilling.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sql_server_excluded_objects": {
              "block": {
                "block_types": {
                  "schemas": {
                    "block": {
                      "attributes": {
                        "schema": {
                          "description": "Schema name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "tables": {
                          "block": {
                            "attributes": {
                              "table": {
                                "description": "Table name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "columns": {
                                "block": {
                                  "attributes": {
                                    "column": {
                                      "description": "Column name.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "data_type": {
                                      "description": "The SQL Server data type. Full data types list can be found here:\nhttps://learn.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql?view=sql-server-ver16",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "length": {
                                      "computed": true,
                                      "description": "Column length.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "nullable": {
                                      "computed": true,
                                      "description": "Whether or not the column can accept a null value.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "ordinal_position": {
                                      "computed": true,
                                      "description": "The ordinal position of the column in the table.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "precision": {
                                      "computed": true,
                                      "description": "Column precision.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "primary_key": {
                                      "computed": true,
                                      "description": "Whether or not the column represents a primary key.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "scale": {
                                      "computed": true,
                                      "description": "Column scale.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "description": "SQL Server columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Tables in the database.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "SQL Server schemas/databases in the database server",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "SQL Server data source objects to avoid backfilling.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "backfill_none": {
        "block": {
          "description": "Backfill strategy to disable automatic backfill for the Stream's objects.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "destination_config": {
        "block": {
          "attributes": {
            "destination_connection_profile": {
              "description": "Destination connection profile resource. Format: projects/{project}/locations/{location}/connectionProfiles/{name}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "bigquery_destination_config": {
              "block": {
                "attributes": {
                  "data_freshness": {
                    "description": "The guaranteed data freshness (in seconds) when querying tables created by the stream.\nEditing this field will only affect new tables created in the future, but existing tables\nwill not be impacted. Lower values mean that queries will return fresher data, but may result in higher cost.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\". Defaults to 900s.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "append_only": {
                    "block": {
                      "description": "AppendOnly mode defines that the stream of changes (INSERT, UPDATE-INSERT, UPDATE-DELETE and DELETE\nevents) to a source table will be written to the destination Google BigQuery table, retaining the\nhistorical state of the data.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "blmt_config": {
                    "block": {
                      "attributes": {
                        "bucket": {
                          "description": "The Cloud Storage bucket name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "connection_name": {
                          "description": "The bigquery connection. Format: '{project}.{location}.{name}'",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "file_format": {
                          "description": "The file format.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "root_path": {
                          "description": "The root path inside the Cloud Storage bucket.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "table_format": {
                          "description": "The table format.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "BigLake Managed Tables configuration for BigQuery streams.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "merge": {
                    "block": {
                      "description": "Merge mode defines that all changes to a table will be merged at the destination Google BigQuery\ntable. This is the default write mode. When selected, BigQuery reflects the way the data is stored\nin the source database. With Merge mode, no historical record of the change events is kept.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "single_target_dataset": {
                    "block": {
                      "attributes": {
                        "dataset_id": {
                          "description": "Dataset ID in the format projects/{project}/datasets/{dataset_id} or\n{project}:{dataset_id}",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "A single target dataset to which all data will be streamed.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "source_hierarchy_datasets": {
                    "block": {
                      "block_types": {
                        "dataset_template": {
                          "block": {
                            "attributes": {
                              "dataset_id_prefix": {
                                "description": "If supplied, every created dataset will have its name prefixed by the provided value.\nThe prefix and name will be separated by an underscore. i.e. _.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "kms_key_name": {
                                "description": "Describes the Cloud KMS encryption key that will be used to protect destination BigQuery\ntable. The BigQuery Service Account associated with your project requires access to this\nencryption key. i.e. projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{cryptoKey}.\nSee https://cloud.google.com/bigquery/docs/customer-managed-encryption for more information.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "location": {
                                "description": "The geographic location where the dataset should reside.\nSee https://cloud.google.com/bigquery/docs/locations for supported locations.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Dataset template used for dynamic dataset creation.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Destination datasets are created so that hierarchy of the destination data objects matches the source hierarchy.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A configuration for how data should be loaded to Google BigQuery.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gcs_destination_config": {
              "block": {
                "attributes": {
                  "file_rotation_interval": {
                    "computed": true,
                    "description": "The maximum duration for which new events are added before a file is closed and a new file is created.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\". Defaults to 900s.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "file_rotation_mb": {
                    "computed": true,
                    "description": "The maximum file size to be saved in the bucket.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "path": {
                    "description": "Path inside the Cloud Storage bucket to write data to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "avro_file_format": {
                    "block": {
                      "description": "AVRO file format configuration.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "json_file_format": {
                    "block": {
                      "attributes": {
                        "compression": {
                          "description": "Compression of the loaded JSON file. Possible values: [\"NO_COMPRESSION\", \"GZIP\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "schema_file_format": {
                          "description": "The schema file format along JSON data files. Possible values: [\"NO_SCHEMA_FILE\", \"AVRO_SCHEMA_FILE\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "JSON file format configuration.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A configuration for how data should be loaded to Cloud Storage.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Destination connection profile configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "source_config": {
        "block": {
          "attributes": {
            "source_connection_profile": {
              "description": "Source connection profile resource. Format: projects/{project}/locations/{location}/connectionProfiles/{name}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "mysql_source_config": {
              "block": {
                "attributes": {
                  "max_concurrent_backfill_tasks": {
                    "computed": true,
                    "description": "Maximum number of concurrent backfill tasks. The number should be non negative.\nIf not set (or set to 0), the system's default value will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_concurrent_cdc_tasks": {
                    "computed": true,
                    "description": "Maximum number of concurrent CDC tasks. The number should be non negative.\nIf not set (or set to 0), the system's default value will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "binary_log_position": {
                    "block": {
                      "description": "CDC reader reads from binary logs replication cdc method.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "exclude_objects": {
                    "block": {
                      "block_types": {
                        "mysql_databases": {
                          "block": {
                            "attributes": {
                              "database": {
                                "description": "Database name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "mysql_tables": {
                                "block": {
                                  "attributes": {
                                    "table": {
                                      "description": "Table name.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "mysql_columns": {
                                      "block": {
                                        "attributes": {
                                          "collation": {
                                            "description": "Column collation.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "column": {
                                            "description": "Column name.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "data_type": {
                                            "description": "The MySQL data type. Full data types list can be found here:\nhttps://dev.mysql.com/doc/refman/8.0/en/data-types.html",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "length": {
                                            "computed": true,
                                            "description": "Column length.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "nullable": {
                                            "description": "Whether or not the column can accept a null value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "ordinal_position": {
                                            "description": "The ordinal position of the column in the table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "primary_key": {
                                            "description": "Whether or not the column represents a primary key.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          }
                                        },
                                        "description": "MySQL columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Tables in the database.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "MySQL databases on the server",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "MySQL objects to exclude from the stream.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "gtid": {
                    "block": {
                      "description": "CDC reader reads from gtid based replication.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "include_objects": {
                    "block": {
                      "block_types": {
                        "mysql_databases": {
                          "block": {
                            "attributes": {
                              "database": {
                                "description": "Database name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "mysql_tables": {
                                "block": {
                                  "attributes": {
                                    "table": {
                                      "description": "Table name.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "mysql_columns": {
                                      "block": {
                                        "attributes": {
                                          "collation": {
                                            "description": "Column collation.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "column": {
                                            "description": "Column name.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "data_type": {
                                            "description": "The MySQL data type. Full data types list can be found here:\nhttps://dev.mysql.com/doc/refman/8.0/en/data-types.html",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "length": {
                                            "computed": true,
                                            "description": "Column length.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "nullable": {
                                            "description": "Whether or not the column can accept a null value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "ordinal_position": {
                                            "description": "The ordinal position of the column in the table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "primary_key": {
                                            "description": "Whether or not the column represents a primary key.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          }
                                        },
                                        "description": "MySQL columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Tables in the database.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "MySQL databases on the server",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "MySQL objects to retrieve from the source.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "MySQL data source configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oracle_source_config": {
              "block": {
                "attributes": {
                  "max_concurrent_backfill_tasks": {
                    "computed": true,
                    "description": "Maximum number of concurrent backfill tasks. The number should be non negative.\nIf not set (or set to 0), the system's default value will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_concurrent_cdc_tasks": {
                    "computed": true,
                    "description": "Maximum number of concurrent CDC tasks. The number should be non negative.\nIf not set (or set to 0), the system's default value will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "drop_large_objects": {
                    "block": {
                      "description": "Configuration to drop large object values.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "exclude_objects": {
                    "block": {
                      "block_types": {
                        "oracle_schemas": {
                          "block": {
                            "attributes": {
                              "schema": {
                                "description": "Schema name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "oracle_tables": {
                                "block": {
                                  "attributes": {
                                    "table": {
                                      "description": "Table name.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "oracle_columns": {
                                      "block": {
                                        "attributes": {
                                          "column": {
                                            "description": "Column name.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "data_type": {
                                            "description": "The Oracle data type. Full data types list can be found here:\nhttps://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/Data-Types.html",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "encoding": {
                                            "computed": true,
                                            "description": "Column encoding.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "length": {
                                            "computed": true,
                                            "description": "Column length.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "nullable": {
                                            "computed": true,
                                            "description": "Whether or not the column can accept a null value.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "ordinal_position": {
                                            "computed": true,
                                            "description": "The ordinal position of the column in the table.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "precision": {
                                            "computed": true,
                                            "description": "Column precision.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "primary_key": {
                                            "computed": true,
                                            "description": "Whether or not the column represents a primary key.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "scale": {
                                            "computed": true,
                                            "description": "Column scale.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          }
                                        },
                                        "description": "Oracle columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Tables in the database.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Oracle schemas/databases in the database server",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Oracle objects to exclude from the stream.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "include_objects": {
                    "block": {
                      "block_types": {
                        "oracle_schemas": {
                          "block": {
                            "attributes": {
                              "schema": {
                                "description": "Schema name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "oracle_tables": {
                                "block": {
                                  "attributes": {
                                    "table": {
                                      "description": "Table name.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "oracle_columns": {
                                      "block": {
                                        "attributes": {
                                          "column": {
                                            "description": "Column name.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "data_type": {
                                            "description": "The Oracle data type. Full data types list can be found here:\nhttps://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/Data-Types.html",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "encoding": {
                                            "computed": true,
                                            "description": "Column encoding.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "length": {
                                            "computed": true,
                                            "description": "Column length.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "nullable": {
                                            "computed": true,
                                            "description": "Whether or not the column can accept a null value.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "ordinal_position": {
                                            "computed": true,
                                            "description": "The ordinal position of the column in the table.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "precision": {
                                            "computed": true,
                                            "description": "Column precision.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "primary_key": {
                                            "computed": true,
                                            "description": "Whether or not the column represents a primary key.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "scale": {
                                            "computed": true,
                                            "description": "Column scale.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          }
                                        },
                                        "description": "Oracle columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Tables in the database.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Oracle schemas/databases in the database server",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Oracle objects to retrieve from the source.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "stream_large_objects": {
                    "block": {
                      "description": "Configuration to drop large object values.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "MySQL data source configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "postgresql_source_config": {
              "block": {
                "attributes": {
                  "max_concurrent_backfill_tasks": {
                    "computed": true,
                    "description": "Maximum number of concurrent backfill tasks. The number should be non\nnegative. If not set (or set to 0), the system's default value will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "publication": {
                    "description": "The name of the publication that includes the set of all tables\nthat are defined in the stream's include_objects.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "replication_slot": {
                    "description": "The name of the logical replication slot that's configured with\nthe pgoutput plugin.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "exclude_objects": {
                    "block": {
                      "block_types": {
                        "postgresql_schemas": {
                          "block": {
                            "attributes": {
                              "schema": {
                                "description": "Database name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "postgresql_tables": {
                                "block": {
                                  "attributes": {
                                    "table": {
                                      "description": "Table name.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "postgresql_columns": {
                                      "block": {
                                        "attributes": {
                                          "column": {
                                            "description": "Column name.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "data_type": {
                                            "description": "The PostgreSQL data type. Full data types list can be found here:\nhttps://www.postgresql.org/docs/current/datatype.html",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "length": {
                                            "computed": true,
                                            "description": "Column length.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "nullable": {
                                            "description": "Whether or not the column can accept a null value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "ordinal_position": {
                                            "description": "The ordinal position of the column in the table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "precision": {
                                            "computed": true,
                                            "description": "Column precision.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "primary_key": {
                                            "description": "Whether or not the column represents a primary key.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "scale": {
                                            "computed": true,
                                            "description": "Column scale.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          }
                                        },
                                        "description": "PostgreSQL columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Tables in the schema.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "PostgreSQL schemas on the server",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "PostgreSQL objects to exclude from the stream.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "include_objects": {
                    "block": {
                      "block_types": {
                        "postgresql_schemas": {
                          "block": {
                            "attributes": {
                              "schema": {
                                "description": "Database name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "postgresql_tables": {
                                "block": {
                                  "attributes": {
                                    "table": {
                                      "description": "Table name.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "postgresql_columns": {
                                      "block": {
                                        "attributes": {
                                          "column": {
                                            "description": "Column name.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "data_type": {
                                            "description": "The PostgreSQL data type. Full data types list can be found here:\nhttps://www.postgresql.org/docs/current/datatype.html",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "length": {
                                            "computed": true,
                                            "description": "Column length.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "nullable": {
                                            "description": "Whether or not the column can accept a null value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "ordinal_position": {
                                            "description": "The ordinal position of the column in the table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "precision": {
                                            "computed": true,
                                            "description": "Column precision.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "primary_key": {
                                            "description": "Whether or not the column represents a primary key.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "scale": {
                                            "computed": true,
                                            "description": "Column scale.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          }
                                        },
                                        "description": "PostgreSQL columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Tables in the schema.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "PostgreSQL schemas on the server",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "PostgreSQL objects to retrieve from the source.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "PostgreSQL data source configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "salesforce_source_config": {
              "block": {
                "attributes": {
                  "polling_interval": {
                    "description": "Salesforce objects polling interval. The interval at which new changes will be polled for each object. The duration must be between 5 minutes and 24 hours.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "exclude_objects": {
                    "block": {
                      "block_types": {
                        "objects": {
                          "block": {
                            "attributes": {
                              "object_name": {
                                "description": "Name of object in Salesforce Org.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "fields": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "Field name.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Fields in the Salesforce object. When unspecified as part of include/exclude objects, includes/excludes everything/nothing.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Salesforce objects in data source.",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Salesforce objects to exclude from the stream.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "include_objects": {
                    "block": {
                      "block_types": {
                        "objects": {
                          "block": {
                            "attributes": {
                              "object_name": {
                                "description": "Name of object in Salesforce Org.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "fields": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "Field name.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Fields in the Salesforce object. When unspecified as part of include/exclude objects, includes/excludes everything/nothing.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Salesforce objects in Salesforce Org.",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Salesforce objects to retrieve from the source.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Salesforce data source configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sql_server_source_config": {
              "block": {
                "attributes": {
                  "max_concurrent_backfill_tasks": {
                    "computed": true,
                    "description": "Max concurrent backfill tasks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_concurrent_cdc_tasks": {
                    "computed": true,
                    "description": "Max concurrent CDC tasks.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "change_tables": {
                    "block": {
                      "description": "CDC reader reads from change tables.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "exclude_objects": {
                    "block": {
                      "block_types": {
                        "schemas": {
                          "block": {
                            "attributes": {
                              "schema": {
                                "description": "Schema name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "tables": {
                                "block": {
                                  "attributes": {
                                    "table": {
                                      "description": "Table name.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "columns": {
                                      "block": {
                                        "attributes": {
                                          "column": {
                                            "description": "Column name.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "data_type": {
                                            "description": "The SQL Server data type. Full data types list can be found here:\nhttps://learn.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql?view=sql-server-ver16",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "length": {
                                            "computed": true,
                                            "description": "Column length.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "nullable": {
                                            "computed": true,
                                            "description": "Whether or not the column can accept a null value.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "ordinal_position": {
                                            "computed": true,
                                            "description": "The ordinal position of the column in the table.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "precision": {
                                            "computed": true,
                                            "description": "Column precision.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "primary_key": {
                                            "computed": true,
                                            "description": "Whether or not the column represents a primary key.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "scale": {
                                            "computed": true,
                                            "description": "Column scale.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          }
                                        },
                                        "description": "SQL Server columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Tables in the database.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "SQL Server schemas/databases in the database server",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "SQL Server objects to exclude from the stream.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "include_objects": {
                    "block": {
                      "block_types": {
                        "schemas": {
                          "block": {
                            "attributes": {
                              "schema": {
                                "description": "Schema name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "tables": {
                                "block": {
                                  "attributes": {
                                    "table": {
                                      "description": "Table name.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "columns": {
                                      "block": {
                                        "attributes": {
                                          "column": {
                                            "description": "Column name.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "data_type": {
                                            "description": "The SQL Server data type. Full data types list can be found here:\nhttps://learn.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql?view=sql-server-ver16",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "length": {
                                            "computed": true,
                                            "description": "Column length.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "nullable": {
                                            "computed": true,
                                            "description": "Whether or not the column can accept a null value.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "ordinal_position": {
                                            "computed": true,
                                            "description": "The ordinal position of the column in the table.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "precision": {
                                            "computed": true,
                                            "description": "Column precision.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "primary_key": {
                                            "computed": true,
                                            "description": "Whether or not the column represents a primary key.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "scale": {
                                            "computed": true,
                                            "description": "Column scale.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          }
                                        },
                                        "description": "SQL Server columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Tables in the database.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "SQL Server schemas/databases in the database server",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "SQL Server objects to retrieve from the source.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "transaction_logs": {
                    "block": {
                      "description": "CDC reader reads from transaction logs.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "SQL Server data source configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Source connection profile configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleDatastreamStreamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDatastreamStream), &result)
	return &result
}
