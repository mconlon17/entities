package entities

// TODO build out publication struct

import (
	"errors"
	"fmt"
)
	
var _ = fmt.Println // remove after test
	
func makePublicationTypeSet() map[string]struct{} {
	a := map[string]struct{}{}
	a["Abstract"] = struct{}{}
	a["Article"] = struct{}{}
	a["Editorial"] = struct{}{}
	a["Book"] = struct{}{}
	a["Booklet"] = struct{}{}
    a["Conference"] = struct{}{}
    a["InBook"] = struct{}{}
    a["InCollection"] = struct{}{}
    a["Manual"] = struct{}{}
    a["MasterThesis"] = struct{}{}
    a["PhDThesis"] = struct{}{}
	a["Proceedings"] = struct{}{}
	a["TechReport"] = struct{}{}
	a["Unpublished"] = struct{}{}
	a["Dataset"] = struct{}{}
	a["Presentation"] = struct{}{}
	return a
}

var PublicationTypeSet = makePublicationTypeSet()


type Publication struct {
	Key *Key
	PublicationType string
	Title string
	Abstract string
	PublicationDate *Date
	Venue *Venue
	DOI *URL
	FullTextURL *URL
	Authors []*Person
	Keywords []*Concept
}

func NewPublication(ty string, ti string) *Publication {
	p := new(Publication)
	p.Key = makeKey("Publication")
	_,ok := PublicationTypeSet[ty]
	if !ok {
		err := errors.New("Unknown publication type: "+ty)
		fmt.Println(err)
	}
	p.PublicationType = ty
	p.Title = ti
	return p
}

func (a *Publication) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Publication",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasPublicationType",(*a).PublicationType,nil}))
	if (*a).Title != "" {t = append(t, makeTriple(Triple{(*a).Key,"hasTitle",(*a).Title,nil}))}
	if (*a).Abstract != "" {t = append(t, makeTriple(Triple{(*a).Key,"hasAbstract",(*a).Abstract,nil}))}

	if (*a).PublicationDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasPublicationDate","",(*a.PublicationDate).Key}))}
	if (*a).Venue != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasVenue","",(*a.Venue).Key}))}
    if (*a).DOI != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasDOI","",(*a.DOI).Key}))}
	if (*a).FullTextURL != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasFullTextURL","",(*a.FullTextURL).Key}))}
	
	for _,cptr := range (*a).Authors {
		t = append(t, makeTriple(Triple{(*a).Key,"hasAuthor","",cptr.Key}))
	}
	for _,cptr := range (*a).Keywords {
		t = append(t, makeTriple(Triple{(*a).Key,"hasKeyword","",cptr.Key}))
	}
	return t
}

func FindPublicationKey (kf *Key) int {
	for i,a := range Publications {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddPublicationFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindPublicationKey(key)
	switch a[1] {
	case "PublicationType":
		Publications[i].PublicationType = a[2]
	case "Title":
		Publications[i].Title = a[2]
	case "Abstract":
		Publications[i].Abstract = a[2]
	case "PublicationDate":
		key.s = a[2]
		j := FindDateKey(key)
		Publications[i].PublicationDate = Dates[j]
	case "Venue":
		key.s = a[2]
		j := FindVenueKey(key)
		Publications[i].Venue = Venues[j]
	case "DOI":
		key.s = a[2]
		j := FindURLKey(key)
		Publications[i].DOI = URLs[j]
	case "FullTextURL":
		key.s = a[2]
		j := FindURLKey(key)
		Publications[i].FullTextURL = URLs[j]
	case "Author":
		key.s = a[2]
		j := FindPersonKey(key)
		Publications[i].Authors = append(Publications[i].Authors,People[j])		
	case "Keyword":
		key.s = a[2]
		j := FindConceptKey(key)
		fmt.Println(j,key,a)
		Publications[i].Keywords = append(Publications[i].Keywords,Concepts[j])	
	}
}


var Publications []*Publication