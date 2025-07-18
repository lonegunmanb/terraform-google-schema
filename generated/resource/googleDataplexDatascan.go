package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataplexDatascan = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time when the scan was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_scan_id": {
        "description": "DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "Description of the scan.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "User friendly display name.",
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
      "execution_status": {
        "computed": true,
        "description": "Status of the data scan execution.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "latest_job_end_time": "string",
              "latest_job_start_time": "string"
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
      "labels": {
        "description": "User-defined labels for the scan. A list of key-\u003evalue pairs.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location where the data scan should reside.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The relative resource name of the scan, of the form: projects/{project}/locations/{locationId}/dataScans/{datascan_id}, where project refers to a project_id or project_number and locationId refers to a GCP region.",
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
        "description": "Current state of the DataScan.",
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
      },
      "type": {
        "computed": true,
        "description": "The type of DataScan.",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "System generated globally unique ID for the scan. This ID will be different if the scan is deleted and re-created with the same name.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time when the scan was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "data": {
        "block": {
          "attributes": {
            "entity": {
              "description": "The Dataplex entity that represents the data source(e.g. BigQuery table) for Datascan.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource": {
              "description": "The service-qualified full resource name of the cloud resource for a DataScan job to scan against. The field could be:\nCloud Storage bucket (//storage.googleapis.com/projects/PROJECT_ID/buckets/BUCKET_ID) for DataDiscoveryScan OR BigQuery table of type \"TABLE\" (/bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID) for DataProfileScan/DataQualityScan.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The data source for DataScan.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "data_discovery_spec": {
        "block": {
          "block_types": {
            "bigquery_publishing_config": {
              "block": {
                "attributes": {
                  "connection": {
                    "description": "The BigQuery connection used to create BigLake tables. Must be in the form 'projects/{projectId}/locations/{locationId}/connections/{connection_id}'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "location": {
                    "description": "The location of the BigQuery dataset to publish BigLake external or non-BigLake external tables to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "project": {
                    "description": "The project of the BigQuery dataset to publish BigLake external or non-BigLake external tables to. If not specified, the project of the Cloud Storage bucket will be used. The format is \"projects/{project_id_or_number}\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_type": {
                    "description": "Determines whether to publish discovered tables as BigLake external tables or non-BigLake external tables. Possible values: [\"TABLE_TYPE_UNSPECIFIED\", \"EXTERNAL\", \"BIGLAKE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Configuration for metadata publishing.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "storage_config": {
              "block": {
                "attributes": {
                  "exclude_patterns": {
                    "description": "Defines the data to exclude during discovery. Provide a list of patterns that identify the data to exclude. For Cloud Storage bucket assets, these patterns are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these patterns are interpreted as patterns to match table names.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "include_patterns": {
                    "description": "Defines the data to include during discovery when only a subset of the data should be considered. Provide a list of patterns that identify the data to include. For Cloud Storage bucket assets, these patterns are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these patterns are interpreted as patterns to match table names.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "csv_options": {
                    "block": {
                      "attributes": {
                        "delimiter": {
                          "description": "The delimiter that is used to separate values. The default is ',' (comma).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "encoding": {
                          "description": "The character encoding of the data. The default is UTF-8.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "header_rows": {
                          "description": "The number of rows to interpret as header rows that should be skipped when reading data rows.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "quote": {
                          "description": "The character used to quote column values. Accepts '\"' (double quotation mark) or ''' (single quotation mark). If unspecified, defaults to '\"' (double quotation mark).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type_inference_disabled": {
                          "description": "Whether to disable the inference of data types for CSV data. If true, all columns are registered as strings.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Configuration for CSV data.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "json_options": {
                    "block": {
                      "attributes": {
                        "encoding": {
                          "description": "The character encoding of the data. The default is UTF-8.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type_inference_disabled": {
                          "description": "Whether to disable the inference of data types for JSON data. If true, all columns are registered as their primitive types (strings, number, or boolean).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Configuration for JSON data.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Configurations related to Cloud Storage as the data source.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "DataDiscoveryScan related setting.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "data_profile_spec": {
        "block": {
          "attributes": {
            "row_filter": {
              "description": "A filter applied to all rows in a single DataScan job. The filter needs to be a valid SQL expression for a WHERE clause in BigQuery standard SQL syntax. Example: col1 \u003e= 0 AND col2 \u003c 10",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sampling_percent": {
              "description": "The percentage of the records to be selected from the dataset for DataScan.\nValue can range between 0.0 and 100.0 with up to 3 significant decimal digits.\nSampling is not applied if 'sampling_percent' is not specified, 0 or 100.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "exclude_fields": {
              "block": {
                "attributes": {
                  "field_names": {
                    "description": "Expected input is a list of fully qualified names of fields as in the schema.\nOnly top-level field names for nested fields are supported.\nFor instance, if 'x' is of nested field type, listing 'x' is supported but 'x.y.z' is not supported. Here 'y' and 'y.z' are nested fields of 'x'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "The fields to exclude from data profile.\nIf specified, the fields will be excluded from data profile, regardless of 'include_fields' value.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "include_fields": {
              "block": {
                "attributes": {
                  "field_names": {
                    "description": "Expected input is a list of fully qualified names of fields as in the schema.\nOnly top-level field names for nested fields are supported.\nFor instance, if 'x' is of nested field type, listing 'x' is supported but 'x.y.z' is not supported. Here 'y' and 'y.z' are nested fields of 'x'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "The fields to include in data profile.\nIf not specified, all fields at the time of profile scan job execution are included, except for ones listed in 'exclude_fields'.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "post_scan_actions": {
              "block": {
                "block_types": {
                  "bigquery_export": {
                    "block": {
                      "attributes": {
                        "results_table": {
                          "description": "The BigQuery table to export DataProfileScan results to.\nFormat://bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "If set, results will be exported to the provided BigQuery table.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Actions to take upon job completion.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "DataProfileScan related setting.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "data_quality_spec": {
        "block": {
          "attributes": {
            "catalog_publishing_enabled": {
              "description": "If set, the latest DataScan job result will be published to Dataplex Catalog.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "row_filter": {
              "description": "A filter applied to all rows in a single DataScan job. The filter needs to be a valid SQL expression for a WHERE clause in BigQuery standard SQL syntax. Example: col1 \u003e= 0 AND col2 \u003c 10",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sampling_percent": {
              "description": "The percentage of the records to be selected from the dataset for DataScan.\nValue can range between 0.0 and 100.0 with up to 3 significant decimal digits.\nSampling is not applied if 'sampling_percent' is not specified, 0 or 100.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "post_scan_actions": {
              "block": {
                "block_types": {
                  "bigquery_export": {
                    "block": {
                      "attributes": {
                        "results_table": {
                          "description": "The BigQuery table to export DataQualityScan results to.\nFormat://bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "If set, results will be exported to the provided BigQuery table.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "notification_report": {
                    "block": {
                      "block_types": {
                        "job_end_trigger": {
                          "block": {
                            "description": "This trigger is triggered whenever a scan job run ends, regardless of the result.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "job_failure_trigger": {
                          "block": {
                            "description": "This trigger is triggered when the scan job itself fails, regardless of the result.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "recipients": {
                          "block": {
                            "attributes": {
                              "emails": {
                                "description": "The email recipients who will receive the DataQualityScan results report.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "The individuals or groups who are designated to receive notifications upon triggers.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "score_threshold_trigger": {
                          "block": {
                            "attributes": {
                              "score_threshold": {
                                "description": "The score range is in [0,100].",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "This trigger is triggered when the DQ score in the job result is less than a specified input score.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The configuration of notification report post scan action.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Actions to take upon job completion.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "rules": {
              "block": {
                "attributes": {
                  "column": {
                    "description": "The unnested column which this rule is evaluated against.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "description": {
                    "description": "Description of the rule.\nThe maximum length is 1,024 characters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "dimension": {
                    "description": "The dimension name a rule belongs to. Custom dimension name is supported with all uppercase letters and maximum length of 30 characters.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "ignore_null": {
                    "description": "Rows with null values will automatically fail a rule, unless ignoreNull is true. In that case, such null rows are trivially considered passing. Only applicable to ColumnMap rules.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "name": {
                    "description": "A mutable name for the rule.\nThe name must contain only letters (a-z, A-Z), numbers (0-9), or hyphens (-).\nThe maximum length is 63 characters.\nMust start with a letter.\nMust end with a number or a letter.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "suspended": {
                    "description": "Whether the Rule is active or suspended. Default = false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "threshold": {
                    "description": "The minimum ratio of passing_rows / total_rows required to pass this rule, with a range of [0.0, 1.0]. 0 indicates default value (i.e. 1.0).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "non_null_expectation": {
                    "block": {
                      "description": "ColumnMap rule which evaluates whether each column value is null.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "range_expectation": {
                    "block": {
                      "attributes": {
                        "max_value": {
                          "description": "The maximum column value allowed for a row to pass this validation. At least one of minValue and maxValue need to be provided.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "min_value": {
                          "description": "The minimum column value allowed for a row to pass this validation. At least one of minValue and maxValue need to be provided.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "strict_max_enabled": {
                          "description": "Whether each value needs to be strictly lesser than ('\u003c') the maximum, or if equality is allowed.\nOnly relevant if a maxValue has been defined. Default = false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "strict_min_enabled": {
                          "description": "Whether each value needs to be strictly greater than ('\u003e') the minimum, or if equality is allowed.\nOnly relevant if a minValue has been defined. Default = false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "ColumnMap rule which evaluates whether each column value lies between a specified range.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "regex_expectation": {
                    "block": {
                      "attributes": {
                        "regex": {
                          "description": "A regular expression the column value is expected to match.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "ColumnMap rule which evaluates whether each column value matches a specified regex.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "row_condition_expectation": {
                    "block": {
                      "attributes": {
                        "sql_expression": {
                          "description": "The SQL expression.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Table rule which evaluates whether each row passes the specified condition.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "set_expectation": {
                    "block": {
                      "attributes": {
                        "values": {
                          "description": "Expected values for the column value.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "ColumnMap rule which evaluates whether each column value is contained by a specified set.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "sql_assertion": {
                    "block": {
                      "attributes": {
                        "sql_statement": {
                          "description": "The SQL statement.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Table rule which evaluates whether any row matches invalid state.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "statistic_range_expectation": {
                    "block": {
                      "attributes": {
                        "max_value": {
                          "description": "The maximum column statistic value allowed for a row to pass this validation.\nAt least one of minValue and maxValue need to be provided.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "min_value": {
                          "description": "The minimum column statistic value allowed for a row to pass this validation.\nAt least one of minValue and maxValue need to be provided.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "statistic": {
                          "description": "column statistics. Possible values: [\"STATISTIC_UNDEFINED\", \"MEAN\", \"MIN\", \"MAX\"]",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "strict_max_enabled": {
                          "description": "Whether column statistic needs to be strictly lesser than ('\u003c') the maximum, or if equality is allowed.\nOnly relevant if a maxValue has been defined. Default = false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "strict_min_enabled": {
                          "description": "Whether column statistic needs to be strictly greater than ('\u003e') the minimum, or if equality is allowed.\nOnly relevant if a minValue has been defined. Default = false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "ColumnAggregate rule which evaluates whether the column aggregate statistic lies between a specified range.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "table_condition_expectation": {
                    "block": {
                      "attributes": {
                        "sql_expression": {
                          "description": "The SQL expression.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Table rule which evaluates whether the provided expression is true.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "uniqueness_expectation": {
                    "block": {
                      "description": "Row-level rule which evaluates whether each column value is unique.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The list of rules to evaluate against a data source. At least one rule is required.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "DataQualityScan related setting.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "execution_spec": {
        "block": {
          "attributes": {
            "field": {
              "description": "The unnested field (of type Date or Timestamp) that contains values which monotonically increase over time. If not specified, a data scan will run for all data in the table.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "trigger": {
              "block": {
                "block_types": {
                  "on_demand": {
                    "block": {
                      "description": "The scan runs once via dataScans.run API.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "schedule": {
                    "block": {
                      "attributes": {
                        "cron": {
                          "description": "Cron schedule for running scans periodically. This field is required for Schedule scans.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The scan is scheduled to run periodically.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Spec related to how often and when a scan should be triggered.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "DataScan execution settings.",
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

func GoogleDataplexDatascanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataplexDatascan), &result)
	return &result
}
