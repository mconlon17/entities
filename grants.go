package entities

import (
	"errors"
	"fmt"
	"strconv"
)
	
var _ = fmt.Println // remove after test
	
func makeGrantTypeSet() map[string]struct{} {
	a := map[string]struct{}{}
	a["Research Grant"] = struct{}{}
	a["Teaching Grant"] = struct{}{}
	a["Service Grant"] = struct{}{}
	a["Unrestricted Grant"] = struct{}{}
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

func NewGrant(ty string, ti string) *Grant {
	p := new(Grant)
	p.Key = makeKey("Grant")
	_,ok := GrantTypeSet[ty]
	if !ok {
		err := errors.New("Unknown grant type: "+ty)
		fmt.Println(err)
	}
	p.GrantType = ty
	p.Title = ti
	return p
}

func NewResearchGrant(n string) *Grant {return NewGrant("Research Grant", n)}
func NewTeachingGrant(n string) *Grant {return NewGrant("Teaching Grant", n)}
func NewServiceGrant(n string) *Grant {return NewGrant("Service Grant", n)}
func NewUnrestrictedGrant(n string) *Grant {return NewGrant("Unrestricted Grant", n)}

func (a *Grant) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Grant",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasGrantType",(*a).GrantType,nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasTitle",(*a).Title,nil}))
	if (*a).Administrator != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasAdministrator","",(*a.Administrator).Key}))}
	if (*a).AwardingOrganization != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasAwardingOrganization","",(*a.AwardingOrganization).Key}))}
	if (*a).StartDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasStartDate","",(*a.StartDate).Key}))}
	if (*a).EndDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasEndDate","",(*a.EndDate).Key}))}
	if (*a).amount != 0.0 { t = append(t, makeTriple(Triple{(*a).Key,"hasamount",strconv.FormatFloat((*a).amount, 'f', -1, 64),nil}))}
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
	case "amount":
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