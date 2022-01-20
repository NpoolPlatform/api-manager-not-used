// Code generated by entc, DO NOT EDIT.

package serviceapi

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/api-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ServiceName applies equality check predicate on the "service_name" field. It's identical to ServiceNameEQ.
func ServiceName(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServiceName), v))
	})
}

// Method applies equality check predicate on the "method" field. It's identical to MethodEQ.
func Method(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMethod), v))
	})
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// Exported applies equality check predicate on the "exported" field. It's identical to ExportedEQ.
func Exported(v bool) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExported), v))
	})
}

// PathPrefix applies equality check predicate on the "path_prefix" field. It's identical to PathPrefixEQ.
func PathPrefix(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPathPrefix), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// ServiceNameEQ applies the EQ predicate on the "service_name" field.
func ServiceNameEQ(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServiceName), v))
	})
}

// ServiceNameNEQ applies the NEQ predicate on the "service_name" field.
func ServiceNameNEQ(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldServiceName), v))
	})
}

// ServiceNameIn applies the In predicate on the "service_name" field.
func ServiceNameIn(vs ...string) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldServiceName), v...))
	})
}

// ServiceNameNotIn applies the NotIn predicate on the "service_name" field.
func ServiceNameNotIn(vs ...string) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldServiceName), v...))
	})
}

// ServiceNameGT applies the GT predicate on the "service_name" field.
func ServiceNameGT(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldServiceName), v))
	})
}

// ServiceNameGTE applies the GTE predicate on the "service_name" field.
func ServiceNameGTE(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldServiceName), v))
	})
}

// ServiceNameLT applies the LT predicate on the "service_name" field.
func ServiceNameLT(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldServiceName), v))
	})
}

// ServiceNameLTE applies the LTE predicate on the "service_name" field.
func ServiceNameLTE(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldServiceName), v))
	})
}

// ServiceNameContains applies the Contains predicate on the "service_name" field.
func ServiceNameContains(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldServiceName), v))
	})
}

// ServiceNameHasPrefix applies the HasPrefix predicate on the "service_name" field.
func ServiceNameHasPrefix(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldServiceName), v))
	})
}

// ServiceNameHasSuffix applies the HasSuffix predicate on the "service_name" field.
func ServiceNameHasSuffix(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldServiceName), v))
	})
}

// ServiceNameEqualFold applies the EqualFold predicate on the "service_name" field.
func ServiceNameEqualFold(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldServiceName), v))
	})
}

// ServiceNameContainsFold applies the ContainsFold predicate on the "service_name" field.
func ServiceNameContainsFold(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldServiceName), v))
	})
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMethod), v))
	})
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMethod), v))
	})
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...string) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMethod), v...))
	})
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...string) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMethod), v...))
	})
}

// MethodGT applies the GT predicate on the "method" field.
func MethodGT(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMethod), v))
	})
}

// MethodGTE applies the GTE predicate on the "method" field.
func MethodGTE(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMethod), v))
	})
}

// MethodLT applies the LT predicate on the "method" field.
func MethodLT(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMethod), v))
	})
}

// MethodLTE applies the LTE predicate on the "method" field.
func MethodLTE(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMethod), v))
	})
}

// MethodContains applies the Contains predicate on the "method" field.
func MethodContains(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMethod), v))
	})
}

// MethodHasPrefix applies the HasPrefix predicate on the "method" field.
func MethodHasPrefix(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMethod), v))
	})
}

// MethodHasSuffix applies the HasSuffix predicate on the "method" field.
func MethodHasSuffix(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMethod), v))
	})
}

// MethodEqualFold applies the EqualFold predicate on the "method" field.
func MethodEqualFold(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMethod), v))
	})
}

// MethodContainsFold applies the ContainsFold predicate on the "method" field.
func MethodContainsFold(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMethod), v))
	})
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPath), v))
	})
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPath), v...))
	})
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPath), v...))
	})
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPath), v))
	})
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPath), v))
	})
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPath), v))
	})
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPath), v))
	})
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPath), v))
	})
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPath), v))
	})
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPath), v))
	})
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPath), v))
	})
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPath), v))
	})
}

// ExportedEQ applies the EQ predicate on the "exported" field.
func ExportedEQ(v bool) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExported), v))
	})
}

// ExportedNEQ applies the NEQ predicate on the "exported" field.
func ExportedNEQ(v bool) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExported), v))
	})
}

// PathPrefixEQ applies the EQ predicate on the "path_prefix" field.
func PathPrefixEQ(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPathPrefix), v))
	})
}

// PathPrefixNEQ applies the NEQ predicate on the "path_prefix" field.
func PathPrefixNEQ(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPathPrefix), v))
	})
}

// PathPrefixIn applies the In predicate on the "path_prefix" field.
func PathPrefixIn(vs ...string) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPathPrefix), v...))
	})
}

// PathPrefixNotIn applies the NotIn predicate on the "path_prefix" field.
func PathPrefixNotIn(vs ...string) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPathPrefix), v...))
	})
}

// PathPrefixGT applies the GT predicate on the "path_prefix" field.
func PathPrefixGT(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPathPrefix), v))
	})
}

// PathPrefixGTE applies the GTE predicate on the "path_prefix" field.
func PathPrefixGTE(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPathPrefix), v))
	})
}

// PathPrefixLT applies the LT predicate on the "path_prefix" field.
func PathPrefixLT(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPathPrefix), v))
	})
}

// PathPrefixLTE applies the LTE predicate on the "path_prefix" field.
func PathPrefixLTE(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPathPrefix), v))
	})
}

// PathPrefixContains applies the Contains predicate on the "path_prefix" field.
func PathPrefixContains(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPathPrefix), v))
	})
}

// PathPrefixHasPrefix applies the HasPrefix predicate on the "path_prefix" field.
func PathPrefixHasPrefix(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPathPrefix), v))
	})
}

// PathPrefixHasSuffix applies the HasSuffix predicate on the "path_prefix" field.
func PathPrefixHasSuffix(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPathPrefix), v))
	})
}

// PathPrefixEqualFold applies the EqualFold predicate on the "path_prefix" field.
func PathPrefixEqualFold(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPathPrefix), v))
	})
}

// PathPrefixContainsFold applies the ContainsFold predicate on the "path_prefix" field.
func PathPrefixContainsFold(v string) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPathPrefix), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...uint32) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...uint32) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...uint32) predicate.ServiceAPI {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceAPI(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ServiceAPI) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ServiceAPI) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ServiceAPI) predicate.ServiceAPI {
	return predicate.ServiceAPI(func(s *sql.Selector) {
		p(s.Not())
	})
}
