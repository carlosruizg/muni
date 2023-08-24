// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/carlosruizg/muni/ent/expert"
	"github.com/carlosruizg/muni/ent/labellingproject"
	"github.com/carlosruizg/muni/ent/labellingtask"
	"github.com/carlosruizg/muni/ent/labellingtaskresponse"
	"github.com/carlosruizg/muni/ent/qualification"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (e *ExpertQuery) CollectFields(ctx context.Context, satisfies ...string) (*ExpertQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return e, nil
	}
	if err := e.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return e, nil
}

func (e *ExpertQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(expert.Columns))
		selectedFields = []string{expert.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "qualifications":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&QualificationClient{config: e.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, qualificationImplementors)...); err != nil {
				return err
			}
			e.WithNamedQualifications(alias, func(wq *QualificationQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[expert.FieldName]; !ok {
				selectedFields = append(selectedFields, expert.FieldName)
				fieldSeen[expert.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		e.Select(selectedFields...)
	}
	return nil
}

type expertPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ExpertPaginateOption
}

func newExpertPaginateArgs(rv map[string]any) *expertPaginateArgs {
	args := &expertPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (lp *LabellingProjectQuery) CollectFields(ctx context.Context, satisfies ...string) (*LabellingProjectQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return lp, nil
	}
	if err := lp.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return lp, nil
}

func (lp *LabellingProjectQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(labellingproject.Columns))
		selectedFields = []string{labellingproject.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "name":
			if _, ok := fieldSeen[labellingproject.FieldName]; !ok {
				selectedFields = append(selectedFields, labellingproject.FieldName)
				fieldSeen[labellingproject.FieldName] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[labellingproject.FieldStatus]; !ok {
				selectedFields = append(selectedFields, labellingproject.FieldStatus)
				fieldSeen[labellingproject.FieldStatus] = struct{}{}
			}
		case "description":
			if _, ok := fieldSeen[labellingproject.FieldDescription]; !ok {
				selectedFields = append(selectedFields, labellingproject.FieldDescription)
				fieldSeen[labellingproject.FieldDescription] = struct{}{}
			}
		case "isPrivate":
			if _, ok := fieldSeen[labellingproject.FieldIsPrivate]; !ok {
				selectedFields = append(selectedFields, labellingproject.FieldIsPrivate)
				fieldSeen[labellingproject.FieldIsPrivate] = struct{}{}
			}
		case "callbackURL":
			if _, ok := fieldSeen[labellingproject.FieldCallbackURL]; !ok {
				selectedFields = append(selectedFields, labellingproject.FieldCallbackURL)
				fieldSeen[labellingproject.FieldCallbackURL] = struct{}{}
			}
		case "workersPerTask":
			if _, ok := fieldSeen[labellingproject.FieldWorkersPerTask]; !ok {
				selectedFields = append(selectedFields, labellingproject.FieldWorkersPerTask)
				fieldSeen[labellingproject.FieldWorkersPerTask] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		lp.Select(selectedFields...)
	}
	return nil
}

type labellingprojectPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []LabellingProjectPaginateOption
}

func newLabellingProjectPaginateArgs(rv map[string]any) *labellingprojectPaginateArgs {
	args := &labellingprojectPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (lt *LabellingTaskQuery) CollectFields(ctx context.Context, satisfies ...string) (*LabellingTaskQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return lt, nil
	}
	if err := lt.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return lt, nil
}

func (lt *LabellingTaskQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(labellingtask.Columns))
		selectedFields = []string{labellingtask.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "name":
			if _, ok := fieldSeen[labellingtask.FieldName]; !ok {
				selectedFields = append(selectedFields, labellingtask.FieldName)
				fieldSeen[labellingtask.FieldName] = struct{}{}
			}
		case "description":
			if _, ok := fieldSeen[labellingtask.FieldDescription]; !ok {
				selectedFields = append(selectedFields, labellingtask.FieldDescription)
				fieldSeen[labellingtask.FieldDescription] = struct{}{}
			}
		case "instructions":
			if _, ok := fieldSeen[labellingtask.FieldInstructions]; !ok {
				selectedFields = append(selectedFields, labellingtask.FieldInstructions)
				fieldSeen[labellingtask.FieldInstructions] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		lt.Select(selectedFields...)
	}
	return nil
}

type labellingtaskPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []LabellingTaskPaginateOption
}

func newLabellingTaskPaginateArgs(rv map[string]any) *labellingtaskPaginateArgs {
	args := &labellingtaskPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ltr *LabellingTaskResponseQuery) CollectFields(ctx context.Context, satisfies ...string) (*LabellingTaskResponseQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ltr, nil
	}
	if err := ltr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ltr, nil
}

func (ltr *LabellingTaskResponseQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(labellingtaskresponse.Columns))
		selectedFields = []string{labellingtaskresponse.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "response":
			if _, ok := fieldSeen[labellingtaskresponse.FieldResponse]; !ok {
				selectedFields = append(selectedFields, labellingtaskresponse.FieldResponse)
				fieldSeen[labellingtaskresponse.FieldResponse] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ltr.Select(selectedFields...)
	}
	return nil
}

type labellingtaskresponsePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []LabellingTaskResponsePaginateOption
}

func newLabellingTaskResponsePaginateArgs(rv map[string]any) *labellingtaskresponsePaginateArgs {
	args := &labellingtaskresponsePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (q *QualificationQuery) CollectFields(ctx context.Context, satisfies ...string) (*QualificationQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return q, nil
	}
	if err := q.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return q, nil
}

func (q *QualificationQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(qualification.Columns))
		selectedFields = []string{qualification.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "experts":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ExpertClient{config: q.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, expertImplementors)...); err != nil {
				return err
			}
			q.WithNamedExperts(alias, func(wq *ExpertQuery) {
				*wq = *query
			})
		case "value":
			if _, ok := fieldSeen[qualification.FieldValue]; !ok {
				selectedFields = append(selectedFields, qualification.FieldValue)
				fieldSeen[qualification.FieldValue] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		q.Select(selectedFields...)
	}
	return nil
}

type qualificationPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []QualificationPaginateOption
}

func newQualificationPaginateArgs(rv map[string]any) *qualificationPaginateArgs {
	args := &qualificationPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
