// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"go.infratographer.com/resource-provider-api/internal/ent/generated/resourceprovider"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (rp *ResourceProviderQuery) CollectFields(ctx context.Context, satisfies ...string) (*ResourceProviderQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return rp, nil
	}
	if err := rp.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return rp, nil
}

func (rp *ResourceProviderQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(resourceprovider.Columns))
		selectedFields = []string{resourceprovider.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "createdAt":
			if _, ok := fieldSeen[resourceprovider.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, resourceprovider.FieldCreatedAt)
				fieldSeen[resourceprovider.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[resourceprovider.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, resourceprovider.FieldUpdatedAt)
				fieldSeen[resourceprovider.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[resourceprovider.FieldName]; !ok {
				selectedFields = append(selectedFields, resourceprovider.FieldName)
				fieldSeen[resourceprovider.FieldName] = struct{}{}
			}
		case "description":
			if _, ok := fieldSeen[resourceprovider.FieldDescription]; !ok {
				selectedFields = append(selectedFields, resourceprovider.FieldDescription)
				fieldSeen[resourceprovider.FieldDescription] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		rp.Select(selectedFields...)
	}
	return nil
}

type resourceproviderPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ResourceProviderPaginateOption
}

func newResourceProviderPaginateArgs(rv map[string]any) *resourceproviderPaginateArgs {
	args := &resourceproviderPaginateArgs{}
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
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &ResourceProviderOrder{Field: &ResourceProviderOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithResourceProviderOrder(order))
			}
		case *ResourceProviderOrder:
			if v != nil {
				args.opts = append(args.opts, WithResourceProviderOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*ResourceProviderWhereInput); ok {
		args.opts = append(args.opts, WithResourceProviderFilter(v.Filter))
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
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
