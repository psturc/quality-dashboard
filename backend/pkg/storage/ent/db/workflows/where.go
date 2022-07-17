// Code generated by entc, DO NOT EDIT.

package workflows

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// WorkflowID applies equality check predicate on the "workflow_id" field. It's identical to WorkflowIDEQ.
func WorkflowID(v uuid.UUID) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkflowID), v))
	})
}

// WorkflowName applies equality check predicate on the "workflow_name" field. It's identical to WorkflowNameEQ.
func WorkflowName(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkflowName), v))
	})
}

// BadgeURL applies equality check predicate on the "badge_url" field. It's identical to BadgeURLEQ.
func BadgeURL(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBadgeURL), v))
	})
}

// HTMLURL applies equality check predicate on the "html_url" field. It's identical to HTMLURLEQ.
func HTMLURL(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHTMLURL), v))
	})
}

// JobURL applies equality check predicate on the "job_url" field. It's identical to JobURLEQ.
func JobURL(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJobURL), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// WorkflowIDEQ applies the EQ predicate on the "workflow_id" field.
func WorkflowIDEQ(v uuid.UUID) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkflowID), v))
	})
}

// WorkflowIDNEQ applies the NEQ predicate on the "workflow_id" field.
func WorkflowIDNEQ(v uuid.UUID) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWorkflowID), v))
	})
}

// WorkflowIDIn applies the In predicate on the "workflow_id" field.
func WorkflowIDIn(vs ...uuid.UUID) predicate.Workflows {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWorkflowID), v...))
	})
}

// WorkflowIDNotIn applies the NotIn predicate on the "workflow_id" field.
func WorkflowIDNotIn(vs ...uuid.UUID) predicate.Workflows {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWorkflowID), v...))
	})
}

// WorkflowIDGT applies the GT predicate on the "workflow_id" field.
func WorkflowIDGT(v uuid.UUID) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWorkflowID), v))
	})
}

// WorkflowIDGTE applies the GTE predicate on the "workflow_id" field.
func WorkflowIDGTE(v uuid.UUID) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWorkflowID), v))
	})
}

// WorkflowIDLT applies the LT predicate on the "workflow_id" field.
func WorkflowIDLT(v uuid.UUID) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWorkflowID), v))
	})
}

// WorkflowIDLTE applies the LTE predicate on the "workflow_id" field.
func WorkflowIDLTE(v uuid.UUID) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWorkflowID), v))
	})
}

// WorkflowNameEQ applies the EQ predicate on the "workflow_name" field.
func WorkflowNameEQ(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkflowName), v))
	})
}

// WorkflowNameNEQ applies the NEQ predicate on the "workflow_name" field.
func WorkflowNameNEQ(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWorkflowName), v))
	})
}

// WorkflowNameIn applies the In predicate on the "workflow_name" field.
func WorkflowNameIn(vs ...string) predicate.Workflows {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWorkflowName), v...))
	})
}

// WorkflowNameNotIn applies the NotIn predicate on the "workflow_name" field.
func WorkflowNameNotIn(vs ...string) predicate.Workflows {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWorkflowName), v...))
	})
}

// WorkflowNameGT applies the GT predicate on the "workflow_name" field.
func WorkflowNameGT(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWorkflowName), v))
	})
}

// WorkflowNameGTE applies the GTE predicate on the "workflow_name" field.
func WorkflowNameGTE(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWorkflowName), v))
	})
}

// WorkflowNameLT applies the LT predicate on the "workflow_name" field.
func WorkflowNameLT(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWorkflowName), v))
	})
}

// WorkflowNameLTE applies the LTE predicate on the "workflow_name" field.
func WorkflowNameLTE(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWorkflowName), v))
	})
}

// WorkflowNameContains applies the Contains predicate on the "workflow_name" field.
func WorkflowNameContains(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWorkflowName), v))
	})
}

// WorkflowNameHasPrefix applies the HasPrefix predicate on the "workflow_name" field.
func WorkflowNameHasPrefix(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWorkflowName), v))
	})
}

// WorkflowNameHasSuffix applies the HasSuffix predicate on the "workflow_name" field.
func WorkflowNameHasSuffix(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWorkflowName), v))
	})
}

// WorkflowNameEqualFold applies the EqualFold predicate on the "workflow_name" field.
func WorkflowNameEqualFold(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWorkflowName), v))
	})
}

// WorkflowNameContainsFold applies the ContainsFold predicate on the "workflow_name" field.
func WorkflowNameContainsFold(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWorkflowName), v))
	})
}

// BadgeURLEQ applies the EQ predicate on the "badge_url" field.
func BadgeURLEQ(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBadgeURL), v))
	})
}

// BadgeURLNEQ applies the NEQ predicate on the "badge_url" field.
func BadgeURLNEQ(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBadgeURL), v))
	})
}

// BadgeURLIn applies the In predicate on the "badge_url" field.
func BadgeURLIn(vs ...string) predicate.Workflows {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBadgeURL), v...))
	})
}

// BadgeURLNotIn applies the NotIn predicate on the "badge_url" field.
func BadgeURLNotIn(vs ...string) predicate.Workflows {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBadgeURL), v...))
	})
}

// BadgeURLGT applies the GT predicate on the "badge_url" field.
func BadgeURLGT(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBadgeURL), v))
	})
}

// BadgeURLGTE applies the GTE predicate on the "badge_url" field.
func BadgeURLGTE(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBadgeURL), v))
	})
}

// BadgeURLLT applies the LT predicate on the "badge_url" field.
func BadgeURLLT(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBadgeURL), v))
	})
}

// BadgeURLLTE applies the LTE predicate on the "badge_url" field.
func BadgeURLLTE(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBadgeURL), v))
	})
}

// BadgeURLContains applies the Contains predicate on the "badge_url" field.
func BadgeURLContains(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBadgeURL), v))
	})
}

// BadgeURLHasPrefix applies the HasPrefix predicate on the "badge_url" field.
func BadgeURLHasPrefix(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBadgeURL), v))
	})
}

// BadgeURLHasSuffix applies the HasSuffix predicate on the "badge_url" field.
func BadgeURLHasSuffix(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBadgeURL), v))
	})
}

// BadgeURLEqualFold applies the EqualFold predicate on the "badge_url" field.
func BadgeURLEqualFold(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBadgeURL), v))
	})
}

// BadgeURLContainsFold applies the ContainsFold predicate on the "badge_url" field.
func BadgeURLContainsFold(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBadgeURL), v))
	})
}

// HTMLURLEQ applies the EQ predicate on the "html_url" field.
func HTMLURLEQ(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHTMLURL), v))
	})
}

// HTMLURLNEQ applies the NEQ predicate on the "html_url" field.
func HTMLURLNEQ(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHTMLURL), v))
	})
}

// HTMLURLIn applies the In predicate on the "html_url" field.
func HTMLURLIn(vs ...string) predicate.Workflows {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHTMLURL), v...))
	})
}

// HTMLURLNotIn applies the NotIn predicate on the "html_url" field.
func HTMLURLNotIn(vs ...string) predicate.Workflows {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHTMLURL), v...))
	})
}

// HTMLURLGT applies the GT predicate on the "html_url" field.
func HTMLURLGT(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHTMLURL), v))
	})
}

// HTMLURLGTE applies the GTE predicate on the "html_url" field.
func HTMLURLGTE(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHTMLURL), v))
	})
}

// HTMLURLLT applies the LT predicate on the "html_url" field.
func HTMLURLLT(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHTMLURL), v))
	})
}

// HTMLURLLTE applies the LTE predicate on the "html_url" field.
func HTMLURLLTE(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHTMLURL), v))
	})
}

// HTMLURLContains applies the Contains predicate on the "html_url" field.
func HTMLURLContains(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHTMLURL), v))
	})
}

// HTMLURLHasPrefix applies the HasPrefix predicate on the "html_url" field.
func HTMLURLHasPrefix(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHTMLURL), v))
	})
}

// HTMLURLHasSuffix applies the HasSuffix predicate on the "html_url" field.
func HTMLURLHasSuffix(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHTMLURL), v))
	})
}

// HTMLURLEqualFold applies the EqualFold predicate on the "html_url" field.
func HTMLURLEqualFold(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHTMLURL), v))
	})
}

// HTMLURLContainsFold applies the ContainsFold predicate on the "html_url" field.
func HTMLURLContainsFold(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHTMLURL), v))
	})
}

// JobURLEQ applies the EQ predicate on the "job_url" field.
func JobURLEQ(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJobURL), v))
	})
}

// JobURLNEQ applies the NEQ predicate on the "job_url" field.
func JobURLNEQ(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldJobURL), v))
	})
}

// JobURLIn applies the In predicate on the "job_url" field.
func JobURLIn(vs ...string) predicate.Workflows {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldJobURL), v...))
	})
}

// JobURLNotIn applies the NotIn predicate on the "job_url" field.
func JobURLNotIn(vs ...string) predicate.Workflows {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldJobURL), v...))
	})
}

// JobURLGT applies the GT predicate on the "job_url" field.
func JobURLGT(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldJobURL), v))
	})
}

// JobURLGTE applies the GTE predicate on the "job_url" field.
func JobURLGTE(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldJobURL), v))
	})
}

// JobURLLT applies the LT predicate on the "job_url" field.
func JobURLLT(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldJobURL), v))
	})
}

// JobURLLTE applies the LTE predicate on the "job_url" field.
func JobURLLTE(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldJobURL), v))
	})
}

// JobURLContains applies the Contains predicate on the "job_url" field.
func JobURLContains(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldJobURL), v))
	})
}

// JobURLHasPrefix applies the HasPrefix predicate on the "job_url" field.
func JobURLHasPrefix(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldJobURL), v))
	})
}

// JobURLHasSuffix applies the HasSuffix predicate on the "job_url" field.
func JobURLHasSuffix(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldJobURL), v))
	})
}

// JobURLEqualFold applies the EqualFold predicate on the "job_url" field.
func JobURLEqualFold(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldJobURL), v))
	})
}

// JobURLContainsFold applies the ContainsFold predicate on the "job_url" field.
func JobURLContainsFold(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldJobURL), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.Workflows {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.Workflows {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflows(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// HasWorkflows applies the HasEdge predicate on the "workflows" edge.
func HasWorkflows() predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowsTable, RepositoryFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkflowsTable, WorkflowsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkflowsWith applies the HasEdge predicate on the "workflows" edge with a given conditions (other predicates).
func HasWorkflowsWith(preds ...predicate.Repository) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowsInverseTable, RepositoryFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkflowsTable, WorkflowsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Workflows) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Workflows) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
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
func Not(p predicate.Workflows) predicate.Workflows {
	return predicate.Workflows(func(s *sql.Selector) {
		p(s.Not())
	})
}
