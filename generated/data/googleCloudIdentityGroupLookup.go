package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudIdentityGroupLookup = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The [resource name](https://cloud.google.com/apis/design/resource_names) of the looked-up Group.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "group_key": {
        "block": {
          "attributes": {
            "id": {
              "description": "The ID of the entity. For Google-managed entities, the id should be the email address of an existing group or user.\nFor external-identity-mapped entities, the id must be a string conforming to the Identity Source's requirements.\nMust be unique within a namespace.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "namespace": {
              "description": "The namespace in which the entity exists. If not specified, the EntityKey represents a Google-managed entity such as a Google user or a Google Group.\nIf specified, the EntityKey represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of identitysources/{identity_source}.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The EntityKey of the Group to lookup. A unique identifier for an entity in the Cloud Identity Groups API.\nAn entity can represent either a group with an optional namespace or a user without a namespace.\nThe combination of id and namespace must be unique; however, the same id can be used with different namespaces.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCloudIdentityGroupLookupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudIdentityGroupLookup), &result)
	return &result
}
