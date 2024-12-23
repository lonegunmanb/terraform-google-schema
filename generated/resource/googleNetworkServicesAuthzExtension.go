package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkServicesAuthzExtension = `{
  "block": {
    "attributes": {
      "authority": {
        "description": "The :authority header in the gRPC request sent from Envoy to the extension service.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp when the resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A human-readable description of the resource.",
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
      "fail_open": {
        "computed": true,
        "description": "Determines how the proxy behaves if the call to the extension fails or times out.\nWhen set to TRUE, request or response processing continues without error. Any subsequent extensions in the extension chain are also executed. When set to FALSE or the default setting of FALSE is used, one of the following happens:\n* If response headers have not been delivered to the downstream client, a generic 500 error is returned to the client. The error response can be tailored by configuring a custom error response in the load balancer.\n* If response headers have been delivered, then the HTTP stream to the downstream client is reset.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "forward_headers": {
        "description": "List of the HTTP headers to forward to the extension (from the client). If omitted, all headers are sent. Each element is a string indicating the header name.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Set of labels associated with the AuthzExtension resource.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "load_balancing_scheme": {
        "description": "All backend services and forwarding rules referenced by this extension must share the same load balancing scheme.\nFor more information, refer to [Backend services overview](https://cloud.google.com/load-balancing/docs/backend-service). Possible values: [\"INTERNAL_MANAGED\", \"EXTERNAL_MANAGED\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metadata": {
        "description": "The metadata provided here is included as part of the metadata_context (of type google.protobuf.Struct) in the ProcessingRequest message sent to the extension server. The metadata is available under the namespace com.google.authz_extension.\u003cresourceName\u003e. The following variables are supported in the metadata Struct:\n\n{forwarding_rule_id} - substituted with the forwarding rule's fully qualified resource name.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "Identifier. Name of the AuthzExtension resource.",
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
      "service": {
        "description": "The reference to the service that runs the extension.\nTo configure a callout extension, service must be a fully-qualified reference to a [backend service](https://cloud.google.com/compute/docs/reference/rest/v1/backendServices) in the format:\nhttps://www.googleapis.com/compute/v1/projects/{project}/regions/{region}/backendServices/{backendService} or https://www.googleapis.com/compute/v1/projects/{project}/global/backendServices/{backendService}.",
        "description_kind": "plain",
        "required": true,
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
      "timeout": {
        "description": "Specifies the timeout for each individual message on the stream. The timeout must be between 10-10000 milliseconds.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp when the resource was updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "wire_format": {
        "description": "The format of communication supported by the callout extension. Default value: \"EXT_PROC_GRPC\" Possible values: [\"WIRE_FORMAT_UNSPECIFIED\", \"EXT_PROC_GRPC\"]",
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

func GoogleNetworkServicesAuthzExtensionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkServicesAuthzExtension), &result)
	return &result
}
