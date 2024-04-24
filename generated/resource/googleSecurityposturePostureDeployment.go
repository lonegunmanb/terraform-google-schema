package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecurityposturePostureDeployment = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the posture deployment was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the posture deployment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "desired_posture_id": {
        "computed": true,
        "description": "This is an output only optional field which will be filled in case when\nPostureDeployment state is UPDATE_FAILED or CREATE_FAILED or DELETE_FAILED.\nIt denotes the desired posture to be deployed.",
        "description_kind": "plain",
        "type": "string"
      },
      "desired_posture_revision_id": {
        "computed": true,
        "description": "This is an output only optional field which will be filled in case when\nPostureDeployment state is UPDATE_FAILED or CREATE_FAILED or DELETE_FAILED.\nIt denotes the desired posture revision_id to be deployed.",
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "For Resource freshness validation (https://google.aip.dev/154)",
        "description_kind": "plain",
        "type": "string"
      },
      "failure_message": {
        "computed": true,
        "description": "This is a output only optional field which will be filled in case where\nPostureDeployment enters a failure state like UPDATE_FAILED or\nCREATE_FAILED or DELETE_FAILED. It will have the failure message for posture deployment's\nCREATE/UPDATE/DELETE methods.",
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
        "description": "The location of the resource, eg. global'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the posture deployment instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The parent of the resource, an organization. Format should be 'organizations/{organization_id}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "posture_deployment_id": {
        "description": "ID of the posture deployment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "posture_id": {
        "description": "Relative name of the posture which needs to be deployed. It should be in the format:\n  organizations/{organization_id}/locations/{location}/postures/{posture_id}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "posture_revision_id": {
        "description": "Revision_id the posture which needs to be deployed.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "If set, there are currently changes in flight to the posture deployment.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "State of the posture deployment. A posture deployment can be in the following terminal states:\nACTIVE, CREATE_FAILED, UPDATE_FAILED, DELETE_FAILED.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_resource": {
        "description": "The resource on which the posture should be deployed. This can be in one of the following formats:\nprojects/{project_number},\nfolders/{folder_number},\norganizations/{organization_id}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the posture deployment was updated in UTC.",
        "description_kind": "plain",
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

func GoogleSecurityposturePostureDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSecurityposturePostureDeployment), &result)
	return &result
}
