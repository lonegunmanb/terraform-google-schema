package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMemorystoreInstanceDesiredUserCreatedEndpoints = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the Memorystore instance these endpoints should be added to.",
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
      "region": {
        "description": "The name of the region of the Memorystore instance these endpoints should be added to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "desired_user_created_endpoints": {
        "block": {
          "block_types": {
            "connections": {
              "block": {
                "block_types": {
                  "psc_connection": {
                    "block": {
                      "attributes": {
                        "connection_type": {
                          "computed": true,
                          "description": "Output Only. Type of a PSC Connection. \n Possible values:\n CONNECTION_TYPE_DISCOVERY \n CONNECTION_TYPE_PRIMARY \n CONNECTION_TYPE_READER",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "forwarding_rule": {
                          "description": "The URI of the consumer side forwarding rule.\nFormat:\nprojects/{project}/regions/{region}/forwardingRules/{forwarding_rule}",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "ip_address": {
                          "description": "The IP allocated on the consumer network for the PSC forwarding rule.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "network": {
                          "description": "The consumer network where the IP address resides, in the form of\nprojects/{project_id}/global/networks/{network_id}.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "project_id": {
                          "computed": true,
                          "description": "The consumer project_id where the forwarding rule is created from.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "psc_connection_id": {
                          "description": "The PSC connection id of the forwarding rule connected to the\nservice attachment.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "psc_connection_status": {
                          "computed": true,
                          "description": "Output Only. The status of the PSC connection: whether a connection exists and ACTIVE or it no longer exists. \n Possible values:\n ACTIVE \n NOT_FOUND",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "service_attachment": {
                          "description": "The service attachment which is the target of the PSC connection, in the form of projects/{project-id}/regions/{region}/serviceAttachments/{service-attachment-id}.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Detailed information of a PSC connection that is created by the customer\nwho owns the cluster.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "A list of desired user endpoints",
          "description_kind": "plain"
        },
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

func GoogleMemorystoreInstanceDesiredUserCreatedEndpointsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMemorystoreInstanceDesiredUserCreatedEndpoints), &result)
	return &result
}
