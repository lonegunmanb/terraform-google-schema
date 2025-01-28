package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkSecuritySecurityProfile = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the security profile was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of the security profile. The Max length is 512 characters.",
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
        "description": "This checksum is computed by the server based on the value of other fields,\nand may be sent on update and delete requests to ensure the client has an up-to-date\nvalue before proceeding.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "A map of key/value label pairs to assign to the resource.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the security profile.\nThe default value is 'global'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the security profile resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "The name of the parent this security profile belongs to.\nFormat: organizations/{organization_id}.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "Server-defined URL of this resource.",
        "description_kind": "plain",
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
      "type": {
        "description": "The type of security profile. Possible values: [\"THREAT_PREVENTION\", \"CUSTOM_MIRRORING\", \"CUSTOM_INTERCEPT\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the security profile was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "custom_intercept_profile": {
        "block": {
          "attributes": {
            "intercept_endpoint_group": {
              "description": "The Intercept Endpoint Group to which matching traffic should be intercepted.\nFormat: projects/{project_id}/locations/global/interceptEndpointGroups/{endpoint_group_id}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The configuration for defining the Intercept Endpoint Group used to\nintercept traffic to third-party firewall appliances.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "custom_mirroring_profile": {
        "block": {
          "attributes": {
            "mirroring_endpoint_group": {
              "description": "The Mirroring Endpoint Group to which matching traffic should be mirrored.\nFormat: projects/{project_id}/locations/global/mirroringEndpointGroups/{endpoint_group_id}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The configuration for defining the Mirroring Endpoint Group used to\nmirror traffic to third-party collectors.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "threat_prevention_profile": {
        "block": {
          "block_types": {
            "severity_overrides": {
              "block": {
                "attributes": {
                  "action": {
                    "description": "Threat action override. Possible values: [\"ALERT\", \"ALLOW\", \"DEFAULT_ACTION\", \"DENY\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "severity": {
                    "description": "Severity level to match. Possible values: [\"CRITICAL\", \"HIGH\", \"INFORMATIONAL\", \"LOW\", \"MEDIUM\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The configuration for overriding threats actions by severity match.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "threat_overrides": {
              "block": {
                "attributes": {
                  "action": {
                    "description": "Threat action. Possible values: [\"ALERT\", \"ALLOW\", \"DEFAULT_ACTION\", \"DENY\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "threat_id": {
                    "description": "Vendor-specific ID of a threat to override.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of threat.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "The configuration for overriding threats actions by threat id match.\nIf a threat is matched both by configuration provided in severity overrides\nand threat overrides, the threat overrides action is applied.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "The threat prevention configuration for the security profile.",
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

func GoogleNetworkSecuritySecurityProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkSecuritySecurityProfile), &result)
	return &result
}
