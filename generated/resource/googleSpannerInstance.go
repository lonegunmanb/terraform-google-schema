package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSpannerInstance = `{
  "block": {
    "attributes": {
      "config": {
        "description": "The name of the instance's configuration (similar but not\nquite the same as a region) which defines the geographic placement and\nreplication of your databases in this instance. It determines where your data\nis stored. Values are typically of the form 'regional-europe-west1' , 'us-central' etc.\nIn order to obtain a valid list please consult the\n[Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_backup_schedule_type": {
        "computed": true,
        "description": "Controls the default backup behavior for new databases within the instance.\nNote that 'AUTOMATIC' is not permitted for free instances, as backups and backup schedules are not allowed for free instances.\nif unset or NONE, no default backup schedule will be created for new databases within the instance. Possible values: [\"NONE\", \"AUTOMATIC\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The descriptive name for this instance as it appears in UIs. Must be\nunique per project and between 4 and 30 characters in length.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "edition": {
        "computed": true,
        "description": "The edition selected for this instance. Different editions provide different capabilities at different price points. Possible values: [\"EDITION_UNSPECIFIED\", \"STANDARD\", \"ENTERPRISE\", \"ENTERPRISE_PLUS\"]",
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
      "force_destroy": {
        "description": "When deleting a spanner instance, this boolean option will delete all backups of this instance.\nThis must be set to true if you created a backup manually in the console.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description": "The type of this instance. The type can be used to distinguish product variants, that can affect aspects like:\nusage restrictions, quotas and billing. Currently this is used to distinguish FREE_INSTANCE vs PROVISIONED instances.\nWhen configured as FREE_INSTANCE, the field 'edition' should not be configured. Possible values: [\"PROVISIONED\", \"FREE_INSTANCE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "An object containing a list of \"key\": value pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "A unique identifier for the instance, which cannot be changed after\nthe instance is created. The name must be between 6 and 30 characters\nin length.\nIf not provided, a random string starting with 'tf-' will be selected.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "num_nodes": {
        "computed": true,
        "description": "The number of nodes allocated to this instance. Exactly one of either num_nodes, processing_units or\nautoscaling_config must be present in terraform except when instance_type = FREE_INSTANCE.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "processing_units": {
        "computed": true,
        "description": "The number of processing units allocated to this instance. Exactly one of either num_nodes,\nprocessing_units or autoscaling_config must be present in terraform except when instance_type = FREE_INSTANCE.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Instance status: 'CREATING' or 'READY'.",
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
    "block_types": {
      "autoscaling_config": {
        "block": {
          "block_types": {
            "asymmetric_autoscaling_options": {
              "block": {
                "block_types": {
                  "overrides": {
                    "block": {
                      "block_types": {
                        "autoscaling_limits": {
                          "block": {
                            "attributes": {
                              "max_nodes": {
                                "description": "The maximum number of nodes for this specific replica.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "min_nodes": {
                                "description": "The minimum number of nodes for this specific replica.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description": "A nested object resource.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "A nested object resource.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "replica_selection": {
                    "block": {
                      "attributes": {
                        "location": {
                          "description": "The location of the replica to apply asymmetric autoscaling options.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "A nested object resource.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Asymmetric autoscaling options for specific replicas.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "autoscaling_limits": {
              "block": {
                "attributes": {
                  "max_nodes": {
                    "description": "Specifies maximum number of nodes allocated to the instance. If set, this number\nshould be greater than or equal to min_nodes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_processing_units": {
                    "description": "Specifies maximum number of processing units allocated to the instance.\nIf set, this number should be multiples of 1000 and be greater than or equal to\nmin_processing_units.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_nodes": {
                    "description": "Specifies number of nodes allocated to the instance. If set, this number\nshould be greater than or equal to 1.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_processing_units": {
                    "description": "Specifies minimum number of processing units allocated to the instance.\nIf set, this number should be multiples of 1000.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Defines scale in controls to reduce the risk of response latency\nand outages due to abrupt scale-in events. Users can define the minimum and\nmaximum compute capacity allocated to the instance, and the autoscaler will\nonly scale within that range. Users can either use nodes or processing\nunits to specify the limits, but should use the same unit to set both the\nmin_limit and max_limit.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "autoscaling_targets": {
              "block": {
                "attributes": {
                  "high_priority_cpu_utilization_percent": {
                    "description": "Specifies the target high priority cpu utilization percentage that the autoscaler\nshould be trying to achieve for the instance.\nThis number is on a scale from 0 (no utilization) to 100 (full utilization)..",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "storage_utilization_percent": {
                    "description": "Specifies the target storage utilization percentage that the autoscaler\nshould be trying to achieve for the instance.\nThis number is on a scale from 0 (no utilization) to 100 (full utilization).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Defines scale in controls to reduce the risk of response latency\nand outages due to abrupt scale-in events",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The autoscaling configuration. Autoscaling is enabled if this field is set.\nExactly one of either num_nodes, processing_units or autoscaling_config must be\npresent in terraform except when instance_type = FREE_INSTANCE.\nWhen autoscaling is enabled, num_nodes and processing_units are treated as,\nOUTPUT_ONLY fields and reflect the current compute capacity allocated to\nthe instance.",
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

func GoogleSpannerInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSpannerInstance), &result)
	return &result
}
