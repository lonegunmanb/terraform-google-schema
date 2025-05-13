package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSpannerDatabase = `{
  "block": {
    "attributes": {
      "database_dialect": {
        "computed": true,
        "description": "The dialect of the Cloud Spanner Database.\nIf it is not provided, \"GOOGLE_STANDARD_SQL\" will be used. Possible values: [\"GOOGLE_STANDARD_SQL\", \"POSTGRESQL\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "ddl": {
        "computed": true,
        "description": "An optional list of DDL statements to run inside the database. Statements can create\ntables, indexes, etc.\n\nDuring creation these statements execute atomically with the creation of the database\nand if there is an error in any statement, the database is not created.\n\nTerraform does not perform drift detection on this field and assumes that the values\nrecorded in state are accurate. Limited updates to this field are supported, and\nnewly appended DDL statements can be executed in an update. However, modifications\nto prior statements will create a plan that marks the resource for recreation.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "default_time_zone": {
        "computed": true,
        "description": "The default time zone for the database. The default time zone must be a valid name\nfrom the tz database. Default value is \"America/Los_angeles\".",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Whether Terraform will be prevented from destroying the database. Defaults to true.\nWhen a'terraform destroy' or 'terraform apply' would delete the database,\nthe command will fail if this field is not set to false in Terraform state.\nWhen the field is set to true or unset in Terraform state, a 'terraform apply'\nor 'terraform destroy' that would delete the database will fail.\nWhen the field is set to false, deleting the database is allowed.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_drop_protection": {
        "computed": true,
        "description": "Whether drop protection is enabled for this database. Defaults to false.\nDrop protection is different from\nthe \"deletion_protection\" attribute in the following ways:\n(1) \"deletion_protection\" only protects the database from deletions in Terraform.\nwhereas setting “enableDropProtection” to true protects the database from deletions in all interfaces.\n(2) Setting \"enableDropProtection\" to true also prevents the deletion of the parent instance containing the database.\n\"deletion_protection\" attribute does not provide protection against the deletion of the parent instance.",
        "description_kind": "plain",
        "type": "bool"
      },
      "encryption_config": {
        "computed": true,
        "description": "Encryption configuration for the database",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key_name": "string",
              "kms_key_names": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance": {
        "description": "The instance to create the database on.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "A unique identifier for the database, which cannot be changed after the\ninstance is created. Values are of the form '[a-z][-_a-z0-9]*[a-z0-9]'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "An explanation of the status of the database.",
        "description_kind": "plain",
        "type": "string"
      },
      "version_retention_period": {
        "computed": true,
        "description": "The retention period for the database. The retention period must be between 1 hour\nand 7 days, and can be specified in days, hours, minutes, or seconds. For example,\nthe values 1d, 24h, 1440m, and 86400s are equivalent. Default value is 1h.\nIf this property is used, you must avoid adding new DDL statements to 'ddl' that\nupdate the database's version_retention_period.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleSpannerDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSpannerDatabase), &result)
	return &result
}
