// Code generated by entc, DO NOT EDIT.

package artist

import (
	"odin/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// ExternalURL applies equality check predicate on the "external_url" field. It's identical to ExternalURLEQ.
func ExternalURL(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExternalURL), v))
	})
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhoneNumber), v))
	})
}

// Discord applies equality check predicate on the "discord" field. It's identical to DiscordEQ.
func Discord(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscord), v))
	})
}

// Recommender applies equality check predicate on the "recommender" field. It's identical to RecommenderEQ.
func Recommender(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecommender), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ExternalURLEQ applies the EQ predicate on the "external_url" field.
func ExternalURLEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExternalURL), v))
	})
}

// ExternalURLNEQ applies the NEQ predicate on the "external_url" field.
func ExternalURLNEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExternalURL), v))
	})
}

// ExternalURLIn applies the In predicate on the "external_url" field.
func ExternalURLIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExternalURL), v...))
	})
}

// ExternalURLNotIn applies the NotIn predicate on the "external_url" field.
func ExternalURLNotIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExternalURL), v...))
	})
}

// ExternalURLGT applies the GT predicate on the "external_url" field.
func ExternalURLGT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExternalURL), v))
	})
}

// ExternalURLGTE applies the GTE predicate on the "external_url" field.
func ExternalURLGTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExternalURL), v))
	})
}

// ExternalURLLT applies the LT predicate on the "external_url" field.
func ExternalURLLT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExternalURL), v))
	})
}

// ExternalURLLTE applies the LTE predicate on the "external_url" field.
func ExternalURLLTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExternalURL), v))
	})
}

// ExternalURLContains applies the Contains predicate on the "external_url" field.
func ExternalURLContains(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExternalURL), v))
	})
}

// ExternalURLHasPrefix applies the HasPrefix predicate on the "external_url" field.
func ExternalURLHasPrefix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExternalURL), v))
	})
}

// ExternalURLHasSuffix applies the HasSuffix predicate on the "external_url" field.
func ExternalURLHasSuffix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExternalURL), v))
	})
}

// ExternalURLEqualFold applies the EqualFold predicate on the "external_url" field.
func ExternalURLEqualFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExternalURL), v))
	})
}

// ExternalURLContainsFold applies the ContainsFold predicate on the "external_url" field.
func ExternalURLContainsFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExternalURL), v))
	})
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhoneNumber), v...))
	})
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhoneNumber), v...))
	})
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhoneNumber), v))
	})
}

// DiscordEQ applies the EQ predicate on the "discord" field.
func DiscordEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscord), v))
	})
}

// DiscordNEQ applies the NEQ predicate on the "discord" field.
func DiscordNEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiscord), v))
	})
}

// DiscordIn applies the In predicate on the "discord" field.
func DiscordIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiscord), v...))
	})
}

// DiscordNotIn applies the NotIn predicate on the "discord" field.
func DiscordNotIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiscord), v...))
	})
}

// DiscordGT applies the GT predicate on the "discord" field.
func DiscordGT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiscord), v))
	})
}

// DiscordGTE applies the GTE predicate on the "discord" field.
func DiscordGTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiscord), v))
	})
}

// DiscordLT applies the LT predicate on the "discord" field.
func DiscordLT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiscord), v))
	})
}

// DiscordLTE applies the LTE predicate on the "discord" field.
func DiscordLTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiscord), v))
	})
}

// DiscordContains applies the Contains predicate on the "discord" field.
func DiscordContains(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDiscord), v))
	})
}

// DiscordHasPrefix applies the HasPrefix predicate on the "discord" field.
func DiscordHasPrefix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDiscord), v))
	})
}

// DiscordHasSuffix applies the HasSuffix predicate on the "discord" field.
func DiscordHasSuffix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDiscord), v))
	})
}

// DiscordIsNil applies the IsNil predicate on the "discord" field.
func DiscordIsNil() predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDiscord)))
	})
}

// DiscordNotNil applies the NotNil predicate on the "discord" field.
func DiscordNotNil() predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDiscord)))
	})
}

// DiscordEqualFold applies the EqualFold predicate on the "discord" field.
func DiscordEqualFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDiscord), v))
	})
}

// DiscordContainsFold applies the ContainsFold predicate on the "discord" field.
func DiscordContainsFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDiscord), v))
	})
}

// RecommenderEQ applies the EQ predicate on the "recommender" field.
func RecommenderEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecommender), v))
	})
}

// RecommenderNEQ applies the NEQ predicate on the "recommender" field.
func RecommenderNEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecommender), v))
	})
}

// RecommenderIn applies the In predicate on the "recommender" field.
func RecommenderIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRecommender), v...))
	})
}

// RecommenderNotIn applies the NotIn predicate on the "recommender" field.
func RecommenderNotIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRecommender), v...))
	})
}

// RecommenderGT applies the GT predicate on the "recommender" field.
func RecommenderGT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecommender), v))
	})
}

// RecommenderGTE applies the GTE predicate on the "recommender" field.
func RecommenderGTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecommender), v))
	})
}

// RecommenderLT applies the LT predicate on the "recommender" field.
func RecommenderLT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecommender), v))
	})
}

// RecommenderLTE applies the LTE predicate on the "recommender" field.
func RecommenderLTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecommender), v))
	})
}

// RecommenderContains applies the Contains predicate on the "recommender" field.
func RecommenderContains(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRecommender), v))
	})
}

// RecommenderHasPrefix applies the HasPrefix predicate on the "recommender" field.
func RecommenderHasPrefix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRecommender), v))
	})
}

// RecommenderHasSuffix applies the HasSuffix predicate on the "recommender" field.
func RecommenderHasSuffix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRecommender), v))
	})
}

// RecommenderIsNil applies the IsNil predicate on the "recommender" field.
func RecommenderIsNil() predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRecommender)))
	})
}

// RecommenderNotNil applies the NotNil predicate on the "recommender" field.
func RecommenderNotNil() predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRecommender)))
	})
}

// RecommenderEqualFold applies the EqualFold predicate on the "recommender" field.
func RecommenderEqualFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRecommender), v))
	})
}

// RecommenderContainsFold applies the ContainsFold predicate on the "recommender" field.
func RecommenderContainsFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRecommender), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Artist) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Artist) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
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
func Not(p predicate.Artist) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		p(s.Not())
	})
}
