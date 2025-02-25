package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDeveloperConnectConnection = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Optional. Allows clients to store small amounts of arbitrary data.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "connection_id": {
        "description": "Required. Id of the requesting object\nIf auto-generating Id server-side, remove this field and\nconnection_id from the method_signature of Create RPC",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. [Output only] Create timestamp",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "Output only. [Output only] Delete timestamp",
        "description_kind": "plain",
        "type": "string"
      },
      "disabled": {
        "description": "Optional. If disabled is set to true, functionality is disabled for this connection.\nRepository based API methods and webhooks processing for repositories in\nthis connection will be disabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "effective_annotations": {
        "computed": true,
        "description": "All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
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
      "etag": {
        "description": "Optional. This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the\nclient has an up-to-date value before proceeding.",
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
      "installation_state": {
        "computed": true,
        "description": "Describes stage and necessary actions to be taken by the\nuser to complete the installation. Used for GitHub and GitHub Enterprise\nbased connections.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "action_uri": "string",
              "message": "string",
              "stage": "string"
            }
          ]
        ]
      },
      "labels": {
        "description": "Optional. Labels as key value pairs\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name of the connection, in the format\n'projects/{project}/locations/{location}/connections/{connection_id}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "Output only. Set to true when the connection is being set up or updated in the\nbackground.",
        "description_kind": "plain",
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
      "uid": {
        "computed": true,
        "description": "Output only. A system-assigned unique identifier for a the GitRepositoryLink.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. [Output only] Update timestamp",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "bitbucket_cloud_config": {
        "block": {
          "attributes": {
            "webhook_secret_secret_version": {
              "description": "Required. Immutable. SecretManager resource containing the webhook secret used to verify webhook\nevents, formatted as 'projects/*/secrets/*/versions/*'. This is used to\nvalidate and create webhooks.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "workspace": {
              "description": "Required. The Bitbucket Cloud Workspace ID to be connected to Google Cloud Platform.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "authorizer_credential": {
              "block": {
                "attributes": {
                  "user_token_secret_version": {
                    "description": "Required. A SecretManager resource containing the user token that authorizes\nthe Developer Connect connection. Format:\n'projects/*/secrets/*/versions/*'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description": "Output only. The username associated with this token.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Represents a personal access token that authorized the Connection,\nand associated metadata.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "read_authorizer_credential": {
              "block": {
                "attributes": {
                  "user_token_secret_version": {
                    "description": "Required. A SecretManager resource containing the user token that authorizes\nthe Developer Connect connection. Format:\n'projects/*/secrets/*/versions/*'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description": "Output only. The username associated with this token.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Represents a personal access token that authorized the Connection,\nand associated metadata.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for connections to an instance of Bitbucket Cloud.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "bitbucket_data_center_config": {
        "block": {
          "attributes": {
            "host_uri": {
              "description": "Required. The URI of the Bitbucket Data Center host this connection is for.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "server_version": {
              "computed": true,
              "description": "Output only. Version of the Bitbucket Data Center server running on the 'host_uri'.",
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_ca_certificate": {
              "description": "Optional. SSL certificate authority to trust when making requests to Bitbucket Data\nCenter.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "webhook_secret_secret_version": {
              "description": "Required. Immutable. SecretManager resource containing the webhook secret used to verify webhook\nevents, formatted as 'projects/*/secrets/*/versions/*'. This is used to\nvalidate webhooks.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "authorizer_credential": {
              "block": {
                "attributes": {
                  "user_token_secret_version": {
                    "description": "Required. A SecretManager resource containing the user token that authorizes\nthe Developer Connect connection. Format:\n'projects/*/secrets/*/versions/*'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description": "Output only. The username associated with this token.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Represents a personal access token that authorized the Connection,\nand associated metadata.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "read_authorizer_credential": {
              "block": {
                "attributes": {
                  "user_token_secret_version": {
                    "description": "Required. A SecretManager resource containing the user token that authorizes\nthe Developer Connect connection. Format:\n'projects/*/secrets/*/versions/*'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description": "Output only. The username associated with this token.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Represents a personal access token that authorized the Connection,\nand associated metadata.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "service_directory_config": {
              "block": {
                "attributes": {
                  "service": {
                    "description": "Required. The Service Directory service name.\nFormat:\nprojects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "ServiceDirectoryConfig represents Service Directory configuration for a\nconnection.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for connections to an instance of Bitbucket Data Center.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "crypto_key_config": {
        "block": {
          "attributes": {
            "key_reference": {
              "description": "Required. The name of the key which is used to encrypt/decrypt customer data. For key\nin Cloud KMS, the key should be in the format of\n'projects/*/locations/*/keyRings/*/cryptoKeys/*'.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The crypto key configuration. This field is used by the Customer-managed\nencryption keys (CMEK) feature.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "github_config": {
        "block": {
          "attributes": {
            "app_installation_id": {
              "computed": true,
              "description": "Optional. GitHub App installation id.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "github_app": {
              "description": "Required. Immutable. The GitHub Application that was installed to the GitHub user or\norganization.\nPossible values:\nGIT_HUB_APP_UNSPECIFIED\nDEVELOPER_CONNECT\nFIREBASE",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "installation_uri": {
              "computed": true,
              "description": "Output only. The URI to navigate to in order to manage the installation associated\nwith this GitHubConfig.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "authorizer_credential": {
              "block": {
                "attributes": {
                  "oauth_token_secret_version": {
                    "description": "Required. A SecretManager resource containing the OAuth token that authorizes\nthe connection. Format: 'projects/*/secrets/*/versions/*'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description": "Output only. The username associated with this token.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Represents an OAuth token of the account that authorized the Connection,\nand associated metadata.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for connections to github.com.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "github_enterprise_config": {
        "block": {
          "attributes": {
            "app_id": {
              "description": "Optional. ID of the GitHub App created from the manifest.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "app_installation_id": {
              "description": "Optional. ID of the installation of the GitHub App.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "app_slug": {
              "computed": true,
              "description": "Output only. The URL-friendly name of the GitHub App.",
              "description_kind": "plain",
              "type": "string"
            },
            "host_uri": {
              "description": "Required. The URI of the GitHub Enterprise host this connection is for.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "installation_uri": {
              "computed": true,
              "description": "Output only. The URI to navigate to in order to manage the installation associated\nwith this GitHubEnterpriseConfig.",
              "description_kind": "plain",
              "type": "string"
            },
            "private_key_secret_version": {
              "description": "Optional. SecretManager resource containing the private key of the GitHub App,\nformatted as 'projects/*/secrets/*/versions/*'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "server_version": {
              "computed": true,
              "description": "Output only. GitHub Enterprise version installed at the host_uri.",
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_ca_certificate": {
              "description": "Optional. SSL certificate to use for requests to GitHub Enterprise.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "webhook_secret_secret_version": {
              "description": "Optional. SecretManager resource containing the webhook secret of the GitHub App,\nformatted as 'projects/*/secrets/*/versions/*'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "service_directory_config": {
              "block": {
                "attributes": {
                  "service": {
                    "description": "Required. The Service Directory service name.\nFormat:\nprojects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "ServiceDirectoryConfig represents Service Directory configuration for a\nconnection.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for connections to an instance of GitHub Enterprise.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "gitlab_config": {
        "block": {
          "attributes": {
            "webhook_secret_secret_version": {
              "description": "Required. Immutable. SecretManager resource containing the webhook secret of a GitLab project,\nformatted as 'projects/*/secrets/*/versions/*'. This is used to validate\nwebhooks.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "authorizer_credential": {
              "block": {
                "attributes": {
                  "user_token_secret_version": {
                    "description": "Required. A SecretManager resource containing the user token that authorizes\nthe Developer Connect connection. Format:\n'projects/*/secrets/*/versions/*'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description": "Output only. The username associated with this token.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Represents a personal access token that authorized the Connection,\nand associated metadata.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "read_authorizer_credential": {
              "block": {
                "attributes": {
                  "user_token_secret_version": {
                    "description": "Required. A SecretManager resource containing the user token that authorizes\nthe Developer Connect connection. Format:\n'projects/*/secrets/*/versions/*'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description": "Output only. The username associated with this token.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Represents a personal access token that authorized the Connection,\nand associated metadata.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for connections to gitlab.com.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "gitlab_enterprise_config": {
        "block": {
          "attributes": {
            "host_uri": {
              "description": "Required. The URI of the GitLab Enterprise host this connection is for.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "server_version": {
              "computed": true,
              "description": "Output only. Version of the GitLab Enterprise server running on the 'host_uri'.",
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_ca_certificate": {
              "description": "Optional. SSL Certificate Authority certificate to use for requests to GitLab\nEnterprise instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "webhook_secret_secret_version": {
              "description": "Required. Immutable. SecretManager resource containing the webhook secret of a GitLab project,\nformatted as 'projects/*/secrets/*/versions/*'. This is used to validate\nwebhooks.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "authorizer_credential": {
              "block": {
                "attributes": {
                  "user_token_secret_version": {
                    "description": "Required. A SecretManager resource containing the user token that authorizes\nthe Developer Connect connection. Format:\n'projects/*/secrets/*/versions/*'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description": "Output only. The username associated with this token.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Represents a personal access token that authorized the Connection,\nand associated metadata.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "read_authorizer_credential": {
              "block": {
                "attributes": {
                  "user_token_secret_version": {
                    "description": "Required. A SecretManager resource containing the user token that authorizes\nthe Developer Connect connection. Format:\n'projects/*/secrets/*/versions/*'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description": "Output only. The username associated with this token.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Represents a personal access token that authorized the Connection,\nand associated metadata.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "service_directory_config": {
              "block": {
                "attributes": {
                  "service": {
                    "description": "Required. The Service Directory service name.\nFormat:\nprojects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "ServiceDirectoryConfig represents Service Directory configuration for a\nconnection.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for connections to an instance of GitLab Enterprise.",
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

func GoogleDeveloperConnectConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDeveloperConnectConnection), &result)
	return &result
}
