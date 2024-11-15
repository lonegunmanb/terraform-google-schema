package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVertexAiEndpoint = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. Timestamp when this Endpoint was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "dedicated_endpoint_dns": {
        "computed": true,
        "description": "Output only. DNS of the dedicated endpoint. Will only be populated if dedicatedEndpointEnabled is true. Format: 'https://{endpointId}.{region}-{projectNumber}.prediction.vertexai.goog'.",
        "description_kind": "plain",
        "type": "string"
      },
      "dedicated_endpoint_enabled": {
        "description": "If true, the endpoint will be exposed through a dedicated DNS [Endpoint.dedicated_endpoint_dns]. Your request to the dedicated DNS will be isolated from other users' traffic and will have better performance and reliability. Note: Once you enabled dedicated endpoint, you won't be able to send request to the shared DNS {region}-aiplatform.googleapis.com. The limitation will be removed soon.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "deployed_models": {
        "computed": true,
        "description": "Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/).",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automatic_resources": [
                "list",
                [
                  "object",
                  {
                    "max_replica_count": "number",
                    "min_replica_count": "number"
                  }
                ]
              ],
              "create_time": "string",
              "dedicated_resources": [
                "list",
                [
                  "object",
                  {
                    "autoscaling_metric_specs": [
                      "list",
                      [
                        "object",
                        {
                          "metric_name": "string",
                          "target": "number"
                        }
                      ]
                    ],
                    "machine_spec": [
                      "list",
                      [
                        "object",
                        {
                          "accelerator_count": "number",
                          "accelerator_type": "string",
                          "machine_type": "string"
                        }
                      ]
                    ],
                    "max_replica_count": "number",
                    "min_replica_count": "number"
                  }
                ]
              ],
              "display_name": "string",
              "enable_access_logging": "bool",
              "enable_container_logging": "bool",
              "id": "string",
              "model": "string",
              "model_version_id": "string",
              "private_endpoints": [
                "list",
                [
                  "object",
                  {
                    "explain_http_uri": "string",
                    "health_http_uri": "string",
                    "predict_http_uri": "string",
                    "service_attachment": "string"
                  }
                ]
              ],
              "service_account": "string",
              "shared_resources": "string"
            }
          ]
        ]
      },
      "description": {
        "description": "The description of the Endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.",
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
      "etag": {
        "computed": true,
        "description": "Used to perform consistent read-modify-write updates. If not set, a blind \"overwrite\" update happens.",
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
        "description": "The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
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
      "model_deployment_monitoring_job": {
        "computed": true,
        "description": "Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: 'projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}'",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): 'projects/{project}/global/networks/{network}'. Where '{project}' is a project number, as in '12345', and '{network}' is network name. Only one of the fields, 'network' or 'privateServiceConnectConfig', can be set.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The region for the resource",
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
      "traffic_split": {
        "computed": true,
        "description": "A map from a DeployedModel's id to the percentage of this Endpoint's traffic that should be forwarded to that DeployedModel.\nIf a DeployedModel's id is not listed in this map, then it receives no traffic.\nThe traffic percentage values must add up to 100, or map must be empty if the Endpoint is to not accept any traffic at a moment. See\nthe 'deployModel' [example](https://cloud.google.com/vertex-ai/docs/general/deployment#deploy_a_model_to_an_endpoint) and\n[documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.endpoints/deployModel) for more information.\n\n~\u003e **Note:** To set the map to empty, set '\"{}\"', apply, and then remove the field from your config.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Timestamp when this Endpoint was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "encryption_spec": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource. Has the form: 'projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key'. The key needs to be in the same region as where the compute resource is created.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "predict_request_response_logging_config": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "If logging is enabled or not.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "sampling_rate": {
              "description": "Percentage of requests to be logged, expressed as a fraction in range(0,1]",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "bigquery_destination": {
              "block": {
                "attributes": {
                  "output_uri": {
                    "description": "BigQuery URI to a project or table, up to 2000 characters long. When only the project is specified, the Dataset and Table is created. When the full table reference is specified, the Dataset must exist and table must not exist. Accepted forms: - BigQuery path. For example: 'bq://projectId' or 'bq://projectId.bqDatasetId' or 'bq://projectId.bqDatasetId.bqTableId'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "BigQuery table for logging. If only given a project, a new dataset will be created with name 'logging_\u003cendpoint-display-name\u003e_\u003cendpoint-id\u003e' where will be made BigQuery-dataset-name compatible (e.g. most special characters will become underscores). If no table name is given, a new table will be created with name 'request_response_logging'",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configures the request-response logging for online prediction.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "private_service_connect_config": {
        "block": {
          "attributes": {
            "enable_private_service_connect": {
              "description": "Required. If true, expose the IndexEndpoint via private service connect.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "enable_secure_private_service_connect": {
              "description": "If set to true, enable secure private service connect with IAM authorization. Otherwise, private service connect will be done without authorization. Note latency will be slightly increased if authorization is enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "project_allowlist": {
              "description": "A list of Projects from which the forwarding rule will target the service attachment.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Configuration for private service connect. 'network' and 'privateServiceConnectConfig' are mutually exclusive.",
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

func GoogleVertexAiEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVertexAiEndpoint), &result)
	return &result
}
