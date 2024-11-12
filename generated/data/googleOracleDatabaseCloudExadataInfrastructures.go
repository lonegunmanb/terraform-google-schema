package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOracleDatabaseCloudExadataInfrastructures = `{
  "block": {
    "attributes": {
      "cloud_exadata_infrastructures": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cloud_exadata_infrastructure_id": "string",
              "create_time": "string",
              "display_name": "string",
              "effective_labels": [
                "map",
                "string"
              ],
              "entitlement_id": "string",
              "gcp_oracle_zone": "string",
              "labels": [
                "map",
                "string"
              ],
              "location": "string",
              "name": "string",
              "project": "string",
              "properties": [
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
              ],
              "terraform_labels": [
                "map",
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
      "location": {
        "description": "location",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description": "The ID of the project in which the dataset is located. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleOracleDatabaseCloudExadataInfrastructuresSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOracleDatabaseCloudExadataInfrastructures), &result)
	return &result
}
