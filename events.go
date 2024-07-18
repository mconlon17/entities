package entities

import (
	"errors"
	"fmt"
)
	
var _ = fmt.Println // remove after test

func makeEventTypeSet() map[string]struct{} {
	a := map[string]struct{}{}
	a["Conference"] = struct{}{}
	a["Symposium"] = struct{}{}
	a["Workshop"] = struct{}{}
	a["Fundraiser"] = struct{}{}
	a["Meeting"] = struct{}{}
	return a
}

var EventTypeSet = makeEventTypeSet()

type Event struct {
	Key *Key
	EventType string
	Name string
	Organizer *Organization
	Location *Location
	StartDate *Date
	EndDate *Date
	Sponsors []*Organization
	Attendees []*Person
}

func NewEvent(t string, n string) (*Event, error) {
	p := new(Event)
	p.Key = makeKey("Event")
	_,ok := EventTypeSet[t]
	if !ok {
		err := errors.New("Unknown Event type: " + t)
		return nil, err
	}
	p.EventType = t
	p.Name = n
	return p,nil
}

func NewConference(n string) (*Event,error) { return NewEvent("Conference", n) }
func NewSymposium(n string) (*Event,error) { return NewEvent("Symposium", n) }
func NewWorkshop(n string) (*Event,error) { return NewEvent("Workshop", n) }
func NewFundraiser(n string) (*Event,error) { return NewEvent("Fundraiser", n) }
func NewMeeting(n string) (*Event,error) { return NewEvent("Meeting", n) }

func (a *Event) Triples () [][3]string {
	var t [][3]string
	t = append(t, makeTriple(Triple{(*a).Key,"hasType","Event",nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasEventType",(*a).EventType,nil}))
	t = append(t, makeTriple(Triple{(*a).Key,"hasName",(*a).Name,nil}))
	if (*a).Organizer != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasOrganizer","",(*a.Organizer).Key}))}
	if (*a).Location != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasLocation","",(*a.Location).Key}))}
	if (*a).StartDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasStartDate","",(*a.StartDate).Key}))}
	if (*a).EndDate != nil {t = append(t, makeTriple(Triple{(*a).Key,"hasEndDate","",(*a.EndDate).Key}))}
	for _,cptr := range (*a).Sponsors {
		t = append(t, makeTriple(Triple{(*a).Key,"hasSponsor","",cptr.Key}))
	}
	for _,cptr := range (*a).Attendees {
		t = append(t, makeTriple(Triple{(*a).Key,"hasAttendee","",cptr.Key}))
	}
	return t
}

func FindEventKey (kf *Key) int {
	for i,a := range Events {
		if kf.s == a.Key.s {
			return i
		}
	}
	return -1
}

func AddEventFact(a []string) {
	key := new(Key)
	key.s = a[0]
	i := FindEventKey(key)
	switch a[1] {
	case "EventType":
		Events[i].EventType = a[2]
	case "Name":
		Events[i].Name = a[2]
	case "Organizer":
		key.s = a[2]
		j := FindOrganizationKey(key)
		Events[i].Organizer = Organizations[j]
	case "Location":
		key.s = a[2]
		j := FindLocationKey(key)
		Events[i].Location = Locations[j]
	case "StartDate":
		key.s = a[2]
		j := FindDateKey(key)
		Events[i].StartDate = Dates[j]
	case "EndDate":
		key.s = a[2]
		j := FindDateKey(key)
		Events[i].EndDate = Dates[j]
	case "Sponsor":
		key.s = a[2]
		j := FindOrganizationKey(key)
		Events[i].Sponsors = append(Events[i].Sponsors,Organizations[j])
	case "Attendee":
		key.s = a[2]
		j := FindPersonKey(key)
		Events[i].Attendees = append(Events[i].Attendees,People[j])	
	}
}

var Events []*Event