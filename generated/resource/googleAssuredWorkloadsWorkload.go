package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAssuredWorkloadsWorkload = `{
  "block": {
    "attributes": {
      "billing_account": {
        "description": "Optional. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form ` + "`" + `billingAccounts/{billing_account_id}` + "`" + `. For example, ` + "`" + `billingAccounts/012345-567890-ABCDEF` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compliance_regime": {
        "description": "Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS, HIPAA, HITRUST, EU_REGIONS_AND_SUPPORT, CA_REGIONS_AND_SUPPORT, ITAR, AU_REGIONS_AND_US_SUPPORT, ASSURED_WORKLOADS_FOR_PARTNERS, ISR_REGIONS, ISR_REGIONS_AND_SUPPORT, CA_PROTECTED_B, IL5, IL2, JP_REGIONS_AND_SUPPORT, KSA_REGIONS_AND_SUPPORT_WITH_SOVEREIGNTY_CONTROLS, REGIONAL_CONTROLS, HEALTHCARE_AND_LIFE_SCIENCES_CONTROLS, HEALTHCARE_AND_LIFE_SCIENCES_CONTROLS_WITH_US_SUPPORT",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "compliance_status": {
        "computed": true,
        "description": "Output only. Count of active Violations in the Workload.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "acknowledged_violation_count": [
                "list",
                "number"
              ],
              "active_violation_count": [
                "list",
                "number"
              ]
            }
          ]
        ]
      },
      "compliant_but_disallowed_services": {
        "computed": true,
        "description": "Output only. Urls for services which are compliant for this Assured Workload, but which are currently disallowed by the ResourceUsageRestriction org policy. Invoke workloads.restrictAllowedResources endpoint to allow your project developers to use these services in their environment.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Immutable. The Workload creation timestamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload",
        "description_kind": "plain",
        "required": true,
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
      "ekm_provisioning_response": {
        "computed": true,
        "description": "Optional. Represents the Ekm Provisioning State of the given workload.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ekm_provisioning_error_domain": "string",
              "ekm_provisioning_error_mapping": "string",
              "ekm_provisioning_state": "string"
            }
          ]
        ]
      },
      "enable_sovereign_controls": {
        "description": "Optional. Indicates the sovereignty status of the given workload. Currently meant to be used by Europe/Canada customers.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kaj_enrollment_state": {
        "computed": true,
        "description": "Output only. Represents the KAJ enrollment state of the given workload. Possible values: KAJ_ENROLLMENT_STATE_UNSPECIFIED, KAJ_ENROLLMENT_STATE_PENDING, KAJ_ENROLLMENT_STATE_COMPLETE",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Optional. Labels applied to the workload.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field ` + "`" + `effective_labels` + "`" + ` for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Output only. The resource name of the workload.",
        "description_kind": "plain",
        "type": "string"
      },
      "organization": {
        "description": "The organization for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "partner": {
        "description": "Optional. Partner regime associated with this workload. Possible values: PARTNER_UNSPECIFIED, LOCAL_CONTROLS_BY_S3NS, SOVEREIGN_CONTROLS_BY_T_SYSTEMS, SOVEREIGN_CONTROLS_BY_SIA_MINSAIT, SOVEREIGN_CONTROLS_BY_PSN, SOVEREIGN_CONTROLS_BY_CNTXT, SOVEREIGN_CONTROLS_BY_CNTXT_NO_EKM",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "partner_services_billing_account": {
        "description": "Optional. Input only. Billing account necessary for purchasing services from Sovereign Partners. This field is required for creating SIA/PSN/CNTXT partner workloads. The caller should have 'billing.resourceAssociations.create' IAM permission on this billing-account. The format of this string is billingAccounts/AAAAAA-BBBBBB-CCCCCC.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provisioned_resources_parent": {
        "description": "Input only. The parent resource for the resources managed by this Assured Workload. May be either empty or a folder resource which is a child of the Workload parent. If not specified all resources are created under the parent organization. Format: folders/{folder_id}",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resources": {
        "computed": true,
        "description": "Output only. The resources associated with this workload. These resources will be created when creating the workload. If any of the projects already exist, the workload creation will fail. Always read only.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "resource_id": "number",
              "resource_type": "string"
            }
          ]
        ]
      },
      "saa_enrollment_response": {
        "computed": true,
        "description": "Output only. Represents the SAA enrollment response of the given workload. SAA enrollment response is queried during workloads.get call. In failure cases, user friendly error message is shown in SAA details page.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "setup_errors": [
                "list",
                "string"
              ],
              "setup_status": "string"
            }
          ]
        ]
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "violation_notifications_enabled": {
        "computed": true,
        "description": "Optional. Indicates whether the e-mail notification for a violation is enabled for a workload. This value will be by default True, and if not present will be considered as true. This should only be updated via updateWorkload call. Any Changes to this field during the createWorkload call will not be honored. This will always be true while creating the workload.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "kms_settings": {
        "block": {
          "attributes": {
            "next_rotation_time": {
              "description": "Required. Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rotation_period": {
              "description": "Required. Input only. Immutable. will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "**DEPRECATED** Input only. Settings used to create a CMEK crypto key. When set, a project with a KMS CMEK key is provisioned. This field is deprecated as of Feb 28, 2022. In order to create a Keyring, callers should specify, ENCRYPTION_KEYS_PROJECT or KEYRING in ResourceSettings.resource_type field.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "partner_permissions": {
        "block": {
          "attributes": {
            "assured_workloads_monitoring": {
              "description": "Optional. Allow partner to view violation alerts.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "data_logs_viewer": {
              "description": "Allow the partner to view inspectability logs and monitoring violations.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "service_access_approver": {
              "description": "Optional. Allow partner to view access approval logs.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Optional. Permissions granted to the AW Partner SA account for the customer workload",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "resource_settings": {
        "block": {
          "attributes": {
            "display_name": {
              "description": "User-assigned resource display name. If not empty it will be used to create a resource with the specified name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_id": {
              "description": "Resource identifier. For a project this represents projectId. If the project is already taken, the workload creation will fail. For KeyRing, this represents the keyring_id. For a folder, don't set this value as folder_id is assigned by Google.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_type": {
              "description": "Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.",
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
      },
      "workload_options": {
        "block": {
          "attributes": {
            "kaj_enrollment_type": {
              "description": "Indicates type of KAJ enrollment for the workload. Currently, only specifiying KEY_ACCESS_TRANSPARENCY_OFF is implemented to not enroll in KAT-level KAJ enrollment for Regional Controls workloads. Possible values: KAJ_ENROLLMENT_TYPE_UNSPECIFIED, FULL_KAJ, EKM_ONLY, KEY_ACCESS_TRANSPARENCY_OFF",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Optional. Used to specify certain options for a workload during workload creation - currently only supporting KAT Optionality for Regional Controls workloads.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleAssuredWorkloadsWorkloadSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAssuredWorkloadsWorkload), &result)
	return &result
}
