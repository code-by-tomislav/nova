package web

import (
	"nova/items"
	"nova/user"
	"nova/utils"
)

type Bouncer struct {
	m *user.Manager
	b *items.Builder
}

func NewBouncer(c *utils.Configuration) *Bouncer {
	b := &Bouncer{
		m: user.NewManager(c),
		b: items.NewBuilder(c),
	}

	return b
}
