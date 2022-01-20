package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// ServiceAPI holds the schema definition for the ServiceAPI entity.
type ServiceAPI struct {
	ent.Schema
}

// Fields of the ServiceAPI.
func (ServiceAPI) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("domain"),
		field.String("method"),
		field.String("path"),
		field.Bool("exported"),
		field.String("path_prefix"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the ServiceAPI.
func (ServiceAPI) Edges() []ent.Edge {
	return nil
}

func (ServiceAPI) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("domain", "method", "path").
			Unique(),
	}
}
