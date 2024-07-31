package entities

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)
	
var _ = fmt.Println // remove after test
	
func makeGrantTypeSet() map[string]struct{} {
	a := map[string]struct{}{}
	a["ResearchGrant"] = struct{}{}
	a["TeachingGrant"] = struct{}{}
	a["ServiceGrant"] = struct{}{}
	a["UnrestrictedGrant"] = struct{}{}
	return a
}

var GrantTypeSet = makeGrantTypeSet()


type Grant struct {
	Key *Key
	GrantType string
	Title string
	Administrator *Organization
	AwardingOrganization *Organization
	StartDate *Date
	EndDate *Date
	amount float64
	PIs []*Person
	CoIs []*Person
	Topics []*Concept
}

func (a *Grant) setAmount(n float64) {
	if n > 0 && n <= 1e+12 {
		a.amount = n
	} else {
		err := errors.New("Invalid grant amount: " + strconv.FormatFloat(n, 'f', -1, 64))
		fmt.Println(err)
	}
}

func (a *Grant) getAmount(n float64) float64 {
	return a.amount
}

func NewGrant(ty string, ti string) (*Grant, error) {
	p := new(Grant)
	p.Key = makeKey("Grant")
	_,ok := GrantTypeSet[ty]
	if !ok {
		err := errors.New("Unknown grant type: "+ty)
		return nil, err
	}
	p.GrantType = ty
	p.Title = ti
	return p, nil
}

func NewResearchGrant(n string) (*Grant,error) {return NewGrant("ResearchGrant", n)}
func NewTeachingGrant(n string) (*Grant,error) {return NewGrant("TeachingGrant", n)}
func NewServiceGrant(n string) (*Grant,error) {return NewGrant("ServiceGrant", n)}
func NewUnrestrictedGrant(n string) (*Grant,error) {return NewGrant("UnrestrictedGrant", n)}

func (a *Grant) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Grant",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasGrantType",(*a).GrantType,nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasTitle",(*a).Title,nil}))
	if (*a).Administrator != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasAdministrator","",(*a.Administrator).Key}))}
	if (*a).AwardingOrganization != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasAwardingOrganization","",(*a.AwardingOrganization).Key}))}
	if (*a).StartDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasStartDate","",(*a.StartDate).Key}))}
	if (*a).EndDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasEndDate","",(*a.EndDate).Key}))}
	if (*a).amount != 0.0 { t = append(t, makeTriple(Triple{(*a).Key,"hasAmount",strconv.FormatFloat((*a).amount, 'f', -1, 64),nil}))}
	for _,cptr := range (*a).PIs {
		t = append(t, makeTriple(Triple{(*a).Key,"hasPI","",cptr.Key}))
	}
	for _,cptr := range (*a).CoIs {
		t = append(t, makeTriple(Triple{(*a).Key,"hasCoI","",cptr.Key}))
	}
	for _,cptr := range (*a).Topics {
		t = append(t, makeTriple(Triple{(*a).Key,"hasTopic","",cptr.Key}))
	}
	return t
}

func (a *Grant) Row() []string {
	var t []string
	var pi []string
	var co []string
	var to []string
	t = append(t, a.Key.s)
	t = append(t, a.GrantType)
	t = append(t, a.Title)
	if a.Administrator != nil { t=append(t, a.Administrator.Key.s)}
	if a.AwardingOrganization != nil { t=append(t, a.AwardingOrganization.Key.s)}
	if a.StartDate != nil { t=append(t, a.StartDate.Key.s)}
	if a.EndDate != nil { t=append(t, a.EndDate.Key.s)}
	if a.amount != 0.0 { t=append(t, strconv.FormatFloat((*a).amount, 'f', -1, 64))}
	for _,cptr := range (*a).PIs {
		pi = append(pi, cptr.Key.s)
	}
	t = append(t,strings.Join(pi,","))
	for _,cptr := range (*a).CoIs {
		co = append(co, cptr.Key.s)
	}
	t = append(t,strings.Join(co,","))
	for _,cptr := range (*a).Topics {
		to = append(to, cptr.Key.s)
	}
	t = append(t,strings.Join(to,","))
	return t
}

func FindGrantKey (kf *Key) int {
	for i,a := range Grants {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddGrantFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindGrantKey(key)
	switch a[1] {
	case "GrantType":
		Grants[i].GrantType = a[2]
	case "Title":
		Grants[i].Title = a[2]
	case "Administrator":
		key.s = a[2]
		j := FindOrganizationKey(key)
		Grants[i].Administrator = Organizations[j]
	case "AwardingOrganization":
		key.s = a[2]
		j := FindOrganizationKey(key)
		Grants[i].AwardingOrganization = Organizations[j]
	case "StartDate":
		key.s = a[2]
		j := FindDateKey(key)
		Grants[i].StartDate = Dates[j]
	case "EndDate":
		key.s = a[2]
		j := FindDateKey(key)
		Grants[i].EndDate = Dates[j]
	case "Amount":
		f,_ := strconv.ParseFloat(a[2], 64)
		Grants[i].setAmount(f)
	case "PI":
		key.s = a[2]
		j := FindPersonKey(key)
		Grants[i].PIs = append(Grants[i].PIs,People[j])	
	case "CoI":
		key.s = a[2]
		j := FindPersonKey(key)
		Grants[i].CoIs = append(Grants[i].CoIs,People[j])	
	case "Topic":
		key.s = a[2]
		j := FindConceptKey(key)
		Grants[i].Topics = append(Grants[i].Topics,Concepts[j])	
	}
}


var Grants []*Grant