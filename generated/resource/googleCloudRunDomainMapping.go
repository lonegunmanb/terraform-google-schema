package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudRunDomainMapping = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the cloud run instance. eg us-central1",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Name should be a [verified](https://support.google.com/webmasters/answer/9008080) domain",
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
      "status": {
        "computed": true,
        "description": "The current status of the DomainMapping.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "conditions": [
                "list",
                [
                  "object",
                  {
                    "message": "string",
                    "reason": "string",
                    "status": "string",
                    "type": "string"
                  }
                ]
              ],
              "mapped_route_name": "string",
              "observed_generation": "number",
              "resource_records": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "rrdata": "string",
                    "type": "string"
                  }
                ]
              ]
            }
          ]
        ]
      }
    },
    "block_types": {
      "metadata": {
        "block": {
          "attributes": {
            "annotations": {
              "computed": true,
              "description": "Annotations is a key value map stored with a resource that\nmay be set by external tools to store and retrieve arbitrary metadata. More\ninfo: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations\n\n**Note**: The Cloud Run API may add additional annotations that were not provided in your config.\nIf terraform plan shows a diff where a server-side annotation is added, you can add it to your config\nor apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "generation": {
              "computed": true,
              "description": "A sequence number representing a specific generation of the desired state.",
              "description_kind": "plain",
              "type": "number"
            },
            "labels": {
              "computed": true,
              "description": "Map of string keys and values that can be used to organize and categorize\n(scope and select) objects. May match selectors of replication controllers\nand routes.\nMore info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "namespace": {
              "description": "In Cloud Run the namespace must be equal to either the\nproject ID or project number.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_version": {
              "computed": true,
              "description": "An opaque value that represents the internal version of this object that\ncan be used by clients to determine when objects have changed. May be used\nfor optimistic concurrency, change detection, and the watch operation on a\nresource or set of resources. They may only be valid for a\nparticular resource or set of resources.\n\nMore info:\nhttps://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
              "description_kind": "plain",
              "type": "string"
            },
            "self_link": {
              "computed": true,
              "description": "SelfLink is a URL representing this object.",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "UID is a unique id generated by the server on successful creation of a resource and is not\nallowed to change on PUT operations.\n\nMore info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Metadata associated with this DomainMapping.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "spec": {
        "block": {
          "attributes": {
            "certificate_mode": {
              "description": "The mode of the certificate. Default value: \"AUTOMATIC\" Possible values: [\"NONE\", \"AUTOMATIC\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "force_override": {
              "description": "If set, the mapping will override any mapping set before this spec was set.\nIt is recommended that the user leaves this empty to receive an error\nwarning about a potential conflict and only set it once the respective UI\nhas given such a warning.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "route_name": {
              "description": "The name of the Cloud Run Service that this DomainMapping applies to.\nThe route must exist.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The spec for this DomainMapping.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
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

func GoogleCloudRunDomainMappingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudRunDomainMapping), &result)
	return &result
}
