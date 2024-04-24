package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeForwardingRule = `{
  "block": {
    "attributes": {
      "all_ports": {
        "computed": true,
        "description": "The 'ports', 'portRange', and 'allPorts' fields are mutually exclusive.\nOnly packets addressed to ports in the specified range will be forwarded\nto the backends configured with this forwarding rule.\n\nThe 'allPorts' field has the following limitations:\n* It requires that the forwarding rule 'IPProtocol' be TCP, UDP, SCTP, or\nL3_DEFAULT.\n* It's applicable only to the following products: internal passthrough\nNetwork Load Balancers, backend service-based external passthrough Network\nLoad Balancers, and internal and external protocol forwarding.\n* Set this field to true to allow packets addressed to any port or packets\nlacking destination port information (for example, UDP fragments after the\nfirst fragment) to be forwarded to the backends configured with this\nforwarding rule. The L3_DEFAULT protocol requires 'allPorts' be set to\ntrue.",
        "description_kind": "plain",
        "type": "bool"
      },
      "allow_global_access": {
        "computed": true,
        "description": "This field is used along with the 'backend_service' field for\ninternal load balancing or with the 'target' field for internal\nTargetInstance.\n\nIf the field is set to 'TRUE', clients can access ILB from all\nregions.\n\nOtherwise only allows access from clients in the same region as the\ninternal load balancer.",
        "description_kind": "plain",
        "type": "bool"
      },
      "allow_psc_global_access": {
        "computed": true,
        "description": "This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.",
        "description_kind": "plain",
        "type": "bool"
      },
      "backend_service": {
        "computed": true,
        "description": "Identifies the backend service to which the forwarding rule sends traffic.\n\nRequired for Internal TCP/UDP Load Balancing and Network Load Balancing;\nmust be omitted for all other load balancer types.",
        "description_kind": "plain",
        "type": "string"
      },
      "base_forwarding_rule": {
        "computed": true,
        "description": "[Output Only] The URL for the corresponding base Forwarding Rule. By base Forwarding Rule, we mean the Forwarding Rule that has the same IP address, protocol, and port settings with the current Forwarding Rule, but without sourceIPRanges specified. Always empty if the current Forwarding Rule does not have sourceIPRanges specified.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_address": {
        "computed": true,
        "description": "IP address for which this forwarding rule accepts traffic. When a client\nsends traffic to this IP address, the forwarding rule directs the traffic\nto the referenced 'target' or 'backendService'.\n\nWhile creating a forwarding rule, specifying an 'IPAddress' is\nrequired under the following circumstances:\n\n* When the 'target' is set to 'targetGrpcProxy' and\n'validateForProxyless' is set to 'true', the\n'IPAddress' should be set to '0.0.0.0'.\n* When the 'target' is a Private Service Connect Google APIs\nbundle, you must specify an 'IPAddress'.\n\n\nOtherwise, you can optionally specify an IP address that references an\nexisting static (reserved) IP address resource. When omitted, Google Cloud\nassigns an ephemeral IP address.\n\nUse one of the following formats to specify an IP address while creating a\nforwarding rule:\n\n* IP address number, as in '100.1.2.3'\n* IPv6 address range, as in '2600:1234::/96'\n* Full resource URL, as in\n'https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name'\n* Partial URL or by name, as in:\n  * 'projects/project_id/regions/region/addresses/address-name'\n  * 'regions/region/addresses/address-name'\n  * 'global/addresses/address-name'\n  * 'address-name'\n\n\nThe forwarding rule's 'target' or 'backendService',\nand in most cases, also the 'loadBalancingScheme', determine the\ntype of IP address that you can use. For detailed information, see\n[IP address\nspecifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).\n\nWhen reading an 'IPAddress', the API always returns the IP\naddress number.",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_protocol": {
        "computed": true,
        "description": "The IP protocol to which this rule applies.\n\nFor protocol forwarding, valid\noptions are 'TCP', 'UDP', 'ESP',\n'AH', 'SCTP', 'ICMP' and\n'L3_DEFAULT'.\n\nThe valid IP protocols are different for different load balancing products\nas described in [Load balancing\nfeatures](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends).\n\nA Forwarding Rule with protocol L3_DEFAULT can attach with target instance or\nbackend service with UNSPECIFIED protocol.\nA forwarding rule with \"L3_DEFAULT\" IPProtocal cannot be attached to a backend service with TCP or UDP. Possible values: [\"TCP\", \"UDP\", \"ESP\", \"AH\", \"SCTP\", \"ICMP\", \"L3_DEFAULT\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_version": {
        "computed": true,
        "description": "The IP address version that will be used by this forwarding rule.\nValid options are IPV4 and IPV6.\n\nIf not set, the IPv4 address will be used by default. Possible values: [\"IPV4\", \"IPV6\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "is_mirroring_collector": {
        "computed": true,
        "description": "Indicates whether or not this load balancer can be used as a collector for\npacket mirroring. To prevent mirroring loops, instances behind this\nload balancer will not have their traffic mirrored even if a\n'PacketMirroring' rule applies to them.\n\nThis can only be set to true for load balancers that have their\n'loadBalancingScheme' set to 'INTERNAL'.",
        "description_kind": "plain",
        "type": "bool"
      },
      "label_fingerprint": {
        "computed": true,
        "description": "The fingerprint used for optimistic locking of this resource.  Used\ninternally during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "Labels to apply to this forwarding rule.  A list of key-\u003evalue pairs.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "load_balancing_scheme": {
        "computed": true,
        "description": "Specifies the forwarding rule type.\n\nFor more information about forwarding rules, refer to\n[Forwarding rule concepts](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts). Default value: \"EXTERNAL\" Possible values: [\"EXTERNAL\", \"EXTERNAL_MANAGED\", \"INTERNAL\", \"INTERNAL_MANAGED\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the resource; provided by the client when the resource is created.\nThe name must be 1-63 characters long, and comply with\n[RFC1035](https://www.ietf.org/rfc/rfc1035.txt).\n\nSpecifically, the name must be 1-63 characters long and match the regular\nexpression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first\ncharacter must be a lowercase letter, and all following characters must\nbe a dash, lowercase letter, or digit, except the last character, which\ncannot be a dash.\n\nFor Private Service Connect forwarding rules that forward traffic to Google\nAPIs, the forwarding rule name must be a 1-20 characters string with\nlowercase letters and numbers and must start with a letter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "This field is not used for external load balancing.\n\nFor Internal TCP/UDP Load Balancing, this field identifies the network that\nthe load balanced IP should belong to for this Forwarding Rule.\nIf the subnetwork is specified, the network of the subnetwork will be used.\nIf neither subnetwork nor this field is specified, the default network will\nbe used.\n\nFor Private Service Connect forwarding rules that forward traffic to Google\nAPIs, a network must be provided.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_tier": {
        "computed": true,
        "description": "This signifies the networking tier used for configuring\nthis load balancer and can only take the following values:\n'PREMIUM', 'STANDARD'.\n\nFor regional ForwardingRule, the valid values are 'PREMIUM' and\n'STANDARD'. For GlobalForwardingRule, the valid value is\n'PREMIUM'.\n\nIf this field is not specified, it is assumed to be 'PREMIUM'.\nIf 'IPAddress' is specified, this value must be equal to the\nnetworkTier of the Address. Possible values: [\"PREMIUM\", \"STANDARD\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "no_automate_dns_zone": {
        "computed": true,
        "description": "This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding rules do not use this field.",
        "description_kind": "plain",
        "type": "bool"
      },
      "port_range": {
        "computed": true,
        "description": "The 'ports', 'portRange', and 'allPorts' fields are mutually exclusive.\nOnly packets addressed to ports in the specified range will be forwarded\nto the backends configured with this forwarding rule.\n\nThe 'portRange' field has the following limitations:\n* It requires that the forwarding rule 'IPProtocol' be TCP, UDP, or SCTP,\nand\n* It's applicable only to the following products: external passthrough\nNetwork Load Balancers, internal and external proxy Network Load\nBalancers, internal and external Application Load Balancers, external\nprotocol forwarding, and Classic VPN.\n* Some products have restrictions on what ports can be used. See\n[port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#port_specifications)\nfor details.\n\nFor external forwarding rules, two or more forwarding rules cannot use the\nsame '[IPAddress, IPProtocol]' pair, and cannot have overlapping\n'portRange's.\n\nFor internal forwarding rules within the same VPC network, two or more\nforwarding rules cannot use the same '[IPAddress, IPProtocol]' pair, and\ncannot have overlapping 'portRange's.\n\n@pattern: \\d+(?:-\\d+)?",
        "description_kind": "plain",
        "type": "string"
      },
      "ports": {
        "computed": true,
        "description": "The 'ports', 'portRange', and 'allPorts' fields are mutually exclusive.\nOnly packets addressed to ports in the specified range will be forwarded\nto the backends configured with this forwarding rule.\n\nThe 'ports' field has the following limitations:\n* It requires that the forwarding rule 'IPProtocol' be TCP, UDP, or SCTP,\nand\n* It's applicable only to the following products: internal passthrough\nNetwork Load Balancers, backend service-based external passthrough Network\nLoad Balancers, and internal protocol forwarding.\n* You can specify a list of up to five ports by number, separated by\ncommas. The ports can be contiguous or discontiguous.\n\nFor external forwarding rules, two or more forwarding rules cannot use the\nsame '[IPAddress, IPProtocol]' pair if they share at least one port\nnumber.\n\nFor internal forwarding rules within the same VPC network, two or more\nforwarding rules cannot use the same '[IPAddress, IPProtocol]' pair if\nthey share at least one port number.\n\n@pattern: \\d+(?:-\\d+)?",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "psc_connection_id": {
        "computed": true,
        "description": "The PSC connection id of the PSC Forwarding Rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "psc_connection_status": {
        "computed": true,
        "description": "The PSC connection status of the PSC Forwarding Rule. Possible values: 'STATUS_UNSPECIFIED', 'PENDING', 'ACCEPTED', 'REJECTED', 'CLOSED'",
        "description_kind": "plain",
        "type": "string"
      },
      "recreate_closed_psc": {
        "computed": true,
        "description": "This is used in PSC consumer ForwardingRule to make terraform recreate the ForwardingRule when the status is closed",
        "description_kind": "plain",
        "type": "bool"
      },
      "region": {
        "description": "A reference to the region where the regional forwarding rule resides.\n\nThis field is not applicable to global forwarding rules.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_directory_registrations": {
        "computed": true,
        "description": "Service Directory resources to register this forwarding rule with.\n\nCurrently, only supports a single Service Directory resource.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "namespace": "string",
              "service": "string"
            }
          ]
        ]
      },
      "service_label": {
        "computed": true,
        "description": "An optional prefix to the service name for this Forwarding Rule.\nIf specified, will be the first label of the fully qualified service\nname.\n\nThe label must be 1-63 characters long, and comply with RFC1035.\nSpecifically, the label must be 1-63 characters long and match the\nregular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first\ncharacter must be a lowercase letter, and all following characters\nmust be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.\n\nThis field is only used for INTERNAL load balancing.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_name": {
        "computed": true,
        "description": "The internal fully qualified service name for this Forwarding Rule.\n\nThis field is only used for INTERNAL load balancing.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_ip_ranges": {
        "computed": true,
        "description": "If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "subnetwork": {
        "computed": true,
        "description": "This field identifies the subnetwork that the load balanced IP should\nbelong to for this Forwarding Rule, used in internal load balancing and\nnetwork load balancing with IPv6.\n\nIf the network specified is in auto subnet mode, this field is optional.\nHowever, a subnetwork must be specified if the network is in custom subnet\nmode or when creating external forwarding rule with IPv6.",
        "description_kind": "plain",
        "type": "string"
      },
      "target": {
        "computed": true,
        "description": "The URL of the target resource to receive the matched traffic.  For\nregional forwarding rules, this target must be in the same region as the\nforwarding rule. For global forwarding rules, this target must be a global\nload balancing resource.\n\nThe forwarded traffic must be of a type appropriate to the target object.\n*  For load balancers, see the \"Target\" column in [Port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).\n*  For Private Service Connect forwarding rules that forward traffic to Google APIs, provide the name of a supported Google API bundle:\n  *  'vpc-sc' - [ APIs that support VPC Service Controls](https://cloud.google.com/vpc-service-controls/docs/supported-products).\n  *  'all-apis' - [All supported Google APIs](https://cloud.google.com/vpc/docs/private-service-connect#supported-apis).\n\n\nFor Private Service Connect forwarding rules that forward traffic to managed services, the target must be a service attachment.",
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
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeForwardingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeForwardingRule), &result)
	return &result
}
