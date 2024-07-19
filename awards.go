package entities

import (
	"errors"
	"fmt"
)
	
var _ = fmt.Println // remove after test

func makeAwardTypeSet() map[string]struct{} {
	a := map[string]struct{}{}
	a["PhD"] = struct{}{}
	a["Masters"] = struct{}{}
	a["Bachelors"] = struct{}{}
	a["Associate"] = struct{}{}
	a["HonorSociety"] = struct{}{}
	a["TeachingAward"] = struct{}{}
	a["ServiceAward"] = struct{}{}
	a["ResearchAward"] = struct{}{}
	a["SportsAward"] = struct{}{}
	return a
}

var AwardTypeSet = makeAwardTypeSet()

type Award struct {
	Key *Key
	AwardType string
	Title string
	AwardingOrganization *Organization
	Awardee *Person
	AwardDate *Date
}

func NewAward(t string, d string) (*Award,error) {
	p := new(Award)
	p.Key = makeKey("Award")
	_,ok := AwardTypeSet[t]
	if !ok {
		err := errors.New("Unknown award type: "+t)
		return nil,err
	}
	p.AwardType = t
	p.Title = d
	return p, nil
}

func (a *Award) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Award",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasAwardType",(*a).AwardType,nil}))
	if (*a).AwardingOrganization != nil { t=append(t, makeTriple(Triple{(*a).Key,"hasAwardingOrganization","",(*a.AwardingOrganization).Key}))}
	if (*a).Awardee != nil { t=append(t, makeTriple(Triple{(*a).Key,"hasAwardee","",(*a.Awardee).Key}))}
	if (*a).AwardDate != nil { t=append(t, makeTriple(Triple{(*a).Key,"hasAwardDate","",(*a.AwardDate).Key}))}
	return t
}

func FindAwardKey (kf *Key) int {
	for i,a := range Awards {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddAwardFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindAwardKey(key)
	switch a[1] {
	case "AwardType":
		Awards[i].AwardType = a[2]
	case "Title":
		Awards[i].Title = a[2]
	case "AwardingOrganization":
		key.s = a[2]
		j := FindOrganizationKey(key)
		Awards[i].AwardingOrganization = Organizations[j]
	case "Awardee":
		key.s = a[2]
		j := FindPersonKey(key)
		Awards[i].Awardee = People[j]
	case "AwardDate":
		key.s = a[2]
		j := FindDateKey(key)
		Awards[i].AwardDate = Dates[j]
	}
}

var Awards []*Award
