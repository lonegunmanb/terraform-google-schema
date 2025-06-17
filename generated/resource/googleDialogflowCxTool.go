package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowCxTool = `{
  "block": {
    "attributes": {
      "description": {
        "description": "High level description of the Tool and its usage.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description": "The human-readable name of the tool, unique within the agent.",
        "description_kind": "plain",
        "required": true,
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
        "description": "The unique identifier of the Tool.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e/tools/\u003cTool ID\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The agent to create a Tool for.\nFormat: projects/\u003cProject ID\u003e/locations/\u003cLocation ID\u003e/agents/\u003cAgent ID\u003e.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tool_type": {
        "computed": true,
        "description": "The tool type.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "data_store_spec": {
        "block": {
          "block_types": {
            "data_store_connections": {
              "block": {
                "attributes": {
                  "data_store": {
                    "description": "The full name of the referenced data store. Formats: projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore} projects/{project}/locations/{location}/dataStores/{dataStore}",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "data_store_type": {
                    "description": "The type of the connected data store.\nSee [DataStoreType](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/DataStoreConnection#datastoretype) for valid values.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "document_processing_mode": {
                    "description": "The document processing mode for the data store connection. Should only be set for PUBLIC_WEB and UNSTRUCTURED data stores. If not set it is considered as DOCUMENTS, as this is the legacy mode.\nSee [DocumentProcessingMode](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/DataStoreConnection#documentprocessingmode) for valid values.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "List of data stores to search.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            },
            "fallback_prompt": {
              "block": {
                "description": "Fallback prompt configurations to use.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Data store search tool specification.\nThis field is part of a union field 'specification': Only one of 'openApiSpec', 'dataStoreSpec', or 'functionSpec' may be set.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "function_spec": {
        "block": {
          "attributes": {
            "input_schema": {
              "description": "Optional. The JSON schema is encapsulated in a [google.protobuf.Struct](https://protobuf.dev/reference/protobuf/google.protobuf/#struct) to describe the input of the function.\nThis input is a JSON object that contains the function's parameters as properties of the object",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "output_schema": {
              "description": "Optional. The JSON schema is encapsulated in a [google.protobuf.Struct](https://protobuf.dev/reference/protobuf/google.protobuf/#struct) to describe the output of the function.\nThis output is a JSON object that contains the function's parameters as properties of the object",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Client side executed function specification.\nThis field is part of a union field 'specification': Only one of 'openApiSpec', 'dataStoreSpec', or 'functionSpec' may be set.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "open_api_spec": {
        "block": {
          "attributes": {
            "text_schema": {
              "description": "The OpenAPI schema specified as a text.\nThis field is part of a union field 'schema': only one of 'textSchema' may be set.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "authentication": {
              "block": {
                "block_types": {
                  "api_key_config": {
                    "block": {
                      "attributes": {
                        "api_key": {
                          "description": "Optional. The API key. If the 'secretVersionForApiKey'' field is set, this field will be ignored.",
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "key_name": {
                          "description": "The parameter name or the header name of the API key.\nE.g., If the API request is \"https://example.com/act?X-Api-Key=\", \"X-Api-Key\" would be the parameter name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "request_location": {
                          "description": "Key location in the request.\nSee [RequestLocation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.tools#requestlocation) for valid values.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_version_for_api_key": {
                          "description": "Optional. The name of the SecretManager secret version resource storing the API key.\nIf this field is set, the apiKey field will be ignored.\nFormat: projects/{project}/secrets/{secret}/versions/{version}",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Config for API key auth.\nThis field is part of a union field 'auth_config': Only one of 'apiKeyConfig', 'oauthConfig', 'serviceAgentAuthConfig', or 'bearerTokenConfig' may be set.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "bearer_token_config": {
                    "block": {
                      "attributes": {
                        "secret_version_for_token": {
                          "description": "Optional. The name of the SecretManager secret version resource storing the Bearer token. If this field is set, the 'token' field will be ignored.\nFormat: projects/{project}/secrets/{secret}/versions/{version}",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "token": {
                          "description": "Optional. The text token appended to the text Bearer to the request Authorization header.\n[Session parameters reference](https://cloud.google.com/dialogflow/cx/docs/concept/parameter#session-ref) can be used to pass the token dynamically, e.g. '$session.params.parameter-id'.",
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        }
                      },
                      "description": "Config for bearer token auth.\nThis field is part of a union field 'auth_config': Only one of 'apiKeyConfig', 'oauthConfig', 'serviceAgentAuthConfig', or 'bearerTokenConfig' may be set.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "oauth_config": {
                    "block": {
                      "attributes": {
                        "client_id": {
                          "description": "The client ID from the OAuth provider.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "client_secret": {
                          "description": "Optional. The client secret from the OAuth provider. If the 'secretVersionForClientSecret' field is set, this field will be ignored.",
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "oauth_grant_type": {
                          "description": "OAuth grant types.\nSee [OauthGrantType](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.tools#oauthgranttype) for valid values",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "scopes": {
                          "description": "Optional. The OAuth scopes to grant.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "secret_version_for_client_secret": {
                          "description": "Optional. The name of the SecretManager secret version resource storing the client secret.\nIf this field is set, the clientSecret field will be ignored.\nFormat: projects/{project}/secrets/{secret}/versions/{version}",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "token_endpoint": {
                          "description": "The token endpoint in the OAuth provider to exchange for an access token.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Config for OAuth.\nThis field is part of a union field 'auth_config': Only one of 'apiKeyConfig', 'oauthConfig', 'serviceAgentAuthConfig', or 'bearerTokenConfig' may be set.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "service_agent_auth_config": {
                    "block": {
                      "attributes": {
                        "service_agent_auth": {
                          "description": "Optional. Indicate the auth token type generated from the Diglogflow service agent.\nThe generated token is sent in the Authorization header.\nSee [ServiceAgentAuth](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.tools#serviceagentauth) for valid values.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Config for [Diglogflow service agent](https://cloud.google.com/iam/docs/service-agents#dialogflow-service-agent) auth.\nThis field is part of a union field 'auth_config': Only one of 'apiKeyConfig', 'oauthConfig', 'serviceAgentAuthConfig', or 'bearerTokenConfig' may be set.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Authentication information required by the API.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "service_directory_config": {
              "block": {
                "attributes": {
                  "service": {
                    "description": "The name of [Service Directory](https://cloud.google.com/service-directory/docs) service.\nFormat: projects/\u003cProjectID\u003e/locations/\u003cLocationID\u003e/namespaces/\u003cNamespaceID\u003e/services/\u003cServiceID\u003e. LocationID of the service directory must be the same as the location of the agent.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Optional. Service Directory configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "tls_config": {
              "block": {
                "block_types": {
                  "ca_certs": {
                    "block": {
                      "attributes": {
                        "cert": {
                          "description": "The allowed custom CA certificates (in DER format) for HTTPS verification. This overrides the default SSL trust store.\nIf this is empty or unspecified, Dialogflow will use Google's default trust store to verify certificates.\nN.B. Make sure the HTTPS server certificates are signed with \"subject alt name\".\nFor instance a certificate can be self-signed using the following command:\n'''\n  openssl x509 -req -days 200 -in example.com.csr \\\n    -signkey example.com.key \\\n    -out example.com.crt \\\n    -extfile \u003c(printf \"\\nsubjectAltName='DNS:www.example.com'\")\n'''\nA base64-encoded string.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "display_name": {
                          "description": "The name of the allowed custom CA certificates. This can be used to disambiguate the custom CA certificates.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Specifies a list of allowed custom CA certificates for HTTPS verification.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. TLS configuration for the HTTPS verification.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "OpenAPI specification of the Tool.\nThis field is part of a union field 'specification': Only one of 'openApiSpec', 'dataStoreSpec', or 'functionSpec' may be set.",
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

func GoogleDialogflowCxToolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowCxTool), &result)
	return &result
}
