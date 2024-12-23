package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeHubMembershipBinding = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the MembershipBinding was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "Time the MembershipBinding was deleted in UTC.",
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
      "labels": {
        "computed": true,
        "description": "Labels for this Membership binding.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Location of the membership",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "membership_binding_id": {
        "description": "The client-provided identifier of the membership binding.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "membership_id": {
        "description": "Id of the membership",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name for the membershipbinding itself",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scope": {
        "computed": true,
        "description": "A Workspace resource name in the format\n'projects/*/locations/*/scopes/*'.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the membership binding resource.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "code": "string"
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
        "description": "Google-generated UUID for this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the MembershipBinding was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleGkeHubMembershipBindingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeHubMembershipBinding), &result)
	return &result
}
