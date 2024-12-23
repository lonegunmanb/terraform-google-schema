package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleManagedKafkaTopic = `{
  "block": {
    "attributes": {
      "cluster": {
        "description": "The cluster name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "configs": {
        "description": "Configuration for the topic that are overridden from the cluster defaults. The key of the map is a Kafka topic property name, for example: 'cleanup.policy=compact', 'compression.type=producer'.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "ID of the location of the Kafka resource. See https://cloud.google.com/managed-kafka/docs/locations for a list of supported locations.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the topic. The 'topic' segment is used when connecting directly to the cluster. Must be in the format 'projects/PROJECT_ID/locations/LOCATION/clusters/CLUSTER_ID/topics/TOPIC_ID'.",
        "description_kind": "plain",
        "type": "string"
      },
      "partition_count": {
        "description": "The number of partitions in a topic. You can increase the partition count for a topic, but you cannot decrease it. Increasing partitions for a topic that uses a key might change how messages are distributed.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replication_factor": {
        "description": "The number of replicas of each partition. A replication factor of 3 is recommended for high availability.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "topic_id": {
        "description": "The ID to use for the topic, which will become the final component of the topic's name. This value is structured like: 'my-topic-name'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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

func GoogleManagedKafkaTopicSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleManagedKafkaTopic), &result)
	return &result
}
