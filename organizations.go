package entities

import (
	"errors"
	"fmt"
)
	
var _ = fmt.Println // remove after test
	
func makeOrganizationTypeSet() map[string]struct{} {
	a := map[string]struct{}{}
	a["Company"] = struct{}{}
	a["University"] = struct{}{}
	a["College"] = struct{}{}
	a["Department"] = struct{}{}
	a["Office"] = struct{}{}
	a["Nonprofit"] = struct{}{}
	a["Research"] = struct{}{}
	a["Education"] = struct{}{}
	a["Religious"] = struct{}{}
	a["Military"] = struct{}{}
	a["Government"] = struct{}{}
	a["Professional"] = struct{}{}
	a["Governance"] = struct{}{}
	a["Club"] = struct{}{}
    a["NGO"] = struct{}{}
	return a
}

var OrganizationTypeSet = makeOrganizationTypeSet()


type Organization struct {
	Key *Key
	OrganizationType string
	Name string
	ChildOrganization *Organization
	ParentOrganization *Organization
	Logo *Image
	Homepage *URL
	Resources []*Resource
}

func NewOrganization(t string, n string) *Organization {
	p := new(Organization)
	p.Key = makeKey("Organization")
	_,ok := OrganizationTypeSet[t]
	if !ok {
		err := errors.New("Unknown organization type: "+t)
		fmt.Println(err)
	}
	p.OrganizationType = t
	p.Name = n
	return p
}

func NewCompany(n string) *Organization {return NewOrganization("Company", n)}
func NewUniversity(n string) *Organization {return NewOrganization("University", n)}
func NewCollege(n string) *Organization {return NewOrganization("College", n)}
func NewDepartment(n string) *Organization {return NewOrganization("Department", n)}
func NewOffice(n string) *Organization {return NewOrganization("Office", n)}
func NewNonprofit(n string) *Organization {return NewOrganization("Nonprofit", n)}
func NewInstitute(n string) *Organization {return NewOrganization("Research", n)}
func NewSchool(n string) *Organization {return NewOrganization("Education", n)}
func NewChurch(n string) *Organization {return NewOrganization("Religious", n)}
func NewMilitary(n string) *Organization {return NewOrganization("Military", n)}
func NewGovernanment(n string) *Organization {return NewOrganization("Government", n)}
func NewAssociation(n string) *Organization {return NewOrganization("Professional", n)}
func NewGovernance(n string) *Organization {return NewOrganization("Governance", n)}
func NewClub(n string) *Organization {return NewOrganization("Club", n)}
func NewNGO(n string) *Organization {return NewOrganization("NGO", n)}

func (a *Organization) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Organization",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasOrganizationType",(*a).OrganizationType,nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasName",(*a).Name,nil}))
	if (*a).ChildOrganization != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasChildOrganization","",(*a.ChildOrganization).Key}))}
	if (*a).ParentOrganization != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasParentOrganization","",(*a.ParentOrganization).Key}))}
	if (*a).Logo != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasLogo","",(*a.Logo).Key}))}
	if (*a).Homepage != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasHomepage","",(*a.Homepage).Key}))}
	for _,cptr := range (*a).Resources {
		t = append(t, makeTriple(Triple{(*a).Key,"hasResource","",cptr.Key}))
	}	
	return t
}

func FindOrganizationKey (kf *Key) int {
	for i,a := range Organizations {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddOrganizationFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindOrganizationKey(key)
	switch a[1] {
	case "OrganizationType":
		Organizations[i].OrganizationType = a[2]
	case "Name":
		Organizations[i].Name = a[2]
	case "ChildOrganization":
		key.s = a[2]
		j := FindOrganizationKey(key)
		Organizations[i].ChildOrganization = Organizations[j]
	case "ParentOrganization":
		key.s = a[2]
		j := FindOrganizationKey(key)
		Organizations[i].ParentOrganization = Organizations[j]	
	case "Logo":
		key.s = a[2]
		j := FindImageKey(key)
		Organizations[i].Logo= Images[j]
	case "Homepage":
		key.s = a[2]
		j := FindURLKey(key)
		Organizations[i].Homepage = URLs[j]
	case "Resource":
		key.s = a[2]
		j := FindResourceKey(key)
		fmt.Println(j,key,a)
		Organizations[i].Resources = append(Organizations[i].Resources,Resources[j])	
	}
}


var Organizations []*Organization
