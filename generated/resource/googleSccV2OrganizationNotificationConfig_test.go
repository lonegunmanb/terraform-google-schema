package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v6/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestGoogleSccV2OrganizationNotificationConfigSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.GoogleSccV2OrganizationNotificationConfigSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
