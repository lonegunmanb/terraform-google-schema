package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOracleDatabaseCloudVmCluster = `{
  "block": {
    "attributes": {
      "backup_odb_subnet": {
        "computed": true,
        "description": "The name of the backup OdbSubnet associated with the VM Cluster.\nFormat:\nprojects/{project}/locations/{location}/odbNetworks/{odb_network}/odbSubnets/{odb_subnet}",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_subnet_cidr": {
        "computed": true,
        "description": "CIDR range of the backup subnet.",
        "description_kind": "plain",
        "type": "string"
      },
      "cidr": {
        "computed": true,
        "description": "Network settings. CIDR to use for cluster IP allocation.",
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_vm_cluster_id": {
        "description": "The ID of the VM Cluster to create. This value is restricted\nto (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$) and must be a maximum of 63\ncharacters in length. The value must start with a letter and end with\na letter or a number.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The date and time that the VM cluster was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Whether Terraform will be prevented from destroying the cluster. Deleting this cluster via terraform destroy or terraform apply will only succeed if this field is false in the Terraform state.",
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
      "exadata_infrastructure": {
        "computed": true,
        "description": "The name of the Exadata Infrastructure resource on which VM cluster\nresource is created, in the following format:\nprojects/{project}/locations/{region}/cloudExadataInfrastuctures/{cloud_extradata_infrastructure}",
        "description_kind": "plain",
        "type": "string"
      },
      "gcp_oracle_zone": {
        "computed": true,
        "description": "GCP location where Oracle Exadata is hosted. It is same as GCP Oracle zone\nof Exadata infrastructure.",
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
        "description": "Labels or tags associated with the VM Cluster. \n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. See documentation for resource type 'oracledatabase.googleapis.com/DbNode'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The name of the VM Cluster resource with the format:\nprojects/{project}/locations/{region}/cloudVmClusters/{cloud_vm_cluster}",
        "description_kind": "plain",
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "The name of the VPC network.\nFormat: projects/{project}/global/networks/{network}",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network": {
        "computed": true,
        "description": "The name of the OdbNetwork associated with the VM Cluster.\nFormat:\nprojects/{project}/locations/{location}/odbNetworks/{odb_network}\nIt is optional but if specified, this should match the parent ODBNetwork of\nthe odb_subnet and backup_odb_subnet.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_subnet": {
        "computed": true,
        "description": "The name of the OdbSubnet associated with the VM Cluster for\nIP allocation. Format:\nprojects/{project}/locations/{location}/odbNetworks/{odb_network}/odbSubnets/{odb_subnet}",
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
        "description": "Various properties and settings associated with Exadata VM cluster.",
        "description_kind": "plain",
        "type": [
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

func GoogleOracleDatabaseCloudVmClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOracleDatabaseCloudVmCluster), &result)
	return &result
}
