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
		field.String("protocol"),
		field.String("service_name"),
		field.JSON("domains", []string{}),
		field.String("method"),
		field.String("method_name"),
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

func (ServiceAPI) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("service_name", "method", "path").
			Unique(),
	}
}
