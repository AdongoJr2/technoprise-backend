package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// BlogPost holds the schema definition for the BlogPost entity.
type BlogPost struct {
	ent.Schema
}

// Fields of the BlogPost.
func (BlogPost) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("slug").Unique().NotEmpty(),
		field.Text("content").NotEmpty(),
		field.String("excerpt").MaxLen(160).NotEmpty(),
	}
}

// Edges of the BlogPost.
func (BlogPost) Edges() []ent.Edge {
	return nil
}

// Mixin of the BlogPost.
func (BlogPost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Indexes of the BlogPost.
func (BlogPost) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("slug"),
	}
}
