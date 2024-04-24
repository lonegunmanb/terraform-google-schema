package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVmwareengineSubnet = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Creation time of this resource.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and\nup to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "dhcp_address_ranges": {
        "computed": true,
        "description": "DHCP address ranges.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "first_address": "string",
              "last_address": "string"
            }
          ]
        ]
      },
      "gateway_id": {
        "computed": true,
        "description": "The canonical identifier of the logical router that this subnet is attached to.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_ip": {
        "computed": true,
        "description": "The IP address of the gateway of this subnet. Must fall within the IP prefix defined above.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_cidr_range": {
        "computed": true,
        "description": "The IP address range of the subnet in CIDR format.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The ID of the subnet. For userDefined subnets, this name should be in the format of \"service-n\",\nwhere n ranges from 1 to 5.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "The resource name of the private cloud to create a new subnet in.\nResource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.\nFor example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "standard_config": {
        "computed": true,
        "description": "Whether the NSX-T configuration in the backend follows the standard configuration supported by Google Cloud.\nIf false, the subnet cannot be modified through Google Cloud, only through NSX-T directly.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "State of the subnet.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of the subnet.",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "System-generated unique identifier for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Last updated time of this resource.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine\nfractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "vlan_id": {
        "computed": true,
        "description": "VLAN ID of the VLAN on which the subnet is configured.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleVmwareengineSubnetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVmwareengineSubnet), &result)
	return &result
}
