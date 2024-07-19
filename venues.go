package entities

import (
	"errors"
	"fmt"
)
	
var _ = fmt.Println // remove after test

func makeVenueTypeSet() map[string]struct{} {
	a := map[string]struct{}{}
	a["Journal"] = struct{}{}
	a["Magazine"] = struct{}{}
	a["Newsletter"] = struct{}{}
	return a
}

var VenueTypeSet = makeVenueTypeSet()

type Venue struct {
	Key *Key
	VenueType string
	Name string
	Publisher *Organization
}

func NewVenue(t string, n string) *Venue {
	p := new(Venue)
	p.Key = makeKey("Venue")
	_,ok := VenueTypeSet[t]
	if !ok {
		err := errors.New("Unknown venue type: "+t)
		fmt.Println(err)
	}
	p.VenueType = t
	p.Name = n
	return p
}

func NewJournal(n string) *Venue {return NewVenue("Journal", n)}
func NewMagazine(n string) *Venue {return NewVenue("Magazine", n)}
func NewNewsletter(n string) *Venue {return NewVenue("Newsletter", n)}

func (a *Venue) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Venue",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasVenueType",(*a).VenueType,nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasName",(*a).Name,nil}))
	if (*a).Publisher != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasPublisher","",(*a.Publisher).Key}))}
	return t
}

func FindVenueKey (kf *Key) int {
	for i,a := range Venues {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddVenueFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindVenueKey(key)
	switch a[1] {
	case "VenueType":
		Venues[i].VenueType = a[2]
	case "Name":
		Venues[i].Name = a[2]
	case "Publisher":
		key.s = a[2]
		j := FindOrganizationKey(key)
		Venues[i].Publisher = Organizations[j]
	}
}

var Venues []*Venue