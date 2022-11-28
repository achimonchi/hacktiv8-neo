package helper

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

type UserLDAPData struct {
	ID       string
	Email    string
	Name     string
	FullName string
}

const (
	LDAP_Server   = "ldap.forumsys.com"
	LDAP_Port     = 389
	LDAP_BindDN   = "cn=read-only-admin,dc=example,dc=com"
	LDAP_Password = "password"
	LDAP_SearchDN = "dc=example,dc=com"
)

func AuthUsingLDAP(username, password string) (bool, *UserLDAPData, error) {
	listen, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", LDAP_Server, LDAP_Port))
	if err != nil {
		return false, nil, err
	}

	defer listen.Close()

	err = listen.Bind(LDAP_BindDN, LDAP_Password)
	if err != nil {
		return false, nil, err
	}

	search := ldap.NewSearchRequest(
		LDAP_SearchDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(mail=%s))", username),
		[]string{"dn", "cn", "sn", "mail", "uid"},
		nil,
	)

	sr, err := listen.Search(search)
	if err != nil {
		return false, nil, err
	}

	if len(sr.Entries) == 0 {
		return false, nil, nil
	}

	entry := sr.Entries[0]

	err = listen.Bind(entry.DN, password)
	if err != nil {
		return false, nil, nil
	}

	fmt.Printf("%+v\n", entry)

	data := new(UserLDAPData)

	data.ID = username
	for _, attr := range entry.Attributes {
		switch attr.Name {
		case "sn":
			data.Name = attr.Values[0]
		case "uid":
			data.ID = attr.Values[0]
		case "mail":
			data.Email = attr.Values[0]
		case "cn":
			data.FullName = attr.Values[0]
		}
	}
	return true, data, nil
}
