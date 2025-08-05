package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDeveloperConnectInsightsConfig = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "User specified annotations. See https://google.aip.dev/148#annotations\nfor more details such as format and size limitations.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "app_hub_application": {
        "description": "The name of the App Hub Application.\nFormat:\nprojects/{project}/locations/{location}/applications/{application}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "[Output only] Create timestamp",
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
      "errors": {
        "computed": true,
        "description": "Any errors that occurred while setting up the InsightsConfig.\nEach error will be in the format: 'field_name: error_message', e.g.\nGetAppHubApplication: Permission denied while getting App Hub\napplication. Please grant permissions to the P4SA.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "code": "number",
              "details": [
                "list",
                [
                  "object",
                  {
                    "detail_message": "string"
                  }
                ]
              ],
              "message": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "insights_config_id": {
        "description": "ID of the requesting InsightsConfig.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "labels": {
        "description": "Set of labels associated with an InsightsConfig.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The name of the InsightsConfig.\nFormat:\nprojects/{project}/locations/{location}/insightsConfigs/{insightsConfig}",
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
        "description": "Reconciling (https://google.aip.dev/128#reconciliation).\nSet to true if the current state of InsightsConfig does not match the\nuser's intended state, and the service is actively updating the resource to\nreconcile them. This can happen due to user-triggered updates or\nsystem actions like failover or maintenance.",
        "description_kind": "plain",
        "type": "bool"
      },
      "runtime_configs": {
        "computed": true,
        "description": "The runtime configurations where the application is deployed.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "app_hub_workload": [
                "list",
                [
                  "object",
                  {
                    "criticality": "string",
                    "environment": "string",
                    "workload": "string"
                  }
                ]
              ],
              "gke_workload": [
                "list",
                [
                  "object",
                  {
                    "cluster": "string",
                    "deployment": "string"
                  }
                ]
              ],
              "state": "string",
              "uri": "string"
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description": "The state of the InsightsConfig.\nPossible values:\nSTATE_UNSPECIFIED\nPENDING\nCOMPLETE\nERROR",
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
        "description": "[Output only] Update timestamp",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "artifact_configs": {
        "block": {
          "attributes": {
            "uri": {
              "description": "The URI of the artifact that is deployed.\ne.g. 'us-docker.pkg.dev/my-project/my-repo/image'.\nThe URI does not include the tag / digest because it captures a lineage of\nartifacts.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "google_artifact_analysis": {
              "block": {
                "attributes": {
                  "project_id": {
                    "description": "The project id of the project where the provenance is stored.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Google Artifact Analysis configurations.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "google_artifact_registry": {
              "block": {
                "attributes": {
                  "artifact_registry_package": {
                    "description": "The name of the artifact registry package.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "project_id": {
                    "description": "The host project of Artifact Registry.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Google Artifact Registry configurations.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The artifact configurations of the artifacts that are deployed.",
          "description_kind": "plain"
        },
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

func GoogleDeveloperConnectInsightsConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDeveloperConnectInsightsConfig), &result)
	return &result
}
