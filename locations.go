package entities

import (
	"errors"
	"fmt"
)
	
var _ = fmt.Println // remove after test
	
func makeLocationTypeSet() map[string]struct{} {
	a := map[string]struct{}{}
	a["Building"] = struct{}{}
	a["City"] = struct{}{}
	a["County"] = struct{}{}
	a["State"] = struct{}{}
	a["Country"] = struct{}{}

	return a
}

var LocationTypeSet = makeLocationTypeSet()


type Location struct {
	Key *Key
	LocationType string
	Name string
	Abbreviation string
	GeoLocation *GeoLocation
	Contains []*Location
	Inside []*Location	
}

func NewLocation(t string, n string) (*Location,error) {
	p := new(Location)
	p.Key = makeKey("Location")
	_,ok := LocationTypeSet[t]
	if !ok {
		err := errors.New("Unknown location type: " + t)
		return nil, err
	}
	p.LocationType = t
	p.Name = n
	return p, nil
}

func NewBuilding(n string) (*Location,error) {return NewLocation("Building", n)}
func NewCity(n string)  (*Location,error) {return NewLocation("City", n)}
func NewCounty(n string)  (*Location,error) {return NewLocation("County", n)}
func NewState(n string)  (*Location,error) {return NewLocation("State", n)}
func NewCountry(n string)  (*Location,error) {return NewLocation("Country", n)}

func (a *Location) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Location",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasLocationType",(*a).LocationType,nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasName",(*a).Name,nil}))
	if (*a).Abbreviation != "" {t = append(t, makeTriple(Triple{(*a).Key,"hasAbbreviation",(*a).Abbreviation,nil}))}
	if (*a).GeoLocation != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasGeoLocation","",(*a.GeoLocation).Key}))}	
		for _,cptr := range (*a).Contains {
		t = append(t, makeTriple(Triple{(*a).Key,"hasContains","",cptr.Key}))
	}
	for _,cptr := range (*a).Inside {
		t = append(t, makeTriple(Triple{(*a).Key,"hasInside","",cptr.Key}))
	}
	return t
}

func FindLocationKey (kf *Key) int {
	for i,a := range Locations {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddLocationFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindLocationKey(key)
	switch a[1] {
	case "LocationType":
		Locations[i].LocationType = a[2]
	case "Name":
		Locations[i].Name = a[2]
	case "Abbreviation":
		Locations[i].Abbreviation = a[2]
	case "GeoLocation":
		key.s = a[2]
		j := FindGeoLocationKey(key)
		Locations[i].GeoLocation = GeoLocations[j]
	case "Contains":
		key.s = a[2]
		j := FindLocationKey(key)
		Locations[i].Contains = append(Locations[i].Contains,Locations[j])
	case "Inside":
		key.s = a[2]
		j := FindPersonKey(key)
		Locations[i].Inside = append(Locations[i].Inside,Locations[j])	
	}
}

var Locations []*Location