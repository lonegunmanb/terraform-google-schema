package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v5/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestGoogleComputeSnapshotSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.GoogleComputeSnapshotSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
