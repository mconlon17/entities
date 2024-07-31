package entities

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

var _ = fmt.Println // remove after test

type Key struct {
	s string
}

func makeKeySet() map[string]string {
	return make(map[string]string)
}

var KeySet = makeKeySet()

type Triple struct {
	Subject *Key
	Predicate string
	Literal string
	Referent *Key
}

func makeKey(t string) *Key {
	var k *Key
	radix := string([]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"))
	
	l := len(radix)
	s := string(radix[rand.Intn(l)]) + string(radix[rand.Intn(l)]) + string(radix[rand.Intn(l)]) + 
			string(radix[rand.Intn(l)]) + string(radix[rand.Intn(l)]) + string(radix[rand.Intn(l)])
	_,ok := KeySet[s] 
	if ok {
		k = makeKey(t) // key value already in KeySet, so try again
	} else {
		KeySet[s] = t // s is a new key value, so record it and return it
		k = new(Key)
		k.s = s
	}
	return k
}

func makeTriple (a Triple) [3]string {
	var o string
	uri := "https://a.b/"
	t := [3]string{"<"+uri+a.Subject.s+">", "<" + uri + a.Predicate + ">",""}
	if a.Literal != "" {
		o = "\"" + a.Literal + "\""
	} else {
		o = "<" + uri + a.Referent.s + ">"
	}
	t[2] = o
	return t
}

// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}

// New returns an error that formats as the given text
func New(text string) error {
    return &errorString{text}
}

func ShowSets() {
 	fmt.Println(CourseTypeSet)
	fmt.Println(LocationTypeSet)
	fmt.Println(PublicationTypeSet)
	fmt.Println(PositionTypeSet)
	fmt.Println(AwardTypeSet)
	fmt.Println(DateTypeSet)		
	fmt.Println(OrganizationTypeSet)
	fmt.Println(GrantTypeSet)
	fmt.Println(RelationshipTypeSet)
	fmt.Println(VenueTypeSet)
}

func ShowEntities() {
 	for i,a := range Addresses      { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Awards         { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Concepts       { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Courses        { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Dates          { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range EmailAddresses { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Events         { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range GeoLocations   { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Grants         { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Images         { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Locations      { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Organizations  { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range People         { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range PhoneNumbers   { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Portfolios     { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Positions      { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Publications   { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Relationships  { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Resources      { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Teachings      { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range URLs           { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
 	for i,a := range Venues         { fmt.Println(i,KeySet[a.Key.s],a.Key.s, a) }
}

func EmptyLists() {
    Addresses      =[]*Address{}
    Awards         =[]*Award{}
    Concepts       =[]*Concept{}
    Courses        =[]*Course{}
    Dates          =[]*Date{}
    EmailAddresses =[]*EmailAddress{}
    Events         =[]*Event{}
    GeoLocations   =[]*GeoLocation{}
    Grants         =[]*Grant{}
    Images         =[]*Image{}
    Locations      =[]*Location{}
    Organizations  =[]*Organization{}
    People         =[]*Person{}
    PhoneNumbers   =[]*PhoneNumber{}
    Portfolios     =[]*Portfolio{}
    Positions      =[]*Position{}
    Publications   =[]*Publication{}
    Relationships  =[]*Relationship{}
    Resources      =[]*Resource{}
    Teachings      =[]*Teaching{}
    URLs           =[]*URL{}
    Venues         =[]*Venue{}
}

func createKey(s string, t string) *Key {
	k := new(Key)
	k.s = s
	KeySet[s] = t
	return k
}

func CreateEntity(a []string){
	switch a[2] {
    case "Address":
        Addresses=append(Addresses,&Address{Key:createKey(a[0],"Address")})
    case "Award":
        Awards=append(Awards,&Award{Key:createKey(a[0],"Award")})
    case "Concept":
        Concepts=append(Concepts,&Concept{Key:createKey(a[0],"Concept")})
    case "Course":
        Courses=append(Courses,&Course{Key:createKey(a[0],"Course")})
    case "Date":
        Dates=append(Dates,&Date{Key:createKey(a[0],"Date")})
    case "EmailAddress":
        EmailAddresses=append(EmailAddresses,&EmailAddress{Key:createKey(a[0],"EmailAddress")})
    case "Event":
        Events=append(Events,&Event{Key:createKey(a[0],"Event")})
    case "GeoLocation":
        GeoLocations=append(GeoLocations,&GeoLocation{Key:createKey(a[0],"GeoLocation")})
    case "Grant":
        Grants=append(Grants,&Grant{Key:createKey(a[0],"Grant")})
    case "Image":
        Images=append(Images,&Image{Key:createKey(a[0],"Image")})
    case "Location":
        Locations=append(Locations,&Location{Key:createKey(a[0],"Location")})
    case "Organization":
        Organizations=append(Organizations,&Organization{Key:createKey(a[0],"Organization")})
    case "Person":
        People=append(People,&Person{Key:createKey(a[0],"Person")})
    case "PhoneNumber":
        PhoneNumbers=append(PhoneNumbers,&PhoneNumber{Key:createKey(a[0],"PhoneNumber")})
    case "Portfolio":
        Portfolios=append(Portfolios,&Portfolio{Key:createKey(a[0],"Portfolio")})
    case "Position":
        Positions=append(Positions,&Position{Key:createKey(a[0],"Position")})
    case "Publication":
        Publications=append(Publications,&Publication{Key:createKey(a[0],"Publication")})
    case "Relationship":
        Relationships=append(Relationships,&Relationship{Key:createKey(a[0],"Relationship")})
    case "Resource":
        Resources=append(Resources,&Resource{Key:createKey(a[0],"Resource")})
    case "Teaching":
        Teachings=append(Teachings,&Teaching{Key:createKey(a[0],"Teaching")})
    case "URL":
        URLs=append(URLs,&URL{Key:createKey(a[0],"URL")})
    case "Venue":
        Venues=append(Venues,&Venue{Key:createKey(a[0],"Venue")})
	}
}

func AddEntityFact(a []string) {
	switch KeySet[a[0]] {
    case "Address":
        AddAddressFact(a)
    case "Award":
        AddAwardFact(a)
    case "Concept":
        AddConceptFact(a)
    case "Course":
        AddCourseFact(a)
    case "Date":
        AddDateFact(a)
    case "EmailAddress":
        AddEmailAddressFact(a)
    case "Event":
        AddEventFact(a)
    case "GeoLocation":
        AddGeoLocationFact(a)
    case "Grant":
        AddGrantFact(a)
    case "Image":
        AddImageFact(a)
    case "Location":
        AddLocationFact(a)
    case "Organization":
        AddOrganizationFact(a)
    case "Person":
        AddPersonFact(a)
    case "PhoneNumber":
        AddPhoneNumberFact(a)
    case "Portfolio":
        AddPortfolioFact(a)
    case "Position":
        AddPositionFact(a)
    case "Publication":
        AddPublicationFact(a)
    case "Relationship":
        AddRelationshipFact(a)
    case "Resource":
        AddResourceFact(a)
    case "Teaching":
        AddTeachingFact(a)
    case "URL":
        AddURLFact(a)
    case "Venue":
        AddVenueFact(a)
	}
}

func LoadEntities(n string) {
	EmptyLists()
	r3 := regexp.MustCompile(`\/(?P<Subj>[A-Z-a-z0-9]{6,6})> <.*has(?P<Pred>[A-Za-z]*)> (?P<Obj>.*)$`)
	r4 := regexp.MustCompile(`(?P<ObjURL>[A-Z-a-z0-9]{6,6})`)
	var t [][]string
	dat,err := os.ReadFile(n)
	if err != nil { log.Fatal(err) }
    lines := strings.Split(string(dat), " .\n")
    
    // First Pass -- extract keys and text
    
   	for i,v := range lines {
   		u := r3.FindStringSubmatch(v)
   		if len(u)== 4 {
			t = append(t,u[1:])
			if t[i][2][0:1] == `"` {
				t[i][2] = t[i][2][1:len(t[i][2])-1]
			} else {
				t[i][2] = r4.FindStringSubmatch(t[i][2])[0]
			}
		}
   	}
   	
   	// Second Pass -- create all entities and KeyMap
   		
   	for _,a := range t {
   		if a[1] == "Type" {
   			CreateEntity(a)
   		}
   	}
   	
   	// Third Pass -- add all facts to the appropriate entities using key values for linkage
   	
   	for _,a := range t {
   		if a[1] != "Type" {
   			AddEntityFact(a)
   		}
   	}
}


func WriteTriples(f *os.File, triples [][3]string) {
	for _,t := range triples {
		_, err := f.WriteString(t[0] + " " + t[1] + " " + t[2] + " .\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func WriteRow(f *os.File, row []string) {
	str := strings.Join(row,"\t")
	_, err := f.WriteString(str + "\n")
		if err != nil {
			log.Fatal(err)
		}
}

func SaveTables(n string) {
    f, err := os.Create(n)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    
    for _,p := range Addresses      { WriteRow(f, p.Row()) }
    for _,p := range Awards         { WriteRow(f, p.Row()) }
    for _,p := range Concepts       { WriteRow(f, p.Row()) }
    for _,p := range Courses        { WriteRow(f, p.Row()) }
    for _,p := range Dates          { WriteRow(f, p.Row()) }
    for _,p := range EmailAddresses { WriteRow(f, p.Row()) }
    for _,p := range Events         { WriteRow(f, p.Row()) }
    for _,p := range GeoLocations   { WriteRow(f, p.Row()) }
    for _,p := range Grants         { WriteRow(f, p.Row()) }
    for _,p := range Images         { WriteRow(f, p.Row()) }
    for _,p := range Locations      { WriteRow(f, p.Row()) }
//     for _,p := range Organizations  { WriteRow(f, p.Row()) }
//     for _,p := range People         { WriteRow(f, p.Row()) }
//     for _,p := range PhoneNumbers   { WriteRow(f, p.Row()) }
//     for _,p := range Portfolios     { WriteRow(f, p.Row()) }
//     for _,p := range Positions      { WriteRow(f, p.Row()) }
//     for _,p := range Publications   { WriteRow(f, p.Row()) }
//     for _,p := range Relationships  { WriteRow(f, p.Row()) }
//     for _,p := range Resources      { WriteRow(f, p.Row()) }
//     for _,p := range Teachings      { WriteRow(f, p.Row()) }
//     for _,p := range URLs           { WriteRow(f, p.Row()) }
//     for _,p := range Venues         { WriteRow(f, p.Row()) }    
}

func SaveEntities(n string) {
    f, err := os.Create(n)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    
    for _,p := range Addresses      { WriteTriples(f, p.Triples()) }
    for _,p := range Awards         { WriteTriples(f, p.Triples()) }
    for _,p := range Concepts       { WriteTriples(f, p.Triples()) }
    for _,p := range Courses        { WriteTriples(f, p.Triples()) }
    for _,p := range Dates          { WriteTriples(f, p.Triples()) }
    for _,p := range EmailAddresses { WriteTriples(f, p.Triples()) }
    for _,p := range Events         { WriteTriples(f, p.Triples()) }
    for _,p := range GeoLocations   { WriteTriples(f, p.Triples()) }
    for _,p := range Grants         { WriteTriples(f, p.Triples()) }
    for _,p := range Images         { WriteTriples(f, p.Triples()) }
    for _,p := range Locations      { WriteTriples(f, p.Triples()) }
    for _,p := range Organizations  { WriteTriples(f, p.Triples()) }
    for _,p := range People         { WriteTriples(f, p.Triples()) }
    for _,p := range PhoneNumbers   { WriteTriples(f, p.Triples()) }
    for _,p := range Portfolios     { WriteTriples(f, p.Triples()) }
    for _,p := range Positions      { WriteTriples(f, p.Triples()) }
    for _,p := range Publications   { WriteTriples(f, p.Triples()) }
    for _,p := range Relationships  { WriteTriples(f, p.Triples()) }
    for _,p := range Resources      { WriteTriples(f, p.Triples()) }
    for _,p := range Teachings      { WriteTriples(f, p.Triples()) }
    for _,p := range URLs           { WriteTriples(f, p.Triples()) }
    for _,p := range Venues         { WriteTriples(f, p.Triples()) }
}