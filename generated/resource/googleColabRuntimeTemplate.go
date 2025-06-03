package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleColabRuntimeTemplate = `{
  "block": {
    "attributes": {
      "description": {
        "description": "The description of the Runtime Template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Required. The display name of the Runtime Template.",
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
        "computed": true,
        "description": "Labels to identify and group the runtime template.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource: https://cloud.google.com/colab/docs/locations",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the Runtime Template",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_tags": {
        "description": "Applies the given Compute Engine tags to the runtime.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "project": {
        "computed": true,
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
      "data_persistent_disk_spec": {
        "block": {
          "attributes": {
            "disk_size_gb": {
              "computed": true,
              "description": "The disk size of the runtime in GB. If specified, the diskType must also be specified. The minimum size is 10GB and the maximum is 65536GB.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_type": {
              "computed": true,
              "description": "The type of the persistent disk.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The configuration for the data disk of the runtime.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "encryption_spec": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "The Cloud KMS encryption key (customer-managed encryption key) used to protect the runtime.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Customer-managed encryption key spec for the notebook runtime.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "euc_config": {
        "block": {
          "attributes": {
            "euc_disabled": {
              "computed": true,
              "description": "Disable end user credential access for the runtime.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "EUC configuration of the NotebookRuntimeTemplate.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "idle_shutdown_config": {
        "block": {
          "attributes": {
            "idle_timeout": {
              "computed": true,
              "description": "The duration after which the runtime is automatically shut down. An input of 0s disables the idle shutdown feature, and a valid range is [10m, 24h].",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Notebook Idle Shutdown configuration for the runtime.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "machine_spec": {
        "block": {
          "attributes": {
            "accelerator_count": {
              "computed": true,
              "description": "The number of accelerators used by the runtime.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "accelerator_type": {
              "description": "The type of hardware accelerator used by the runtime. If specified, acceleratorCount must also be specified.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "machine_type": {
              "computed": true,
              "description": "The Compute Engine machine type selected for the runtime.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "'The machine configuration of the runtime.'",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_spec": {
        "block": {
          "attributes": {
            "enable_internet_access": {
              "description": "Enable public internet access for the runtime.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "network": {
              "computed": true,
              "description": "The name of the VPC that this runtime is in.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnetwork": {
              "description": "The name of the subnetwork that this runtime is in.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The network configuration for the runtime.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "shielded_vm_config": {
        "block": {
          "attributes": {
            "enable_secure_boot": {
              "computed": true,
              "description": "Enables secure boot for the runtime.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Runtime Shielded VM spec.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "software_config": {
        "block": {
          "block_types": {
            "env": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name of the environment variable. Must be a valid C identifier.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Variables that reference a $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Environment variables to be passed to the container.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "post_startup_script_config": {
              "block": {
                "attributes": {
                  "post_startup_script": {
                    "description": "Post startup script to run after runtime is started.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "post_startup_script_behavior": {
                    "description": "Post startup script behavior that defines download and execution behavior. Possible values: [\"RUN_ONCE\", \"RUN_EVERY_START\", \"DOWNLOAD_AND_RUN_EVERY_START\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "post_startup_script_url": {
                    "description": "Post startup script url to download. Example: https://bucket/script.sh.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "deprecated": true,
                "description": "Post startup script config.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The notebook software configuration of the notebook runtime.",
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

func GoogleColabRuntimeTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleColabRuntimeTemplate), &result)
	return &result
}
