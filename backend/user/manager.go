package user

import (
	"nova/storage"
	"nova/utils"
	"time"

	"github.com/patrickmn/go-cache"
)

type Manager struct {
	cache  *cache.Cache
	config *utils.Configuration
	gitlab *storage.Gitlab
	ldap   *storage.LDAP
	mongo  *storage.Mongo
}

func NewManager(c *utils.Configuration) *Manager {
	m := &Manager{
		cache:  cache.New(24*time.Hour, 48*time.Hour),
		config: c,
		gitlab: storage.NewGitlab(c),
		mongo:  storage.NewMongo(c),
		ldap:   storage.NewLDAP(c),
	}
	return m
}
