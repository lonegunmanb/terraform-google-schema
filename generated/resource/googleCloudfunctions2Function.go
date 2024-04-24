package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudfunctions2Function = `{
  "block": {
    "attributes": {
      "description": {
        "description": "User-provided description of a function.",
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
      "environment": {
        "computed": true,
        "description": "The environment the function is hosted on.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_name": {
        "description": "Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.\nIt must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "A set of key/value label pairs associated with this Cloud Function.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of this cloud function.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "A user-defined name of the function. Function names must\nbe unique globally and match pattern 'projects/*/locations/*/functions/*'.",
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
      "state": {
        "computed": true,
        "description": "Describes the current state of the function.",
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
      "update_time": {
        "computed": true,
        "description": "The last update timestamp of a Cloud Function.",
        "description_kind": "plain",
        "type": "string"
      },
      "url": {
        "computed": true,
        "description": "Output only. The deployed url for the function.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "build_config": {
        "block": {
          "attributes": {
            "build": {
              "computed": true,
              "description": "The Cloud Build name of the latest successful\ndeployment of the function.",
              "description_kind": "plain",
              "type": "string"
            },
            "docker_repository": {
              "computed": true,
              "description": "User managed repository created in Artifact Registry optionally with a customer managed encryption key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "entry_point": {
              "description": "The name of the function (as defined in source code) that will be executed.\nDefaults to the resource name suffix, if not specified. For backward\ncompatibility, if function with given name is not found, then the system\nwill try to use function named \"function\". For Node.js this is name of a\nfunction exported by the module specified in source_location.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "environment_variables": {
              "computed": true,
              "description": "User-provided build-time environment variables for the function.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "runtime": {
              "description": "The runtime in which to run the function. Required when deploying a new\nfunction, optional when updating an existing function.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_account": {
              "computed": true,
              "description": "The fully-qualified name of the service account to be used for building the container.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "worker_pool": {
              "description": "Name of the Cloud Build Custom Worker Pool that should be used to build the function.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "source": {
              "block": {
                "block_types": {
                  "repo_source": {
                    "block": {
                      "attributes": {
                        "branch_name": {
                          "description": "Regex matching branches to build.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "commit_sha": {
                          "description": "Regex matching tags to build.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "dir": {
                          "description": "Directory, relative to the source root, in which to run the build.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "invert_regex": {
                          "description": "Only trigger a build if the revision regex does\nNOT match the revision regex.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "project_id": {
                          "description": "ID of the project that owns the Cloud Source Repository. If omitted, the\nproject ID requesting the build is assumed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "repo_name": {
                          "description": "Name of the Cloud Source Repository.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "tag_name": {
                          "description": "Regex matching tags to build.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "If provided, get the source from this location in a Cloud Source Repository.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "storage_source": {
                    "block": {
                      "attributes": {
                        "bucket": {
                          "description": "Google Cloud Storage bucket containing the source",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "generation": {
                          "computed": true,
                          "description": "Google Cloud Storage generation for the object. If the generation\nis omitted, the latest generation will be used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "object": {
                          "description": "Google Cloud Storage object containing the source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "If provided, get the source from this location in Google Cloud Storage.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The location of the function source code.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Describes the Build step of the function that builds a container\nfrom the given source.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "event_trigger": {
        "block": {
          "attributes": {
            "event_type": {
              "description": "Required. The type of event to observe.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pubsub_topic": {
              "computed": true,
              "description": "The name of a Pub/Sub topic in the same project that will be used\nas the transport topic for the event delivery.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "retry_policy": {
              "description": "Describes the retry policy in case of function's execution failure.\nRetried execution is charged as any other execution. Possible values: [\"RETRY_POLICY_UNSPECIFIED\", \"RETRY_POLICY_DO_NOT_RETRY\", \"RETRY_POLICY_RETRY\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_account_email": {
              "computed": true,
              "description": "Optional. The email of the trigger's service account. The service account\nmust have permission to invoke Cloud Run services. If empty, defaults to the\nCompute Engine default service account: {project_number}-compute@developer.gserviceaccount.com.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "trigger": {
              "computed": true,
              "description": "Output only. The resource name of the Eventarc trigger.",
              "description_kind": "plain",
              "type": "string"
            },
            "trigger_region": {
              "computed": true,
              "description": "The region that the trigger will be in. The trigger will only receive\nevents originating in this region. It can be the same\nregion as the function, a different region or multi-region, or the global\nregion. If not provided, defaults to the same region as the function.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "event_filters": {
              "block": {
                "attributes": {
                  "attribute": {
                    "description": "'Required. The name of a CloudEvents attribute.\nCurrently, only a subset of attributes are supported for filtering. Use the 'gcloud eventarc providers describe' command to learn more about events and their attributes.\nDo not filter for the 'type' attribute here, as this is already achieved by the resource's 'event_type' attribute.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "operator": {
                    "description": "Optional. The operator used for matching the events with the value of\nthe filter. If not specified, only events that have an exact key-value\npair specified in the filter are matched.\nThe only allowed value is 'match-path-pattern'.\n[See documentation on path patterns here](https://cloud.google.com/eventarc/docs/path-patterns)'",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Required. The value for the attribute.\nIf the operator field is set as 'match-path-pattern', this value can be a path pattern instead of an exact value.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Criteria used to filter events.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "An Eventarc trigger managed by Google Cloud Functions that fires events in\nresponse to a condition in another service.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "service_config": {
        "block": {
          "attributes": {
            "all_traffic_on_latest_revision": {
              "description": "Whether 100% of traffic is routed to the latest revision. Defaults to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "available_cpu": {
              "computed": true,
              "description": "The number of CPUs used in a single container instance. Default value is calculated from available memory.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "available_memory": {
              "computed": true,
              "description": "The amount of memory available for a function.\nDefaults to 256M. Supported units are k, M, G, Mi, Gi. If no unit is\nsupplied the value is interpreted as bytes.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "environment_variables": {
              "description": "Environment variables that shall be available during function execution.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "gcf_uri": {
              "computed": true,
              "description": "URIs of the Service deployed",
              "description_kind": "plain",
              "type": "string"
            },
            "ingress_settings": {
              "description": "Available ingress settings. Defaults to \"ALLOW_ALL\" if unspecified. Default value: \"ALLOW_ALL\" Possible values: [\"ALLOW_ALL\", \"ALLOW_INTERNAL_ONLY\", \"ALLOW_INTERNAL_AND_GCLB\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_instance_count": {
              "computed": true,
              "description": "The limit on the maximum number of function instances that may coexist at a\ngiven time.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_instance_request_concurrency": {
              "computed": true,
              "description": "Sets the maximum number of concurrent requests that each instance can receive. Defaults to 1.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_instance_count": {
              "description": "The limit on the minimum number of function instances that may coexist at a\ngiven time.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "service": {
              "computed": true,
              "description": "Name of the service associated with a Function.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_account_email": {
              "computed": true,
              "description": "The email of the service account for this function.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "timeout_seconds": {
              "computed": true,
              "description": "The function execution timeout. Execution is considered failed and\ncan be terminated if the function is not completed at the end of the\ntimeout period. Defaults to 60 seconds.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "uri": {
              "computed": true,
              "description": "URI of the Service deployed.",
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_connector": {
              "description": "The Serverless VPC Access connector that this cloud function can connect to.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_connector_egress_settings": {
              "description": "Available egress settings. Possible values: [\"VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED\", \"PRIVATE_RANGES_ONLY\", \"ALL_TRAFFIC\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "secret_environment_variables": {
              "block": {
                "attributes": {
                  "key": {
                    "description": "Name of the environment variable.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "project_id": {
                    "description": "Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "secret": {
                    "description": "Name of the secret in secret manager (not the full resource name).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "version": {
                    "description": "Version of the secret (version number or the string 'latest'). It is recommended to use a numeric version for secret environment variables as any updates to the secret value is not reflected until new instances start.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Secret environment variables configuration.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "secret_volumes": {
              "block": {
                "attributes": {
                  "mount_path": {
                    "description": "The path within the container to mount the secret volume. For example, setting the mountPath as /etc/secrets would mount the secret value files under the /etc/secrets directory. This directory will also be completely shadowed and unavailable to mount any other secrets. Recommended mount path: /etc/secrets",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "project_id": {
                    "description": "Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "secret": {
                    "description": "Name of the secret in secret manager (not the full resource name).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "versions": {
                    "block": {
                      "attributes": {
                        "path": {
                          "description": "Relative path of the file under the mount path where the secret value for this version will be fetched and made available. For example, setting the mountPath as '/etc/secrets' and path as secret_foo would mount the secret value file at /etc/secrets/secret_foo.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "version": {
                          "description": "Version of the secret (version number or the string 'latest'). It is preferable to use latest version with secret volumes as secret value changes are reflected immediately.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "List of secret versions to mount for this secret. If empty, the latest version of the secret will be made available in a file named after the secret under the mount point.'",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Secret volumes configuration.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Describes the Service being deployed.",
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

func GoogleCloudfunctions2FunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudfunctions2Function), &result)
	return &result
}
