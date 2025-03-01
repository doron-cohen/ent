// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/json/ent/schema"
	"entgo.io/ent/entc/integration/json/ent/user"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetT sets the "t" field.
func (uc *UserCreate) SetT(s *schema.T) *UserCreate {
	uc.mutation.SetT(s)
	return uc
}

// SetURL sets the "url" field.
func (uc *UserCreate) SetURL(u *url.URL) *UserCreate {
	uc.mutation.SetURL(u)
	return uc
}

// SetURLs sets the "URLs" field.
func (uc *UserCreate) SetURLs(u []*url.URL) *UserCreate {
	uc.mutation.SetURLs(u)
	return uc
}

// SetRaw sets the "raw" field.
func (uc *UserCreate) SetRaw(jm json.RawMessage) *UserCreate {
	uc.mutation.SetRaw(jm)
	return uc
}

// SetDirs sets the "dirs" field.
func (uc *UserCreate) SetDirs(h []http.Dir) *UserCreate {
	uc.mutation.SetDirs(h)
	return uc
}

// SetInts sets the "ints" field.
func (uc *UserCreate) SetInts(i []int) *UserCreate {
	uc.mutation.SetInts(i)
	return uc
}

// SetFloats sets the "floats" field.
func (uc *UserCreate) SetFloats(f []float64) *UserCreate {
	uc.mutation.SetFloats(f)
	return uc
}

// SetStrings sets the "strings" field.
func (uc *UserCreate) SetStrings(s []string) *UserCreate {
	uc.mutation.SetStrings(s)
	return uc
}

// SetAddr sets the "addr" field.
func (uc *UserCreate) SetAddr(s schema.Addr) *UserCreate {
	uc.mutation.SetAddr(s)
	return uc
}

// SetNillableAddr sets the "addr" field if the given value is not nil.
func (uc *UserCreate) SetNillableAddr(s *schema.Addr) *UserCreate {
	if s != nil {
		uc.SetAddr(*s)
	}
	return uc
}

// SetUnknown sets the "unknown" field.
func (uc *UserCreate) SetUnknown(a any) *UserCreate {
	uc.mutation.SetUnknown(a)
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks[*User, UserMutation](ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Dirs(); !ok {
		v := user.DefaultDirs()
		uc.mutation.SetDirs(v)
	}
	if _, ok := uc.mutation.Ints(); !ok {
		v := user.DefaultInts
		uc.mutation.SetInts(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Dirs(); !ok {
		return &ValidationError{Name: "dirs", err: errors.New(`ent: missing required field "User.dirs"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.T(); ok {
		_spec.SetField(user.FieldT, field.TypeJSON, value)
		_node.T = value
	}
	if value, ok := uc.mutation.URL(); ok {
		_spec.SetField(user.FieldURL, field.TypeJSON, value)
		_node.URL = value
	}
	if value, ok := uc.mutation.URLs(); ok {
		_spec.SetField(user.FieldURLs, field.TypeJSON, value)
		_node.URLs = value
	}
	if value, ok := uc.mutation.Raw(); ok {
		_spec.SetField(user.FieldRaw, field.TypeJSON, value)
		_node.Raw = value
	}
	if value, ok := uc.mutation.Dirs(); ok {
		_spec.SetField(user.FieldDirs, field.TypeJSON, value)
		_node.Dirs = value
	}
	if value, ok := uc.mutation.Ints(); ok {
		_spec.SetField(user.FieldInts, field.TypeJSON, value)
		_node.Ints = value
	}
	if value, ok := uc.mutation.Floats(); ok {
		_spec.SetField(user.FieldFloats, field.TypeJSON, value)
		_node.Floats = value
	}
	if value, ok := uc.mutation.Strings(); ok {
		_spec.SetField(user.FieldStrings, field.TypeJSON, value)
		_node.Strings = value
	}
	if value, ok := uc.mutation.Addr(); ok {
		_spec.SetField(user.FieldAddr, field.TypeJSON, value)
		_node.Addr = value
	}
	if value, ok := uc.mutation.Unknown(); ok {
		_spec.SetField(user.FieldUnknown, field.TypeJSON, value)
		_node.Unknown = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
