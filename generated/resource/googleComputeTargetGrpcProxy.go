package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeTargetGrpcProxy = `{
  "block": {
    "attributes": {
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
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource. A hash of the contents stored in\nthis object. This field is used in optimistic locking. This field\nwill be ignored when inserting a TargetGrpcProxy. An up-to-date\nfingerprint must be provided in order to patch/update the\nTargetGrpcProxy; otherwise, the request will fail with error\n412 conditionNotMet. To see the latest fingerprint, make a get()\nrequest to retrieve the TargetGrpcProxy. A base64-encoded string.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource\nis created. The name must be 1-63 characters long, and comply\nwith RFC1035. Specifically, the name must be 1-63 characters long\nand match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which\nmeans the first character must be a lowercase letter, and all\nfollowing characters must be a dash, lowercase letter, or digit,\nexcept the last character, which cannot be a dash.",
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
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "self_link_with_id": {
        "computed": true,
        "description": "Server-defined URL with id for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "url_map": {
        "description": "URL to the UrlMap resource that defines the mapping from URL to\nthe BackendService. The protocol field in the BackendService\nmust be set to GRPC.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "validate_for_proxyless": {
        "description": "If true, indicates that the BackendServices referenced by\nthe urlMap may be accessed by gRPC applications without using\na sidecar proxy. This will enable configuration checks on urlMap\nand its referenced BackendServices to not allow unsupported features.\nA gRPC application must use \"xds:///\" scheme in the target URI\nof the service it is connecting to. If false, indicates that the\nBackendServices referenced by the urlMap will be accessed by gRPC\napplications via a sidecar proxy. In this case, a gRPC application\nmust not use \"xds:///\" scheme in the target URI of the service\nit is connecting to",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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

func GoogleComputeTargetGrpcProxySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeTargetGrpcProxy), &result)
	return &result
}
