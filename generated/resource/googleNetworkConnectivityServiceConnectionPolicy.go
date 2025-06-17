package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkConnectivityServiceConnectionPolicy = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The timestamp when the resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Free-text description of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_labels": {
        "computed": true,
        "description": "All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "etag": {
        "computed": true,
        "description": "The etag is computed by the server, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "infrastructure": {
        "computed": true,
        "description": "The type of underlying resources used to create the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "User-defined labels.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the ServiceConnectionPolicy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of a ServiceConnectionPolicy. Format: projects/{project}/locations/{location}/serviceConnectionPolicies/{service_connection_policy} See: https://google.aip.dev/122#fields-representing-resource-names",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.",
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
      "psc_connections": {
        "computed": true,
        "description": "Information about each Private Service Connect connection.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "consumer_address": "string",
              "consumer_forwarding_rule": "string",
              "consumer_target_project": "string",
              "error": [
                "list",
                [
                  "object",
                  {
                    "code": "number",
                    "details": [
                      "list",
                      [
                        "map",
                        "string"
                      ]
                    ],
                    "message": "string"
                  }
                ]
              ],
              "error_info": [
                "list",
                [
                  "object",
                  {
                    "domain": "string",
                    "metadata": [
                      "map",
                      "string"
                    ],
                    "reason": "string"
                  }
                ]
              ],
              "error_type": "string",
              "gce_operation": "string",
              "psc_connection_id": "string",
              "state": "string"
            }
          ]
        ]
      },
      "service_class": {
        "description": "The service class identifier for which this ServiceConnectionPolicy is for. The service class identifier is a unique, symbolic representation of a ServiceClass.\nIt is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp when the resource was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "psc_config": {
        "block": {
          "attributes": {
            "allowed_google_producers_resource_hierarchy_level": {
              "description": "List of Projects, Folders, or Organizations from where the Producer instance can be within. For example,\na network administrator can provide both 'organizations/foo' and 'projects/bar' as\nallowed_google_producers_resource_hierarchy_levels. This allowlists this network to connect with any Producer\ninstance within the 'foo' organization or the 'bar' project. By default,\nallowedGoogleProducersResourceHierarchyLevel is empty. The format for each\nallowedGoogleProducersResourceHierarchyLevel is / where is one of 'projects', 'folders', or 'organizations'\nand is either the ID or the number of the resource type. Format for each\nallowedGoogleProducersResourceHierarchyLevel value: 'projects/' or 'folders/' or 'organizations/' Eg.\n[projects/my-project-id, projects/567, folders/891, organizations/123]",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "limit": {
              "description": "Max number of PSC connections for this policy.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "producer_instance_location": {
              "computed": true,
              "description": "ProducerInstanceLocation is used to specify which authorization mechanism to use to determine which projects\nthe Producer instance can be within. Possible values: [\"PRODUCER_INSTANCE_LOCATION_UNSPECIFIED\", \"CUSTOM_RESOURCE_HIERARCHY_LEVELS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnetworks": {
              "description": "IDs of the subnetworks or fully qualified identifiers for the subnetworks",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Configuration used for Private Service Connect connections. Used when Infrastructure is PSC.",
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GoogleNetworkConnectivityServiceConnectionPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkConnectivityServiceConnectionPolicy), &result)
	return &result
}
