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
