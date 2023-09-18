package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataPipelinePipeline = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The timestamp when the pipeline was initially created. Set by the Data Pipelines service.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "The display name of the pipeline. It can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), and underscores (_).",
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
      "job_count": {
        "computed": true,
        "description": "Number of jobs.",
        "description_kind": "plain",
        "type": "number"
      },
      "last_update_time": {
        "computed": true,
        "description": "The timestamp when the pipeline was last modified. Set by the Data Pipelines service.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "\"The pipeline name. For example': 'projects/PROJECT_ID/locations/LOCATION_ID/pipelines/PIPELINE_ID.\"\n\"- PROJECT_ID can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), and periods (.). For more information, see Identifying projects.\"\n\"LOCATION_ID is the canonical ID for the pipeline's location. The list of available locations can be obtained by calling google.cloud.location.Locations.ListLocations. Note that the Data Pipelines service is not available in all regions. It depends on Cloud Scheduler, an App Engine application, so it's only available in App Engine regions.\"\n\"PIPELINE_ID is the ID of the pipeline. Must be unique for the selected project and location.\"",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pipeline_sources": {
        "description": "The sources of the pipeline (for example, Dataplex). The keys and values are set by the corresponding sources during pipeline creation.\nAn object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "A reference to the region",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheduler_service_account_email": {
        "description": "Optional. A service account email to be used with the Cloud Scheduler job. If not specified, the default compute engine service account will be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "description": "The state of the pipeline. When the pipeline is created, the state is set to 'PIPELINE_STATE_ACTIVE' by default. State changes can be requested by setting the state to stopping, paused, or resuming. State cannot be changed through pipelines.patch requests.\nhttps://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#state Possible values: [\"STATE_UNSPECIFIED\", \"STATE_RESUMING\", \"STATE_ACTIVE\", \"STATE_STOPPING\", \"STATE_ARCHIVED\", \"STATE_PAUSED\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description": "The type of the pipeline. This field affects the scheduling of the pipeline and the type of metrics to show for the pipeline.\nhttps://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#pipelinetype Possible values: [\"PIPELINE_TYPE_UNSPECIFIED\", \"PIPELINE_TYPE_BATCH\", \"PIPELINE_TYPE_STREAMING\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "schedule_info": {
        "block": {
          "attributes": {
            "next_job_time": {
              "computed": true,
              "description": "When the next Scheduler job is going to run.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
              "description_kind": "plain",
              "type": "string"
            },
            "schedule": {
              "description": "Unix-cron format of the schedule. This information is retrieved from the linked Cloud Scheduler.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "time_zone": {
              "description": "Timezone ID. This matches the timezone IDs used by the Cloud Scheduler API. If empty, UTC time is assumed.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Internal scheduling information for a pipeline. If this information is provided, periodic jobs will be created per the schedule. If not, users are responsible for creating jobs externally.\nhttps://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#schedulespec",
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
      },
      "workload": {
        "block": {
          "block_types": {
            "dataflow_flex_template_request": {
              "block": {
                "attributes": {
                  "location": {
                    "description": "The regional endpoint to which to direct the request. For example, us-central1, us-west1.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "project_id": {
                    "description": "The ID of the Cloud Platform project that the job belongs to.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "validate_only": {
                    "description": "If true, the request is validated but not actually executed. Defaults to false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "launch_parameter": {
                    "block": {
                      "attributes": {
                        "container_spec_gcs_path": {
                          "description": "Cloud Storage path to a file with a JSON-serialized ContainerSpec as content.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "job_name": {
                          "description": "The job name to use for the created job. For an update job request, the job name should be the same as the existing running job.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "launch_options": {
                          "description": "Launch options for this Flex Template job. This is a common set of options across languages and templates. This should not be used to pass job parameters.\n'An object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "parameters": {
                          "description": "'The parameters for the Flex Template. Example: {\"numWorkers\":\"5\"}'\n'An object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "transform_name_mappings": {
                          "description": "'Use this to pass transform name mappings for streaming update jobs. Example: {\"oldTransformName\":\"newTransformName\",...}'\n'An object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "update": {
                          "description": "Set this to true if you are sending a request to update a running streaming job. When set, the job name should be the same as the running job.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "environment": {
                          "block": {
                            "attributes": {
                              "additional_experiments": {
                                "description": "Additional experiment flags for the job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "additional_user_labels": {
                                "description": "Additional user labels to be specified for the job. Keys and values should follow the restrictions specified in the labeling restrictions page. An object containing a list of key/value pairs.\n'Example: { \"name\": \"wrench\", \"mass\": \"1kg\", \"count\": \"3\" }.'\n'An object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.'",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "enable_streaming_engine": {
                                "description": "Whether to enable Streaming Engine for the job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "flexrs_goal": {
                                "description": "Set FlexRS goal for the job. https://cloud.google.com/dataflow/docs/guides/flexrs\nhttps://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#FlexResourceSchedulingGoal Possible values: [\"FLEXRS_UNSPECIFIED\", \"FLEXRS_SPEED_OPTIMIZED\", \"FLEXRS_COST_OPTIMIZED\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "ip_configuration": {
                                "description": "Configuration for VM IPs.\nhttps://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#WorkerIPAddressConfiguration Possible values: [\"WORKER_IP_UNSPECIFIED\", \"WORKER_IP_PUBLIC\", \"WORKER_IP_PRIVATE\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "kms_key_name": {
                                "description": "'Name for the Cloud KMS key for the job. The key format is: projects//locations//keyRings//cryptoKeys/'",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "machine_type": {
                                "description": "The machine type to use for the job. Defaults to the value from the template if not specified.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "max_workers": {
                                "description": "The maximum number of Compute Engine instances to be made available to your pipeline during execution, from 1 to 1000.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "network": {
                                "description": "Network to which VMs will be assigned. If empty or unspecified, the service will use the network \"default\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "num_workers": {
                                "description": "The initial number of Compute Engine instances for the job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "service_account_email": {
                                "description": "The email address of the service account to run the job as.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "subnetwork": {
                                "description": "Subnetwork to which VMs will be assigned, if desired. You can specify a subnetwork using either a complete URL or an abbreviated path. Expected to be of the form \"https://www.googleapis.com/compute/v1/projects/HOST_PROJECT_ID/regions/REGION/subnetworks/SUBNETWORK\" or \"regions/REGION/subnetworks/SUBNETWORK\". If the subnetwork is located in a Shared VPC network, you must use the complete URL.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "temp_location": {
                                "description": "The Cloud Storage path to use for temporary files. Must be a valid Cloud Storage URL, beginning with gs://.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "worker_region": {
                                "description": "The Compute Engine region (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. \"us-west1\". Mutually exclusive with workerZone. If neither workerRegion nor workerZone is specified, default to the control plane's region.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "worker_zone": {
                                "description": "The Compute Engine zone (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. \"us-west1-a\". Mutually exclusive with workerRegion. If neither workerRegion nor workerZone is specified, a zone in the control plane's region is chosen based on available capacity. If both workerZone and zone are set, workerZone takes precedence.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "zone": {
                                "description": "The Compute Engine availability zone for launching worker instances to run your pipeline. In the future, workerZone will take precedence.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The runtime environment for the Flex Template job.\nhttps://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#FlexTemplateRuntimeEnvironment",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Parameter to launch a job from a Flex Template.\nhttps://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#launchflextemplateparameter",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Template information and additional parameters needed to launch a Dataflow job using the flex launch API.\nhttps://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#launchflextemplaterequest",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "dataflow_launch_template_request": {
              "block": {
                "attributes": {
                  "gcs_path": {
                    "description": "A Cloud Storage path to the template from which to create the job. Must be a valid Cloud Storage URL, beginning with 'gs://'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "location": {
                    "description": "The regional endpoint to which to direct the request.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "project_id": {
                    "description": "The ID of the Cloud Platform project that the job belongs to.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "validate_only": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "launch_parameters": {
                    "block": {
                      "attributes": {
                        "job_name": {
                          "description": "The job name to use for the created job.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "parameters": {
                          "description": "The runtime parameters to pass to the job.\n'An object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "transform_name_mapping": {
                          "description": "Map of transform name prefixes of the job to be replaced to the corresponding name prefixes of the new job. Only applicable when updating a pipeline.\n'An object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "update": {
                          "description": "If set, replace the existing pipeline with the name specified by jobName with this pipeline, preserving state.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "environment": {
                          "block": {
                            "attributes": {
                              "additional_experiments": {
                                "description": "Additional experiment flags for the job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "additional_user_labels": {
                                "description": "Additional user labels to be specified for the job. Keys and values should follow the restrictions specified in the labeling restrictions page. An object containing a list of key/value pairs.\n'Example: { \"name\": \"wrench\", \"mass\": \"1kg\", \"count\": \"3\" }.'\n'An object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.'",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "bypass_temp_dir_validation": {
                                "description": "Whether to bypass the safety checks for the job's temporary directory. Use with caution.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "enable_streaming_engine": {
                                "description": "Whether to enable Streaming Engine for the job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "ip_configuration": {
                                "description": "Configuration for VM IPs.\nhttps://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#WorkerIPAddressConfiguration Possible values: [\"WORKER_IP_UNSPECIFIED\", \"WORKER_IP_PUBLIC\", \"WORKER_IP_PRIVATE\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "kms_key_name": {
                                "description": "'Name for the Cloud KMS key for the job. The key format is: projects//locations//keyRings//cryptoKeys/'",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "machine_type": {
                                "description": "The machine type to use for the job. Defaults to the value from the template if not specified.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "max_workers": {
                                "description": "The maximum number of Compute Engine instances to be made available to your pipeline during execution, from 1 to 1000.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "network": {
                                "computed": true,
                                "description": "Network to which VMs will be assigned. If empty or unspecified, the service will use the network \"default\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "num_workers": {
                                "description": "The initial number of Compute Engine instances for the job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "service_account_email": {
                                "description": "The email address of the service account to run the job as.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "subnetwork": {
                                "description": "Subnetwork to which VMs will be assigned, if desired. You can specify a subnetwork using either a complete URL or an abbreviated path. Expected to be of the form \"https://www.googleapis.com/compute/v1/projects/HOST_PROJECT_ID/regions/REGION/subnetworks/SUBNETWORK\" or \"regions/REGION/subnetworks/SUBNETWORK\". If the subnetwork is located in a Shared VPC network, you must use the complete URL.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "temp_location": {
                                "description": "The Cloud Storage path to use for temporary files. Must be a valid Cloud Storage URL, beginning with gs://.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "worker_region": {
                                "description": "The Compute Engine region (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. \"us-west1\". Mutually exclusive with workerZone. If neither workerRegion nor workerZone is specified, default to the control plane's region.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "worker_zone": {
                                "description": "The Compute Engine zone (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. \"us-west1-a\". Mutually exclusive with workerRegion. If neither workerRegion nor workerZone is specified, a zone in the control plane's region is chosen based on available capacity. If both workerZone and zone are set, workerZone takes precedence.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "zone": {
                                "description": "The Compute Engine availability zone for launching worker instances to run your pipeline. In the future, workerZone will take precedence.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The runtime environment for the job.\nhttps://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#RuntimeEnvironment",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The parameters of the template to launch. This should be part of the body of the POST request.\nhttps://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#launchtemplateparameters",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Template information and additional parameters needed to launch a Dataflow job using the standard launch API.\nhttps://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#launchtemplaterequest",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Workload information for creating new jobs.\nhttps://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#workload",
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

func GoogleDataPipelinePipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataPipelinePipeline), &result)
	return &result
}
