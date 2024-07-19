package entities

import (
	"errors"
	"fmt"
	"regexp"
)
	
var _ = fmt.Println // remove after test
var r6 = regexp.MustCompile(`^(?P<Num>[2-9]\d{2}\.[2-9]\d{2}\.\d{4})$`)

type PhoneNumber struct {
	Key *Key
	PhoneNumberText string
}

func NewPhoneNumber(e string) (*PhoneNumber,error) {
	p := new(PhoneNumber)
	p.Key = makeKey("PhoneNumber")
	u := r6.FindStringSubmatch(e)
	if len(u) == 2 {
		p.PhoneNumberText = e
		return p,nil
	} else {
		err := errors.New("Invalid phone number: " + e)
		return nil,err
	}
}

func (a *PhoneNumber) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","PhoneNumber",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasPhoneNumber",(*a).PhoneNumberText,nil}))
	return t
}

func FindPhoneNumberKey (kf *Key) int {
	for i,a := range PhoneNumbers {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddPhoneNumberFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindPhoneNumberKey(key)
	switch a[1] {
	case "PhoneNumber":
		PhoneNumbers[i].PhoneNumberText = a[2]
	}
}

var PhoneNumbers []*PhoneNumber