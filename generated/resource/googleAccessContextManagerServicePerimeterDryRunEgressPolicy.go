package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAccessContextManagerServicePerimeterDryRunEgressPolicy = `{
  "block": {
    "attributes": {
      "access_policy_id": {
        "computed": true,
        "description": "The name of the Access Policy this resource belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "The perimeter etag is internally used to prevent overwriting the list of policies on PATCH calls. It is retrieved from the same GET perimeter API call that's used to get the current list of policies. The policy defined in this resource is added or removed from that list, and then this etag is sent with the PATCH call along with the updated policies.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "perimeter": {
        "description": "The name of the Service Perimeter to add this resource to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "title": {
        "description": "Human readable title. Must be unique within the perimeter. Does not affect behavior.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "egress_from": {
        "block": {
          "attributes": {
            "identities": {
              "description": "Identities can be an individual user, service account, Google group,\nor third-party identity. For third-party identity, only single identities\nare supported and other identity types are not supported.The v1 identities\nthat have the prefix user, group and serviceAccount in\nhttps://cloud.google.com/iam/docs/principal-identifiers#v1 are supported.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "identity_type": {
              "description": "Specifies the type of identities that are allowed access to outside the\nperimeter. If left unspecified, then members of 'identities' field will\nbe allowed access. Possible values: [\"ANY_IDENTITY\", \"ANY_USER_ACCOUNT\", \"ANY_SERVICE_ACCOUNT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_restriction": {
              "description": "Whether to enforce traffic restrictions based on 'sources' field. If the 'sources' field is non-empty, then this field must be set to 'SOURCE_RESTRICTION_ENABLED'. Possible values: [\"SOURCE_RESTRICTION_ENABLED\", \"SOURCE_RESTRICTION_DISABLED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "sources": {
              "block": {
                "attributes": {
                  "access_level": {
                    "description": "An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resource": {
                    "description": "A Google Cloud resource that is allowed to egress the perimeter.\nRequests from these resources are allowed to access data outside the perimeter.\nCurrently only projects are allowed. Project format: 'projects/{project_number}'.\nThe resource may be in any Google Cloud organization, not just the\norganization that the perimeter is defined in. '*' is not allowed, the\ncase of allowing all Google Cloud resources only is not supported.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Sources that this EgressPolicy authorizes access from.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Defines conditions on the source of a request causing this 'EgressPolicy' to apply.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "egress_to": {
        "block": {
          "attributes": {
            "external_resources": {
              "description": "A list of external resources that are allowed to be accessed. A request\nmatches if it contains an external resource in this list (Example:\ns3://bucket/path). Currently '*' is not allowed.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resources": {
              "description": "A list of resources, currently only projects in the form\n'projects/\u003cprojectnumber\u003e', that match this to stanza. A request matches\nif it contains a resource in this list. If * is specified for resources,\nthen this 'EgressTo' rule will authorize access to all resources outside\nthe perimeter.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "roles": {
              "description": "A list of IAM roles that represent the set of operations that the sources\nspecified in the corresponding 'EgressFrom'\nare allowed to perform.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "operations": {
              "block": {
                "attributes": {
                  "service_name": {
                    "description": "The name of the API whose methods or permissions the 'IngressPolicy' or\n'EgressPolicy' want to allow. A single 'ApiOperation' with serviceName\nfield set to '*' will allow all methods AND permissions for all services.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "method_selectors": {
                    "block": {
                      "attributes": {
                        "method": {
                          "description": "Value for 'method' should be a valid method name for the corresponding\n'serviceName' in 'ApiOperation'. If '*' used as value for method,\nthen ALL methods and permissions are allowed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "permission": {
                          "description": "Value for permission should be a valid Cloud IAM permission for the\ncorresponding 'serviceName' in 'ApiOperation'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "API methods or permissions to allow. Method or permission must belong\nto the service specified by 'serviceName' field. A single MethodSelector\nentry with '*' specified for the 'method' field will allow all methods\nAND permissions for the service specified in 'serviceName'.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "A list of 'ApiOperations' that this egress rule applies to. A request matches\nif it contains an operation/service in this list.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Defines the conditions on the 'ApiOperation' and destination resources that\ncause this 'EgressPolicy' to apply.",
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

func GoogleAccessContextManagerServicePerimeterDryRunEgressPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAccessContextManagerServicePerimeterDryRunEgressPolicy), &result)
	return &result
}
