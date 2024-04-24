package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVmwareengineCluster = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "management": {
        "computed": true,
        "description": "True if the cluster is a management cluster; false otherwise.\nThere can only be one management cluster in a private cloud and it has to be the first one.",
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "description": "The ID of the Cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_type_configs": {
        "computed": true,
        "description": "The map of cluster node types in this cluster,\nwhere the key is canonical identifier of the node type (corresponds to the NodeType).",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "custom_core_count": "number",
              "node_count": "number",
              "node_type_id": "string"
            }
          ]
        ]
      },
      "parent": {
        "description": "The resource name of the private cloud to create a new cluster in.\nResource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.\nFor example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the Cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "System-generated unique identifier for the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleVmwareengineClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVmwareengineCluster), &result)
	return &result
}
