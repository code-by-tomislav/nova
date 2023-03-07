package user

import (
	"nova/storage"

	"github.com/patrickmn/go-cache"
	"github.com/xanzy/go-gitlab"
)

type Student struct {
	Gitlab *gitlab.User
	LDAP   *storage.LdapEntry
}

func (m *Manager) ReadStudent(id string) (*Student, error) {
	cached, found := m.cache.Get(id)
	if found {
		return cached.(*Student), nil
	}

	ldap := m.ldap.RetrieveUser(id)
	gitlab, _ := m.gitlab.GetUser(ldap.LdapUID)

	s := &Student{
		LDAP:   ldap,
		Gitlab: gitlab,
	}

	m.cache.Set(id, s, cache.DefaultExpiration)
	return s, nil
}

func (m *Manager) ListStudents() ([]*Student, error) {
	cached, found := m.cache.Get("students")
	if found {
		return cached.([]*Student), nil
	}

	var s []*Student
	list := m.ldap.ListUsers(m.config.LDAP.Filter.Students)
	for _, entry := range list {
		s = append(s, &Student{
			LDAP: entry,
		})
	}

	m.cache.Set("students", s, cache.DefaultExpiration)
	return s, nil
}
