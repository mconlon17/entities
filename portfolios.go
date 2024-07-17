package entities

import (
	"fmt"
)
	
var _ = fmt.Println // remove after test
	
type Portfolio struct {
	Key *Key
	Person *Person
	Positions []*Position
	Awards []*Award
	Publications []*Publication
	Teachings []*Teaching
	Relationships []*Relationship
	Grants []*Grant
}

func NewPortfolio() *Portfolio {
	p := new(Portfolio)
	p.Key = makeKey("Portfolio")
	return p
}

func (a *Portfolio) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Portfolio",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasPerson","",(*a.Person).Key}))
	for _,cptr := range (*a).Positions {
		t = append(t, makeTriple(Triple{(*a).Key,"hasPosition","",cptr.Key}))
	}
	for _,cptr := range (*a).Awards {
		t = append(t, makeTriple(Triple{(*a).Key,"hasAward","",cptr.Key}))
	}
	for _,cptr := range (*a).Publications {
		t = append(t, makeTriple(Triple{(*a).Key,"hasPublication","",cptr.Key}))
	}
	for _,cptr := range (*a).Teachings {
		t = append(t, makeTriple(Triple{(*a).Key,"hasTeaching","",cptr.Key}))
	}
	for _,cptr := range (*a).Relationships {
		t = append(t, makeTriple(Triple{(*a).Key,"hasRelationship","",cptr.Key}))
	}
	for _,cptr := range (*a).Grants {
		t = append(t, makeTriple(Triple{(*a).Key,"hasGrant","",cptr.Key}))
	}
	return t
}

func FindPortfolioKey (kf *Key) int {
	for i,a := range Portfolios {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddPortfolioFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindPortfolioKey(key)
	switch a[1] {
	case "Person":
		key.s = a[2]
		j := FindPersonKey(key)
		Portfolios[i].Person = People[j]
	case "Position":
		key.s = a[2]
		j := FindPositionKey(key)
		Portfolios[i].Positions = append(Portfolios[i].Positions,Positions[j])	
	case "Award":
		key.s = a[2]
		j := FindAwardKey(key)
		Portfolios[i].Awards = append(Portfolios[i].Awards,Awards[j])
	case "Publication":
		key.s = a[2]
		j := FindPublicationKey(key)
		Portfolios[i].Publications = append(Portfolios[i].Publications,Publications[j])
	case "Teaching":
		key.s = a[2]
		j := FindTeachingKey(key)
		Portfolios[i].Teachings = append(Portfolios[i].Teachings,Teachings[j])
	case "Relationship":
		key.s = a[2]
		j := FindRelationshipKey(key)
		Portfolios[i].Relationships = append(Portfolios[i].Relationships,Relationships[j])					
	case "Grant":
		key.s = a[2]
		j := FindGrantKey(key)
		Portfolios[i].Grants = append(Portfolios[i].Grants,Grants[j])	
	}
}

var Portfolios []*Portfolio