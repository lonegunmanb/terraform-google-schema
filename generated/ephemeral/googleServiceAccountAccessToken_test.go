package ephemeral_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v6/generated/ephemeral"
	"github.com/stretchr/testify/assert"
)

func TestGoogleServiceAccountAccessTokenSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := ephemeral.GoogleServiceAccountAccessTokenSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
