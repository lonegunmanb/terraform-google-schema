package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVertexAiRagEngineConfig = `{
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
        "description": "The resource name of the Dataset. This value is set by Google.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The region of the RagEngineConfig. eg us-central1",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "rag_managed_db_config": {
        "block": {
          "block_types": {
            "basic": {
              "block": {
                "description": "Basic tier is a cost-effective and low compute tier suitable for the following cases: Experimenting with RagManagedDb, Small data size, Latency insensitive workload, Only using RAG Engine with external vector DBs.\nNOTE: This is the default tier if not explicitly chosen.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "scaled": {
              "block": {
                "description": "Scaled tier offers production grade performance along with autoscaling functionality. It is suitable for customers with large amounts of data or performance sensitive workloads.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "unprovisioned": {
              "block": {
                "description": "Disables the RAG Engine service and deletes all your data held within this service. This will halt the billing of the service.\nNOTE: Once deleted the data cannot be recovered. To start using RAG Engine again, you will need to update the tier by calling the UpdateRagEngineConfig API.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Required. The config of the RagManagedDb used by RagEngine.",
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

func GoogleVertexAiRagEngineConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVertexAiRagEngineConfig), &result)
	return &result
}
