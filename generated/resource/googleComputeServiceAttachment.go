package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeServiceAttachment = `{
  "block": {
    "attributes": {
      "connected_endpoints": {
        "computed": true,
        "description": "An array of the consumer forwarding rules connected to this service\nattachment.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "consumer_network": "string",
              "endpoint": "string",
              "propagated_connection_count": "number",
              "psc_connection_id": "string",
              "status": "string"
            }
          ]
        ]
      },
      "connection_preference": {
        "description": "The connection preference to use for this service attachment. Valid\nvalues include \"ACCEPT_AUTOMATIC\", \"ACCEPT_MANUAL\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "consumer_reject_lists": {
        "description": "An array of projects that are not allowed to connect to this service\nattachment.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_names": {
        "description": "If specified, the domain name will be used during the integration between\nthe PSC connected endpoints and the Cloud DNS. For example, this is a\nvalid domain name: \"p.mycompany.com.\". Current max number of domain names\nsupported is 1.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "enable_proxy_protocol": {
        "description": "If true, enable the proxy protocol which is for supplying client TCP/IP\naddress data in TCP connections that traverse proxies on their way to\ndestination servers.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource. This field is used internally during\nupdates of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the resource. The name must be 1-63 characters long, and\ncomply with RFC1035. Specifically, the name must be 1-63 characters\nlong and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'\nwhich means the first character must be a lowercase letter, and all\nfollowing characters must be a dash, lowercase letter, or digit,\nexcept the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nat_subnets": {
        "description": "An array of subnets that is provided for NAT in this service attachment.",
        "description_kind": "plain",
        "required": true,
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
      "propagated_connection_limit": {
        "computed": true,
        "description": "The number of consumer spokes that connected Private Service Connect endpoints can be propagated to through Network Connectivity Center.\nThis limit lets the service producer limit how many propagated Private Service Connect connections can be established to this service attachment from a single consumer.\n\nIf the connection preference of the service attachment is ACCEPT_MANUAL, the limit applies to each project or network that is listed in the consumer accept list.\nIf the connection preference of the service attachment is ACCEPT_AUTOMATIC, the limit applies to each project that contains a connected endpoint.\n\nIf unspecified, the default propagated connection limit is 250. To explicitly send a zero value, set 'send_propagated_connection_limit_if_zero = true'.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "reconcile_connections": {
        "computed": true,
        "description": "This flag determines whether a consumer accept/reject list change can reconcile the statuses of existing ACCEPTED or REJECTED PSC endpoints.\n\nIf false, connection policy update will only affect existing PENDING PSC endpoints. Existing ACCEPTED/REJECTED endpoints will remain untouched regardless how the connection policy is modified .\nIf true, update will affect both PENDING and ACCEPTED/REJECTED PSC endpoints. For example, an ACCEPTED PSC endpoint will be moved to REJECTED if its project is added to the reject list.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "URL of the region where the resource resides.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "send_propagated_connection_limit_if_zero": {
        "description": "Controls the behavior of propagated_connection_limit.\nWhen false, setting propagated_connection_limit to zero causes the provider to use to the API's default value.\nWhen true, the provider will set propagated_connection_limit to zero.\nDefaults to false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "target_service": {
        "description": "The URL of a service serving the endpoint identified by this service attachment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "consumer_accept_lists": {
        "block": {
          "attributes": {
            "connection_limit": {
              "description": "The number of consumer forwarding rules the consumer project can\ncreate.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "network_url": {
              "description": "The network that is allowed to connect to this service attachment.\nOnly one of project_id_or_num and network_url may be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project_id_or_num": {
              "description": "A project that is allowed to connect to this service attachment.\nOnly one of project_id_or_num and network_url may be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "An array of projects that are allowed to connect to this service\nattachment.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func GoogleComputeServiceAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeServiceAttachment), &result)
	return &result
}
