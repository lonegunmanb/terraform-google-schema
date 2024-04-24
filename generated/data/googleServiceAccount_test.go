package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v5/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestGoogleServiceAccountSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.GoogleServiceAccountSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
