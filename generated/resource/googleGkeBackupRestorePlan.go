package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeBackupRestorePlan = `{
  "block": {
    "attributes": {
      "backup_plan": {
        "description": "A reference to the BackupPlan from which Backups may be used\nas the source for Restores created via this RestorePlan.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster": {
        "description": "The source cluster from which Restores will be created via this RestorePlan.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "User specified descriptive string for this RestorePlan.",
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
      "labels": {
        "description": "Description: A set of custom labels supplied by the user.\nA list of key-\u003evalue pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The region of the Restore Plan.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The full name of the BackupPlan Resource.",
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
      "state": {
        "computed": true,
        "description": "The State of the RestorePlan.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_reason": {
        "computed": true,
        "description": "Detailed description of why RestorePlan is in its current state.",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Server generated, unique identifier of UUID format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "restore_config": {
        "block": {
          "attributes": {
            "all_namespaces": {
              "description": "If True, restore all namespaced resources in the Backup.\nSetting this field to False will result in an error.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "cluster_resource_conflict_policy": {
              "description": "Defines the behavior for handling the situation where cluster-scoped resources\nbeing restored already exist in the target cluster.\nThis MUST be set to a value other than 'CLUSTER_RESOURCE_CONFLICT_POLICY_UNSPECIFIED'\nif 'clusterResourceRestoreScope' is anyting other than 'noGroupKinds'.\nSee https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/RestoreConfig#clusterresourceconflictpolicy\nfor more information on each policy option. Possible values: [\"USE_EXISTING_VERSION\", \"USE_BACKUP_VERSION\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespaced_resource_restore_mode": {
              "description": "Defines the behavior for handling the situation where sets of namespaced resources\nbeing restored already exist in the target cluster.\nThis MUST be set to a value other than 'NAMESPACED_RESOURCE_RESTORE_MODE_UNSPECIFIED'\nif the 'namespacedResourceRestoreScope' is anything other than 'noNamespaces'.\nSee https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/RestoreConfig#namespacedresourcerestoremode\nfor more information on each mode. Possible values: [\"DELETE_AND_RESTORE\", \"FAIL_ON_CONFLICT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "no_namespaces": {
              "description": "Do not restore any namespaced resources if set to \"True\".\nSpecifying this field to \"False\" is not allowed.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "volume_data_restore_policy": {
              "description": "Specifies the mechanism to be used to restore volume data.\nThis should be set to a value other than 'NAMESPACED_RESOURCE_RESTORE_MODE_UNSPECIFIED'\nif the 'namespacedResourceRestoreScope' is anything other than 'noNamespaces'.\nIf not specified, it will be treated as 'NO_VOLUME_DATA_RESTORATION'.\nSee https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/RestoreConfig#VolumeDataRestorePolicy\nfor more information on each policy option. Possible values: [\"RESTORE_VOLUME_DATA_FROM_BACKUP\", \"REUSE_VOLUME_HANDLE_FROM_BACKUP\", \"NO_VOLUME_DATA_RESTORATION\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "cluster_resource_restore_scope": {
              "block": {
                "attributes": {
                  "all_group_kinds": {
                    "description": "If True, all valid cluster-scoped resources will be restored.\nMutually exclusive to any other field in 'clusterResourceRestoreScope'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "no_group_kinds": {
                    "description": "If True, no cluster-scoped resources will be restored.\nMutually exclusive to any other field in 'clusterResourceRestoreScope'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "excluded_group_kinds": {
                    "block": {
                      "attributes": {
                        "resource_group": {
                          "description": "API Group string of a Kubernetes resource, e.g.\n\"apiextensions.k8s.io\", \"storage.k8s.io\", etc.\nUse empty string for core group.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "resource_kind": {
                          "description": "Kind of a Kubernetes resource, e.g.\n\"CustomResourceDefinition\", \"StorageClass\", etc.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "A list of cluster-scoped resource group kinds to NOT restore from the backup.\nIf specified, all valid cluster-scoped resources will be restored except\nfor those specified in the list.\nMutually exclusive to any other field in 'clusterResourceRestoreScope'.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "selected_group_kinds": {
                    "block": {
                      "attributes": {
                        "resource_group": {
                          "description": "API Group string of a Kubernetes resource, e.g.\n\"apiextensions.k8s.io\", \"storage.k8s.io\", etc.\nUse empty string for core group.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "resource_kind": {
                          "description": "Kind of a Kubernetes resource, e.g.\n\"CustomResourceDefinition\", \"StorageClass\", etc.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "A list of cluster-scoped resource group kinds to restore from the backup.\nIf specified, only the selected resources will be restored.\nMutually exclusive to any other field in the 'clusterResourceRestoreScope'.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Identifies the cluster-scoped resources to restore from the Backup.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "excluded_namespaces": {
              "block": {
                "attributes": {
                  "namespaces": {
                    "description": "A list of Kubernetes Namespaces.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "A list of selected namespaces excluded from restoration.\nAll namespaces except those in this list will be restored.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "selected_applications": {
              "block": {
                "block_types": {
                  "namespaced_names": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "The name of a Kubernetes Resource.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "namespace": {
                          "description": "The namespace of a Kubernetes Resource.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "A list of namespaced Kubernetes resources.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A list of selected ProtectedApplications to restore.\nThe listed ProtectedApplications and all the resources\nto which they refer will be restored.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "selected_namespaces": {
              "block": {
                "attributes": {
                  "namespaces": {
                    "description": "A list of Kubernetes Namespaces.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "A list of selected namespaces to restore from the Backup.\nThe listed Namespaces and all resources contained in them will be restored.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "transformation_rules": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "The description is a user specified string description\nof the transformation rule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "field_actions": {
                    "block": {
                      "attributes": {
                        "from_path": {
                          "description": "A string containing a JSON Pointer value that references the\nlocation in the target document to move the value from.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "op": {
                          "description": "Specifies the operation to perform. Possible values: [\"REMOVE\", \"MOVE\", \"COPY\", \"ADD\", \"TEST\", \"REPLACE\"]",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "path": {
                          "description": "A string containing a JSON-Pointer value that references a\nlocation within the target document where the operation is performed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "A string that specifies the desired value in string format\nto use for transformation.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "A list of transformation rule actions to take against candidate\nresources. Actions are executed in order defined - this order\nmatters, as they could potentially interfere with each other and\nthe first operation could affect the outcome of the second operation.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "resource_filter": {
                    "block": {
                      "attributes": {
                        "json_path": {
                          "description": "This is a JSONPath expression that matches specific fields of\ncandidate resources and it operates as a filtering parameter\n(resources that are not matched with this expression will not\nbe candidates for transformation).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "namespaces": {
                          "description": "(Filtering parameter) Any resource subject to transformation must\nbe contained within one of the listed Kubernetes Namespace in the\nBackup. If this field is not provided, no namespace filtering will\nbe performed (all resources in all Namespaces, including all\ncluster-scoped resources, will be candidates for transformation).\nTo mix cluster-scoped and namespaced resources in the same rule,\nuse an empty string (\"\") as one of the target namespaces.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "group_kinds": {
                          "block": {
                            "attributes": {
                              "resource_group": {
                                "description": "API Group string of a Kubernetes resource, e.g.\n\"apiextensions.k8s.io\", \"storage.k8s.io\", etc.\nUse empty string for core group.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "resource_kind": {
                                "description": "Kind of a Kubernetes resource, e.g.\n\"CustomResourceDefinition\", \"StorageClass\", etc.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "(Filtering parameter) Any resource subject to transformation must\nbelong to one of the listed \"types\". If this field is not provided,\nno type filtering will be performed\n(all resources of all types matching previous filtering parameters\nwill be candidates for transformation).",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "This field is used to specify a set of fields that should be used to\ndetermine which resources in backup should be acted upon by the\nsupplied transformation rule actions, and this will ensure that only\nspecific resources are affected by transformation rule actions.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A list of transformation rules to be applied against Kubernetes\nresources as they are selected for restoration from a Backup.\nRules are executed in order defined - this order matters,\nas changes made by a rule may impact the filtering logic of subsequent\nrules. An empty list means no transformation will occur.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Defines the configuration of Restores created via this RestorePlan.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleGkeBackupRestorePlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeBackupRestorePlan), &result)
	return &result
}
