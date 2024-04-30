package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIntegrationsAuthConfig = `{
  "block": {
    "attributes": {
      "certificate_id": {
        "computed": true,
        "description": "Certificate id for client certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp when the auth config is created.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "creator_email": {
        "computed": true,
        "description": "The creator's email address. Generated based on the End User Credentials/LOAS role of the user making the call.",
        "description_kind": "plain",
        "type": "string"
      },
      "credential_type": {
        "computed": true,
        "description": "Credential type of the encrypted credential.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A description of the auth config.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The name of the auth config.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "encrypted_credential": {
        "computed": true,
        "description": "Auth credential encrypted by Cloud KMS. Can be decrypted as Credential with proper KMS key.\n\nA base64-encoded string.",
        "description_kind": "plain",
        "type": "string"
      },
      "expiry_notification_duration": {
        "description": "User can define the time to receive notification after which the auth config becomes invalid. Support up to 30 days. Support granularity in hours.\n\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.5s\".",
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
      "last_modifier_email": {
        "computed": true,
        "description": "The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "Location in which client needs to be provisioned.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Resource name of the auth config.",
        "description_kind": "plain",
        "type": "string"
      },
      "override_valid_time": {
        "description": "User provided expiry time to override. For the example of Salesforce, username/password credentials can be valid for 6 months depending on the instance settings.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reason": {
        "computed": true,
        "description": "The reason / details of the current status.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The status of the auth config.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp when the auth config is modified.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "valid_time": {
        "computed": true,
        "description": "The time until the auth config is valid. Empty or max value is considered the auth config won't expire.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "visibility": {
        "description": "The visibility of the auth config. Possible values: [\"PRIVATE\", \"CLIENT_VISIBLE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "client_certificate": {
        "block": {
          "attributes": {
            "encrypted_private_key": {
              "description": "The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "passphrase": {
              "description": "'passphrase' should be left unset if private key is not encrypted.\nNote that 'passphrase' is not the password for web server, but an extra layer of security to protected private key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_certificate": {
              "description": "The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Raw client certificate",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "decrypted_credential": {
        "block": {
          "attributes": {
            "credential_type": {
              "description": "Credential type associated with auth configs.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "auth_token": {
              "block": {
                "attributes": {
                  "token": {
                    "description": "The token for the auth type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "Authentication type, e.g. \"Basic\", \"Bearer\", etc.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Auth token credential.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "jwt": {
              "block": {
                "attributes": {
                  "jwt": {
                    "computed": true,
                    "description": "The token calculated by the header, payload and signature.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "jwt_header": {
                    "description": "Identifies which algorithm is used to generate the signature.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "jwt_payload": {
                    "description": "Contains a set of claims. The JWT specification defines seven Registered Claim Names which are the standard fields commonly included in tokens. Custom claims are usually also included, depending on the purpose of the token.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secret": {
                    "description": "User's pre-shared secret to sign the token.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "JWT credential.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oauth2_authorization_code": {
              "block": {
                "attributes": {
                  "auth_endpoint": {
                    "description": "The auth url endpoint to send the auth code request to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "client_id": {
                    "description": "The client's id.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "client_secret": {
                    "description": "The client's secret.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scope": {
                    "description": "A space-delimited list of requested scope permissions.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "token_endpoint": {
                    "description": "The token url endpoint to send the token request to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "OAuth2 authorization code credential.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oauth2_client_credentials": {
              "block": {
                "attributes": {
                  "client_id": {
                    "description": "The client's ID.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "client_secret": {
                    "description": "The client's secret.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "request_type": {
                    "description": "Represent how to pass parameters to fetch access token Possible values: [\"REQUEST_TYPE_UNSPECIFIED\", \"REQUEST_BODY\", \"QUERY_PARAMETERS\", \"ENCODED_HEADER\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scope": {
                    "description": "A space-delimited list of requested scope permissions.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "token_endpoint": {
                    "description": "The token endpoint is used by the client to obtain an access token by presenting its authorization grant or refresh token.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "token_params": {
                    "block": {
                      "block_types": {
                        "entries": {
                          "block": {
                            "block_types": {
                              "key": {
                                "block": {
                                  "block_types": {
                                    "literal_value": {
                                      "block": {
                                        "attributes": {
                                          "string_value": {
                                            "description": "String.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Passing a literal value",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Key of the map entry.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "value": {
                                "block": {
                                  "block_types": {
                                    "literal_value": {
                                      "block": {
                                        "attributes": {
                                          "string_value": {
                                            "description": "String.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Passing a literal value",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Value of the map entry.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A list of parameter map entries.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Token parameters for the auth request.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "OAuth2 client credentials.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oidc_token": {
              "block": {
                "attributes": {
                  "audience": {
                    "description": "Audience to be used when generating OIDC token. The audience claim identifies the recipients that the JWT is intended for.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_account_email": {
                    "description": "The service account email to be used as the identity for the token.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "token": {
                    "computed": true,
                    "description": "ID token obtained for the service account.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "token_expire_time": {
                    "computed": true,
                    "description": "The approximate time until the token retrieved is valid.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Google OIDC ID Token.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "service_account_credentials": {
              "block": {
                "attributes": {
                  "scope": {
                    "description": "A space-delimited list of requested scope permissions.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_account": {
                    "description": "Name of the service account that has the permission to make the request.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Service account credential.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "username_and_password": {
              "block": {
                "attributes": {
                  "password": {
                    "description": "Password to be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "username": {
                    "description": "Username to be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Username and password credential.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Raw auth credentials.",
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

func GoogleIntegrationsAuthConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIntegrationsAuthConfig), &result)
	return &result
}
