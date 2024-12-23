package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeTargetHttpsProxy = `{
  "block": {
    "attributes": {
      "certificate_manager_certificates": {
        "description": "URLs to certificate manager certificate resources that are used to authenticate connections between users and the load balancer.\nCertificate manager certificates only apply when the load balancing scheme is set to INTERNAL_MANAGED.\nFor EXTERNAL and EXTERNAL_MANAGED, use certificate_map instead.\nsslCertificates and certificateManagerCertificates fields can not be defined together.\nAccepted format is '//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificates/{resourceName}' or just the self_link 'projects/{project}/locations/{location}/certificates/{resourceName}'",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "certificate_map": {
        "description": "A reference to the CertificateMap resource uri that identifies a certificate map\nassociated with the given target proxy. This field is only supported for EXTERNAL and EXTERNAL_MANAGED load balancing schemes.\nFor INTERNAL_MANAGED, use certificate_manager_certificates instead.\nAccepted format is '//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "http_keep_alive_timeout_sec": {
        "description": "Specifies how long to keep a connection open, after completing a response,\nwhile there is no matching traffic (in seconds). If an HTTP keepalive is\nnot specified, a default value will be used. For Global\nexternal HTTP(S) load balancer, the default value is 610 seconds, the\nminimum allowed value is 5 seconds and the maximum allowed value is 1200\nseconds. For cross-region internal HTTP(S) load balancer, the default\nvalue is 600 seconds, the minimum allowed value is 5 seconds, and the\nmaximum allowed value is 600 seconds. For Global external HTTP(S) load\nbalancer (classic), this option is not available publicly.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
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
      "proxy_bind": {
        "computed": true,
        "description": "This field only applies when the forwarding rule that references\nthis target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "proxy_id": {
        "computed": true,
        "description": "The unique identifier for the resource.",
        "description_kind": "plain",
        "type": "number"
      },
      "quic_override": {
        "description": "Specifies the QUIC override policy for this resource. This determines\nwhether the load balancer will attempt to negotiate QUIC with clients\nor not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is\nspecified, Google manages whether QUIC is used. Default value: \"NONE\" Possible values: [\"NONE\", \"ENABLE\", \"DISABLE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "server_tls_policy": {
        "description": "A URL referring to a networksecurity.ServerTlsPolicy\nresource that describes how the proxy should authenticate inbound\ntraffic. serverTlsPolicy only applies to a global TargetHttpsProxy\nattached to globalForwardingRules with the loadBalancingScheme\nset to INTERNAL_SELF_MANAGED or EXTERNAL or EXTERNAL_MANAGED.\nFor details which ServerTlsPolicy resources are accepted with\nINTERNAL_SELF_MANAGED and which with EXTERNAL, EXTERNAL_MANAGED\nloadBalancingScheme consult ServerTlsPolicy documentation.\nIf left blank, communications are not encrypted.\n\nIf you remove this field from your configuration at the same time as\ndeleting or recreating a referenced ServerTlsPolicy resource, you will\nreceive a resourceInUseByAnotherResource error. Use lifecycle.create_before_destroy\nwithin the ServerTlsPolicy resource to avoid this.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_certificates": {
        "description": "URLs to SslCertificate resources that are used to authenticate connections between users and the load balancer.\nCurrently, you may specify up to 15 SSL certificates. sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.\nsslCertificates and certificateManagerCertificates can not be defined together.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "ssl_policy": {
        "description": "A reference to the SslPolicy resource that will be associated with\nthe TargetHttpsProxy resource. If not set, the TargetHttpsProxy\nresource will not have any SSL policy configured.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tls_early_data": {
        "computed": true,
        "description": "Specifies whether TLS 1.3 0-RTT Data (“Early Data”) should be accepted for this service.\nEarly Data allows a TLS resumption handshake to include the initial application payload\n(a HTTP request) alongside the handshake, reducing the effective round trips to “zero”.\nThis applies to TLS 1.3 connections over TCP (HTTP/2) as well as over UDP (QUIC/h3). Possible values: [\"STRICT\", \"PERMISSIVE\", \"DISABLED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "url_map": {
        "description": "A reference to the UrlMap resource that defines the mapping from URL\nto the BackendService.",
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

func GoogleComputeTargetHttpsProxySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeTargetHttpsProxy), &result)
	return &result
}
