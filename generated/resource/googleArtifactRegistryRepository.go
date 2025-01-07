package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleArtifactRegistryRepository = `{
  "block": {
    "attributes": {
      "cleanup_policy_dry_run": {
        "description": "If true, the cleanup pipeline is prevented from deleting versions in this\nrepository.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "create_time": {
        "computed": true,
        "description": "The time when the repository was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "The user-provided description of the repository.",
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
      "format": {
        "description": "The format of packages that are stored in the repository. Supported formats\ncan be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).\nYou can only create alpha formats if you are a member of the\n[alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).",
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
      "kms_key_name": {
        "description": "The Cloud KMS resource name of the customer managed encryption key that’s\nused to encrypt the contents of the Repository. Has the form:\n'projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key'.\nThis value may not be changed after the Repository has been created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels with user-defined metadata.\nThis field may contain up to 64 entries. Label keys and values may be no\nlonger than 63 characters. Label keys must begin with a lowercase letter\nand may only contain lowercase letters, numeric characters, underscores,\nand dashes.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "computed": true,
        "description": "The name of the repository's location. In addition to specific regions,\nspecial values for multi-region locations are 'asia', 'europe', and 'us'.\nSee [here](https://cloud.google.com/artifact-registry/docs/repositories/repo-locations),\nor use the\n[google_artifact_registry_locations](https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/artifact_registry_locations)\ndata source for possible values.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mode": {
        "description": "The mode configures the repository to serve artifacts from different sources. Default value: \"STANDARD_REPOSITORY\" Possible values: [\"STANDARD_REPOSITORY\", \"VIRTUAL_REPOSITORY\", \"REMOTE_REPOSITORY\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the repository, for example:\n\"repo1\"",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      }
    },
    "block_types": {
      "cleanup_policies": {
        "block": {
          "attributes": {
            "action": {
              "description": "Policy action. Possible values: [\"DELETE\", \"KEEP\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "condition": {
              "block": {
                "attributes": {
                  "newer_than": {
                    "description": "Match versions newer than a duration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "older_than": {
                    "description": "Match versions older than a duration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "package_name_prefixes": {
                    "description": "Match versions by package prefix. Applied on any prefix match.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "tag_prefixes": {
                    "description": "Match versions by tag prefix. Applied on any prefix match.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "tag_state": {
                    "description": "Match versions by tag status. Default value: \"ANY\" Possible values: [\"TAGGED\", \"UNTAGGED\", \"ANY\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version_name_prefixes": {
                    "description": "Match versions by version name prefix. Applied on any prefix match.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Policy condition for matching versions.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "most_recent_versions": {
              "block": {
                "attributes": {
                  "keep_count": {
                    "description": "Minimum number of versions to keep.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "package_name_prefixes": {
                    "description": "Match versions by package prefix. Applied on any prefix match.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Policy condition for retaining a minimum number of versions. May only be\nspecified with a Keep action.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Cleanup policies for this repository. Cleanup policies indicate when\ncertain package versions can be automatically deleted.\nMap keys are policy IDs supplied by users during policy creation. They must\nunique within a repository and be under 128 characters in length.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "docker_config": {
        "block": {
          "attributes": {
            "immutable_tags": {
              "description": "The repository which enabled this flag prevents all tags from being modified, moved or deleted. This does not prevent tags from being created.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Docker repository config contains repository level configuration for the repositories of docker type.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "maven_config": {
        "block": {
          "attributes": {
            "allow_snapshot_overwrites": {
              "description": "The repository with this flag will allow publishing the same\nsnapshot versions.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "version_policy": {
              "description": "Version policy defines the versions that the registry will accept. Default value: \"VERSION_POLICY_UNSPECIFIED\" Possible values: [\"VERSION_POLICY_UNSPECIFIED\", \"RELEASE\", \"SNAPSHOT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "MavenRepositoryConfig is maven related repository details.\nProvides additional configuration details for repositories of the maven\nformat type.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "remote_repository_config": {
        "block": {
          "attributes": {
            "description": {
              "description": "The description of the remote source.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disable_upstream_validation": {
              "description": "If true, the remote repository upstream and upstream credentials will\nnot be validated.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "apt_repository": {
              "block": {
                "block_types": {
                  "public_repository": {
                    "block": {
                      "attributes": {
                        "repository_base": {
                          "description": "A common public repository base for Apt, e.g. '\"debian/dists/buster\"' Possible values: [\"DEBIAN\", \"UBUNTU\"]",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "repository_path": {
                          "description": "Specific repository from the base.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "One of the publicly available Apt repositories supported by Artifact Registry.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specific settings for an Apt remote repository.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "common_repository": {
              "block": {
                "attributes": {
                  "uri": {
                    "description": "One of:\na. Artifact Registry Repository resource, e.g. 'projects/UPSTREAM_PROJECT_ID/locations/REGION/repositories/UPSTREAM_REPOSITORY'\nb. URI to the registry, e.g. '\"https://registry-1.docker.io\"'\nc. URI to Artifact Registry Repository, e.g. '\"https://REGION-docker.pkg.dev/UPSTREAM_PROJECT_ID/UPSTREAM_REPOSITORY\"'",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Specific settings for an Artifact Registory remote repository.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "docker_repository": {
              "block": {
                "attributes": {
                  "public_repository": {
                    "description": "Address of the remote repository. Default value: \"DOCKER_HUB\" Possible values: [\"DOCKER_HUB\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "custom_repository": {
                    "block": {
                      "attributes": {
                        "uri": {
                          "description": "Specific uri to the registry, e.g. '\"https://registry-1.docker.io\"'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "[Deprecated, please use commonRepository instead] Settings for a remote repository with a custom uri.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specific settings for a Docker remote repository.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "maven_repository": {
              "block": {
                "attributes": {
                  "public_repository": {
                    "description": "Address of the remote repository. Default value: \"MAVEN_CENTRAL\" Possible values: [\"MAVEN_CENTRAL\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "custom_repository": {
                    "block": {
                      "attributes": {
                        "uri": {
                          "description": "Specific uri to the registry, e.g. '\"https://repo.maven.apache.org/maven2\"'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "[Deprecated, please use commonRepository instead] Settings for a remote repository with a custom uri.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specific settings for a Maven remote repository.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "npm_repository": {
              "block": {
                "attributes": {
                  "public_repository": {
                    "description": "Address of the remote repository. Default value: \"NPMJS\" Possible values: [\"NPMJS\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "custom_repository": {
                    "block": {
                      "attributes": {
                        "uri": {
                          "description": "Specific uri to the registry, e.g. '\"https://registry.npmjs.org\"'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "[Deprecated, please use commonRepository instead] Settings for a remote repository with a custom uri.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specific settings for an Npm remote repository.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "python_repository": {
              "block": {
                "attributes": {
                  "public_repository": {
                    "description": "Address of the remote repository. Default value: \"PYPI\" Possible values: [\"PYPI\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "custom_repository": {
                    "block": {
                      "attributes": {
                        "uri": {
                          "description": "Specific uri to the registry, e.g. '\"https://pypi.io\"'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "[Deprecated, please use commonRepository instead] Settings for a remote repository with a custom uri.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specific settings for a Python remote repository.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "upstream_credentials": {
              "block": {
                "block_types": {
                  "username_password_credentials": {
                    "block": {
                      "attributes": {
                        "password_secret_version": {
                          "description": "The Secret Manager key version that holds the password to access the\nremote repository. Must be in the format of\n'projects/{project}/secrets/{secret}/versions/{version}'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "username": {
                          "description": "The username to access the remote repository.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Use username and password to access the remote repository.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The credentials used to access the remote repository.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "yum_repository": {
              "block": {
                "block_types": {
                  "public_repository": {
                    "block": {
                      "attributes": {
                        "repository_base": {
                          "description": "A common public repository base for Yum. Possible values: [\"CENTOS\", \"CENTOS_DEBUG\", \"CENTOS_VAULT\", \"CENTOS_STREAM\", \"ROCKY\", \"EPEL\"]",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "repository_path": {
                          "description": "Specific repository from the base, e.g. '\"pub/rocky/9/BaseOS/x86_64/os\"'",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "One of the publicly available Yum repositories supported by Artifact Registry.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specific settings for an Yum remote repository.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration specific for a Remote Repository.",
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
      },
      "virtual_repository_config": {
        "block": {
          "block_types": {
            "upstream_policies": {
              "block": {
                "attributes": {
                  "id": {
                    "description": "The user-provided ID of the upstream policy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "priority": {
                    "description": "Entries with a greater priority value take precedence in the pull order.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "repository": {
                    "description": "A reference to the repository resource, for example:\n\"projects/p1/locations/us-central1/repository/repo1\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Policies that configure the upstream artifacts distributed by the Virtual\nRepository. Upstream policies cannot be set on a standard repository.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Configuration specific for a Virtual Repository.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "vulnerability_scanning_config": {
        "block": {
          "attributes": {
            "enablement_config": {
              "description": "This configures whether vulnerability scanning is automatically performed for artifacts pushed to this repository. Possible values: [\"INHERITED\", \"DISABLED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enablement_state": {
              "computed": true,
              "description": "This field returns whether scanning is active for this repository.",
              "description_kind": "plain",
              "type": "string"
            },
            "enablement_state_reason": {
              "computed": true,
              "description": "This provides an explanation for the state of scanning on this repository.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Configuration for vulnerability scanning of artifacts stored in this repository.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
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
