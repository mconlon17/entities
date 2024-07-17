package entities

import (
	"fmt"
)
	
var _ = fmt.Println // remove after test

type Concept struct {
	Key *Key
	Name string
	Broader *Concept
	Narrower *Concept
}

func NewConcept(d string) *Concept {
	p := new(Concept)
	p.Key = makeKey("Concept")
	p.Name = d
	return p
}

func (a *Concept) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Concept",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasName",(*a).Name,nil}))
	if (*a).Broader != nil { t=append(t, makeTriple(Triple{(*a).Key,"hasBroader","",(*a.Broader).Key}))}
	if (*a).Narrower != nil { t=append(t, makeTriple(Triple{(*a).Key,"hasNarrower","",(*a.Narrower).Key}))}
	return t
}

func FindConceptKey (kf *Key) int {
	for i,a := range Concepts {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddConceptFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindConceptKey(key)
	switch a[1] {
	case "Name":
		Concepts[i].Name = a[2]
	case "Broader":
		key.s = a[2]
		j := FindConceptKey(key)
		Concepts[i].Broader = Concepts[j]
	case "Narrower":
		key.s = a[2]
		j := FindConceptKey(key)
		Concepts[i].Narrower = Concepts[j]
	}
}

var Concepts []*Concept