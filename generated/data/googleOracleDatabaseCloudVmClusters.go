package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOracleDatabaseCloudVmClusters = `{
  "block": {
    "attributes": {
      "cloud_vm_clusters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_subnet_cidr": "string",
              "cidr": "string",
              "cloud_vm_cluster_id": "string",
              "create_time": "string",
              "display_name": "string",
              "effective_labels": [
                "map",
                "string"
              ],
              "exadata_infrastructure": "string",
              "gcp_oracle_zone": "string",
              "labels": [
                "map",
                "string"
              ],
              "location": "string",
              "name": "string",
              "network": "string",
              "project": "string",
              "properties": [
                "list",
                [
                  "object",
                  {
                    "cluster_name": "string",
                    "compartment_id": "string",
                    "cpu_core_count": "number",
                    "data_storage_size_tb": "number",
                    "db_node_storage_size_gb": "number",
                    "db_server_ocids": [
                      "list",
                      "string"
                    ],
                    "diagnostics_data_collection_options": [
                      "list",
                      [
                        "object",
                        {
                          "diagnostics_events_enabled": "bool",
                          "health_monitoring_enabled": "bool",
                          "incident_logs_enabled": "bool"
                        }
                      ]
                    ],
                    "disk_redundancy": "string",
                    "dns_listener_ip": "string",
                    "domain": "string",
                    "gi_version": "string",
                    "hostname": "string",
                    "hostname_prefix": "string",
                    "license_type": "string",
                    "local_backup_enabled": "bool",
                    "memory_size_gb": "number",
                    "node_count": "number",
                    "oci_url": "string",
                    "ocid": "string",
                    "ocpu_count": "number",
                    "scan_dns": "string",
                    "scan_dns_record_id": "string",
                    "scan_ip_ids": [
                      "list",
                      "string"
                    ],
                    "scan_listener_port_tcp": "number",
                    "scan_listener_port_tcp_ssl": "number",
                    "shape": "string",
                    "sparse_diskgroup_enabled": "bool",
                    "ssh_public_keys": [
                      "list",
                      "string"
                    ],
                    "state": "string",
                    "storage_size_gb": "number",
                    "system_version": "string",
                    "time_zone": [
                      "list",
                      [
                        "object",
                        {
                          "id": "string"
                        }
                      ]
                    ]
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

func GoogleOracleDatabaseCloudVmClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOracleDatabaseCloudVmClusters), &result)
	return &result
}
