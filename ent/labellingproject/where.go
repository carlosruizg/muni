// Code generated by ent, DO NOT EDIT.

package labellingproject

import (
	"entgo.io/ent/dialect/sql"
	"github.com/carlosruizg/muni/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEQ(FieldDescription, v))
}

// IsPrivate applies equality check predicate on the "is_private" field. It's identical to IsPrivateEQ.
func IsPrivate(v bool) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEQ(FieldIsPrivate, v))
}

// CallbackURL applies equality check predicate on the "callback_url" field. It's identical to CallbackURLEQ.
func CallbackURL(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEQ(FieldCallbackURL, v))
}

// WorkersPerTask applies equality check predicate on the "workers_per_task" field. It's identical to WorkersPerTaskEQ.
func WorkersPerTask(v int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEQ(FieldWorkersPerTask, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldContainsFold(FieldName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNotIn(FieldStatus, vs...))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldContainsFold(FieldDescription, v))
}

// IsPrivateEQ applies the EQ predicate on the "is_private" field.
func IsPrivateEQ(v bool) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEQ(FieldIsPrivate, v))
}

// IsPrivateNEQ applies the NEQ predicate on the "is_private" field.
func IsPrivateNEQ(v bool) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNEQ(FieldIsPrivate, v))
}

// CallbackURLEQ applies the EQ predicate on the "callback_url" field.
func CallbackURLEQ(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEQ(FieldCallbackURL, v))
}

// CallbackURLNEQ applies the NEQ predicate on the "callback_url" field.
func CallbackURLNEQ(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNEQ(FieldCallbackURL, v))
}

// CallbackURLIn applies the In predicate on the "callback_url" field.
func CallbackURLIn(vs ...string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldIn(FieldCallbackURL, vs...))
}

// CallbackURLNotIn applies the NotIn predicate on the "callback_url" field.
func CallbackURLNotIn(vs ...string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNotIn(FieldCallbackURL, vs...))
}

// CallbackURLGT applies the GT predicate on the "callback_url" field.
func CallbackURLGT(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldGT(FieldCallbackURL, v))
}

// CallbackURLGTE applies the GTE predicate on the "callback_url" field.
func CallbackURLGTE(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldGTE(FieldCallbackURL, v))
}

// CallbackURLLT applies the LT predicate on the "callback_url" field.
func CallbackURLLT(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldLT(FieldCallbackURL, v))
}

// CallbackURLLTE applies the LTE predicate on the "callback_url" field.
func CallbackURLLTE(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldLTE(FieldCallbackURL, v))
}

// CallbackURLContains applies the Contains predicate on the "callback_url" field.
func CallbackURLContains(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldContains(FieldCallbackURL, v))
}

// CallbackURLHasPrefix applies the HasPrefix predicate on the "callback_url" field.
func CallbackURLHasPrefix(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldHasPrefix(FieldCallbackURL, v))
}

// CallbackURLHasSuffix applies the HasSuffix predicate on the "callback_url" field.
func CallbackURLHasSuffix(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldHasSuffix(FieldCallbackURL, v))
}

// CallbackURLIsNil applies the IsNil predicate on the "callback_url" field.
func CallbackURLIsNil() predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldIsNull(FieldCallbackURL))
}

// CallbackURLNotNil applies the NotNil predicate on the "callback_url" field.
func CallbackURLNotNil() predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNotNull(FieldCallbackURL))
}

// CallbackURLEqualFold applies the EqualFold predicate on the "callback_url" field.
func CallbackURLEqualFold(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEqualFold(FieldCallbackURL, v))
}

// CallbackURLContainsFold applies the ContainsFold predicate on the "callback_url" field.
func CallbackURLContainsFold(v string) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldContainsFold(FieldCallbackURL, v))
}

// WorkersPerTaskEQ applies the EQ predicate on the "workers_per_task" field.
func WorkersPerTaskEQ(v int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldEQ(FieldWorkersPerTask, v))
}

// WorkersPerTaskNEQ applies the NEQ predicate on the "workers_per_task" field.
func WorkersPerTaskNEQ(v int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNEQ(FieldWorkersPerTask, v))
}

// WorkersPerTaskIn applies the In predicate on the "workers_per_task" field.
func WorkersPerTaskIn(vs ...int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldIn(FieldWorkersPerTask, vs...))
}

// WorkersPerTaskNotIn applies the NotIn predicate on the "workers_per_task" field.
func WorkersPerTaskNotIn(vs ...int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldNotIn(FieldWorkersPerTask, vs...))
}

// WorkersPerTaskGT applies the GT predicate on the "workers_per_task" field.
func WorkersPerTaskGT(v int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldGT(FieldWorkersPerTask, v))
}

// WorkersPerTaskGTE applies the GTE predicate on the "workers_per_task" field.
func WorkersPerTaskGTE(v int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldGTE(FieldWorkersPerTask, v))
}

// WorkersPerTaskLT applies the LT predicate on the "workers_per_task" field.
func WorkersPerTaskLT(v int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldLT(FieldWorkersPerTask, v))
}

// WorkersPerTaskLTE applies the LTE predicate on the "workers_per_task" field.
func WorkersPerTaskLTE(v int) predicate.LabellingProject {
	return predicate.LabellingProject(sql.FieldLTE(FieldWorkersPerTask, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LabellingProject) predicate.LabellingProject {
	return predicate.LabellingProject(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LabellingProject) predicate.LabellingProject {
	return predicate.LabellingProject(func(s *sql.Selector) {
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
func Not(p predicate.LabellingProject) predicate.LabellingProject {
	return predicate.LabellingProject(func(s *sql.Selector) {
		p(s.Not())
	})
}