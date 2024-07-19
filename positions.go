package entities

import (
	"errors"
	"fmt"
)
	
var _ = fmt.Println // remove after test
	
func makePositionTypeSet() map[string]struct{} {
	a := map[string]struct{}{}
	a["Faculty"] = struct{}{}
	a["Staff"] = struct{}{}
	a["Director"] = struct{}{}
	a["Volunteer"] = struct{}{}
	a["Member"] = struct{}{}
	a["Retired"] = struct{}{}
	return a
}

var PositionTypeSet = makePositionTypeSet()


type Position struct {
	Key *Key
	PositionType string
	Title string
	Description string
	Organization *Organization
	Person *Person
	StartDate *Date
	EndDate *Date
}

func NewPosition(t, e string) (*Position,error) {
	p := new(Position)
	p.Key = makeKey("Position")
	_,ok := PositionTypeSet[t]
	if !ok {
		err := errors.New("Unknown position type: " + t)
		return nil,err
	}
	p.PositionType = t
	p.Title = e
	return p, nil
}

func NewFaculty(n string) (*Position,error) {return NewPosition("Faculty", n)}
func NewStaff(n string) (*Position,error) {return NewPosition("Staff", n)}
func NewDirector(n string) (*Position,error) {return NewPosition("Director", n)}
func NewVolunteer(n string) (*Position,error) {return NewPosition("Volunteer", n)}
func NewMember(n string) (*Position,error) {return NewPosition("Member", n)}
func NewRetired(n string) (*Position,error) {return NewPosition("Retired", n)}

func (a *Position) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Position",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasPositionType",(*a).PositionType,nil}))
	if (*a).Title != "" {t = append(t, makeTriple(Triple{(*a).Key,"hasTitle",(*a).Title,nil}))}
	if (*a).Description != "" {t = append(t, makeTriple(Triple{(*a).Key,"hasDescription",(*a).Description,nil}))}
	if (*a).Organization != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasOrganization","",(*a.Organization).Key}))}
	if (*a).Person != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasPerson","",(*a.Person).Key}))}
    if (*a).StartDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasStartDate","",(*a.StartDate).Key}))}
	if (*a).EndDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasEndDate","",(*a.EndDate).Key}))}
	return t
}

func FindPositionKey (kf *Key) int {
	for i,a := range Positions {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddPositionFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindPositionKey(key)
	switch a[1] {
	case "PositionType":
		Positions[i].PositionType = a[2]
	case "Title":
		Positions[i].Title = a[2]
	case "Description":
		Positions[i].Description = a[2]
	case "Organization":
		key.s = a[2]
		j := FindOrganizationKey(key)
		Positions[i].Organization = Organizations[j]
	case "Person":
		key.s = a[2]
		j := FindPersonKey(key)
		Positions[i].Person = People[j]
	case "StartDate":
		key.s = a[2]
		j := FindDateKey(key)
		Positions[i].StartDate = Dates[j]
	case "EndDate":
		key.s = a[2]
		j := FindDateKey(key)
		Positions[i].EndDate = Dates[j]			
	}
}

var Positions []*Position
