package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleEdgenetworkInterconnectAttachment = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time when the resource was created.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine\nfractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A free-text description of the resource. Max length 1024 characters.",
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
      "interconnect": {
        "description": "The ID of the underlying interconnect that this attachment's traffic will traverse through.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "interconnect_attachment_id": {
        "description": "A unique ID that identifies this interconnect attachment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels associated with this resource.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The Google Cloud region to which the target Distributed Cloud Edge zone belongs.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mtu": {
        "description": "IP (L3) MTU value of the virtual edge cloud. Default value is '1500'. Possible values are: '1500', '9000'.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "The canonical name of this resource, with format\n'projects/{{project}}/locations/{{location}}/zones/{{zone}}/interconnectAttachments/{{interconnect_attachment_id}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "network": {
        "description": "The ID of the network to which this interconnect attachment belongs.\nMust be of the form: 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}'",
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
        "description": "The time when the resource was last updated.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine\nfractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.",
        "description_kind": "plain",
        "type": "string"
      },
      "vlan_id": {
        "description": "VLAN ID provided by user. Must be site-wise unique.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "zone": {
        "description": "The name of the target Distributed Cloud Edge zone.",
        "description_kind": "plain",
        "required": true,
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

func GoogleEdgenetworkInterconnectAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleEdgenetworkInterconnectAttachment), &result)
	return &result
}
