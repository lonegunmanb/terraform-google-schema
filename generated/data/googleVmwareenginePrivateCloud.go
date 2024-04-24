package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVmwareenginePrivateCloud = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "User-provided description for this private cloud.",
        "description_kind": "plain",
        "type": "string"
      },
      "hcx": {
        "computed": true,
        "description": "Details about a HCX Cloud Manager appliance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "fqdn": "string",
              "internal_ip": "string",
              "state": "string",
              "version": "string"
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
      "location": {
        "description": "The location where the PrivateCloud should reside.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "management_cluster": {
        "computed": true,
        "description": "The management cluster for this private cloud. This used for creating and managing the default cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cluster_id": "string",
              "node_type_configs": [
                "set",
                [
                  "object",
                  {
                    "custom_core_count": "number",
                    "node_count": "number",
                    "node_type_id": "string"
                  }
                ]
              ],
              "stretched_cluster_config": [
                "list",
                [
                  "object",
                  {
                    "preferred_location": "string",
                    "secondary_location": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "name": {
        "description": "The ID of the PrivateCloud.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_config": {
        "computed": true,
        "description": "Network configuration in the consumer project with which the peering has to be done.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dns_server_ip": "string",
              "management_cidr": "string",
              "management_ip_address_layout_version": "number",
              "vmware_engine_network": "string",
              "vmware_engine_network_canonical": "string"
            }
          ]
        ]
      },
      "nsx": {
        "computed": true,
        "description": "Details about a NSX Manager appliance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "fqdn": "string",
              "internal_ip": "string",
              "state": "string",
              "version": "string"
            }
          ]
        ]
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the resource. New values may be added to this enum when appropriate.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "Initial type of the private cloud. Possible values: [\"STANDARD\", \"TIME_LIMITED\", \"STRETCHED\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "System-generated unique identifier for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "vcenter": {
        "computed": true,
        "description": "Details about a vCenter Server management appliance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "fqdn": "string",
              "internal_ip": "string",
              "state": "string",
              "version": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleVmwareenginePrivateCloudSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVmwareenginePrivateCloud), &result)
	return &result
}
