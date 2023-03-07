package items

import (
	"nova/user"

	"github.com/xanzy/go-gitlab"
)

type Project struct {
	Gitlab      *gitlab.Project `json:"gitlab,omitempty"`
	Student     *user.Student   `json:"student,omitempty"`
	Supervisior *user.Member    `json:"supervisior,omitempty"`
	Course      *Course         `json:"course,omitempty"`
}

func (b *Builder) WriteProject(s *user.Student, c *Course, n string) (*Project, error) {
	cpo := &gitlab.CreateProjectOptions{
		Name:        gitlab.String(n),
		NamespaceID: gitlab.Int(c.gitlab.ID),
	}
	p, err := b.gitlab.AddProject(cpo)
	if err != nil {
		return nil, err
	}

	apmo := &gitlab.AddProjectMemberOptions{
		UserID:      gitlab.Int(s.Gitlab.ID),
		AccessLevel: gitlab.AccessLevel(30),
	}
	err = b.gitlab.AddMember(p, apmo)
	if err != nil {
		return nil, err
	}

	project := &Project{
		Gitlab:  p,
		Student: s,
		Course:  c,
	}

	return project, nil
}

func (b *Builder) ListProjects(c *Course) ([]*Project, error) {
	projects, err := b.gitlab.ListProjects(c.gitlab)
	if err != nil {
		return nil, err
	}

	var p []*Project
	for _, project := range projects {
		p = append(p, &Project{
			Gitlab: project,
		})
	}
	return p, nil
}

func (b *Builder) ReadProject(id int) (*Project, error) {
	project, err := b.gitlab.GetProject(id)
	if err != nil {
		return nil, err
	}

	p := &Project{
		Gitlab: project,
	}
	return p, nil
}

func (b *Builder) AddSupervisor(p *Project, m *user.Member) error {
	apmo := &gitlab.AddProjectMemberOptions{
		UserID:      gitlab.Int(m.Gitlab.ID),
		AccessLevel: gitlab.AccessLevel(40),
	}
	err := b.gitlab.AddMember(p.Gitlab, apmo)
	if err != nil {
		return err
	}

	return nil
}
