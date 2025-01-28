package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudRunV2Service = `{
  "block": {
    "attributes": {
      "annotations": {
        "computed": true,
        "description": "Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects.\n\nCloud Run API v2 does not support annotations with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected in new resources.\nAll system annotations in v1 now have a corresponding field in v2 Service.\n\nThis field follows Kubernetes annotations' namespacing, limits, and rules.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "binary_authorization": {
        "computed": true,
        "description": "Settings for the Binary Authorization feature.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "breakglass_justification": "string",
              "policy": "string",
              "use_default": "bool"
            }
          ]
        ]
      },
      "build_config": {
        "computed": true,
        "description": "Configuration for building a Cloud Run function.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "base_image": "string",
              "enable_automatic_updates": "bool",
              "environment_variables": [
                "map",
                "string"
              ],
              "function_target": "string",
              "image_uri": "string",
              "name": "string",
              "service_account": "string",
              "source_location": "string",
              "worker_pool": "string"
            }
          ]
        ]
      },
      "client": {
        "computed": true,
        "description": "Arbitrary identifier for the API client.",
        "description_kind": "plain",
        "type": "string"
      },
      "client_version": {
        "computed": true,
        "description": "Arbitrary version identifier for the API client.",
        "description_kind": "plain",
        "type": "string"
      },
      "conditions": {
        "computed": true,
        "description": "The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Service does not reach its Serving state. See comments in reconciling for additional information on reconciliation process in Cloud Run.",
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
      "custom_audiences": {
        "computed": true,
        "description": "One or more custom audiences that you want this service to support. Specify each custom audience as the full URL in a string. The custom audiences are encoded in the token and used to authenticate requests.\nFor more information, see https://cloud.google.com/run/docs/configuring/custom-audiences.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "delete_time": {
        "computed": true,
        "description": "The deletion time.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Whether Terraform will be prevented from destroying the service. Defaults to true.\nWhen a'terraform destroy' or 'terraform apply' would delete the service,\nthe command will fail if this field is not set to false in Terraform state.\nWhen the field is set to true or unset in Terraform state, a 'terraform apply'\nor 'terraform destroy' that would delete the service will fail.\nWhen the field is set to false, deleting the service is allowed.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "User-provided description of the Service. This field currently has a 512-character limit.",
        "description_kind": "plain",
        "type": "string"
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
      "expire_time": {
        "computed": true,
        "description": "For a deleted resource, the time after which it will be permanently deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "generation": {
        "computed": true,
        "description": "A number that monotonically increases every time the user modifies the desired state. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a string instead of an integer.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ingress": {
        "computed": true,
        "description": "Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active. Possible values: [\"INGRESS_TRAFFIC_ALL\", \"INGRESS_TRAFFIC_INTERNAL_ONLY\", \"INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "invoker_iam_disabled": {
        "computed": true,
        "description": "Disables IAM permission check for run.routes.invoke for callers of this service. This feature is available by invitation only. For more information, visit https://cloud.google.com/run/docs/securing/managing-access#invoker_check.",
        "description_kind": "plain",
        "type": "bool"
      },
      "labels": {
        "computed": true,
        "description": "Unstructured key value map that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component,\nenvironment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels.\n\nCloud Run API v2 does not support labels with  'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.\nAll system labels in v1 now have a corresponding field in v2 Service.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
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
      "latest_created_revision": {
        "computed": true,
        "description": "Name of the last created revision. See comments in reconciling for additional information on reconciliation process in Cloud Run.",
        "description_kind": "plain",
        "type": "string"
      },
      "latest_ready_revision": {
        "computed": true,
        "description": "Name of the latest revision that is serving traffic. See comments in reconciling for additional information on reconciliation process in Cloud Run.",
        "description_kind": "plain",
        "type": "string"
      },
      "launch_stage": {
        "computed": true,
        "description": "The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/products#product-launch-stages). Cloud Run supports ALPHA, BETA, and GA.\nIf no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features.\n\nFor example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output. Possible values: [\"UNIMPLEMENTED\", \"PRELAUNCH\", \"EARLY_ACCESS\", \"ALPHA\", \"BETA\", \"GA\", \"DEPRECATED\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "The location of the cloud run service",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the Service.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "observed_generation": {
        "computed": true,
        "description": "The generation of this Service currently serving traffic. See comments in reconciling for additional information on reconciliation process in Cloud Run. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a string instead of an integer.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "Returns true if the Service is currently being acted upon by the system to bring it into the desired state.\n\nWhen a new Service is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Service to the desired serving state. This process is called reconciliation. While reconciliation is in process, observedGeneration, latest_ready_revison, trafficStatuses, and uri will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the serving state matches the Service, or there was an error, and reconciliation failed. This state can be found in terminalCondition.state.\n\nIf reconciliation succeeded, the following fields will match: traffic and trafficStatuses, observedGeneration and generation, latestReadyRevision and latestCreatedRevision.\n\nIf reconciliation failed, trafficStatuses, observedGeneration, and latestReadyRevision will have the state of the last serving revision, or empty for newly created Services. Additional information on the failure can be found in terminalCondition and conditions.",
        "description_kind": "plain",
        "type": "bool"
      },
      "scaling": {
        "computed": true,
        "description": "Scaling settings that apply to the whole service",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "min_instance_count": "number"
            }
          ]
        ]
      },
      "template": {
        "computed": true,
        "description": "The template used to create revisions for this Service.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "annotations": [
                "map",
                "string"
              ],
              "containers": [
                "list",
                [
                  "object",
                  {
                    "args": [
                      "list",
                      "string"
                    ],
                    "command": [
                      "list",
                      "string"
                    ],
                    "depends_on": [
                      "list",
                      "string"
                    ],
                    "env": [
                      "set",
                      [
                        "object",
                        {
                          "name": "string",
                          "value": "string",
                          "value_source": [
                            "list",
                            [
                              "object",
                              {
                                "secret_key_ref": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "secret": "string",
                                      "version": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "image": "string",
                    "liveness_probe": [
                      "list",
                      [
                        "object",
                        {
                          "failure_threshold": "number",
                          "grpc": [
                            "list",
                            [
                              "object",
                              {
                                "port": "number",
                                "service": "string"
                              }
                            ]
                          ],
                          "http_get": [
                            "list",
                            [
                              "object",
                              {
                                "http_headers": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value": "string"
                                    }
                                  ]
                                ],
                                "path": "string",
                                "port": "number"
                              }
                            ]
                          ],
                          "initial_delay_seconds": "number",
                          "period_seconds": "number",
                          "tcp_socket": [
                            "list",
                            [
                              "object",
                              {
                                "port": "number"
                              }
                            ]
                          ],
                          "timeout_seconds": "number"
                        }
                      ]
                    ],
                    "name": "string",
                    "ports": [
                      "list",
                      [
                        "object",
                        {
                          "container_port": "number",
                          "name": "string"
                        }
                      ]
                    ],
                    "resources": [
                      "list",
                      [
                        "object",
                        {
                          "cpu_idle": "bool",
                          "limits": [
                            "map",
                            "string"
                          ],
                          "startup_cpu_boost": "bool"
                        }
                      ]
                    ],
                    "startup_probe": [
                      "list",
                      [
                        "object",
                        {
                          "failure_threshold": "number",
                          "grpc": [
                            "list",
                            [
                              "object",
                              {
                                "port": "number",
                                "service": "string"
                              }
                            ]
                          ],
                          "http_get": [
                            "list",
                            [
                              "object",
                              {
                                "http_headers": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value": "string"
                                    }
                                  ]
                                ],
                                "path": "string",
                                "port": "number"
                              }
                            ]
                          ],
                          "initial_delay_seconds": "number",
                          "period_seconds": "number",
                          "tcp_socket": [
                            "list",
                            [
                              "object",
                              {
                                "port": "number"
                              }
                            ]
                          ],
                          "timeout_seconds": "number"
                        }
                      ]
                    ],
                    "volume_mounts": [
                      "list",
                      [
                        "object",
                        {
                          "mount_path": "string",
                          "name": "string"
                        }
                      ]
                    ],
                    "working_dir": "string"
                  }
                ]
              ],
              "encryption_key": "string",
              "execution_environment": "string",
              "labels": [
                "map",
                "string"
              ],
              "max_instance_request_concurrency": "number",
              "revision": "string",
              "scaling": [
                "list",
                [
                  "object",
                  {
                    "max_instance_count": "number",
                    "min_instance_count": "number"
                  }
                ]
              ],
              "service_account": "string",
              "session_affinity": "bool",
              "timeout": "string",
              "volumes": [
                "list",
                [
                  "object",
                  {
                    "cloud_sql_instance": [
                      "list",
                      [
                        "object",
                        {
                          "instances": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "empty_dir": [
                      "list",
                      [
                        "object",
                        {
                          "medium": "string",
                          "size_limit": "string"
                        }
                      ]
                    ],
                    "gcs": [
                      "list",
                      [
                        "object",
                        {
                          "bucket": "string",
                          "read_only": "bool"
                        }
                      ]
                    ],
                    "name": "string",
                    "nfs": [
                      "list",
                      [
                        "object",
                        {
                          "path": "string",
                          "read_only": "bool",
                          "server": "string"
                        }
                      ]
                    ],
                    "secret": [
                      "list",
                      [
                        "object",
                        {
                          "default_mode": "number",
                          "items": [
                            "list",
                            [
                              "object",
                              {
                                "mode": "number",
                                "path": "string",
                                "version": "string"
                              }
                            ]
                          ],
                          "secret": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "vpc_access": [
                "list",
                [
                  "object",
                  {
                    "connector": "string",
                    "egress": "string",
                    "network_interfaces": [
                      "list",
                      [
                        "object",
                        {
                          "network": "string",
                          "subnetwork": "string",
                          "tags": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "terminal_condition": {
        "computed": true,
        "description": "The Condition of this Service, containing its readiness status, and detailed error information in case it did not reach a serving state. See comments in reconciling for additional information on reconciliation process in Cloud Run.",
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
      "traffic": {
        "computed": true,
        "description": "Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100% traffic to the latest Ready Revision.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "percent": "number",
              "revision": "string",
              "tag": "string",
              "type": "string"
            }
          ]
        ]
      },
      "traffic_statuses": {
        "computed": true,
        "description": "Detailed status information for corresponding traffic targets. See comments in reconciling for additional information on reconciliation process in Cloud Run.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "percent": "number",
              "revision": "string",
              "tag": "string",
              "type": "string",
              "uri": "string"
            }
          ]
        ]
      },
      "uid": {
        "computed": true,
        "description": "Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The last-modified time.",
        "description_kind": "plain",
        "type": "string"
      },
      "uri": {
        "computed": true,
        "description": "The main URI in which this Service is serving traffic.",
        "description_kind": "plain",
        "type": "string"
      },
      "urls": {
        "computed": true,
        "description": "All URLs serving traffic for this Service.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCloudRunV2ServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudRunV2Service), &result)
	return &result
}
