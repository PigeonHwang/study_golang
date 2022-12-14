// Code generated by ent, DO NOT EDIT.

package member

import (
	"gqlgen-crud/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Nick applies equality check predicate on the "nick" field. It's identical to NickEQ.
func Nick(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNick), v))
	})
}

// Team applies equality check predicate on the "team" field. It's identical to TeamEQ.
func Team(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeam), v))
	})
}

// Detail applies equality check predicate on the "detail" field. It's identical to DetailEQ.
func Detail(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// Img applies equality check predicate on the "img" field. It's identical to ImgEQ.
func Img(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImg), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// NickEQ applies the EQ predicate on the "nick" field.
func NickEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNick), v))
	})
}

// NickNEQ applies the NEQ predicate on the "nick" field.
func NickNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNick), v))
	})
}

// NickIn applies the In predicate on the "nick" field.
func NickIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNick), v...))
	})
}

// NickNotIn applies the NotIn predicate on the "nick" field.
func NickNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNick), v...))
	})
}

// NickGT applies the GT predicate on the "nick" field.
func NickGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNick), v))
	})
}

// NickGTE applies the GTE predicate on the "nick" field.
func NickGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNick), v))
	})
}

// NickLT applies the LT predicate on the "nick" field.
func NickLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNick), v))
	})
}

// NickLTE applies the LTE predicate on the "nick" field.
func NickLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNick), v))
	})
}

// NickContains applies the Contains predicate on the "nick" field.
func NickContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNick), v))
	})
}

// NickHasPrefix applies the HasPrefix predicate on the "nick" field.
func NickHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNick), v))
	})
}

// NickHasSuffix applies the HasSuffix predicate on the "nick" field.
func NickHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNick), v))
	})
}

// NickEqualFold applies the EqualFold predicate on the "nick" field.
func NickEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNick), v))
	})
}

// NickContainsFold applies the ContainsFold predicate on the "nick" field.
func NickContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNick), v))
	})
}

// TeamEQ applies the EQ predicate on the "team" field.
func TeamEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeam), v))
	})
}

// TeamNEQ applies the NEQ predicate on the "team" field.
func TeamNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTeam), v))
	})
}

// TeamIn applies the In predicate on the "team" field.
func TeamIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTeam), v...))
	})
}

// TeamNotIn applies the NotIn predicate on the "team" field.
func TeamNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTeam), v...))
	})
}

// TeamGT applies the GT predicate on the "team" field.
func TeamGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTeam), v))
	})
}

// TeamGTE applies the GTE predicate on the "team" field.
func TeamGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTeam), v))
	})
}

// TeamLT applies the LT predicate on the "team" field.
func TeamLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTeam), v))
	})
}

// TeamLTE applies the LTE predicate on the "team" field.
func TeamLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTeam), v))
	})
}

// TeamContains applies the Contains predicate on the "team" field.
func TeamContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTeam), v))
	})
}

// TeamHasPrefix applies the HasPrefix predicate on the "team" field.
func TeamHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTeam), v))
	})
}

// TeamHasSuffix applies the HasSuffix predicate on the "team" field.
func TeamHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTeam), v))
	})
}

// TeamEqualFold applies the EqualFold predicate on the "team" field.
func TeamEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTeam), v))
	})
}

// TeamContainsFold applies the ContainsFold predicate on the "team" field.
func TeamContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTeam), v))
	})
}

// DetailEQ applies the EQ predicate on the "detail" field.
func DetailEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// DetailNEQ applies the NEQ predicate on the "detail" field.
func DetailNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDetail), v))
	})
}

// DetailIn applies the In predicate on the "detail" field.
func DetailIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDetail), v...))
	})
}

// DetailNotIn applies the NotIn predicate on the "detail" field.
func DetailNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDetail), v...))
	})
}

// DetailGT applies the GT predicate on the "detail" field.
func DetailGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDetail), v))
	})
}

// DetailGTE applies the GTE predicate on the "detail" field.
func DetailGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDetail), v))
	})
}

// DetailLT applies the LT predicate on the "detail" field.
func DetailLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDetail), v))
	})
}

// DetailLTE applies the LTE predicate on the "detail" field.
func DetailLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDetail), v))
	})
}

// DetailContains applies the Contains predicate on the "detail" field.
func DetailContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDetail), v))
	})
}

// DetailHasPrefix applies the HasPrefix predicate on the "detail" field.
func DetailHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDetail), v))
	})
}

// DetailHasSuffix applies the HasSuffix predicate on the "detail" field.
func DetailHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDetail), v))
	})
}

// DetailEqualFold applies the EqualFold predicate on the "detail" field.
func DetailEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDetail), v))
	})
}

// DetailContainsFold applies the ContainsFold predicate on the "detail" field.
func DetailContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDetail), v))
	})
}

// ImgEQ applies the EQ predicate on the "img" field.
func ImgEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImg), v))
	})
}

// ImgNEQ applies the NEQ predicate on the "img" field.
func ImgNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImg), v))
	})
}

// ImgIn applies the In predicate on the "img" field.
func ImgIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldImg), v...))
	})
}

// ImgNotIn applies the NotIn predicate on the "img" field.
func ImgNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldImg), v...))
	})
}

// ImgGT applies the GT predicate on the "img" field.
func ImgGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImg), v))
	})
}

// ImgGTE applies the GTE predicate on the "img" field.
func ImgGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImg), v))
	})
}

// ImgLT applies the LT predicate on the "img" field.
func ImgLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImg), v))
	})
}

// ImgLTE applies the LTE predicate on the "img" field.
func ImgLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImg), v))
	})
}

// ImgContains applies the Contains predicate on the "img" field.
func ImgContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImg), v))
	})
}

// ImgHasPrefix applies the HasPrefix predicate on the "img" field.
func ImgHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImg), v))
	})
}

// ImgHasSuffix applies the HasSuffix predicate on the "img" field.
func ImgHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImg), v))
	})
}

// ImgEqualFold applies the EqualFold predicate on the "img" field.
func ImgEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImg), v))
	})
}

// ImgContainsFold applies the ContainsFold predicate on the "img" field.
func ImgContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImg), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
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
func Not(p predicate.Member) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		p(s.Not())
	})
}
