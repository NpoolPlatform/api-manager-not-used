// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/api-manager/pkg/db/ent/schema"
	"github.com/NpoolPlatform/api-manager/pkg/db/ent/serviceapi"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	serviceapiFields := schema.ServiceAPI{}.Fields()
	_ = serviceapiFields
	// serviceapiDescCreateAt is the schema descriptor for create_at field.
	serviceapiDescCreateAt := serviceapiFields[9].Descriptor()
	// serviceapi.DefaultCreateAt holds the default value on creation for the create_at field.
	serviceapi.DefaultCreateAt = serviceapiDescCreateAt.Default.(func() uint32)
	// serviceapiDescUpdateAt is the schema descriptor for update_at field.
	serviceapiDescUpdateAt := serviceapiFields[10].Descriptor()
	// serviceapi.DefaultUpdateAt holds the default value on creation for the update_at field.
	serviceapi.DefaultUpdateAt = serviceapiDescUpdateAt.Default.(func() uint32)
	// serviceapi.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	serviceapi.UpdateDefaultUpdateAt = serviceapiDescUpdateAt.UpdateDefault.(func() uint32)
	// serviceapiDescDeleteAt is the schema descriptor for delete_at field.
	serviceapiDescDeleteAt := serviceapiFields[11].Descriptor()
	// serviceapi.DefaultDeleteAt holds the default value on creation for the delete_at field.
	serviceapi.DefaultDeleteAt = serviceapiDescDeleteAt.Default.(func() uint32)
	// serviceapiDescID is the schema descriptor for id field.
	serviceapiDescID := serviceapiFields[0].Descriptor()
	// serviceapi.DefaultID holds the default value on creation for the id field.
	serviceapi.DefaultID = serviceapiDescID.Default.(func() uuid.UUID)
}
