package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v6/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestGoogleKmsKeyHandlesSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.GoogleKmsKeyHandlesSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
