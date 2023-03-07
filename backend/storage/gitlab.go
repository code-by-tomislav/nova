package storage

import (
	"log"
	"nova/utils"

	"github.com/xanzy/go-gitlab"
)

type Gitlab struct {
	client *gitlab.Client
}

func NewGitlab(c *utils.Configuration) *Gitlab {
	git, err := gitlab.NewClient(c.Gitlab.Token, gitlab.WithBaseURL(c.Gitlab.Host))
	if err != nil {
		log.Fatalf("Failed to create gitlab client: %v", err)
	}

	gl := &Gitlab{client: git}
	return gl
}

func (gl *Gitlab) AddGroup(options *gitlab.CreateGroupOptions) (*gitlab.Group, error) {
	group, _, err := gl.client.Groups.CreateGroup(options)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (gl *Gitlab) AddProject(options *gitlab.CreateProjectOptions) (*gitlab.Project, error) {
	project, _, err := gl.client.Projects.CreateProject(options)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (gl *Gitlab) GetProject(id int) (*gitlab.Project, error) {
	options := &gitlab.GetProjectOptions{}
	project, _, err := gl.client.Projects.GetProject(id, options)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (gl *Gitlab) ListProjects(group *gitlab.Group) ([]*gitlab.Project, error) {
	options := &gitlab.ListGroupProjectsOptions{}
	projects, _, err := gl.client.Groups.ListGroupProjects(group.ID, options)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (gl *Gitlab) AddMember(p *gitlab.Project, options *gitlab.AddProjectMemberOptions) error {
	_, _, err := gl.client.ProjectMembers.AddProjectMember(p.ID, options)
	return err
}

func (gl *Gitlab) GetUser(ldapuid string) (*gitlab.User, error) {
	options := &gitlab.ListUsersOptions{
		ExternalUID: gitlab.String(ldapuid),
	}
	list, _, err := gl.client.Users.ListUsers(options)
	if err != nil {
		return nil, err
	} else if len(list) == 0 {
		return nil, nil
	}

	return list[0], nil
}

func (gl *Gitlab) ListIssues(p *gitlab.Project) ([]*gitlab.Issue, error) {
	options := &gitlab.ListProjectIssuesOptions{}
	issues, _, err := gl.client.Issues.ListProjectIssues(p.ID, options)
	if err != nil {
		return nil, err
	}
	return issues, nil
}

func (gl *Gitlab) GetIssue(p *gitlab.Project, i int) (*gitlab.Issue, error) {
	issue, _, err := gl.client.Issues.GetIssue(p.ID, i)
	if err != nil {
		return nil, err
	}
	return issue, nil
}

func (gl *Gitlab) AddIssue(p *gitlab.Project, options *gitlab.CreateIssueOptions) (*gitlab.Issue, error) {
	issue, _, err := gl.client.Issues.CreateIssue(p.ID, options)
	if err != nil {
		return nil, err
	}
	return issue, nil
}
