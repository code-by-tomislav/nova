package items

import (
	"nova/storage"
	"nova/utils"
)

type Builder struct {
	gitlab *storage.Gitlab
	mongo  *storage.Mongo
}

func NewBuilder(c *utils.Configuration) *Builder {
	m := &Builder{
		gitlab: storage.NewGitlab(c),
		mongo:  storage.NewMongo(c),
	}
	return m
}
