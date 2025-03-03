package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAlloydbInstance = `{
  "block": {
    "attributes": {
      "annotations": {
        "computed": true,
        "description": "Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "availability_type": {
        "computed": true,
        "description": "'Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.\nNote that primary and read instances can have different availability types.\nPrimary instances can be either ZONAL or REGIONAL. Read Pool instances can also be either ZONAL or REGIONAL.\nRead pools of size 1 can only have zonal availability. Read pools with a node count of 2 or more\ncan have regional availability (nodes are present in 2 or more zones in a region).\nPossible values are: 'AVAILABILITY_TYPE_UNSPECIFIED', 'ZONAL', 'REGIONAL'.' Possible values: [\"AVAILABILITY_TYPE_UNSPECIFIED\", \"ZONAL\", \"REGIONAL\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "client_connection_config": {
        "computed": true,
        "description": "Client connection specific configurations.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "require_connectors": "bool",
              "ssl_config": [
                "list",
                [
                  "object",
                  {
                    "ssl_mode": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "cluster": {
        "computed": true,
        "description": "Identifies the alloydb cluster. Must be in the format\n'projects/{project}/locations/{location}/clusters/{cluster_id}'",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_id": {
        "description": "The ID of the alloydb cluster that the instance belongs to.'alloydb_cluster_id'",
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
        "type": [
          "map",
          "string"
        ]
      },
      "display_name": {
        "computed": true,
        "description": "User-settable and human-readable display name for the Instance.",
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
      "gce_zone": {
        "computed": true,
        "description": "The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "The type of the instance.\nIf the instance type is READ_POOL, provide the associated PRIMARY/SECONDARY instance in the 'depends_on' meta-data attribute.\nIf the instance type is SECONDARY, point to the cluster_type of the associated secondary cluster instead of mentioning SECONDARY.\nExample: {instance_type = google_alloydb_cluster.\u003csecondary_cluster_name\u003e.cluster_type} instead of {instance_type = SECONDARY}\nIf the instance type is SECONDARY, the terraform delete instance operation does not delete the secondary instance but abandons it instead.\nUse deletion_policy = \"FORCE\" in the associated secondary cluster and delete the cluster forcefully to delete the secondary cluster as well its associated secondary instance.\nUsers can undo the delete secondary instance action by importing the deleted secondary instance by calling terraform import. Possible values: [\"PRIMARY\", \"READ_POOL\", \"SECONDARY\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_address": {
        "computed": true,
        "description": "The IP address for the Instance. This is the connection endpoint for an end-user application.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "User-defined labels for the alloydb instance.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The canonical ID for the location. For example: \"us-east1\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "machine_config": {
        "computed": true,
        "description": "Configurations for the machines that host the underlying database engine.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cpu_count": "number"
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description": "The name of the instance resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_config": {
        "computed": true,
        "description": "Instance level network configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "authorized_external_networks": [
                "list",
                [
                  "object",
                  {
                    "cidr_range": "string"
                  }
                ]
              ],
              "enable_outbound_public_ip": "bool",
              "enable_public_ip": "bool"
            }
          ]
        ]
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
      "project": {
        "description": "Project ID of the project.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "psc_instance_config": {
        "computed": true,
        "description": "Configuration for Private Service Connect (PSC) for the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allowed_consumer_projects": [
                "list",
                "string"
              ],
              "psc_dns_name": "string",
              "service_attachment_link": "string"
            }
          ]
        ]
      },
      "public_ip_address": {
        "computed": true,
        "description": "The public IP addresses for the Instance. This is available ONLY when\nnetworkConfig.enablePublicIp is set to true. This is the connection\nendpoint for an end-user application.",
        "description_kind": "plain",
        "type": "string"
      },
      "query_insights_config": {
        "computed": true,
        "description": "Configuration for query insights.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "query_plans_per_minute": "number",
              "query_string_length": "number",
              "record_application_tags": "bool",
              "record_client_address": "bool"
            }
          ]
        ]
      },
      "read_pool_config": {
        "computed": true,
        "description": "Read pool specific config. If the instance type is READ_POOL, this configuration must be provided.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "node_count": "number"
            }
          ]
        ]
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
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleAlloydbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAlloydbInstance), &result)
	return &result
}
