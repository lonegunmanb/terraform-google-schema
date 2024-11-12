package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOracleDatabaseCloudVmCluster = `{
  "block": {
    "attributes": {
      "backup_subnet_cidr": {
        "description": "CIDR range of the backup subnet.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cidr": {
        "description": "Network settings. CIDR to use for cluster IP allocation.",
        "description_kind": "plain",
        "required": true,
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
      "display_name": {
        "description": "User friendly name for this resource.",
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
      "exadata_infrastructure": {
        "description": "The name of the Exadata Infrastructure resource on which VM cluster\nresource is created, in the following format:\nprojects/{project}/locations/{region}/cloudExadataInfrastuctures/{cloud_extradata_infrastructure}",
        "description_kind": "plain",
        "required": true,
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
        "description": "Labels or tags associated with the VM Cluster. \n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The name of the VPC network.\nFormat: projects/{project}/global/networks/{network}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      }
    },
    "block_types": {
      "properties": {
        "block": {
          "attributes": {
            "cluster_name": {
              "computed": true,
              "description": "OCI Cluster name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "compartment_id": {
              "computed": true,
              "description": "Compartment ID of cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "cpu_core_count": {
              "description": "Number of enabled CPU cores.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "data_storage_size_tb": {
              "computed": true,
              "description": "The data disk group size to be allocated in TBs.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "db_node_storage_size_gb": {
              "computed": true,
              "description": "Local storage per VM",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "db_server_ocids": {
              "computed": true,
              "description": "OCID of database servers.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "disk_redundancy": {
              "computed": true,
              "description": "The type of redundancy. \n Possible values:\n DISK_REDUNDANCY_UNSPECIFIED\nHIGH\nNORMAL",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dns_listener_ip": {
              "computed": true,
              "description": "DNS listener IP.",
              "description_kind": "plain",
              "type": "string"
            },
            "domain": {
              "computed": true,
              "description": "Parent DNS domain where SCAN DNS and hosts names are qualified.\nex: ocispdelegated.ocisp10jvnet.oraclevcn.com",
              "description_kind": "plain",
              "type": "string"
            },
            "gi_version": {
              "description": "Grid Infrastructure Version.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hostname": {
              "computed": true,
              "description": "host name without domain.\nformat: \"-\" with some suffix.\nex: sp2-yi0xq where \"sp2\" is the hostname_prefix.",
              "description_kind": "plain",
              "type": "string"
            },
            "hostname_prefix": {
              "description": "Prefix for VM cluster host names.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "license_type": {
              "description": "License type of VM Cluster. \n Possible values:\n LICENSE_TYPE_UNSPECIFIED\nLICENSE_INCLUDED\nBRING_YOUR_OWN_LICENSE",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "local_backup_enabled": {
              "description": "Use local backup.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "memory_size_gb": {
              "computed": true,
              "description": "Memory allocated in GBs.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "node_count": {
              "computed": true,
              "description": "Number of database servers.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "oci_url": {
              "computed": true,
              "description": "Deep link to the OCI console to view this resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "ocid": {
              "computed": true,
              "description": "Oracle Cloud Infrastructure ID of VM Cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "ocpu_count": {
              "computed": true,
              "description": "OCPU count per VM. Minimum is 0.1.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "scan_dns": {
              "computed": true,
              "description": "SCAN DNS name.\nex: sp2-yi0xq-scan.ocispdelegated.ocisp10jvnet.oraclevcn.com",
              "description_kind": "plain",
              "type": "string"
            },
            "scan_dns_record_id": {
              "computed": true,
              "description": "OCID of scan DNS record.",
              "description_kind": "plain",
              "type": "string"
            },
            "scan_ip_ids": {
              "computed": true,
              "description": "OCIDs of scan IPs.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "scan_listener_port_tcp": {
              "computed": true,
              "description": "SCAN listener port - TCP",
              "description_kind": "plain",
              "type": "number"
            },
            "scan_listener_port_tcp_ssl": {
              "computed": true,
              "description": "SCAN listener port - TLS",
              "description_kind": "plain",
              "type": "number"
            },
            "shape": {
              "computed": true,
              "description": "Shape of VM Cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "sparse_diskgroup_enabled": {
              "computed": true,
              "description": "Use exadata sparse snapshots.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ssh_public_keys": {
              "description": "SSH public keys to be stored with cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "state": {
              "computed": true,
              "description": "State of the cluster. \n Possible values:\n STATE_UNSPECIFIED\nPROVISIONING\nAVAILABLE\nUPDATING\nTERMINATING\nTERMINATED\nFAILED\nMAINTENANCE_IN_PROGRESS",
              "description_kind": "plain",
              "type": "string"
            },
            "storage_size_gb": {
              "computed": true,
              "description": "The storage allocation for the disk group, in gigabytes (GB).",
              "description_kind": "plain",
              "type": "number"
            },
            "system_version": {
              "computed": true,
              "description": "Operating system version of the image.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "diagnostics_data_collection_options": {
              "block": {
                "attributes": {
                  "diagnostics_events_enabled": {
                    "description": "Indicates whether diagnostic collection is enabled for the VM cluster",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "health_monitoring_enabled": {
                    "description": "Indicates whether health monitoring is enabled for the VM cluster",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "incident_logs_enabled": {
                    "description": "Indicates whether incident logs and trace collection are enabled for the VM\ncluster",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Data collection options for diagnostics.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "time_zone": {
              "block": {
                "attributes": {
                  "id": {
                    "computed": true,
                    "description": "IANA Time Zone Database time zone, e.g. \"America/New_York\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Represents a time zone from the\n[IANA Time Zone Database](https://www.iana.org/time-zones).",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Various properties and settings associated with Exadata VM cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
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
