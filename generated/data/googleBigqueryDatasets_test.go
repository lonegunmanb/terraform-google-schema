package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v6/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestGoogleBigqueryDatasetsSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.GoogleBigqueryDatasetsSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
