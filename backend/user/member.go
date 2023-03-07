package user

import (
	"context"
	"nova/storage"

	"github.com/patrickmn/go-cache"
	"github.com/xanzy/go-gitlab"
)

type Member struct {
	Id     string             `json:"id,omitempty"     bson:"_id"`
	Gitlab *gitlab.User       `json:"gitlab,omitempty"`
	LDAP   *storage.LdapEntry `json:"ldap,omitempty"`
}

func (m *Manager) RemoveMember(id string) {
	m.cache.Delete(id)
	m.mongo.DeleteObject("team", id)
}

func (m *Manager) ReadMember(id string) (*Member, error) {
	cached, found := m.cache.Get(id)
	if found {
		return cached.(*Member), nil
	}

	object := m.mongo.RetrieveObject("team", id)

	var member *Member
	err := object.Decode(member)
	if err != nil {
		return nil, err
	}

	m.cache.Set(id, member, cache.DefaultExpiration)
	return member, nil
}

func (m *Manager) ListMember() ([]*Member, error) {
	cached, found := m.cache.Get("team")
	if found {
		return cached.([]*Member), nil
	}

	cursor := m.mongo.ListObjects("team")

	var team []*Member
	err := cursor.All(context.TODO(), team)
	if err != nil {
		return nil, err
	}

	m.cache.Set("team", team, cache.DefaultExpiration)
	return team, nil
}

func (m *Manager) UpdateMember(s *Member) {
	m.cache.Delete(s.Id)
	m.cache.Set(s.Id, s, cache.DefaultExpiration)
	m.mongo.RefreshObject("team", s.Id, s)
}
