package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOracleDatabaseCloudExadataInfrastructure = `{
  "block": {
    "attributes": {
      "cloud_exadata_infrastructure_id": {
        "description": "The ID of the Exadata Infrastructure to create. This value is restricted\nto (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$) and must be a maximum of 63\ncharacters in length. The value must start with a letter and end with\na letter or a number.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The date and time that the Exadata Infrastructure was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Whether or not to allow Terraform to destroy the instance. Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.",
        "description_kind": "plain",
        "type": "bool"
      },
      "display_name": {
        "computed": true,
        "description": "User friendly name for this resource.",
        "description_kind": "plain",
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
      "entitlement_id": {
        "computed": true,
        "description": "Entitlement ID of the private offer against which this infrastructure\nresource is provisioned.",
        "description_kind": "plain",
        "type": "string"
      },
      "gcp_oracle_zone": {
        "computed": true,
        "description": "GCP location where Oracle Exadata is hosted.",
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
        "computed": true,
        "description": "Labels or tags associated with the resource. \n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. See documentation for resource type 'oracledatabase.googleapis.com/DbServer'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The name of the Exadata Infrastructure resource with the following format:\nprojects/{project}/locations/{region}/cloudExadataInfrastructures/{cloud_exadata_infrastructure}",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "properties": {
        "computed": true,
        "description": "Various properties of Exadata Infrastructure.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "activated_storage_count": "number",
              "additional_storage_count": "number",
              "available_storage_size_gb": "number",
              "compute_count": "number",
              "cpu_count": "number",
              "customer_contacts": [
                "list",
                [
                  "object",
                  {
                    "email": "string"
                  }
                ]
              ],
              "data_storage_size_tb": "number",
              "db_node_storage_size_gb": "number",
              "db_server_version": "string",
              "maintenance_window": [
                "list",
                [
                  "object",
                  {
                    "custom_action_timeout_mins": "number",
                    "days_of_week": [
                      "list",
                      "string"
                    ],
                    "hours_of_day": [
                      "list",
                      "number"
                    ],
                    "is_custom_action_timeout_enabled": "bool",
                    "lead_time_week": "number",
                    "months": [
                      "list",
                      "string"
                    ],
                    "patching_mode": "string",
                    "preference": "string",
                    "weeks_of_month": [
                      "list",
                      "number"
                    ]
                  }
                ]
              ],
              "max_cpu_count": "number",
              "max_data_storage_tb": "number",
              "max_db_node_storage_size_gb": "number",
              "max_memory_gb": "number",
              "memory_size_gb": "number",
              "monthly_db_server_version": "string",
              "monthly_storage_server_version": "string",
              "next_maintenance_run_id": "string",
              "next_maintenance_run_time": "string",
              "next_security_maintenance_run_time": "string",
              "oci_url": "string",
              "ocid": "string",
              "shape": "string",
              "state": "string",
              "storage_count": "number",
              "storage_server_version": "string",
              "total_storage_size_gb": "number"
            }
          ]
        ]
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleOracleDatabaseCloudExadataInfrastructureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOracleDatabaseCloudExadataInfrastructure), &result)
	return &result
}
