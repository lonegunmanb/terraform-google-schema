package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataprocSessionTemplate = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time when the session template was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "creator": {
        "computed": true,
        "description": "The email address of the user who created the session template.",
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "The labels to associate with this session template.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location in which the session template will be created in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the session template in the following format:\nprojects/{project}/locations/{location}/sessionTemplates/{template_id}",
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
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "The time when the session template was updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "uuid": {
        "computed": true,
        "description": "A session template UUID (Unique Universal Identifier). The service generates this value when it creates the session template.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "environment_config": {
        "block": {
          "block_types": {
            "execution_config": {
              "block": {
                "attributes": {
                  "kms_key": {
                    "description": "The Cloud KMS key to use for encryption.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "network_tags": {
                    "description": "Tags used for network traffic control.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "service_account": {
                    "computed": true,
                    "description": "Service account that used to execute workload.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "staging_bucket": {
                    "description": "A Cloud Storage bucket used to stage workload dependencies, config files, and store\nworkload output and other ephemeral data, such as Spark history files. If you do not specify a staging bucket,\nCloud Dataproc will determine a Cloud Storage location according to the region where your workload is running,\nand then create and manage project-level, per-location staging and temporary buckets.\nThis field requires a Cloud Storage bucket name, not a gs://... URI to a Cloud Storage bucket.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subnetwork_uri": {
                    "description": "Subnetwork configuration for workload execution.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ttl": {
                    "computed": true,
                    "description": "The duration after which the workload will be terminated.\nWhen the workload exceeds this duration, it will be unconditionally terminated without waiting for ongoing\nwork to finish. If ttl is not specified for a session workload, the workload will be allowed to run until it\nexits naturally (or run forever without exiting). If ttl is not specified for an interactive session,\nit defaults to 24 hours. If ttl is not specified for a batch that uses 2.1+ runtime version, it defaults to 4 hours.\nMinimum value is 10 minutes; maximum value is 14 days. If both ttl and idleTtl are specified (for an interactive session),\nthe conditions are treated as OR conditions: the workload will be terminated when it has been idle for idleTtl or\nwhen ttl has been exceeded, whichever occurs first.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Execution configuration for a workload.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "peripherals_config": {
              "block": {
                "attributes": {
                  "metastore_service": {
                    "description": "Resource name of an existing Dataproc Metastore service.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "spark_history_server_config": {
                    "block": {
                      "attributes": {
                        "dataproc_cluster": {
                          "description": "Resource name of an existing Dataproc Cluster to act as a Spark History Server for the workload.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The Spark History Server configuration for the workload.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Peripherals configuration that workload has access to.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Environment configuration for the session execution.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "jupyter_session": {
        "block": {
          "attributes": {
            "display_name": {
              "description": "Display name, shown in the Jupyter kernelspec card.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kernel": {
              "description": "Kernel to be used with Jupyter interactive session. Possible values: [\"PYTHON\", \"SCALA\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Jupyter configuration for an interactive session.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "runtime_config": {
        "block": {
          "attributes": {
            "container_image": {
              "description": "Optional custom container image for the job runtime environment. If not specified, a default container image will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "effective_properties": {
              "computed": true,
              "description": "A mapping of property names to values, which are used to configure workload execution.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "properties": {
              "description": "A mapping of property names to values, which are used to configure workload execution.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "version": {
              "description": "Version of the session runtime.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Runtime configuration for the session template.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "spark_connect_session": {
        "block": {
          "description": "Spark connect configuration for an interactive session.",
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

func GoogleDataprocSessionTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataprocSessionTemplate), &result)
	return &result
}
