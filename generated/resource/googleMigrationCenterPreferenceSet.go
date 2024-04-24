package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMigrationCenterPreferenceSet = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The timestamp when the preference set was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A description of the preference set.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "User-friendly display name. Maximum length is 63 characters.",
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
      "location": {
        "description": "Part of 'parent'. See documentation of 'projectsId'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Output only. Name of the preference set.",
        "description_kind": "plain",
        "type": "string"
      },
      "preference_set_id": {
        "description": "Required. User specified ID for the preference set. It will become the last component of the preference set name. The ID must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. The ID must match the regular expression '[a-z]([a-z0-9-]{0,61}[a-z0-9])?'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The timestamp when the preference set was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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
      },
      "virtual_machine_preferences": {
        "block": {
          "attributes": {
            "commitment_plan": {
              "description": "Commitment plan to consider when calculating costs for virtual machine insights and recommendations. If you are unsure which value to set, a 3 year commitment plan is often a good value to start with. \n Possible values:\n COMMITMENT_PLAN_UNSPECIFIED\nCOMMITMENT_PLAN_NONE\nCOMMITMENT_PLAN_ONE_YEAR\nCOMMITMENT_PLAN_THREE_YEARS",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sizing_optimization_strategy": {
              "description": "Sizing optimization strategy specifies the preferred strategy used when extrapolating usage data to calculate insights and recommendations for a virtual machine. If you are unsure which value to set, a moderate sizing optimization strategy is often a good value to start with. \n Possible values:\n SIZING_OPTIMIZATION_STRATEGY_UNSPECIFIED\nSIZING_OPTIMIZATION_STRATEGY_SAME_AS_SOURCE\nSIZING_OPTIMIZATION_STRATEGY_MODERATE\nSIZING_OPTIMIZATION_STRATEGY_AGGRESSIVE",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_product": {
              "description": "Target product for assets using this preference set. Specify either target product or business goal, but not both. \n Possible values:\n COMPUTE_MIGRATION_TARGET_PRODUCT_UNSPECIFIED\nCOMPUTE_MIGRATION_TARGET_PRODUCT_COMPUTE_ENGINE\nCOMPUTE_MIGRATION_TARGET_PRODUCT_VMWARE_ENGINE\nCOMPUTE_MIGRATION_TARGET_PRODUCT_SOLE_TENANCY",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "compute_engine_preferences": {
              "block": {
                "attributes": {
                  "license_type": {
                    "description": "License type to consider when calculating costs for virtual machine insights and recommendations. If unspecified, costs are calculated based on the default licensing plan. \n Possible values:\n LICENSE_TYPE_UNSPECIFIED\nLICENSE_TYPE_DEFAULT\nLICENSE_TYPE_BRING_YOUR_OWN_LICENSE",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "machine_preferences": {
                    "block": {
                      "block_types": {
                        "allowed_machine_series": {
                          "block": {
                            "attributes": {
                              "code": {
                                "description": "Code to identify a Compute Engine machine series. Consult https://cloud.google.com/compute/docs/machine-resource#machine_type_comparison for more details on the available series.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Compute Engine machine series to consider for insights and recommendations. If empty, no restriction is applied on the machine series.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The type of machines to consider when calculating virtual machine migration insights and recommendations. Not all machine types are available in all zones and regions.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The user preferences relating to Compute Engine target platform.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "region_preferences": {
              "block": {
                "attributes": {
                  "preferred_regions": {
                    "description": "A list of preferred regions, ordered by the most preferred region first. Set only valid Google Cloud region names. See https://cloud.google.com/compute/docs/regions-zones for available regions.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "The user preferences relating to target regions.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sole_tenancy_preferences": {
              "block": {
                "attributes": {
                  "commitment_plan": {
                    "description": "Commitment plan to consider when calculating costs for virtual machine insights and recommendations. If you are unsure which value to set, a 3 year commitment plan is often a good value to start with. \n Possible values:\n COMMITMENT_PLAN_UNSPECIFIED\nON_DEMAND\nCOMMITMENT_1_YEAR\nCOMMITMENT_3_YEAR",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cpu_overcommit_ratio": {
                    "description": "CPU overcommit ratio. Acceptable values are between 1.0 and 2.0 inclusive.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "host_maintenance_policy": {
                    "description": "Sole Tenancy nodes maintenance policy. \n Possible values:\n HOST_MAINTENANCE_POLICY_UNSPECIFIED\nHOST_MAINTENANCE_POLICY_DEFAULT\nHOST_MAINTENANCE_POLICY_RESTART_IN_PLACE\nHOST_MAINTENANCE_POLICY_MIGRATE_WITHIN_NODE_GROUP",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "node_types": {
                    "block": {
                      "attributes": {
                        "node_name": {
                          "description": "Name of the Sole Tenant node. Consult https://cloud.google.com/compute/docs/nodes/sole-tenant-nodes",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "A list of sole tenant node types. An empty list means that all possible node types will be considered.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Preferences concerning Sole Tenancy nodes and VMs.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "vmware_engine_preferences": {
              "block": {
                "attributes": {
                  "commitment_plan": {
                    "description": "Commitment plan to consider when calculating costs for virtual machine insights and recommendations. If you are unsure which value to set, a 3 year commitment plan is often a good value to start with. \n Possible values:\n COMMITMENT_PLAN_UNSPECIFIED\nON_DEMAND\nCOMMITMENT_1_YEAR_MONTHLY_PAYMENTS\nCOMMITMENT_3_YEAR_MONTHLY_PAYMENTS\nCOMMITMENT_1_YEAR_UPFRONT_PAYMENT\nCOMMITMENT_3_YEAR_UPFRONT_PAYMENT",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cpu_overcommit_ratio": {
                    "description": "CPU overcommit ratio. Acceptable values are between 1.0 and 8.0, with 0.1 increment.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "memory_overcommit_ratio": {
                    "description": "Memory overcommit ratio. Acceptable values are 1.0, 1.25, 1.5, 1.75 and 2.0.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "storage_deduplication_compression_ratio": {
                    "description": "The Deduplication and Compression ratio is based on the logical (Used Before) space required to store data before applying deduplication and compression, in relation to the physical (Used After) space required after applying deduplication and compression. Specifically, the ratio is the Used Before space divided by the Used After space. For example, if the Used Before space is 3 GB, but the physical Used After space is 1 GB, the deduplication and compression ratio is 3x. Acceptable values are between 1.0 and 4.0.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "The user preferences relating to Google Cloud VMware Engine target platform.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "VirtualMachinePreferences enables you to create sets of assumptions, for example, a geographical location and pricing track, for your migrated virtual machines. The set of preferences influence recommendations for migrating virtual machine assets.",
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

func GoogleMigrationCenterPreferenceSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMigrationCenterPreferenceSet), &result)
	return &result
}
