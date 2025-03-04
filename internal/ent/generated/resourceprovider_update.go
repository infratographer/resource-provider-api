// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/resource-provider-api/internal/ent/generated/predicate"
	"go.infratographer.com/resource-provider-api/internal/ent/generated/resourceprovider"
)

// ResourceProviderUpdate is the builder for updating ResourceProvider entities.
type ResourceProviderUpdate struct {
	config
	hooks    []Hook
	mutation *ResourceProviderMutation
}

// Where appends a list predicates to the ResourceProviderUpdate builder.
func (rpu *ResourceProviderUpdate) Where(ps ...predicate.ResourceProvider) *ResourceProviderUpdate {
	rpu.mutation.Where(ps...)
	return rpu
}

// SetName sets the "name" field.
func (rpu *ResourceProviderUpdate) SetName(s string) *ResourceProviderUpdate {
	rpu.mutation.SetName(s)
	return rpu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (rpu *ResourceProviderUpdate) SetNillableName(s *string) *ResourceProviderUpdate {
	if s != nil {
		rpu.SetName(*s)
	}
	return rpu
}

// SetDescription sets the "description" field.
func (rpu *ResourceProviderUpdate) SetDescription(s string) *ResourceProviderUpdate {
	rpu.mutation.SetDescription(s)
	return rpu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rpu *ResourceProviderUpdate) SetNillableDescription(s *string) *ResourceProviderUpdate {
	if s != nil {
		rpu.SetDescription(*s)
	}
	return rpu
}

// ClearDescription clears the value of the "description" field.
func (rpu *ResourceProviderUpdate) ClearDescription() *ResourceProviderUpdate {
	rpu.mutation.ClearDescription()
	return rpu
}

// Mutation returns the ResourceProviderMutation object of the builder.
func (rpu *ResourceProviderUpdate) Mutation() *ResourceProviderMutation {
	return rpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rpu *ResourceProviderUpdate) Save(ctx context.Context) (int, error) {
	rpu.defaults()
	return withHooks(ctx, rpu.sqlSave, rpu.mutation, rpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rpu *ResourceProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := rpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rpu *ResourceProviderUpdate) Exec(ctx context.Context) error {
	_, err := rpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpu *ResourceProviderUpdate) ExecX(ctx context.Context) {
	if err := rpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rpu *ResourceProviderUpdate) defaults() {
	if _, ok := rpu.mutation.UpdatedAt(); !ok {
		v := resourceprovider.UpdateDefaultUpdatedAt()
		rpu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rpu *ResourceProviderUpdate) check() error {
	if v, ok := rpu.mutation.Name(); ok {
		if err := resourceprovider.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "ResourceProvider.name": %w`, err)}
		}
	}
	return nil
}

func (rpu *ResourceProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(resourceprovider.Table, resourceprovider.Columns, sqlgraph.NewFieldSpec(resourceprovider.FieldID, field.TypeString))
	if ps := rpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rpu.mutation.UpdatedAt(); ok {
		_spec.SetField(resourceprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rpu.mutation.Name(); ok {
		_spec.SetField(resourceprovider.FieldName, field.TypeString, value)
	}
	if value, ok := rpu.mutation.Description(); ok {
		_spec.SetField(resourceprovider.FieldDescription, field.TypeString, value)
	}
	if rpu.mutation.DescriptionCleared() {
		_spec.ClearField(resourceprovider.FieldDescription, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourceprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rpu.mutation.done = true
	return n, nil
}

// ResourceProviderUpdateOne is the builder for updating a single ResourceProvider entity.
type ResourceProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ResourceProviderMutation
}

// SetName sets the "name" field.
func (rpuo *ResourceProviderUpdateOne) SetName(s string) *ResourceProviderUpdateOne {
	rpuo.mutation.SetName(s)
	return rpuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (rpuo *ResourceProviderUpdateOne) SetNillableName(s *string) *ResourceProviderUpdateOne {
	if s != nil {
		rpuo.SetName(*s)
	}
	return rpuo
}

// SetDescription sets the "description" field.
func (rpuo *ResourceProviderUpdateOne) SetDescription(s string) *ResourceProviderUpdateOne {
	rpuo.mutation.SetDescription(s)
	return rpuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rpuo *ResourceProviderUpdateOne) SetNillableDescription(s *string) *ResourceProviderUpdateOne {
	if s != nil {
		rpuo.SetDescription(*s)
	}
	return rpuo
}

// ClearDescription clears the value of the "description" field.
func (rpuo *ResourceProviderUpdateOne) ClearDescription() *ResourceProviderUpdateOne {
	rpuo.mutation.ClearDescription()
	return rpuo
}

// Mutation returns the ResourceProviderMutation object of the builder.
func (rpuo *ResourceProviderUpdateOne) Mutation() *ResourceProviderMutation {
	return rpuo.mutation
}

// Where appends a list predicates to the ResourceProviderUpdate builder.
func (rpuo *ResourceProviderUpdateOne) Where(ps ...predicate.ResourceProvider) *ResourceProviderUpdateOne {
	rpuo.mutation.Where(ps...)
	return rpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rpuo *ResourceProviderUpdateOne) Select(field string, fields ...string) *ResourceProviderUpdateOne {
	rpuo.fields = append([]string{field}, fields...)
	return rpuo
}

// Save executes the query and returns the updated ResourceProvider entity.
func (rpuo *ResourceProviderUpdateOne) Save(ctx context.Context) (*ResourceProvider, error) {
	rpuo.defaults()
	return withHooks(ctx, rpuo.sqlSave, rpuo.mutation, rpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rpuo *ResourceProviderUpdateOne) SaveX(ctx context.Context) *ResourceProvider {
	node, err := rpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rpuo *ResourceProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := rpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpuo *ResourceProviderUpdateOne) ExecX(ctx context.Context) {
	if err := rpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rpuo *ResourceProviderUpdateOne) defaults() {
	if _, ok := rpuo.mutation.UpdatedAt(); !ok {
		v := resourceprovider.UpdateDefaultUpdatedAt()
		rpuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rpuo *ResourceProviderUpdateOne) check() error {
	if v, ok := rpuo.mutation.Name(); ok {
		if err := resourceprovider.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "ResourceProvider.name": %w`, err)}
		}
	}
	return nil
}

func (rpuo *ResourceProviderUpdateOne) sqlSave(ctx context.Context) (_node *ResourceProvider, err error) {
	if err := rpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(resourceprovider.Table, resourceprovider.Columns, sqlgraph.NewFieldSpec(resourceprovider.FieldID, field.TypeString))
	id, ok := rpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "ResourceProvider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resourceprovider.FieldID)
		for _, f := range fields {
			if !resourceprovider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != resourceprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(resourceprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rpuo.mutation.Name(); ok {
		_spec.SetField(resourceprovider.FieldName, field.TypeString, value)
	}
	if value, ok := rpuo.mutation.Description(); ok {
		_spec.SetField(resourceprovider.FieldDescription, field.TypeString, value)
	}
	if rpuo.mutation.DescriptionCleared() {
		_spec.ClearField(resourceprovider.FieldDescription, field.TypeString)
	}
	_node = &ResourceProvider{config: rpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourceprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rpuo.mutation.done = true
	return _node, nil
}
