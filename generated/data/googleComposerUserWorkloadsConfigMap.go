package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComposerUserWorkloadsConfigMap = `{
  "block": {
    "attributes": {
      "data": {
        "computed": true,
        "description": "The \"data\" field of Kubernetes ConfigMap, organized in key-value pairs.\nFor details see: https://kubernetes.io/docs/concepts/configuration/configmap/",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "environment": {
        "description": "Environment where the Kubernetes ConfigMap will be stored and used.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the Kubernetes ConfigMap.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The location or Compute Engine region for the environment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComposerUserWorkloadsConfigMapSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComposerUserWorkloadsConfigMap), &result)
	return &result
}