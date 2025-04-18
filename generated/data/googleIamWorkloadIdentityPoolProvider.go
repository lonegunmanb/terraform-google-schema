package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIamWorkloadIdentityPoolProvider = `{
  "block": {
    "attributes": {
      "attribute_condition": {
        "computed": true,
        "description": "[A Common Expression Language](https://opensource.google/projects/cel) expression, in\nplain text, to restrict what otherwise valid authentication credentials issued by the\nprovider should not be accepted.\n\nThe expression must output a boolean representing whether to allow the federation.\n\nThe following keywords may be referenced in the expressions:\n  * 'assertion': JSON representing the authentication credential issued by the provider.\n  * 'google': The Google attributes mapped from the assertion in the 'attribute_mappings'.\n  * 'attribute': The custom attributes mapped from the assertion in the 'attribute_mappings'.\n\nThe maximum length of the attribute condition expression is 4096 characters. If\nunspecified, all valid authentication credential are accepted.\n\nThe following example shows how to only allow credentials with a mapped 'google.groups'\nvalue of 'admins':\n'''\n\"'admins' in google.groups\"\n'''",
        "description_kind": "plain",
        "type": "string"
      },
      "attribute_mapping": {
        "computed": true,
        "description": "Maps attributes from authentication credentials issued by an external identity provider\nto Google Cloud attributes, such as 'subject' and 'segment'.\n\nEach key must be a string specifying the Google Cloud IAM attribute to map to.\n\nThe following keys are supported:\n  * 'google.subject': The principal IAM is authenticating. You can reference this value\n    in IAM bindings. This is also the subject that appears in Cloud Logging logs.\n    Cannot exceed 127 characters.\n  * 'google.groups': Groups the external identity belongs to. You can grant groups\n    access to resources using an IAM 'principalSet' binding; access applies to all\n    members of the group.\n\nYou can also provide custom attributes by specifying 'attribute.{custom_attribute}',\nwhere '{custom_attribute}' is the name of the custom attribute to be mapped. You can\ndefine a maximum of 50 custom attributes. The maximum length of a mapped attribute key\nis 100 characters, and the key may only contain the characters [a-z0-9_].\n\nYou can reference these attributes in IAM policies to define fine-grained access for a\nworkload to Google Cloud resources. For example:\n  * 'google.subject':\n    'principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}'\n  * 'google.groups':\n    'principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}'\n  * 'attribute.{custom_attribute}':\n    'principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}'\n\nEach value must be a [Common Expression Language](https://opensource.google/projects/cel)\nfunction that maps an identity provider credential to the normalized attribute specified\nby the corresponding map key.\n\nYou can use the 'assertion' keyword in the expression to access a JSON representation of\nthe authentication credential issued by the provider.\n\nThe maximum length of an attribute mapping expression is 2048 characters. When evaluated,\nthe total size of all mapped attributes must not exceed 8KB.\n\nFor AWS providers, the following rules apply:\n  - If no attribute mapping is defined, the following default mapping applies:\n    '''\n    {\n      \"google.subject\":\"assertion.arn\",\n      \"attribute.aws_role\":\n        \"assertion.arn.contains('assumed-role')\"\n        \" ? assertion.arn.extract('{account_arn}assumed-role/')\"\n        \"   + 'assumed-role/'\"\n        \"   + assertion.arn.extract('assumed-role/{role_name}/')\"\n        \" : assertion.arn\",\n    }\n    '''\n  - If any custom attribute mappings are defined, they must include a mapping to the\n    'google.subject' attribute.\n\nFor OIDC providers, the following rules apply:\n  - Custom attribute mappings must be defined, and must include a mapping to the\n    'google.subject' attribute. For example, the following maps the 'sub' claim of the\n    incoming credential to the 'subject' attribute on a Google token.\n    '''\n    {\"google.subject\": \"assertion.sub\"}\n    '''",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "aws": {
        "computed": true,
        "description": "An Amazon Web Services identity provider. Not compatible with the property oidc or saml.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "account_id": "string"
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description": "A description for the provider. Cannot exceed 256 characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "disabled": {
        "computed": true,
        "description": "Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.\nHowever, existing tokens still grant access.",
        "description_kind": "plain",
        "type": "bool"
      },
      "display_name": {
        "computed": true,
        "description": "A display name for the provider. Cannot exceed 32 characters.",
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
        "computed": true,
        "description": "The resource name of the provider as\n'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}/providers/{workload_identity_pool_provider_id}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "oidc": {
        "computed": true,
        "description": "An OpenId Connect 1.0 identity provider. Not compatible with the property aws or saml.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allowed_audiences": [
                "list",
                "string"
              ],
              "issuer_uri": "string",
              "jwks_json": "string"
            }
          ]
        ]
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "saml": {
        "computed": true,
        "description": "An SAML 2.0 identity provider. Not compatible with the property oidc or aws.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "idp_metadata_xml": "string"
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description": "The state of the provider.\n* STATE_UNSPECIFIED: State unspecified.\n* ACTIVE: The provider is active, and may be used to validate authentication credentials.\n* DELETED: The provider is soft-deleted. Soft-deleted providers are permanently deleted\n  after approximately 30 days. You can restore a soft-deleted provider using\n  UndeleteWorkloadIdentityPoolProvider. You cannot reuse the ID of a soft-deleted provider\n  until it is permanently deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "workload_identity_pool_id": {
        "description": "The ID used for the pool, which is the final component of the pool resource name. This\nvalue should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix\n'gcp-' is reserved for use by Google, and may not be specified.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "workload_identity_pool_provider_id": {
        "description": "The ID for the provider, which becomes the final component of the resource name. This\nvalue must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix\n'gcp-' is reserved for use by Google, and may not be specified.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "x509": {
        "computed": true,
        "description": "An X.509-type identity provider represents a CA. It is trusted to assert a\nclient identity if the client has a certificate that chains up to this CA.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "trust_store": [
                "list",
                [
                  "object",
                  {
                    "intermediate_cas": [
                      "list",
                      [
                        "object",
                        {
                          "pem_certificate": "string"
                        }
                      ]
                    ],
                    "trust_anchors": [
                      "list",
                      [
                        "object",
                        {
                          "pem_certificate": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleIamWorkloadIdentityPoolProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIamWorkloadIdentityPoolProvider), &result)
	return &result
}
