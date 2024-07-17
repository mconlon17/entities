package entities

// add regex check to email adress text

// ^([a-zA-Z0-9_\-\.]+)@([a-zA-Z0-9_\-\.]+)\.([a-zA-Z]{2,5})$ 

import (
	"fmt"
)
	
var _ = fmt.Println // remove after test
	
type EmailAddress struct {
	Key *Key
	EmailAddressText string
}

func NewEmailAddress(e string) *EmailAddress {
	p := new(EmailAddress)
	p.Key = makeKey("EmailAddress")
	p.EmailAddressText = e
	return p
}

func (a *EmailAddress) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","EmailAddress",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasEmailAddress",(*a).EmailAddressText,nil}))
	return t
}

func FindEmailAddressKey (kf *Key) int {
	for i,a := range EmailAddresses {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddEmailAddressFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindEmailAddressKey(key)
	switch a[1] {
	case "EmailAddress":
		EmailAddresses[i].EmailAddressText = a[2]
	}
}

var EmailAddresses []*EmailAddress
