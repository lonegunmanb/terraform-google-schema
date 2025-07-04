package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeInterconnect = `{
  "block": {
    "attributes": {
      "admin_enabled": {
        "description": "Administrative status of the interconnect. When this is set to true, the Interconnect is\nfunctional and can carry traffic. When set to false, no packets can be carried over the\ninterconnect and no BGP routes are exchanged over it. By default, the status is set to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "available_features": {
        "computed": true,
        "description": "interconnects.list of features available for this Interconnect connection. Can take the value:\nMACSEC. If present then the Interconnect connection is provisioned on MACsec capable hardware\nports. If not present then the Interconnect connection is provisioned on non-MACsec capable\nports and MACsec isn't supported and enabling MACsec fails).",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "circuit_infos": {
        "computed": true,
        "description": "A list of CircuitInfo objects, that describe the individual circuits in this LAG.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "customer_demarc_id": "string",
              "google_circuit_id": "string",
              "google_demarc_id": "string"
            }
          ]
        ]
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_name": {
        "description": "Customer name, to put in the Letter of Authorization as the party authorized to request a\ncrossconnect. This field is required for Dedicated and Partner Interconnect, should not be specified\nfor cross-cloud interconnect.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when you create the resource.",
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
      "expected_outages": {
        "computed": true,
        "description": "A list of outages expected for this Interconnect.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "affected_circuits": [
                "list",
                "string"
              ],
              "description": "string",
              "end_time": "string",
              "issue_type": "string",
              "name": "string",
              "source": "string",
              "start_time": "string",
              "state": "string"
            }
          ]
        ]
      },
      "google_ip_address": {
        "computed": true,
        "description": "IP address configured on the Google side of the Interconnect link.\nThis can be used only for ping tests.",
        "description_kind": "plain",
        "type": "string"
      },
      "google_reference_id": {
        "computed": true,
        "description": "Google reference ID to be used when raising support tickets with Google or otherwise to debug\nbackend connectivity issues.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "interconnect_attachments": {
        "computed": true,
        "description": "A list of the URLs of all InterconnectAttachments configured to use this Interconnect.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "interconnect_groups": {
        "computed": true,
        "description": "URLs of InterconnectGroups that include this Interconnect.\nOrder is arbitrary and items are unique.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "interconnect_type": {
        "description": "Type of interconnect. Note that a value IT_PRIVATE has been deprecated in favor of DEDICATED.\nCan take one of the following values:\n  - PARTNER: A partner-managed interconnection shared between customers though a partner.\n  - DEDICATED: A dedicated physical interconnection with the customer. Possible values: [\"DEDICATED\", \"PARTNER\", \"IT_PRIVATE\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "label_fingerprint": {
        "computed": true,
        "description": "A fingerprint for the labels being applied to this Interconnect, which is essentially a hash\nof the labels set used for optimistic locking. The fingerprint is initially generated by\nCompute Engine and changes after every request to modify or update labels.\nYou must always provide an up-to-date fingerprint hash in order to update or change labels,\notherwise the request will fail with error 412 conditionNotMet.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Labels for this resource. These can only be added or modified by the setLabels\nmethod. Each label key/value pair must comply with RFC1035. Label values may be empty.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "link_type": {
        "description": "Type of link requested. Note that this field indicates the speed of each of the links in the\nbundle, not the speed of the entire bundle. Can take one of the following values:\n  - LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics.\n  - LINK_TYPE_ETHERNET_100G_LR: A 100G Ethernet with LR optics.\n  - LINK_TYPE_ETHERNET_400G_LR4: A 400G Ethernet with LR4 optics Possible values: [\"LINK_TYPE_ETHERNET_10G_LR\", \"LINK_TYPE_ETHERNET_100G_LR\", \"LINK_TYPE_ETHERNET_400G_LR4\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "URL of the InterconnectLocation object that represents where this connection is to be provisioned.\nSpecifies the location inside Google's Networks.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "macsec_enabled": {
        "description": "Enable or disable MACsec on this Interconnect connection.\nMACsec enablement fails if the MACsec object is not specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is created. The name must be\n1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters\nlong and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first\ncharacter must be a lowercase letter, and all following characters must be a dash,\nlowercase letter, or digit, except the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "noc_contact_email": {
        "description": "Email address to contact the customer NOC for operations and maintenance notifications\nregarding this Interconnect. If specified, this will be used for notifications in addition to\nall other forms described, such as Cloud Monitoring logs alerting and Cloud Notifications.\nThis field is required for users who sign up for Cloud Interconnect using workforce identity\nfederation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "operational_status": {
        "computed": true,
        "description": "The current status of this Interconnect's functionality, which can take one of the following:\n  - OS_ACTIVE: A valid Interconnect, which is turned up and is ready to use. Attachments may\n  be provisioned on this Interconnect.\n  - OS_UNPROVISIONED: An Interconnect that has not completed turnup. No attachments may be\n  provisioned on this Interconnect.\n  - OS_UNDER_MAINTENANCE: An Interconnect that is undergoing internal maintenance. No\n  attachments may be provisioned or updated on this Interconnect.",
        "description_kind": "plain",
        "type": "string"
      },
      "peer_ip_address": {
        "computed": true,
        "description": "IP address configured on the customer side of the Interconnect link.\nThe customer should configure this IP address during turnup when prompted by Google NOC.\nThis can be used only for ping tests.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provisioned_link_count": {
        "computed": true,
        "description": "Number of links actually provisioned in this interconnect.",
        "description_kind": "plain",
        "type": "number"
      },
      "remote_location": {
        "description": "Indicates that this is a Cross-Cloud Interconnect. This field specifies the location outside\nof Google's network that the interconnect is connected to.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "requested_features": {
        "description": "interconnects.list of features requested for this Interconnect connection. Options: IF_MACSEC (\nIf specified then the connection is created on MACsec capable hardware ports. If not\nspecified, the default value is false, which allocates non-MACsec capable ports first if\navailable). Note that MACSEC is still technically allowed for compatibility reasons, but it\ndoes not work with the API, and will be removed in an upcoming major version. Possible values: [\"MACSEC\", \"CROSS_SITE_NETWORK\", \"IF_MACSEC\"]",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "requested_link_count": {
        "description": "Target number of physical links in the link bundle, as requested by the customer.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "satisfies_pzs": {
        "computed": true,
        "description": "Reserved for future use.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The current state of Interconnect functionality, which can take one of the following values:\n  - ACTIVE: The Interconnect is valid, turned up and ready to use.\n  Attachments may be provisioned on this Interconnect.\n  - UNPROVISIONED: The Interconnect has not completed turnup. No attachments may b\n   provisioned on this Interconnect.\n  - UNDER_MAINTENANCE: The Interconnect is undergoing internal maintenance. No attachments may\n   be provisioned or updated on this Interconnect.",
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
      }
    },
    "block_types": {
      "macsec": {
        "block": {
          "attributes": {
            "fail_open": {
              "description": "If set to true, the Interconnect connection is configured with a should-secure\nMACsec security policy, that allows the Google router to fallback to cleartext\ntraffic if the MKA session cannot be established. By default, the Interconnect\nconnection is configured with a must-secure security policy that drops all traffic\nif the MKA session cannot be established with your router.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "pre_shared_keys": {
              "block": {
                "attributes": {
                  "fail_open": {
                    "deprecated": true,
                    "description": "If set to true, the Interconnect connection is configured with a should-secure\nMACsec security policy, that allows the Google router to fallback to cleartext\ntraffic if the MKA session cannot be established. By default, the Interconnect\nconnection is configured with a must-secure security policy that drops all traffic\nif the MKA session cannot be established with your router.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "name": {
                    "description": "A name for this pre-shared key. The name must be 1-63 characters long, and\n comply with RFC1035. Specifically, the name must be 1-63 characters long and match\n the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character\n must be a lowercase letter, and all following characters must be a dash, lowercase\n letter, or digit, except the last character, which cannot be a dash.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "start_time": {
                    "description": "A RFC3339 timestamp on or after which the key is valid. startTime can be in the\nfuture. If the keychain has a single key, startTime can be omitted. If the keychain\nhas multiple keys, startTime is mandatory for each key. The start times of keys must\nbe in increasing order. The start times of two consecutive keys must be at least 6\nhours apart.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A keychain placeholder describing a set of named key objects along with their\nstart times. A MACsec CKN/CAK is generated for each key in the key chain.\nGoogle router automatically picks the key with the most recent startTime when establishing\nor re-establishing a MACsec secure link.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration that enables Media Access Control security (MACsec) on the Cloud\nInterconnect connection between Google and your on-premises router.",
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

func GoogleComputeInterconnectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeInterconnect), &result)
	return &result
}
