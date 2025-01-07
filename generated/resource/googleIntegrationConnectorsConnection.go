package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIntegrationConnectorsConnection = `{
  "block": {
    "attributes": {
      "connection_revision": {
        "computed": true,
        "description": "Connection revision. This field is only updated when the connection is created or updated by User.",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_version": {
        "description": "connectorVersion of the Connector.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connector_version_infra_config": {
        "computed": true,
        "description": "This configuration provides infra configs like rate limit threshold which need to be configurable for every connector version.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ratelimit_threshold": "string"
            }
          ]
        ]
      },
      "connector_version_launch_stage": {
        "computed": true,
        "description": "Flag to mark the version indicating the launch stage.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Time the Namespace was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An arbitrary description for the Connection.",
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
      "eventing_enablement_type": {
        "description": "Eventing enablement type. Will be nil if eventing is not enabled. Possible values: [\"EVENTING_AND_CONNECTION\", \"ONLY_EVENTING\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "eventing_runtime_data": {
        "computed": true,
        "description": "Eventing Runtime Data.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "events_listener_endpoint": "string",
              "status": [
                "list",
                [
                  "object",
                  {
                    "description": "string",
                    "state": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Resource labels to represent user provided metadata.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Location in which Connection needs to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Name of Connection needs to be created.",
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
      "service_account": {
        "computed": true,
        "description": "Service account needed for runtime plane to access Google Cloud resources.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_directory": {
        "computed": true,
        "description": "The name of the Service Directory service name. Used for Private Harpoon to resolve the ILB address.\ne.g. \"projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors\"",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Status of the Integration Connector.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "state": "string",
              "status": "string"
            }
          ]
        ]
      },
      "subscription_type": {
        "computed": true,
        "description": "This subscription type enum states the subscription type of the project.",
        "description_kind": "plain",
        "type": "string"
      },
      "suspended": {
        "description": "Suspended indicates if a user has suspended a connection or not.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "update_time": {
        "computed": true,
        "description": "Time the Namespace was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "auth_config": {
        "block": {
          "attributes": {
            "auth_key": {
              "description": "The type of authentication configured.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "auth_type": {
              "description": "authType of the Connection Possible values: [\"AUTH_TYPE_UNSPECIFIED\", \"USER_PASSWORD\", \"OAUTH2_JWT_BEARER\", \"OAUTH2_CLIENT_CREDENTIALS\", \"SSH_PUBLIC_KEY\", \"OAUTH2_AUTH_CODE_FLOW\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "additional_variable": {
              "block": {
                "attributes": {
                  "boolean_value": {
                    "description": "Boolean Value of configVariable.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "integer_value": {
                    "description": "Integer Value of configVariable.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "key": {
                    "description": "Key for the configVariable",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "string_value": {
                    "description": "String Value of configVariabley.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "encryption_key_value": {
                    "block": {
                      "attributes": {
                        "kms_key_name": {
                          "description": "The [KMS key name] with which the content of the Operation is encrypted. The\nexpected format: projects/*/locations/*/keyRings/*/cryptoKeys/*.\nWill be empty string if google managed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description": "Type of Encryption Key Possible values: [\"GOOGLE_MANAGED\", \"CUSTOMER_MANAGED\"]",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Encryption key value of configVariable.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "secret_value": {
                    "block": {
                      "attributes": {
                        "secret_version": {
                          "description": "Secret version of Secret Value for Config variable.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Secret value of configVariable.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List containing additional auth configs.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "oauth2_auth_code_flow": {
              "block": {
                "attributes": {
                  "auth_uri": {
                    "description": "Auth URL for Authorization Code Flow.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "client_id": {
                    "description": "Client ID for user-provided OAuth app.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_pkce": {
                    "description": "Whether to enable PKCE when the user performs the auth code flow.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "scopes": {
                    "description": "Scopes the connection will request when the user performs the auth code flow.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "client_secret": {
                    "block": {
                      "attributes": {
                        "secret_version": {
                          "description": "The resource name of the secret version in the format,\nformat as: projects/*/secrets/*/versions/*.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Client secret for user-provided OAuth app.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Parameters to support Oauth 2.0 Auth Code Grant Authentication.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oauth2_client_credentials": {
              "block": {
                "attributes": {
                  "client_id": {
                    "description": "Secret version of Password for Authentication.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "client_secret": {
                    "block": {
                      "attributes": {
                        "secret_version": {
                          "description": "The resource name of the secret version in the format,\nformat as: projects/*/secrets/*/versions/*.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Secret version reference containing the client secret.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "OAuth3 Client Credentials for Authentication.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oauth2_jwt_bearer": {
              "block": {
                "block_types": {
                  "client_key": {
                    "block": {
                      "attributes": {
                        "secret_version": {
                          "description": "The resource name of the secret version in the format,\nformat as: projects/*/secrets/*/versions/*.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Secret version reference containing a PKCS#8 PEM-encoded private key associated with the Client Certificate.\nThis private key will be used to sign JWTs used for the jwt-bearer authorization grant.\nSpecified in the form as: projects/*/secrets/*/versions/*.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "jwt_claims": {
                    "block": {
                      "attributes": {
                        "audience": {
                          "description": "Value for the \"aud\" claim.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "issuer": {
                          "description": "Value for the \"iss\" claim.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "subject": {
                          "description": "Value for the \"sub\" claim.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "JwtClaims providers fields to generate the token.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "OAuth2 JWT Bearer for Authentication.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ssh_public_key": {
              "block": {
                "attributes": {
                  "cert_type": {
                    "description": "Format of SSH Client cert.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "username": {
                    "description": "The user account used to authenticate.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "ssh_client_cert": {
                    "block": {
                      "attributes": {
                        "secret_version": {
                          "description": "The resource name of the secret version in the format,\nformat as: projects/*/secrets/*/versions/*.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "SSH Client Cert. It should contain both public and private key.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "ssh_client_cert_pass": {
                    "block": {
                      "attributes": {
                        "secret_version": {
                          "description": "The resource name of the secret version in the format,\nformat as: projects/*/secrets/*/versions/*.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Password (passphrase) for ssh client certificate if it has one.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "SSH Public Key for Authentication.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "user_password": {
              "block": {
                "attributes": {
                  "username": {
                    "description": "Username for Authentication.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "password": {
                    "block": {
                      "attributes": {
                        "secret_version": {
                          "description": "The resource name of the secret version in the format,\nformat as: projects/*/secrets/*/versions/*.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Password for Authentication.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "User password for Authentication.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "authConfig for the connection.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "config_variable": {
        "block": {
          "attributes": {
            "boolean_value": {
              "description": "Boolean Value of configVariable",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "integer_value": {
              "description": "Integer Value of configVariable",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "key": {
              "description": "Key for the configVariable",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "string_value": {
              "description": "String Value of configVariabley",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "encryption_key_value": {
              "block": {
                "attributes": {
                  "kms_key_name": {
                    "description": "The [KMS key name] with which the content of the Operation is encrypted. The\nexpected format: projects/*/locations/*/keyRings/*/cryptoKeys/*.\nWill be empty string if google managed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "Type of Encryption Key Possible values: [\"GOOGLE_MANAGED\", \"CUSTOMER_MANAGED\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Encryption key value of configVariable.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "secret_value": {
              "block": {
                "attributes": {
                  "secret_version": {
                    "description": "Secret version of Secret Value for Config variable.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Secret value of configVariable.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Config Variables for the connection.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "destination_config": {
        "block": {
          "attributes": {
            "key": {
              "description": "The key is the destination identifier that is supported by the Connector.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "destination": {
              "block": {
                "attributes": {
                  "host": {
                    "description": "For publicly routable host.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description": "The port is the target port number that is accepted by the destination.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "service_attachment": {
                    "description": "PSC service attachments. Format: projects/*/regions/*/serviceAttachments/*",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The destinations for the key.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Define the Connectors target endpoint.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "eventing_config": {
        "block": {
          "attributes": {
            "enrichment_enabled": {
              "description": "Enrichment Enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "additional_variable": {
              "block": {
                "attributes": {
                  "boolean_value": {
                    "description": "Boolean Value of configVariable.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "integer_value": {
                    "description": "Integer Value of configVariable.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "key": {
                    "description": "Key for the configVariable",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "string_value": {
                    "description": "String Value of configVariabley.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "encryption_key_value": {
                    "block": {
                      "attributes": {
                        "kms_key_name": {
                          "description": "The [KMS key name] with which the content of the Operation is encrypted. The\nexpected format: projects/*/locations/*/keyRings/*/cryptoKeys/*.\nWill be empty string if google managed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description": "Type of Encryption Key Possible values: [\"GOOGLE_MANAGED\", \"CUSTOMER_MANAGED\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Encryption key value of configVariable.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "secret_value": {
                    "block": {
                      "attributes": {
                        "secret_version": {
                          "description": "Secret version of Secret Value for Config variable.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Secret value of configVariable",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List containing additional auth configs.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "auth_config": {
              "block": {
                "attributes": {
                  "auth_key": {
                    "description": "The type of authentication configured.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "auth_type": {
                    "description": "authType of the Connection Possible values: [\"USER_PASSWORD\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "additional_variable": {
                    "block": {
                      "attributes": {
                        "boolean_value": {
                          "description": "Boolean Value of configVariable.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "integer_value": {
                          "description": "Integer Value of configVariable.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "key": {
                          "description": "Key for the configVariable",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "string_value": {
                          "description": "String Value of configVariabley.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "encryption_key_value": {
                          "block": {
                            "attributes": {
                              "kms_key_name": {
                                "description": "The [KMS key name] with which the content of the Operation is encrypted. The\nexpected format: projects/*/locations/*/keyRings/*/cryptoKeys/*.\nWill be empty string if google managed.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "Type of Encryption Key Possible values: [\"GOOGLE_MANAGED\", \"CUSTOMER_MANAGED\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Encryption key value of configVariable",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "secret_value": {
                          "block": {
                            "attributes": {
                              "secret_version": {
                                "description": "Secret version of Secret Value for Config variable.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Secret value of configVariable",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "List containing additional auth configs.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "user_password": {
                    "block": {
                      "attributes": {
                        "username": {
                          "description": "Username for Authentication.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "password": {
                          "block": {
                            "attributes": {
                              "secret_version": {
                                "description": "The resource name of the secret version in the format,\nformat as: projects/*/secrets/*/versions/*.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Password for Authentication.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "User password for Authentication.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "authConfig for Eventing Configuration.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "registration_destination_config": {
              "block": {
                "attributes": {
                  "key": {
                    "description": "Key for the connection",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "destination": {
                    "block": {
                      "attributes": {
                        "host": {
                          "description": "Host",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "port": {
                          "description": "port number",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "service_attachment": {
                          "description": "Service Attachment",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "destinations for the connection",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "registrationDestinationConfig",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Eventing Configuration of a connection",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "lock_config": {
        "block": {
          "attributes": {
            "locked": {
              "description": "Indicates whether or not the connection is locked.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "reason": {
              "description": "Describes why a connection is locked.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Determines whether or no a connection is locked. If locked, a reason must be specified.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "log_config": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Enabled represents whether logging is enabled or not for a connection.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Log configuration for the connection.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "node_config": {
        "block": {
          "attributes": {
            "max_node_count": {
              "computed": true,
              "description": "Minimum number of nodes in the runtime nodes.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_node_count": {
              "computed": true,
              "description": "Minimum number of nodes in the runtime nodes.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Node configuration for the connection.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ssl_config": {
        "block": {
          "attributes": {
            "client_cert_type": {
              "description": "Type of Client Cert (PEM/JKS/.. etc.) Possible values: [\"PEM\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "server_cert_type": {
              "description": "Type of Server Cert (PEM/JKS/.. etc.) Possible values: [\"PEM\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "trust_model": {
              "description": "Enum for Trust Model Possible values: [\"PUBLIC\", \"PRIVATE\", \"INSECURE\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "Enum for controlling the SSL Type (TLS/MTLS) Possible values: [\"TLS\", \"MTLS\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "use_ssl": {
              "description": "Bool for enabling SSL",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "additional_variable": {
              "block": {
                "attributes": {
                  "boolean_value": {
                    "description": "Boolean Value of configVariable.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "integer_value": {
                    "description": "Integer Value of configVariable.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "key": {
                    "description": "Key for the configVariable",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "string_value": {
                    "description": "String Value of configVariabley.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "encryption_key_value": {
                    "block": {
                      "attributes": {
                        "kms_key_name": {
                          "description": "The [KMS key name] with which the content of the Operation is encrypted. The\nexpected format: projects/*/locations/*/keyRings/*/cryptoKeys/*.\nWill be empty string if google managed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description": "Type of Encryption Key Possible values: [\"GOOGLE_MANAGED\", \"CUSTOMER_MANAGED\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Encryption key value of configVariable",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "secret_value": {
                    "block": {
                      "attributes": {
                        "secret_version": {
                          "description": "Secret version of Secret Value for Config variable.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Secret value of configVariable",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Additional SSL related field values.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "client_certificate": {
              "block": {
                "attributes": {
                  "secret_version": {
                    "description": "Secret version of Secret Value for Config variable.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Client Certificate",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "client_private_key": {
              "block": {
                "attributes": {
                  "secret_version": {
                    "description": "Secret version of Secret Value for Config variable.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Client Private Key",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "client_private_key_pass": {
              "block": {
                "attributes": {
                  "secret_version": {
                    "description": "Secret version of Secret Value for Config variable.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Secret containing the passphrase protecting the Client Private Key",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "private_server_certificate": {
              "block": {
                "attributes": {
                  "secret_version": {
                    "description": "Secret version of Secret Value for Config variable.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Private Server Certificate. Needs to be specified if trust model is PRIVATE.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "SSL Configuration of a connection",
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

func GoogleIntegrationConnectorsConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIntegrationConnectorsConnection), &result)
	return &result
}
