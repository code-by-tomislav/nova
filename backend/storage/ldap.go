package storage

import (
	"fmt"
	"log"
	"nova/utils"

	"github.com/go-ldap/ldap"
)

type LDAP struct {
	config     *utils.Configuration
	connection *ldap.Conn
}

type LdapEntry struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Mail      string `json:"mail,omitempty"`
	LdapUID   string `json:"ldap_uid,omitempty"`
}

func NewLDAP(c *utils.Configuration) *LDAP {
	l, err := ldap.DialURL(c.LDAP.Host)
	if err != nil {
		log.Fatalln("failed to connect to ldap")
	}

	ldap := &LDAP{
		config:     c,
		connection: l,
	}
	return ldap
}

// public functions

func (l *LDAP) RetrieveUser(id string) *LdapEntry {
	l.bind()
	defer l.connection.Close()

	filter := fmt.Sprintf("(%s=%s)", l.config.LDAP.Field.Id, id)

	searchReq := ldap.NewSearchRequest(
		l.config.LDAP.Domain.Base,
		ldap.ScopeWholeSubtree,
		0,
		0,
		0,
		false,
		filter,
		[]string{
			l.config.LDAP.Field.Id,
			l.config.LDAP.Field.FirstName,
			l.config.LDAP.Field.LastName,
			l.config.LDAP.Field.Mail,
		},
		[]ldap.Control{},
	)

	result, err := l.connection.Search(searchReq)
	if err != nil {
		log.Fatal(err)
	}

	return l.convertEntry(result.Entries[0])
}

func (l *LDAP) ListUsers(f string) []*LdapEntry {
	l.bind()
	defer l.connection.Close()

	searchReq := ldap.NewSearchRequest(
		l.config.LDAP.Domain.Base,
		ldap.ScopeWholeSubtree,
		0,
		0,
		0,
		false,
		f,
		[]string{
			l.config.LDAP.Field.Id,
			l.config.LDAP.Field.FirstName,
			l.config.LDAP.Field.LastName,
			l.config.LDAP.Field.Mail,
		},
		[]ldap.Control{},
	)

	result, err := l.connection.Search(searchReq)
	if err != nil {
		log.Fatal(err)
	}

	return l.convertEntries(result.Entries)
}

// private functions

func (l *LDAP) bind() {
	err := l.connection.Bind(l.config.LDAP.Domain.Query, l.config.LDAP.Pass)
	if err != nil {
		log.Fatalln("failed to bind to ldap")
	}
}

func (l *LDAP) convertEntry(e *ldap.Entry) *LdapEntry {
	return &LdapEntry{
		ID:        e.GetAttributeValue(l.config.LDAP.Field.Id),
		FirstName: e.GetAttributeValue(l.config.LDAP.Field.FirstName),
		LastName:  e.GetAttributeValue(l.config.LDAP.Field.LastName),
		Mail:      e.GetAttributeValue(l.config.LDAP.Field.Mail),
		LdapUID:   e.GetAttributeValue(l.config.LDAP.Field.LdapUID),
	}
}

func (l *LDAP) convertEntries(e []*ldap.Entry) []*LdapEntry {
	var entries []*LdapEntry
	for _, entry := range e {
		entries = append(entries, l.convertEntry(entry))
	}
	return entries
}
