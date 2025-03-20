package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkServicesGateway = `{
  "block": {
    "attributes": {
      "addresses": {
        "computed": true,
        "description": "Zero or one IPv4 or IPv6 address on which the Gateway will receive the traffic.\nWhen no address is provided, an IP from the subnetwork is allocated.\n\nThis field only applies to gateways of type 'SECURE_WEB_GATEWAY'.\nGateways of type 'OPEN_MESH' listen on 0.0.0.0 for IPv4 and :: for IPv6.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "certificate_urls": {
        "description": "A fully-qualified Certificates URL reference. The proxy presents a Certificate (selected based on SNI) when establishing a TLS connection.\nThis feature only applies to gateways of type 'SECURE_WEB_GATEWAY'.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp when the resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_swg_autogen_router_on_destroy": {
        "description": "When deleting a gateway of type 'SECURE_WEB_GATEWAY', this boolean option will also delete auto generated router by the gateway creation.\nIf there is no other gateway of type 'SECURE_WEB_GATEWAY' remaining for that region and network it will be deleted.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "A free-text description of the resource. Max length 1024 characters.",
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
      "envoy_headers": {
        "description": "Determines if envoy will insert internal debug headers into upstream requests.\nOther Envoy headers may still be injected.\nBy default, envoy will not insert any debug headers. Possible values: [\"NONE\", \"DEBUG_HEADERS\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateway_security_policy": {
        "description": "A fully-qualified GatewaySecurityPolicy URL reference. Defines how a server should apply security policy to inbound (VM to Proxy) initiated connections.\nFor example: 'projects/*/locations/*/gatewaySecurityPolicies/swg-policy'.\nThis policy is specific to gateways of type 'SECURE_WEB_GATEWAY'.",
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
      "ip_version": {
        "description": "The IP Version that will be used by this gateway. Possible values: [\"IPV4\", \"IPV6\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Set of label tags associated with the Gateway resource.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the gateway.\nThe default value is 'global'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the Gateway resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The relative resource name identifying the VPC network that is using this configuration.\nFor example: 'projects/*/global/networks/network-1'.\n\nCurrently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ports": {
        "description": "One or more port numbers (1-65535), on which the Gateway will receive traffic.\nThe proxy binds to the specified ports. Gateways of type 'SECURE_WEB_GATEWAY' are limited to 1 port.\n Gateways of type 'OPEN_MESH' listen on 0.0.0.0 for IPv4 and :: for IPv6 and support multiple ports.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "number"
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_mode": {
        "description": "The routing mode of the Gateway. This field is configurable only for gateways of type SECURE_WEB_GATEWAY. This field is required for gateways of type SECURE_WEB_GATEWAY. Possible values: [\"NEXT_HOP_ROUTING_MODE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scope": {
        "description": "Immutable. Scope determines how configuration across multiple Gateway instances are merged.\nThe configuration for multiple Gateway instances with the same scope will be merged as presented as a single coniguration to the proxy/load balancer.\n\nMax length 64 characters. Scope should start with a letter and can only have letters, numbers, hyphens.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "Server-defined URL of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "server_tls_policy": {
        "description": "A fully-qualified ServerTLSPolicy URL reference. Specifies how TLS traffic is terminated. If empty, TLS termination is disabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnetwork": {
        "description": "The relative resource name identifying the subnetwork in which this SWG is allocated.\nFor example: projects/*/regions/us-central1/subnetworks/network-1.\n\nCurrently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY'.",
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
      },
      "type": {
        "description": "Immutable. The type of the customer managed gateway. Possible values: [\"OPEN_MESH\", \"SECURE_WEB_GATEWAY\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp when the resource was updated.",
        "description_kind": "plain",
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

func GoogleNetworkServicesGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkServicesGateway), &result)
	return &result
}
