package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleClouddeployDeliveryPipeline = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field ` + "`" + `effective_annotations` + "`" + ` for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "condition": {
        "computed": true,
        "description": "Output only. Information around the state of the Delivery Pipeline.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "pipeline_ready_condition": [
                "list",
                [
                  "object",
                  {
                    "status": "bool",
                    "update_time": "string"
                  }
                ]
              ],
              "targets_present_condition": [
                "list",
                [
                  "object",
                  {
                    "missing_targets": [
                      "list",
                      "string"
                    ],
                    "status": "bool",
                    "update_time": "string"
                  }
                ]
              ],
              "targets_type_condition": [
                "list",
                [
                  "object",
                  {
                    "error_details": "string",
                    "status": "bool"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Time at which the pipeline was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the ` + "`" + `DeliveryPipeline` + "`" + `. Max length is 255 characters.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.",
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
        "description": "Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be \u003c= 128 bytes.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field ` + "`" + `effective_labels` + "`" + ` for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the ` + "`" + `DeliveryPipeline` + "`" + `. Format is ` + "`" + `[a-z]([a-z0-9-]{0,61}[a-z0-9])?` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "suspended": {
        "description": "When suspended, no new releases or rollouts can be created, but in-progress ones will complete.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "uid": {
        "computed": true,
        "description": "Output only. Unique identifier of the ` + "`" + `DeliveryPipeline` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Most recent time at which the pipeline was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "serial_pipeline": {
        "block": {
          "block_types": {
            "stages": {
              "block": {
                "attributes": {
                  "profiles": {
                    "description": "Skaffold profiles to use when rendering the manifest for this stage's ` + "`" + `Target` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "target_id": {
                    "description": "The target_id to which this stage points. This field refers exclusively to the last segment of a target name. For example, this field would just be ` + "`" + `my-target` + "`" + ` (rather than ` + "`" + `projects/project/locations/location/targets/my-target` + "`" + `). The location of the ` + "`" + `Target` + "`" + ` is inferred to be the same as the location of the ` + "`" + `DeliveryPipeline` + "`" + ` that contains this ` + "`" + `Stage` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "deploy_parameters": {
                    "block": {
                      "attributes": {
                        "match_target_labels": {
                          "description": "Optional. Deploy parameters are applied to targets with match labels. If unspecified, deploy parameters are applied to all targets (including child targets of a multi-target).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "values": {
                          "description": "Required. Values are deploy parameters in key-value pairs.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description": "Optional. The deploy parameters to use for the target in this stage.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "strategy": {
                    "block": {
                      "block_types": {
                        "canary": {
                          "block": {
                            "block_types": {
                              "canary_deployment": {
                                "block": {
                                  "attributes": {
                                    "percentages": {
                                      "description": "Required. The percentage based deployments that will occur as a part of a ` + "`" + `Rollout` + "`" + `. List is expected in ascending order and each integer n is 0 \u003c= n \u003c 100.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "number"
                                      ]
                                    },
                                    "verify": {
                                      "description": "Whether to run verify tests after each percentage deployment.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "block_types": {
                                    "postdeploy": {
                                      "block": {
                                        "attributes": {
                                          "actions": {
                                            "description": "Optional. A sequence of skaffold custom actions to invoke during execution of the postdeploy job.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "Optional. Configuration for the postdeploy job of the last phase. If this is not configured, postdeploy job will not be present.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "predeploy": {
                                      "block": {
                                        "attributes": {
                                          "actions": {
                                            "description": "Optional. A sequence of skaffold custom actions to invoke during execution of the predeploy job.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "Optional. Configuration for the predeploy job of the first phase. If this is not configured, predeploy job will not be present.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Configures the progressive based deployment for a Target.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "custom_canary_deployment": {
                                "block": {
                                  "block_types": {
                                    "phase_configs": {
                                      "block": {
                                        "attributes": {
                                          "percentage": {
                                            "description": "Required. Percentage deployment for the phase.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "phase_id": {
                                            "description": "Required. The ID to assign to the ` + "`" + `Rollout` + "`" + ` phase. This value must consist of lower-case letters, numbers, and hyphens, start with a letter and end with a letter or a number, and have a max length of 63 characters. In other words, it must match the following regex: ` + "`" + `^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$` + "`" + `.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "profiles": {
                                            "description": "Skaffold profiles to use when rendering the manifest for this phase. These are in addition to the profiles list specified in the ` + "`" + `DeliveryPipeline` + "`" + ` stage.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "verify": {
                                            "description": "Whether to run verify tests after the deployment.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          }
                                        },
                                        "block_types": {
                                          "postdeploy": {
                                            "block": {
                                              "attributes": {
                                                "actions": {
                                                  "description": "Optional. A sequence of skaffold custom actions to invoke during execution of the postdeploy job.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "description": "Optional. Configuration for the postdeploy job of this phase. If this is not configured, postdeploy job will not be present for this phase.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "predeploy": {
                                            "block": {
                                              "attributes": {
                                                "actions": {
                                                  "description": "Optional. A sequence of skaffold custom actions to invoke during execution of the predeploy job.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "description": "Optional. Configuration for the predeploy job of this phase. If this is not configured, predeploy job will not be present for this phase.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Required. Configuration for each phase in the canary deployment in the order executed.",
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Configures the progressive based deployment for a Target, but allows customizing at the phase level where a phase represents each of the percentage deployments.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "runtime_config": {
                                "block": {
                                  "block_types": {
                                    "cloud_run": {
                                      "block": {
                                        "attributes": {
                                          "automatic_traffic_control": {
                                            "description": "Whether Cloud Deploy should update the traffic stanza in a Cloud Run Service on the user's behalf to facilitate traffic splitting. This is required to be true for CanaryDeployments, but optional for CustomCanaryDeployments.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "canary_revision_tags": {
                                            "description": "Optional. A list of tags that are added to the canary revision while the canary phase is in progress.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "prior_revision_tags": {
                                            "description": "Optional. A list of tags that are added to the prior revision while the canary phase is in progress.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "stable_revision_tags": {
                                            "description": "Optional. A list of tags that are added to the final stable revision when the stable phase is applied.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "Cloud Run runtime configuration.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "kubernetes": {
                                      "block": {
                                        "block_types": {
                                          "gateway_service_mesh": {
                                            "block": {
                                              "attributes": {
                                                "deployment": {
                                                  "description": "Required. Name of the Kubernetes Deployment whose traffic is managed by the specified HTTPRoute and Service.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "http_route": {
                                                  "description": "Required. Name of the Gateway API HTTPRoute.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "pod_selector_label": {
                                                  "description": "Optional. The label to use when selecting Pods for the Deployment and Service resources. This label must already be present in both resources.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "route_update_wait_time": {
                                                  "description": "Optional. The time to wait for route updates to propagate. The maximum configurable time is 3 hours, in seconds format. If unspecified, there is no wait time.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "service": {
                                                  "description": "Required. Name of the Kubernetes Service.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "stable_cutback_duration": {
                                                  "description": "Optional. The amount of time to migrate traffic back from the canary Service to the original Service during the stable phase deployment. If specified, must be between 15s and 3600s. If unspecified, there is no cutback time.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "route_destinations": {
                                                  "block": {
                                                    "attributes": {
                                                      "destination_ids": {
                                                        "description": "Required. The clusters where the Gateway API HTTPRoute resource will be deployed to. Valid entries include the associated entities IDs configured in the Target resource and \"@self\" to include the Target cluster.",
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": [
                                                          "list",
                                                          "string"
                                                        ]
                                                      },
                                                      "propagate_service": {
                                                        "description": "Optional. Whether to propagate the Kubernetes Service to the route destination clusters. The Service will always be deployed to the Target cluster even if the HTTPRoute is not. This option may be used to facilitiate successful DNS lookup in the route destination clusters. Can only be set to true if destinations are specified.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "bool"
                                                      }
                                                    },
                                                    "description": "Optional. Route destinations allow configuring the Gateway API HTTPRoute to be deployed to additional clusters. This option is available for multi-cluster service mesh set ups that require the route to exist in the clusters that call the service. If unspecified, the HTTPRoute will only be deployed to the Target cluster.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                }
                                              },
                                              "description": "Kubernetes Gateway API service mesh configuration.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "service_networking": {
                                            "block": {
                                              "attributes": {
                                                "deployment": {
                                                  "description": "Required. Name of the Kubernetes Deployment whose traffic is managed by the specified Service.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "disable_pod_overprovisioning": {
                                                  "description": "Optional. Whether to disable Pod overprovisioning. If Pod overprovisioning is disabled then Cloud Deploy will limit the number of total Pods used for the deployment strategy to the number of Pods the Deployment has on the cluster.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
                                                },
                                                "pod_selector_label": {
                                                  "description": "Optional. The label to use when selecting Pods for the Deployment resource. This label must already be present in the Deployment.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "service": {
                                                  "description": "Required. Name of the Kubernetes Service.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Kubernetes Service networking configuration.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Kubernetes runtime configuration.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Optional. Runtime specific configurations for the deployment strategy. The runtime configuration is used to determine how Cloud Deploy will split traffic to enable a progressive deployment.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Canary deployment strategy provides progressive percentage based deployments to a Target.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "standard": {
                          "block": {
                            "attributes": {
                              "verify": {
                                "description": "Whether to verify a deployment.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "postdeploy": {
                                "block": {
                                  "attributes": {
                                    "actions": {
                                      "description": "Optional. A sequence of skaffold custom actions to invoke during execution of the postdeploy job.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "Optional. Configuration for the postdeploy job. If this is not configured, postdeploy job will not be present.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "predeploy": {
                                "block": {
                                  "attributes": {
                                    "actions": {
                                      "description": "Optional. A sequence of skaffold custom actions to invoke during execution of the predeploy job.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "Optional. Configuration for the predeploy job. If this is not configured, predeploy job will not be present.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Standard deployment strategy executes a single deploy and allows verifying the deployment.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Optional. The strategy to use for a ` + "`" + `Rollout` + "`" + ` to this stage.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Each stage specifies configuration for a ` + "`" + `Target` + "`" + `. The ordering of this list defines the promotion flow.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "SerialPipeline defines a sequential set of stages for a ` + "`" + `DeliveryPipeline` + "`" + `.",
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

func GoogleClouddeployDeliveryPipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleClouddeployDeliveryPipeline), &result)
	return &result
}
