package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAlloydbInstance = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "availability_type": {
        "computed": true,
        "description": "'Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.\nNote that primary and read instances can have different availability types.\nPrimary instances can be either ZONAL or REGIONAL. Read Pool instances can also be either ZONAL or REGIONAL.\nRead pools of size 1 can only have zonal availability. Read pools with a node count of 2 or more\ncan have regional availability (nodes are present in 2 or more zones in a region).\nPossible values are: 'AVAILABILITY_TYPE_UNSPECIFIED', 'ZONAL', 'REGIONAL'.' Possible values: [\"AVAILABILITY_TYPE_UNSPECIFIED\", \"ZONAL\", \"REGIONAL\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster": {
        "description": "Identifies the alloydb cluster. Must be in the format\n'projects/{project}/locations/{location}/clusters/{cluster_id}'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Time the Instance was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_flags": {
        "computed": true,
        "description": "Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "display_name": {
        "description": "User-settable and human-readable display name for the Instance.",
        "description_kind": "plain",
        "optional": true,
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
      "gce_zone": {
        "description": "The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "description": "The ID of the alloydb instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_type": {
        "description": "The type of the instance.\nIf the instance type is READ_POOL, provide the associated PRIMARY/SECONDARY instance in the 'depends_on' meta-data attribute.\nIf the instance type is SECONDARY, point to the cluster_type of the associated secondary cluster instead of mentioning SECONDARY.\nExample: {instance_type = google_alloydb_cluster.\u003csecondary_cluster_name\u003e.cluster_type} instead of {instance_type = SECONDARY}\nIf the instance type is SECONDARY, the terraform delete instance operation does not delete the secondary instance but abandons it instead.\nUse deletion_policy = \"FORCE\" in the associated secondary cluster and delete the cluster forcefully to delete the secondary cluster as well its associated secondary instance.\nUsers can undo the delete secondary instance action by importing the deleted secondary instance by calling terraform import. Possible values: [\"PRIMARY\", \"READ_POOL\", \"SECONDARY\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_address": {
        "computed": true,
        "description": "The IP address for the Instance. This is the connection endpoint for an end-user application.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "User-defined labels for the alloydb instance.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "The name of the instance resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "outbound_public_ip_addresses": {
        "computed": true,
        "description": "The outbound public IP addresses for the instance. This is available ONLY when\nnetworkConfig.enableOutboundPublicIp is set to true. These IP addresses are used\nfor outbound connections.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "public_ip_address": {
        "computed": true,
        "description": "The public IP addresses for the Instance. This is available ONLY when\nnetworkConfig.enablePublicIp is set to true. This is the connection\nendpoint for an end-user application.",
        "description_kind": "plain",
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "Set to true if the current state of Instance does not match the user's intended state, and the service is actively updating the resource to reconcile them. This can happen due to user-triggered updates or system actions like failover or maintenance.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The current state of the alloydb instance.",
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
      "uid": {
        "computed": true,
        "description": "The system-generated UID of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the Instance was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "client_connection_config": {
        "block": {
          "attributes": {
            "require_connectors": {
              "description": "Configuration to enforce connectors only (ex: AuthProxy) connections to the database.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "ssl_config": {
              "block": {
                "attributes": {
                  "ssl_mode": {
                    "computed": true,
                    "description": "SSL mode. Specifies client-server SSL/TLS connection behavior. Possible values: [\"ENCRYPTED_ONLY\", \"ALLOW_UNENCRYPTED_AND_ENCRYPTED\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "SSL config option for this instance.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Client connection specific configurations.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "machine_config": {
        "block": {
          "attributes": {
            "cpu_count": {
              "computed": true,
              "description": "The number of CPU's in the VM instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Configurations for the machines that host the underlying database engine.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_config": {
        "block": {
          "attributes": {
            "enable_outbound_public_ip": {
              "description": "Enabling outbound public ip for the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_public_ip": {
              "description": "Enabling public ip for the instance. If a user wishes to disable this,\nplease also clear the list of the authorized external networks set on\nthe same instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "authorized_external_networks": {
              "block": {
                "attributes": {
                  "cidr_range": {
                    "description": "CIDR range for one authorized network of the instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A list of external networks authorized to access this instance. This\nfield is only allowed to be set when 'enable_public_ip' is set to\ntrue.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Instance level network configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "psc_instance_config": {
        "block": {
          "attributes": {
            "allowed_consumer_projects": {
              "description": "List of consumer projects that are allowed to create PSC endpoints to service-attachments to this instance.\nThese should be specified as project numbers only.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "psc_dns_name": {
              "computed": true,
              "description": "The DNS name of the instance for PSC connectivity.\nName convention: \u003cuid\u003e.\u003cuid\u003e.\u003cregion\u003e.alloydb-psc.goog",
              "description_kind": "plain",
              "type": "string"
            },
            "service_attachment_link": {
              "computed": true,
              "description": "The service attachment created when Private Service Connect (PSC) is enabled for the instance.\nThe name of the resource will be in the format of\n'projects/\u003calloydb-tenant-project-number\u003e/regions/\u003cregion-name\u003e/serviceAttachments/\u003cservice-attachment-name\u003e'",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Configuration for Private Service Connect (PSC) for the instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "query_insights_config": {
        "block": {
          "attributes": {
            "query_plans_per_minute": {
              "description": "Number of query execution plans captured by Insights per minute for all queries combined. The default value is 5. Any integer between 0 and 20 is considered valid.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "query_string_length": {
              "description": "Query string length. The default value is 1024. Any integer between 256 and 4500 is considered valid.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "record_application_tags": {
              "description": "Record application tags for an instance. This flag is turned \"on\" by default.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "record_client_address": {
              "description": "Record client address for an instance. Client address is PII information. This flag is turned \"on\" by default.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Configuration for query insights.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "read_pool_config": {
        "block": {
          "attributes": {
            "node_count": {
              "description": "Read capacity, i.e. number of nodes in a read pool instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Read pool specific config. If the instance type is READ_POOL, this configuration must be provided.",
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

func GoogleAlloydbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAlloydbInstance), &result)
	return &result
}
