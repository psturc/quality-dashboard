// Code generated by entc, DO NOT EDIT.

package codecov

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// RepositoryName applies equality check predicate on the "repository_name" field. It's identical to RepositoryNameEQ.
func RepositoryName(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRepositoryName), v))
	})
}

// GitOrganization applies equality check predicate on the "git_organization" field. It's identical to GitOrganizationEQ.
func GitOrganization(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGitOrganization), v))
	})
}

// CoveragePercentage applies equality check predicate on the "coverage_percentage" field. It's identical to CoveragePercentageEQ.
func CoveragePercentage(v float64) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoveragePercentage), v))
	})
}

// RepositoryNameEQ applies the EQ predicate on the "repository_name" field.
func RepositoryNameEQ(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRepositoryName), v))
	})
}

// RepositoryNameNEQ applies the NEQ predicate on the "repository_name" field.
func RepositoryNameNEQ(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRepositoryName), v))
	})
}

// RepositoryNameIn applies the In predicate on the "repository_name" field.
func RepositoryNameIn(vs ...string) predicate.CodeCov {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodeCov(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRepositoryName), v...))
	})
}

// RepositoryNameNotIn applies the NotIn predicate on the "repository_name" field.
func RepositoryNameNotIn(vs ...string) predicate.CodeCov {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodeCov(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRepositoryName), v...))
	})
}

// RepositoryNameGT applies the GT predicate on the "repository_name" field.
func RepositoryNameGT(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRepositoryName), v))
	})
}

// RepositoryNameGTE applies the GTE predicate on the "repository_name" field.
func RepositoryNameGTE(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRepositoryName), v))
	})
}

// RepositoryNameLT applies the LT predicate on the "repository_name" field.
func RepositoryNameLT(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRepositoryName), v))
	})
}

// RepositoryNameLTE applies the LTE predicate on the "repository_name" field.
func RepositoryNameLTE(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRepositoryName), v))
	})
}

// RepositoryNameContains applies the Contains predicate on the "repository_name" field.
func RepositoryNameContains(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRepositoryName), v))
	})
}

// RepositoryNameHasPrefix applies the HasPrefix predicate on the "repository_name" field.
func RepositoryNameHasPrefix(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRepositoryName), v))
	})
}

// RepositoryNameHasSuffix applies the HasSuffix predicate on the "repository_name" field.
func RepositoryNameHasSuffix(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRepositoryName), v))
	})
}

// RepositoryNameEqualFold applies the EqualFold predicate on the "repository_name" field.
func RepositoryNameEqualFold(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRepositoryName), v))
	})
}

// RepositoryNameContainsFold applies the ContainsFold predicate on the "repository_name" field.
func RepositoryNameContainsFold(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRepositoryName), v))
	})
}

// GitOrganizationEQ applies the EQ predicate on the "git_organization" field.
func GitOrganizationEQ(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGitOrganization), v))
	})
}

// GitOrganizationNEQ applies the NEQ predicate on the "git_organization" field.
func GitOrganizationNEQ(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGitOrganization), v))
	})
}

// GitOrganizationIn applies the In predicate on the "git_organization" field.
func GitOrganizationIn(vs ...string) predicate.CodeCov {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodeCov(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGitOrganization), v...))
	})
}

// GitOrganizationNotIn applies the NotIn predicate on the "git_organization" field.
func GitOrganizationNotIn(vs ...string) predicate.CodeCov {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodeCov(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGitOrganization), v...))
	})
}

// GitOrganizationGT applies the GT predicate on the "git_organization" field.
func GitOrganizationGT(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGitOrganization), v))
	})
}

// GitOrganizationGTE applies the GTE predicate on the "git_organization" field.
func GitOrganizationGTE(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGitOrganization), v))
	})
}

// GitOrganizationLT applies the LT predicate on the "git_organization" field.
func GitOrganizationLT(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGitOrganization), v))
	})
}

// GitOrganizationLTE applies the LTE predicate on the "git_organization" field.
func GitOrganizationLTE(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGitOrganization), v))
	})
}

// GitOrganizationContains applies the Contains predicate on the "git_organization" field.
func GitOrganizationContains(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGitOrganization), v))
	})
}

// GitOrganizationHasPrefix applies the HasPrefix predicate on the "git_organization" field.
func GitOrganizationHasPrefix(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGitOrganization), v))
	})
}

// GitOrganizationHasSuffix applies the HasSuffix predicate on the "git_organization" field.
func GitOrganizationHasSuffix(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGitOrganization), v))
	})
}

// GitOrganizationEqualFold applies the EqualFold predicate on the "git_organization" field.
func GitOrganizationEqualFold(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGitOrganization), v))
	})
}

// GitOrganizationContainsFold applies the ContainsFold predicate on the "git_organization" field.
func GitOrganizationContainsFold(v string) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGitOrganization), v))
	})
}

// CoveragePercentageEQ applies the EQ predicate on the "coverage_percentage" field.
func CoveragePercentageEQ(v float64) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoveragePercentage), v))
	})
}

// CoveragePercentageNEQ applies the NEQ predicate on the "coverage_percentage" field.
func CoveragePercentageNEQ(v float64) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoveragePercentage), v))
	})
}

// CoveragePercentageIn applies the In predicate on the "coverage_percentage" field.
func CoveragePercentageIn(vs ...float64) predicate.CodeCov {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodeCov(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoveragePercentage), v...))
	})
}

// CoveragePercentageNotIn applies the NotIn predicate on the "coverage_percentage" field.
func CoveragePercentageNotIn(vs ...float64) predicate.CodeCov {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodeCov(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoveragePercentage), v...))
	})
}

// CoveragePercentageGT applies the GT predicate on the "coverage_percentage" field.
func CoveragePercentageGT(v float64) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoveragePercentage), v))
	})
}

// CoveragePercentageGTE applies the GTE predicate on the "coverage_percentage" field.
func CoveragePercentageGTE(v float64) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoveragePercentage), v))
	})
}

// CoveragePercentageLT applies the LT predicate on the "coverage_percentage" field.
func CoveragePercentageLT(v float64) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoveragePercentage), v))
	})
}

// CoveragePercentageLTE applies the LTE predicate on the "coverage_percentage" field.
func CoveragePercentageLTE(v float64) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoveragePercentage), v))
	})
}

// HasCodecov applies the HasEdge predicate on the "codecov" edge.
func HasCodecov() predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CodecovTable, RepositoryFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CodecovTable, CodecovColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCodecovWith applies the HasEdge predicate on the "codecov" edge with a given conditions (other predicates).
func HasCodecovWith(preds ...predicate.Repository) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CodecovInverseTable, RepositoryFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CodecovTable, CodecovColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CodeCov) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CodeCov) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
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
func Not(p predicate.CodeCov) predicate.CodeCov {
	return predicate.CodeCov(func(s *sql.Selector) {
		p(s.Not())
	})
}
