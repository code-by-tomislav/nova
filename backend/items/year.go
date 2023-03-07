package items

import (
	"github.com/xanzy/go-gitlab"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Year struct {
	Id       primitive.ObjectID   `json:"-"                  bson:"_id"`
	Gitlab   *gitlab.Group        `json:"gitlab,omitempty"`
	Parent   primitive.ObjectID   `json:"parent,omitempty"`
	Children []primitive.ObjectID `json:"children,omitempty"`
}

func (b *Builder) WriteYear(f *Field, n string) *Year {
	cgo := &gitlab.CreateGroupOptions{
		Name:     gitlab.String(n),
		ParentID: gitlab.Int(f.Gitlab.ID),
	}
	c, _ := b.gitlab.AddGroup(cgo)

	year := &Year{
		Gitlab: c,
		Parent: f.Id,
	}

	f.Children = append(f.Children, year.Id)
	b.updateFieldCache(f)

	return year
}
