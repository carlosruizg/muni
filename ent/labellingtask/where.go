// Code generated by ent, DO NOT EDIT.

package labellingtask

import (
	"entgo.io/ent/dialect/sql"
	"github.com/carlosruizg/muni/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldEQ(FieldDescription, v))
}

// Instructions applies equality check predicate on the "instructions" field. It's identical to InstructionsEQ.
func Instructions(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldEQ(FieldInstructions, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldContainsFold(FieldDescription, v))
}

// InstructionsEQ applies the EQ predicate on the "instructions" field.
func InstructionsEQ(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldEQ(FieldInstructions, v))
}

// InstructionsNEQ applies the NEQ predicate on the "instructions" field.
func InstructionsNEQ(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldNEQ(FieldInstructions, v))
}

// InstructionsIn applies the In predicate on the "instructions" field.
func InstructionsIn(vs ...string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldIn(FieldInstructions, vs...))
}

// InstructionsNotIn applies the NotIn predicate on the "instructions" field.
func InstructionsNotIn(vs ...string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldNotIn(FieldInstructions, vs...))
}

// InstructionsGT applies the GT predicate on the "instructions" field.
func InstructionsGT(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldGT(FieldInstructions, v))
}

// InstructionsGTE applies the GTE predicate on the "instructions" field.
func InstructionsGTE(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldGTE(FieldInstructions, v))
}

// InstructionsLT applies the LT predicate on the "instructions" field.
func InstructionsLT(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldLT(FieldInstructions, v))
}

// InstructionsLTE applies the LTE predicate on the "instructions" field.
func InstructionsLTE(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldLTE(FieldInstructions, v))
}

// InstructionsContains applies the Contains predicate on the "instructions" field.
func InstructionsContains(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldContains(FieldInstructions, v))
}

// InstructionsHasPrefix applies the HasPrefix predicate on the "instructions" field.
func InstructionsHasPrefix(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldHasPrefix(FieldInstructions, v))
}

// InstructionsHasSuffix applies the HasSuffix predicate on the "instructions" field.
func InstructionsHasSuffix(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldHasSuffix(FieldInstructions, v))
}

// InstructionsEqualFold applies the EqualFold predicate on the "instructions" field.
func InstructionsEqualFold(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldEqualFold(FieldInstructions, v))
}

// InstructionsContainsFold applies the ContainsFold predicate on the "instructions" field.
func InstructionsContainsFold(v string) predicate.LabellingTask {
	return predicate.LabellingTask(sql.FieldContainsFold(FieldInstructions, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LabellingTask) predicate.LabellingTask {
	return predicate.LabellingTask(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LabellingTask) predicate.LabellingTask {
	return predicate.LabellingTask(func(s *sql.Selector) {
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
func Not(p predicate.LabellingTask) predicate.LabellingTask {
	return predicate.LabellingTask(func(s *sql.Selector) {
		p(s.Not())
	})
}
