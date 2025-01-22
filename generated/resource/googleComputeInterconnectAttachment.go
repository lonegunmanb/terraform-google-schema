package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeInterconnectAttachment = `{
  "block": {
    "attributes": {
      "admin_enabled": {
        "description": "Whether the VLAN attachment is enabled or disabled.  When using\nPARTNER type this will Pre-Activate the interconnect attachment",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "bandwidth": {
        "computed": true,
        "description": "Provisioned bandwidth capacity for the interconnect attachment.\nFor attachments of type DEDICATED, the user can set the bandwidth.\nFor attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.\nOutput only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,\nDefaults to BPS_10G Possible values: [\"BPS_50M\", \"BPS_100M\", \"BPS_200M\", \"BPS_300M\", \"BPS_400M\", \"BPS_500M\", \"BPS_1G\", \"BPS_2G\", \"BPS_5G\", \"BPS_10G\", \"BPS_20G\", \"BPS_50G\", \"BPS_100G\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "candidate_subnets": {
        "description": "Up to 16 candidate prefixes that can be used to restrict the allocation\nof cloudRouterIpAddress and customerRouterIpAddress for this attachment.\nAll prefixes must be within link-local address space (169.254.0.0/16)\nand must be /29 or shorter (/28, /27, etc). Google will attempt to select\nan unused /29 from the supplied candidate prefix(es). The request will\nfail if all possible /29s are in use on Google's edge. If not supplied,\nGoogle will randomly select an unused /29 from all of link-local space.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "cloud_router_ip_address": {
        "computed": true,
        "description": "IPv4 address + prefix length to be configured on Cloud Router\nInterface for this interconnect attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_router_ipv6_address": {
        "computed": true,
        "description": "IPv6 address + prefix length to be configured on Cloud Router\nInterface for this interconnect attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_router_ip_address": {
        "computed": true,
        "description": "IPv4 address + prefix length to be configured on the customer\nrouter subinterface for this interconnect attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_router_ipv6_address": {
        "computed": true,
        "description": "IPv6 address + prefix length to be configured on the customer\nrouter subinterface for this interconnect attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "edge_availability_domain": {
        "computed": true,
        "description": "Desired availability domain for the attachment. Only available for type\nPARTNER, at creation time. For improved reliability, customers should\nconfigure a pair of attachments with one per availability domain. The\nselected availability domain will be provided to the Partner via the\npairing key so that the provisioned circuit will lie in the specified\ndomain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption": {
        "description": "Indicates the user-supplied encryption option of this interconnect\nattachment. Can only be specified at attachment creation for PARTNER or\nDEDICATED attachments.\n* NONE - This is the default value, which means that the VLAN attachment\ncarries unencrypted traffic. VMs are able to send traffic to, or receive\ntraffic from, such a VLAN attachment.\n* IPSEC - The VLAN attachment carries only encrypted traffic that is\nencrypted by an IPsec device, such as an HA VPN gateway or third-party\nIPsec VPN. VMs cannot directly send traffic to, or receive traffic from,\nsuch a VLAN attachment. To use HA VPN over Cloud Interconnect, the VLAN\nattachment must be created with this option. Default value: \"NONE\" Possible values: [\"NONE\", \"IPSEC\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "google_reference_id": {
        "computed": true,
        "description": "Google reference ID, to be used when raising support tickets with\nGoogle or otherwise to debug backend connectivity issues.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "interconnect": {
        "description": "URL of the underlying Interconnect object that this attachment's\ntraffic will traverse through. Required if type is DEDICATED, must not\nbe set if type is PARTNER.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipsec_internal_addresses": {
        "description": "URL of addresses that have been reserved for the interconnect attachment,\nUsed only for interconnect attachment that has the encryption option as\nIPSEC.\nThe addresses must be RFC 1918 IP address ranges. When creating HA VPN\ngateway over the interconnect attachment, if the attachment is configured\nto use an RFC 1918 IP address, then the VPN gateway's IP address will be\nallocated from the IP address range specified here.\nFor example, if the HA VPN gateway's interface 0 is paired to this\ninterconnect attachment, then an RFC 1918 IP address for the VPN gateway\ninterface 0 will be allocated from the IP address specified for this\ninterconnect attachment.\nIf this field is not specified for interconnect attachment that has\nencryption option as IPSEC, later on when creating HA VPN gateway on this\ninterconnect attachment, the HA VPN gateway's IP address will be\nallocated from regional external IP address pool.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "mtu": {
        "computed": true,
        "description": "Maximum Transmission Unit (MTU), in bytes, of packets passing through\nthis interconnect attachment. Currently, only 1440 and 1500 are allowed. If not specified, the value will default to 1440.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is created. The\nname must be 1-63 characters long, and comply with RFC1035. Specifically, the\nname must be 1-63 characters long and match the regular expression\n'[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a\nlowercase letter, and all following characters must be a dash, lowercase\nletter, or digit, except the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pairing_key": {
        "computed": true,
        "description": "[Output only for type PARTNER. Not present for DEDICATED]. The opaque\nidentifier of an PARTNER attachment used to initiate provisioning with\na selected partner. Of the form \"XXXXX/region/domain\"",
        "description_kind": "plain",
        "type": "string"
      },
      "partner_asn": {
        "computed": true,
        "description": "[Output only for type PARTNER. Not present for DEDICATED]. Optional\nBGP ASN for the router that should be supplied by a layer 3 Partner if\nthey configured BGP on behalf of the customer.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_interconnect_info": {
        "computed": true,
        "description": "Information specific to an InterconnectAttachment. This property\nis populated if the interconnect that this is attached to is of type DEDICATED.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "tag8021q": "number"
            }
          ]
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where the regional interconnect attachment resides.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "router": {
        "description": "URL of the cloud router to be used for dynamic routing. This router must be in\nthe same region as this InterconnectAttachment. The InterconnectAttachment will\nautomatically connect the Interconnect to the network \u0026 region within which the\nCloud Router is configured.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stack_type": {
        "computed": true,
        "description": "The stack type for this interconnect attachment to identify whether the IPv6\nfeature is enabled or not. If not specified, IPV4_ONLY will be used.\nThis field can be both set at interconnect attachments creation and update\ninterconnect attachment operations. Possible values: [\"IPV4_IPV6\", \"IPV4_ONLY\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "[Output Only] The current state of this attachment's functionality.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_length": {
        "description": "Length of the IPv4 subnet mask. Allowed values: 29 (default), 30. The default value is 29,\nexcept for Cross-Cloud Interconnect connections that use an InterconnectRemoteLocation with a\nconstraints.subnetLengthRange.min equal to 30. For example, connections that use an Azure\nremote location fall into this category. In these cases, the default value is 30, and\nrequesting 29 returns an error. Where both 29 and 30 are allowed, 29 is preferred, because it\ngives Google Cloud Support more debugging visibility.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "type": {
        "computed": true,
        "description": "The type of InterconnectAttachment you wish to create. Defaults to\nDEDICATED. Possible values: [\"DEDICATED\", \"PARTNER\", \"PARTNER_PROVIDER\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vlan_tag8021q": {
        "computed": true,
        "description": "The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When\nusing PARTNER type this will be managed upstream.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func GoogleComputeInterconnectAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeInterconnectAttachment), &result)
	return &result
}
