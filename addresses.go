package entities

import (
	"errors"
	"fmt"
	"regexp"
)
	
var _ = fmt.Println // remove after test

var r1 = regexp.MustCompile(`^(?P<Zip>[0-9]{5,5})$`)

	
type Address struct {
	Key *Key
	AddressLine string
	Location *Location
	PostalCode string
}

func NewAddress(a string, l *Location, c string) (*Address,error) {
	p := new(Address)
	p.Key = makeKey("Address")
	p.AddressLine = a
	p.Location = l
	u := r1.FindStringSubmatch(c)
	if len(u) == 2 {
		p.PostalCode = c
		return p,nil
	} else {
		err := errors.New("Invalid postal code: " + c)
		return nil,err
	}
}

func (a *Address) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Address",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasAddressLine",(*a).AddressLine,nil}))
	if (*a).Location != nil { t=append(t, makeTriple(Triple{(*a).Key,"hasLocation","",(*a.Location).Key}))}
	if (*a).PostalCode != "" { t=append(t, makeTriple(Triple{(*a).Key,"hasPostalCode",(*a).PostalCode,nil}))}
	return t
}

func (a *Address) Row() []string {
	var t []string
	t = append(t, a.Key.s)
	t = append(t, a.AddressLine)
	if (*a).Location != nil { t=append(t, a.Location.Key.s)}
	if (*a).PostalCode != "" { t=append(t, a.PostalCode)}
	return t
}

func FindAddressKey (kf *Key) int {
	for i,a := range Addresses {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddAddressFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindAddressKey(key)
	switch a[1] {
	case "AddressLine":
		Addresses[i].AddressLine = a[2]
	case "Location":
		key.s = a[2]
		j := FindLocationKey(key)
		Addresses[i].Location = Locations[j]
	case "PostalCode":
		Addresses[i].PostalCode = a[2]
	}
}

var Addresses []*Address