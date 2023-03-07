package items

import (
	"github.com/xanzy/go-gitlab"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Field struct {
	Id       primitive.ObjectID
	Gitlab   *gitlab.Group        `json:"gitlab,omitempty"`
	Children []primitive.ObjectID `json:"children,omitempty"`
}

// public functions

func (b *Builder) WriteField(n string) *Field {
	cgo := &gitlab.CreateGroupOptions{
		Name: gitlab.String(n),
	}
	c, _ := b.gitlab.AddGroup(cgo)

	field := &Field{
		Gitlab: c,
	}

	b.mongo.StoreObject("fields", field)

	return field
}

// private functions

func (b *Builder) updateFieldCache(f *Field) {
	b.mongo.RefreshObject("fields", f.Id, f)
}
