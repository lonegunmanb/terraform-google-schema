package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVmwareengineExternalAddress = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Creation time of this resource.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and\nup to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "User-provided description for this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "external_ip": {
        "computed": true,
        "description": "The external IP address of a workload VM.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internal_ip": {
        "computed": true,
        "description": "The internal IP address of a workload VM.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The ID of the external IP Address.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "The resource name of the private cloud to create a new external address in.\nResource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.\nFor example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the resource.",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleVmwareengineExternalAddressSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVmwareengineExternalAddress), &result)
	return &result
}
