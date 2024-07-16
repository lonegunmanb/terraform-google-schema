package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googlePrivilegedAccessManagerEntitlement = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. Create time stamp. A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\"",
        "description_kind": "plain",
        "type": "string"
      },
      "entitlement_id": {
        "description": "The ID to use for this Entitlement. This will become the last part of the resource name.\nThis value should be 4-63 characters, and valid characters are \"[a-z]\", \"[0-9]\", and \"-\". The first character should be from [a-z].\nThis value should be unique among all other Entitlements under the specified 'parent'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "For Resource freshness validation (https://google.aip.dev/154)",
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
        "description": "The region of the Entitlement resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_request_duration": {
        "description": "The maximum amount of time for which access would be granted for a request.\nA requester can choose to ask for access for less than this duration but never more.\nFormat: calculate the time in seconds and concatenate it with 's' i.e. 2 hours = \"7200s\", 45 minutes = \"2700s\"",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Output Only. The entitlement's name follows a hierarchical structure, comprising the organization, folder, or project, alongside the region and a unique entitlement ID.\nFormats: organizations/{organization-number}/locations/{region}/entitlements/{entitlement-id}, folders/{folder-number}/locations/{region}/entitlements/{entitlement-id}, and projects/{project-id|project-number}/locations/{region}/entitlements/{entitlement-id}.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "Format: projects/{project-id|project-number} or organizations/{organization-number} or folders/{folder-number}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. The current state of the Entitlement.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Update time stamp. A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "additional_notification_targets": {
        "block": {
          "attributes": {
            "admin_email_recipients": {
              "description": "Optional. Additional email addresses to be notified when a principal(requester) is granted access.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "requester_email_recipients": {
              "description": "Optional. Additional email address to be notified about an eligible entitlement.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "AdditionalNotificationTargets includes email addresses to be notified.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "approval_workflow": {
        "block": {
          "block_types": {
            "manual_approvals": {
              "block": {
                "attributes": {
                  "require_approver_justification": {
                    "description": "Optional. Do the approvers need to provide a justification for their actions?",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "steps": {
                    "block": {
                      "attributes": {
                        "approvals_needed": {
                          "description": "How many users from the above list need to approve.\nIf there are not enough distinct users in the list above then the workflow\nwill indefinitely block. Should always be greater than 0. Currently 1 is the only\nsupported value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "approver_email_recipients": {
                          "description": "Optional. Additional email addresses to be notified when a grant is pending approval.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "approvers": {
                          "block": {
                            "attributes": {
                              "principals": {
                                "description": "Users who are being allowed for the operation. Each entry should be a valid v1 IAM Principal Identifier. Format for these is documented at: https://cloud.google.com/iam/docs/principal-identifiers#v1",
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              }
                            },
                            "description": "The potential set of approvers in this step. This list should contain at only one entry.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "List of approval steps in this workflow. These steps would be followed in the specified order sequentially.  1 step is supported for now.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A manual approval workflow where users who are designated as approvers need to call the ApproveGrant/DenyGrant APIs for an Grant.\nThe workflow can consist of multiple serial steps where each step defines who can act as Approver in that step and how many of those users should approve before the workflow moves to the next step.\nThis can be used to create approval workflows such as\n* Require an approval from any user in a group G.\n* Require an approval from any k number of users from a Group G.\n* Require an approval from any user in a group G and then from a user U. etc.\nA single user might be part of 'approvers' ACL for multiple steps in this workflow but they can only approve once and that approval will only be considered to satisfy the approval step at which it was granted.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The approvals needed before access will be granted to a requester.\nNo approvals will be needed if this field is null. Different types of approval workflows that can be used to gate privileged access granting.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "eligible_users": {
        "block": {
          "attributes": {
            "principals": {
              "description": "Users who are being allowed for the operation. Each entry should be a valid v1 IAM Principal Identifier. Format for these is documented at \"https://cloud.google.com/iam/docs/principal-identifiers#v1\"",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "Who can create Grants using Entitlement. This list should contain at most one entry",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "privileged_access": {
        "block": {
          "block_types": {
            "gcp_iam_access": {
              "block": {
                "attributes": {
                  "resource": {
                    "description": "Name of the resource.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "resource_type": {
                    "description": "The type of this resource.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "role_bindings": {
                    "block": {
                      "attributes": {
                        "condition_expression": {
                          "description": "The expression field of the IAM condition to be associated with the role. If specified, a user with an active grant for this entitlement would be able to access the resource only if this condition evaluates to true for their request.\nhttps://cloud.google.com/iam/docs/conditions-overview#attributes.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role": {
                          "description": "IAM role to be granted. https://cloud.google.com/iam/docs/roles-overview.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Role bindings to be created on successful grant.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "GcpIamAccess represents IAM based access control on a GCP resource. Refer to https://cloud.google.com/iam/docs to understand more about IAM.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Privileged access that this service can be used to gate.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "requester_justification_config": {
        "block": {
          "block_types": {
            "not_mandatory": {
              "block": {
                "description": "The justification is not mandatory but can be provided in any of the supported formats.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "unstructured": {
              "block": {
                "description": "The requester has to provide a justification in the form of free flowing text.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Defines the ways in which a requester should provide the justification while requesting for access.",
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

func GooglePrivilegedAccessManagerEntitlementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePrivilegedAccessManagerEntitlement), &result)
	return &result
}
