package items

import "github.com/xanzy/go-gitlab"

type Course struct {
	gitlab *gitlab.Group
	year   *Year
}

func (b *Builder) WriteCourse(y *Year, n string) *Course {
	cgo := &gitlab.CreateGroupOptions{
		Name:     gitlab.String(n),
		ParentID: gitlab.Int(y.Gitlab.ID),
	}
	c, _ := b.gitlab.AddGroup(cgo)

	course := &Course{
		gitlab: c,
		year:   y,
	}

	return course
}
