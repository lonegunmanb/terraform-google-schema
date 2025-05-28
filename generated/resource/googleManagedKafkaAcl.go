package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleManagedKafkaAcl = `{
  "block": {
    "attributes": {
      "acl_id": {
        "description": "The ID to use for the acl, which will become the final component of the acl's name. The structure of 'aclId' defines the Resource Pattern (resource_type, resource_name, pattern_type) of the acl. 'aclId' is structured like one of the following:\nFor acls on the cluster: 'cluster'\nFor acls on a single resource within the cluster: 'topic/{resource_name}' 'consumerGroup/{resource_name}' 'transactionalId/{resource_name}'\nFor acls on all resources that match a prefix: 'topicPrefixed/{resource_name}' 'consumerGroupPrefixed/{resource_name}' 'transactionalIdPrefixed/{resource_name}'\nFor acls on all resources of a given type (i.e. the wildcard literal '*''): 'allTopics' (represents 'topic/*') 'allConsumerGroups' (represents 'consumerGroup/*') 'allTransactionalIds' (represents 'transactionalId/*').",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster": {
        "description": "The cluster name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "'etag' is used for concurrency control. An 'etag' is returned in the\nresponse to 'GetAcl' and 'CreateAcl'. Callers are required to put that etag\nin the request to 'UpdateAcl' to ensure that their change will be applied\nto the same version of the acl that exists in the Kafka Cluster.\n\nA terminal 'T' character in the etag indicates that the AclEntries were\ntruncated due to repeated field limits.",
        "description_kind": "plain",
        "type": "string"
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
        "description": "The name of the acl. The 'ACL_ID' segment is used when connecting directly to the cluster. Must be in the format 'projects/PROJECT_ID/locations/LOCATION/clusters/CLUSTER_ID/acls/ACL_ID'.",
        "description_kind": "plain",
        "type": "string"
      },
      "pattern_type": {
        "computed": true,
        "description": "The acl pattern type derived from the name. One of: LITERAL, PREFIXED.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_name": {
        "computed": true,
        "description": "The acl resource name derived from the name. For cluster resource_type, this is always \"kafka-cluster\". Can be the wildcard literal \"*\".",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The acl resource type derived from the name. One of: CLUSTER, TOPIC, GROUP, TRANSACTIONAL_ID.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "acl_entries": {
        "block": {
          "attributes": {
            "host": {
              "description": "The host. Must be set to \"*\" for Managed Service for Apache Kafka.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "operation": {
              "description": "The operation type. Allowed values are (case insensitive): ALL, READ,\nWRITE, CREATE, DELETE, ALTER, DESCRIBE, CLUSTER_ACTION, DESCRIBE_CONFIGS,\nALTER_CONFIGS, and IDEMPOTENT_WRITE. See https://kafka.apache.org/documentation/#operations_resources_and_protocols\nfor valid combinations of resource_type and operation for different Kafka API requests.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "permission_type": {
              "description": "The permission type. Accepted values are (case insensitive): ALLOW, DENY.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "principal": {
              "description": "The principal. Specified as Google Cloud account, with the Kafka StandardAuthorizer prefix User:\". For example: \"User:test-kafka-client@test-project.iam.gserviceaccount.com\". Can be the wildcard \"User:*\" to refer to all users.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The acl entries that apply to the resource pattern. The maximum number of allowed entries is 100.",
          "description_kind": "plain"
        },
        "min_items": 1,
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

func GoogleManagedKafkaAclSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleManagedKafkaAcl), &result)
	return &result
}
