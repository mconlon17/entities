package entities

import(
	"errors"
	"fmt"
	"regexp"
)
	
var _ = fmt.Println // remove after test
var r2 = regexp.MustCompile(`^(?P<Name>[a-zA-Z0-9_\-\.]+)@(?P<Address>[a-zA-Z0-9_\-\.]+)\.(?P<Domain>[a-zA-Z]{2,5})$`)
	
type EmailAddress struct {
	Key *Key
	EmailAddressText string
}

func NewEmailAddress(e string) (*EmailAddress, error) {
	p := new(EmailAddress)
	p.Key = makeKey("EmailAddress")
	u := r2.FindStringSubmatch(e)
	if len(u) == 4 {
		p.EmailAddressText = e
		return p,nil
	} else {
		err := errors.New("Invalid email address: " + e)
		return nil,err
	}
}

func (a *EmailAddress) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","EmailAddress",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasEmailAddress",(*a).EmailAddressText,nil}))
	return t
}

func (a *EmailAddress) Row() []string {
	var t []string
	t = append(t, a.Key.s)
	t = append(t, a.EmailAddressText)
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
