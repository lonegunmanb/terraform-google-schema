package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkSecurityInterceptEndpointGroupAssociation = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The timestamp when the resource was created.\nSee https://google.aip.dev/148#timestamps.",
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
      "intercept_endpoint_group": {
        "description": "The endpoint group that this association is connected to, for example:\n'projects/123456789/locations/global/interceptEndpointGroups/my-eg'.\nSee https://google.aip.dev/124.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "intercept_endpoint_group_association_id": {
        "description": "The ID to use for the new association, which will become the final\ncomponent of the endpoint group's resource name. If not provided, the\nserver will generate a unique ID.",
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
        "description": "The cloud location of the association, currently restricted to 'global'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "locations": {
        "computed": true,
        "description": "The list of locations where the association is configured. This information\nis retrieved from the linked endpoint group.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "location": "string",
              "state": "string"
            }
          ]
        ]
      },
      "locations_details": {
        "computed": true,
        "deprecated": true,
        "description": "The list of locations where the association is present. This information\nis retrieved from the linked endpoint group, and not configured as part\nof the association itself.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "location": "string",
              "state": "string"
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description": "The resource name of this endpoint group association, for example:\n'projects/123456789/locations/global/interceptEndpointGroupAssociations/my-eg-association'.\nSee https://google.aip.dev/122 for more details.",
        "description_kind": "plain",
        "type": "string"
      },
      "network": {
        "description": "The VPC network that is associated. for example:\n'projects/123456789/global/networks/my-network'.\nSee https://google.aip.dev/124.",
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
      "reconciling": {
        "computed": true,
        "description": "The current state of the resource does not match the user's intended state,\nand the system is working to reconcile them. This part of the normal\noperation (e.g. adding a new location to the target deployment group).\nSee https://google.aip.dev/128.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "Current state of the endpoint group association.\nPossible values:\nSTATE_UNSPECIFIED\nACTIVE\nCREATING\nDELETING\nCLOSED\nOUT_OF_SYNC\nDELETE_FAILED",
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

func GoogleNetworkSecurityInterceptEndpointGroupAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkSecurityInterceptEndpointGroupAssociation), &result)
	return &result
}
