package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googlePrivilegedAccessManagerEntitlement = `{
  "block": {
    "attributes": {
      "additional_notification_targets": {
        "computed": true,
        "description": "AdditionalNotificationTargets includes email addresses to be notified.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "admin_email_recipients": [
                "set",
                "string"
              ],
              "requester_email_recipients": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "approval_workflow": {
        "computed": true,
        "description": "The approvals needed before access will be granted to a requester.\nNo approvals will be needed if this field is null. Different types of approval workflows that can be used to gate privileged access granting.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "manual_approvals": [
                "list",
                [
                  "object",
                  {
                    "require_approver_justification": "bool",
                    "steps": [
                      "list",
                      [
                        "object",
                        {
                          "approvals_needed": "number",
                          "approver_email_recipients": [
                            "set",
                            "string"
                          ],
                          "approvers": [
                            "list",
                            [
                              "object",
                              {
                                "principals": [
                                  "set",
                                  "string"
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Create time stamp. A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\"",
        "description_kind": "plain",
        "type": "string"
      },
      "eligible_users": {
        "computed": true,
        "description": "Who can create Grants using Entitlement. This list should contain at most one entry",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "principals": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "entitlement_id": {
        "description": "The ID to use for this Entitlement. This will become the last part of the resource name.\nThis value should be 4-63 characters, and valid characters are \"[a-z]\", \"[0-9]\", and \"-\". The first character should be from [a-z].\nThis value should be unique among all other Entitlements under the specified 'parent'.",
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "max_request_duration": {
        "computed": true,
        "description": "The maximum amount of time for which access would be granted for a request.\nA requester can choose to ask for access for less than this duration but never more.\nFormat: calculate the time in seconds and concatenate it with 's' i.e. 2 hours = \"7200s\", 45 minutes = \"2700s\"",
        "description_kind": "plain",
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
        "optional": true,
        "type": "string"
      },
      "privileged_access": {
        "computed": true,
        "description": "Privileged access that this service can be used to gate.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "gcp_iam_access": [
                "list",
                [
                  "object",
                  {
                    "resource": "string",
                    "resource_type": "string",
                    "role_bindings": [
                      "list",
                      [
                        "object",
                        {
                          "condition_expression": "string",
                          "role": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "requester_justification_config": {
        "computed": true,
        "description": "Defines the ways in which a requester should provide the justification while requesting for access.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "not_mandatory": [
                "list",
                [
                  "object",
                  {}
                ]
              ],
              "unstructured": [
                "list",
                [
                  "object",
                  {}
                ]
              ]
            }
          ]
        ]
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
    "description_kind": "plain"
  },
  "version": 0
}`

func GooglePrivilegedAccessManagerEntitlementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePrivilegedAccessManagerEntitlement), &result)
	return &result
}
