package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSqlSourceRepresentationInstance = `{
  "block": {
    "attributes": {
      "ca_certificate": {
        "description": "The CA certificate on the external server. Include only if SSL/TLS is used on the external server.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_certificate": {
        "description": "The client certificate on the external server. Required only for server-client authentication. Include only if SSL/TLS is used on the external server.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_key": {
        "description": "The private key file for the client certificate on the external server. Required only for server-client authentication. Include only if SSL/TLS is used on the external server.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_version": {
        "description": "The MySQL, PostgreSQL or SQL Server (beta) version to use. Supported values include MYSQL_5_6, MYSQL_5_7, MYSQL_8_0, MYSQL_8_4, POSTGRES_9_6, POSTGRES_10, POSTGRES_11, POSTGRES_12, POSTGRES_13, POSTGRES_14, POSTGRES_15, POSTGRES_16, POSTGRES_17. Database Version Policies includes an up-to-date reference of supported versions.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dump_file_path": {
        "description": "A file in the bucket that contains the data from the external server.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host": {
        "description": "The IPv4 address and port for the external server, or the the DNS address for the external server. If the external server is hosted on Cloud SQL, the port is 5432.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the source representation instance. Use any valid Cloud SQL instance name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "password": {
        "description": "The password for the replication user account.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "port": {
        "description": "The externally accessible port for the source database server.\nDefaults to 3306.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The Region in which the created instance should reside.\nIf it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "username": {
        "description": "The replication user account on the external server.",
        "description_kind": "plain",
        "optional": true,
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

func GoogleSqlSourceRepresentationInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSqlSourceRepresentationInstance), &result)
	return &result
}
