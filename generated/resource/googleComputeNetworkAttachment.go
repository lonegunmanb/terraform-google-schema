package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeNetworkAttachment = `{
  "block": {
    "attributes": {
      "connection_endpoints": {
        "computed": true,
        "description": "An array of connections for all the producers connected to this network attachment.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ip_address": "string",
              "project_id_or_num": "string",
              "secondary_ip_cidr_ranges": "string",
              "status": "string",
              "subnetwork": "string"
            }
          ]
        ]
      },
      "connection_preference": {
        "description": "The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules. Possible values: [\"ACCEPT_AUTOMATIC\", \"ACCEPT_MANUAL\", \"INVALID\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when you create the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource. A hash of the contents stored in this object. This\nfield is used in optimistic locking. An up-to-date fingerprint must be provided in order to patch.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "The unique identifier for the resource type. The server generates this identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "kind": {
        "computed": true,
        "description": "Type of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "The URL of the network which the Network Attachment belongs to. Practically it is inferred by fetching the network of the first subnetwork associated.\nBecause it is required that all the subnetworks must be from the same network, it is assured that the Network Attachment belongs to the same network as all the subnetworks.",
        "description_kind": "plain",
        "type": "string"
      },
      "producer_accept_lists": {
        "description": "Projects that are allowed to connect to this network attachment. The project can be specified using its id or number.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "producer_reject_lists": {
        "description": "Projects that are not allowed to connect to this network attachment. The project can be specified using its id or number.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "URL of the region where the network attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "Server-defined URL for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "self_link_with_id": {
        "computed": true,
        "description": "Server-defined URL for this resource's resource id.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnetworks": {
        "description": "An array of URLs where each entry is the URL of a subnet provided by the service consumer to use for endpoints in the producers that connect to this network attachment.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
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

func GoogleComputeNetworkAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeNetworkAttachment), &result)
	return &result
}
