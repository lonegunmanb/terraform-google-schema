package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkSecurityMirroringEndpointGroup = `{
  "block": {
    "attributes": {
      "associations": {
        "computed": true,
        "description": "List of associations to this endpoint group.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "name": "string",
              "network": "string",
              "state": "string"
            }
          ]
        ]
      },
      "connected_deployment_groups": {
        "computed": true,
        "description": "List of details about the connected deployment groups to this endpoint\ngroup.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "locations": [
                "set",
                [
                  "object",
                  {
                    "location": "string",
                    "state": "string"
                  }
                ]
              ],
              "name": "string"
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp when the resource was created.\nSee https://google.aip.dev/148#timestamps.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "User-provided description of the endpoint group.\nUsed as additional context for the endpoint group.",
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels are key/value pairs that help to organize and filter resources.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The cloud location of the endpoint group, currently restricted to 'global'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mirroring_deployment_group": {
        "description": "The deployment group that this DIRECT endpoint group is connected to, for example:\n'projects/123456789/locations/global/mirroringDeploymentGroups/my-dg'.\nSee https://google.aip.dev/124.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mirroring_endpoint_group_id": {
        "description": "The ID to use for the endpoint group, which will become the final component\nof the endpoint group's resource name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of this endpoint group, for example:\n'projects/123456789/locations/global/mirroringEndpointGroups/my-eg'.\nSee https://google.aip.dev/122 for more details.",
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
        "description": "The current state of the resource does not match the user's intended state,\nand the system is working to reconcile them. This is part of the normal\noperation (e.g. adding a new association to the group).\nSee https://google.aip.dev/128.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The current state of the endpoint group.\nSee https://google.aip.dev/216.\nPossible values:\nSTATE_UNSPECIFIED\nACTIVE\nCLOSED\nCREATING\nDELETING\nOUT_OF_SYNC\nDELETE_FAILED",
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
        "description": "The timestamp when the resource was most recently updated.\nSee https://google.aip.dev/148#timestamps.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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

func GoogleNetworkSecurityMirroringEndpointGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkSecurityMirroringEndpointGroup), &result)
	return &result
}
