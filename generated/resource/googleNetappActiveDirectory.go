package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetappActiveDirectory = `{
  "block": {
    "attributes": {
      "administrators": {
        "description": "Domain user accounts to be added to the local Administrators group of the SMB service. Comma-separated list of domain users or groups. The Domain Admin group is automatically added when the service joins your domain as a hidden group.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "aes_encryption": {
        "description": "Enables AES-128 and AES-256 encryption for Kerberos-based communication with Active Directory.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "backup_operators": {
        "description": "Domain user/group accounts to be added to the Backup Operators group of the SMB service. The Backup Operators group allows members to backup and restore files regardless of whether they have read or write access to the files. Comma-separated list.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Create time of the active directory. A timestamp in RFC3339 UTC \"Zulu\" format. Examples: \"2023-06-22T09:13:01.617Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns": {
        "description": "Comma separated list of DNS server IP addresses for the Active Directory domain.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain": {
        "description": "Fully qualified domain name for the Active Directory domain.",
        "description_kind": "plain",
        "required": true,
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
      "encrypt_dc_connections": {
        "description": "If enabled, traffic between the SMB server to Domain Controller (DC) will be encrypted.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kdc_hostname": {
        "description": "Hostname of the Active Directory server used as Kerberos Key Distribution Center. Only requried for volumes using kerberized NFSv4.1",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kdc_ip": {
        "description": "IP address of the Active Directory server used as Kerberos Key Distribution Center.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels as key value pairs. Example: '{ \"owner\": \"Bob\", \"department\": \"finance\", \"purpose\": \"testing\" }'.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "ldap_signing": {
        "description": "Specifies whether or not the LDAP traffic needs to be signed.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "description": "Name of the region for the policy to apply to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the Active Directory pool. Needs to be unique per location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "net_bios_prefix": {
        "description": "NetBIOS name prefix of the server to be created.\nA five-character random ID is generated automatically, for example, -6f9a, and appended to the prefix. The full UNC share path will have the following format:\n'\\\\NetBIOS_PREFIX-ABCD.DOMAIN_NAME\\SHARE_NAME'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nfs_users_with_ldap": {
        "description": "Local UNIX users on clients without valid user information in Active Directory are blocked from access to LDAP enabled volumes.\nThis option can be used to temporarily switch such volumes to AUTH_SYS authentication (user ID + 1-16 groups).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "organizational_unit": {
        "computed": true,
        "description": "Name of the Organizational Unit where you intend to create the computer account for NetApp Volumes.\nDefaults to 'CN=Computers' if left empty.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "password": {
        "description": "Password for specified username. Note - Manual changes done to the password will not be detected. Terraform will not re-apply the password, unless you use a new password in Terraform.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_operators": {
        "description": "Domain accounts that require elevated privileges such as 'SeSecurityPrivilege' to manage security logs. Comma-separated list.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "site": {
        "description": "Specifies an Active Directory site to manage domain controller selection.\nUse when Active Directory domain controllers in multiple regions are configured. Defaults to 'Default-First-Site-Name' if left empty.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the Active Directory policy (not the Active Directory itself).",
        "description_kind": "plain",
        "type": "string"
      },
      "state_details": {
        "computed": true,
        "description": "The state details of the Active Directory.",
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
      "username": {
        "description": "Username for the Active Directory account with permissions to create the compute account within the specified organizational unit.",
        "description_kind": "plain",
        "required": true,
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

func GoogleNetappActiveDirectorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetappActiveDirectory), &result)
	return &result
}
