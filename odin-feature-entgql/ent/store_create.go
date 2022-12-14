// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"odin/ent/store"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StoreCreate is the builder for creating a Store entity.
type StoreCreate struct {
	config
	mutation *StoreMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *StoreCreate) SetName(s string) *StoreCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetLocation sets the "location" field.
func (sc *StoreCreate) SetLocation(s string) *StoreCreate {
	sc.mutation.SetLocation(s)
	return sc
}

// SetFood sets the "food" field.
func (sc *StoreCreate) SetFood(s []string) *StoreCreate {
	sc.mutation.SetFood(s)
	return sc
}

// SetOnFoot sets the "on_foot" field.
func (sc *StoreCreate) SetOnFoot(i int) *StoreCreate {
	sc.mutation.SetOnFoot(i)
	return sc
}

// SetNillableOnFoot sets the "on_foot" field if the given value is not nil.
func (sc *StoreCreate) SetNillableOnFoot(i *int) *StoreCreate {
	if i != nil {
		sc.SetOnFoot(*i)
	}
	return sc
}

// SetSentAt sets the "sent_at" field.
func (sc *StoreCreate) SetSentAt(t time.Time) *StoreCreate {
	sc.mutation.SetSentAt(t)
	return sc
}

// SetNillableSentAt sets the "sent_at" field if the given value is not nil.
func (sc *StoreCreate) SetNillableSentAt(t *time.Time) *StoreCreate {
	if t != nil {
		sc.SetSentAt(*t)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *StoreCreate) SetCreatedAt(t time.Time) *StoreCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StoreCreate) SetNillableCreatedAt(t *time.Time) *StoreCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StoreCreate) SetUpdatedAt(t time.Time) *StoreCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *StoreCreate) SetNillableUpdatedAt(t *time.Time) *StoreCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// Mutation returns the StoreMutation object of the builder.
func (sc *StoreCreate) Mutation() *StoreMutation {
	return sc.mutation
}

// Save creates the Store in the database.
func (sc *StoreCreate) Save(ctx context.Context) (*Store, error) {
	var (
		err  error
		node *Store
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Store)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from StoreMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StoreCreate) SaveX(ctx context.Context) *Store {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StoreCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StoreCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StoreCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := store.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := store.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StoreCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Store.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := store.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Store.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`ent: missing required field "Store.location"`)}
	}
	if v, ok := sc.mutation.Location(); ok {
		if err := store.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`ent: validator failed for field "Store.location": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Food(); !ok {
		return &ValidationError{Name: "food", err: errors.New(`ent: missing required field "Store.food"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Store.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Store.updated_at"`)}
	}
	return nil
}

func (sc *StoreCreate) sqlSave(ctx context.Context) (*Store, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *StoreCreate) createSpec() (*Store, *sqlgraph.CreateSpec) {
	var (
		_node = &Store{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: store.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: store.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: store.FieldName,
		})
		_node.Name = value
	}
	if value, ok := sc.mutation.Location(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: store.FieldLocation,
		})
		_node.Location = value
	}
	if value, ok := sc.mutation.Food(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: store.FieldFood,
		})
		_node.Food = value
	}
	if value, ok := sc.mutation.OnFoot(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: store.FieldOnFoot,
		})
		_node.OnFoot = value
	}
	if value, ok := sc.mutation.SentAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: store.FieldSentAt,
		})
		_node.SentAt = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: store.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: store.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// StoreCreateBulk is the builder for creating many Store entities in bulk.
type StoreCreateBulk struct {
	config
	builders []*StoreCreate
}

// Save creates the Store entities in the database.
func (scb *StoreCreateBulk) Save(ctx context.Context) ([]*Store, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Store, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StoreMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StoreCreateBulk) SaveX(ctx context.Context) []*Store {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StoreCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StoreCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
