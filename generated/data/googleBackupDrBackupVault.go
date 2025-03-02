package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBackupDrBackupVault = `{
  "block": {
    "attributes": {
      "access_restriction": {
        "computed": true,
        "description": "Access restriction for the backup vault. Default value is 'WITHIN_ORGANIZATION' if not provided during creation. Default value: \"WITHIN_ORGANIZATION\" Possible values: [\"ACCESS_RESTRICTION_UNSPECIFIED\", \"WITHIN_PROJECT\", \"WITHIN_ORGANIZATION\", \"UNRESTRICTED\", \"WITHIN_ORG_BUT_UNRESTRICTED_FOR_BA\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "allow_missing": {
        "computed": true,
        "description": "Allow idempotent deletion of backup vault. The request will still succeed in case the backup vault does not exist.",
        "description_kind": "plain",
        "type": "bool"
      },
      "annotations": {
        "computed": true,
        "description": "Optional. User annotations. See https://google.aip.dev/128#annotations\nStores small amounts of arbitrary data. \n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "backup_count": {
        "computed": true,
        "description": "Output only. The number of backups in this backup vault.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_minimum_enforced_retention_duration": {
        "computed": true,
        "description": "Required. The default and minimum enforced retention for each backup within the backup vault. The enforced retention for each backup can be extended.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_vault_id": {
        "description": "Required. ID of the requesting object.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The time when the instance was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletable": {
        "computed": true,
        "description": "Output only. Set to true when there are no backups nested under this resource.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "Optional. The description of the BackupVault instance (2048 characters or less).",
        "description_kind": "plain",
        "type": "string"
      },
      "effective_annotations": {
        "computed": true,
        "description": "All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
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
      "effective_time": {
        "computed": true,
        "description": "Optional. Time after which the BackupVault resource is locked.",
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Optional. Server specified ETag for the backup vault resource to prevent simultaneous updates from overwiting each other.",
        "description_kind": "plain",
        "type": "string"
      },
      "force_delete": {
        "computed": true,
        "description": "If set, the following restrictions against deletion of the backup vault instance can be overridden:\n   * deletion of a backup vault instance containing no backups, but still containing empty datasources.\n   * deletion of a backup vault instance that is being referenced by an active backup plan.",
        "description_kind": "plain",
        "type": "bool"
      },
      "force_update": {
        "computed": true,
        "description": "If set, allow update to extend the minimum enforced retention for backup vault. This overrides\n the restriction against conflicting retention periods. This conflict may occur when the\n expiration schedule defined by the associated backup plan is shorter than the minimum\n retention set by the backup vault.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ignore_backup_plan_references": {
        "computed": true,
        "description": "If set, the following restrictions against deletion of the backup vault instance can be overridden:\n   * deletion of a backup vault instance that is being referenced by an active backup plan.",
        "description_kind": "plain",
        "type": "bool"
      },
      "ignore_inactive_datasources": {
        "computed": true,
        "description": "If set, the following restrictions against deletion of the backup vault instance can be overridden:\n   * deletion of a backup vault instance containing no backups, but still containing empty datasources.",
        "description_kind": "plain",
        "type": "bool"
      },
      "labels": {
        "computed": true,
        "description": "Optional. Resource labels to represent user provided metadata. \n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The GCP location for the backup vault.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Output only. Identifier. The resource name.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account": {
        "computed": true,
        "description": "Output only. Service account used by the BackupVault Service for this BackupVault.  The user should grant this account permissions in their workload project to enable the service to run backups and restores there.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. The BackupVault resource instance state. \n Possible values:\n STATE_UNSPECIFIED\n CREATING\n ACTIVE\n DELETING\n ERROR",
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
      "total_stored_bytes": {
        "computed": true,
        "description": "Output only. Total size of the storage used by all backup resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. Output only Immutable after resource creation until resource deletion.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time when the instance was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBackupDrBackupVaultSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBackupDrBackupVault), &result)
	return &result
}
