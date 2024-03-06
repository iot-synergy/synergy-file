// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/iot-synergy/synergy-file/ent/cloudfiletag"
	"github.com/iot-synergy/synergy-file/ent/predicate"
)

// CloudFileTagDelete is the builder for deleting a CloudFileTag entity.
type CloudFileTagDelete struct {
	config
	hooks    []Hook
	mutation *CloudFileTagMutation
}

// Where appends a list predicates to the CloudFileTagDelete builder.
func (cftd *CloudFileTagDelete) Where(ps ...predicate.CloudFileTag) *CloudFileTagDelete {
	cftd.mutation.Where(ps...)
	return cftd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cftd *CloudFileTagDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cftd.sqlExec, cftd.mutation, cftd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cftd *CloudFileTagDelete) ExecX(ctx context.Context) int {
	n, err := cftd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cftd *CloudFileTagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cloudfiletag.Table, sqlgraph.NewFieldSpec(cloudfiletag.FieldID, field.TypeUint64))
	if ps := cftd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cftd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cftd.mutation.done = true
	return affected, err
}

// CloudFileTagDeleteOne is the builder for deleting a single CloudFileTag entity.
type CloudFileTagDeleteOne struct {
	cftd *CloudFileTagDelete
}

// Where appends a list predicates to the CloudFileTagDelete builder.
func (cftdo *CloudFileTagDeleteOne) Where(ps ...predicate.CloudFileTag) *CloudFileTagDeleteOne {
	cftdo.cftd.mutation.Where(ps...)
	return cftdo
}

// Exec executes the deletion query.
func (cftdo *CloudFileTagDeleteOne) Exec(ctx context.Context) error {
	n, err := cftdo.cftd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cloudfiletag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cftdo *CloudFileTagDeleteOne) ExecX(ctx context.Context) {
	if err := cftdo.Exec(ctx); err != nil {
		panic(err)
	}
}
