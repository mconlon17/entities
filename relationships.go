package entities

import (
	"errors"
	"fmt"
)
	
var _ = fmt.Println // remove after test

func makeRelationshipTypeSet() map[string]struct{} {
	a := map[string]struct{}{}
	a["Friend"] = struct{}{}
	a["Mentor"] = struct{}{}
	a["Parent"] = struct{}{}
	a["Spouse"] = struct{}{}
	a["Supervisor"] = struct{}{}
	return a
}

var RelationshipTypeSet = makeRelationshipTypeSet()

type Relationship struct {
	Key *Key
	RelationshipType string
	PersonA *Person
	PersonB *Person
	StartDate *Date
	EndDate *Date
}

func NewRelationship(t string, pa *Person, pb *Person) (*Relationship,error) {
	p := new(Relationship)
	p.Key = makeKey("Relationship")
	_,ok := RelationshipTypeSet[t]
	if !ok {
		err := errors.New("Unknown relationship type: "+t)
		return nil,err
	}
	p.RelationshipType = t
	p.PersonA = pa
	p.PersonB = pb
	return p, nil
}

func NewFriend(pa *Person, pb *Person)     (*Relationship,error) {return NewRelationship("Friend", pa, pb)}
func NewMentor(pa *Person, pb *Person)     (*Relationship,error) {return NewRelationship("Mentor", pa, pb)}
func NewParent(pa *Person, pb *Person)     (*Relationship,error) {return NewRelationship("Parent", pa, pb)}
func NewSpouse(pa *Person, pb *Person)     (*Relationship,error) {return NewRelationship("Spouse", pa, pb)}
func NewSupervisor(pa *Person, pb *Person) (*Relationship,error) {return NewRelationship("Supervisor", pa, pb)}

func (a *Relationship) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Relationship",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasRelationshipType",(*a).RelationshipType,nil}))
	if (*a).PersonA != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasPersonA","",(*a.PersonA).Key}))}
	if (*a).PersonB != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasPersonB","",(*a.PersonB).Key}))}
    if (*a).StartDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasStartDate","",(*a.StartDate).Key}))}
	if (*a).EndDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasEndDate","",(*a.EndDate).Key}))}
	return t
}

func FindRelationshipKey (kf *Key) int {
	for i,a := range Relationships {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddRelationshipFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindRelationshipKey(key)
	switch a[1] {
	case "RelationshipType":
		Relationships[i].RelationshipType = a[2]
	case "PersonA":
		key.s = a[2]
		j := FindPersonKey(key)
		Relationships[i].PersonA = People[j]
	case "PersonB":
		key.s = a[2]
		j := FindPersonKey(key)
		Relationships[i].PersonB = People[j]		
	case "StartDate":
		key.s = a[2]
		j := FindDateKey(key)
		Relationships[i].StartDate = Dates[j]
	case "EndDate":
		key.s = a[2]
		j := FindDateKey(key)
		Relationships[i].EndDate = Dates[j]
	}
}

var Relationships []*Relationship
