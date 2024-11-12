package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataprocBatch = `{
  "block": {
    "attributes": {
      "batch_id": {
        "description": "The ID to use for the batch, which will become the final component of the batch's resource name.\nThis value must be 4-63 characters. Valid characters are /[a-z][0-9]-/.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time when the batch was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "creator": {
        "computed": true,
        "description": "The email address of the user who created the batch.",
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
        "description": "The labels to associate with this batch.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location in which the batch will be created in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the batch.",
        "description_kind": "plain",
        "type": "string"
      },
      "operation": {
        "computed": true,
        "description": "The resource name of the operation associated with this batch.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime_info": {
        "computed": true,
        "description": "Runtime information about batch execution.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "approximate_usage": [
                "list",
                [
                  "object",
                  {
                    "accelerator_type": "string",
                    "milli_accelerator_seconds": "string",
                    "milli_dcu_seconds": "string",
                    "shuffle_storage_gb_seconds": "string"
                  }
                ]
              ],
              "current_usage": [
                "list",
                [
                  "object",
                  {
                    "accelerator_type": "string",
                    "milli_accelerator": "string",
                    "milli_dcu": "string",
                    "milli_dcu_premium": "string",
                    "shuffle_storage_gb": "string",
                    "shuffle_storage_gb_premium": "string",
                    "snapshot_time": "string"
                  }
                ]
              ],
              "diagnostic_output_uri": "string",
              "endpoints": [
                "map",
                "string"
              ],
              "output_uri": "string"
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description": "The state of the batch. For possible values, see the [API documentation](https://cloud.google.com/dataproc-serverless/docs/reference/rest/v1/projects.locations.batches#State).",
        "description_kind": "plain",
        "type": "string"
      },
      "state_history": {
        "computed": true,
        "description": "Historical state information for the batch.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "state": "string",
              "state_message": "string",
              "state_start_time": "string"
            }
          ]
        ]
      },
      "state_message": {
        "computed": true,
        "description": "Batch state details, such as a failure description if the state is FAILED.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_time": {
        "computed": true,
        "description": "Batch state details, such as a failure description if the state is FAILED.",
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
      "uuid": {
        "computed": true,
        "description": "A batch UUID (Unique Universal Identifier). The service generates this value when it creates the batch.",
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
                  "network_uri": {
                    "description": "Network configuration for workload execution.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
                    "description": "The duration after which the workload will be terminated.\nWhen the workload exceeds this duration, it will be unconditionally terminated without waiting for ongoing\nwork to finish. If ttl is not specified for a batch workload, the workload will be allowed to run until it\nexits naturally (or run forever without exiting). If ttl is not specified for an interactive session,\nit defaults to 24 hours. If ttl is not specified for a batch that uses 2.1+ runtime version, it defaults to 4 hours.\nMinimum value is 10 minutes; maximum value is 14 days. If both ttl and idleTtl are specified (for an interactive session),\nthe conditions are treated as OR conditions: the workload will be terminated when it has been idle for idleTtl or\nwhen ttl has been exceeded, whichever occurs first.",
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
          "description": "Environment configuration for the batch execution.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "pyspark_batch": {
        "block": {
          "attributes": {
            "archive_uris": {
              "description": "HCFS URIs of archives to be extracted into the working directory of each executor.\nSupported file types: .jar, .tar, .tar.gz, .tgz, and .zip.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "args": {
              "description": "The arguments to pass to the driver. Do not include arguments that can be set as batch\nproperties, such as --conf, since a collision can occur that causes an incorrect batch submission.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "file_uris": {
              "description": "HCFS URIs of files to be placed in the working directory of each executor.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "jar_file_uris": {
              "description": "HCFS URIs of jar files to add to the classpath of the Spark driver and tasks.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "main_python_file_uri": {
              "description": "The HCFS URI of the main Python file to use as the Spark driver. Must be a .py file.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "python_file_uris": {
              "description": "HCFS file URIs of Python files to pass to the PySpark framework.\nSupported file types: .py, .egg, and .zip.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "PySpark batch config.",
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
              "computed": true,
              "description": "Version of the batch runtime.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Runtime configuration for the batch execution.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "spark_batch": {
        "block": {
          "attributes": {
            "archive_uris": {
              "description": "HCFS URIs of archives to be extracted into the working directory of each executor.\nSupported file types: .jar, .tar, .tar.gz, .tgz, and .zip.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "args": {
              "description": "The arguments to pass to the driver. Do not include arguments that can be set as batch\nproperties, such as --conf, since a collision can occur that causes an incorrect batch submission.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "file_uris": {
              "description": "HCFS URIs of files to be placed in the working directory of each executor.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "jar_file_uris": {
              "description": "HCFS URIs of jar files to add to the classpath of the Spark driver and tasks.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "main_class": {
              "description": "The name of the driver main class. The jar file that contains the class must be in the\nclasspath or specified in jarFileUris.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "main_jar_file_uri": {
              "description": "The HCFS URI of the jar file that contains the main class.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Spark batch config.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "spark_r_batch": {
        "block": {
          "attributes": {
            "archive_uris": {
              "description": "HCFS URIs of archives to be extracted into the working directory of each executor.\nSupported file types: .jar, .tar, .tar.gz, .tgz, and .zip.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "args": {
              "description": "The arguments to pass to the driver. Do not include arguments that can be set as batch\nproperties, such as --conf, since a collision can occur that causes an incorrect batch submission.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "file_uris": {
              "description": "HCFS URIs of files to be placed in the working directory of each executor.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "main_r_file_uri": {
              "description": "The HCFS URI of the main R file to use as the driver. Must be a .R or .r file.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "SparkR batch config.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "spark_sql_batch": {
        "block": {
          "attributes": {
            "jar_file_uris": {
              "description": "HCFS URIs of jar files to be added to the Spark CLASSPATH.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "query_file_uri": {
              "description": "The HCFS URI of the script that contains Spark SQL queries to execute.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "query_variables": {
              "description": "Mapping of query variable names to values (equivalent to the Spark SQL command: SET name=\"value\";).",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "description": "Spark SQL batch config.",
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

func GoogleDataprocBatchSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataprocBatch), &result)
	return &result
}
