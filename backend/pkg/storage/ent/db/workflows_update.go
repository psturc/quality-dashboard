// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/workflows"
	"github.com/google/uuid"
)

// WorkflowsUpdate is the builder for updating Workflows entities.
type WorkflowsUpdate struct {
	config
	hooks    []Hook
	mutation *WorkflowsMutation
}

// Where appends a list predicates to the WorkflowsUpdate builder.
func (wu *WorkflowsUpdate) Where(ps ...predicate.Workflows) *WorkflowsUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetWorkflowID sets the "workflow_id" field.
func (wu *WorkflowsUpdate) SetWorkflowID(u uuid.UUID) *WorkflowsUpdate {
	wu.mutation.SetWorkflowID(u)
	return wu
}

// SetWorkflowName sets the "workflow_name" field.
func (wu *WorkflowsUpdate) SetWorkflowName(s string) *WorkflowsUpdate {
	wu.mutation.SetWorkflowName(s)
	return wu
}

// SetBadgeURL sets the "badge_url" field.
func (wu *WorkflowsUpdate) SetBadgeURL(s string) *WorkflowsUpdate {
	wu.mutation.SetBadgeURL(s)
	return wu
}

// SetHTMLURL sets the "html_url" field.
func (wu *WorkflowsUpdate) SetHTMLURL(s string) *WorkflowsUpdate {
	wu.mutation.SetHTMLURL(s)
	return wu
}

// SetJobURL sets the "job_url" field.
func (wu *WorkflowsUpdate) SetJobURL(s string) *WorkflowsUpdate {
	wu.mutation.SetJobURL(s)
	return wu
}

// SetState sets the "state" field.
func (wu *WorkflowsUpdate) SetState(s string) *WorkflowsUpdate {
	wu.mutation.SetState(s)
	return wu
}

// SetWorkflowsID sets the "workflows" edge to the Repository entity by ID.
func (wu *WorkflowsUpdate) SetWorkflowsID(id uuid.UUID) *WorkflowsUpdate {
	wu.mutation.SetWorkflowsID(id)
	return wu
}

// SetNillableWorkflowsID sets the "workflows" edge to the Repository entity by ID if the given value is not nil.
func (wu *WorkflowsUpdate) SetNillableWorkflowsID(id *uuid.UUID) *WorkflowsUpdate {
	if id != nil {
		wu = wu.SetWorkflowsID(*id)
	}
	return wu
}

// SetWorkflows sets the "workflows" edge to the Repository entity.
func (wu *WorkflowsUpdate) SetWorkflows(r *Repository) *WorkflowsUpdate {
	return wu.SetWorkflowsID(r.ID)
}

// Mutation returns the WorkflowsMutation object of the builder.
func (wu *WorkflowsUpdate) Mutation() *WorkflowsMutation {
	return wu.mutation
}

// ClearWorkflows clears the "workflows" edge to the Repository entity.
func (wu *WorkflowsUpdate) ClearWorkflows() *WorkflowsUpdate {
	wu.mutation.ClearWorkflows()
	return wu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkflowsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkflowsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkflowsUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkflowsUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkflowsUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WorkflowsUpdate) check() error {
	if v, ok := wu.mutation.WorkflowName(); ok {
		if err := workflows.WorkflowNameValidator(v); err != nil {
			return &ValidationError{Name: "workflow_name", err: fmt.Errorf(`db: validator failed for field "Workflows.workflow_name": %w`, err)}
		}
	}
	return nil
}

func (wu *WorkflowsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workflows.Table,
			Columns: workflows.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workflows.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.WorkflowID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: workflows.FieldWorkflowID,
		})
	}
	if value, ok := wu.mutation.WorkflowName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflows.FieldWorkflowName,
		})
	}
	if value, ok := wu.mutation.BadgeURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflows.FieldBadgeURL,
		})
	}
	if value, ok := wu.mutation.HTMLURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflows.FieldHTMLURL,
		})
	}
	if value, ok := wu.mutation.JobURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflows.FieldJobURL,
		})
	}
	if value, ok := wu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflows.FieldState,
		})
	}
	if wu.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflows.WorkflowsTable,
			Columns: []string{workflows.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.WorkflowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflows.WorkflowsTable,
			Columns: []string{workflows.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflows.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WorkflowsUpdateOne is the builder for updating a single Workflows entity.
type WorkflowsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkflowsMutation
}

// SetWorkflowID sets the "workflow_id" field.
func (wuo *WorkflowsUpdateOne) SetWorkflowID(u uuid.UUID) *WorkflowsUpdateOne {
	wuo.mutation.SetWorkflowID(u)
	return wuo
}

// SetWorkflowName sets the "workflow_name" field.
func (wuo *WorkflowsUpdateOne) SetWorkflowName(s string) *WorkflowsUpdateOne {
	wuo.mutation.SetWorkflowName(s)
	return wuo
}

// SetBadgeURL sets the "badge_url" field.
func (wuo *WorkflowsUpdateOne) SetBadgeURL(s string) *WorkflowsUpdateOne {
	wuo.mutation.SetBadgeURL(s)
	return wuo
}

// SetHTMLURL sets the "html_url" field.
func (wuo *WorkflowsUpdateOne) SetHTMLURL(s string) *WorkflowsUpdateOne {
	wuo.mutation.SetHTMLURL(s)
	return wuo
}

// SetJobURL sets the "job_url" field.
func (wuo *WorkflowsUpdateOne) SetJobURL(s string) *WorkflowsUpdateOne {
	wuo.mutation.SetJobURL(s)
	return wuo
}

// SetState sets the "state" field.
func (wuo *WorkflowsUpdateOne) SetState(s string) *WorkflowsUpdateOne {
	wuo.mutation.SetState(s)
	return wuo
}

// SetWorkflowsID sets the "workflows" edge to the Repository entity by ID.
func (wuo *WorkflowsUpdateOne) SetWorkflowsID(id uuid.UUID) *WorkflowsUpdateOne {
	wuo.mutation.SetWorkflowsID(id)
	return wuo
}

// SetNillableWorkflowsID sets the "workflows" edge to the Repository entity by ID if the given value is not nil.
func (wuo *WorkflowsUpdateOne) SetNillableWorkflowsID(id *uuid.UUID) *WorkflowsUpdateOne {
	if id != nil {
		wuo = wuo.SetWorkflowsID(*id)
	}
	return wuo
}

// SetWorkflows sets the "workflows" edge to the Repository entity.
func (wuo *WorkflowsUpdateOne) SetWorkflows(r *Repository) *WorkflowsUpdateOne {
	return wuo.SetWorkflowsID(r.ID)
}

// Mutation returns the WorkflowsMutation object of the builder.
func (wuo *WorkflowsUpdateOne) Mutation() *WorkflowsMutation {
	return wuo.mutation
}

// ClearWorkflows clears the "workflows" edge to the Repository entity.
func (wuo *WorkflowsUpdateOne) ClearWorkflows() *WorkflowsUpdateOne {
	wuo.mutation.ClearWorkflows()
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorkflowsUpdateOne) Select(field string, fields ...string) *WorkflowsUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Workflows entity.
func (wuo *WorkflowsUpdateOne) Save(ctx context.Context) (*Workflows, error) {
	var (
		err  error
		node *Workflows
	)
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkflowsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkflowsUpdateOne) SaveX(ctx context.Context) *Workflows {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkflowsUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkflowsUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WorkflowsUpdateOne) check() error {
	if v, ok := wuo.mutation.WorkflowName(); ok {
		if err := workflows.WorkflowNameValidator(v); err != nil {
			return &ValidationError{Name: "workflow_name", err: fmt.Errorf(`db: validator failed for field "Workflows.workflow_name": %w`, err)}
		}
	}
	return nil
}

func (wuo *WorkflowsUpdateOne) sqlSave(ctx context.Context) (_node *Workflows, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workflows.Table,
			Columns: workflows.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workflows.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Workflows.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflows.FieldID)
		for _, f := range fields {
			if !workflows.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != workflows.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.WorkflowID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: workflows.FieldWorkflowID,
		})
	}
	if value, ok := wuo.mutation.WorkflowName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflows.FieldWorkflowName,
		})
	}
	if value, ok := wuo.mutation.BadgeURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflows.FieldBadgeURL,
		})
	}
	if value, ok := wuo.mutation.HTMLURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflows.FieldHTMLURL,
		})
	}
	if value, ok := wuo.mutation.JobURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflows.FieldJobURL,
		})
	}
	if value, ok := wuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflows.FieldState,
		})
	}
	if wuo.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflows.WorkflowsTable,
			Columns: []string{workflows.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.WorkflowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflows.WorkflowsTable,
			Columns: []string{workflows.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Workflows{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflows.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
