package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleArtifactRegistryDockerImage = `{
  "block": {
    "attributes": {
      "build_time": {
        "computed": true,
        "description": "The time, as a RFC 3339 string, this image was built.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_name": {
        "description": "The image name to fetch.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "image_size_bytes": {
        "computed": true,
        "description": "Calculated size of the image in bytes.",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "The region of the artifact registry repository. For example, \"us-west1\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "media_type": {
        "computed": true,
        "description": "Media type of this image.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The fully qualified name of the fetched image.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description": "Project ID of the project.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_id": {
        "description": "The last part of the repository name to fetch from.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The URI to access the image.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "All tags associated with the image.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "The time, as a RFC 3339 string, this image was updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "upload_time": {
        "computed": true,
        "description": "The time, as a RFC 3339 string, the image was uploaded.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleArtifactRegistryDockerImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleArtifactRegistryDockerImage), &result)
	return &result
}
