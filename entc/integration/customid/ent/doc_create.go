// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/doc"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	"entgo.io/ent/schema/field"
)

// DocCreate is the builder for creating a Doc entity.
type DocCreate struct {
	config
	mutation *DocMutation
	hooks    []Hook
}

// SetText sets the "text" field.
func (dc *DocCreate) SetText(s string) *DocCreate {
	dc.mutation.SetText(s)
	return dc
}

// SetNillableText sets the "text" field if the given value is not nil.
func (dc *DocCreate) SetNillableText(s *string) *DocCreate {
	if s != nil {
		dc.SetText(*s)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DocCreate) SetID(si schema.DocID) *DocCreate {
	dc.mutation.SetID(si)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DocCreate) SetNillableID(si *schema.DocID) *DocCreate {
	if si != nil {
		dc.SetID(*si)
	}
	return dc
}

// SetParentID sets the "parent" edge to the Doc entity by ID.
func (dc *DocCreate) SetParentID(id schema.DocID) *DocCreate {
	dc.mutation.SetParentID(id)
	return dc
}

// SetNillableParentID sets the "parent" edge to the Doc entity by ID if the given value is not nil.
func (dc *DocCreate) SetNillableParentID(id *schema.DocID) *DocCreate {
	if id != nil {
		dc = dc.SetParentID(*id)
	}
	return dc
}

// SetParent sets the "parent" edge to the Doc entity.
func (dc *DocCreate) SetParent(d *Doc) *DocCreate {
	return dc.SetParentID(d.ID)
}

// AddChildIDs adds the "children" edge to the Doc entity by IDs.
func (dc *DocCreate) AddChildIDs(ids ...schema.DocID) *DocCreate {
	dc.mutation.AddChildIDs(ids...)
	return dc
}

// AddChildren adds the "children" edges to the Doc entity.
func (dc *DocCreate) AddChildren(d ...*Doc) *DocCreate {
	ids := make([]schema.DocID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddChildIDs(ids...)
}

// Mutation returns the DocMutation object of the builder.
func (dc *DocCreate) Mutation() *DocMutation {
	return dc.mutation
}

// Save creates the Doc in the database.
func (dc *DocCreate) Save(ctx context.Context) (*Doc, error) {
	var (
		err  error
		node *Doc
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DocMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DocCreate) SaveX(ctx context.Context) *Doc {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DocCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DocCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DocCreate) defaults() {
	if _, ok := dc.mutation.ID(); !ok {
		v := doc.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DocCreate) check() error {
	if v, ok := dc.mutation.ID(); ok {
		if err := doc.IDValidator(string(v)); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "id": %w`, err)}
		}
	}
	return nil
}

func (dc *DocCreate) sqlSave(ctx context.Context) (*Doc, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

func (dc *DocCreate) createSpec() (*Doc, *sqlgraph.CreateSpec) {
	var (
		_node = &Doc{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: doc.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: doc.FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.Text(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doc.FieldText,
		})
		_node.Text = value
	}
	if nodes := dc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doc.ParentTable,
			Columns: []string{doc.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.doc_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doc.ChildrenTable,
			Columns: []string{doc.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DocCreateBulk is the builder for creating many Doc entities in bulk.
type DocCreateBulk struct {
	config
	builders []*DocCreate
}

// Save creates the Doc entities in the database.
func (dcb *DocCreateBulk) Save(ctx context.Context) ([]*Doc, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Doc, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DocMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DocCreateBulk) SaveX(ctx context.Context) []*Doc {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DocCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DocCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
