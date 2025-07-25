package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudRunV2Job = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects.\n\nCloud Run API v2 does not support annotations with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected on new resources.\nAll system annotations in v1 now have a corresponding field in v2 Job.\n\nThis field follows Kubernetes annotations' namespacing, limits, and rules.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "client": {
        "description": "Arbitrary identifier for the API client.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_version": {
        "description": "Arbitrary version identifier for the API client.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "conditions": {
        "computed": true,
        "description": "The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Job does not reach its desired state. See comments in reconciling for additional information on 'reconciliation' process in Cloud Run.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "execution_reason": "string",
              "last_transition_time": "string",
              "message": "string",
              "reason": "string",
              "revision_reason": "string",
              "severity": "string",
              "state": "string",
              "type": "string"
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description": "The creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "creator": {
        "computed": true,
        "description": "Email address of the authenticated creator.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "The deletion time.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "description": "Whether Terraform will be prevented from destroying the job. Defaults to true.\nWhen a'terraform destroy' or 'terraform apply' would delete the job,\nthe command will fail if this field is not set to false in Terraform state.\nWhen the field is set to true or unset in Terraform state, a 'terraform apply'\nor 'terraform destroy' that would delete the job will fail.\nWhen the field is set to false, deleting the job is allowed.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "effective_annotations": {
        "computed": true,
        "description": "All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
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
      "etag": {
        "computed": true,
        "description": "A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "execution_count": {
        "computed": true,
        "description": "Number of executions created for this job.",
        "description_kind": "plain",
        "type": "number"
      },
      "expire_time": {
        "computed": true,
        "description": "For a deleted resource, the time after which it will be permanently deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "generation": {
        "computed": true,
        "description": "A number that monotonically increases every time the user modifies the desired state.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Unstructured key value map that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component,\nenvironment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels.\n\nCloud Run API v2 does not support labels with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.\nAll system labels in v1 now have a corresponding field in v2 Job.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "last_modifier": {
        "computed": true,
        "description": "Email address of the last authenticated modifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "latest_created_execution": {
        "computed": true,
        "description": "Name of the last created execution.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "completion_time": "string",
              "create_time": "string",
              "name": "string"
            }
          ]
        ]
      },
      "launch_stage": {
        "computed": true,
        "description": "The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/products#product-launch-stages). Cloud Run supports ALPHA, BETA, and GA.\nIf no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features.\n\nFor example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output. Possible values: [\"UNIMPLEMENTED\", \"PRELAUNCH\", \"EARLY_ACCESS\", \"ALPHA\", \"BETA\", \"GA\", \"DEPRECATED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the cloud run job",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the Job.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "observed_generation": {
        "computed": true,
        "description": "The generation of this Job. See comments in reconciling for additional information on reconciliation process in Cloud Run.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "Returns true if the Job is currently being acted upon by the system to bring it into the desired state.\n\nWhen a new Job is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Job to the desired state. This process is called reconciliation. While reconciliation is in process, observedGeneration and latest_succeeded_execution, will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the state matches the Job, or there was an error, and reconciliation failed. This state can be found in terminalCondition.state.\n\nIf reconciliation succeeded, the following fields will match: observedGeneration and generation, latest_succeeded_execution and latestCreatedExecution.\n\nIf reconciliation failed, observedGeneration and latest_succeeded_execution will have the state of the last succeeded execution or empty for newly created Job. Additional information on the failure can be found in terminalCondition and conditions",
        "description_kind": "plain",
        "type": "bool"
      },
      "terminal_condition": {
        "computed": true,
        "description": "The Condition of this Job, containing its readiness status, and detailed error information in case it did not reach the desired state",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "execution_reason": "string",
              "last_transition_time": "string",
              "message": "string",
              "reason": "string",
              "revision_reason": "string",
              "severity": "string",
              "state": "string",
              "type": "string"
            }
          ]
        ]
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
      "uid": {
        "computed": true,
        "description": "Server assigned unique identifier for the Execution. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The last-modified time.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "binary_authorization": {
        "block": {
          "attributes": {
            "breakglass_justification": {
              "description": "If present, indicates to use Breakglass using this justification. If useDefault is False, then it must be empty. For more information on breakglass, see https://cloud.google.com/binary-authorization/docs/using-breakglass",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "policy": {
              "description": "The path to a binary authorization policy. Format: projects/{project}/platforms/cloudRun/{policy-name}",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_default": {
              "description": "If True, indicates to use the default project's binary authorization policy. If False, binary authorization will be disabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Settings for the Binary Authorization feature.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "template": {
        "block": {
          "attributes": {
            "annotations": {
              "description": "Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects.\n\nCloud Run API v2 does not support annotations with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.\nAll system annotations in v1 now have a corresponding field in v2 ExecutionTemplate.\n\nThis field follows Kubernetes annotations' namespacing, limits, and rules.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "labels": {
              "description": "Unstructured key value map that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter,\nor break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or\nhttps://cloud.google.com/run/docs/configuring/labels.\n\nCloud Run API v2 does not support labels with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.\nAll system labels in v1 now have a corresponding field in v2 ExecutionTemplate.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "parallelism": {
              "computed": true,
              "description": "Specifies the maximum desired number of tasks the execution should run at given time. Must be \u003c= taskCount. When the job is run, if this field is 0 or unset, the maximum possible value will be used for that execution. The actual number of tasks running in steady state will be less than this number when there are fewer tasks waiting to be completed remaining, i.e. when the work left to do is less than max parallelism.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "task_count": {
              "computed": true,
              "description": "Specifies the desired number of tasks the execution should run. Setting to 1 means that parallelism is limited to 1 and the success of that task signals the success of the execution. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "template": {
              "block": {
                "attributes": {
                  "encryption_key": {
                    "description": "A reference to a customer managed encryption key (CMEK) to use to encrypt this container image. For more information, go to https://cloud.google.com/run/docs/securing/using-cmek",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "execution_environment": {
                    "computed": true,
                    "description": "The execution environment being used to host this Task. Possible values: [\"EXECUTION_ENVIRONMENT_GEN1\", \"EXECUTION_ENVIRONMENT_GEN2\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "max_retries": {
                    "description": "Number of retries allowed per Task, before marking this Task failed. Defaults to 3. Minimum value is 0.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "service_account": {
                    "computed": true,
                    "description": "Email address of the IAM service account associated with the Task of a Job. The service account represents the identity of the running task, and determines what permissions the task has. If not provided, the task will use the project's default service account.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "timeout": {
                    "computed": true,
                    "description": "Max allowed time duration the Task may be active before the system will actively try to mark it failed and kill associated containers. This applies per attempt of a task, meaning each retry can run for the full timeout.\n\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "containers": {
                    "block": {
                      "attributes": {
                        "args": {
                          "description": "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references are not supported in Cloud Run.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "command": {
                          "description": "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "depends_on": {
                          "description": "Names of the containers that must start before this container.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "image": {
                          "description": "URL of the Container image in Google Container Registry or Google Artifact Registry. More info: https://kubernetes.io/docs/concepts/containers/images",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "name": {
                          "description": "Name of the container specified as a DNS_LABEL.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "working_dir": {
                          "description": "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "env": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the environment variable. Must be a C_IDENTIFIER, and mnay not exceed 32768 characters.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "Literal value of the environment variable. Defaults to \"\" and the maximum allowed length is 32768 characters. Variable references are not supported in Cloud Run.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "value_source": {
                                "block": {
                                  "block_types": {
                                    "secret_key_ref": {
                                      "block": {
                                        "attributes": {
                                          "secret": {
                                            "description": "The name of the secret in Cloud Secret Manager. Format: {secretName} if the secret is in the same project. projects/{project}/secrets/{secretName} if the secret is in a different project.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "version": {
                                            "description": "The Cloud Secret Manager secret version. Can be 'latest' for the latest value or an integer for a specific version.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Selects a secret and a specific version from Cloud Secret Manager.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Source for the environment variable's value.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "List of environment variables to set in the container.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "set"
                        },
                        "ports": {
                          "block": {
                            "attributes": {
                              "container_port": {
                                "description": "Port number the container listens on. This must be a valid TCP port number, 0 \u003c containerPort \u003c 65536.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "name": {
                                "description": "If specified, used to specify which protocol to use. Allowed values are \"http1\" and \"h2c\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "List of ports to expose from the container. Only a single port can be specified. The specified ports must be listening on all interfaces (0.0.0.0) within the container to be accessible.\n\nIf omitted, a port number will be chosen and passed to the container through the PORT environment variable for the container to listen on",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "resources": {
                          "block": {
                            "attributes": {
                              "limits": {
                                "computed": true,
                                "description": "Only memory, CPU, and nvidia.com/gpu are supported. Use key 'cpu' for CPU limit, 'memory' for memory limit, 'nvidia.com/gpu' for gpu limit. Note: The only supported values for CPU are '1', '2', '4', and '8'. Setting 4 CPU requires at least 2Gi of memory. The values of the map is string form of the 'quantity' k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "description": "Compute Resource requirements by this container. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "startup_probe": {
                          "block": {
                            "attributes": {
                              "failure_threshold": {
                                "description": "Minimum consecutive failures for the probe to be considered failed after\nhaving succeeded. Defaults to 3. Minimum value is 1.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "initial_delay_seconds": {
                                "description": "Number of seconds after the container has started before the probe is\ninitiated.\nDefaults to 0 seconds. Minimum value is 0. Maximum value is 240.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "period_seconds": {
                                "description": "How often (in seconds) to perform the probe.\nDefault to 10 seconds. Minimum value is 1. Maximum value is 240.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "timeout_seconds": {
                                "description": "Number of seconds after which the probe times out.\nDefaults to 1 second. Minimum value is 1. Maximum value is 3600.\nMust be smaller than periodSeconds.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "grpc": {
                                "block": {
                                  "attributes": {
                                    "port": {
                                      "computed": true,
                                      "description": "Port number to access on the container. Number must be in the range 1 to 65535.\nIf not specified, defaults to the same value as container.ports[0].containerPort.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "service": {
                                      "description": "The name of the service to place in the gRPC HealthCheckRequest\n(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).\nIf this is not specified, the default behavior is defined by gRPC.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "GRPC specifies an action involving a GRPC port.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "http_get": {
                                "block": {
                                  "attributes": {
                                    "path": {
                                      "description": "Path to access on the HTTP server. If set, it should not be empty string.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "port": {
                                      "computed": true,
                                      "description": "Port number to access on the container. Number must be in the range 1 to 65535.\nIf not specified, defaults to the same value as container.ports[0].containerPort.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "http_headers": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "The header field name.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "description": "The header field value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Custom headers to set in the request. HTTP allows repeated headers.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "HttpGet specifies the http request to perform.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "tcp_socket": {
                                "block": {
                                  "attributes": {
                                    "port": {
                                      "computed": true,
                                      "description": "Port number to access on the container. Number must be in the range 1 to 65535.\nIf not specified, defaults to the same value as container.ports[0].containerPort.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "description": "TcpSocket specifies an action involving a TCP port.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Startup probe of application within the container.\nAll other probes are disabled if a startup probe is provided, until it\nsucceeds. Container will not be added to service endpoints if the probe fails.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "volume_mounts": {
                          "block": {
                            "attributes": {
                              "mount_path": {
                                "description": "Path within the container at which the volume should be mounted. Must not contain ':'. For Cloud SQL volumes, it can be left empty, or must otherwise be /cloudsql. All instances defined in the Volume will be available as /cloudsql/[instance]. For more information on Cloud SQL volumes, visit https://cloud.google.com/sql/docs/mysql/connect-run",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "name": {
                                "description": "This must match the Name of a Volume.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Volume to mount into the container's filesystem.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Holds the single container that defines the unit of execution for this task.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "node_selector": {
                    "block": {
                      "attributes": {
                        "accelerator": {
                          "description": "The GPU to attach to an instance. See https://cloud.google.com/run/docs/configuring/jobs/gpu for configuring GPU.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Node Selector describes the hardware requirements of the resources.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "volumes": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Volume's name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "cloud_sql_instance": {
                          "block": {
                            "attributes": {
                              "instances": {
                                "description": "The Cloud SQL instance connection names, as can be found in https://console.cloud.google.com/sql/instances. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run. Format: {project}:{location}:{instance}",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "For Cloud SQL volumes, contains the specific instances that should be mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "empty_dir": {
                          "block": {
                            "attributes": {
                              "medium": {
                                "description": "The different types of medium supported for EmptyDir. Default value: \"MEMORY\" Possible values: [\"MEMORY\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "size_limit": {
                                "description": "Limit on the storage usable by this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. This field's values are of the 'Quantity' k8s type: https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Ephemeral storage used as a shared volume.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "gcs": {
                          "block": {
                            "attributes": {
                              "bucket": {
                                "description": "Name of the cloud storage bucket to back the volume. The resource service account must have permission to access the bucket.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "read_only": {
                                "description": "If true, mount this volume as read-only in all mounts. If false, mount this volume as read-write.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "Cloud Storage bucket mounted as a volume using GCSFuse.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "nfs": {
                          "block": {
                            "attributes": {
                              "path": {
                                "description": "Path that is exported by the NFS server.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "read_only": {
                                "description": "If true, mount this volume as read-only in all mounts.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "server": {
                                "description": "Hostname or IP address of the NFS server.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "NFS share mounted as a volume.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "secret": {
                          "block": {
                            "attributes": {
                              "default_mode": {
                                "description": "Integer representation of mode bits to use on created files by default. Must be a value between 0000 and 0777 (octal), defaulting to 0444. Directories within the path are not affected by this setting.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "secret": {
                                "description": "The name of the secret in Cloud Secret Manager. Format: {secret} if the secret is in the same project. projects/{project}/secrets/{secret} if the secret is in a different project.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "items": {
                                "block": {
                                  "attributes": {
                                    "mode": {
                                      "description": "Integer octal mode bits to use on this file, must be a value between 01 and 0777 (octal). If 0 or not set, the Volume's default mode will be used.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "path": {
                                      "description": "The relative path of the secret in the container.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "version": {
                                      "description": "The Cloud Secret Manager secret version. Can be 'latest' for the latest value or an integer for a specific version",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "If unspecified, the volume will expose a file whose name is the secret, relative to VolumeMount.mount_path. If specified, the key will be used as the version to fetch from Cloud Secret Manager and the path will be the name of the file exposed in the volume. When items are defined, they must specify a path and a version.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "A list of Volumes to make available to containers.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "vpc_access": {
                    "block": {
                      "attributes": {
                        "connector": {
                          "description": "VPC Access connector name. Format: projects/{project}/locations/{location}/connectors/{connector}, where {project} can be project id or number.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "egress": {
                          "computed": true,
                          "description": "Traffic VPC egress settings. Possible values: [\"ALL_TRAFFIC\", \"PRIVATE_RANGES_ONLY\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "network_interfaces": {
                          "block": {
                            "attributes": {
                              "network": {
                                "computed": true,
                                "description": "The VPC network that the Cloud Run resource will be able to send traffic to. At least one of network or subnetwork must be specified. If both\nnetwork and subnetwork are specified, the given VPC subnetwork must belong to the given VPC network. If network is not specified, it will be\nlooked up from the subnetwork.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "subnetwork": {
                                "computed": true,
                                "description": "The VPC subnetwork that the Cloud Run resource will get IPs from. At least one of network or subnetwork must be specified. If both\nnetwork and subnetwork are specified, the given VPC subnetwork must belong to the given VPC network. If subnetwork is not specified, the\nsubnetwork with the same name with the network will be used.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "tags": {
                                "description": "Network tags applied to this Cloud Run job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "Direct VPC egress settings. Currently only single network interface is supported.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "VPC Access configuration to use for this Task. For more information, visit https://cloud.google.com/run/docs/configuring/connecting-vpc.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Describes the task(s) that will be created when executing an execution",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The template used to create executions for this Job.",
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

func GoogleCloudRunV2JobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudRunV2Job), &result)
	return &result
}
