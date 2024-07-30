package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataLossPreventionDiscoveryConfig = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The creation timestamp of a DiscoveryConfig.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "Display Name (max 1000 Chars)",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "errors": {
        "computed": true,
        "description": "Output only. A stream of errors encountered when the config was activated. Repeated errors may result in the config automatically being paused. Output only field. Will return the last 100 errors. Whenever the config is modified this list will be cleared.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "details": [
                "list",
                [
                  "object",
                  {
                    "code": "number",
                    "details": [
                      "list",
                      [
                        "map",
                        "string"
                      ]
                    ],
                    "message": "string"
                  }
                ]
              ],
              "timestamp": "string"
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
      "inspect_templates": {
        "description": "Detection logic for profile generation",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "last_run_time": {
        "computed": true,
        "description": "Output only. The timestamp of the last time this config was executed",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "Location to create the discovery config in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Unique resource name for the DiscoveryConfig, assigned by the service when the DiscoveryConfig is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The parent of the discovery config in any of the following formats:\n\n* 'projects/{{project}}/locations/{{location}}'\n* 'organizations/{{organization_id}}/locations/{{location}}'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "description": "Required. A status for this configuration Possible values: [\"RUNNING\", \"PAUSED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The last update timestamp of a DiscoveryConfig.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "actions": {
        "block": {
          "block_types": {
            "export_data": {
              "block": {
                "block_types": {
                  "profile_table": {
                    "block": {
                      "attributes": {
                        "dataset_id": {
                          "description": "Dataset Id of the table",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "project_id": {
                          "description": "The Google Cloud Platform project ID of the project containing the table. If omitted, the project ID is inferred from the API call.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "table_id": {
                          "description": "Name of the table",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Store all table and column profiles in an existing table or a new table in an existing dataset. Each re-generation will result in a new row in BigQuery",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Export data profiles into a provided location",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "pub_sub_notification": {
              "block": {
                "attributes": {
                  "detail_of_message": {
                    "description": "How much data to include in the pub/sub message. Possible values: [\"TABLE_PROFILE\", \"RESOURCE_NAME\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "event": {
                    "description": "The type of event that triggers a Pub/Sub. At most one PubSubNotification per EventType is permitted. Possible values: [\"NEW_PROFILE\", \"CHANGED_PROFILE\", \"SCORE_INCREASED\", \"ERROR_CHANGED\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "topic": {
                    "description": "Cloud Pub/Sub topic to send notifications to. Format is projects/{project}/topics/{topic}.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "pubsub_condition": {
                    "block": {
                      "block_types": {
                        "expressions": {
                          "block": {
                            "attributes": {
                              "logical_operator": {
                                "description": "The operator to apply to the collection of conditions Possible values: [\"OR\", \"AND\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "conditions": {
                                "block": {
                                  "attributes": {
                                    "minimum_risk_score": {
                                      "description": "The minimum data risk score that triggers the condition. Possible values: [\"HIGH\", \"MEDIUM_OR_HIGH\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "minimum_sensitivity_score": {
                                      "description": "The minimum sensitivity level that triggers the condition. Possible values: [\"HIGH\", \"MEDIUM_OR_HIGH\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Conditions to apply to the expression",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "An expression",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Conditions for triggering pubsub",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Publish a message into the Pub/Sub topic.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Actions to execute at the completion of scanning",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "org_config": {
        "block": {
          "attributes": {
            "project_id": {
              "description": "The project that will run the scan. The DLP service account that exists within this project must have access to all resources that are profiled, and the cloud DLP API must be enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "location": {
              "block": {
                "attributes": {
                  "folder_id": {
                    "description": "The ID for the folder within an organization to scan",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "organization_id": {
                    "description": "The ID of an organization to scan",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The data to scan folder org or project",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A nested object resource",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "targets": {
        "block": {
          "block_types": {
            "big_query_target": {
              "block": {
                "block_types": {
                  "cadence": {
                    "block": {
                      "block_types": {
                        "schema_modified_cadence": {
                          "block": {
                            "attributes": {
                              "frequency": {
                                "description": "How frequently profiles may be updated when schemas are modified. Default to monthly Possible values: [\"UPDATE_FREQUENCY_NEVER\", \"UPDATE_FREQUENCY_DAILY\", \"UPDATE_FREQUENCY_MONTHLY\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "types": {
                                "description": "The type of events to consider when deciding if the table's schema has been modified and should have the profile updated. Defaults to NEW_COLUMN. Possible values: [\"SCHEMA_NEW_COLUMNS\", \"SCHEMA_REMOVED_COLUMNS\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "Governs when to update data profiles when a schema is modified",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "table_modified_cadence": {
                          "block": {
                            "attributes": {
                              "frequency": {
                                "description": "How frequently data profiles can be updated when tables are modified. Defaults to never. Possible values: [\"UPDATE_FREQUENCY_NEVER\", \"UPDATE_FREQUENCY_DAILY\", \"UPDATE_FREQUENCY_MONTHLY\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "types": {
                                "description": "The type of events to consider when deciding if the table has been modified and should have the profile updated. Defaults to MODIFIED_TIMESTAMP Possible values: [\"TABLE_MODIFIED_TIMESTAMP\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "Governs when to update profile when a table is modified.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "How often and when to update profiles. New tables that match both the fiter and conditions are scanned as quickly as possible depending on system capacity.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "conditions": {
                    "block": {
                      "attributes": {
                        "created_after": {
                          "description": "A timestamp in RFC3339 UTC \"Zulu\" format with nanosecond resolution and upto nine fractional digits.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type_collection": {
                          "description": "Restrict discovery to categories of table types. Currently view, materialized view, snapshot and non-biglake external tables are supported. Possible values: [\"BIG_QUERY_COLLECTION_ALL_TYPES\", \"BIG_QUERY_COLLECTION_ONLY_SUPPORTED_TYPES\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "or_conditions": {
                          "block": {
                            "attributes": {
                              "min_age": {
                                "description": "Duration format. The minimum age a table must have before Cloud DLP can profile it. Value greater than 1.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "min_row_count": {
                                "description": "Minimum number of rows that should be present before Cloud DLP profiles as a table.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "At least one of the conditions must be true for a table to be scanned.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "types": {
                          "block": {
                            "attributes": {
                              "types": {
                                "description": "A set of BiqQuery table types Possible values: [\"BIG_QUERY_TABLE_TYPE_TABLE\", \"BIG_QUERY_TABLE_TYPE_EXTERNAL_BIG_LAKE\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "Restrict discovery to specific table type",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "In addition to matching the filter, these conditions must be true before a profile is generated",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "disabled": {
                    "block": {
                      "description": "Tables that match this filter will not have profiles created.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "filter": {
                    "block": {
                      "block_types": {
                        "other_tables": {
                          "block": {
                            "description": "Catch-all. This should always be the last filter in the list because anything above it will apply first.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "table_reference": {
                          "block": {
                            "attributes": {
                              "dataset_id": {
                                "description": "Dataset ID of the table.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "table_id": {
                                "description": "Name of the table.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "The table to scan. Discovery configurations including this can only include one DiscoveryTarget (the DiscoveryTarget with this TableReference).",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "tables": {
                          "block": {
                            "block_types": {
                              "include_regexes": {
                                "block": {
                                  "block_types": {
                                    "patterns": {
                                      "block": {
                                        "attributes": {
                                          "dataset_id_regex": {
                                            "description": "if unset, this property matches all datasets",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "project_id_regex": {
                                            "description": "For organizations, if unset, will match all projects. Has no effect for data profile configurations created within a project.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "table_id_regex": {
                                            "description": "if unset, this property matches all tables",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "A single BigQuery regular expression pattern to match against one or more tables, datasets, or projects that contain BigQuery tables.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "A collection of regular expressions to match a BQ table against.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A specific set of tables for this filter to apply to. A table collection must be specified in only one filter per config.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Required. The tables the discovery cadence applies to. The first target with a matching filter will be the one to apply to a table",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "BigQuery target for Discovery. The first target to match a table will be the one applied.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "cloud_sql_target": {
              "block": {
                "block_types": {
                  "conditions": {
                    "block": {
                      "attributes": {
                        "database_engines": {
                          "description": "Database engines that should be profiled. Optional. Defaults to ALL_SUPPORTED_DATABASE_ENGINES if unspecified. Possible values: [\"ALL_SUPPORTED_DATABASE_ENGINES\", \"MYSQL\", \"POSTGRES\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "types": {
                          "description": "Data profiles will only be generated for the database resource types specified in this field. If not specified, defaults to [DATABASE_RESOURCE_TYPE_ALL_SUPPORTED_TYPES]. Possible values: [\"DATABASE_RESOURCE_TYPE_ALL_SUPPORTED_TYPES\", \"DATABASE_RESOURCE_TYPE_TABLE\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "In addition to matching the filter, these conditions must be true before a profile is generated.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "disabled": {
                    "block": {
                      "description": "Disable profiling for database resources that match this filter.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "filter": {
                    "block": {
                      "block_types": {
                        "collection": {
                          "block": {
                            "block_types": {
                              "include_regexes": {
                                "block": {
                                  "block_types": {
                                    "patterns": {
                                      "block": {
                                        "attributes": {
                                          "database_regex": {
                                            "description": "Regex to test the database name against. If empty, all databases match.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "database_resource_name_regex": {
                                            "description": "Regex to test the database resource's name against. An example of a database resource name is a table's name. Other database resource names like view names could be included in the future. If empty, all database resources match.'",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "instance_regex": {
                                            "description": "Regex to test the instance name against. If empty, all instances match.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "project_id_regex": {
                                            "description": "For organizations, if unset, will match all projects. Has no effect for data profile configurations created within a project.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "A group of regular expression patterns to match against one or more database resources. Maximum of 100 entries. The sum of all regular expressions' length can't exceed 10 KiB.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "A collection of regular expressions to match a database resource against.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A specific set of database resources for this filter to apply to.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "database_resource_reference": {
                          "block": {
                            "attributes": {
                              "database": {
                                "description": "Required. Name of a database within the instance.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "database_resource": {
                                "description": "Required. Name of a database resource, for example, a table within the database.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "instance": {
                                "description": "Required. The instance where this resource is located. For example: Cloud SQL instance ID.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "project_id": {
                                "description": "Required. If within a project-level config, then this must match the config's project ID.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "The database resource to scan. Targets including this can only include one target (the target with this database resource reference).",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "others": {
                          "block": {
                            "description": "Catch-all. This should always be the last target in the list because anything above it will apply first. Should only appear once in a configuration. If none is specified, a default one will be added automatically.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Required. The tables the discovery cadence applies to. The first target with a matching filter will be the one to apply to a table.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "generation_cadence": {
                    "block": {
                      "attributes": {
                        "refresh_frequency": {
                          "description": "Data changes (non-schema changes) in Cloud SQL tables can't trigger reprofiling. If you set this field, profiles are refreshed at this frequency regardless of whether the underlying tables have changes. Defaults to never. Possible values: [\"UPDATE_FREQUENCY_NEVER\", \"UPDATE_FREQUENCY_DAILY\", \"UPDATE_FREQUENCY_MONTHLY\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "schema_modified_cadence": {
                          "block": {
                            "attributes": {
                              "frequency": {
                                "description": "Frequency to regenerate data profiles when the schema is modified. Defaults to monthly. Possible values: [\"UPDATE_FREQUENCY_NEVER\", \"UPDATE_FREQUENCY_DAILY\", \"UPDATE_FREQUENCY_MONTHLY\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "types": {
                                "description": "The types of schema modifications to consider. Defaults to NEW_COLUMNS. Possible values: [\"NEW_COLUMNS\", \"REMOVED_COLUMNS\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "Governs when to update data profiles when a schema is modified",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "How often and when to update profiles. New tables that match both the filter and conditions are scanned as quickly as possible depending on system capacity.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Cloud SQL target for Discovery. The first target to match a table will be the one applied.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "cloud_storage_target": {
              "block": {
                "block_types": {
                  "conditions": {
                    "block": {
                      "attributes": {
                        "created_after": {
                          "description": "File store must have been created after this date. Used to avoid backfilling. A timestamp in RFC3339 UTC \"Zulu\" format with nanosecond resolution and upto nine fractional digits.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "min_age": {
                          "description": "Duration format. Minimum age a file store must have. If set, the value must be 1 hour or greater.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "cloud_storage_conditions": {
                          "block": {
                            "attributes": {
                              "included_bucket_attributes": {
                                "description": "Only objects with the specified attributes will be scanned. Defaults to [ALL_SUPPORTED_BUCKETS] if unset. Possible values: [\"ALL_SUPPORTED_BUCKETS\", \"AUTOCLASS_DISABLED\", \"AUTOCLASS_ENABLED\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "included_object_attributes": {
                                "description": "Only objects with the specified attributes will be scanned. If an object has one of the specified attributes but is inside an excluded bucket, it will not be scanned. Defaults to [ALL_SUPPORTED_OBJECTS]. A profile will be created even if no objects match the included_object_attributes. Possible values: [\"ALL_SUPPORTED_OBJECTS\", \"STANDARD\", \"NEARLINE\", \"COLDLINE\", \"ARCHIVE\", \"REGIONAL\", \"MULTI_REGIONAL\", \"DURABLE_REDUCED_AVAILABILITY\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "Cloud Storage conditions.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "In addition to matching the filter, these conditions must be true before a profile is generated.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "disabled": {
                    "block": {
                      "description": "Disable profiling for buckets that match this filter.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "filter": {
                    "block": {
                      "block_types": {
                        "cloud_storage_resource_reference": {
                          "block": {
                            "attributes": {
                              "bucket_name": {
                                "description": "The bucket to scan.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "project_id": {
                                "description": "If within a project-level config, then this must match the config's project id.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The bucket to scan. Targets including this can only include one target (the target with this bucket). This enables profiling the contents of a single bucket, while the other options allow for easy profiling of many buckets within a project or an organization.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "collection": {
                          "block": {
                            "block_types": {
                              "include_regexes": {
                                "block": {
                                  "block_types": {
                                    "patterns": {
                                      "block": {
                                        "block_types": {
                                          "cloud_storage_regex": {
                                            "block": {
                                              "attributes": {
                                                "bucket_name_regex": {
                                                  "description": "Regex to test the bucket name against. If empty, all buckets match. Example: \"marketing2021\" or \"(marketing)\\d{4}\" will both match the bucket gs://marketing2021",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "project_id_regex": {
                                                  "description": "For organizations, if unset, will match all projects.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Regex for Cloud Storage.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "The group of regular expression patterns to match against one or more file stores. Maximum of 100 entries. The sum of all lengths of regular expressions can't exceed 10 KiB.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "A collection of regular expressions to match a file store against.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A specific set of buckets for this filter to apply to.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "others": {
                          "block": {
                            "description": "Match discovery resources not covered by any other filter.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The buckets the generation_cadence applies to. The first target with a matching filter will be the one to apply to a bucket.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "generation_cadence": {
                    "block": {
                      "attributes": {
                        "refresh_frequency": {
                          "description": "Data changes in Cloud Storage can't trigger reprofiling. If you set this field, profiles are refreshed at this frequency regardless of whether the underlying buckets have changes. Defaults to never. Possible values: [\"UPDATE_FREQUENCY_NEVER\", \"UPDATE_FREQUENCY_DAILY\", \"UPDATE_FREQUENCY_MONTHLY\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "inspect_template_modified_cadence": {
                          "block": {
                            "attributes": {
                              "frequency": {
                                "description": "How frequently data profiles can be updated when the template is modified. Defaults to never. Possible values: [\"UPDATE_FREQUENCY_NEVER\", \"UPDATE_FREQUENCY_DAILY\", \"UPDATE_FREQUENCY_MONTHLY\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Governs when to update data profiles when the inspection rules defined by the 'InspectTemplate' change. If not set, changing the template will not cause a data profile to update.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "How often and when to update profiles. New buckets that match both the filter and conditions are scanned as quickly as possible depending on system capacity.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Cloud Storage target for Discovery. The first target to match a bucket will be the one applied.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "secrets_target": {
              "block": {
                "description": "Discovery target that looks for credentials and secrets stored in cloud resource metadata and reports them as vulnerabilities to Security Command Center. Only one target of this type is allowed.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Target to match against for determining what to scan and how frequently",
          "description_kind": "plain"
        },
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

func GoogleDataLossPreventionDiscoveryConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataLossPreventionDiscoveryConfig), &result)
	return &result
}
