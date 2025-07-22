package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxWebhook = `{
  "block": {
    "attributes": {
      "disabled": {
        "description": "Indicates whether the webhook is disabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "description": "The human-readable name of the webhook, unique within the agent.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enable_spell_correction": {
        "description": "Deprecated. Indicates if automatic spell correction is enabled in detect intent requests.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_stackdriver_logging": {
        "description": "Deprecated. Determines whether this agent should log conversation queries.",
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
      "name": {
        "computed": true,
        "description": "The unique identifier of the webhook.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/webhooks/\u003cWebhook ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The agent to create a webhook for.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_settings": {
        "description": "Deprecated. Name of the SecuritySettings reference for the agent. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/securitySettings/\u003cSecurity Settings ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_flow": {
        "computed": true,
        "description": "Deprecated. Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/flows/\u003cFlow ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "timeout": {
        "description": "Webhook execution timeout.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "generic_web_service": {
        "block": {
          "attributes": {
            "allowed_ca_certs": {
              "description": "Specifies a list of allowed custom CA certificates (in DER format) for\nHTTPS verification. This overrides the default SSL trust store. If this\nis empty or unspecified, Dialogflow will use Google's default trust store\nto verify certificates.\nN.B. Make sure the HTTPS server certificates are signed with \"subject alt\nname\". For instance a certificate can be self-signed using the following\ncommand,\nopenssl x509 -req -days 200 -in example.com.csr \\\n-signkey example.com.key \\\n-out example.com.crt \\\n-extfile \u003c(printf \"\\nsubjectAltName='DNS:www.example.com'\")",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "http_method": {
              "description": "HTTP method for the flexible webhook calls. Standard webhook always uses\nPOST. Possible values: [\"POST\", \"GET\", \"HEAD\", \"PUT\", \"DELETE\", \"PATCH\", \"OPTIONS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parameter_mapping": {
              "description": "Maps the values extracted from specific fields of the flexible webhook\nresponse into session parameters.\n- Key: session parameter name\n- Value: field path in the webhook response",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "request_body": {
              "description": "Defines a custom JSON object as request body to send to flexible webhook.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "request_headers": {
              "description": "The HTTP request headers to send together with webhook requests.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "secret_version_for_username_password": {
              "description": "The SecretManager secret version resource storing the username:password\npair for HTTP Basic authentication.\nFormat: 'projects/{project}/secrets/{secret}/versions/{version}'",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_agent_auth": {
              "description": "Indicate the auth token type generated from the [Diglogflow service\nagent](https://cloud.google.com/iam/docs/service-agents#dialogflow-service-agent).\nThe generated token is sent in the Authorization header. Possible values: [\"NONE\", \"ID_TOKEN\", \"ACCESS_TOKEN\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "uri": {
              "description": "The webhook URI for receiving POST requests. It must use https protocol.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "webhook_type": {
              "description": "Type of the webhook. Possible values: [\"STANDARD\", \"FLEXIBLE\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "oauth_config": {
              "block": {
                "attributes": {
                  "client_id": {
                    "description": "The client ID provided by the 3rd party platform.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret": {
                    "description": "The client secret provided by the 3rd party platform.  If the\n'secret_version_for_client_secret' field is set, this field will be\nignored.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scopes": {
                    "description": "The OAuth scopes to grant.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "secret_version_for_client_secret": {
                    "description": "The name of the SecretManager secret version resource storing the\nclient secret. If this field is set, the 'client_secret' field will be\nignored.\nFormat: 'projects/{project}/secrets/{secret}/versions/{version}'",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "token_endpoint": {
                    "description": "The token endpoint provided by the 3rd party platform to exchange an\naccess token.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Represents configuration of OAuth client credential flow for 3rd party\nAPI authentication.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "secret_versions_for_request_headers": {
              "block": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "secret_version": {
                    "description": "The SecretManager secret version resource storing the header value.\nFormat: 'projects/{project}/secrets/{secret}/versions/{version}'",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The HTTP request headers to send together with webhook requests. Header\nvalues are stored in SecretManager secret versions.\n\nWhen the same header name is specified in both 'request_headers' and\n'secret_versions_for_request_headers', the value in\n'secret_versions_for_request_headers' will be used.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "Represents configuration for a generic web service.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "service_directory": {
        "block": {
          "attributes": {
            "service": {
              "description": "The name of Service Directory service.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "generic_web_service": {
              "block": {
                "attributes": {
                  "allowed_ca_certs": {
                    "description": "Specifies a list of allowed custom CA certificates (in DER format) for\nHTTPS verification. This overrides the default SSL trust store. If this\nis empty or unspecified, Dialogflow will use Google's default trust store\nto verify certificates.\nN.B. Make sure the HTTPS server certificates are signed with \"subject alt\nname\". For instance a certificate can be self-signed using the following\ncommand,\nopenssl x509 -req -days 200 -in example.com.csr \\\n-signkey example.com.key \\\n-out example.com.crt \\\n-extfile \u003c(printf \"\\nsubjectAltName='DNS:www.example.com'\")",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "http_method": {
                    "description": "HTTP method for the flexible webhook calls. Standard webhook always uses\nPOST. Possible values: [\"POST\", \"GET\", \"HEAD\", \"PUT\", \"DELETE\", \"PATCH\", \"OPTIONS\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "parameter_mapping": {
                    "description": "Maps the values extracted from specific fields of the flexible webhook\nresponse into session parameters.\n- Key: session parameter name\n- Value: field path in the webhook response",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "request_body": {
                    "description": "Defines a custom JSON object as request body to send to flexible webhook.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "request_headers": {
                    "description": "The HTTP request headers to send together with webhook requests.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "secret_version_for_username_password": {
                    "description": "The SecretManager secret version resource storing the username:password\npair for HTTP Basic authentication.\nFormat: 'projects/{project}/secrets/{secret}/versions/{version}'",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_agent_auth": {
                    "description": "Indicate the auth token type generated from the [Diglogflow service\nagent](https://cloud.google.com/iam/docs/service-agents#dialogflow-service-agent).\nThe generated token is sent in the Authorization header. Possible values: [\"NONE\", \"ID_TOKEN\", \"ACCESS_TOKEN\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "uri": {
                    "description": "The webhook URI for receiving POST requests. It must use https protocol.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "webhook_type": {
                    "description": "Type of the webhook. Possible values: [\"STANDARD\", \"FLEXIBLE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "oauth_config": {
                    "block": {
                      "attributes": {
                        "client_id": {
                          "description": "The client ID provided by the 3rd party platform.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "client_secret": {
                          "description": "The client secret provided by the 3rd party platform.  If the\n'secret_version_for_client_secret' field is set, this field will be\nignored.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "scopes": {
                          "description": "The OAuth scopes to grant.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "secret_version_for_client_secret": {
                          "description": "The name of the SecretManager secret version resource storing the\nclient secret. If this field is set, the 'client_secret' field will be\nignored.\nFormat: 'projects/{project}/secrets/{secret}/versions/{version}'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "token_endpoint": {
                          "description": "The token endpoint provided by the 3rd party platform to exchange an\naccess token.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents configuration of OAuth client credential flow for 3rd party\nAPI authentication.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "secret_versions_for_request_headers": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_version": {
                          "description": "The SecretManager secret version resource storing the header value.\nFormat: 'projects/{project}/secrets/{secret}/versions/{version}'",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The HTTP request headers to send together with webhook requests. Header\nvalues are stored in SecretManager secret versions.\n\nWhen the same header name is specified in both 'request_headers' and\n'secret_versions_for_request_headers', the value in\n'secret_versions_for_request_headers' will be used.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  }
                },
                "description": "Represents configuration for a generic web service.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for a Service Directory service.",
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

func GoogleDialogflowCxWebhookSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxWebhook), &result)
	return &result
}
