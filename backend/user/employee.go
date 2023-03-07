package user

import (
	"nova/storage"

	"github.com/patrickmn/go-cache"
	"github.com/xanzy/go-gitlab"
)

type Employee struct {
	Gitlab *gitlab.User
	LDAP   *storage.LdapEntry
}

func (m *Manager) AddMember(e *Employee) *Member {
	member := &Member{
		Id:   e.LDAP.ID,
		LDAP: e.LDAP,
	}

	m.cache.Set(member.Id, e, cache.DefaultExpiration)
	return member
}

func (m *Manager) ReadEmployee(id string) (*Employee, error) {
	cached, found := m.cache.Get(id)
	if found {
		return cached.(*Employee), nil
	}

	e := &Employee{
		LDAP: m.ldap.RetrieveUser(id),
	}

	m.cache.Set(id, e, cache.DefaultExpiration)
	return e, nil
}

func (m *Manager) ListEmployees() ([]*Employee, error) {
	cached, found := m.cache.Get("employees")
	if found {
		return cached.([]*Employee), nil
	}

	var e []*Employee
	list := m.ldap.ListUsers(m.config.LDAP.Filter.Employees)
	for _, entry := range list {
		e = append(e, &Employee{
			LDAP: entry,
		})
	}

	m.cache.Set("employees", e, cache.DefaultExpiration)
	return e, nil
}
