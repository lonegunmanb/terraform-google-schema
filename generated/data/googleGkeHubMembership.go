package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeHubMembership = `{
  "block": {
    "attributes": {
      "authority": {
        "computed": true,
        "description": "Authority encodes how Google will recognize identities from this Membership.\nSee the workload identity documentation for more details:\nhttps://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "issuer": "string"
            }
          ]
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
      "endpoint": {
        "computed": true,
        "description": "If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "gke_cluster": [
                "list",
                [
                  "object",
                  {
                    "resource_link": "string"
                  }
                ]
              ]
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
      "labels": {
        "computed": true,
        "description": "Labels to apply to this membership.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Location of the membership.\nThe default value is 'global'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "membership_id": {
        "description": "The client-provided identifier of the membership.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique identifier of the membership.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleGkeHubMembershipSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeHubMembership), &result)
	return &result
}
