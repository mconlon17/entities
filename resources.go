package entities

import (
	"errors"
	"fmt"
)
	
var _ = fmt.Println // remove after test

func makeResourceTypeSet() map[string]struct{} {
	a := map[string]struct{}{}
	a["Equipment"] = struct{}{}
	a["Service"] = struct{}{}
	return a
}

var ResourceTypeSet = makeResourceTypeSet()

type Resource struct {
	Key *Key
	Name string
	ResourceType string
	Description string
	Contact *Person
	Admin *Organization
}

func NewResource(t string, n string) *Resource {
	p := new(Resource)
	p.Key = makeKey("Resource")
	_,ok := ResourceTypeSet[t]
	if !ok {
		err := errors.New("Unknown resource type: "+t)
		fmt.Println(err)
	}
	p.ResourceType = t
	p.Name = n
	return p
}

func NewEquipment(n string) *Resource {return NewResource("Equipment", n)}
func NewService(n string) *Resource {return NewResource("Service", n)}

func (a *Resource) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Resource",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasResourceType",(*a).ResourceType,nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasName",(*a).Name,nil}))
	if (*a).Description != "" {t = append(t, makeTriple(Triple{(*a).Key,"hasDescription",(*a).Description,nil}))}
	if (*a).Contact != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasContact","",(*a.Contact).Key}))}
	if (*a).Admin != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasAdmin","",(*a.Admin).Key}))}
	return t
}

func FindResourceKey (kf *Key) int {
	for i,a := range Resources {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddResourceFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindResourceKey(key)
	switch a[1] {
	case "ResourceType":
		Resources[i].ResourceType = a[2]
	case "Name":
		Resources[i].Name = a[2]
	case "Description":
		Resources[i].Description = a[2]
	case "Contact":
		key.s = a[2]
		j := FindPersonKey(key)
		Resources[i].Contact = People[j]		
	case "Admin":
		key.s = a[2]
		j := FindOrganizationKey(key)
		Resources[i].Admin = Organizations[j]
	}
}


var Resources []*Resource