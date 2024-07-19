package entities

import (
	"fmt"
)

var _ = fmt.Println // remove after test	

func SampleData() {
	u1, _ := NewURL("Orcid","https://orcid.org/0000-0002-1304-8447")
	URLs =[]*URL{u1}
	
	course1,_ := NewCourse("AcademicCourse","STA5106", "Research Data Management")
	course1.CourseDescription = "Research projects collect data from which to create evidence regarding scientific hypotheses. This course will show you how."
	course2,_ := NewAcademicCourse("PMT", "Statistical Methods for Pharmacy")
	Courses = []*Course{course1,course2}
	
	c1 := NewConcept("Informatics")
	c2 := NewConcept("Biostatistics")
	c3 := NewConcept("Data Representation")
	c4 := NewConcept("Pharmacogenetics")
	c5 := NewConcept("Social Networks")
	c6 := NewConcept("Clinical and Translational Science")
	Concepts = []*Concept{c1,c2,c3,c4,c5,c6}
	
	p1,_ := NewPhoneNumber("352.214.7882")
	PhoneNumbers = []*PhoneNumber{p1}
	
	g1,_ := NewGeoLocation(29.65163000,-82.32483000)
	GeoLocations = []*GeoLocation{g1}
	
	i1, err := NewImage("images/uf-logo.jpg")
	if err != nil {
		fmt.Println(err)
		panic("Ending execution")
	}
	i1.Caption = "UF Logo"
	i1.AltText = "UF Logo"
	Images = []*Image{i1}

	l1,_ := NewCity("Gainesville")
	l1.GeoLocation = g1
	l2,_ := NewCity("Washington, D.C.")
	l3,_ := NewState("Florida")
	l3.Abbreviation = "FL"
	l3.Contains = []*Location{l1}
	Locations = []*Location{l1,l2, l3}
	
	a1, _ := NewAddress("6715 NW 67th Ave", l1, "32653")
	Addresses = []*Address{a1}
	
	e1,_ := NewEmailAddress("conlon.m@gmail.com")
	e2,_ := NewEmailAddress("mconlon@ufl.edu")
	e3,_ := NewEmailAddress("cpb@ufl.edu")
	EmailAddresses = []*EmailAddress{e1,e2,e3}
	
	d1, _ := NewDate("1953-06-08")
	d2, _ := NewYear("2004")
	d3, _ := NewYear("1983")
	d4, _ := NewYear("1997")
	d5, _ := NewYear("2014")
	d6, _ := NewYearMonth("1975-06")
	d7, _ := NewYearMonth("1979-05")
	d8, _ := NewYearMonth("1982-05")
	d9, _ := NewDate("2015-05-16")
	d10, _ := NewDate("1980-09-03")
	d11, _ := NewDateTime("2024-06-11 18:00:00")
	d12, _ := NewYear("2016")
	d13, _ := NewDate("1977-08-20")
	Dates = []*Date{d1,d2,d3,d4,d5,d6,d7,d8,d9,d10,d11,d12,d13}
	
	o1, _ := NewUniversity("University of Florida")
	course2.CourseOrganization = o1
	o2, _ := NewUniversity("Bucknell University")
	o3, _ := NewOrganization("Company", "Wiley and Sons")
	o4, _ := NewOrganization("Company", "Elsevier")
	o5, _ := NewOrganization("Department", "Oxford University Press")
	o6, _ := NewUniversity("Oxford University")
	o6.ChildOrganization = o5
	o5.ParentOrganization = o6
	o7, _ := NewOrganization("Governance", "Gainesville Opportunity Center Board of Directors")
	o8, _ := NewOrganization("Religious", "Holy Faith Catholic Church")
	o9, _ := NewOrganization("Professional", "American Statistical Association")
	Organizations = []*Organization{o1,o2,o3,o4,o5,o6,o7,o8,o9}
	
	v1 := NewJournal("New England Journal of Medicine")
	v2 := NewJournal("Journal of the American Medical Informatics Association")
	v2.Publisher = o5
	v3 := NewJournal("Clinical and Translational Science")
	v3.Publisher = o3
	Venues = []*Venue{v1,v2,v3}
	
	r1 := NewPortfolio()
	r2 := NewPortfolio()
	Portfolios = []*Portfolio{r1,r2}
	
	Publications = append(Publications, &Publication{Key: makeKey("Publication"), PublicationType: "AcademicArticle", Keywords: []*Concept{c1, c5}, PublicationDate: d2, Venue: v2, Title: "Direct2Experts: a pilot national network to demonstrate interoperability among research-networking platforms"})
	Publications = append(Publications, &Publication{Key: makeKey("Publication"), PublicationType: "AcademicArticle", Keywords: []*Concept{c4, c1}, PublicationDate: d3, Venue: v3, Title: "Clinical Pharmacogenetics implementation: approaches, successes, and challenges"})
	Publications = append(Publications, &Publication{Key: makeKey("Publication"), PublicationType: "AcademicArticle", Keywords: []*Concept{c4, c6}, PublicationDate: d4, Venue: v3, Title: "Institutional Profile: University of Florida and Shands Hospital Personalized Medicine Program: Clinical implementation of pharmacogenetics"})
	Publications = append(Publications, &Publication{Key: makeKey("Publication"), PublicationType: "AcademicArticle", Keywords: []*Concept{c5, c6}, PublicationDate: d5, Venue: v3, Title: "Designing a CTSA-based Social Network Intervention to Foster Cross-Disciplinary Team Science"})
	
	pub,_ := NewPhDThesis("M-estimation in the Linear Model")
	Publications = append(Publications, pub)
	
	
	People = append(People, &Person{Key: makeKey("Person"), Name: "Michael Conlon", Orcid: u1, HomePhone: p1, HomeAddress: a1, BirthDate: d1,
		HomeEmail: e1, WorkEmail: e2, Portfolio: r1, ResearchAreas: []*Concept{c1, c2}})
	People = append(People, &Person{Key: makeKey("Person"), Name: "Chris Barnes", Portfolio: r2, WorkEmail: e3, ResearchAreas: []*Concept{c1, c3}})
	p, _ := NewPerson("Deborah Conlon")
	People = append(People, p)

	
	award1,_ := NewAward("Bachelors", "BA in Mathematics")
	award1.Awardee = People[0]
	award1.AwardingOrganization = o2
	award1.AwardDate = d6
	award2,_ := NewAward("Bachelors", "BA in Economics")
	award2.Awardee = People[0]
	award2.AwardingOrganization = o2
	award2.AwardDate = d6
	award3,_ := NewAward("Masters", "Masters in Statistics")
	award3.Awardee = People[0]
	award3.AwardingOrganization = o1
	award3.AwardDate = d7
	award4,_ := NewAward("PhD", "PhD, Statistics")
	award4.Awardee = People[0]
	award4.AwardingOrganization = o1
	award4.AwardDate = d8
	Awards = []*Award{award1,award2,award3,award4}
	
	gr1,_ := NewResearchGrant("INVEST: The International Verapamil/Trandolapril Study")
	gr1.setAmount(37000000)
	gr2,_ := NewResearchGrant("Project Care: Cocaine Abuse in the Rural Environment")
	gr3,_ := NewResearchGrant("VIVO: Enabling National Networking of Scientists")
	gr3.setAmount(10800000)
	Grants = []*Grant{gr1,gr2,gr3}
	
	po1,_ := NewMember("Member")
	po1.Person = People[0]
	po1.Organization = o7
	
	po2,_ := NewMember("Member")
	po2.Person = People[0]
	po2.Organization = o8
	
	po3,_ := NewMember("Member")
	po2.Person = People[0]
	po2.Organization = o9
	
	po4,_ := NewFaculty("Emeritus Faculty Member")
	po4.Person = People[0]
	po4.Organization = o1
	po4.Description = "Dr. Conlon is an Emeritus Faculty Member of Health Outcomes Policy and Biomedical Informatics at the University of Florida. As such he is entitled to library privileges, an email address, and free parking."
	po4.StartDate = d9
	
	po5,_ := NewDirector("Director, Clinical and Translational Science Informatics and Technology (CTS-IT)")
	po5.Person = People[1]
	po5.Organization = o1
	po5.Description = "Chris leads a team of talented IT professionals who design, build, and operate IT solutions for research."
	po5.StartDate = d9
	
	po6,_ := NewRetired("")
	po6.Person = People[2]
	po6.Organization = o1
	po6.StartDate = d12

	
	Positions = []*Position{po1,po2,po3,po4,po5,po6}
	
	t1 := NewTeaching(course1, People[0])
	t1.StartDate = d10
	t1.Description = "My first course.  I had taught short courses, guest lectures, and served as a substitue previously. Students did well.  Test writing was nerve wracking. Test grading was rewarding."
	Teachings = []*Teaching{t1}
	
	rel1, _ := NewRelationship("Friend", People[0], People[1])
	rel2, _ := NewSpouse(People[0],People[2])
	rel2.StartDate = d13
	Relationships = []*Relationship{rel1,rel2}
	
	res1, _ := NewEquipment("HiperGator")
	res1.Admin = o1
	res1.Description = "HiperGator is a supercomputer at the University of Florida.  It has often been mentioned as the most powerful computer in academia. HiperGator uses Nvidia chips from a generous gift by Nvidia Corporation."
	res2, _ := NewService("IRB")
	res2.Admin = o1
	res2.Description = "The Institutional Review Board (IRB) is the human subject protection committee of the University of Florida.  All human subject research must be reviewed and approved by the IRB prior to enrollment of subjects in research."
	o1.Resources = append(o1.Resources,res1)
	o1.Resources = append(o1.Resources,res2)
	Resources = []*Resource{res1, res2}
	
	r1.Person = People[0]
	r1.Publications = []*Publication{Publications[0], Publications[1], Publications[2], Publications[3], Publications[4]}
	r1.Awards = []*Award{award1,award2,award3,award4}
	r1.Grants = []*Grant{gr1,gr2,gr3}
	r1.Positions = []*Position{po1,po2,po3,po4}
	r1.Teachings = []*Teaching{t1}
	
	r2.Person = People[1]
	r2.Positions = []*Position{po5}
	
	ev1,_ := NewConference("2nd Annual VIVO Conference")
	ev1.Location = l2
	ev1.Attendees = []*Person{People[0],People[1]}
	ev2,_ := NewMeeting("Board Meeting")
	ev2.Organizer = o7
	ev2.StartDate = d11
	Events = []*Event{ev1,ev2}
}