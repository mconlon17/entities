package entities

import (
	"errors"
	"fmt"
	"regexp"
)
	
var _ = fmt.Println // remove after test
var r5 = regexp.MustCompile(`^(?P<Name>[a-zA-Z''-'\s]{1,40})$`)
	
type Person struct {
	Key *Key
	Name string
	Orcid *URL
	HomeAddress *Address
	Portfolio *Portfolio
	HomeEmail *EmailAddress
	WorkEmail *EmailAddress
	HomePhone *PhoneNumber
	ResearchAreas []*Concept
	BirthDate *Date
}

func NewPerson(n string) (*Person,error) {
	p := new(Person)
	p.Key = makeKey("Person")
	u := r5.FindStringSubmatch(n)
	if len(u) == 2 {
		p.Name = n
		return p,nil
	} else {
		err := errors.New("Invalid name: " + n)
		return nil,err
	}
}

func (a *Person) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Person",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasName",(*a).Name,nil}))
	if (*a).Orcid != nil { t = append(t, makeTriple(Triple{(*a).Key,"hasOrcid","",(*a.Orcid).Key}))}
	if (*a).HomeAddress != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasHomeAddress","",(*a.HomeAddress).Key}))}
    if (*a).Portfolio != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasPortfolio","",(*a.Portfolio).Key}))}
    if (*a).HomeEmail != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasHomeEmail","",(*a.HomeEmail).Key}))}
	if (*a).WorkEmail != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasWorkEmail","",(*a.WorkEmail).Key}))}
    if (*a).HomePhone != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasHomePhone","",(*a.HomePhone).Key}))}
	if (*a).BirthDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasBirthDate","",(*a.BirthDate).Key}))}
	for _,cptr := range (*a).ResearchAreas {
		t = append(t, makeTriple(Triple{(*a).Key,"hasResearchArea","",cptr.Key}))
	}
	return t
}

func FindPersonKey (kf *Key) int {
	for i,a := range People {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddPersonFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindPersonKey(key)
	switch a[1] {
	case "Name":
		People[i].Name = a[2]
	case "Orcid":
		key.s = a[2]
		j := FindURLKey(key)
		People[i].Orcid = URLs[j]
	case "HomeAddress":
		key.s = a[2]
		j := FindAddressKey(key)
		People[i].HomeAddress = Addresses[j]
	case "Portfolio":
		key.s = a[2]
		j := FindPortfolioKey(key)
		People[i].Portfolio = Portfolios[j]
	case "HomeEmail":
		key.s = a[2]
		j := FindEmailAddressKey(key)
		People[i].HomeEmail = EmailAddresses[j]
	case "WorkEmail":
		key.s = a[2]
		j := FindEmailAddressKey(key)
		People[i].WorkEmail = EmailAddresses[j]
	case "HomePhone":
		key.s = a[2]
		j := FindPhoneNumberKey(key)
		People[i].HomePhone = PhoneNumbers[j]
	case "BirthDate":
		key.s = a[2]
		j := FindDateKey(key)
		People[i].BirthDate = Dates[j]	
	case "ReseachArea":
		key.s = a[2]
		j := FindConceptKey(key)
		People[i].ResearchAreas = append(People[i].ResearchAreas,Concepts[j])	
	}
}

var People []*Person