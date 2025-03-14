package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleArtifactRegistryRepository = `{
  "block": {
    "attributes": {
      "cleanup_policies": {
        "computed": true,
        "description": "Cleanup policies for this repository. Cleanup policies indicate when\ncertain package versions can be automatically deleted.\nMap keys are policy IDs supplied by users during policy creation. They must\nunique within a repository and be under 128 characters in length.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "action": "string",
              "condition": [
                "list",
                [
                  "object",
                  {
                    "newer_than": "string",
                    "older_than": "string",
                    "package_name_prefixes": [
                      "list",
                      "string"
                    ],
                    "tag_prefixes": [
                      "list",
                      "string"
                    ],
                    "tag_state": "string",
                    "version_name_prefixes": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "id": "string",
              "most_recent_versions": [
                "list",
                [
                  "object",
                  {
                    "keep_count": "number",
                    "package_name_prefixes": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "cleanup_policy_dry_run": {
        "computed": true,
        "description": "If true, the cleanup pipeline is prevented from deleting versions in this\nrepository.",
        "description_kind": "plain",
        "type": "bool"
      },
      "create_time": {
        "computed": true,
        "description": "The time when the repository was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The user-provided description of the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "docker_config": {
        "computed": true,
        "description": "Docker repository config contains repository level configuration for the repositories of docker type.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "immutable_tags": "bool"
            }
          ]
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
      "format": {
        "computed": true,
        "description": "The format of packages that are stored in the repository. Supported formats\ncan be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).\nYou can only create alpha formats if you are a member of the\n[alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_name": {
        "computed": true,
        "description": "The Cloud KMS resource name of the customer managed encryption key that’s\nused to encrypt the contents of the Repository. Has the form:\n'projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key'.\nThis value may not be changed after the Repository has been created.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "Labels with user-defined metadata.\nThis field may contain up to 64 entries. Label keys and values may be no\nlonger than 63 characters. Label keys must begin with a lowercase letter\nand may only contain lowercase letters, numeric characters, underscores,\nand dashes.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The name of the repository's location. In addition to specific regions,\nspecial values for multi-region locations are 'asia', 'europe', and 'us'.\nSee [here](https://cloud.google.com/artifact-registry/docs/repositories/repo-locations),\nor use the\n[google_artifact_registry_locations](https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/artifact_registry_locations)\ndata source for possible values.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maven_config": {
        "computed": true,
        "description": "MavenRepositoryConfig is maven related repository details.\nProvides additional configuration details for repositories of the maven\nformat type.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allow_snapshot_overwrites": "bool",
              "version_policy": "string"
            }
          ]
        ]
      },
      "mode": {
        "computed": true,
        "description": "The mode configures the repository to serve artifacts from different sources. Default value: \"STANDARD_REPOSITORY\" Possible values: [\"STANDARD_REPOSITORY\", \"VIRTUAL_REPOSITORY\", \"REMOTE_REPOSITORY\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the repository, for example:\n\"repo1\"",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remote_repository_config": {
        "computed": true,
        "description": "Configuration specific for a Remote Repository.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "apt_repository": [
                "list",
                [
                  "object",
                  {
                    "public_repository": [
                      "list",
                      [
                        "object",
                        {
                          "repository_base": "string",
                          "repository_path": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "common_repository": [
                "list",
                [
                  "object",
                  {
                    "uri": "string"
                  }
                ]
              ],
              "description": "string",
              "disable_upstream_validation": "bool",
              "docker_repository": [
                "list",
                [
                  "object",
                  {
                    "custom_repository": [
                      "list",
                      [
                        "object",
                        {
                          "uri": "string"
                        }
                      ]
                    ],
                    "public_repository": "string"
                  }
                ]
              ],
              "maven_repository": [
                "list",
                [
                  "object",
                  {
                    "custom_repository": [
                      "list",
                      [
                        "object",
                        {
                          "uri": "string"
                        }
                      ]
                    ],
                    "public_repository": "string"
                  }
                ]
              ],
              "npm_repository": [
                "list",
                [
                  "object",
                  {
                    "custom_repository": [
                      "list",
                      [
                        "object",
                        {
                          "uri": "string"
                        }
                      ]
                    ],
                    "public_repository": "string"
                  }
                ]
              ],
              "python_repository": [
                "list",
                [
                  "object",
                  {
                    "custom_repository": [
                      "list",
                      [
                        "object",
                        {
                          "uri": "string"
                        }
                      ]
                    ],
                    "public_repository": "string"
                  }
                ]
              ],
              "upstream_credentials": [
                "list",
                [
                  "object",
                  {
                    "username_password_credentials": [
                      "list",
                      [
                        "object",
                        {
                          "password_secret_version": "string",
                          "username": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "yum_repository": [
                "list",
                [
                  "object",
                  {
                    "public_repository": [
                      "list",
                      [
                        "object",
                        {
                          "repository_base": "string",
                          "repository_path": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "repository_id": {
        "description": "The last part of the repository name, for example:\n\"repo1\"",
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
      "update_time": {
        "computed": true,
        "description": "The time when the repository was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "virtual_repository_config": {
        "computed": true,
        "description": "Configuration specific for a Virtual Repository.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "upstream_policies": [
                "list",
                [
                  "object",
                  {
                    "id": "string",
                    "priority": "number",
                    "repository": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "vulnerability_scanning_config": {
        "computed": true,
        "description": "Configuration for vulnerability scanning of artifacts stored in this repository.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enablement_config": "string",
              "enablement_state": "string",
              "enablement_state_reason": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleArtifactRegistryRepositorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleArtifactRegistryRepository), &result)
	return &result
}
