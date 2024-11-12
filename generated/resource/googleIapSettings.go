package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIapSettings = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the IAP protected resource. Name can have below resources:\n* organizations/{organization_id}\n* folders/{folder_id}\n* projects/{projects_id}\n* projects/{projects_id}/iap_web\n* projects/{projects_id}/iap_web/compute\n* projects/{projects_id}/iap_web/compute-{region}\n* projects/{projects_id}/iap_web/compute/service/{service_id}\n* projects/{projects_id}/iap_web/compute-{region}/service/{service_id}\n* projects/{projects_id}/iap_web/appengine-{app_id}\n* projects/{projects_id}/iap_web/appengine-{app_id}/service/{service_id}\n* projects/{projects_id}/iap_web/appengine-{app_id}/service/{service_id}/version/{version_id}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "access_settings": {
        "block": {
          "attributes": {
            "identity_sources": {
              "description": "Identity sources that IAP can use to authenticate the end user. Only one identity source\ncan be configured. The possible values are:\n\n* 'WORKFORCE_IDENTITY_FEDERATION': Use external identities set up on Google Cloud Workforce\n  \t\t\t\t     Identity Federation. Possible values: [\"WORKFORCE_IDENTITY_FEDERATION\"]",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "allowed_domains_settings": {
              "block": {
                "attributes": {
                  "domains": {
                    "description": "List of trusted domains.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "enable": {
                    "description": "Configuration for customers to opt in for the feature.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Settings to configure and enable allowed domains.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "cors_settings": {
              "block": {
                "attributes": {
                  "allow_http_options": {
                    "description": "Configuration to allow HTTP OPTIONS calls to skip authorization.\nIf undefined, IAP will not apply any special logic to OPTIONS requests.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration to allow cross-origin requests via IAP.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gcip_settings": {
              "block": {
                "attributes": {
                  "login_page_uri": {
                    "description": "Login page URI associated with the GCIP tenants. Typically, all resources within\nthe same project share the same login page, though it could be overridden at the\nsub resource level.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tenant_ids": {
                    "description": "GCIP tenant ids that are linked to the IAP resource. tenantIds could be a string\nbeginning with a number character to indicate authenticating with GCIP tenant flow,\nor in the format of _ to indicate authenticating with GCIP agent flow. If agent flow\nis used, tenantIds should only contain one single element, while for tenant flow,\ntenantIds can contain multiple elements.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "GCIP claims and endpoint configurations for 3p identity providers.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oauth_settings": {
              "block": {
                "attributes": {
                  "login_hint": {
                    "description": "Domain hint to send as hd=? parameter in OAuth request flow.\nEnables redirect to primary IDP by skipping Google's login screen.\n(https://developers.google.com/identity/protocols/OpenIDConnect#hd-param)\nNote: IAP does not verify that the id token's hd claim matches this value\nsince access behavior is managed by IAM policies.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "programmatic_clients": {
                    "description": "List of client ids allowed to use IAP programmatically.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Settings to configure IAP's OAuth behavior.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "reauth_settings": {
              "block": {
                "attributes": {
                  "max_age": {
                    "description": "Reauth session lifetime, how long before a user has to reauthenticate again.\nA duration in seconds with up to nine fractional digits, ending with 's'.\nExample: \"3.5s\".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "method": {
                    "description": "Reauth method requested. The possible values are:\n\n* 'LOGIN': Prompts the user to log in again.\n* 'SECURE_KEY': User must use their secure key 2nd factor device.\n* 'ENROLLED_SECOND_FACTORS': User can use any enabled 2nd factor. Possible values: [\"LOGIN\", \"SECURE_KEY\", \"ENROLLED_SECOND_FACTORS\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "policy_type": {
                    "description": "How IAP determines the effective policy in cases of hierarchical policies.\nPolicies are merged from higher in the hierarchy to lower in the hierarchy.\nThe possible values are:\n\n* 'MINIMUM': This policy acts as a minimum to other policies, lower in the hierarchy.\n\t\t   Effective policy may only be the same or stricter.\n* 'DEFAULT': This policy acts as a default if no other reauth policy is set. Possible values: [\"MINIMUM\", \"DEFAULT\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Settings to configure reauthentication policies in IAP.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "workforce_identity_settings": {
              "block": {
                "attributes": {
                  "workforce_pools": {
                    "description": "The workforce pool resources. Only one workforce pool is accepted.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "oauth2": {
                    "block": {
                      "attributes": {
                        "client_id": {
                          "description": "The OAuth 2.0 client ID registered in the workforce identity\nfederation OAuth 2.0 Server.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "client_secret": {
                          "description": "Input only. The OAuth 2.0 client secret created while registering\nthe client ID.",
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "client_secret_sha256": {
                          "computed": true,
                          "description": "Output only. SHA256 hash value for the client secret. This field\nis returned by IAP when the settings are retrieved.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "description": "OAuth 2.0 settings for IAP to perform OIDC flow with workforce identity\nfederation services.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Settings to configure the workforce identity federation, including workforce pools\nand OAuth 2.0 settings.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Top level wrapper for all access related setting in IAP.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "application_settings": {
        "block": {
          "attributes": {
            "cookie_domain": {
              "description": "The Domain value to set for cookies generated by IAP. This value is not validated by the API,\nbut will be ignored at runtime if invalid.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "access_denied_page_settings": {
              "block": {
                "attributes": {
                  "access_denied_page_uri": {
                    "description": "The URI to be redirected to when access is denied.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "generate_troubleshooting_uri": {
                    "description": "Whether to generate a troubleshooting URL on access denied events to this application.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "remediation_token_generation_enabled": {
                    "description": "Whether to generate remediation token on access denied events to this application.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Customization for Access Denied page. IAP allows customers to define a custom URI\nto use as the error page when access is denied to users. If IAP prevents access\nto this page, the default IAP error page will be displayed instead.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "attribute_propagation_settings": {
              "block": {
                "attributes": {
                  "enable": {
                    "description": "Whether the provided attribute propagation settings should be evaluated on user requests.\nIf set to true, attributes returned from the expression will be propagated in the set output credentials.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "expression": {
                    "description": "Raw string CEL expression. Must return a list of attributes. A maximum of 45 attributes can\nbe selected. Expressions can select different attribute types from attributes:\nattributes.saml_attributes, attributes.iap_attributes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_credentials": {
                    "description": "Which output credentials attributes selected by the CEL expression should be propagated in.\nAll attributes will be fully duplicated in each selected output credential.\nPossible values are:\n\n* 'HEADER': Propagate attributes in the headers with \"x-goog-iap-attr-\" prefix.\n* 'JWT': Propagate attributes in the JWT of the form:\n         \"additional_claims\": { \"my_attribute\": [\"value1\", \"value2\"] }\n* 'RCTOKEN': Propagate attributes in the RCToken of the form: \"\n             additional_claims\": { \"my_attribute\": [\"value1\", \"value2\"] } Possible values: [\"HEADER\", \"JWT\", \"RCTOKEN\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Settings to configure attribute propagation.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "csm_settings": {
              "block": {
                "attributes": {
                  "rctoken_aud": {
                    "description": "Audience claim set in the generated RCToken. This value is not validated by IAP.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Settings to configure IAP's behavior for a service mesh.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Top level wrapper for all application related settings in IAP.",
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

func GoogleIapSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIapSettings), &result)
	return &result
}
