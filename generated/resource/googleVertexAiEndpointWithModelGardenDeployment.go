package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVertexAiEndpointWithModelGardenDeployment = `{
  "block": {
    "attributes": {
      "endpoint": {
        "computed": true,
        "description": "Resource ID segment making up resource 'endpoint'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
        "description_kind": "plain",
        "type": "string"
      },
      "hugging_face_model_id": {
        "description": "The Hugging Face model to deploy.\nFormat: Hugging Face model ID like 'google/gemma-2-2b-it'.",
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
        "description": "Resource ID segment making up resource 'location'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
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
      "publisher_model_name": {
        "description": "The Model Garden model to deploy.\nFormat:\n'publishers/{publisher}/models/{publisher_model}@{version_id}', or\n'publishers/hf-{hugging-face-author}/models/{hugging-face-model-name}@001'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "deploy_config": {
        "block": {
          "attributes": {
            "fast_tryout_enabled": {
              "description": "If true, enable the QMT fast tryout feature for this model if possible.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "system_labels": {
              "description": "System labels for Model Garden deployments.\nThese labels are managed by Google and for tracking purposes only.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "block_types": {
            "dedicated_resources": {
              "block": {
                "attributes": {
                  "max_replica_count": {
                    "description": "The maximum number of replicas that may be deployed on when the traffic\nagainst it increases. If the requested value is too large, the deployment\nwill error, but if deployment succeeds then the ability to scale to that\nmany replicas is guaranteed (barring service outages). If traffic increases\nbeyond what its replicas at maximum may handle, a portion of the traffic\nwill be dropped. If this value is not provided, will use\nmin_replica_count as the default value.\n\nThe value of this field impacts the charge against Vertex CPU and GPU\nquotas. Specifically, you will be charged for (max_replica_count *\nnumber of cores in the selected machine type) and (max_replica_count *\nnumber of GPUs per replica in the selected machine type).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_replica_count": {
                    "description": "The minimum number of machine replicas that will be always deployed on.\nThis value must be greater than or equal to 1.\n\nIf traffic increases, it may dynamically be deployed onto more replicas,\nand as traffic decreases, some of these extra replicas may be freed.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "required_replica_count": {
                    "description": "Number of required available replicas for the deployment to succeed.\nThis field is only needed when partial deployment/mutation is\ndesired. If set, the deploy/mutate operation will succeed once\navailable_replica_count reaches required_replica_count, and the rest of\nthe replicas will be retried. If not set, the default\nrequired_replica_count will be min_replica_count.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "spot": {
                    "description": "If true, schedule the deployment workload on [spot\nVMs](https://cloud.google.com/kubernetes-engine/docs/concepts/spot-vms).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "autoscaling_metric_specs": {
                    "block": {
                      "attributes": {
                        "metric_name": {
                          "description": "The resource metric name.\nSupported metrics:\n\n* For Online Prediction:\n* 'aiplatform.googleapis.com/prediction/online/accelerator/duty_cycle'\n* 'aiplatform.googleapis.com/prediction/online/cpu/utilization'",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "target": {
                          "description": "The target resource utilization in percentage (1% - 100%) for the given\nmetric; once the real usage deviates from the target by a certain\npercentage, the machine replicas change. The default value is 60\n(representing 60%) if not provided.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "The metric specifications that overrides a resource\nutilization metric (CPU utilization, accelerator's duty cycle, and so on)\ntarget value (default to 60 if not set). At most one entry is allowed per\nmetric.\n\nIf machine_spec.accelerator_count is\nabove 0, the autoscaling will be based on both CPU utilization and\naccelerator's duty cycle metrics and scale up when either metrics exceeds\nits target value while scale down if both metrics are under their target\nvalue. The default target value is 60 for both metrics.\n\nIf machine_spec.accelerator_count is\n0, the autoscaling will be based on CPU utilization metric only with\ndefault target value 60 if not explicitly set.\n\nFor example, in the case of Online Prediction, if you want to override\ntarget CPU utilization to 80, you should set\nautoscaling_metric_specs.metric_name\nto 'aiplatform.googleapis.com/prediction/online/cpu/utilization' and\nautoscaling_metric_specs.target to '80'.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "machine_spec": {
                    "block": {
                      "attributes": {
                        "accelerator_count": {
                          "description": "The number of accelerators to attach to the machine.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "accelerator_type": {
                          "description": "Possible values:\nACCELERATOR_TYPE_UNSPECIFIED\nNVIDIA_TESLA_K80\nNVIDIA_TESLA_P100\nNVIDIA_TESLA_V100\nNVIDIA_TESLA_P4\nNVIDIA_TESLA_T4\nNVIDIA_TESLA_A100\nNVIDIA_A100_80GB\nNVIDIA_L4\nNVIDIA_H100_80GB\nNVIDIA_H100_MEGA_80GB\nNVIDIA_H200_141GB\nNVIDIA_B200\nTPU_V2\nTPU_V3\nTPU_V4_POD\nTPU_V5_LITEPOD",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "machine_type": {
                          "description": "The type of the machine.\n\nSee the [list of machine types supported for\nprediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types)\n\nSee the [list of machine types supported for custom\ntraining](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types).\n\nFor DeployedModel this field is optional, and the default\nvalue is 'n1-standard-2'. For BatchPredictionJob or as part of\nWorkerPoolSpec this field is required.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "multihost_gpu_node_count": {
                          "description": "The number of nodes per replica for multihost GPU deployments.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "tpu_topology": {
                          "description": "The topology of the TPUs. Corresponds to the TPU topologies available from\nGKE. (Example: tpu_topology: \"2x2x1\").",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "reservation_affinity": {
                          "block": {
                            "attributes": {
                              "key": {
                                "description": "Corresponds to the label key of a reservation resource. To target a\nSPECIFIC_RESERVATION by name, use 'compute.googleapis.com/reservation-name'\nas the key and specify the name of your reservation as its value.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "reservation_affinity_type": {
                                "description": "Specifies the reservation affinity type.\nPossible values:\nTYPE_UNSPECIFIED\nNO_RESERVATION\nANY_RESERVATION\nSPECIFIC_RESERVATION",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "values": {
                                "description": "Corresponds to the label values of a reservation resource. This must be the\nfull resource name of the reservation or reservation block.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "A ReservationAffinity can be used to configure a Vertex AI resource (e.g., a\nDeployedModel) to draw its Compute Engine resources from a Shared\nReservation, or exclusively from on-demand capacity.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Specification of a single machine.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A description of resources that are dedicated to a DeployedModel or\nDeployedIndex, and that need a higher degree of manual configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The deploy config to use for the deployment.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "endpoint_config": {
        "block": {
          "attributes": {
            "dedicated_endpoint_enabled": {
              "description": "If true, the endpoint will be exposed through a dedicated\nDNS [Endpoint.dedicated_endpoint_dns]. Your request to the dedicated DNS\nwill be isolated from other users' traffic and will have better\nperformance and reliability. Note: Once you enabled dedicated endpoint,\nyou won't be able to send request to the shared DNS\n{region}-aiplatform.googleapis.com. The limitations will be removed soon.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "endpoint_display_name": {
              "description": "The user-specified display name of the endpoint. If not set, a\ndefault name will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The endpoint config to use for the deployment.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "model_config": {
        "block": {
          "attributes": {
            "accept_eula": {
              "description": "Whether the user accepts the End User License Agreement (EULA)\nfor the model.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hugging_face_access_token": {
              "description": "The Hugging Face read access token used to access the model\nartifacts of gated models.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hugging_face_cache_enabled": {
              "description": "If true, the model will deploy with a cached version instead of directly\ndownloading the model artifacts from Hugging Face. This is suitable for\nVPC-SC users with limited internet access.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "model_display_name": {
              "description": "The user-specified display name of the uploaded model. If not\nset, a default name will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "container_spec": {
              "block": {
                "attributes": {
                  "args": {
                    "description": "Specifies arguments for the command that runs when the container starts.\nThis overrides the container's\n['CMD'](https://docs.docker.com/engine/reference/builder/#cmd). Specify\nthis field as an array of executable and arguments, similar to a Docker\n'CMD''s \"default parameters\" form.\n\nIf you don't specify this field but do specify the\ncommand field, then the command from the\n'command' field runs without any additional arguments. See the\n[Kubernetes documentation about how the\n'command' and 'args' fields interact with a container's 'ENTRYPOINT' and\n'CMD'](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes).\n\nIf you don't specify this field and don't specify the 'command' field,\nthen the container's\n['ENTRYPOINT'](https://docs.docker.com/engine/reference/builder/#cmd) and\n'CMD' determine what runs based on their default behavior. See the Docker\ndocumentation about [how 'CMD' and 'ENTRYPOINT'\ninteract](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact).\n\nIn this field, you can reference [environment variables\nset by Vertex\nAI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables)\nand environment variables set in the env field.\nYou cannot reference environment variables set in the Docker image. In\norder for environment variables to be expanded, reference them by using the\nfollowing syntax:$(VARIABLE_NAME)\nNote that this differs from Bash variable expansion, which does not use\nparentheses. If a variable cannot be resolved, the reference in the input\nstring is used unchanged. To avoid variable expansion, you can escape this\nsyntax with '$$'; for example:$$(VARIABLE_NAME)\nThis field corresponds to the 'args' field of the Kubernetes Containers\n[v1 core\nAPI](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "command": {
                    "description": "Specifies the command that runs when the container starts. This overrides\nthe container's\n[ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint).\nSpecify this field as an array of executable and arguments, similar to a\nDocker 'ENTRYPOINT''s \"exec\" form, not its \"shell\" form.\n\nIf you do not specify this field, then the container's 'ENTRYPOINT' runs,\nin conjunction with the args field or the\ncontainer's ['CMD'](https://docs.docker.com/engine/reference/builder/#cmd),\nif either exists. If this field is not specified and the container does not\nhave an 'ENTRYPOINT', then refer to the Docker documentation about [how\n'CMD' and 'ENTRYPOINT'\ninteract](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact).\n\nIf you specify this field, then you can also specify the 'args' field to\nprovide additional arguments for this command. However, if you specify this\nfield, then the container's 'CMD' is ignored. See the\n[Kubernetes documentation about how the\n'command' and 'args' fields interact with a container's 'ENTRYPOINT' and\n'CMD'](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes).\n\nIn this field, you can reference [environment variables set by Vertex\nAI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables)\nand environment variables set in the env field.\nYou cannot reference environment variables set in the Docker image. In\norder for environment variables to be expanded, reference them by using the\nfollowing syntax:$(VARIABLE_NAME)\nNote that this differs from Bash variable expansion, which does not use\nparentheses. If a variable cannot be resolved, the reference in the input\nstring is used unchanged. To avoid variable expansion, you can escape this\nsyntax with '$$'; for example:$$(VARIABLE_NAME)\nThis field corresponds to the 'command' field of the Kubernetes Containers\n[v1 core\nAPI](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "deployment_timeout": {
                    "description": "Deployment timeout.\nLimit for deployment timeout is 2 hours.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "health_route": {
                    "description": "HTTP path on the container to send health checks to. Vertex AI\nintermittently sends GET requests to this path on the container's IP\naddress and port to check that the container is healthy. Read more about\n[health\nchecks](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#health).\n\nFor example, if you set this field to '/bar', then Vertex AI\nintermittently sends a GET request to the '/bar' path on the port of your\ncontainer specified by the first value of this 'ModelContainerSpec''s\nports field.\n\nIf you don't specify this field, it defaults to the following value when\nyou deploy this Model to an Endpoint:/v1/endpoints/ENDPOINT/deployedModels/DEPLOYED_MODEL:predict\nThe placeholders in this value are replaced as follows:\n\n* ENDPOINT: The last segment (following 'endpoints/')of the\nEndpoint.name][] field of the Endpoint where this Model has been\ndeployed. (Vertex AI makes this value available to your container code\nas the ['AIP_ENDPOINT_ID' environment\nvariable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)\n\n* DEPLOYED_MODEL: DeployedModel.id of the 'DeployedModel'.\n(Vertex AI makes this value available to your container code as the\n['AIP_DEPLOYED_MODEL_ID' environment\nvariable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "image_uri": {
                    "description": "URI of the Docker image to be used as the custom container for serving\npredictions. This URI must identify an image in Artifact Registry or\nContainer Registry. Learn more about the [container publishing\nrequirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#publishing),\nincluding permissions requirements for the Vertex AI Service Agent.\n\nThe container image is ingested upon ModelService.UploadModel, stored\ninternally, and this original path is afterwards not used.\n\nTo learn about the requirements for the Docker image itself, see\n[Custom container\nrequirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#).\n\nYou can use the URI to one of Vertex AI's [pre-built container images for\nprediction](https://cloud.google.com/vertex-ai/docs/predictions/pre-built-containers)\nin this field.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "predict_route": {
                    "description": "HTTP path on the container to send prediction requests to. Vertex AI\nforwards requests sent using\nprojects.locations.endpoints.predict to this\npath on the container's IP address and port. Vertex AI then returns the\ncontainer's response in the API response.\n\nFor example, if you set this field to '/foo', then when Vertex AI\nreceives a prediction request, it forwards the request body in a POST\nrequest to the '/foo' path on the port of your container specified by the\nfirst value of this 'ModelContainerSpec''s\nports field.\n\nIf you don't specify this field, it defaults to the following value when\nyou deploy this Model to an Endpoint:/v1/endpoints/ENDPOINT/deployedModels/DEPLOYED_MODEL:predict\nThe placeholders in this value are replaced as follows:\n\n* ENDPOINT: The last segment (following 'endpoints/')of the\nEndpoint.name][] field of the Endpoint where this Model has been\ndeployed. (Vertex AI makes this value available to your container code\nas the ['AIP_ENDPOINT_ID' environment\nvariable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)\n\n* DEPLOYED_MODEL: DeployedModel.id of the 'DeployedModel'.\n(Vertex AI makes this value available to your container code\nas the ['AIP_DEPLOYED_MODEL_ID' environment\nvariable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "shared_memory_size_mb": {
                    "description": "The amount of the VM memory to reserve as the shared memory for the model\nin megabytes.",
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
                          "description": "Name of the environment variable. Must be a valid C identifier.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "Variables that reference a $(VAR_NAME) are expanded\nusing the previous defined environment variables in the container and\nany service environment variables. If a variable cannot be resolved,\nthe reference in the input string will be unchanged. The $(VAR_NAME)\nsyntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped\nreferences will never be expanded, regardless of whether the variable\nexists or not.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "List of environment variables to set in the container. After the container\nstarts running, code running in the container can read these environment\nvariables.\n\nAdditionally, the command and\nargs fields can reference these variables. Later\nentries in this list can also reference earlier entries. For example, the\nfollowing example sets the variable 'VAR_2' to have the value 'foo bar':\n\n'''json\n[\n{\n\"name\": \"VAR_1\",\n\"value\": \"foo\"\n},\n{\n\"name\": \"VAR_2\",\n\"value\": \"$(VAR_1) bar\"\n}\n]\n'''\n\nIf you switch the order of the variables in the example, then the expansion\ndoes not occur.\n\nThis field corresponds to the 'env' field of the Kubernetes Containers\n[v1 core\nAPI](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "grpc_ports": {
                    "block": {
                      "attributes": {
                        "container_port": {
                          "description": "The number of the port to expose on the pod's IP address.\nMust be a valid port number, between 1 and 65535 inclusive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "List of ports to expose from the container. Vertex AI sends gRPC\nprediction requests that it receives to the first port on this list. Vertex\nAI also sends liveness and health checks to this port.\n\nIf you do not specify this field, gRPC requests to the container will be\ndisabled.\n\nVertex AI does not use ports other than the first one listed. This field\ncorresponds to the 'ports' field of the Kubernetes Containers v1 core API.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "health_probe": {
                    "block": {
                      "attributes": {
                        "failure_threshold": {
                          "description": "Number of consecutive failures before the probe is considered failed.\nDefaults to 3. Minimum value is 1.\n\nMaps to Kubernetes probe argument 'failureThreshold'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "initial_delay_seconds": {
                          "description": "Number of seconds to wait before starting the probe. Defaults to 0.\nMinimum value is 0.\n\nMaps to Kubernetes probe argument 'initialDelaySeconds'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "period_seconds": {
                          "description": "How often (in seconds) to perform the probe. Default to 10 seconds.\nMinimum value is 1. Must be less than timeout_seconds.\n\nMaps to Kubernetes probe argument 'periodSeconds'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "success_threshold": {
                          "description": "Number of consecutive successes before the probe is considered successful.\nDefaults to 1. Minimum value is 1.\n\nMaps to Kubernetes probe argument 'successThreshold'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_seconds": {
                          "description": "Number of seconds after which the probe times out. Defaults to 1 second.\nMinimum value is 1. Must be greater or equal to period_seconds.\n\nMaps to Kubernetes probe argument 'timeoutSeconds'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "exec": {
                          "block": {
                            "attributes": {
                              "command": {
                                "description": "Command is the command line to execute inside the container, the working\ndirectory for the command is root ('/') in the container's filesystem.\nThe command is simply exec'd, it is not run inside a shell, so\ntraditional shell instructions ('|', etc) won't work. To use a shell, you\nneed to explicitly call out to that shell. Exit status of 0 is treated as\nlive/healthy and non-zero is unhealthy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "ExecAction specifies a command to execute.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "grpc": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Port number of the gRPC service. Number must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "service": {
                                "description": "Service is the name of the service to place in the gRPC\nHealthCheckRequest. See\nhttps://github.com/grpc/grpc/blob/master/doc/health-checking.md.\n\nIf this is not specified, the default behavior is defined by gRPC.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "GrpcAction checks the health of a container using a gRPC service.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "http_get": {
                          "block": {
                            "attributes": {
                              "host": {
                                "description": "Host name to connect to, defaults to the model serving container's IP.\nYou probably want to set \"Host\" in httpHeaders instead.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "description": "Path to access on the HTTP server.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "Number of the port to access on the container.\nNumber must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "scheme": {
                                "description": "Scheme to use for connecting to the host.\nDefaults to HTTP. Acceptable values are \"HTTP\" or \"HTTPS\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "http_headers": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "The header field name.\nThis will be canonicalized upon output, so case-variant names will be\nunderstood as the same header.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "The header field value",
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
                            "description": "HttpGetAction describes an action based on HTTP Get requests.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "tcp_socket": {
                          "block": {
                            "attributes": {
                              "host": {
                                "description": "Optional: Host name to connect to, defaults to the model serving\ncontainer's IP.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "Number of the port to access on the container.\nNumber must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "TcpSocketAction probes the health of a container by opening a TCP socket\nconnection.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Probe describes a health check to be performed against a container to\ndetermine whether it is alive or ready to receive traffic.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "liveness_probe": {
                    "block": {
                      "attributes": {
                        "failure_threshold": {
                          "description": "Number of consecutive failures before the probe is considered failed.\nDefaults to 3. Minimum value is 1.\n\nMaps to Kubernetes probe argument 'failureThreshold'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "initial_delay_seconds": {
                          "description": "Number of seconds to wait before starting the probe. Defaults to 0.\nMinimum value is 0.\n\nMaps to Kubernetes probe argument 'initialDelaySeconds'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "period_seconds": {
                          "description": "How often (in seconds) to perform the probe. Default to 10 seconds.\nMinimum value is 1. Must be less than timeout_seconds.\n\nMaps to Kubernetes probe argument 'periodSeconds'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "success_threshold": {
                          "description": "Number of consecutive successes before the probe is considered successful.\nDefaults to 1. Minimum value is 1.\n\nMaps to Kubernetes probe argument 'successThreshold'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_seconds": {
                          "description": "Number of seconds after which the probe times out. Defaults to 1 second.\nMinimum value is 1. Must be greater or equal to period_seconds.\n\nMaps to Kubernetes probe argument 'timeoutSeconds'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "exec": {
                          "block": {
                            "attributes": {
                              "command": {
                                "description": "Command is the command line to execute inside the container, the working\ndirectory for the command is root ('/') in the container's filesystem.\nThe command is simply exec'd, it is not run inside a shell, so\ntraditional shell instructions ('|', etc) won't work. To use a shell, you\nneed to explicitly call out to that shell. Exit status of 0 is treated as\nlive/healthy and non-zero is unhealthy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "ExecAction specifies a command to execute.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "grpc": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Port number of the gRPC service. Number must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "service": {
                                "description": "Service is the name of the service to place in the gRPC\nHealthCheckRequest. See\nhttps://github.com/grpc/grpc/blob/master/doc/health-checking.md.\n\nIf this is not specified, the default behavior is defined by gRPC.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "GrpcAction checks the health of a container using a gRPC service.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "http_get": {
                          "block": {
                            "attributes": {
                              "host": {
                                "description": "Host name to connect to, defaults to the model serving container's IP.\nYou probably want to set \"Host\" in httpHeaders instead.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "description": "Path to access on the HTTP server.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "Number of the port to access on the container.\nNumber must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "scheme": {
                                "description": "Scheme to use for connecting to the host.\nDefaults to HTTP. Acceptable values are \"HTTP\" or \"HTTPS\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "http_headers": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "The header field name.\nThis will be canonicalized upon output, so case-variant names will be\nunderstood as the same header.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "The header field value",
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
                            "description": "HttpGetAction describes an action based on HTTP Get requests.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "tcp_socket": {
                          "block": {
                            "attributes": {
                              "host": {
                                "description": "Optional: Host name to connect to, defaults to the model serving\ncontainer's IP.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "Number of the port to access on the container.\nNumber must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "TcpSocketAction probes the health of a container by opening a TCP socket\nconnection.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Probe describes a health check to be performed against a container to\ndetermine whether it is alive or ready to receive traffic.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "ports": {
                    "block": {
                      "attributes": {
                        "container_port": {
                          "description": "The number of the port to expose on the pod's IP address.\nMust be a valid port number, between 1 and 65535 inclusive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "List of ports to expose from the container. Vertex AI sends any\nprediction requests that it receives to the first port on this list. Vertex\nAI also sends\n[liveness and health\nchecks](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#liveness)\nto this port.\n\nIf you do not specify this field, it defaults to following value:\n\n'''json\n[\n{\n\"containerPort\": 8080\n}\n]\n'''\n\nVertex AI does not use ports other than the first one listed. This field\ncorresponds to the 'ports' field of the Kubernetes Containers\n[v1 core\nAPI](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "startup_probe": {
                    "block": {
                      "attributes": {
                        "failure_threshold": {
                          "description": "Number of consecutive failures before the probe is considered failed.\nDefaults to 3. Minimum value is 1.\n\nMaps to Kubernetes probe argument 'failureThreshold'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "initial_delay_seconds": {
                          "description": "Number of seconds to wait before starting the probe. Defaults to 0.\nMinimum value is 0.\n\nMaps to Kubernetes probe argument 'initialDelaySeconds'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "period_seconds": {
                          "description": "How often (in seconds) to perform the probe. Default to 10 seconds.\nMinimum value is 1. Must be less than timeout_seconds.\n\nMaps to Kubernetes probe argument 'periodSeconds'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "success_threshold": {
                          "description": "Number of consecutive successes before the probe is considered successful.\nDefaults to 1. Minimum value is 1.\n\nMaps to Kubernetes probe argument 'successThreshold'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_seconds": {
                          "description": "Number of seconds after which the probe times out. Defaults to 1 second.\nMinimum value is 1. Must be greater or equal to period_seconds.\n\nMaps to Kubernetes probe argument 'timeoutSeconds'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "exec": {
                          "block": {
                            "attributes": {
                              "command": {
                                "description": "Command is the command line to execute inside the container, the working\ndirectory for the command is root ('/') in the container's filesystem.\nThe command is simply exec'd, it is not run inside a shell, so\ntraditional shell instructions ('|', etc) won't work. To use a shell, you\nneed to explicitly call out to that shell. Exit status of 0 is treated as\nlive/healthy and non-zero is unhealthy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "description": "ExecAction specifies a command to execute.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "grpc": {
                          "block": {
                            "attributes": {
                              "port": {
                                "description": "Port number of the gRPC service. Number must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "service": {
                                "description": "Service is the name of the service to place in the gRPC\nHealthCheckRequest. See\nhttps://github.com/grpc/grpc/blob/master/doc/health-checking.md.\n\nIf this is not specified, the default behavior is defined by gRPC.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "GrpcAction checks the health of a container using a gRPC service.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "http_get": {
                          "block": {
                            "attributes": {
                              "host": {
                                "description": "Host name to connect to, defaults to the model serving container's IP.\nYou probably want to set \"Host\" in httpHeaders instead.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "description": "Path to access on the HTTP server.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "Number of the port to access on the container.\nNumber must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "scheme": {
                                "description": "Scheme to use for connecting to the host.\nDefaults to HTTP. Acceptable values are \"HTTP\" or \"HTTPS\".",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "http_headers": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "The header field name.\nThis will be canonicalized upon output, so case-variant names will be\nunderstood as the same header.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "The header field value",
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
                            "description": "HttpGetAction describes an action based on HTTP Get requests.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "tcp_socket": {
                          "block": {
                            "attributes": {
                              "host": {
                                "description": "Optional: Host name to connect to, defaults to the model serving\ncontainer's IP.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "Number of the port to access on the container.\nNumber must be in the range 1 to 65535.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "TcpSocketAction probes the health of a container by opening a TCP socket\nconnection.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Probe describes a health check to be performed against a container to\ndetermine whether it is alive or ready to receive traffic.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specification of a container for serving predictions. Some fields in this\nmessage correspond to fields in the [Kubernetes Container v1 core\nspecification](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The model config to use for the deployment.",
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

func GoogleVertexAiEndpointWithModelGardenDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVertexAiEndpointWithModelGardenDeployment), &result)
	return &result
}
