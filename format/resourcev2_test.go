package format

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchemaMapping_Format(t *testing.T) {
	var err error
	// data source
	schemaMapping1 := NewSchemaMapping("1.173.0", "data", "kubernetes_version")
	assert.NotNil(t, schemaMapping1)
	assert.Equal(t, "cs_kubernetes_version", schemaMapping1.Name)
	assert.Equal(t, "data", schemaMapping1.Type)
	assert.Equal(t, "1.173.0", schemaMapping1.Version)
	err = schemaMapping1.Format()
	assert.Nil(t, err)
	assert.NotNil(t, schemaMapping1.Fields)
	assert.NotNil(t, schemaMapping1.Fields[0])
	fmt.Println(schemaMapping1.Fields)

	// resource
	schemaMapping2 := NewSchemaMapping("1.173.0", "resource", "managed_kubernetes")
	assert.NotNil(t, schemaMapping2)
	assert.Equal(t, "cs_managed_kubernetes", schemaMapping2.Name)
	assert.Equal(t, "resource", schemaMapping2.Type)
	assert.Equal(t, "1.173.0", schemaMapping2.Version)
	err = schemaMapping2.Format()
	assert.Nil(t, err)
	assert.NotNil(t, schemaMapping2.Fields)
	assert.NotNil(t, schemaMapping2.Fields[0])
	fmt.Println(schemaMapping2.Fields)
}
