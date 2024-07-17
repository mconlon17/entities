package entities

// TODO Add regex checking to phone numbers

// 	^[01]?[- .]?(\([2-9]\d{2}\)|[2-9]\d{2})[- .]?\d{3}[- .]?\d{4}$
import (
	"fmt"
)
	
var _ = fmt.Println // remove after test

type PhoneNumber struct {
	Key *Key
	PhoneNumberText string
}

func NewPhoneNumber(e string) *PhoneNumber {
	p := new(PhoneNumber)
	p.Key = makeKey("PhoneNumber")
	p.PhoneNumberText = e
	return p
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