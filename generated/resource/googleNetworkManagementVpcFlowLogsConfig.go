package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkManagementVpcFlowLogsConfig = `{
  "block": {
    "attributes": {
      "aggregation_interval": {
        "computed": true,
        "description": "Optional. The aggregation interval for the logs. Default value is\nINTERVAL_5_SEC.   Possible values:  AGGREGATION_INTERVAL_UNSPECIFIED INTERVAL_5_SEC INTERVAL_30_SEC INTERVAL_1_MIN INTERVAL_5_MIN INTERVAL_10_MIN INTERVAL_15_MIN",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The time the config was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Optional. The user-supplied description of the VPC Flow Logs configuration. Maximum\nof 512 characters.",
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
      "filter_expr": {
        "description": "Optional. Export filter used to define which VPC Flow Logs should be logged.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flow_sampling": {
        "computed": true,
        "description": "Optional. The value of the field must be in (0, 1]. The sampling rate\nof VPC Flow Logs where 1.0 means all collected logs are reported. Setting the\nsampling rate to 0.0 is not allowed. If you want to disable VPC Flow Logs, use\nthe state field instead. Default value is 1.0.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "interconnect_attachment": {
        "description": "Traffic will be logged from the Interconnect Attachment. Format: projects/{project_id}/regions/{region}/interconnectAttachments/{name}",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Optional. Resource labels to represent user-provided metadata.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource\nwithin its parent collection as described in https://google.aip.dev/122. See documentation\nfor resource type 'networkmanagement.googleapis.com/VpcFlowLogsConfig'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metadata": {
        "computed": true,
        "description": "Optional. Configures whether all, none or a subset of metadata fields\nshould be added to the reported VPC flow logs. Default value is INCLUDE_ALL_METADATA.\n  Possible values:  METADATA_UNSPECIFIED INCLUDE_ALL_METADATA EXCLUDE_ALL_METADATA CUSTOM_METADATA",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metadata_fields": {
        "description": "Optional. Custom metadata fields to include in the reported VPC flow\nlogs. Can only be specified if \\\"metadata\\\" was set to CUSTOM_METADATA.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "Identifier. Unique name of the configuration using the form:     'projects/{project_id}/locations/global/vpcFlowLogsConfigs/{vpc_flow_logs_config_id}'",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Optional. The state of the VPC Flow Log configuration. Default value\nis ENABLED. When creating a new configuration, it must be enabled.\nPossible values: STATE_UNSPECIFIED ENABLED DISABLED",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_resource_state": {
        "computed": true,
        "description": "Describes the state of the configured target resource for diagnostic\npurposes.\nPossible values:\nTARGET_RESOURCE_STATE_UNSPECIFIED\nTARGET_RESOURCE_EXISTS\nTARGET_RESOURCE_DOES_NOT_EXIST",
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
      "update_time": {
        "computed": true,
        "description": "Output only. The time the config was updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_flow_logs_config_id": {
        "description": "Required. ID of the 'VpcFlowLogsConfig'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpn_tunnel": {
        "description": "Traffic will be logged from the VPN Tunnel. Format: projects/{project_id}/regions/{region}/vpnTunnels/{name}",
        "description_kind": "plain",
        "optional": true,
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

func GoogleNetworkManagementVpcFlowLogsConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkManagementVpcFlowLogsConfig), &result)
	return &result
}
