// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/carlosruizg/muni/ent/labellingtask"
	"github.com/carlosruizg/muni/ent/labellingtaskresponse"
	"github.com/carlosruizg/muni/ent/predicate"
	"github.com/carlosruizg/muni/ent/qualification"
)

// LabellingTaskUpdate is the builder for updating LabellingTask entities.
type LabellingTaskUpdate struct {
	config
	hooks    []Hook
	mutation *LabellingTaskMutation
}

// Where appends a list predicates to the LabellingTaskUpdate builder.
func (ltu *LabellingTaskUpdate) Where(ps ...predicate.LabellingTask) *LabellingTaskUpdate {
	ltu.mutation.Where(ps...)
	return ltu
}

// SetTitle sets the "title" field.
func (ltu *LabellingTaskUpdate) SetTitle(s string) *LabellingTaskUpdate {
	ltu.mutation.SetTitle(s)
	return ltu
}

// SetDescription sets the "description" field.
func (ltu *LabellingTaskUpdate) SetDescription(s string) *LabellingTaskUpdate {
	ltu.mutation.SetDescription(s)
	return ltu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ltu *LabellingTaskUpdate) SetNillableDescription(s *string) *LabellingTaskUpdate {
	if s != nil {
		ltu.SetDescription(*s)
	}
	return ltu
}

// ClearDescription clears the value of the "description" field.
func (ltu *LabellingTaskUpdate) ClearDescription() *LabellingTaskUpdate {
	ltu.mutation.ClearDescription()
	return ltu
}

// AddResponseIDs adds the "responses" edge to the LabellingTaskResponse entity by IDs.
func (ltu *LabellingTaskUpdate) AddResponseIDs(ids ...int) *LabellingTaskUpdate {
	ltu.mutation.AddResponseIDs(ids...)
	return ltu
}

// AddResponses adds the "responses" edges to the LabellingTaskResponse entity.
func (ltu *LabellingTaskUpdate) AddResponses(l ...*LabellingTaskResponse) *LabellingTaskUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ltu.AddResponseIDs(ids...)
}

// AddExpertRequirementIDs adds the "expert_requirements" edge to the Qualification entity by IDs.
func (ltu *LabellingTaskUpdate) AddExpertRequirementIDs(ids ...int) *LabellingTaskUpdate {
	ltu.mutation.AddExpertRequirementIDs(ids...)
	return ltu
}

// AddExpertRequirements adds the "expert_requirements" edges to the Qualification entity.
func (ltu *LabellingTaskUpdate) AddExpertRequirements(q ...*Qualification) *LabellingTaskUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return ltu.AddExpertRequirementIDs(ids...)
}

// Mutation returns the LabellingTaskMutation object of the builder.
func (ltu *LabellingTaskUpdate) Mutation() *LabellingTaskMutation {
	return ltu.mutation
}

// ClearResponses clears all "responses" edges to the LabellingTaskResponse entity.
func (ltu *LabellingTaskUpdate) ClearResponses() *LabellingTaskUpdate {
	ltu.mutation.ClearResponses()
	return ltu
}

// RemoveResponseIDs removes the "responses" edge to LabellingTaskResponse entities by IDs.
func (ltu *LabellingTaskUpdate) RemoveResponseIDs(ids ...int) *LabellingTaskUpdate {
	ltu.mutation.RemoveResponseIDs(ids...)
	return ltu
}

// RemoveResponses removes "responses" edges to LabellingTaskResponse entities.
func (ltu *LabellingTaskUpdate) RemoveResponses(l ...*LabellingTaskResponse) *LabellingTaskUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ltu.RemoveResponseIDs(ids...)
}

// ClearExpertRequirements clears all "expert_requirements" edges to the Qualification entity.
func (ltu *LabellingTaskUpdate) ClearExpertRequirements() *LabellingTaskUpdate {
	ltu.mutation.ClearExpertRequirements()
	return ltu
}

// RemoveExpertRequirementIDs removes the "expert_requirements" edge to Qualification entities by IDs.
func (ltu *LabellingTaskUpdate) RemoveExpertRequirementIDs(ids ...int) *LabellingTaskUpdate {
	ltu.mutation.RemoveExpertRequirementIDs(ids...)
	return ltu
}

// RemoveExpertRequirements removes "expert_requirements" edges to Qualification entities.
func (ltu *LabellingTaskUpdate) RemoveExpertRequirements(q ...*Qualification) *LabellingTaskUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return ltu.RemoveExpertRequirementIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ltu *LabellingTaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ltu.sqlSave, ltu.mutation, ltu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ltu *LabellingTaskUpdate) SaveX(ctx context.Context) int {
	affected, err := ltu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ltu *LabellingTaskUpdate) Exec(ctx context.Context) error {
	_, err := ltu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltu *LabellingTaskUpdate) ExecX(ctx context.Context) {
	if err := ltu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ltu *LabellingTaskUpdate) check() error {
	if v, ok := ltu.mutation.Title(); ok {
		if err := labellingtask.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "LabellingTask.title": %w`, err)}
		}
	}
	return nil
}

func (ltu *LabellingTaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ltu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(labellingtask.Table, labellingtask.Columns, sqlgraph.NewFieldSpec(labellingtask.FieldID, field.TypeInt))
	if ps := ltu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ltu.mutation.Title(); ok {
		_spec.SetField(labellingtask.FieldTitle, field.TypeString, value)
	}
	if value, ok := ltu.mutation.Description(); ok {
		_spec.SetField(labellingtask.FieldDescription, field.TypeString, value)
	}
	if ltu.mutation.DescriptionCleared() {
		_spec.ClearField(labellingtask.FieldDescription, field.TypeString)
	}
	if ltu.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   labellingtask.ResponsesTable,
			Columns: []string{labellingtask.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(labellingtaskresponse.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ltu.mutation.RemovedResponsesIDs(); len(nodes) > 0 && !ltu.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   labellingtask.ResponsesTable,
			Columns: []string{labellingtask.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(labellingtaskresponse.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ltu.mutation.ResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   labellingtask.ResponsesTable,
			Columns: []string{labellingtask.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(labellingtaskresponse.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ltu.mutation.ExpertRequirementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   labellingtask.ExpertRequirementsTable,
			Columns: labellingtask.ExpertRequirementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(qualification.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ltu.mutation.RemovedExpertRequirementsIDs(); len(nodes) > 0 && !ltu.mutation.ExpertRequirementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   labellingtask.ExpertRequirementsTable,
			Columns: labellingtask.ExpertRequirementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(qualification.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ltu.mutation.ExpertRequirementsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   labellingtask.ExpertRequirementsTable,
			Columns: labellingtask.ExpertRequirementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(qualification.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ltu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{labellingtask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ltu.mutation.done = true
	return n, nil
}

// LabellingTaskUpdateOne is the builder for updating a single LabellingTask entity.
type LabellingTaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LabellingTaskMutation
}

// SetTitle sets the "title" field.
func (ltuo *LabellingTaskUpdateOne) SetTitle(s string) *LabellingTaskUpdateOne {
	ltuo.mutation.SetTitle(s)
	return ltuo
}

// SetDescription sets the "description" field.
func (ltuo *LabellingTaskUpdateOne) SetDescription(s string) *LabellingTaskUpdateOne {
	ltuo.mutation.SetDescription(s)
	return ltuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ltuo *LabellingTaskUpdateOne) SetNillableDescription(s *string) *LabellingTaskUpdateOne {
	if s != nil {
		ltuo.SetDescription(*s)
	}
	return ltuo
}

// ClearDescription clears the value of the "description" field.
func (ltuo *LabellingTaskUpdateOne) ClearDescription() *LabellingTaskUpdateOne {
	ltuo.mutation.ClearDescription()
	return ltuo
}

// AddResponseIDs adds the "responses" edge to the LabellingTaskResponse entity by IDs.
func (ltuo *LabellingTaskUpdateOne) AddResponseIDs(ids ...int) *LabellingTaskUpdateOne {
	ltuo.mutation.AddResponseIDs(ids...)
	return ltuo
}

// AddResponses adds the "responses" edges to the LabellingTaskResponse entity.
func (ltuo *LabellingTaskUpdateOne) AddResponses(l ...*LabellingTaskResponse) *LabellingTaskUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ltuo.AddResponseIDs(ids...)
}

// AddExpertRequirementIDs adds the "expert_requirements" edge to the Qualification entity by IDs.
func (ltuo *LabellingTaskUpdateOne) AddExpertRequirementIDs(ids ...int) *LabellingTaskUpdateOne {
	ltuo.mutation.AddExpertRequirementIDs(ids...)
	return ltuo
}

// AddExpertRequirements adds the "expert_requirements" edges to the Qualification entity.
func (ltuo *LabellingTaskUpdateOne) AddExpertRequirements(q ...*Qualification) *LabellingTaskUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return ltuo.AddExpertRequirementIDs(ids...)
}

// Mutation returns the LabellingTaskMutation object of the builder.
func (ltuo *LabellingTaskUpdateOne) Mutation() *LabellingTaskMutation {
	return ltuo.mutation
}

// ClearResponses clears all "responses" edges to the LabellingTaskResponse entity.
func (ltuo *LabellingTaskUpdateOne) ClearResponses() *LabellingTaskUpdateOne {
	ltuo.mutation.ClearResponses()
	return ltuo
}

// RemoveResponseIDs removes the "responses" edge to LabellingTaskResponse entities by IDs.
func (ltuo *LabellingTaskUpdateOne) RemoveResponseIDs(ids ...int) *LabellingTaskUpdateOne {
	ltuo.mutation.RemoveResponseIDs(ids...)
	return ltuo
}

// RemoveResponses removes "responses" edges to LabellingTaskResponse entities.
func (ltuo *LabellingTaskUpdateOne) RemoveResponses(l ...*LabellingTaskResponse) *LabellingTaskUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ltuo.RemoveResponseIDs(ids...)
}

// ClearExpertRequirements clears all "expert_requirements" edges to the Qualification entity.
func (ltuo *LabellingTaskUpdateOne) ClearExpertRequirements() *LabellingTaskUpdateOne {
	ltuo.mutation.ClearExpertRequirements()
	return ltuo
}

// RemoveExpertRequirementIDs removes the "expert_requirements" edge to Qualification entities by IDs.
func (ltuo *LabellingTaskUpdateOne) RemoveExpertRequirementIDs(ids ...int) *LabellingTaskUpdateOne {
	ltuo.mutation.RemoveExpertRequirementIDs(ids...)
	return ltuo
}

// RemoveExpertRequirements removes "expert_requirements" edges to Qualification entities.
func (ltuo *LabellingTaskUpdateOne) RemoveExpertRequirements(q ...*Qualification) *LabellingTaskUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return ltuo.RemoveExpertRequirementIDs(ids...)
}

// Where appends a list predicates to the LabellingTaskUpdate builder.
func (ltuo *LabellingTaskUpdateOne) Where(ps ...predicate.LabellingTask) *LabellingTaskUpdateOne {
	ltuo.mutation.Where(ps...)
	return ltuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ltuo *LabellingTaskUpdateOne) Select(field string, fields ...string) *LabellingTaskUpdateOne {
	ltuo.fields = append([]string{field}, fields...)
	return ltuo
}

// Save executes the query and returns the updated LabellingTask entity.
func (ltuo *LabellingTaskUpdateOne) Save(ctx context.Context) (*LabellingTask, error) {
	return withHooks(ctx, ltuo.sqlSave, ltuo.mutation, ltuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ltuo *LabellingTaskUpdateOne) SaveX(ctx context.Context) *LabellingTask {
	node, err := ltuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ltuo *LabellingTaskUpdateOne) Exec(ctx context.Context) error {
	_, err := ltuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltuo *LabellingTaskUpdateOne) ExecX(ctx context.Context) {
	if err := ltuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ltuo *LabellingTaskUpdateOne) check() error {
	if v, ok := ltuo.mutation.Title(); ok {
		if err := labellingtask.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "LabellingTask.title": %w`, err)}
		}
	}
	return nil
}

func (ltuo *LabellingTaskUpdateOne) sqlSave(ctx context.Context) (_node *LabellingTask, err error) {
	if err := ltuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(labellingtask.Table, labellingtask.Columns, sqlgraph.NewFieldSpec(labellingtask.FieldID, field.TypeInt))
	id, ok := ltuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LabellingTask.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ltuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, labellingtask.FieldID)
		for _, f := range fields {
			if !labellingtask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != labellingtask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ltuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ltuo.mutation.Title(); ok {
		_spec.SetField(labellingtask.FieldTitle, field.TypeString, value)
	}
	if value, ok := ltuo.mutation.Description(); ok {
		_spec.SetField(labellingtask.FieldDescription, field.TypeString, value)
	}
	if ltuo.mutation.DescriptionCleared() {
		_spec.ClearField(labellingtask.FieldDescription, field.TypeString)
	}
	if ltuo.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   labellingtask.ResponsesTable,
			Columns: []string{labellingtask.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(labellingtaskresponse.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ltuo.mutation.RemovedResponsesIDs(); len(nodes) > 0 && !ltuo.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   labellingtask.ResponsesTable,
			Columns: []string{labellingtask.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(labellingtaskresponse.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ltuo.mutation.ResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   labellingtask.ResponsesTable,
			Columns: []string{labellingtask.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(labellingtaskresponse.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ltuo.mutation.ExpertRequirementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   labellingtask.ExpertRequirementsTable,
			Columns: labellingtask.ExpertRequirementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(qualification.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ltuo.mutation.RemovedExpertRequirementsIDs(); len(nodes) > 0 && !ltuo.mutation.ExpertRequirementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   labellingtask.ExpertRequirementsTable,
			Columns: labellingtask.ExpertRequirementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(qualification.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ltuo.mutation.ExpertRequirementsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   labellingtask.ExpertRequirementsTable,
			Columns: labellingtask.ExpertRequirementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(qualification.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LabellingTask{config: ltuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ltuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{labellingtask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ltuo.mutation.done = true
	return _node, nil
}
